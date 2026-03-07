package benchmark

import (
	"testing"

	"github.com/samber/lo"
)

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.RandomString(64, lo.AlphanumericCharset)
	}
}

func BenchmarkSubstring(b *testing.B) {
	s := lo.RandomString(1000, lo.LettersCharset)
	for i := 0; i < b.N; i++ {
		_ = lo.Substring(s, 100, 200)
	}
}

func BenchmarkChunkString(b *testing.B) {
	s := lo.RandomString(1000, lo.LettersCharset)
	for i := 0; i < b.N; i++ {
		_ = lo.ChunkString(s, 10)
	}
}

func BenchmarkRuneLength(b *testing.B) {
	s := lo.RandomString(1000, lo.LettersCharset)
	for i := 0; i < b.N; i++ {
		_ = lo.RuneLength(s)
	}
}

func BenchmarkPascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.PascalCase("some_long_variable_name")
	}
}

func BenchmarkCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.CamelCase("some_long_variable_name")
	}
}

func BenchmarkKebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.KebabCase("someLongVariableName")
	}
}

func BenchmarkSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.SnakeCase("someLongVariableName")
	}
}

func BenchmarkWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Words("someLongVariableName")
	}
}

func BenchmarkCapitalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Capitalize("hello world")
	}
}

func BenchmarkEllipsis(b *testing.B) {
	s := lo.RandomString(200, lo.LettersCharset)
	for i := 0; i < b.N; i++ {
		_ = lo.Ellipsis(s, 50)
	}
}
