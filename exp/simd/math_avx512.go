//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"

	"github.com/samber/lo"
)

// AVX-512 (512-bit) SIMD sum functions - 64/32/16/8 lanes

// SumInt8x64 sums a slice of int8 using AVX-512 SIMD (Int8x64, 64 lanes).
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt8x64[T ~int8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes64

	base := unsafeSliceInt8(collection, length)
	var acc archsimd.Int8x64

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x64Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]int8
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumInt16x32 sums a slice of int16 using AVX-512 SIMD (Int16x32, 32 lanes).
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt16x32[T ~int16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes32

	base := unsafeSliceInt16(collection, length)
	var acc archsimd.Int16x32

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x32Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]int16
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumInt32x16 sums a slice of int32 using AVX-512 SIMD (Int32x16, 16 lanes).
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumInt64x8 or handle overflow detection externally.
func SumInt32x16[T ~int32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes16

	base := unsafeSliceInt32(collection, length)
	var acc archsimd.Int32x16

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x16Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]int32
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumInt64x8 sums a slice of int64 using AVX-512 SIMD (Int64x8, 8 lanes).
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumInt64x8[T ~int64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes8

	base := unsafeSliceInt64(collection, length)
	var acc archsimd.Int64x8

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]int64
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumUint8x64 sums a slice of uint8 using AVX-512 SIMD (Uint8x64, 64 lanes).
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint8x64[T ~uint8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes64

	base := unsafeSliceUint8(collection, length)
	var acc archsimd.Uint8x64

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x64Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]uint8
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumUint16x32 sums a slice of uint16 using AVX-512 SIMD (Uint16x32, 32 lanes).
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint16x32[T ~uint16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes32

	base := unsafeSliceUint16(collection, length)
	var acc archsimd.Uint16x32

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x32Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]uint16
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumUint32x16 sums a slice of uint32 using AVX-512 SIMD (Uint32x16, 16 lanes).
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumUint64x8 or handle overflow detection externally.
func SumUint32x16[T ~uint32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes16

	base := unsafeSliceUint32(collection, length)
	var acc archsimd.Uint32x16

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x16Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]uint32
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumUint64x8 sums a slice of uint64 using AVX-512 SIMD (Uint64x8, 8 lanes).
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumUint64x8[T ~uint64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes8

	base := unsafeSliceUint64(collection, length)
	var acc archsimd.Uint64x8

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]uint64
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumFloat32x16 sums a slice of float32 using AVX-512 SIMD (Float32x16, 16 lanes).
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumFloat64x8.
func SumFloat32x16[T ~float32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes16

	base := unsafeSliceFloat32(collection, length)
	var acc archsimd.Float32x16

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x16Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]float32
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumFloat64x8 sums a slice of float64 using AVX-512 SIMD (Float64x8, 8 lanes).
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
func SumFloat64x8[T ~float64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes8

	base := unsafeSliceFloat64(collection, length)
	var acc archsimd.Float64x8

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [lanes]float64
	acc.Store(&buf)
	var sum T
	for k := uint(0); k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// MeanInt8x64 calculates the mean of a slice of int8 using AVX-512 SIMD
func MeanInt8x64[T ~int8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumInt8x64(collection)
	return sum / T(length)
}

// MeanInt16x32 calculates the mean of a slice of int16 using AVX-512 SIMD
func MeanInt16x32[T ~int16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumInt16x32(collection)
	return sum / T(length)
}

// MeanInt32x16 calculates the mean of a slice of int32 using AVX-512 SIMD
func MeanInt32x16[T ~int32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumInt32x16(collection)
	return sum / T(length)
}

// MeanInt64x8 calculates the mean of a slice of int64 using AVX-512 SIMD
func MeanInt64x8[T ~int64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumInt64x8(collection)
	return sum / T(length)
}

// MeanUint8x64 calculates the mean of a slice of uint8 using AVX-512 SIMD
func MeanUint8x64[T ~uint8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumUint8x64(collection)
	return sum / T(length)
}

// MeanUint16x32 calculates the mean of a slice of uint16 using AVX-512 SIMD
func MeanUint16x32[T ~uint16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumUint16x32(collection)
	return sum / T(length)
}

// MeanUint32x16 calculates the mean of a slice of uint32 using AVX-512 SIMD
func MeanUint32x16[T ~uint32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumUint32x16(collection)
	return sum / T(length)
}

// MeanUint64x8 calculates the mean of a slice of uint64 using AVX-512 SIMD
func MeanUint64x8[T ~uint64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumUint64x8(collection)
	return sum / T(length)
}

// MeanFloat32x16 calculates the mean of a slice of float32 using AVX-512 SIMD
func MeanFloat32x16[T ~float32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumFloat32x16(collection)
	return sum / T(length)
}

// MeanFloat64x8 calculates the mean of a slice of float64 using AVX-512 SIMD
func MeanFloat64x8[T ~float64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumFloat64x8(collection)
	return sum / T(length)
}

// ClampInt8x64 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt8x64[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes64

	minVec := archsimd.BroadcastInt8x64(int8(min))
	maxVec := archsimd.BroadcastInt8x64(int8(max))

	base := unsafeSliceInt8(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadInt8x64Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]int8)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampInt16x32 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt16x32[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes32

	minVec := archsimd.BroadcastInt16x32(int16(min))
	maxVec := archsimd.BroadcastInt16x32(int16(max))

	base := unsafeSliceInt16(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadInt16x32Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]int16)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampInt32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt32x16[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes16

	minVec := archsimd.BroadcastInt32x16(int32(min))
	maxVec := archsimd.BroadcastInt32x16(int32(max))

	base := unsafeSliceInt32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadInt32x16Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]int32)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampInt64x2 clamps each element in collection between min and max values using AVX-512 SIMD.
// Int64x2 Min/Max operations in archsimd require AVX-512 (VPMAXSQ/VPMINSQ).
func ClampInt64x2[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes2

	base := unsafeSliceInt64(collection, length)

	minVec := archsimd.BroadcastInt64x2(int64(min))
	maxVec := archsimd.BroadcastInt64x2(int64(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x2Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]int64)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampUint64x2 clamps each element in collection between min and max values using AVX-512 SIMD.
// Uint64x2 Min/Max operations in archsimd require AVX-512.
func ClampUint64x2[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes2

	base := unsafeSliceUint64(collection, length)

	minVec := archsimd.BroadcastUint64x2(uint64(min))
	maxVec := archsimd.BroadcastUint64x2(uint64(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x2Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]uint64)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampInt64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt64x8[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes8

	minVec := archsimd.BroadcastInt64x8(int64(min))
	maxVec := archsimd.BroadcastInt64x8(int64(max))

	base := unsafeSliceInt64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadInt64x8Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]int64)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampUint8x64 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint8x64[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes64

	minVec := archsimd.BroadcastUint8x64(uint8(min))
	maxVec := archsimd.BroadcastUint8x64(uint8(max))

	base := unsafeSliceUint8(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadUint8x64Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]uint8)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampUint16x32 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint16x32[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes32

	minVec := archsimd.BroadcastUint16x32(uint16(min))
	maxVec := archsimd.BroadcastUint16x32(uint16(max))

	base := unsafeSliceUint16(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadUint16x32Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]uint16)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampUint32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint32x16[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes16

	minVec := archsimd.BroadcastUint32x16(uint32(min))
	maxVec := archsimd.BroadcastUint32x16(uint32(max))

	base := unsafeSliceUint32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadUint32x16Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]uint32)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampUint64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint64x8[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes8

	minVec := archsimd.BroadcastUint64x8(uint64(min))
	maxVec := archsimd.BroadcastUint64x8(uint64(max))

	base := unsafeSliceUint64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadUint64x8Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]uint64)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampFloat32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampFloat32x16[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes16

	minVec := archsimd.BroadcastFloat32x16(float32(min))
	maxVec := archsimd.BroadcastFloat32x16(float32(max))

	base := unsafeSliceFloat32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadFloat32x16Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]float32)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// ClampFloat64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampFloat64x8[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes8

	minVec := archsimd.BroadcastFloat64x8(float64(min))
	maxVec := archsimd.BroadcastFloat64x8(float64(max))

	base := unsafeSliceFloat64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadFloat64x8Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		// bearer:disable go_gosec_unsafe_unsafe
		clamped.Store((*[lanes]float64)(unsafe.Pointer(&result[i])))
	}

	for ; i < length; i++ {
		val := collection[i]
		if val < min {
			val = min
		} else if val > max {
			val = max
		}
		result[i] = val
	}

	return result
}

// MinInt8x64 finds the minimum value in a collection of int8 using AVX-512 SIMD
func MinInt8x64[T ~int8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes64

	base := unsafeSliceInt8(collection, length)

	var minVec archsimd.Int8x64
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x64Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal int8
	if firstInitialized {
		var buf [lanes]int8
		minVec.Store(&buf)
		minVal = min(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
			buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23],
			buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31],
			buf[32], buf[33], buf[34], buf[35], buf[36], buf[37], buf[38], buf[39],
			buf[40], buf[41], buf[42], buf[43], buf[44], buf[45], buf[46], buf[47],
			buf[48], buf[49], buf[50], buf[51], buf[52], buf[53], buf[54], buf[55],
			buf[56], buf[57], buf[58], buf[59], buf[60], buf[61], buf[62], buf[63],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = int8(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinInt16x32 finds the minimum value in a collection of int16 using AVX-512 SIMD
func MinInt16x32[T ~int16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes32

	base := unsafeSliceInt16(collection, length)

	var minVec archsimd.Int16x32
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x32Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal int16
	if firstInitialized {
		var buf [lanes]int16
		minVec.Store(&buf)
		minVal = min(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
			buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23],
			buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = int16(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinInt32x16 finds the minimum value in a collection of int32 using AVX-512 SIMD
func MinInt32x16[T ~int32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16

	base := unsafeSliceInt32(collection, length)

	var minVec archsimd.Int32x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x16Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal int32
	if firstInitialized {
		var buf [lanes]int32
		minVec.Store(&buf)
		minVal = min(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = int32(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinInt64x2 finds the minimum value in a collection of int64 using AVX-512 SIMD.
// Int64x2 Min operations in archsimd require AVX-512.
func MinInt64x2[T ~int64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes2
	base := unsafeSliceInt64(collection, length)

	var minVec archsimd.Int64x2
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x2Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal int64
	if firstInitialized {
		var buf [lanes]int64
		minVec.Store(&buf)
		minVal = min(buf[0], buf[1])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = int64(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinUint64x2 finds the minimum value in a collection of uint64 using AVX-512 SIMD.
// Uint64x2 Min operations in archsimd require AVX-512.
func MinUint64x2[T ~uint64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes2
	base := unsafeSliceUint64(collection, length)

	var minVec archsimd.Uint64x2
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x2Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal uint64
	if firstInitialized {
		var buf [lanes]uint64
		minVec.Store(&buf)
		minVal = min(buf[0], buf[1])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = uint64(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinInt64x8 finds the minimum value in a collection of int64 using AVX-512 SIMD
func MinInt64x8[T ~int64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8

	base := unsafeSliceInt64(collection, length)

	var minVec archsimd.Int64x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x8Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal int64
	if firstInitialized {
		var buf [lanes]int64
		minVec.Store(&buf)
		minVal = min(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = int64(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinUint8x64 finds the minimum value in a collection of uint8 using AVX-512 SIMD
func MinUint8x64[T ~uint8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes64

	base := unsafeSliceUint8(collection, length)

	var minVec archsimd.Uint8x64
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x64Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal uint8
	if firstInitialized {
		var buf [lanes]uint8
		minVec.Store(&buf)
		minVal = min(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
			buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23],
			buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31],
			buf[32], buf[33], buf[34], buf[35], buf[36], buf[37], buf[38], buf[39],
			buf[40], buf[41], buf[42], buf[43], buf[44], buf[45], buf[46], buf[47],
			buf[48], buf[49], buf[50], buf[51], buf[52], buf[53], buf[54], buf[55],
			buf[56], buf[57], buf[58], buf[59], buf[60], buf[61], buf[62], buf[63],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = uint8(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinUint16x32 finds the minimum value in a collection of uint16 using AVX-512 SIMD
func MinUint16x32[T ~uint16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes32

	base := unsafeSliceUint16(collection, length)

	var minVec archsimd.Uint16x32
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x32Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal uint16
	if firstInitialized {
		var buf [lanes]uint16
		minVec.Store(&buf)
		minVal = min(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
			buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23],
			buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = uint16(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinUint32x16 finds the minimum value in a collection of uint32 using AVX-512 SIMD
func MinUint32x16[T ~uint32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16

	base := unsafeSliceUint32(collection, length)

	var minVec archsimd.Uint32x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x16Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal uint32
	if firstInitialized {
		var buf [lanes]uint32
		minVec.Store(&buf)
		minVal = min(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = uint32(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinUint64x8 finds the minimum value in a collection of uint64 using AVX-512 SIMD
func MinUint64x8[T ~uint64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8

	base := unsafeSliceUint64(collection, length)

	var minVec archsimd.Uint64x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x8Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal uint64
	if firstInitialized {
		var buf [lanes]uint64
		minVec.Store(&buf)
		minVal = min(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = uint64(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinFloat32x16 finds the minimum value in a collection of float32 using AVX-512 SIMD
func MinFloat32x16[T ~float32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16

	base := unsafeSliceFloat32(collection, length)

	var minVec archsimd.Float32x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x16Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal float32
	if firstInitialized {
		var buf [lanes]float32
		minVec.Store(&buf)
		minVal = min(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = float32(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MinFloat64x8 finds the minimum value in a collection of float64 using AVX-512 SIMD
func MinFloat64x8[T ~float64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8

	base := unsafeSliceFloat64(collection, length)

	var minVec archsimd.Float64x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x8Slice(base[i : i+lanes])

		if !firstInitialized {
			minVec = v
			firstInitialized = true
		} else {
			minVec = minVec.Min(v)
		}
	}

	// Find minimum in the vector (only if we processed any vectors)
	var minVal float64
	if firstInitialized {
		var buf [lanes]float64
		minVec.Store(&buf)
		minVal = min(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] < T(minVal) {
			minVal = float64(collection[i])
			firstInitialized = true
		}
	}

	return T(minVal)
}

// MaxInt8x64 finds the maximum value in a collection of int8 using AVX-512 SIMD
func MaxInt8x64[T ~int8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes64

	base := unsafeSliceInt8(collection, length)

	var maxVec archsimd.Int8x64
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x64Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal int8
	if firstInitialized {
		var buf [lanes]int8
		maxVec.Store(&buf)
		maxVal = max(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
			buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23],
			buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31],
			buf[32], buf[33], buf[34], buf[35], buf[36], buf[37], buf[38], buf[39],
			buf[40], buf[41], buf[42], buf[43], buf[44], buf[45], buf[46], buf[47],
			buf[48], buf[49], buf[50], buf[51], buf[52], buf[53], buf[54], buf[55],
			buf[56], buf[57], buf[58], buf[59], buf[60], buf[61], buf[62], buf[63],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = int8(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxInt16x32 finds the maximum value in a collection of int16 using AVX-512 SIMD
func MaxInt16x32[T ~int16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes32

	base := unsafeSliceInt16(collection, length)

	var maxVec archsimd.Int16x32
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x32Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal int16
	if firstInitialized {
		var buf [lanes]int16
		maxVec.Store(&buf)
		maxVal = max(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
			buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23],
			buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = int16(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxInt32x16 finds the maximum value in a collection of int32 using AVX-512 SIMD
func MaxInt32x16[T ~int32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16

	base := unsafeSliceInt32(collection, length)

	var maxVec archsimd.Int32x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x16Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal int32
	if firstInitialized {
		var buf [lanes]int32
		maxVec.Store(&buf)
		maxVal = max(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = int32(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxInt64x2 finds the maximum value in a collection of int64 using AVX-512 SIMD.
// Int64x2 Max operations in archsimd require AVX-512.
func MaxInt64x2[T ~int64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes2
	base := unsafeSliceInt64(collection, length)

	var maxVec archsimd.Int64x2
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x2Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal int64
	if firstInitialized {
		var buf [lanes]int64
		maxVec.Store(&buf)
		maxVal = max(buf[0], buf[1])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = int64(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxUint64x2 finds the maximum value in a collection of uint64 using AVX-512 SIMD.
// Uint64x2 Max operations in archsimd require AVX-512.
func MaxUint64x2[T ~uint64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes2
	base := unsafeSliceUint64(collection, length)

	var maxVec archsimd.Uint64x2
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x2Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal uint64
	if firstInitialized {
		var buf [lanes]uint64
		maxVec.Store(&buf)
		maxVal = max(buf[0], buf[1])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = uint64(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxInt64x8 finds the maximum value in a collection of int64 using AVX-512 SIMD
func MaxInt64x8[T ~int64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8

	base := unsafeSliceInt64(collection, length)

	var maxVec archsimd.Int64x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x8Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal int64
	if firstInitialized {
		var buf [lanes]int64
		maxVec.Store(&buf)
		maxVal = max(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = int64(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxUint8x64 finds the maximum value in a collection of uint8 using AVX-512 SIMD
func MaxUint8x64[T ~uint8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes64

	base := unsafeSliceUint8(collection, length)

	var maxVec archsimd.Uint8x64
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x64Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal uint8
	if firstInitialized {
		var buf [lanes]uint8
		maxVec.Store(&buf)
		maxVal = max(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
			buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23],
			buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31],
			buf[32], buf[33], buf[34], buf[35], buf[36], buf[37], buf[38], buf[39],
			buf[40], buf[41], buf[42], buf[43], buf[44], buf[45], buf[46], buf[47],
			buf[48], buf[49], buf[50], buf[51], buf[52], buf[53], buf[54], buf[55],
			buf[56], buf[57], buf[58], buf[59], buf[60], buf[61], buf[62], buf[63],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = uint8(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxUint16x32 finds the maximum value in a collection of uint16 using AVX-512 SIMD
func MaxUint16x32[T ~uint16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes32

	base := unsafeSliceUint16(collection, length)

	var maxVec archsimd.Uint16x32
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x32Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal uint16
	if firstInitialized {
		var buf [lanes]uint16
		maxVec.Store(&buf)
		maxVal = max(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
			buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23],
			buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = uint16(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxUint32x16 finds the maximum value in a collection of uint32 using AVX-512 SIMD
func MaxUint32x16[T ~uint32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16

	base := unsafeSliceUint32(collection, length)

	var maxVec archsimd.Uint32x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x16Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal uint32
	if firstInitialized {
		var buf [lanes]uint32
		maxVec.Store(&buf)
		maxVal = max(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = uint32(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxUint64x8 finds the maximum value in a collection of uint64 using AVX-512 SIMD
func MaxUint64x8[T ~uint64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8

	base := unsafeSliceUint64(collection, length)

	var maxVec archsimd.Uint64x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x8Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal uint64
	if firstInitialized {
		var buf [lanes]uint64
		maxVec.Store(&buf)
		maxVal = max(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = uint64(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxFloat32x16 finds the maximum value in a collection of float32 using AVX-512 SIMD
func MaxFloat32x16[T ~float32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16

	base := unsafeSliceFloat32(collection, length)

	var maxVec archsimd.Float32x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x16Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal float32
	if firstInitialized {
		var buf [lanes]float32
		maxVec.Store(&buf)
		maxVal = max(
			buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7],
			buf[8], buf[9], buf[10], buf[11], buf[12], buf[13], buf[14], buf[15],
		)
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = float32(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// MaxFloat64x8 finds the maximum value in a collection of float64 using AVX-512 SIMD
func MaxFloat64x8[T ~float64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8

	base := unsafeSliceFloat64(collection, length)

	var maxVec archsimd.Float64x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x8Slice(base[i : i+lanes])

		if !firstInitialized {
			maxVec = v
			firstInitialized = true
		} else {
			maxVec = maxVec.Max(v)
		}
	}

	// Find maximum in the vector (only if we processed any vectors)
	var maxVal float64
	if firstInitialized {
		var buf [lanes]float64
		maxVec.Store(&buf)
		maxVal = max(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if !firstInitialized || collection[i] > T(maxVal) {
			maxVal = float64(collection[i])
			firstInitialized = true
		}
	}

	return T(maxVal)
}

// AVX-512 (512-bit) SIMD sumBy functions - 64/32/16/8 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD sum functions.

// SumByInt8x64 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByInt8x64[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumInt8x64(mapped)
}

// SumByInt16x32 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByInt16x32[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumInt16x32(mapped)
}

// SumByInt32x16 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByInt32x16[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumInt32x16(mapped)
}

// SumByInt64x8 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByInt64x8[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumInt64x8(mapped)
}

// SumByUint8x64 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByUint8x64[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumUint8x64(mapped)
}

// SumByUint16x32 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByUint16x32[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumUint16x32(mapped)
}

// SumByUint32x16 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByUint32x16[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumUint32x16(mapped)
}

// SumByUint64x8 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByUint64x8[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumUint64x8(mapped)
}

// SumByFloat32x16 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByFloat32x16[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumFloat32x16(mapped)
}

// SumByFloat64x8 sums the values extracted by iteratee from a slice using AVX-512 SIMD.
func SumByFloat64x8[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumFloat64x8(mapped)
}

// AVX-512 (512-bit) SIMD meanBy functions - 64/32/16/8 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD mean functions.

// MeanByInt8x64 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByInt8x64[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanInt8x64(mapped)
}

// MeanByInt16x32 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByInt16x32[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanInt16x32(mapped)
}

// MeanByInt32x16 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByInt32x16[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanInt32x16(mapped)
}

// MeanByInt64x8 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByInt64x8[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanInt64x8(mapped)
}

// MeanByUint8x64 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByUint8x64[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanUint8x64(mapped)
}

// MeanByUint16x32 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByUint16x32[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanUint16x32(mapped)
}

// MeanByUint32x16 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByUint32x16[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanUint32x16(mapped)
}

// MeanByUint64x8 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByUint64x8[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanUint64x8(mapped)
}

// MeanByFloat32x16 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByFloat32x16[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanFloat32x16(mapped)
}

// MeanByFloat64x8 calculates the mean of values extracted by iteratee from a slice using AVX-512 SIMD.
func MeanByFloat64x8[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanFloat64x8(mapped)
}
