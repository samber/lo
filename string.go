package lo

import (
	"github.com/samber/lo/internal/rand"
	"math"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	LowerCaseLettersCharset = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset          = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset          = []rune("0123456789")
	AlphanumericCharset     = append(LettersCharset, NumbersCharset...)
	SpecialCharset          = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset              = append(AlphanumericCharset, SpecialCharset...)

	// bearer:disable go_lang_permissive_regex_validation
	splitWordReg = regexp.MustCompile(`([a-z])([A-Z0-9])|([a-zA-Z])([0-9])|([0-9])([a-zA-Z])|([A-Z])([A-Z])([a-z])`)
	// bearer:disable go_lang_permissive_regex_validation
	splitNumberLetterReg = regexp.MustCompile(`([0-9])([a-zA-Z])`)
	maximumCapacity      = math.MaxInt>>1 + 1
)

// RandomString return a random string.
// Play: https://go.dev/play/p/rRseOQVVum4
func RandomString(size int, charset []rune) string {
	if size <= 0 {
		panic("lo.RandomString: Size parameter must be greater than 0")
	}
	if len(charset) <= 0 {
		panic("lo.RandomString: Charset parameter must not be empty")
	}

	// see https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	sb := strings.Builder{}
	sb.Grow(size)
	// Calculate the number of bits required to represent the charset,
	// e.g., for 62 characters, it would need 6 bits (since 62 -> 64 = 2^6)
	letterIdBits := int(math.Log2(float64(nearestPowerOfTwo(len(charset)))))
	// Determine the corresponding bitmask,
	// e.g., for 62 characters, the bitmask would be 111111.
	var letterIdMask int64 = 1<<letterIdBits - 1
	// Available count, since rand.Int64() returns a non-negative number, the first bit is fixed, so there are 63 random bits
	// e.g., for 62 characters, this value is 10 (63 / 6).
	letterIdMax := 63 / letterIdBits
	// Generate the random string in a loop.
	for i, cache, remain := size-1, rand.Int64(), letterIdMax; i >= 0; {
		// Regenerate the random number if all available bits have been used
		if remain == 0 {
			cache, remain = rand.Int64(), letterIdMax
		}
		// Select a character from the charset
		if idx := int(cache & letterIdMask); idx < len(charset) {
			sb.WriteRune(charset[idx])
			i--
		}
		// Shift the bits to the right to prepare for the next character selection,
		// e.g., for 62 characters, shift by 6 bits.
		cache >>= letterIdBits
		// Decrease the remaining number of uses for the current random number.
		remain--
	}
	return sb.String()
}

// nearestPowerOfTwo returns the nearest power of two.
func nearestPowerOfTwo(cap int) int {
	n := cap - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 {
		return 1
	}
	if n >= maximumCapacity {
		return maximumCapacity
	}
	return n + 1
}

// Substring return part of a string.
// Play: https://go.dev/play/p/TQlxQi82Lu1
func Substring[T ~string](str T, offset int, length uint) T {
	rs := []rune(str)
	size := len(rs)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}

	if offset >= size {
		return Empty[T]()
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	return T(strings.Replace(string(rs[offset:offset+int(length)]), "\x00", "", -1))
}

// ChunkString returns an array of strings split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
// Play: https://go.dev/play/p/__FLTuJVz54
func ChunkString[T ~string](str T, size int) []T {
	if size <= 0 {
		panic("lo.ChunkString: Size parameter must be greater than 0")
	}

	if len(str) == 0 {
		return []T{""}
	}

	if size >= len(str) {
		return []T{str}
	}

	var chunks = make([]T, 0, ((len(str)-1)/size)+1)
	currentLen := 0
	currentStart := 0
	for i := range str {
		if currentLen == size {
			chunks = append(chunks, str[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, str[currentStart:])
	return chunks
}

// RuneLength is an alias to utf8.RuneCountInString which returns the number of runes in string.
// Play: https://go.dev/play/p/tuhgW_lWY8l
func RuneLength(str string) int {
	return utf8.RuneCountInString(str)
}

// PascalCase converts string to pascal case.
func PascalCase(str string) string {
	items := Words(str)
	for i := range items {
		items[i] = Capitalize(items[i])
	}
	return strings.Join(items, "")
}

// CamelCase converts string to camel case.
func CamelCase(str string) string {
	items := Words(str)
	for i, item := range items {
		item = strings.ToLower(item)
		if i > 0 {
			item = Capitalize(item)
		}
		items[i] = item
	}
	return strings.Join(items, "")
}

// KebabCase converts string to kebab case.
func KebabCase(str string) string {
	items := Words(str)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	return strings.Join(items, "-")
}

// SnakeCase converts string to snake case.
func SnakeCase(str string) string {
	items := Words(str)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	return strings.Join(items, "_")
}

// Words splits string into an array of its words.
func Words(str string) []string {
	str = splitWordReg.ReplaceAllString(str, `$1$3$5$7 $2$4$6$8$9`)
	// example: Int8Value => Int 8Value => Int 8 Value
	str = splitNumberLetterReg.ReplaceAllString(str, "$1 $2")
	var result strings.Builder
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
		} else {
			result.WriteRune(' ')
		}
	}
	return strings.Fields(result.String())
}

// Capitalize converts the first character of string to upper case and the remaining to lower case.
func Capitalize(str string) string {
	return cases.Title(language.English).String(str)
}

// Ellipsis trims and truncates a string to a specified length and appends an ellipsis if truncated.
func Ellipsis(str string, length int) string {
	str = strings.TrimSpace(str)

	if len(str) > length {
		if len(str) < 3 || length < 3 {
			return "..."
		}
		return strings.TrimSpace(str[0:length-3]) + "..."
	}

	return str
}

// Elipse trims and truncates a string to a specified length and appends an ellipsis if truncated.
//
// Deprecated: Use Ellipsis instead.
func Elipse(str string, length int) string {
	return Ellipsis(str, length)
}
