//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"github.com/samber/lo"
)

// SumInt8 sums a slice of int8 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt8[T ~int8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumInt8x64(collection)
	case simdFeatureAVX2:
		return SumInt8x32(collection)
	case simdFeatureAVX:
		return SumInt8x16(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumInt16 sums a slice of int16 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt16[T ~int16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumInt16x32(collection)
	case simdFeatureAVX2:
		return SumInt16x16(collection)
	case simdFeatureAVX:
		return SumInt16x8(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumInt32 sums a slice of int32 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumInt64 or handle overflow detection externally.
func SumInt32[T ~int32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumInt32x16(collection)
	case simdFeatureAVX2:
		return SumInt32x8(collection)
	case simdFeatureAVX:
		return SumInt32x4(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumInt64 sums a slice of int64 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumInt64[T ~int64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumInt64x8(collection)
	case simdFeatureAVX2:
		return SumInt64x4(collection)
	case simdFeatureAVX:
		return SumInt64x2(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumUint8 sums a slice of uint8 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint8[T ~uint8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumUint8x64(collection)
	case simdFeatureAVX2:
		return SumUint8x32(collection)
	case simdFeatureAVX:
		return SumUint8x16(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumUint16 sums a slice of uint16 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint16[T ~uint16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumUint16x32(collection)
	case simdFeatureAVX2:
		return SumUint16x16(collection)
	case simdFeatureAVX:
		return SumUint16x8(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumUint32 sums a slice of uint32 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumUint64 or handle overflow detection externally.
func SumUint32[T ~uint32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumUint32x16(collection)
	case simdFeatureAVX2:
		return SumUint32x8(collection)
	case simdFeatureAVX:
		return SumUint32x4(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumUint64 sums a slice of uint64 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumUint64[T ~uint64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumUint64x8(collection)
	case simdFeatureAVX2:
		return SumUint64x4(collection)
	case simdFeatureAVX:
		return SumUint64x2(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumFloat32 sums a slice of float32 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumFloat64.
func SumFloat32[T ~float32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumFloat32x16(collection)
	case simdFeatureAVX2:
		return SumFloat32x8(collection)
	case simdFeatureAVX:
		return SumFloat32x4(collection)
	default:
		return lo.Sum(collection)
	}
}

// SumFloat64 sums a slice of float64 using the best available SIMD instruction set.
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
func SumFloat64[T ~float64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumFloat64x8(collection)
	case simdFeatureAVX2:
		return SumFloat64x4(collection)
	case simdFeatureAVX:
		return SumFloat64x2(collection)
	default:
		return lo.Sum(collection)
	}
}

// MeanInt8 calculates the mean of a slice of int8 using the best available SIMD instruction set.
func MeanInt8[T ~int8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanInt8x64(collection)
	case simdFeatureAVX2:
		return MeanInt8x32(collection)
	case simdFeatureAVX:
		return MeanInt8x16(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanInt16 calculates the mean of a slice of int16 using the best available SIMD instruction set.
func MeanInt16[T ~int16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanInt16x32(collection)
	case simdFeatureAVX2:
		return MeanInt16x16(collection)
	case simdFeatureAVX:
		return MeanInt16x8(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanInt32 calculates the mean of a slice of int32 using the best available SIMD instruction set.
func MeanInt32[T ~int32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanInt32x16(collection)
	case simdFeatureAVX2:
		return MeanInt32x8(collection)
	case simdFeatureAVX:
		return MeanInt32x4(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanInt64 calculates the mean of a slice of int64 using the best available SIMD instruction set.
func MeanInt64[T ~int64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanInt64x8(collection)
	case simdFeatureAVX2:
		return MeanInt64x4(collection)
	case simdFeatureAVX:
		return MeanInt64x2(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanUint8 calculates the mean of a slice of uint8 using the best available SIMD instruction set.
func MeanUint8[T ~uint8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanUint8x64(collection)
	case simdFeatureAVX2:
		return MeanUint8x32(collection)
	case simdFeatureAVX:
		return MeanUint8x16(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanUint16 calculates the mean of a slice of uint16 using the best available SIMD instruction set.
func MeanUint16[T ~uint16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanUint16x32(collection)
	case simdFeatureAVX2:
		return MeanUint16x16(collection)
	case simdFeatureAVX:
		return MeanUint16x8(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanUint32 calculates the mean of a slice of uint32 using the best available SIMD instruction set.
func MeanUint32[T ~uint32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanUint32x16(collection)
	case simdFeatureAVX2:
		return MeanUint32x8(collection)
	case simdFeatureAVX:
		return MeanUint32x4(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanUint64 calculates the mean of a slice of uint64 using the best available SIMD instruction set.
func MeanUint64[T ~uint64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanUint64x8(collection)
	case simdFeatureAVX2:
		return MeanUint64x4(collection)
	case simdFeatureAVX:
		return MeanUint64x2(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanFloat32 calculates the mean of a slice of float32 using the best available SIMD instruction set.
func MeanFloat32[T ~float32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanFloat32x16(collection)
	case simdFeatureAVX2:
		return MeanFloat32x8(collection)
	case simdFeatureAVX:
		return MeanFloat32x4(collection)
	default:
		return lo.Mean(collection)
	}
}

// MeanFloat64 calculates the mean of a slice of float64 using the best available SIMD instruction set.
func MeanFloat64[T ~float64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanFloat64x8(collection)
	case simdFeatureAVX2:
		return MeanFloat64x4(collection)
	case simdFeatureAVX:
		return MeanFloat64x2(collection)
	default:
		return lo.Mean(collection)
	}
}

// MinInt8 finds the minimum value in a collection of int8 using the best available SIMD instruction set.
func MinInt8[T ~int8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinInt8x64(collection)
	case simdFeatureAVX2:
		return MinInt8x32(collection)
	case simdFeatureAVX:
		return MinInt8x16(collection)
	default:
		return lo.Min(collection)
	}
}

// MinInt16 finds the minimum value in a collection of int16 using the best available SIMD instruction set.
func MinInt16[T ~int16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinInt16x32(collection)
	case simdFeatureAVX2:
		return MinInt16x16(collection)
	case simdFeatureAVX:
		return MinInt16x8(collection)
	default:
		return lo.Min(collection)
	}
}

// MinInt32 finds the minimum value in a collection of int32 using the best available SIMD instruction set.
func MinInt32[T ~int32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinInt32x16(collection)
	case simdFeatureAVX2:
		return MinInt32x8(collection)
	case simdFeatureAVX:
		return MinInt32x4(collection)
	default:
		return lo.Min(collection)
	}
}

// MinInt64 finds the minimum value in a collection of int64 using the best available SIMD instruction set.
func MinInt64[T ~int64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinInt64x8(collection)
	case simdFeatureAVX2:
		return MinInt64x4(collection)
	case simdFeatureAVX:
		// MinInt64x2 requires AVX-512 (archsimd Int64x2.Min); use scalar fallback
		fallthrough
	default:
		return lo.Min(collection)
	}
}

// MinUint8 finds the minimum value in a collection of uint8 using the best available SIMD instruction set.
func MinUint8[T ~uint8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinUint8x64(collection)
	case simdFeatureAVX2:
		return MinUint8x32(collection)
	case simdFeatureAVX:
		return MinUint8x16(collection)
	default:
		return lo.Min(collection)
	}
}

// MinUint16 finds the minimum value in a collection of uint16 using the best available SIMD instruction set.
func MinUint16[T ~uint16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinUint16x32(collection)
	case simdFeatureAVX2:
		return MinUint16x16(collection)
	case simdFeatureAVX:
		return MinUint16x8(collection)
	default:
		return lo.Min(collection)
	}
}

// MinUint32 finds the minimum value in a collection of uint32 using the best available SIMD instruction set.
func MinUint32[T ~uint32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinUint32x16(collection)
	case simdFeatureAVX2:
		return MinUint32x8(collection)
	case simdFeatureAVX:
		return MinUint32x4(collection)
	default:
		return lo.Min(collection)
	}
}

// MinUint64 finds the minimum value in a collection of uint64 using the best available SIMD instruction set.
func MinUint64[T ~uint64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinUint64x8(collection)
	case simdFeatureAVX2:
		return MinUint64x4(collection)
	case simdFeatureAVX:
		// MinUint64x2 requires AVX-512; use scalar fallback
		fallthrough
	default:
		return lo.Min(collection)
	}
}

// MinFloat32 finds the minimum value in a collection of float32 using the best available SIMD instruction set.
func MinFloat32[T ~float32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinFloat32x16(collection)
	case simdFeatureAVX2:
		return MinFloat32x8(collection)
	case simdFeatureAVX:
		return MinFloat32x4(collection)
	default:
		return lo.Min(collection)
	}
}

// MinFloat64 finds the minimum value in a collection of float64 using the best available SIMD instruction set.
func MinFloat64[T ~float64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MinFloat64x8(collection)
	case simdFeatureAVX2:
		return MinFloat64x4(collection)
	case simdFeatureAVX:
		return MinFloat64x2(collection)
	default:
		return lo.Min(collection)
	}
}

// MaxInt8 finds the maximum value in a collection of int8 using the best available SIMD instruction set.
func MaxInt8[T ~int8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxInt8x64(collection)
	case simdFeatureAVX2:
		return MaxInt8x32(collection)
	case simdFeatureAVX:
		return MaxInt8x16(collection)
	default:
		return lo.Max(collection)
	}
}

// MaxInt16 finds the maximum value in a collection of int16 using the best available SIMD instruction set.
func MaxInt16[T ~int16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxInt16x32(collection)
	case simdFeatureAVX2:
		return MaxInt16x16(collection)
	case simdFeatureAVX:
		return MaxInt16x8(collection)
	default:
		return lo.Max(collection)
	}
}

// MaxInt32 finds the maximum value in a collection of int32 using the best available SIMD instruction set.
func MaxInt32[T ~int32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxInt32x16(collection)
	case simdFeatureAVX2:
		return MaxInt32x8(collection)
	case simdFeatureAVX:
		return MaxInt32x4(collection)
	default:
		return lo.Max(collection)
	}
}

// MaxInt64 finds the maximum value in a collection of int64 using the best available SIMD instruction set.
func MaxInt64[T ~int64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxInt64x8(collection)
	case simdFeatureAVX2:
		return MaxInt64x4(collection)
	case simdFeatureAVX:
		// MaxInt64x2 requires AVX-512; use scalar fallback
		fallthrough
	default:
		return lo.Max(collection)
	}
}

// MaxUint8 finds the maximum value in a collection of uint8 using the best available SIMD instruction set.
func MaxUint8[T ~uint8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxUint8x64(collection)
	case simdFeatureAVX2:
		return MaxUint8x32(collection)
	case simdFeatureAVX:
		return MaxUint8x16(collection)
	default:
		return lo.Max(collection)
	}
}

// MaxUint16 finds the maximum value in a collection of uint16 using the best available SIMD instruction set.
func MaxUint16[T ~uint16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxUint16x32(collection)
	case simdFeatureAVX2:
		return MaxUint16x16(collection)
	case simdFeatureAVX:
		return MaxUint16x8(collection)
	default:
		return lo.Max(collection)
	}
}

// MaxUint32 finds the maximum value in a collection of uint32 using the best available SIMD instruction set.
func MaxUint32[T ~uint32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxUint32x16(collection)
	case simdFeatureAVX2:
		return MaxUint32x8(collection)
	case simdFeatureAVX:
		return MaxUint32x4(collection)
	default:
		return lo.Max(collection)
	}
}

// MaxUint64 finds the maximum value in a collection of uint64 using the best available SIMD instruction set.
func MaxUint64[T ~uint64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxUint64x8(collection)
	case simdFeatureAVX2:
		return MaxUint64x4(collection)
	case simdFeatureAVX:
		// MaxUint64x2 requires AVX-512; use scalar fallback
		fallthrough
	default:
		return lo.Max(collection)
	}
}

// MaxFloat32 finds the maximum value in a collection of float32 using the best available SIMD instruction set.
func MaxFloat32[T ~float32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxFloat32x16(collection)
	case simdFeatureAVX2:
		return MaxFloat32x8(collection)
	case simdFeatureAVX:
		return MaxFloat32x4(collection)
	default:
		return lo.Max(collection)
	}
}

// MaxFloat64 finds the maximum value in a collection of float64 using the best available SIMD instruction set.
func MaxFloat64[T ~float64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MaxFloat64x8(collection)
	case simdFeatureAVX2:
		return MaxFloat64x4(collection)
	case simdFeatureAVX:
		return MaxFloat64x2(collection)
	default:
		return lo.Max(collection)
	}
}

// ClampInt8 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampInt8[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampInt8x64(collection, min, max)
	case simdFeatureAVX2:
		return ClampInt8x32(collection, min, max)
	case simdFeatureAVX:
		return ClampInt8x16(collection, min, max)
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampInt16 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampInt16[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampInt16x32(collection, min, max)
	case simdFeatureAVX2:
		return ClampInt16x16(collection, min, max)
	case simdFeatureAVX:
		return ClampInt16x8(collection, min, max)
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampInt32 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampInt32[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampInt32x16(collection, min, max)
	case simdFeatureAVX2:
		return ClampInt32x8(collection, min, max)
	case simdFeatureAVX:
		return ClampInt32x4(collection, min, max)
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampInt64 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampInt64[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampInt64x8(collection, min, max)
	case simdFeatureAVX2:
		return ClampInt64x4(collection, min, max)
	case simdFeatureAVX:
		// ClampInt64x2 requires AVX-512; use scalar fallback
		fallthrough
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampUint8 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampUint8[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampUint8x64(collection, min, max)
	case simdFeatureAVX2:
		return ClampUint8x32(collection, min, max)
	case simdFeatureAVX:
		return ClampUint8x16(collection, min, max)
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampUint16 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampUint16[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampUint16x32(collection, min, max)
	case simdFeatureAVX2:
		return ClampUint16x16(collection, min, max)
	case simdFeatureAVX:
		return ClampUint16x8(collection, min, max)
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampUint32 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampUint32[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampUint32x16(collection, min, max)
	case simdFeatureAVX2:
		return ClampUint32x8(collection, min, max)
	case simdFeatureAVX:
		return ClampUint32x4(collection, min, max)
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampUint64 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampUint64[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampUint64x8(collection, min, max)
	case simdFeatureAVX2:
		return ClampUint64x4(collection, min, max)
	case simdFeatureAVX:
		// ClampUint64x2 requires AVX-512; use scalar fallback
		fallthrough
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampFloat32 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampFloat32[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampFloat32x16(collection, min, max)
	case simdFeatureAVX2:
		return ClampFloat32x8(collection, min, max)
	case simdFeatureAVX:
		return ClampFloat32x4(collection, min, max)
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// ClampFloat64 clamps each element in collection between min and max values using the best available SIMD instruction set.
func ClampFloat64[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return ClampFloat64x8(collection, min, max)
	case simdFeatureAVX2:
		return ClampFloat64x4(collection, min, max)
	case simdFeatureAVX:
		return ClampFloat64x2(collection, min, max)
	default:
		result := make(Slice, len(collection))
		for i, v := range collection {
			if v < min {
				result[i] = min
			} else if v > max {
				result[i] = max
			} else {
				result[i] = v
			}
		}
		return result
	}
}

// SumByInt8 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByInt8[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByInt8x64(collection, iteratee)
	case simdFeatureAVX2:
		return SumByInt8x32(collection, iteratee)
	case simdFeatureAVX:
		return SumByInt8x16(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByInt16 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByInt16[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByInt16x32(collection, iteratee)
	case simdFeatureAVX2:
		return SumByInt16x16(collection, iteratee)
	case simdFeatureAVX:
		return SumByInt16x8(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByInt32 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumByInt64 or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByInt32[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByInt32x16(collection, iteratee)
	case simdFeatureAVX2:
		return SumByInt32x8(collection, iteratee)
	case simdFeatureAVX:
		return SumByInt32x4(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByInt64 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
// Play: https://go.dev/play/p/TBD
func SumByInt64[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByInt64x8(collection, iteratee)
	case simdFeatureAVX2:
		return SumByInt64x4(collection, iteratee)
	case simdFeatureAVX:
		return SumByInt64x2(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByUint8 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByUint8[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByUint8x64(collection, iteratee)
	case simdFeatureAVX2:
		return SumByUint8x32(collection, iteratee)
	case simdFeatureAVX:
		return SumByUint8x16(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByUint16 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByUint16[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByUint16x32(collection, iteratee)
	case simdFeatureAVX2:
		return SumByUint16x16(collection, iteratee)
	case simdFeatureAVX:
		return SumByUint16x8(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByUint32 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumByUint64 or handle overflow detection externally.
// Play: https://go.dev/play/p/TBD
func SumByUint32[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByUint32x16(collection, iteratee)
	case simdFeatureAVX2:
		return SumByUint32x8(collection, iteratee)
	case simdFeatureAVX:
		return SumByUint32x4(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByUint64 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
// Play: https://go.dev/play/p/TBD
func SumByUint64[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByUint64x8(collection, iteratee)
	case simdFeatureAVX2:
		return SumByUint64x4(collection, iteratee)
	case simdFeatureAVX:
		return SumByUint64x2(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByFloat32 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumByFloat64.
// Play: https://go.dev/play/p/TBD
func SumByFloat32[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByFloat32x16(collection, iteratee)
	case simdFeatureAVX2:
		return SumByFloat32x8(collection, iteratee)
	case simdFeatureAVX:
		return SumByFloat32x4(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// SumByFloat64 sums the values extracted by iteratee from a slice using the best available SIMD instruction set.
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
// Play: https://go.dev/play/p/TBD
func SumByFloat64[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumByFloat64x8(collection, iteratee)
	case simdFeatureAVX2:
		return SumByFloat64x4(collection, iteratee)
	case simdFeatureAVX:
		return SumByFloat64x2(collection, iteratee)
	default:
		return lo.SumBy(collection, iteratee)
	}
}

// MeanByInt8 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByInt8[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByInt8x64(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByInt8x32(collection, iteratee)
	case simdFeatureAVX:
		return MeanByInt8x16(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByInt16 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByInt16[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByInt16x32(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByInt16x16(collection, iteratee)
	case simdFeatureAVX:
		return MeanByInt16x8(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByInt32 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByInt32[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByInt32x16(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByInt32x8(collection, iteratee)
	case simdFeatureAVX:
		return MeanByInt32x4(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByInt64 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByInt64[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByInt64x8(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByInt64x4(collection, iteratee)
	case simdFeatureAVX:
		return MeanByInt64x2(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByUint8 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByUint8[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByUint8x64(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByUint8x32(collection, iteratee)
	case simdFeatureAVX:
		return MeanByUint8x16(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByUint16 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByUint16[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByUint16x32(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByUint16x16(collection, iteratee)
	case simdFeatureAVX:
		return MeanByUint16x8(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByUint32 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByUint32[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByUint32x16(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByUint32x8(collection, iteratee)
	case simdFeatureAVX:
		return MeanByUint32x4(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByUint64 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByUint64[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByUint64x8(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByUint64x4(collection, iteratee)
	case simdFeatureAVX:
		return MeanByUint64x2(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByFloat32 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByFloat32[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByFloat32x16(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByFloat32x8(collection, iteratee)
	case simdFeatureAVX:
		return MeanByFloat32x4(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}

// MeanByFloat64 calculates the mean of values extracted by iteratee from a slice using the best available SIMD instruction set.
// Play: https://go.dev/play/p/TBD
func MeanByFloat64[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return MeanByFloat64x8(collection, iteratee)
	case simdFeatureAVX2:
		return MeanByFloat64x4(collection, iteratee)
	case simdFeatureAVX:
		return MeanByFloat64x2(collection, iteratee)
	default:
		return lo.MeanBy(collection, iteratee)
	}
}
