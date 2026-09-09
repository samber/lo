//go:build goexperiment.simd

package simd

import "github.com/samber/lo"

// SumInt8 sums a slice of int8 using SIMD instructions when available.
// Overflow: the accumulation is performed using int8, which can overflow for large
// collections. If the sum exceeds the int8 range (-128 to 127), the result wraps around
// silently, matching lo.Sum. For collections that may overflow, use a wider type.
func SumInt8[T ~int8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumInt8(asInt8(collection)))
}

// SumInt16 is SumInt8 for int16.
func SumInt16[T ~int16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumInt16(asInt16(collection)))
}

// SumInt32 is SumInt8 for int32.
func SumInt32[T ~int32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumInt32(asInt32(collection)))
}

// SumInt64 is SumInt8 for int64.
func SumInt64[T ~int64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumInt64(asInt64(collection)))
}

// SumUint8 is SumInt8 for uint8.
func SumUint8[T ~uint8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumUint8(asUint8(collection)))
}

// SumUint16 is SumInt8 for uint16.
func SumUint16[T ~uint16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumUint16(asUint16(collection)))
}

// SumUint32 is SumInt8 for uint32.
func SumUint32[T ~uint32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumUint32(asUint32(collection)))
}

// SumUint64 is SumInt8 for uint64.
func SumUint64[T ~uint64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumUint64(asUint64(collection)))
}

// SumFloat32 sums a slice of float32 using SIMD instructions when available.
// Precision: floating-point addition is not associative, and the SIMD path accumulates
// several independent lane totals before combining them, instead of lo.Sum's strictly
// sequential left-to-right addition. The result is numerically equivalent (within normal
// floating-point rounding error) but not always bit-identical to lo.Sum's — the same
// trade-off every vectorized numeric library makes.
func SumFloat32[T ~float32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumFloat32(asFloat32(collection)))
}

// SumFloat64 is SumFloat32 for float64.
func SumFloat64[T ~float64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.Sum(collection)
	}
	return T(sumFloat64(asFloat64(collection)))
}

// MeanInt8 calculates the mean of a slice of int8 using SIMD instructions when available.
func MeanInt8[T ~int8](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumInt8(collection) / length
}

// MeanInt16 is MeanInt8 for int16.
func MeanInt16[T ~int16](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumInt16(collection) / length
}

// MeanInt32 is MeanInt8 for int32.
func MeanInt32[T ~int32](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumInt32(collection) / length
}

// MeanInt64 is MeanInt8 for int64.
func MeanInt64[T ~int64](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumInt64(collection) / length
}

// MeanUint8 is MeanInt8 for uint8.
func MeanUint8[T ~uint8](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumUint8(collection) / length
}

// MeanUint16 is MeanInt8 for uint16.
func MeanUint16[T ~uint16](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumUint16(collection) / length
}

// MeanUint32 is MeanInt8 for uint32.
func MeanUint32[T ~uint32](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumUint32(collection) / length
}

// MeanUint64 is MeanInt8 for uint64.
func MeanUint64[T ~uint64](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumUint64(collection) / length
}

// MeanFloat32 is MeanInt8 for float32. See SumFloat32 for the floating-point
// non-associativity caveat: the division is exact, but the numerator it divides may differ
// from lo.Sum's in the last bit or two.
func MeanFloat32[T ~float32](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumFloat32(collection) / length
}

// MeanFloat64 is MeanInt8 for float64.
func MeanFloat64[T ~float64](collection []T) T {
	length := T(len(collection))
	if length == 0 {
		return 0
	}
	return SumFloat64(collection) / length
}

// SumByInt8 summarizes a collection into an int8 using the given iteratee, using SIMD
// instructions when available. See SumInt8 for overflow behaviour.
func SumByInt8[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumInt8(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByInt16 is SumByInt8 for int16.
func SumByInt16[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumInt16(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByInt32 is SumByInt8 for int32.
func SumByInt32[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumInt32(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByInt64 is SumByInt8 for int64.
func SumByInt64[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumInt64(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByUint8 is SumByInt8 for uint8.
func SumByUint8[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumUint8(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByUint16 is SumByInt8 for uint16.
func SumByUint16[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumUint16(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByUint32 is SumByInt8 for uint32.
func SumByUint32[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumUint32(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByUint64 is SumByInt8 for uint64.
func SumByUint64[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumUint64(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByFloat32 is SumByInt8 for float32. See SumFloat32 for the floating-point
// non-associativity caveat that applies to the sum of the mapped values.
func SumByFloat32[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumFloat32(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// SumByFloat64 is SumByInt8 for float64.
func SumByFloat64[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	if len(collection) == 0 {
		return 0
	}
	if !useSIMD {
		return lo.SumBy(collection, iteratee)
	}
	return SumFloat64(lo.Map(collection, func(item T, _ int) R { return iteratee(item) }))
}

// MeanByInt8 calculates the mean of a collection of int8 using the given iteratee, using
// SIMD instructions when available.
func MeanByInt8[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByInt8(collection, iteratee) / length
}

// MeanByInt16 is MeanByInt8 for int16.
func MeanByInt16[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByInt16(collection, iteratee) / length
}

// MeanByInt32 is MeanByInt8 for int32.
func MeanByInt32[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByInt32(collection, iteratee) / length
}

// MeanByInt64 is MeanByInt8 for int64.
func MeanByInt64[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByInt64(collection, iteratee) / length
}

// MeanByUint8 is MeanByInt8 for uint8.
func MeanByUint8[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByUint8(collection, iteratee) / length
}

// MeanByUint16 is MeanByInt8 for uint16.
func MeanByUint16[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByUint16(collection, iteratee) / length
}

// MeanByUint32 is MeanByInt8 for uint32.
func MeanByUint32[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByUint32(collection, iteratee) / length
}

// MeanByUint64 is MeanByInt8 for uint64.
func MeanByUint64[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByUint64(collection, iteratee) / length
}

// MeanByFloat32 is MeanByInt8 for float32. See SumFloat32 for the floating-point
// non-associativity caveat that applies to the underlying sum.
func MeanByFloat32[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByFloat32(collection, iteratee) / length
}

// MeanByFloat64 is MeanByInt8 for float64.
func MeanByFloat64[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return SumByFloat64(collection, iteratee) / length
}

// ClampInt8 clamps each element in collection between mn and mx using SIMD instructions
// when available. If mn > mx, every element is clamped to mn: a well-defined convention
// chosen so the SIMD and scalar-fallback code paths always agree, since lo.Clamp itself is
// value-dependent (not consistently mn) when its bounds are swapped.
func ClampInt8[T ~int8, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampInt8(asInt8(result), asInt8(collection), int8(mn), int8(mx))
	return result
}

// ClampInt16 is ClampInt8 for int16.
func ClampInt16[T ~int16, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampInt16(asInt16(result), asInt16(collection), int16(mn), int16(mx))
	return result
}

// ClampInt32 is ClampInt8 for int32.
func ClampInt32[T ~int32, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampInt32(asInt32(result), asInt32(collection), int32(mn), int32(mx))
	return result
}

// ClampInt64 is ClampInt8 for int64.
func ClampInt64[T ~int64, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampInt64(asInt64(result), asInt64(collection), int64(mn), int64(mx))
	return result
}

// ClampUint8 is ClampInt8 for uint8.
func ClampUint8[T ~uint8, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampUint8(asUint8(result), asUint8(collection), uint8(mn), uint8(mx))
	return result
}

// ClampUint16 is ClampInt8 for uint16.
func ClampUint16[T ~uint16, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampUint16(asUint16(result), asUint16(collection), uint16(mn), uint16(mx))
	return result
}

// ClampUint32 is ClampInt8 for uint32.
func ClampUint32[T ~uint32, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampUint32(asUint32(result), asUint32(collection), uint32(mn), uint32(mx))
	return result
}

// ClampUint64 is ClampInt8 for uint64.
func ClampUint64[T ~uint64, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampUint64(asUint64(result), asUint64(collection), uint64(mn), uint64(mx))
	return result
}

// ClampFloat32 clamps each element in collection between mn and mx using SIMD instructions
// when available. If mn > mx, every element is clamped to mn: a well-defined convention
// chosen so the SIMD and scalar-fallback code paths always agree, since lo.Clamp itself is
// value-dependent (not consistently mn) when its bounds are swapped. A NaN element is
// returned unchanged, matching lo.Clamp exactly on every architecture.
func ClampFloat32[T ~float32, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampFloat32(asFloat32(result), asFloat32(collection), float32(mn), float32(mx))
	return result
}

// ClampFloat64 is ClampFloat32 for float64.
func ClampFloat64[T ~float64, Slice ~[]T](collection Slice, mn, mx T) Slice {
	if len(collection) == 0 {
		return collection
	}
	result := make(Slice, len(collection))
	if !useSIMD {
		for i, v := range collection {
			result[i] = clampScalar(v, mn, mx)
		}
		return result
	}
	clampFloat64(asFloat64(result), asFloat64(collection), float64(mn), float64(mx))
	return result
}
