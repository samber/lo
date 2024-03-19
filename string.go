package lo

import (
	"math/rand"
	"strings"
	"unicode/utf8"
)

var (
	LowerCaseLettersCharset = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset          = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset          = []rune("0123456789")
	AlphanumericCharset     = append(LettersCharset, NumbersCharset...)
	SpecialCharset          = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset              = append(AlphanumericCharset, SpecialCharset...)
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

	b := make([]rune, size)
	possibleCharactersCount := len(charset)
	for i := range b {
		b[i] = charset[rand.Intn(possibleCharactersCount)]
	}
	return string(b)
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

	if offset > size {
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

	var chunks []T = make([]T, 0, ((len(str)-1)/size)+1)
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
