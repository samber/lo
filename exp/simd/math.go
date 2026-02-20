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
