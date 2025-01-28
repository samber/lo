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
	str15 := Substring("hello", 5, 1)

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
	is.Equal("", str15)
}

func TestRuneLength(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(5, RuneLength("hellô"))
	is.Equal(6, len("hellô"))
}

func TestAllCase(t *testing.T) {
	type output struct {
		PascalCase string
		CamelCase  string
		KebabCase  string
		SnakeCase  string
	}
	name := ""
	tests := []struct {
		name   string
		input  string
		output output
	}{
		{name: name, output: output{}},
		{name: name, input: ".", output: output{}},
		{name: name, input: "Hello world!", output: output{
			PascalCase: "HelloWorld",
			CamelCase:  "helloWorld",
			KebabCase:  "hello-world",
			SnakeCase:  "hello_world",
		}},
		{name: name, input: "A", output: output{
			PascalCase: "A",
			CamelCase:  "a",
			KebabCase:  "a",
			SnakeCase:  "a",
		}},
		{name: name, input: "a", output: output{
			PascalCase: "A",
			CamelCase:  "a",
			KebabCase:  "a",
			SnakeCase:  "a",
		}},
		{name: name, input: "foo", output: output{
			PascalCase: "Foo",
			CamelCase:  "foo",
			KebabCase:  "foo",
			SnakeCase:  "foo",
		}},
		{name: name, input: "snake_case", output: output{
			PascalCase: "SnakeCase",
			CamelCase:  "snakeCase",
			KebabCase:  "snake-case",
			SnakeCase:  "snake_case",
		}},
		{name: name, input: "SNAKE_CASE", output: output{
			PascalCase: "SnakeCase",
			CamelCase:  "snakeCase",
			KebabCase:  "snake-case",
			SnakeCase:  "snake_case",
		}},
		{name: name, input: "kebab-case", output: output{
			PascalCase: "KebabCase",
			CamelCase:  "kebabCase",
			KebabCase:  "kebab-case",
			SnakeCase:  "kebab_case",
		}},
		{name: name, input: "PascalCase", output: output{
			PascalCase: "PascalCase",
			CamelCase:  "pascalCase",
			KebabCase:  "pascal-case",
			SnakeCase:  "pascal_case",
		}},
		{name: name, input: "camelCase", output: output{
			PascalCase: "CamelCase",
			CamelCase:  "camelCase",
			KebabCase:  `camel-case`,
			SnakeCase:  "camel_case",
		}},
		{name: name, input: "Title Case", output: output{
			PascalCase: "TitleCase",
			CamelCase:  "titleCase",
			KebabCase:  "title-case",
			SnakeCase:  "title_case",
		}},
		{name: name, input: "point.case", output: output{
			PascalCase: "PointCase",
			CamelCase:  "pointCase",
			KebabCase:  "point-case",
			SnakeCase:  "point_case",
		}},
		{name: name, input: "snake_case_with_more_words", output: output{
			PascalCase: "SnakeCaseWithMoreWords",
			CamelCase:  "snakeCaseWithMoreWords",
			KebabCase:  "snake-case-with-more-words",
			SnakeCase:  "snake_case_with_more_words",
		}},
		{name: name, input: "SNAKE_CASE_WITH_MORE_WORDS", output: output{
			PascalCase: "SnakeCaseWithMoreWords",
			CamelCase:  "snakeCaseWithMoreWords",
			KebabCase:  "snake-case-with-more-words",
			SnakeCase:  "snake_case_with_more_words",
		}},
		{name: name, input: "kebab-case-with-more-words", output: output{
			PascalCase: "KebabCaseWithMoreWords",
			CamelCase:  "kebabCaseWithMoreWords",
			KebabCase:  "kebab-case-with-more-words",
			SnakeCase:  "kebab_case_with_more_words",
		}},
		{name: name, input: "PascalCaseWithMoreWords", output: output{
			PascalCase: "PascalCaseWithMoreWords",
			CamelCase:  "pascalCaseWithMoreWords",
			KebabCase:  "pascal-case-with-more-words",
			SnakeCase:  "pascal_case_with_more_words",
		}},
		{name: name, input: "camelCaseWithMoreWords", output: output{
			PascalCase: "CamelCaseWithMoreWords",
			CamelCase:  "camelCaseWithMoreWords",
			KebabCase:  "camel-case-with-more-words",
			SnakeCase:  "camel_case_with_more_words",
		}},
		{name: name, input: "Title Case With More Words", output: output{
			PascalCase: "TitleCaseWithMoreWords",
			CamelCase:  "titleCaseWithMoreWords",
			KebabCase:  "title-case-with-more-words",
			SnakeCase:  "title_case_with_more_words",
		}},
		{name: name, input: "point.case.with.more.words", output: output{
			PascalCase: "PointCaseWithMoreWords",
			CamelCase:  "pointCaseWithMoreWords",
			KebabCase:  "point-case-with-more-words",
			SnakeCase:  "point_case_with_more_words",
		}},
		{name: name, input: "snake_case__with___multiple____delimiters", output: output{
			PascalCase: "SnakeCaseWithMultipleDelimiters",
			CamelCase:  "snakeCaseWithMultipleDelimiters",
			KebabCase:  "snake-case-with-multiple-delimiters",
			SnakeCase:  "snake_case_with_multiple_delimiters",
		}},
		{name: name, input: "SNAKE_CASE__WITH___multiple____DELIMITERS", output: output{
			PascalCase: "SnakeCaseWithMultipleDelimiters",
			CamelCase:  "snakeCaseWithMultipleDelimiters",
			KebabCase:  "snake-case-with-multiple-delimiters",
			SnakeCase:  "snake_case_with_multiple_delimiters",
		}},
		{name: name, input: "kebab-case--with---multiple----delimiters", output: output{
			PascalCase: "KebabCaseWithMultipleDelimiters",
			CamelCase:  "kebabCaseWithMultipleDelimiters",
			KebabCase:  "kebab-case-with-multiple-delimiters",
			SnakeCase:  "kebab_case_with_multiple_delimiters",
		}},
		{name: name, input: "Title Case  With   Multiple    Delimiters", output: output{
			PascalCase: "TitleCaseWithMultipleDelimiters",
			CamelCase:  "titleCaseWithMultipleDelimiters",
			KebabCase:  "title-case-with-multiple-delimiters",
			SnakeCase:  "title_case_with_multiple_delimiters",
		}},
		{name: name, input: "point.case..with...multiple....delimiters", output: output{
			PascalCase: "PointCaseWithMultipleDelimiters",
			CamelCase:  "pointCaseWithMultipleDelimiters",
			KebabCase:  "point-case-with-multiple-delimiters",
			SnakeCase:  "point_case_with_multiple_delimiters",
		}},
		{name: name, input: " leading space", output: output{
			PascalCase: "LeadingSpace",
			CamelCase:  "leadingSpace",
			KebabCase:  "leading-space",
			SnakeCase:  "leading_space",
		}},
		{name: name, input: "   leading spaces", output: output{
			PascalCase: "LeadingSpaces",
			CamelCase:  "leadingSpaces",
			KebabCase:  "leading-spaces",
			SnakeCase:  "leading_spaces",
		}},
		{name: name, input: "\t\t\r\n leading whitespaces", output: output{
			PascalCase: "LeadingWhitespaces",
			CamelCase:  "leadingWhitespaces",
			KebabCase:  "leading-whitespaces",
			SnakeCase:  "leading_whitespaces",
		}},
		{name: name, input: "trailing space ", output: output{
			PascalCase: "TrailingSpace",
			CamelCase:  "trailingSpace",
			KebabCase:  "trailing-space",
			SnakeCase:  "trailing_space",
		}},
		{name: name, input: "trailing spaces   ", output: output{
			PascalCase: "TrailingSpaces",
			CamelCase:  "trailingSpaces",
			KebabCase:  "trailing-spaces",
			SnakeCase:  "trailing_spaces",
		}},
		{name: name, input: "trailing whitespaces\t\t\r\n", output: output{
			PascalCase: "TrailingWhitespaces",
			CamelCase:  "trailingWhitespaces",
			KebabCase:  "trailing-whitespaces",
			SnakeCase:  "trailing_whitespaces",
		}},
		{name: name, input: " on both sides ", output: output{
			PascalCase: "OnBothSides",
			CamelCase:  "onBothSides",
			KebabCase:  "on-both-sides",
			SnakeCase:  "on_both_sides",
		}},
		{name: name, input: "    many on both sides  ", output: output{
			PascalCase: "ManyOnBothSides",
			CamelCase:  "manyOnBothSides",
			KebabCase:  "many-on-both-sides",
			SnakeCase:  "many_on_both_sides",
		}},
		{name: name, input: "\r whitespaces on both sides\t\t\r\n", output: output{
			PascalCase: "WhitespacesOnBothSides",
			CamelCase:  "whitespacesOnBothSides",
			KebabCase:  "whitespaces-on-both-sides",
			SnakeCase:  "whitespaces_on_both_sides",
		}},
		{name: name, input: "  extraSpaces in_This TestCase Of MIXED_CASES\t", output: output{
			PascalCase: "ExtraSpacesInThisTestCaseOfMixedCases",
			CamelCase:  "extraSpacesInThisTestCaseOfMixedCases",
			KebabCase:  "extra-spaces-in-this-test-case-of-mixed-cases",
			SnakeCase:  "extra_spaces_in_this_test_case_of_mixed_cases",
		}},
		{name: name, input: "CASEBreak", output: output{
			PascalCase: "CaseBreak",
			CamelCase:  "caseBreak",
			KebabCase:  "case-break",
			SnakeCase:  "case_break",
		}},
		{name: name, input: "ID", output: output{
			PascalCase: "Id",
			CamelCase:  "id",
			KebabCase:  "id",
			SnakeCase:  "id",
		}},
		{name: name, input: "userID", output: output{
			PascalCase: "UserId",
			CamelCase:  "userId",
			KebabCase:  "user-id",
			SnakeCase:  "user_id",
		}},
		{name: name, input: "JSON_blob", output: output{
			PascalCase: "JsonBlob",
			CamelCase:  "jsonBlob",
			KebabCase:  "json-blob",
			SnakeCase:  "json_blob",
		}},
		{name: name, input: "HTTPStatusCode", output: output{
			PascalCase: "HttpStatusCode",
			CamelCase:  "httpStatusCode",
			KebabCase:  "http-status-code",
			SnakeCase:  "http_status_code",
		}},
		{name: name, input: "FreeBSD and SSLError are not golang initialisms", output: output{
			PascalCase: "FreeBsdAndSslErrorAreNotGolangInitialisms",
			CamelCase:  "freeBsdAndSslErrorAreNotGolangInitialisms",
			KebabCase:  "free-bsd-and-ssl-error-are-not-golang-initialisms",
			SnakeCase:  "free_bsd_and_ssl_error_are_not_golang_initialisms",
		}},
		{name: name, input: "David's Computer", output: output{
			PascalCase: "DavidSComputer",
			CamelCase:  "davidSComputer",
			KebabCase:  "david-s-computer",
			SnakeCase:  "david_s_computer",
		}},
		{name: name, input: "http200", output: output{
			PascalCase: "Http200",
			CamelCase:  "http200",
			KebabCase:  "http-200",
			SnakeCase:  "http_200",
		}},
		{name: name, input: "NumberSplittingVersion1.0r3", output: output{
			PascalCase: "NumberSplittingVersion10R3",
			CamelCase:  "numberSplittingVersion10R3",
			KebabCase:  "number-splitting-version-1-0-r3",
			SnakeCase:  "number_splitting_version_1_0_r3",
		}},
		{name: name, input: "When you have a comma, odd results", output: output{
			PascalCase: "WhenYouHaveACommaOddResults",
			CamelCase:  "whenYouHaveACommaOddResults",
			KebabCase:  "when-you-have-a-comma-odd-results",
			SnakeCase:  "when_you_have_a_comma_odd_results",
		}},
		{name: name, input: "Ordinal numbers work: 1st 2nd and 3rd place", output: output{
			PascalCase: "OrdinalNumbersWork1St2NdAnd3RdPlace",
			CamelCase:  "ordinalNumbersWork1St2NdAnd3RdPlace",
			KebabCase:  "ordinal-numbers-work-1-st-2-nd-and-3-rd-place",
			SnakeCase:  "ordinal_numbers_work_1_st_2_nd_and_3_rd_place",
		}},
		{name: name, input: "BadUTF8\xe2\xe2\xa1", output: output{
			PascalCase: "BadUtf8",
			CamelCase:  "badUtf8",
			KebabCase:  "bad-utf-8",
			SnakeCase:  "bad_utf_8",
		}},
		{name: name, input: "IDENT3", output: output{
			PascalCase: "Ident3",
			CamelCase:  "ident3",
			KebabCase:  "ident-3",
			SnakeCase:  "ident_3",
		}},
		{name: name, input: "LogRouterS3BucketName", output: output{
			PascalCase: "LogRouterS3BucketName",
			CamelCase:  "logRouterS3BucketName",
			KebabCase:  "log-router-s3-bucket-name",
			SnakeCase:  "log_router_s3_bucket_name",
		}},
		{name: name, input: "PINEAPPLE", output: output{
			PascalCase: "Pineapple",
			CamelCase:  "pineapple",
			KebabCase:  "pineapple",
			SnakeCase:  "pineapple",
		}},
		{name: name, input: "Int8Value", output: output{
			PascalCase: "Int8Value",
			CamelCase:  "int8Value",
			KebabCase:  "int-8-value",
			SnakeCase:  "int_8_value",
		}},
		{name: name, input: "first.last", output: output{
			PascalCase: "FirstLast",
			CamelCase:  "firstLast",
			KebabCase:  "first-last",
			SnakeCase:  "first_last",
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pascal := PascalCase(test.input)
			if pascal != test.output.PascalCase {
				t.Errorf("PascalCase(%q) = %q; expected %q", test.input, pascal, test.output.PascalCase)
			}
			camel := CamelCase(test.input)
			if camel != test.output.CamelCase {
				t.Errorf("CamelCase(%q) = %q; expected %q", test.input, camel, test.output.CamelCase)
			}
			kebab := KebabCase(test.input)
			if kebab != test.output.KebabCase {
				t.Errorf("KebabCase(%q) = %q; expected %q", test.input, kebab, test.output.KebabCase)
			}
			snake := SnakeCase(test.input)
			if snake != test.output.SnakeCase {
				t.Errorf("SnakeCase(%q) = %q; expected %q", test.input, snake, test.output.SnakeCase)
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
		{"", args{"PascalCase"}, []string{"Pascal", "Case"}},
		{"", args{"camelCase"}, []string{"camel", "Case"}},
		{"", args{"snake_case"}, []string{"snake", "case"}},
		{"", args{"kebab_case"}, []string{"kebab", "case"}},
		{"", args{"_test text_"}, []string{"test", "text"}},
		{"", args{"UPPERCASE"}, []string{"UPPERCASE"}},
		{"", args{"HTTPCode"}, []string{"HTTP", "Code"}},
		{"", args{"Int8Value"}, []string{"Int", "8", "Value"}},
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
		{"", args{"heLLO"}, "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Capitalize(tt.args.word), "Capitalize(%v)", tt.args.word)
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
