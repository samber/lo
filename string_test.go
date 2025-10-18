package lo

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := RandomString(100, LowerCaseLettersCharset)
	is.Equal(100, RuneLength(str1))
	is.Subset(LowerCaseLettersCharset, []rune(str1))

	str2 := RandomString(100, LowerCaseLettersCharset)
	is.NotEqual(str1, str2)

	noneUtf8Charset := []rune("明1好休2林森")
	str3 := RandomString(100, noneUtf8Charset)
	is.Equal(100, RuneLength(str3))
	is.Subset(noneUtf8Charset, []rune(str3))

	is.PanicsWithValue("lo.RandomString: charset must not be empty", func() { RandomString(100, []rune{}) })
	is.PanicsWithValue("lo.RandomString: size must be greater than 0", func() { RandomString(0, LowerCaseLettersCharset) })

	str4 := RandomString(10, []rune{65})
	is.Equal(10, RuneLength(str4))
	is.Subset([]rune{65, 65, 65, 65, 65, 65, 65, 65, 65, 65}, []rune(str4))
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

	is.PanicsWithValue("lo.ChunkString: size must be greater than 0", func() {
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
	str15 := Substring("hello", 5, 1)

	is.Empty(str1)
	is.Empty(str2)
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
	is.Empty(str15)
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

func TestEllipsis(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal("12345", Ellipsis("12345", 5))
	is.Equal("1...", Ellipsis("12345", 4))
	is.Equal("1...", Ellipsis("	12345  ", 4))
	is.Equal("12345", Ellipsis("12345", 6))
	is.Equal("12345", Ellipsis("12345", 10))
	is.Equal("12345", Ellipsis("  12345  ", 10))
	is.Equal("...", Ellipsis("12345", 3))
	is.Equal("...", Ellipsis("12345", 2))
	is.Equal("...", Ellipsis("12345", -1))
	is.Equal("hello...", Ellipsis(" hello   world ", 9))
}
