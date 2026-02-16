//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// MinInt8x32 finds the minimum value in a collection of int8 using AVX2 SIMD
func MinInt8x32[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	var minVec archsimd.Int8x32
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt8x32Slice(s)

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

	var minVec archsimd.Int16x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt16x16Slice(s)

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

	var minVec archsimd.Int32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt32x8Slice(s)

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

	var minVec archsimd.Int64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt64x4Slice(s)

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

	var minVec archsimd.Uint8x32
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint8x32Slice(s)

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

	var minVec archsimd.Uint16x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint16x16Slice(s)

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

	var minVec archsimd.Uint32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint32x8Slice(s)

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

	var minVec archsimd.Uint64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint64x4Slice(s)

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

	var minVec archsimd.Float32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat32x8Slice(s)

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

	var minVec archsimd.Float64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat64x4Slice(s)

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
