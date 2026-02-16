//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// AVX2 (256-bit) SIMD sum functions - 32/16/8/4 lanes

// SumInt8x32 sums a slice of int8 using AVX2 SIMD (Int8x32, 32 lanes)
func SumInt8x32[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Int8x32

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x32Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [32]int8
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumInt16x16 sums a slice of int16 using AVX2 SIMD (Int16x16, 16 lanes)
func SumInt16x16[T ~int16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Int16x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x16Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [16]int16
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumInt32x8 sums a slice of int32 using AVX2 SIMD (Int32x8, 8 lanes)
func SumInt32x8[T ~int32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Int32x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [8]int32
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumInt64x4 sums a slice of int64 using AVX2 SIMD (Int64x4, 4 lanes)
func SumInt64x4[T ~int64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Int64x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x4Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [4]int64
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumUint8x32 sums a slice of uint8 using AVX2 SIMD (Uint8x32, 32 lanes)
func SumUint8x32[T ~uint8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Uint8x32

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x32Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [32]uint8
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumUint16x16 sums a slice of uint16 using AVX2 SIMD (Uint16x16, 16 lanes)
func SumUint16x16[T ~uint16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Uint16x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x16Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [16]uint16
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumUint32x8 sums a slice of uint32 using AVX2 SIMD (Uint32x8, 8 lanes)
func SumUint32x8[T ~uint32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Uint32x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [8]uint32
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumUint64x4 sums a slice of uint64 using AVX2 SIMD (Uint64x4, 4 lanes)
func SumUint64x4[T ~uint64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Uint64x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x4Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [4]uint64
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumFloat32x8 sums a slice of float32 using AVX2 SIMD (Float32x8, 8 lanes)
func SumFloat32x8[T ~float32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Float32x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [8]float32
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// SumFloat64x4 sums a slice of float64 using AVX2 SIMD (Float64x4, 4 lanes)
func SumFloat64x4[T ~float64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Float64x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x4Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [4]float64
	acc.Store(&buf)
	var sum T
	for k := 0; k < lanes; k++ {
		sum += T(buf[k])
	}

	for ; i < length; i++ {
		sum += collection[i]
	}

	return sum
}

// MeanInt8x32 calculates the mean of a slice of int8 using AVX2 SIMD
func MeanInt8x32[T ~int8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt8x32(collection)
	return sum / T(len(collection))
}

// MeanInt16x16 calculates the mean of a slice of int16 using AVX2 SIMD
func MeanInt16x16[T ~int16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt16x16(collection)
	return sum / T(len(collection))
}

// MeanInt32x8 calculates the mean of a slice of int32 using AVX2 SIMD
func MeanInt32x8[T ~int32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt32x8(collection)
	return sum / T(len(collection))
}

// MeanInt64x4 calculates the mean of a slice of int64 using AVX2 SIMD
func MeanInt64x4[T ~int64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt64x4(collection)
	return sum / T(len(collection))
}

// MeanUint8x32 calculates the mean of a slice of uint8 using AVX2 SIMD
func MeanUint8x32[T ~uint8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint8x32(collection)
	return sum / T(len(collection))
}

// MeanUint16x16 calculates the mean of a slice of uint16 using AVX2 SIMD
func MeanUint16x16[T ~uint16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint16x16(collection)
	return sum / T(len(collection))
}

// MeanUint32x8 calculates the mean of a slice of uint32 using AVX2 SIMD
func MeanUint32x8[T ~uint32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint32x8(collection)
	return sum / T(len(collection))
}

// MeanUint64x4 calculates the mean of a slice of uint64 using AVX2 SIMD
func MeanUint64x4[T ~uint64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint64x4(collection)
	return sum / T(len(collection))
}

// MeanFloat32x8 calculates the mean of a slice of float32 using AVX2 SIMD
func MeanFloat32x8[T ~float32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumFloat32x8(collection)
	return sum / T(len(collection))
}

// MeanFloat64x4 calculates the mean of a slice of float64 using AVX2 SIMD
func MeanFloat64x4[T ~float64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumFloat64x4(collection)
	return sum / T(len(collection))
}

// ClampInt8x32 clamps each element in collection between min and max values using AVX2 SIMD
func ClampInt8x32[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastInt8x32(int8(min))
	maxVec := archsimd.BroadcastInt8x32(int8(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadInt8x32Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [32]int8
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampInt16x16 clamps each element in collection between min and max values using AVX2 SIMD
func ClampInt16x16[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastInt16x16(int16(min))
	maxVec := archsimd.BroadcastInt16x16(int16(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadInt16x16Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [16]int16
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampInt32x8 clamps each element in collection between min and max values using AVX2 SIMD
func ClampInt32x8[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastInt32x8(int32(min))
	maxVec := archsimd.BroadcastInt32x8(int32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadInt32x8Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [8]int32
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampInt64x4 clamps each element in collection between min and max values using AVX2 SIMD and AVX-512 SIMD.
func ClampInt64x4[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastInt64x4(int64(min))
	maxVec := archsimd.BroadcastInt64x4(int64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadInt64x4Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [4]int64
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampUint8x32 clamps each element in collection between min and max values using AVX2 SIMD
func ClampUint8x32[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastUint8x32(uint8(min))
	maxVec := archsimd.BroadcastUint8x32(uint8(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadUint8x32Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [32]uint8
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampUint16x16 clamps each element in collection between min and max values using AVX2 SIMD
func ClampUint16x16[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastUint16x16(uint16(min))
	maxVec := archsimd.BroadcastUint16x16(uint16(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadUint16x16Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [16]uint16
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampUint32x8 clamps each element in collection between min and max values using AVX2 SIMD
func ClampUint32x8[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastUint32x8(uint32(min))
	maxVec := archsimd.BroadcastUint32x8(uint32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadUint32x8Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [8]uint32
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampUint64x4 clamps each element in collection between min and max values using AVX2 SIMD and AVX-512 SIMD.
func ClampUint64x4[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastUint64x4(uint64(min))
	maxVec := archsimd.BroadcastUint64x4(uint64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadUint64x4Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [4]uint64
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampFloat32x8 clamps each element in collection between min and max values using AVX2 SIMD
func ClampFloat32x8[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastFloat32x8(float32(min))
	maxVec := archsimd.BroadcastFloat32x8(float32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadFloat32x8Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [8]float32
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// ClampFloat64x4 clamps each element in collection between min and max values using AVX2 SIMD
func ClampFloat64x4[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), len(collection))
	minVec := archsimd.BroadcastFloat64x4(float64(min))
	maxVec := archsimd.BroadcastFloat64x4(float64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		v := archsimd.LoadFloat64x4Slice(base[i : i+lanes])

		clamped := v.Max(minVec).Min(maxVec)

		var buf [4]float64
		clamped.Store(&buf)
		for j := 0; j < lanes; j++ {
			result[i+j] = T(buf[j])
		}
	}

	for ; i < len(collection); i++ {
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

// MinInt8x32 finds the minimum value in a collection of int8 using AVX2 SIMD
func MinInt8x32[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Int8x32
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x32Slice(base[i : i+lanes])

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
		var buf [32]int8
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinInt16x16 finds the minimum value in a collection of int16 using AVX2 SIMD
func MinInt16x16[T ~int16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Int16x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x16Slice(base[i : i+lanes])

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
		var buf [16]int16
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinInt32x8 finds the minimum value in a collection of int32 using AVX2 SIMD
func MinInt32x8[T ~int32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Int32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x8Slice(base[i : i+lanes])

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
		var buf [8]int32
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinInt64x4 finds the minimum value in a collection of int64 using AVX2 SIMD
func MinInt64x4[T ~int64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Int64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x4Slice(base[i : i+lanes])

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
		var buf [4]int64
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinUint8x32 finds the minimum value in a collection of uint8 using AVX2 SIMD
func MinUint8x32[T ~uint8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Uint8x32
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x32Slice(base[i : i+lanes])

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
		var buf [32]uint8
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinUint16x16 finds the minimum value in a collection of uint16 using AVX2 SIMD
func MinUint16x16[T ~uint16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Uint16x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x16Slice(base[i : i+lanes])

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
		var buf [16]uint16
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinUint32x8 finds the minimum value in a collection of uint32 using AVX2 SIMD
func MinUint32x8[T ~uint32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Uint32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x8Slice(base[i : i+lanes])

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
		var buf [8]uint32
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinUint64x4 finds the minimum value in a collection of uint64 using AVX2 SIMD
func MinUint64x4[T ~uint64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Uint64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x4Slice(base[i : i+lanes])

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
		var buf [4]uint64
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinFloat32x8 finds the minimum value in a collection of float32 using AVX2 SIMD
func MinFloat32x8[T ~float32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Float32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x8Slice(base[i : i+lanes])

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
		var buf [8]float32
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MinFloat64x4 finds the minimum value in a collection of float64 using AVX2 SIMD
func MinFloat64x4[T ~float64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), length)
	var minVec archsimd.Float64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x4Slice(base[i : i+lanes])

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
		var buf [4]float64
		minVec.Store(&buf)
		minVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] < minVal {
				minVal = buf[j]
			}
		}
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

// MaxInt8x32 finds the maximum value in a collection of int8 using AVX2 SIMD
func MaxInt8x32[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Int8x32
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x32Slice(base[i : i+lanes])

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
		var buf [32]int8
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxInt16x16 finds the maximum value in a collection of int16 using AVX2 SIMD
func MaxInt16x16[T ~int16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Int16x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x16Slice(base[i : i+lanes])

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
		var buf [16]int16
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxInt32x8 finds the maximum value in a collection of int32 using AVX2 SIMD
func MaxInt32x8[T ~int32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Int32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x8Slice(base[i : i+lanes])

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
		var buf [8]int32
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxInt64x4 finds the maximum value in a collection of int64 using AVX2 SIMD
func MaxInt64x4[T ~int64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Int64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x4Slice(base[i : i+lanes])

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
		var buf [4]int64
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxUint8x32 finds the maximum value in a collection of uint8 using AVX2 SIMD
func MaxUint8x32[T ~uint8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Uint8x32
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x32Slice(base[i : i+lanes])

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
		var buf [32]uint8
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxUint16x16 finds the maximum value in a collection of uint16 using AVX2 SIMD
func MaxUint16x16[T ~uint16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Uint16x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x16Slice(base[i : i+lanes])

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
		var buf [16]uint16
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxUint32x8 finds the maximum value in a collection of uint32 using AVX2 SIMD
func MaxUint32x8[T ~uint32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Uint32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x8Slice(base[i : i+lanes])

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
		var buf [8]uint32
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxUint64x4 finds the maximum value in a collection of uint64 using AVX2 SIMD
func MaxUint64x4[T ~uint64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Uint64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x4Slice(base[i : i+lanes])

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
		var buf [4]uint64
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxFloat32x8 finds the maximum value in a collection of float32 using AVX2 SIMD
func MaxFloat32x8[T ~float32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Float32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x8Slice(base[i : i+lanes])

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
		var buf [8]float32
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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

// MaxFloat64x4 finds the maximum value in a collection of float64 using AVX2 SIMD
func MaxFloat64x4[T ~float64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), length)
	var maxVec archsimd.Float64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x4Slice(base[i : i+lanes])

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
		var buf [4]float64
		maxVec.Store(&buf)
		maxVal = buf[0]
		for j := 1; j < lanes; j++ {
			if buf[j] > maxVal {
				maxVal = buf[j]
			}
		}
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
