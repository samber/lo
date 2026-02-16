//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// AVX-512 (512-bit) SIMD sum functions - 64/32/16/8 lanes

// SumInt8x64 sums a slice of int8 using AVX-512 SIMD (Int8x64, 64 lanes)
func SumInt8x64[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 64

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Int8x64

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt8x64Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [64]int8
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

// SumInt16x32 sums a slice of int16 using AVX-512 SIMD (Int16x32, 32 lanes)
func SumInt16x32[T ~int16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Int16x32

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt16x32Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [32]int16
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

// SumInt32x16 sums a slice of int32 using AVX-512 SIMD (Int32x16, 16 lanes)
func SumInt32x16[T ~int32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Int32x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt32x16Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [16]int32
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

// SumInt64x8 sums a slice of int64 using AVX-512 SIMD (Int64x8, 8 lanes)
func SumInt64x8[T ~int64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Int64x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadInt64x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [8]int64
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

// SumUint8x64 sums a slice of uint8 using AVX-512 SIMD (Uint8x64, 64 lanes)
func SumUint8x64[T ~uint8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 64

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Uint8x64

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint8x64Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [64]uint8
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

// SumUint16x32 sums a slice of uint16 using AVX-512 SIMD (Uint16x32, 32 lanes)
func SumUint16x32[T ~uint16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Uint16x32

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint16x32Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [32]uint16
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

// SumUint32x16 sums a slice of uint32 using AVX-512 SIMD (Uint32x16, 16 lanes)
func SumUint32x16[T ~uint32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Uint32x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint32x16Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [16]uint32
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

// SumUint64x8 sums a slice of uint64 using AVX-512 SIMD (Uint64x8, 8 lanes)
func SumUint64x8[T ~uint64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Uint64x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadUint64x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [8]uint64
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

// SumFloat32x16 sums a slice of float32 using AVX-512 SIMD (Float32x16, 16 lanes)
func SumFloat32x16[T ~float32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Float32x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat32x16Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [16]float32
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

// SumFloat64x8 sums a slice of float64 using AVX-512 SIMD (Float64x8, 8 lanes)
func SumFloat64x8[T ~float64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}
	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), length)
	var acc archsimd.Float64x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		v := archsimd.LoadFloat64x8Slice(base[i : i+lanes])
		acc = acc.Add(v)
	}

	var buf [8]float64
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

// MeanInt8x64 calculates the mean of a slice of int8 using AVX-512 SIMD
func MeanInt8x64[T ~int8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt8x64(collection)
	return sum / T(len(collection))
}

// MeanInt16x32 calculates the mean of a slice of int16 using AVX-512 SIMD
func MeanInt16x32[T ~int16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt16x32(collection)
	return sum / T(len(collection))
}

// MeanInt32x16 calculates the mean of a slice of int32 using AVX-512 SIMD
func MeanInt32x16[T ~int32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt32x16(collection)
	return sum / T(len(collection))
}

// MeanInt64x8 calculates the mean of a slice of int64 using AVX-512 SIMD
func MeanInt64x8[T ~int64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt64x8(collection)
	return sum / T(len(collection))
}

// MeanUint8x64 calculates the mean of a slice of uint8 using AVX-512 SIMD
func MeanUint8x64[T ~uint8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint8x64(collection)
	return sum / T(len(collection))
}

// MeanUint16x32 calculates the mean of a slice of uint16 using AVX-512 SIMD
func MeanUint16x32[T ~uint16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint16x32(collection)
	return sum / T(len(collection))
}

// MeanUint32x16 calculates the mean of a slice of uint32 using AVX-512 SIMD
func MeanUint32x16[T ~uint32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint32x16(collection)
	return sum / T(len(collection))
}

// MeanUint64x8 calculates the mean of a slice of uint64 using AVX-512 SIMD
func MeanUint64x8[T ~uint64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint64x8(collection)
	return sum / T(len(collection))
}

// MeanFloat32x16 calculates the mean of a slice of float32 using AVX-512 SIMD
func MeanFloat32x16[T ~float32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumFloat32x16(collection)
	return sum / T(len(collection))
}

// MeanFloat64x8 calculates the mean of a slice of float64 using AVX-512 SIMD
func MeanFloat64x8[T ~float64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumFloat64x8(collection)
	return sum / T(len(collection))
}

// ClampInt8x64 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt8x64[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 64

	minVec := archsimd.BroadcastInt8x64(int8(min))
	maxVec := archsimd.BroadcastInt8x64(int8(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadInt8x64Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [64]int8
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

// ClampInt16x32 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt16x32[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 32

	minVec := archsimd.BroadcastInt16x32(int16(min))
	maxVec := archsimd.BroadcastInt16x32(int16(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadInt16x32Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [32]int16
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

// ClampInt32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt32x16[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 16

	minVec := archsimd.BroadcastInt32x16(int32(min))
	maxVec := archsimd.BroadcastInt32x16(int32(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadInt32x16Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [16]int32
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

// ClampInt64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampInt64x8[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 8

	minVec := archsimd.BroadcastInt64x8(int64(min))
	maxVec := archsimd.BroadcastInt64x8(int64(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadInt64x8Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [8]int64
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

// ClampUint8x64 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint8x64[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 64

	minVec := archsimd.BroadcastUint8x64(uint8(min))
	maxVec := archsimd.BroadcastUint8x64(uint8(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadUint8x64Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [64]uint8
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

// ClampUint16x32 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint16x32[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 32

	minVec := archsimd.BroadcastUint16x32(uint16(min))
	maxVec := archsimd.BroadcastUint16x32(uint16(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadUint16x32Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [32]uint16
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

// ClampUint32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint32x16[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 16

	minVec := archsimd.BroadcastUint32x16(uint32(min))
	maxVec := archsimd.BroadcastUint32x16(uint32(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadUint32x16Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [16]uint32
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

// ClampUint64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampUint64x8[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 8

	minVec := archsimd.BroadcastUint64x8(uint64(min))
	maxVec := archsimd.BroadcastUint64x8(uint64(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadUint64x8Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [8]uint64
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

// ClampFloat32x16 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampFloat32x16[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 16

	minVec := archsimd.BroadcastFloat32x16(float32(min))
	maxVec := archsimd.BroadcastFloat32x16(float32(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadFloat32x16Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [16]float32
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

// ClampFloat64x8 clamps each element in collection between min and max values using AVX-512 SIMD
func ClampFloat64x8[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 8

	minVec := archsimd.BroadcastFloat64x8(float64(min))
	maxVec := archsimd.BroadcastFloat64x8(float64(max))

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), len(collection))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := base[i : i+lanes]
		v := archsimd.LoadFloat64x8Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [8]float64
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

// MinInt8x64 finds the minimum value in a collection of int8 using AVX-512 SIMD
func MinInt8x64[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 64

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Int8x64
	firstInitialized := false

	i := 0
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
		var buf [64]int8
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

// MinInt16x32 finds the minimum value in a collection of int16 using AVX-512 SIMD
func MinInt16x32[T ~int16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Int16x32
	firstInitialized := false

	i := 0
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
		var buf [32]int16
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

// MinInt32x16 finds the minimum value in a collection of int32 using AVX-512 SIMD
func MinInt32x16[T ~int32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Int32x16
	firstInitialized := false

	i := 0
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
		var buf [16]int32
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

// MinInt64x8 finds the minimum value in a collection of int64 using AVX-512 SIMD
func MinInt64x8[T ~int64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Int64x8
	firstInitialized := false

	i := 0
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
		var buf [8]int64
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

// MinUint8x64 finds the minimum value in a collection of uint8 using AVX-512 SIMD
func MinUint8x64[T ~uint8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 64

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Uint8x64
	firstInitialized := false

	i := 0
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
		var buf [64]uint8
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

// MinUint16x32 finds the minimum value in a collection of uint16 using AVX-512 SIMD
func MinUint16x32[T ~uint16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Uint16x32
	firstInitialized := false

	i := 0
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
		var buf [32]uint16
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

// MinUint32x16 finds the minimum value in a collection of uint32 using AVX-512 SIMD
func MinUint32x16[T ~uint32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Uint32x16
	firstInitialized := false

	i := 0
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
		var buf [16]uint32
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

// MinUint64x8 finds the minimum value in a collection of uint64 using AVX-512 SIMD
func MinUint64x8[T ~uint64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Uint64x8
	firstInitialized := false

	i := 0
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
		var buf [8]uint64
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

// MinFloat32x16 finds the minimum value in a collection of float32 using AVX-512 SIMD
func MinFloat32x16[T ~float32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Float32x16
	firstInitialized := false

	i := 0
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
		var buf [16]float32
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

// MinFloat64x8 finds the minimum value in a collection of float64 using AVX-512 SIMD
func MinFloat64x8[T ~float64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), length)

	var minVec archsimd.Float64x8
	firstInitialized := false

	i := 0
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
		var buf [8]float64
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

// MaxInt8x64 finds the maximum value in a collection of int8 using AVX-512 SIMD
func MaxInt8x64[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 64

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Int8x64
	firstInitialized := false

	i := 0
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
		var buf [64]int8
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

// MaxInt16x32 finds the maximum value in a collection of int16 using AVX-512 SIMD
func MaxInt16x32[T ~int16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Int16x32
	firstInitialized := false

	i := 0
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
		var buf [32]int16
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

// MaxInt32x16 finds the maximum value in a collection of int32 using AVX-512 SIMD
func MaxInt32x16[T ~int32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Int32x16
	firstInitialized := false

	i := 0
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
		var buf [16]int32
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

// MaxInt64x8 finds the maximum value in a collection of int64 using AVX-512 SIMD
func MaxInt64x8[T ~int64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Int64x8
	firstInitialized := false

	i := 0
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
		var buf [8]int64
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

// MaxUint8x64 finds the maximum value in a collection of uint8 using AVX-512 SIMD
func MaxUint8x64[T ~uint8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 64

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Uint8x64
	firstInitialized := false

	i := 0
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
		var buf [64]uint8
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

// MaxUint16x32 finds the maximum value in a collection of uint16 using AVX-512 SIMD
func MaxUint16x32[T ~uint16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Uint16x32
	firstInitialized := false

	i := 0
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
		var buf [32]uint16
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

// MaxUint32x16 finds the maximum value in a collection of uint32 using AVX-512 SIMD
func MaxUint32x16[T ~uint32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Uint32x16
	firstInitialized := false

	i := 0
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
		var buf [16]uint32
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

// MaxUint64x8 finds the maximum value in a collection of uint64 using AVX-512 SIMD
func MaxUint64x8[T ~uint64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Uint64x8
	firstInitialized := false

	i := 0
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
		var buf [8]uint64
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

// MaxFloat32x16 finds the maximum value in a collection of float32 using AVX-512 SIMD
func MaxFloat32x16[T ~float32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Float32x16
	firstInitialized := false

	i := 0
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
		var buf [16]float32
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

// MaxFloat64x8 finds the maximum value in a collection of float64 using AVX-512 SIMD
func MaxFloat64x8[T ~float64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	// bearer:disable go_gosec_unsafe_unsafe
	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), length)

	var maxVec archsimd.Float64x8
	firstInitialized := false

	i := 0
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
		var buf [8]float64
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
