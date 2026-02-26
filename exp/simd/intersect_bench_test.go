//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"testing"
)

// Benchmark suite for SIMD Contains operations compared to core lo package fallbacks.
// These benchmarks measure the performance of element lookup operations
// across different SIMD implementations (AVX, AVX2, AVX512) and data sizes.

// Benchmark sizes for Contains operations
var containsBenchmarkSizes = []struct {
	name string
	size int
}{
	{"tiny", 4},       // Smaller than AVX width (16 lanes for int8)
	{"small", 16},     // Exactly AVX width for int8
	{"medium", 64},    // Multiple of AVX, between AVX and AVX2 for int8
	{"large", 256},    // Multiple of AVX2 (32 lanes for int8)
	{"xlarge", 1024},  // Multiple of AVX512 (64 lanes for int8)
	{"massive", 8192}, // Very large dataset
}

// ============================================================================
// CONTAINS INT8 BENCHMARKS
// ============================================================================

func BenchmarkContainsInt8(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt8(bs.size)
			target := int8(42)

			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b) // ContainsInt8x16 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt8x16(data, target)
				}
			})
			b.Run("AVX512-x32", func(b *testing.B) {
				requireAVX512(b) // ContainsInt8x32 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt8x32(data, target)
				}
			})
			b.Run("AVX512-x64", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt8x64(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS INT16 BENCHMARKS
// ============================================================================

func BenchmarkContainsInt16(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt16(bs.size)
			target := int16(42)

			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b) // ContainsInt16x8 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt16x8(data, target)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b) // ContainsInt16x16 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt16x16(data, target)
				}
			})
			b.Run("AVX512-x32", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt16x32(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS INT32 BENCHMARKS
// ============================================================================

func BenchmarkContainsInt32(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt32(bs.size)
			target := int32(42)

			b.Run("AVX512-x4", func(b *testing.B) {
				requireAVX512(b) // ContainsInt32x4 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt32x4(data, target)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b) // ContainsInt32x8 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt32x8(data, target)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt32x16(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS INT64 BENCHMARKS
// ============================================================================

func BenchmarkContainsInt64(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt64(bs.size)
			target := int64(42)

			b.Run("AVX512-x2", func(b *testing.B) {
				requireAVX512(b) // ContainsInt64x2 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt64x2(data, target)
				}
			})
			b.Run("AVX512-x4", func(b *testing.B) {
				requireAVX512(b) // ContainsInt64x4 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt64x4(data, target)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt64x8(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS UINT8 BENCHMARKS
// ============================================================================

func BenchmarkContainsUint8(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateUint8(bs.size)
			target := uint8(255)

			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b) // ContainsUint8x16 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint8x16(data, target)
				}
			})
			b.Run("AVX512-x32", func(b *testing.B) {
				requireAVX512(b) // ContainsUint8x32 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint8x32(data, target)
				}
			})
			b.Run("AVX512-x64", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint8x64(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS UINT16 BENCHMARKS
// ============================================================================

func BenchmarkContainsUint16(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateUint16(bs.size)
			target := uint16(42)

			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b) // ContainsUint16x8 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint16x8(data, target)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b) // ContainsUint16x16 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint16x16(data, target)
				}
			})
			b.Run("AVX512-x32", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint16x32(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS UINT32 BENCHMARKS
// ============================================================================

func BenchmarkContainsUint32(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateUint32(bs.size)
			target := uint32(42)

			b.Run("AVX512-x4", func(b *testing.B) {
				requireAVX512(b) // ContainsUint32x4 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint32x4(data, target)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b) // ContainsUint32x8 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint32x8(data, target)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint32x16(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS UINT64 BENCHMARKS
// ============================================================================

func BenchmarkContainsUint64(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateUint64(bs.size)
			target := uint64(42)

			b.Run("AVX512-x2", func(b *testing.B) {
				requireAVX512(b) // ContainsUint64x2 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint64x2(data, target)
				}
			})
			b.Run("AVX512-x4", func(b *testing.B) {
				requireAVX512(b) // ContainsUint64x4 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint64x4(data, target)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsUint64x8(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS FLOAT32 BENCHMARKS
// ============================================================================

func BenchmarkContainsFloat32(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateFloat32(bs.size)
			target := float32(42.5)

			b.Run("AVX512-x4", func(b *testing.B) {
				requireAVX512(b) // ContainsFloat32x4 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsFloat32x4(data, target)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b) // ContainsFloat32x8 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsFloat32x8(data, target)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsFloat32x16(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS FLOAT64 BENCHMARKS
// ============================================================================

func BenchmarkContainsFloat64(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateFloat64(bs.size)
			target := float64(42.5)

			b.Run("AVX512-x2", func(b *testing.B) {
				requireAVX512(b) // ContainsFloat64x2 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsFloat64x2(data, target)
				}
			})
			b.Run("AVX512-x4", func(b *testing.B) {
				requireAVX512(b) // ContainsFloat64x4 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsFloat64x4(data, target)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsFloat64x8(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS WORST-CASE BENCHMARKS (target at end)
// ============================================================================

// These benchmarks test worst-case performance where target is at the very end
func BenchmarkContainsWorstCase(b *testing.B) {
	size := 1024
	data := make([]int32, size)
	for i := range data {
		data[i] = int32(i)
	}
	target := int32(size - 1) // Target at the very end

	b.Run("AVX512-x4", func(b *testing.B) {
		requireAVX512(b) // ContainsInt32x4 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt32x4(data, target)
		}
	})
	b.Run("AVX512-x8", func(b *testing.B) {
		requireAVX512(b) // ContainsInt32x8 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt32x8(data, target)
		}
	})
	b.Run("AVX512-x16", func(b *testing.B) {
		requireAVX512(b) // ContainsInt32x16 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt32x16(data, target)
		}
	})
}

// ============================================================================
// CONTAINS BEST-CASE BENCHMARKS (target at beginning)
// ============================================================================

// These benchmarks test best-case performance where target is at the beginning
func BenchmarkContainsBestCase(b *testing.B) {
	size := 1024
	data := make([]int32, size)
	for i := range data {
		data[i] = int32(i)
	}
	target := int32(0) // Target at the very beginning

	b.Run("AVX512-x4", func(b *testing.B) {
		requireAVX512(b) // ContainsInt32x4 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt32x4(data, target)
		}
	})
	b.Run("AVX512-x8", func(b *testing.B) {
		requireAVX512(b) // ContainsInt32x8 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt32x8(data, target)
		}
	})
	b.Run("AVX512-x16", func(b *testing.B) {
		requireAVX512(b) // ContainsInt32x16 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt32x16(data, target)
		}
	})
}

// ============================================================================
// CONTAINS NEGATIVE-CASE BENCHMARKS (target not present)
// ============================================================================

// These benchmarks test performance when target is not in the collection
func BenchmarkContainsNegative(b *testing.B) {
	for _, bs := range containsBenchmarkSizes {
		b.Run(bs.name, func(b *testing.B) {
			data := generateInt32(bs.size)
			target := int32(999999) // Target that's unlikely to be in the data

			b.Run("AVX512-x4", func(b *testing.B) {
				requireAVX512(b) // ContainsInt32x4 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt32x4(data, target)
				}
			})
			b.Run("AVX512-x8", func(b *testing.B) {
				requireAVX512(b) // ContainsInt32x8 is in intersect_avx512.go which uses AVX-512
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt32x8(data, target)
				}
			})
			b.Run("AVX512-x16", func(b *testing.B) {
				requireAVX512(b)
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = ContainsInt32x16(data, target)
				}
			})
		})
	}
}

// ============================================================================
// CONTAINS LANE WIDTH COMPARISON BENCHMARK
// ============================================================================

// This benchmark shows how performance scales with SIMD register width
func BenchmarkContainsInt8ByWidth(b *testing.B) {
	requireAVX512(b)

	size := 4096
	data := generateInt8(size)
	target := int8(42)

	benchmarks := []struct {
		name string
		fn   func() bool
	}{
		{"AVX512-x16", func() bool { return ContainsInt8x16(data, target) }},
		{"AVX512-x32", func() bool { return ContainsInt8x32(data, target) }},
		{"AVX512-x64", func() bool { return ContainsInt8x64(data, target) }},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = bm.fn()
			}
		})
	}
}

// ============================================================================
// STEADY STATE BENCHMARK
// ============================================================================

// This benchmark demonstrates the steady-state performance after warmup
func BenchmarkContainsInt64SteadyState(b *testing.B) {
	requireAVX512(b)

	size := 8192
	data := generateInt64(size)
	target := int64(42)

	// Warmup phase
	for i := 0; i < 1000; i++ {
		ContainsInt64x2(data, target)
		ContainsInt64x4(data, target)
		ContainsInt64x8(data, target)
	}

	b.ResetTimer() // Reset timer to exclude warmup

	b.Run("AVX512-x2", func(b *testing.B) {
		requireAVX512(b) // ContainsInt64x2 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt64x2(data, target)
		}
	})
	b.Run("AVX512-x4", func(b *testing.B) {
		requireAVX512(b) // ContainsInt64x4 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt64x4(data, target)
		}
	})
	b.Run("AVX512-x8", func(b *testing.B) {
		requireAVX512(b) // ContainsInt64x8 is in intersect_avx512.go which uses AVX-512
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ContainsInt64x8(data, target)
		}
	})
}
