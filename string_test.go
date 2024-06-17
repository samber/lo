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
	str13 := Substring("🏠🐶🐱", 0, 2)
	str14 := Substring("你好，世界", 0, 3)

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

	is.Equal(5, RuneLength("hellô"))
	is.Equal(6, len("hellô"))
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"", "hello_world", "HelloWorld"},
		{"", "helloWorld", "HelloWorld"},
		{"", "__hello_world-example string--", "HelloWorldExampleString"},
		{"", "WITH UPPERCASE LETTERS", "WithUppercaseLetters"},
		{"", "test123_string", "Test123String"},
		{"", "test123string", "Test123String"},
		{"", "", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := PascalCase(test.input)
			if actual != test.expected {
				t.Errorf("PascalCase(%q) = %q; expected %q", test.input, actual, test.expected)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"", "hello_world", "helloWorld"},
		{"", "helloWorld", "helloWorld"},
		{"", "__hello_world-example string--", "helloWorldExampleString"},
		{"", "WITH UPPERCASE LETTERS", "withUppercaseLetters"},
		{"", "test123_string", "test123String"},
		{"", "test123string", "test123String"},
		{"", "", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CamelCase(test.input)
			if result != test.expected {
				t.Errorf("CamelCase(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"", "hello world", "hello-world"},
		{"", "HelloWorld", "hello-world"},
		{"", "KebabCase", "kebab-case"},
		{"", "already-kebab-case", "already-kebab-case"},
		{"", "Already-Kebab-Case", "already-kebab-case"},
		{"", "multiple   spaces", "multiple-spaces"},
		{"", "", ""},
		{"", "Single", "single"},
		{"", "123_abs", "123-abs"},
		{"", "SINGLE", "single"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := KebabCase(test.input)
			if result != test.expected {
				t.Errorf("KebabCase(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"", "CamelCase", "camel_case"},
		{"", "snakeCase", "snake_case"},
		{"", "snake-case", "snake_case"},
		{"", "SnakeCaseTest", "snake_case_test"},
		{"", "Snake_Case_With_Underscores", "snake_case_with_underscores"},
		{"", "lowercase", "lowercase"},
		{"", "UPPERCASE", "uppercase"},
		{"", "", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := SnakeCase(test.input)
			if got != test.expected {
				t.Errorf("SnakeCase(%q) = %q; want %q", test.input, got, test.expected)
			}
		})
	}
}

func TestWords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{"CamelCase"}, []string{"Camel", "Case"}},
		{"", args{"snakeCase"}, []string{"snake", "Case"}},
		{"", args{"snake-case"}, []string{"snake", "case"}},
		{"", args{"test123string"}, []string{"test", "123", "string"}},
		{"", args{"UPPERCASE"}, []string{"UPPERCASE"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Words(tt.args.str), "words(%v)", tt.args.str)
		})
	}
}

func TestCapitalize(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"hello"}, "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Capitalize(tt.args.word), "Capitalize(%v)", tt.args.word)
		})
	}
}
