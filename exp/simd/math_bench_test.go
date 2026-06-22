//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"math/rand"
	"simd/archsimd"
	"testing"
	"time"

	"github.com/samber/lo"
)

// Benchmark suite for SIMD math operations compared to core lo package fallbacks.
// These benchmarks measure the performance of Sum, Mean, Min, and Max operations
// across different SIMD implementations (AVX, AVX2, AVX512) and data sizes.

// Benchmark sizes to demonstrate performance characteristics at different scales
var benchmarkSizes = []struct {
	name string
	size int
}{
	{"small", 8},     // Smaller than AVX width (16 lanes for int8)
	{"medium", 128},  // Between AVX (16) and AVX2 (32) width for int8
	{"large", 1024},  // Well above SIMD register widths
	{"xlarge", 8192}, // Large dataset for real-world performance
}

func init() {
	// Seeded for reproducibility
	rand.Seed(time.Now().UnixNano())
}

// Helper function to generate random test data
type benchDataGenerator[T any] func(n int) []T

func generateInt8(n int) []int8 {
	data := make([]int8, n)
	for i := range data {
		data[i] = int8(rand.Intn(127) - 64)
	}
	return data
}

func generateInt16(n int) []int16 {
	data := make([]int16, n)
	for i := range data {
		data[i] = int16(rand.Intn(32767) - 16384)
	}
	return data
}

func generateInt32(n int) []int32 {
	data := make([]int32, n)
	for i := range data {
		data[i] = int32(rand.Intn(1000) - 500)
	}
	return data
}

func generateInt64(n int) []int64 {
	data := make([]int64, n)
	for i := range data {
		data[i] = rand.Int63() % 10000
	}
	return data
}

func generateUint8(n int) []uint8 {
	data := make([]uint8, n)
	for i := range data {
		data[i] = uint8(rand.Uint32() % 256)
	}
	return data
}

func generateUint16(n int) []uint16 {
	data := make([]uint16, n)
	for i := range data {
		data[i] = uint16(rand.Uint32() % 65536)
	}
	return data
}

func generateUint32(n int) []uint32 {
	data := make([]uint32, n)
	for i := range data {
		data[i] = rand.Uint32() % 10000
	}
	return data
}

func generateUint64(n int) []uint64 {
	data := make([]uint64, n)
	for i := range data {
		data[i] = rand.Uint64() % 10000
	}
	return data
}

func generateFloat32(n int) []float32 {
	data := make([]float32, n)
	for i := range data {
		data[i] = rand.Float32()*100 - 50
	}
	return data
}

func generateFloat64(n int) []float64 {
	data := make([]float64, n)
	for i := range data {
		data[i] = rand.Float64()*100 - 50
	}
	return data
}

// ============================================================================
// SUM BENCHMARKS
// ============================================================================

func BenchmarkSumInt8(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt8(bs.size)
			b.Run("Fallback-lo", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = lo.Sum(data)
				}
			})
			b.Run("AVX-x16", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt8x16(data)
				}
			})
			b.Run("AVX2-x32", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt8x32(data)
				}
			})
			b.Run("AVX512-x64", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt8x64(data)
				}
			})
		})
	}
}

func BenchmarkSumInt16(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt16(bs.size)
			b.Run("Fallback-lo", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = lo.Sum(data)
				}
			})
			b.Run("AVX-x8", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt16x8(data)
				}
			})
			b.Run("AVX2-x16", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt16x16(data)
				}
			})
			b.Run("AVX512-x32", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt16x32(data)
				}
			})
		})
	}
}

func BenchmarkSumInt32(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt32(bs.size)
			b.Run("Fallback-lo", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = lo.Sum(data)
				}
			})
			b.Run("AVX-x4", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt32x4(data)
				}
			})
			b.Run("AVX2-x8", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt32x8(data)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt32x16(data)
				}
			})
		})
	}
}

func BenchmarkSumInt64(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt64(bs.size)
			b.Run("Fallback-lo", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = lo.Sum(data)
				}
			})
			b.Run("AVX-x2", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt64x2(data)
				}
			})
			b.Run("AVX2-x4", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt64x4(data)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumInt64x8(data)
				}
			})
		})
	}
}

func BenchmarkSumFloat32(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateFloat32(bs.size)
			b.Run("Fallback-lo", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = lo.Sum(data)
				}
			})
			b.Run("AVX-x4", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumFloat32x4(data)
				}
			})
			b.Run("AVX2-x8", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumFloat32x8(data)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumFloat32x16(data)
				}
			})
		})
	}
}

func BenchmarkSumFloat64(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateFloat64(bs.size)
			b.Run("Fallback-lo", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = lo.Sum(data)
				}
			})
			b.Run("AVX-x2", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumFloat64x2(data)
				}
			})
			b.Run("AVX2-x4", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumFloat64x4(data)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = SumFloat64x8(data)
				}
			})
		})
	}
}

// ============================================================================
// MEAN BENCHMARKS
// ============================================================================

func BenchmarkMeanInt32(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt32(bs.size)
			b.Run("Fallback-lo", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = lo.Mean(data)
				}
			})
			b.Run("AVX-x4", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MeanInt32x4(data)
				}
			})
			b.Run("AVX2-x8", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MeanInt32x8(data)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MeanInt32x16(data)
				}
			})
		})
	}
}

func BenchmarkMeanFloat64(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateFloat64(bs.size)
			b.Run("Fallback-lo", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = lo.Mean(data)
				}
			})
			b.Run("AVX-x2", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MeanFloat64x2(data)
				}
			})
			b.Run("AVX2-x4", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MeanFloat64x4(data)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MeanFloat64x8(data)
				}
			})
		})
	}
}

// ============================================================================
// MIN BENCHMARKS
// ============================================================================

func BenchmarkMinInt32(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt32(bs.size)
			b.Run("AVX-x4", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MinInt32x4(data)
				}
			})
			b.Run("AVX2-x8", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MinInt32x8(data)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MinInt32x16(data)
				}
			})
		})
	}
}

func BenchmarkMinFloat64(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateFloat64(bs.size)
			b.Run("AVX-x2", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MinFloat64x2(data)
				}
			})
			b.Run("AVX2-x4", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MinFloat64x4(data)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MinFloat64x8(data)
				}
			})
		})
	}
}

// ============================================================================
// MAX BENCHMARKS
// ============================================================================

func BenchmarkMaxInt32(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt32(bs.size)
			b.Run("AVX-x4", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MaxInt32x4(data)
				}
			})
			b.Run("AVX2-x8", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MaxInt32x8(data)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MaxInt32x16(data)
				}
			})
		})
	}
}

func BenchmarkMaxFloat64(b *testing.B) {
	for _, bs := range benchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateFloat64(bs.size)
			b.Run("AVX-x2", func(b *testing.B) {
				requireAVX(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MaxFloat64x2(data)
				}
			})
			b.Run("AVX2-x4", func(b *testing.B) {
				requireAVX2(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MaxFloat64x4(data)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = MaxFloat64x8(data)
				}
			})
		})
	}
}

// ============================================================================
// LANE WIDTH COMPARISON BENCHMARKS
// ============================================================================

// These benchmarks show how performance scales with SIMD register width
func BenchmarkSumInt8ByWidth(b *testing.B) {
	size := 4096 // Large enough to see differences across implementations
	data := generateInt8(size)

	benchmarks := []struct {
		name string
		fn   func() int8
	}{
		{"Fallback-lo", func() int8 { return lo.Sum(data) }},
		{"AVX-x16", func() int8 { return SumInt8x16(data) }},
		{"AVX2-x32", func() int8 { return SumInt8x32(data) }},
		{"AVX512-x64", func() int8 { return SumInt8x64(data) }},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			if bm.name == "AVX-x16" {
				requireAVX(b)
			}
			if bm.name == "AVX2-x32" {
				requireAVX2(b)
			}
			if bm.name == "AVX512-x64" {
				requireAVX512(b)
			}
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = bm.fn()
			}
		})
	}
}

// ============================================================================
// COMPARATIVE BENCHMARK WITH WARMUP
// ============================================================================

// This benchmark demonstrates the steady-state performance after warmup
func BenchmarkSumInt64SteadyState(b *testing.B) {
	size := 8192
	data := generateInt64(size)

	// Warmup phase to ensure JIT compilation if applicable
	for i := 0; i < 1000; i++ {
		lo.Sum(data)
		SumInt64x2(data)
		if archsimd.X86.AVX2() {
			SumInt64x4(data)
		}
		if archsimd.X86.AVX512() {
			SumInt64x8(data)
		}
	}

	b.ResetTimer() // Reset timer to exclude warmup

	b.Run("Fallback-lo", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = lo.Sum(data)
		}
	})
	b.Run("AVX-x2", func(b *testing.B) {
		requireAVX(b)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = SumInt64x2(data)
		}
	})
	b.Run("AVX2-x4", func(b *testing.B) {
		requireAVX2(b)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = SumInt64x4(data)
		}
	})
	b.Run("AVX512-x8", func(b *testing.B) {
		requireAVX512(b)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = SumInt64x8(data)
		}
	})
}
