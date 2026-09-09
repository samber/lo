//go:build goexperiment.simd

package simd

import "github.com/samber/lo"

// MinInt8 returns the minimum of a slice of int8 using SIMD instructions when available.
// If collection is empty, 0 is returned.
func MinInt8[T ~int8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minInt8(asInt8(collection)))
}

// MinInt16 is MinInt8 for int16.
func MinInt16[T ~int16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minInt16(asInt16(collection)))
}

// MinInt32 is MinInt8 for int32.
func MinInt32[T ~int32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minInt32(asInt32(collection)))
}

// MinInt64 is MinInt8 for int64.
func MinInt64[T ~int64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minInt64(asInt64(collection)))
}

// MinUint8 is MinInt8 for uint8.
func MinUint8[T ~uint8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minUint8(asUint8(collection)))
}

// MinUint16 is MinInt8 for uint16.
func MinUint16[T ~uint16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minUint16(asUint16(collection)))
}

// MinUint32 is MinInt8 for uint32.
func MinUint32[T ~uint32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minUint32(asUint32(collection)))
}

// MinUint64 is MinInt8 for uint64.
func MinUint64[T ~uint64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minUint64(asUint64(collection)))
}

// MinFloat32 returns the minimum of a slice of float32 using SIMD instructions when
// available. NaN handling matches lo.Min exactly, on every architecture: if collection[0] is
// NaN, the result is NaN; otherwise NaN elsewhere in collection is ignored.
func MinFloat32[T ~float32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minFloat32(asFloat32(collection)))
}

// MinFloat64 is MinFloat32 for float64.
func MinFloat64[T ~float64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Min(collection)
	}
	return T(minFloat64(asFloat64(collection)))
}

// MaxInt8 returns the maximum of a slice of int8 using SIMD instructions when available.
// If collection is empty, 0 is returned.
func MaxInt8[T ~int8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxInt8(asInt8(collection)))
}

// MaxInt16 is MaxInt8 for int16.
func MaxInt16[T ~int16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxInt16(asInt16(collection)))
}

// MaxInt32 is MaxInt8 for int32.
func MaxInt32[T ~int32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxInt32(asInt32(collection)))
}

// MaxInt64 is MaxInt8 for int64.
func MaxInt64[T ~int64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxInt64(asInt64(collection)))
}

// MaxUint8 is MaxInt8 for uint8.
func MaxUint8[T ~uint8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxUint8(asUint8(collection)))
}

// MaxUint16 is MaxInt8 for uint16.
func MaxUint16[T ~uint16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxUint16(asUint16(collection)))
}

// MaxUint32 is MaxInt8 for uint32.
func MaxUint32[T ~uint32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxUint32(asUint32(collection)))
}

// MaxUint64 is MaxInt8 for uint64.
func MaxUint64[T ~uint64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxUint64(asUint64(collection)))
}

// MaxFloat32 returns the maximum of a slice of float32 using SIMD instructions when
// available. NaN handling matches lo.Max exactly, on every architecture: if collection[0] is
// NaN, the result is NaN; otherwise NaN elsewhere in collection is ignored.
func MaxFloat32[T ~float32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxFloat32(asFloat32(collection)))
}

// MaxFloat64 is MaxFloat32 for float64.
func MaxFloat64[T ~float64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Max(collection)
	}
	return T(maxFloat64(asFloat64(collection)))
}
