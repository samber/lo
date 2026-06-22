//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"

	"github.com/samber/lo"
)

// AVX (128-bit) SIMD sum functions - 16/8/4/2 lanes

// SumInt8x16 sums a slice of int8 using AVX SIMD (Int8x16, 16 lanes).
// Overflow: The accumulation is performed using int8, which can overflow for large collections.
// If the sum exceeds the int8 range (-128 to 127), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt8x16[T ~int8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes16

	base := unsafeSliceInt8(collection, length)
	var acc archsimd.Int8x16

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x16Slice(base[i : i+lanes])
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

// SumInt16x8 sums a slice of int16 using AVX SIMD (Int16x8, 8 lanes).
// Overflow: The accumulation is performed using int16, which can overflow for large collections.
// If the sum exceeds the int16 range (-32768 to 32767), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumInt16x8[T ~int16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes8

	base := unsafeSliceInt16(collection, length)
	var acc archsimd.Int16x8

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x8Slice(base[i : i+lanes])
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

// SumInt32x4 sums a slice of int32 using AVX SIMD (Int32x4, 4 lanes).
// Overflow: The accumulation is performed using int32, which can overflow for very large collections.
// If the sum exceeds the int32 range (-2147483648 to 2147483647), the result will wrap around silently.
// For collections that may overflow, consider using SumInt64x2 or handle overflow detection externally.
func SumInt32x4[T ~int32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes4

	base := unsafeSliceInt32(collection, length)
	var acc archsimd.Int32x4

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x4Slice(base[i : i+lanes])
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

// SumInt64x2 sums a slice of int64 using AVX SIMD (Int64x2, 2 lanes).
// Overflow: The accumulation is performed using int64, which can overflow for extremely large collections.
// If the sum exceeds the int64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumInt64x2[T ~int64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes2

	base := unsafeSliceInt64(collection, length)
	var acc archsimd.Int64x2

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x2Slice(base[i : i+lanes])
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

// SumUint8x16 sums a slice of uint8 using AVX SIMD (Uint8x16, 16 lanes).
// Overflow: The accumulation is performed using uint8, which can overflow for large collections.
// If the sum exceeds the uint8 range (0 to 255), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint8x16[T ~uint8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes16

	base := unsafeSliceUint8(collection, length)
	var acc archsimd.Uint8x16

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x16Slice(base[i : i+lanes])
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

// SumUint16x8 sums a slice of uint16 using AVX SIMD (Uint16x8, 8 lanes).
// Overflow: The accumulation is performed using uint16, which can overflow for large collections.
// If the sum exceeds the uint16 range (0 to 65535), the result will wrap around silently.
// For collections that may overflow, consider using a wider type or handle overflow detection externally.
func SumUint16x8[T ~uint16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes8

	base := unsafeSliceUint16(collection, length)
	var acc archsimd.Uint16x8

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x8Slice(base[i : i+lanes])
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

// SumUint32x4 sums a slice of uint32 using AVX SIMD (Uint32x4, 4 lanes).
// Overflow: The accumulation is performed using uint32, which can overflow for very large collections.
// If the sum exceeds the uint32 range (0 to 4294967295), the result will wrap around silently.
// For collections that may overflow, consider using SumUint64x2 or handle overflow detection externally.
func SumUint32x4[T ~uint32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes4

	base := unsafeSliceUint32(collection, length)
	var acc archsimd.Uint32x4

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x4Slice(base[i : i+lanes])
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

// SumUint64x2 sums a slice of uint64 using AVX SIMD (Uint64x2, 2 lanes).
// Overflow: The accumulation is performed using uint64, which can overflow for extremely large collections.
// If the sum exceeds the uint64 range, the result will wrap around silently.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Int).
func SumUint64x2[T ~uint64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes2

	base := unsafeSliceUint64(collection, length)
	var acc archsimd.Uint64x2

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x2Slice(base[i : i+lanes])
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

// SumFloat32x4 sums a slice of float32 using AVX SIMD (Float32x4, 4 lanes).
// Overflow: The accumulation is performed using float32. Overflow will result in +/-Inf rather than wrapping.
// For collections requiring high precision or large sums, consider using SumFloat64x2.
func SumFloat32x4[T ~float32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes4

	base := unsafeSliceFloat32(collection, length)
	var acc archsimd.Float32x4

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x4Slice(base[i : i+lanes])
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

// SumFloat64x2 sums a slice of float64 using AVX SIMD (Float64x2, 2 lanes).
// Overflow: The accumulation is performed using float64. Overflow will result in +/-Inf rather than wrapping.
// For collections that may overflow, handle overflow detection externally (e.g., using big.Float).
func SumFloat64x2[T ~float64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	const lanes = simdLanes2

	base := unsafeSliceFloat64(collection, length)
	var acc archsimd.Float64x2

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x2Slice(base[i : i+lanes])
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

// MeanInt8x16 calculates the mean of a slice of int8 using AVX SIMD
func MeanInt8x16[T ~int8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumInt8x16(collection)
	return sum / T(length)
}

// MeanInt16x8 calculates the mean of a slice of int16 using AVX SIMD
func MeanInt16x8[T ~int16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumInt16x8(collection)
	return sum / T(length)
}

// MeanInt32x4 calculates the mean of a slice of int32 using AVX SIMD
func MeanInt32x4[T ~int32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumInt32x4(collection)
	return sum / T(length)
}

// MeanInt64x2 calculates the mean of a slice of int64 using AVX SIMD
func MeanInt64x2[T ~int64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumInt64x2(collection)
	return sum / T(length)
}

// MeanUint8x16 calculates the mean of a slice of uint8 using AVX SIMD
func MeanUint8x16[T ~uint8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumUint8x16(collection)
	return sum / T(length)
}

// MeanUint16x8 calculates the mean of a slice of uint16 using AVX SIMD
func MeanUint16x8[T ~uint16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumUint16x8(collection)
	return sum / T(length)
}

// MeanUint32x4 calculates the mean of a slice of uint32 using AVX SIMD
func MeanUint32x4[T ~uint32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumUint32x4(collection)
	return sum / T(length)
}

// MeanUint64x2 calculates the mean of a slice of uint64 using AVX SIMD
func MeanUint64x2[T ~uint64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumUint64x2(collection)
	return sum / T(length)
}

// MeanFloat32x4 calculates the mean of a slice of float32 using AVX SIMD
func MeanFloat32x4[T ~float32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumFloat32x4(collection)

	return sum / T(length)
}

// MeanFloat64x2 calculates the mean of a slice of float64 using AVX SIMD
func MeanFloat64x2[T ~float64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}
	sum := SumFloat64x2(collection)
	return sum / T(length)
}

// ClampInt8x16 clamps each element in collection between min and max values using AVX SIMD
func ClampInt8x16[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes16

	base := unsafeSliceInt8(collection, length)

	minVec := archsimd.BroadcastInt8x16(int8(min))
	maxVec := archsimd.BroadcastInt8x16(int8(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x16Slice(base[i : i+lanes])

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

// ClampInt16x8 clamps each element in collection between min and max values using AVX SIMD
func ClampInt16x8[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes8

	base := unsafeSliceInt16(collection, length)

	minVec := archsimd.BroadcastInt16x8(int16(min))
	maxVec := archsimd.BroadcastInt16x8(int16(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x8Slice(base[i : i+lanes])

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

// ClampInt32x4 clamps each element in collection between min and max values using AVX SIMD
func ClampInt32x4[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes4

	base := unsafeSliceInt32(collection, length)

	minVec := archsimd.BroadcastInt32x4(int32(min))
	maxVec := archsimd.BroadcastInt32x4(int32(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x4Slice(base[i : i+lanes])

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

// ClampUint8x16 clamps each element in collection between min and max values using AVX SIMD
func ClampUint8x16[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes16

	base := unsafeSliceUint8(collection, length)

	minVec := archsimd.BroadcastUint8x16(uint8(min))
	maxVec := archsimd.BroadcastUint8x16(uint8(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x16Slice(base[i : i+lanes])

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

// ClampUint16x8 clamps each element in collection between min and max values using AVX SIMD
func ClampUint16x8[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes8

	base := unsafeSliceUint16(collection, length)

	minVec := archsimd.BroadcastUint16x8(uint16(min))
	maxVec := archsimd.BroadcastUint16x8(uint16(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x8Slice(base[i : i+lanes])

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

// ClampUint32x4 clamps each element in collection between min and max values using AVX SIMD
func ClampUint32x4[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes4

	base := unsafeSliceUint32(collection, length)

	minVec := archsimd.BroadcastUint32x4(uint32(min))
	maxVec := archsimd.BroadcastUint32x4(uint32(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x4Slice(base[i : i+lanes])

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

// ClampFloat32x4 clamps each element in collection between min and max values using AVX SIMD
func ClampFloat32x4[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes4

	base := unsafeSliceFloat32(collection, length)

	minVec := archsimd.BroadcastFloat32x4(float32(min))
	maxVec := archsimd.BroadcastFloat32x4(float32(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x4Slice(base[i : i+lanes])

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

// ClampFloat64x2 clamps each element in collection between min and max values using AVX SIMD
func ClampFloat64x2[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	length := uint(len(collection))
	if length == 0 {
		return collection
	}

	result := make(Slice, length)
	const lanes = simdLanes2

	base := unsafeSliceFloat64(collection, length)

	minVec := archsimd.BroadcastFloat64x2(float64(min))
	maxVec := archsimd.BroadcastFloat64x2(float64(max))

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x2Slice(base[i : i+lanes])

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

// MinInt8x16 finds the minimum value in a collection of int8 using AVX SIMD
func MinInt8x16[T ~int8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16
	base := unsafeSliceInt8(collection, length)

	var minVec archsimd.Int8x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x16Slice(base[i : i+lanes])

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

// MinInt16x8 finds the minimum value in a collection of int16 using AVX SIMD
func MinInt16x8[T ~int16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8
	base := unsafeSliceInt16(collection, length)

	var minVec archsimd.Int16x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x8Slice(base[i : i+lanes])

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
		minVal = min(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
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

// MinInt32x4 finds the minimum value in a collection of int32 using AVX SIMD
func MinInt32x4[T ~int32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes4
	base := unsafeSliceInt32(collection, length)

	var minVec archsimd.Int32x4
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x4Slice(base[i : i+lanes])

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
		minVal = min(buf[0], buf[1], buf[2], buf[3])
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

// MinUint8x16 finds the minimum value in a collection of uint8 using AVX SIMD
func MinUint8x16[T ~uint8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16
	base := unsafeSliceUint8(collection, length)

	var minVec archsimd.Uint8x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x16Slice(base[i : i+lanes])

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

// MinUint16x8 finds the minimum value in a collection of uint16 using AVX SIMD
func MinUint16x8[T ~uint16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8
	base := unsafeSliceUint16(collection, length)

	var minVec archsimd.Uint16x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x8Slice(base[i : i+lanes])

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
		minVal = min(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
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

// MinUint32x4 finds the minimum value in a collection of uint32 using AVX SIMD
func MinUint32x4[T ~uint32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes4
	base := unsafeSliceUint32(collection, length)

	var minVec archsimd.Uint32x4
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x4Slice(base[i : i+lanes])

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
		minVal = min(buf[0], buf[1], buf[2], buf[3])
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

// MinFloat32x4 finds the minimum value in a collection of float32 using AVX SIMD
func MinFloat32x4[T ~float32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes4
	base := unsafeSliceFloat32(collection, length)

	var minVec archsimd.Float32x4
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x4Slice(base[i : i+lanes])

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
		minVal = min(buf[0], buf[1], buf[2], buf[3])
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

// MinFloat64x2 finds the minimum value in a collection of float64 using AVX SIMD
func MinFloat64x2[T ~float64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes2
	base := unsafeSliceFloat64(collection, length)

	var minVec archsimd.Float64x2
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x2Slice(base[i : i+lanes])

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
		minVal = min(buf[0], buf[1])
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

// MaxInt8x16 finds the maximum value in a collection of int8 using AVX SIMD
func MaxInt8x16[T ~int8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16
	base := unsafeSliceInt8(collection, length)

	var maxVec archsimd.Int8x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x16Slice(base[i : i+lanes])

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

// MaxInt16x8 finds the maximum value in a collection of int16 using AVX SIMD
func MaxInt16x8[T ~int16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8
	base := unsafeSliceInt16(collection, length)

	var maxVec archsimd.Int16x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x8Slice(base[i : i+lanes])

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
		maxVal = max(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
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

// MaxInt32x4 finds the maximum value in a collection of int32 using AVX SIMD
func MaxInt32x4[T ~int32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes4
	base := unsafeSliceInt32(collection, length)

	var maxVec archsimd.Int32x4
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x4Slice(base[i : i+lanes])

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
		maxVal = max(buf[0], buf[1], buf[2], buf[3])
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

// MaxUint8x16 finds the maximum value in a collection of uint8 using AVX SIMD
func MaxUint8x16[T ~uint8](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes16
	base := unsafeSliceUint8(collection, length)

	var maxVec archsimd.Uint8x16
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x16Slice(base[i : i+lanes])

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

// MaxUint16x8 finds the maximum value in a collection of uint16 using AVX SIMD
func MaxUint16x8[T ~uint16](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes8
	base := unsafeSliceUint16(collection, length)

	var maxVec archsimd.Uint16x8
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x8Slice(base[i : i+lanes])

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
		maxVal = max(buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7])
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

// MaxUint32x4 finds the maximum value in a collection of uint32 using AVX SIMD
func MaxUint32x4[T ~uint32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes4
	base := unsafeSliceUint32(collection, length)

	var maxVec archsimd.Uint32x4
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x4Slice(base[i : i+lanes])

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
		maxVal = max(buf[0], buf[1], buf[2], buf[3])
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

// MaxFloat32x4 finds the maximum value in a collection of float32 using AVX SIMD
func MaxFloat32x4[T ~float32](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes4
	base := unsafeSliceFloat32(collection, length)

	var maxVec archsimd.Float32x4
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x4Slice(base[i : i+lanes])

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
		maxVal = max(buf[0], buf[1], buf[2], buf[3])
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

// MaxFloat64x2 finds the maximum value in a collection of float64 using AVX SIMD
func MaxFloat64x2[T ~float64](collection []T) T {
	length := uint(len(collection))
	if length == 0 {
		return 0
	}

	const lanes = simdLanes2
	base := unsafeSliceFloat64(collection, length)

	var maxVec archsimd.Float64x2
	firstInitialized := false

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x2Slice(base[i : i+lanes])

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
		maxVal = max(buf[0], buf[1])
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

// AVX (128-bit) SIMD sumBy functions - 16/8/4/2 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD sum functions.

// SumByInt8x16 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByInt8x16[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumInt8x16(mapped)
}

// SumByInt16x8 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByInt16x8[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumInt16x8(mapped)
}

// SumByInt32x4 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByInt32x4[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumInt32x4(mapped)
}

// SumByInt64x2 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByInt64x2[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumInt64x2(mapped)
}

// SumByUint8x16 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByUint8x16[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumUint8x16(mapped)
}

// SumByUint16x8 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByUint16x8[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumUint16x8(mapped)
}

// SumByUint32x4 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByUint32x4[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumUint32x4(mapped)
}

// SumByUint64x2 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByUint64x2[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumUint64x2(mapped)
}

// SumByFloat32x4 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByFloat32x4[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumFloat32x4(mapped)
}

// SumByFloat64x2 sums the values extracted by iteratee from a slice using AVX SIMD.
func SumByFloat64x2[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return SumFloat64x2(mapped)
}

// AVX (128-bit) SIMD meanBy functions - 16/8/4/2 lanes
// These implementations use lo.Map to apply the iteratee, then chain with SIMD mean functions.

// MeanByInt8x16 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByInt8x16[T any, R ~int8](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanInt8x16(mapped)
}

// MeanByInt16x8 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByInt16x8[T any, R ~int16](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanInt16x8(mapped)
}

// MeanByInt32x4 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByInt32x4[T any, R ~int32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanInt32x4(mapped)
}

// MeanByInt64x2 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByInt64x2[T any, R ~int64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanInt64x2(mapped)
}

// MeanByUint8x16 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByUint8x16[T any, R ~uint8](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanUint8x16(mapped)
}

// MeanByUint16x8 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByUint16x8[T any, R ~uint16](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanUint16x8(mapped)
}

// MeanByUint32x4 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByUint32x4[T any, R ~uint32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanUint32x4(mapped)
}

// MeanByUint64x2 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByUint64x2[T any, R ~uint64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanUint64x2(mapped)
}

// MeanByFloat32x4 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByFloat32x4[T any, R ~float32](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanFloat32x4(mapped)
}

// MeanByFloat64x2 calculates the mean of values extracted by iteratee from a slice using AVX SIMD.
func MeanByFloat64x2[T any, R ~float64](collection []T, iteratee func(item T) R) R {
	mapped := lo.Map(collection, func(item T, _ int) R { return iteratee(item) })
	return MeanFloat64x2(mapped)
}
