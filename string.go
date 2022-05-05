package lo

import (
	"strings"
	"unicode/utf8"
)

// Substring return part of a string.
func Substring[T ~string](str T, offset int, length uint) T {
	size := len(str)

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

	return str[offset : offset+int(length)]
}

// RuneLength is an alias to utf8.RuneCountInString which returns the number of runes in string.
func RuneLength(str string) int {
	return utf8.RuneCountInString(str)
}

// Chop returns str with the last character removed.
// If str ends with \r\n, both characters are removed.
// If str is empty, Chop returns an empty string.
func Chop(str string) string {
	if len(str) < 2 {
		return ""
	}

	if strings.HasSuffix(str, "\r\n") {
		return str[:len(str)-2]
	}

	return str[:len(str)-1]
}
