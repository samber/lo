package lo

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	str1 := RandomString(100, LowerCaseLettersCharset)
	is.Equal(100, RuneLength(str1))
	is.Subset(LowerCaseLettersCharset, []rune(str1))

	str2 := RandomString(100, LowerCaseLettersCharset)
	is.NotEqual(str1, str2)

	noneUtf8Charset := []rune("明1好休2林森")
	str3 := RandomString(100, noneUtf8Charset)
	is.Equal(100, RuneLength(str3))
	is.Subset(noneUtf8Charset, []rune(str3))

	is.PanicsWithValue("lo.RandomString: Charset parameter must not be empty", func() { RandomString(100, []rune{}) })
	is.PanicsWithValue("lo.RandomString: Size parameter must be greater than 0", func() { RandomString(0, LowerCaseLettersCharset) })
}

func TestChunkString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ChunkString("12345", 2)
	is.Equal([]string{"12", "34", "5"}, result1)

	result2 := ChunkString("123456", 2)
	is.Equal([]string{"12", "34", "56"}, result2)

	result3 := ChunkString("123456", 6)
	is.Equal([]string{"123456"}, result3)

	result4 := ChunkString("123456", 10)
	is.Equal([]string{"123456"}, result4)

	result5 := ChunkString("", 2)
	is.Equal([]string{""}, result5)

	result6 := ChunkString("明1好休2林森", 2)
	is.Equal([]string{"明1", "好休", "2林", "森"}, result6)

	is.Panics(func() {
		ChunkString("12345", 0)
	})
}

func TestSubstring(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	is := assert.New(t)

	is.Equal(5, RuneLength("hellô"))
	is.Equal(6, len("hellô"))
}
