package benchmark

import (
	"strconv"
	"testing"

	"github.com/samber/lo"
)

func BenchmarkRange(b *testing.B) {
	for _, n := range lengths {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Range(n)
			}
		})
	}
}

func BenchmarkRangeFrom(b *testing.B) {
	for _, n := range lengths {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.RangeFrom(0, n)
			}
		})
	}
}

func BenchmarkRangeWithSteps(b *testing.B) {
	for _, n := range lengths {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.RangeWithSteps(0, n, 1)
			}
		})
	}
}

func BenchmarkClamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lo.Clamp(15, 0, 10)
	}
}

func BenchmarkSum(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Sum(ints)
			}
		})
	}
}

func BenchmarkSumBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.SumBy(ints, func(v int) int { return v })
			}
		})
	}
}

func BenchmarkProduct(b *testing.B) {
	for _, n := range lengths {
		floats := make([]float64, n)
		for j := range floats {
			floats[j] = 1.0001
		}
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Product(floats)
			}
		})
	}
}

func BenchmarkProductBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.ProductBy(ints, func(v int) float64 { return float64(v) * 0.001 })
			}
		})
	}
}

func BenchmarkMean(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Mean(ints)
			}
		})
	}
}

func BenchmarkMeanBy(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.MeanBy(ints, func(v int) int { return v })
			}
		})
	}
}

func BenchmarkMode(b *testing.B) {
	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lo.Mode(ints)
			}
		})
	}
}
