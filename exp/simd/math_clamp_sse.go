//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// ClampInt8x16 clamps each element in collection between min and max values using SSE SIMD
func ClampInt8x16[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 16

	minVec := archsimd.BroadcastInt8x16(int8(min))
	maxVec := archsimd.BroadcastInt8x16(int8(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*int8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt8x16Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [16]int8
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

// ClampInt16x8 clamps each element in collection between min and max values using SSE SIMD
func ClampInt16x8[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 8

	minVec := archsimd.BroadcastInt16x8(int16(min))
	maxVec := archsimd.BroadcastInt16x8(int16(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*int16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt16x8Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [8]int16
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

// ClampInt32x4 clamps each element in collection between min and max values using SSE SIMD
func ClampInt32x4[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 4

	minVec := archsimd.BroadcastInt32x4(int32(min))
	maxVec := archsimd.BroadcastInt32x4(int32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*int32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt32x4Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [4]int32
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

// ClampInt64x2 clamps each element in collection between min and max values using SSE SIMD and AVX-512 SIMD.
func ClampInt64x2[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 2

	minVec := archsimd.BroadcastInt64x2(int64(min))
	maxVec := archsimd.BroadcastInt64x2(int64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*int64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt64x2Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [2]int64
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

// ClampUint8x16 clamps each element in collection between min and max values using SSE SIMD
func ClampUint8x16[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 16

	minVec := archsimd.BroadcastUint8x16(uint8(min))
	maxVec := archsimd.BroadcastUint8x16(uint8(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint8x16Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [16]uint8
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

// ClampUint16x8 clamps each element in collection between min and max values using SSE SIMD
func ClampUint16x8[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 8

	minVec := archsimd.BroadcastUint16x8(uint16(min))
	maxVec := archsimd.BroadcastUint16x8(uint16(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint16x8Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [8]uint16
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

// ClampUint32x4 clamps each element in collection between min and max values using SSE SIMD
func ClampUint32x4[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 4

	minVec := archsimd.BroadcastUint32x4(uint32(min))
	maxVec := archsimd.BroadcastUint32x4(uint32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint32x4Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [4]uint32
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

// ClampUint64x2 clamps each element in collection between min and max values using SSE SIMD and AVX-512 SIMD.
func ClampUint64x2[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 2

	minVec := archsimd.BroadcastUint64x2(uint64(min))
	maxVec := archsimd.BroadcastUint64x2(uint64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint64x2Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [2]uint64
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

// ClampFloat32x4 clamps each element in collection between min and max values using SSE SIMD
func ClampFloat32x4[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 4

	minVec := archsimd.BroadcastFloat32x4(float32(min))
	maxVec := archsimd.BroadcastFloat32x4(float32(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*float32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat32x4Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [4]float32
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

// ClampFloat64x2 clamps each element in collection between min and max values using SSE SIMD
func ClampFloat64x2[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice {
	if len(collection) == 0 {
		return collection
	}

	result := make(Slice, len(collection))
	lanes := 2

	minVec := archsimd.BroadcastFloat64x2(float64(min))
	maxVec := archsimd.BroadcastFloat64x2(float64(max))

	i := 0
	for ; i+lanes <= len(collection); i += lanes {
		c := unsafe.Slice((*float64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat64x2Slice(c)

		clamped := v.Max(minVec).Min(maxVec)

		var buf [2]float64
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
