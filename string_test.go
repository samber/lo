package lo

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestRandomString(t *testing.T) {
	t.Parallel()

	t.Run("length and charset", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name    string
			size    int
			charset []rune
		}{
			{name: "lowercase letters", size: 100, charset: LowerCaseLettersCharset},
			{name: "non-utf8 charset", size: 100, charset: []rune("明1好休2林森")},
			{name: "single-rune charset", size: 10, charset: []rune{65}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				str := RandomString(tt.size, tt.charset)
				is.Equal(tt.size, RuneLength(str))
				is.Subset(tt.charset, []rune(str))
			})
		}
	})

	t.Run("distinct calls produce distinct strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		str1 := RandomString(100, LowerCaseLettersCharset)
		str2 := RandomString(100, LowerCaseLettersCharset)
		is.NotEqual(str1, str2)
	})

	t.Run("panics", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.PanicsWithValue("lo.RandomString: charset must not be empty", func() { RandomString(100, []rune{}) })
		is.PanicsWithValue("lo.RandomString: size must be greater than 0", func() { RandomString(0, LowerCaseLettersCharset) })
	})
}

func TestChunkString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		size     int
		expected []string
	}{
		{name: "even split", str: "12345", size: 2, expected: []string{"12", "34", "5"}},
		{name: "exact multiple", str: "123456", size: 2, expected: []string{"12", "34", "56"}},
		{name: "size equals length", str: "123456", size: 6, expected: []string{"123456"}},
		{name: "size greater than length", str: "123456", size: 10, expected: []string{"123456"}},
		{name: "empty string", str: "", size: 2, expected: []string{""}}, // @TODO: should be [] - see https://github.com/samber/lo/issues/788
		{name: "unicode string", str: "明1好休2林森", size: 2, expected: []string{"明1", "好休", "2林", "森"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, ChunkString(tt.str, tt.size))
		})
	}

	t.Run("panics on zero size", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.PanicsWithValue("lo.ChunkString: size must be greater than 0", func() {
			ChunkString("12345", 0)
		})
	})
}

func TestSubstring(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		offset   int
		length   uint
		expected string
	}{
		{name: "offset beyond length, large requested length", str: "hello", offset: 5, length: 10, expected: ""},
		{name: "zero length", str: "hello", offset: 0, length: 0, expected: ""},
		{name: "offset beyond length", str: "hello", offset: 10, length: 2, expected: ""},
		{name: "negative offset within bounds", str: "hello", offset: -10, length: 2, expected: "he"},
		{name: "length exceeds string", str: "hello", offset: 0, length: 10, expected: "hello"},
		{name: "simple prefix", str: "hello", offset: 0, length: 2, expected: "he"},
		{name: "middle slice", str: "hello", offset: 2, length: 2, expected: "ll"},
		{name: "middle to end, exact length", str: "hello", offset: 2, length: 5, expected: "llo"},
		{name: "middle, length 3", str: "hello", offset: 2, length: 3, expected: "llo"},
		{name: "middle, length 4 clipped", str: "hello", offset: 2, length: 4, expected: "llo"},
		{name: "negative offset, length 4", str: "hello", offset: -2, length: 4, expected: "lo"},
		{name: "negative offset, length 1", str: "hello", offset: -4, length: 1, expected: "e"},
		{name: "negative offset, max length", str: "hello", offset: -4, length: math.MaxUint, expected: "ello"},
		{name: "emoji prefix", str: "🏠🐶🐱", offset: 0, length: 2, expected: "🏠🐶"},
		{name: "cjk prefix", str: "你好，世界", offset: 0, length: 3, expected: "你好，"},
		{name: "emoji offset 1", str: "🏠🐶🐱", offset: 1, length: 2, expected: "🐶🐱"},
		{name: "emoji negative offset", str: "🏠🐶🐱", offset: -2, length: 2, expected: "🐶🐱"},
		{name: "emoji offset at rune count", str: "🏠🐶🐱", offset: 3, length: 3, expected: ""},
		{name: "emoji offset beyond rune count", str: "🏠🐶🐱", offset: 4, length: 3, expected: ""},
		{name: "offset at length, length 1", str: "hello", offset: 5, length: 1, expected: ""},
		{name: "negative offset, full length", str: "hello", offset: -5, length: 5, expected: "hello"},
		{name: "negative offset, length 4 of full string", str: "hello", offset: -5, length: 4, expected: "hell"},
		{name: "negative offset, max length of full string", str: "hello", offset: -5, length: math.MaxUint, expected: "hello"},
		{name: "null bytes", str: "\x00\x00\x00", offset: 0, length: math.MaxUint, expected: ""},
		{name: "utf8 rune error", str: string(utf8.RuneError), offset: 0, length: math.MaxUint, expected: "�"},
		{name: "invalid utf8 from byte offset", str: "привет"[1:], offset: 0, length: 6, expected: "�ривет"},
		{name: "invalid utf8 truncated tail", str: "привет"[:2*5+1], offset: 0, length: 6, expected: "приве�"},
		{name: "invalid utf8 negative offset", str: "привет"[:2*5+1], offset: -2, length: math.MaxUint, expected: "е�"},
		{name: "invalid utf8 emoji tail, max length", str: "🏠🐶🐱"[1:], offset: 0, length: math.MaxUint, expected: "���🐶🐱"},
		{name: "invalid utf8 emoji tail, length 2", str: "🏠🐶🐱"[1:], offset: 0, length: 2, expected: "��"},
		{name: "cyrillic offset at rune count", str: "привет", offset: 6, length: math.MaxUint, expected: ""},
		{name: "cyrillic offset beyond rune count", str: "привет", offset: 6 + 1, length: math.MaxUint, expected: ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, Substring(tt.str, tt.offset, tt.length))
		})
	}
}

func BenchmarkSubstring(b *testing.B) {
	str := strings.Repeat("1", 100)

	for _, test := range []struct {
		offset int
		length uint
	}{
		{10, 10},
		{50, 50},
		{50, 45},
		{-50, 50},
		{-10, 10},
	} {
		fmt.Println(test)
		b.Run(fmt.Sprint(test), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Substring(str, test.offset, test.length)
			}
		})
	}
}

func TestRuneLength(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(5, RuneLength("hellô"))
	is.Len("hellô", 6)
}

func TestAllCase(t *testing.T) {
	t.Parallel()

	type output struct {
		PascalCase string
		CamelCase  string
		KebabCase  string
		SnakeCase  string
	}
	testCases := []struct {
		name string
		in   string
		want output
	}{
		{want: output{}},
		{in: ".", want: output{}},
		{in: "Hello world!", want: output{
			PascalCase: "HelloWorld",
			CamelCase:  "helloWorld",
			KebabCase:  "hello-world",
			SnakeCase:  "hello_world",
		}},
		{in: "A", want: output{
			PascalCase: "A",
			CamelCase:  "a",
			KebabCase:  "a",
			SnakeCase:  "a",
		}},
		{in: "a", want: output{
			PascalCase: "A",
			CamelCase:  "a",
			KebabCase:  "a",
			SnakeCase:  "a",
		}},
		{in: "foo", want: output{
			PascalCase: "Foo",
			CamelCase:  "foo",
			KebabCase:  "foo",
			SnakeCase:  "foo",
		}},
		{in: "snake_case", want: output{
			PascalCase: "SnakeCase",
			CamelCase:  "snakeCase",
			KebabCase:  "snake-case",
			SnakeCase:  "snake_case",
		}},
		{in: "SNAKE_CASE", want: output{
			PascalCase: "SnakeCase",
			CamelCase:  "snakeCase",
			KebabCase:  "snake-case",
			SnakeCase:  "snake_case",
		}},
		{in: "kebab-case", want: output{
			PascalCase: "KebabCase",
			CamelCase:  "kebabCase",
			KebabCase:  "kebab-case",
			SnakeCase:  "kebab_case",
		}},
		{in: "PascalCase", want: output{
			PascalCase: "PascalCase",
			CamelCase:  "pascalCase",
			KebabCase:  "pascal-case",
			SnakeCase:  "pascal_case",
		}},
		{in: "camelCase", want: output{
			PascalCase: "CamelCase",
			CamelCase:  "camelCase",
			KebabCase:  `camel-case`,
			SnakeCase:  "camel_case",
		}},
		{in: "Title Case", want: output{
			PascalCase: "TitleCase",
			CamelCase:  "titleCase",
			KebabCase:  "title-case",
			SnakeCase:  "title_case",
		}},
		{in: "point.case", want: output{
			PascalCase: "PointCase",
			CamelCase:  "pointCase",
			KebabCase:  "point-case",
			SnakeCase:  "point_case",
		}},
		{in: "snake_case_with_more_words", want: output{
			PascalCase: "SnakeCaseWithMoreWords",
			CamelCase:  "snakeCaseWithMoreWords",
			KebabCase:  "snake-case-with-more-words",
			SnakeCase:  "snake_case_with_more_words",
		}},
		{in: "SNAKE_CASE_WITH_MORE_WORDS", want: output{
			PascalCase: "SnakeCaseWithMoreWords",
			CamelCase:  "snakeCaseWithMoreWords",
			KebabCase:  "snake-case-with-more-words",
			SnakeCase:  "snake_case_with_more_words",
		}},
		{in: "kebab-case-with-more-words", want: output{
			PascalCase: "KebabCaseWithMoreWords",
			CamelCase:  "kebabCaseWithMoreWords",
			KebabCase:  "kebab-case-with-more-words",
			SnakeCase:  "kebab_case_with_more_words",
		}},
		{in: "PascalCaseWithMoreWords", want: output{
			PascalCase: "PascalCaseWithMoreWords",
			CamelCase:  "pascalCaseWithMoreWords",
			KebabCase:  "pascal-case-with-more-words",
			SnakeCase:  "pascal_case_with_more_words",
		}},
		{in: "camelCaseWithMoreWords", want: output{
			PascalCase: "CamelCaseWithMoreWords",
			CamelCase:  "camelCaseWithMoreWords",
			KebabCase:  "camel-case-with-more-words",
			SnakeCase:  "camel_case_with_more_words",
		}},
		{in: "Title Case With More Words", want: output{
			PascalCase: "TitleCaseWithMoreWords",
			CamelCase:  "titleCaseWithMoreWords",
			KebabCase:  "title-case-with-more-words",
			SnakeCase:  "title_case_with_more_words",
		}},
		{in: "point.case.with.more.words", want: output{
			PascalCase: "PointCaseWithMoreWords",
			CamelCase:  "pointCaseWithMoreWords",
			KebabCase:  "point-case-with-more-words",
			SnakeCase:  "point_case_with_more_words",
		}},
		{in: "snake_case__with___multiple____delimiters", want: output{
			PascalCase: "SnakeCaseWithMultipleDelimiters",
			CamelCase:  "snakeCaseWithMultipleDelimiters",
			KebabCase:  "snake-case-with-multiple-delimiters",
			SnakeCase:  "snake_case_with_multiple_delimiters",
		}},
		{in: "SNAKE_CASE__WITH___multiple____DELIMITERS", want: output{
			PascalCase: "SnakeCaseWithMultipleDelimiters",
			CamelCase:  "snakeCaseWithMultipleDelimiters",
			KebabCase:  "snake-case-with-multiple-delimiters",
			SnakeCase:  "snake_case_with_multiple_delimiters",
		}},
		{in: "kebab-case--with---multiple----delimiters", want: output{
			PascalCase: "KebabCaseWithMultipleDelimiters",
			CamelCase:  "kebabCaseWithMultipleDelimiters",
			KebabCase:  "kebab-case-with-multiple-delimiters",
			SnakeCase:  "kebab_case_with_multiple_delimiters",
		}},
		{in: "Title Case  With   Multiple    Delimiters", want: output{
			PascalCase: "TitleCaseWithMultipleDelimiters",
			CamelCase:  "titleCaseWithMultipleDelimiters",
			KebabCase:  "title-case-with-multiple-delimiters",
			SnakeCase:  "title_case_with_multiple_delimiters",
		}},
		{in: "point.case..with...multiple....delimiters", want: output{
			PascalCase: "PointCaseWithMultipleDelimiters",
			CamelCase:  "pointCaseWithMultipleDelimiters",
			KebabCase:  "point-case-with-multiple-delimiters",
			SnakeCase:  "point_case_with_multiple_delimiters",
		}},
		{in: " leading space", want: output{
			PascalCase: "LeadingSpace",
			CamelCase:  "leadingSpace",
			KebabCase:  "leading-space",
			SnakeCase:  "leading_space",
		}},
		{in: "   leading spaces", want: output{
			PascalCase: "LeadingSpaces",
			CamelCase:  "leadingSpaces",
			KebabCase:  "leading-spaces",
			SnakeCase:  "leading_spaces",
		}},
		{in: "\t\t\r\n leading whitespaces", want: output{
			PascalCase: "LeadingWhitespaces",
			CamelCase:  "leadingWhitespaces",
			KebabCase:  "leading-whitespaces",
			SnakeCase:  "leading_whitespaces",
		}},
		{in: "trailing space ", want: output{
			PascalCase: "TrailingSpace",
			CamelCase:  "trailingSpace",
			KebabCase:  "trailing-space",
			SnakeCase:  "trailing_space",
		}},
		{in: "trailing spaces   ", want: output{
			PascalCase: "TrailingSpaces",
			CamelCase:  "trailingSpaces",
			KebabCase:  "trailing-spaces",
			SnakeCase:  "trailing_spaces",
		}},
		{in: "trailing whitespaces\t\t\r\n", want: output{
			PascalCase: "TrailingWhitespaces",
			CamelCase:  "trailingWhitespaces",
			KebabCase:  "trailing-whitespaces",
			SnakeCase:  "trailing_whitespaces",
		}},
		{in: " on both sides ", want: output{
			PascalCase: "OnBothSides",
			CamelCase:  "onBothSides",
			KebabCase:  "on-both-sides",
			SnakeCase:  "on_both_sides",
		}},
		{in: "    many on both sides  ", want: output{
			PascalCase: "ManyOnBothSides",
			CamelCase:  "manyOnBothSides",
			KebabCase:  "many-on-both-sides",
			SnakeCase:  "many_on_both_sides",
		}},
		{in: "\r whitespaces on both sides\t\t\r\n", want: output{
			PascalCase: "WhitespacesOnBothSides",
			CamelCase:  "whitespacesOnBothSides",
			KebabCase:  "whitespaces-on-both-sides",
			SnakeCase:  "whitespaces_on_both_sides",
		}},
		{in: "  extraSpaces in_This TestCase Of MIXED_CASES\t", want: output{
			PascalCase: "ExtraSpacesInThisTestCaseOfMixedCases",
			CamelCase:  "extraSpacesInThisTestCaseOfMixedCases",
			KebabCase:  "extra-spaces-in-this-test-case-of-mixed-cases",
			SnakeCase:  "extra_spaces_in_this_test_case_of_mixed_cases",
		}},
		{in: "CASEBreak", want: output{
			PascalCase: "CaseBreak",
			CamelCase:  "caseBreak",
			KebabCase:  "case-break",
			SnakeCase:  "case_break",
		}},
		{in: "ID", want: output{
			PascalCase: "Id",
			CamelCase:  "id",
			KebabCase:  "id",
			SnakeCase:  "id",
		}},
		{in: "userID", want: output{
			PascalCase: "UserId",
			CamelCase:  "userId",
			KebabCase:  "user-id",
			SnakeCase:  "user_id",
		}},
		{in: "JSON_blob", want: output{
			PascalCase: "JsonBlob",
			CamelCase:  "jsonBlob",
			KebabCase:  "json-blob",
			SnakeCase:  "json_blob",
		}},
		{in: "HTTPStatusCode", want: output{
			PascalCase: "HttpStatusCode",
			CamelCase:  "httpStatusCode",
			KebabCase:  "http-status-code",
			SnakeCase:  "http_status_code",
		}},
		{in: "FreeBSD and SSLError are not golang initialisms", want: output{
			PascalCase: "FreeBsdAndSslErrorAreNotGolangInitialisms",
			CamelCase:  "freeBsdAndSslErrorAreNotGolangInitialisms",
			KebabCase:  "free-bsd-and-ssl-error-are-not-golang-initialisms",
			SnakeCase:  "free_bsd_and_ssl_error_are_not_golang_initialisms",
		}},
		{in: "David's Computer", want: output{
			PascalCase: "DavidSComputer",
			CamelCase:  "davidSComputer",
			KebabCase:  "david-s-computer",
			SnakeCase:  "david_s_computer",
		}},
		{in: "http200", want: output{
			PascalCase: "Http200",
			CamelCase:  "http200",
			KebabCase:  "http-200",
			SnakeCase:  "http_200",
		}},
		{in: "NumberSplittingVersion1.0r3", want: output{
			PascalCase: "NumberSplittingVersion10R3",
			CamelCase:  "numberSplittingVersion10R3",
			KebabCase:  "number-splitting-version-1-0-r3",
			SnakeCase:  "number_splitting_version_1_0_r3",
		}},
		{in: "When you have a comma, odd results", want: output{
			PascalCase: "WhenYouHaveACommaOddResults",
			CamelCase:  "whenYouHaveACommaOddResults",
			KebabCase:  "when-you-have-a-comma-odd-results",
			SnakeCase:  "when_you_have_a_comma_odd_results",
		}},
		{in: "Ordinal numbers work: 1st 2nd and 3rd place", want: output{
			PascalCase: "OrdinalNumbersWork1St2NdAnd3RdPlace",
			CamelCase:  "ordinalNumbersWork1St2NdAnd3RdPlace",
			KebabCase:  "ordinal-numbers-work-1-st-2-nd-and-3-rd-place",
			SnakeCase:  "ordinal_numbers_work_1_st_2_nd_and_3_rd_place",
		}},
		{in: "BadUTF8\xe2\xe2\xa1", want: output{
			PascalCase: "BadUtf8",
			CamelCase:  "badUtf8",
			KebabCase:  "bad-utf-8",
			SnakeCase:  "bad_utf_8",
		}},
		{in: "IDENT3", want: output{
			PascalCase: "Ident3",
			CamelCase:  "ident3",
			KebabCase:  "ident-3",
			SnakeCase:  "ident_3",
		}},
		{in: "LogRouterS3BucketName", want: output{
			PascalCase: "LogRouterS3BucketName",
			CamelCase:  "logRouterS3BucketName",
			KebabCase:  "log-router-s3-bucket-name",
			SnakeCase:  "log_router_s3_bucket_name",
		}},
		{in: "PINEAPPLE", want: output{
			PascalCase: "Pineapple",
			CamelCase:  "pineapple",
			KebabCase:  "pineapple",
			SnakeCase:  "pineapple",
		}},
		{in: "Int8Value", want: output{
			PascalCase: "Int8Value",
			CamelCase:  "int8Value",
			KebabCase:  "int-8-value",
			SnakeCase:  "int_8_value",
		}},
		{in: "first.last", want: output{
			PascalCase: "FirstLast",
			CamelCase:  "firstLast",
			KebabCase:  "first-last",
			SnakeCase:  "first_last",
		}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			is.Equalf(tc.want.PascalCase, PascalCase(tc.in), "PascalCase(%v)", tc.in)
			is.Equalf(tc.want.CamelCase, CamelCase(tc.in), "CamelCase(%v)", tc.in)
			is.Equalf(tc.want.KebabCase, KebabCase(tc.in), "KebabCase(%v)", tc.in)
			is.Equalf(tc.want.SnakeCase, SnakeCase(tc.in), "SnakeCase(%v)", tc.in)
		})
	}
}

func TestWords(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in   string
		want []string
	}{
		{"PascalCase", []string{"Pascal", "Case"}},
		{"camelCase", []string{"camel", "Case"}},
		{"snake_case", []string{"snake", "case"}},
		{"kebab_case", []string{"kebab", "case"}},
		{"_test text_", []string{"test", "text"}},
		{"UPPERCASE", []string{"UPPERCASE"}},
		{"HTTPCode", []string{"HTTP", "Code"}},
		{"Int8Value", []string{"Int", "8", "Value"}},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tc.want, Words(tc.in), "Words(%v)", tc.in)
		})
	}
}

func TestCapitalize(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		in   string
		want string
	}{
		{"lower case", "hello", "Hello"},
		{"mixed case", "heLLO", "Hello"},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tc.want, Capitalize(tc.in), "Capitalize(%v)", tc.in)
		})
	}
}

func TestCapitalizeWithLanguage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		lang language.Tag
		want string
	}{
		{name: "english plain I", in: "istanbul", lang: language.English, want: "Istanbul"},
		{name: "english mixed case", in: "heLLO", lang: language.English, want: "Hello"},
		// Turkish: lowercase i → title İ (dotted capital I, U+0130)
		{name: "turkish lowercase i", in: "istanbul", lang: language.Turkish, want: "İstanbul"},
		// Turkish: ISTANBUL starts with capital I (the uppercase form of ı); title case
		// keeps it as I and lowercases the rest → "Istanbul" (not "İstanbul").
		{name: "turkish uppercase ISTANBUL", in: "ISTANBUL", lang: language.Turkish, want: "Istanbul"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, CapitalizeWithLanguage(tt.in, tt.lang))
		})
	}
}

func TestPascalCaseWithLanguage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		lang language.Tag
		want string
	}{
		{name: "english", in: "istanbul city", lang: language.English, want: "IstanbulCity"},
		// Turkish: lowercase i → title İ (dotted capital I, U+0130)
		{name: "turkish lowercase i", in: "istanbul city", lang: language.Turkish, want: "İstanbulCity"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, PascalCaseWithLanguage(tt.in, tt.lang))
		})
	}
}

func TestCamelCaseWithLanguage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		lang language.Tag
		want string
	}{
		{name: "english", in: "istanbul city", lang: language.English, want: "istanbulCity"},
		// Turkish: capital I lowercases to ı (dotless i, U+0131), so first word → "ıstanbul".
		// Second word title-cased: C stays C, capital I → lowercase ı → "Cıty".
		{name: "turkish uppercase ISTANBUL CITY", in: "ISTANBUL CITY", lang: language.Turkish, want: "ıstanbulCıty"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, CamelCaseWithLanguage(tt.in, tt.lang))
		})
	}
}

func TestKebabCaseWithLanguage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		lang language.Tag
		want string
	}{
		{name: "english", in: "ISTANBUL CITY", lang: language.English, want: "istanbul-city"},
		// Turkish: I lowercases to ı (dotless i, U+0131)
		{name: "turkish", in: "ISTANBUL CITY", lang: language.Turkish, want: "ıstanbul-cıty"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, KebabCaseWithLanguage(tt.in, tt.lang))
		})
	}
}

func TestSnakeCaseWithLanguage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		lang language.Tag
		want string
	}{
		{name: "english", in: "ISTANBUL CITY", lang: language.English, want: "istanbul_city"},
		// Turkish: I lowercases to ı (dotless i, U+0131)
		{name: "turkish", in: "ISTANBUL CITY", lang: language.Turkish, want: "ıstanbul_cıty"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, SnakeCaseWithLanguage(tt.in, tt.lang))
		})
	}
}

func TestEllipsis(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		str      string
		length   int
		expected string
	}{
		{name: "length 0", str: "12", length: 0, expected: "..."},
		{name: "length 1", str: "12", length: 1, expected: "..."},
		{name: "length equals string length", str: "12", length: 2, expected: "12"},
		{name: "length exceeds string length", str: "12", length: 3, expected: "12"},
		{name: "exact length", str: "12345", length: 5, expected: "12345"},
		{name: "truncate to 1 char", str: "12345", length: 4, expected: "1..."},
		{name: "leading tab trimmed before truncation", str: "\t12345  ", length: 4, expected: "1..."},
		{name: "length exceeds string length by 1", str: "12345", length: 6, expected: "12345"},
		{name: "length well beyond string length", str: "12345", length: 10, expected: "12345"},
		{name: "surrounding whitespace trimmed", str: "  12345  ", length: 10, expected: "12345"},
		{name: "length 3 on 5-char string", str: "12345", length: 3, expected: "..."},
		{name: "length 2 on 5-char string", str: "12345", length: 2, expected: "..."},
		{name: "negative length", str: "12345", length: -1, expected: "..."},
		{name: "internal whitespace collapsed and truncated", str: " hello   world ", length: 9, expected: "hello..."},

		// Unicode: rune-based truncation (not byte-based)
		{name: "CJK characters: hello (5 runes) + ... = 8 runes", str: "hello 世界! 你好", length: 8, expected: "hello..."},
		{name: "truncate within CJK text", str: "hello 世界! 你好", length: 11, expected: "hello 世界..."},
		{name: "CJK exact length, no truncation", str: "hello 世界! 你好", length: 12, expected: "hello 世界! 你好"},
		{name: "CJK length exceeds string, no truncation", str: "hello 世界! 你好", length: 20, expected: "hello 世界! 你好"},
		{name: "length > rune count, no truncation", str: "🏠🐶🐱🌟", length: 5, expected: "🏠🐶🐱🌟"},
		{name: "emoji exact length, no truncation", str: "🏠🐶🐱🌟", length: 4, expected: "🏠🐶🐱🌟"},
		{name: "emoji length == 3, returns ...", str: "🏠🐶🐱🌟", length: 3, expected: "..."},
		{name: "emoji length < 3, returns ...", str: "🏠🐶🐱🌟", length: 2, expected: "..."},
		{name: "6 emoji, truncate to 2 + ...", str: "🏠🐶🐱🌟🎉🌈", length: 5, expected: "🏠🐶..."},
		{name: "accented char counts as 1 rune", str: "café", length: 4, expected: "café"},
		{name: "accented, length == 3, returns ...", str: "café", length: 3, expected: "..."},
		{name: "mixed ASCII and accented", str: "café au lait", length: 5, expected: "ca..."},

		// Combining emoji (Rainbow Flag is 4 runes: U+1F3F3 + U+FE0F + U+200D + U+1F308)
		// "aà😁🏳️‍🌈pabc" = 1 + 1 + 1 + 4 + 1 + 1 + 1 + 1 = 11 runes total
		{name: "combining emoji, length 2: only ...", str: "aà😁🏳️‍🌈pabc", length: 2, expected: "..."},
		{name: "combining emoji, length 3: only ...", str: "aà😁🏳️‍🌈pabc", length: 3, expected: "..."},
		{name: "combining emoji, 1 rune + ...", str: "aà😁🏳️‍🌈pabc", length: 4, expected: "a..."},
		{name: "combining emoji, 2 runes + ...", str: "aà😁🏳️‍🌈pabc", length: 5, expected: "aà..."},
		{name: "combining emoji, 3 runes + ...", str: "aà😁🏳️‍🌈pabc", length: 6, expected: "aà😁..."},
		// @TODO: fix these cases
		// {name: "4 runes + ...", str: "aà😁🏳️‍🌈pabc", length: 7, expected: "aà😁🏳️‍🌈..."},
		// {name: "5 runes + ...", str: "aà😁🏳️‍🌈pabc", length: 8, expected: "aà😁🏳️‍🌈p..."},
		// {name: "exact length, no truncation", str: "aà😁🏳️‍🌈pabc", length: 9, expected: "aà😁🏳️‍🌈pabc"},
		// {name: "length exceeds string, no truncation", str: "aà😁🏳️‍🌈pabc", length: 10, expected: "aà😁🏳️‍🌈pabc"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, Ellipsis(tt.str, tt.length))
		})
	}
}
