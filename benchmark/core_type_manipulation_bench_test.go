package benchmark

import (
	"strconv"
	"testing"

	"github.com/samber/lo"
)

func BenchmarkToPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.ToPtr(42)
	}
}

func BenchmarkFromPtr(b *testing.B) {
	p := lo.ToPtr(42)
	for i := 0; i < b.N; i++ {
		_ = lo.FromPtr(p)
	}
}

func BenchmarkFromPtrOr(b *testing.B) {
	var p *int
	for i := 0; i < b.N; i++ {
		_ = lo.FromPtrOr(p, 99)
	}
}

func BenchmarkToSlicePtr(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ToSlicePtr(ints)
			}
		})
	}
}

func BenchmarkToAnySlice(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ToAnySlice(ints)
			}
		})
	}
}

func BenchmarkFromAnySlice(b *testing.B) {
	for _, n := range lengths {
		anys := lo.ToAnySlice(genSliceInt(n))
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = lo.FromAnySlice[int](anys)
			}
		})
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.IsEmpty(0)
	}
}

func BenchmarkIsNotEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.IsNotEmpty(42)
	}
}

func BenchmarkCoalesce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = lo.Coalesce(0, 0, 0, 42, 99)
	}
}
