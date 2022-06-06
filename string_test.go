package lo

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstring(t *testing.T) {
	is := assert.New(t)

	str1 := Substring("hello", 0, 0)
	str2 := Substring("hello", 10, 2)
	str3 := Substring("hello", -10, 2)
	str4 := Substring("hello", 0, 10)
	str5 := Substring("hello", 0, 2)
	str6 := Substring("hello", 2, 2)
	str7 := Substring("hello", 2, 5)
	str8 := Substring("hello", 2, 3)
	str9 := Substring("hello", 2, 4)
	str10 := Substring("hello", -2, 4)
	str11 := Substring("hello", -4, 1)
	str12 := Substring("hello", -4, math.MaxUint)

	is.Equal("", str1)
	is.Equal("", str2)
	is.Equal("he", str3)
	is.Equal("hello", str4)
	is.Equal("he", str5)
	is.Equal("ll", str6)
	is.Equal("llo", str7)
	is.Equal("llo", str8)
	is.Equal("llo", str9)
	is.Equal("lo", str10)
	is.Equal("e", str11)
	is.Equal("ello", str12)
}

func TestRuneLength(t *testing.T) {
	is := assert.New(t)

	is.Equal(5, RuneLength("hellô"))
	is.Equal(6, len("hellô"))
}
