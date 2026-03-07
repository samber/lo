package benchmark

import (
	"fmt"
	"testing"

	"github.com/samber/lo/mutable"
)

func BenchmarkMutableFilter(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				cp := make([]int, len(src))
				copy(cp, src)
				_ = mutable.Filter(cp, func(x int) bool { return x%2 == 0 })
			}
		})
	}
}

func BenchmarkMutableFilterI(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				cp := make([]int, len(src))
				copy(cp, src)
				_ = mutable.FilterI(cp, func(x, _ int) bool { return x%2 == 0 })
			}
		})
	}
}

func BenchmarkMutableMap(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				cp := make([]int, len(src))
				copy(cp, src)
				mutable.Map(cp, func(x int) int { return x * 2 })
			}
		})
	}
}

func BenchmarkMutableMapI(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				cp := make([]int, len(src))
				copy(cp, src)
				mutable.MapI(cp, func(x, i int) int { return x * i })
			}
		})
	}
}

func BenchmarkMutableShuffle(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				cp := make([]int, len(src))
				copy(cp, src)
				mutable.Shuffle(cp)
			}
		})
	}
}

func BenchmarkMutableReverse(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				cp := make([]int, len(src))
				copy(cp, src)
				mutable.Reverse(cp)
			}
		})
	}
}
