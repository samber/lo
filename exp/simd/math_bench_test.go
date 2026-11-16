//go:build goexperiment.simd

package simd

import (
	"testing"

	"github.com/samber/lo"
)

// benchmarkSizes spans below, around, and well above every vector width this package
// supports (128/256/512 bits), so the crossover point between the scalar fallback and the
// SIMD kernels is visible in the results.
var benchmarkSizes = []struct {
	name string
	size int
}{
	{"small", 8},
	{"medium", 128},
	{"large", 1024},
	{"xlarge", 8192},
}

func benchSum[T numeric](b *testing.B, wrapper, scalar func([]T) T) {
	b.Helper()
	for _, bs := range benchmarkSizes {
		data := makeData[T](bs.size, 99)
		b.Run(bs.name, func(b *testing.B) {
			b.Run("lo", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					_ = scalar(data)
				}
			})
			b.Run("simd", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					_ = wrapper(data)
				}
			})
		})
	}
}

func BenchmarkSumInt8(b *testing.B)    { benchSum(b, SumInt8[int8], lo.Sum[int8]) }
func BenchmarkSumInt16(b *testing.B)   { benchSum(b, SumInt16[int16], lo.Sum[int16]) }
func BenchmarkSumInt32(b *testing.B)   { benchSum(b, SumInt32[int32], lo.Sum[int32]) }
func BenchmarkSumInt64(b *testing.B)   { benchSum(b, SumInt64[int64], lo.Sum[int64]) }
func BenchmarkSumFloat32(b *testing.B) { benchSum(b, SumFloat32[float32], lo.Sum[float32]) }
func BenchmarkSumFloat64(b *testing.B) { benchSum(b, SumFloat64[float64], lo.Sum[float64]) }

func benchMinMax[T numeric](b *testing.B, minWrapper, maxWrapper, loMin, loMax func([]T) T) {
	b.Helper()
	for _, bs := range benchmarkSizes {
		data := makeData[T](bs.size, 98)
		b.Run(bs.name, func(b *testing.B) {
			b.Run("min/lo", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					_ = loMin(data)
				}
			})
			b.Run("min/simd", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					_ = minWrapper(data)
				}
			})
			b.Run("max/lo", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					_ = loMax(data)
				}
			})
			b.Run("max/simd", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					_ = maxWrapper(data)
				}
			})
		})
	}
}

func BenchmarkMinMaxInt32(b *testing.B) {
	benchMinMax(b, MinInt32[int32], MaxInt32[int32], lo.Min[int32], lo.Max[int32])
}

func BenchmarkMinMaxFloat64(b *testing.B) {
	benchMinMax(b, MinFloat64[float64], MaxFloat64[float64], lo.Min[float64], lo.Max[float64])
}

func benchClamp[T numeric](b *testing.B, wrapper func([]T, T, T) []T, scalar func(T, T, T) T, mn, mx T) {
	b.Helper()
	for _, bs := range benchmarkSizes {
		data := makeData[T](bs.size, 97)
		b.Run(bs.name, func(b *testing.B) {
			b.Run("lo", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					out := make([]T, len(data))
					for j, v := range data {
						out[j] = scalar(v, mn, mx)
					}
				}
			})
			b.Run("simd", func(b *testing.B) {
				b.ReportAllocs()
				for range b.N {
					_ = wrapper(data, mn, mx)
				}
			})
		})
	}
}

func BenchmarkClampInt32(b *testing.B) {
	benchClamp(b, ClampInt32[int32, []int32], lo.Clamp[int32], -500, 500)
}

func BenchmarkClampFloat64(b *testing.B) {
	benchClamp(b, ClampFloat64[float64, []float64], lo.Clamp[float64], -500, 500)
}
