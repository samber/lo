package benchmark

import (
	"fmt"
	"testing"

	lop "github.com/samber/lo/parallel"
)

func BenchmarkParallelMap(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				_ = lop.Map(src, func(x, _ int) int { return x * 2 })
			}
		})
	}
}

func BenchmarkParallelForEach(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				lop.ForEach(src, func(_, _ int) {})
			}
		})
	}
}

func BenchmarkParallelTimes(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lop.Times(n, func(i int) int { return i * 2 })
			}
		})
	}
}

func BenchmarkParallelGroupBy(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				_ = lop.GroupBy(src, func(x int) int { return x % 10 })
			}
		})
	}
}

func BenchmarkParallelPartitionBy(b *testing.B) {
	for _, n := range lengths {
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			src := genSliceInt(n)
			for i := 0; i < b.N; i++ {
				_ = lop.PartitionBy(src, func(x int) int { return x % 10 })
			}
		})
	}
}
