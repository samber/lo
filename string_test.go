package lo_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	str1 := lo.RandomString(100, lo.LowerCaseLettersCharset)
	is.Equal(100, lo.RuneLength(str1))
	is.Subset(lo.LowerCaseLettersCharset, []rune(str1))

	str2 := lo.RandomString(100, lo.LowerCaseLettersCharset)
	is.NotEqual(str1, str2)

	noneUtf8Charset := []rune("明1好休2林森")
	str3 := lo.RandomString(100, noneUtf8Charset)
	is.Equal(100, lo.RuneLength(str3))
	is.Subset(noneUtf8Charset, []rune(str3))

	is.PanicsWithValue("lo.RandomString: Charset parameter must not be empty", func() { lo.RandomString(100, []rune{}) })
	is.PanicsWithValue("lo.RandomString: Size parameter must be greater than 0", func() { lo.RandomString(0, lo.LowerCaseLettersCharset) })
}

func TestChunkString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.ChunkString("12345", 2)
	is.Equal([]string{"12", "34", "5"}, result1)

	result2 := lo.ChunkString("123456", 2)
	is.Equal([]string{"12", "34", "56"}, result2)

	result3 := lo.ChunkString("123456", 6)
	is.Equal([]string{"123456"}, result3)

	result4 := lo.ChunkString("123456", 10)
	is.Equal([]string{"123456"}, result4)

	result5 := lo.ChunkString("", 2)
	is.Equal([]string{""}, result5)

	result6 := lo.ChunkString("明1好休2林森", 2)
	is.Equal([]string{"明1", "好休", "2林", "森"}, result6)

	is.Panics(func() {
		lo.ChunkString("12345", 0)
	})
}

func TestSubstring(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := lo.Substring("hello", 0, 0)
	str2 := lo.Substring("hello", 10, 2)
	str3 := lo.Substring("hello", -10, 2)
	str4 := lo.Substring("hello", 0, 10)
	str5 := lo.Substring("hello", 0, 2)
	str6 := lo.Substring("hello", 2, 2)
	str7 := lo.Substring("hello", 2, 5)
	str8 := lo.Substring("hello", 2, 3)
	str9 := lo.Substring("hello", 2, 4)
	str10 := lo.Substring("hello", -2, 4)
	str11 := lo.Substring("hello", -4, 1)
	str12 := lo.Substring("hello", -4, math.MaxUint)
	str13 := lo.Substring("🏠🐶🐱", 0, 2)
	str14 := lo.Substring("你好，世界", 0, 3)

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
	is.Equal("🏠🐶", str13)
	is.Equal("你好，", str14)
}

func TestRuneLength(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(5, lo.RuneLength("hellô"))
	is.Equal(6, len("hellô"))
}
