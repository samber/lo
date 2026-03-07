package benchmark

import (
	"testing"

	"github.com/samber/lo"
)

func BenchmarkIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.If(i%2 == 0, "even").Else("odd")
	}
}

func BenchmarkIfElseIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.If(i%3 == 0, "fizz").
			ElseIf(i%3 == 1, "buzz").
			Else("none")
	}
}

func BenchmarkIfElseIfChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.If(i%5 == 0, "a").
			ElseIf(i%5 == 1, "b").
			ElseIf(i%5 == 2, "c").
			ElseIf(i%5 == 3, "d").
			Else("e")
	}
}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Switch[int, string](i%3).
			Case(0, "zero").
			Case(1, "one").
			Default("other")
	}
}

func BenchmarkSwitchChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Switch[int, string](i%5).
			Case(0, "zero").
			Case(1, "one").
			Case(2, "two").
			Case(3, "three").
			Default("other")
	}
}
