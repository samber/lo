//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// ClampInt8x32 clamps each element in collection between min and max values using AVX2 SIMD
func ClampInt8x32[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 32

	minVec := archsimd.BroadcastInt8x32(int8(min))
	maxVec := archsimd.BroadcastInt8x32(int8(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*int8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt8x32Slice(c)

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

	minVec := archsimd.BroadcastInt16x16(int16(min))
	maxVec := archsimd.BroadcastInt16x16(int16(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*int16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt16x16Slice(c)

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

	minVec := archsimd.BroadcastInt32x8(int32(min))
	maxVec := archsimd.BroadcastInt32x8(int32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*int32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt32x8Slice(c)

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

	minVec := archsimd.BroadcastInt64x4(int64(min))
	maxVec := archsimd.BroadcastInt64x4(int64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*int64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt64x4Slice(c)

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

	minVec := archsimd.BroadcastUint8x32(uint8(min))
	maxVec := archsimd.BroadcastUint8x32(uint8(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint8x32Slice(c)

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

	minVec := archsimd.BroadcastUint16x16(uint16(min))
	maxVec := archsimd.BroadcastUint16x16(uint16(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint16x16Slice(c)

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

	minVec := archsimd.BroadcastUint32x8(uint32(min))
	maxVec := archsimd.BroadcastUint32x8(uint32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint32x8Slice(c)

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

	minVec := archsimd.BroadcastUint64x4(uint64(min))
	maxVec := archsimd.BroadcastUint64x4(uint64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint64x4Slice(c)

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

	minVec := archsimd.BroadcastFloat32x8(float32(min))
	maxVec := archsimd.BroadcastFloat32x8(float32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*float32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat32x8Slice(c)

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

	minVec := archsimd.BroadcastFloat64x4(float64(min))
	maxVec := archsimd.BroadcastFloat64x4(float64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*float64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat64x4Slice(c)

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
