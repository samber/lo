//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// MaxInt8x16 finds the maximum value in a collection of int8 using SSE SIMD
func MaxInt8x16[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	var maxVec archsimd.Int8x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt8x16Slice(s)

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
		var buf [16]int8
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

// MaxInt16x8 finds the maximum value in a collection of int16 using SSE SIMD
func MaxInt16x8[T ~int16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	var maxVec archsimd.Int16x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt16x8Slice(s)

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
		var buf [8]int16
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

// MaxInt32x4 finds the maximum value in a collection of int32 using SSE SIMD
func MaxInt32x4[T ~int32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	var maxVec archsimd.Int32x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt32x4Slice(s)

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
		var buf [4]int32
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

// MaxInt64x2 finds the maximum value in a collection of int64 using SSE SIMD
func MaxInt64x2[T ~int64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 2

	var maxVec archsimd.Int64x2
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt64x2Slice(s)

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
		var buf [2]int64
		maxVec.Store(&buf)
		maxVal = buf[0]
		if buf[1] > maxVal {
			maxVal = buf[1]
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

// MaxUint8x16 finds the maximum value in a collection of uint8 using SSE SIMD
func MaxUint8x16[T ~uint8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 16

	var maxVec archsimd.Uint8x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint8x16Slice(s)

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
		var buf [16]uint8
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

// MaxUint16x8 finds the maximum value in a collection of uint16 using SSE SIMD
func MaxUint16x8[T ~uint16](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 8

	var maxVec archsimd.Uint16x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint16x8Slice(s)

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
		var buf [8]uint16
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

// MaxUint32x4 finds the maximum value in a collection of uint32 using SSE SIMD
func MaxUint32x4[T ~uint32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	var maxVec archsimd.Uint32x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint32x4Slice(s)

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
		var buf [4]uint32
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

// MaxUint64x2 finds the maximum value in a collection of uint64 using SSE SIMD
func MaxUint64x2[T ~uint64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 2

	var maxVec archsimd.Uint64x2
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint64x2Slice(s)

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
		var buf [2]uint64
		maxVec.Store(&buf)
		maxVal = buf[0]
		if buf[1] > maxVal {
			maxVal = buf[1]
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

// MaxFloat32x4 finds the maximum value in a collection of float32 using SSE SIMD
func MaxFloat32x4[T ~float32](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 4

	var maxVec archsimd.Float32x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat32x4Slice(s)

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
		var buf [4]float32
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

// MaxFloat64x2 finds the maximum value in a collection of float64 using SSE SIMD
func MaxFloat64x2[T ~float64](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 2

	var maxVec archsimd.Float64x2
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat64x2Slice(s)

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
		var buf [2]float64
		maxVec.Store(&buf)
		maxVal = buf[0]
		if buf[1] > maxVal {
			maxVal = buf[1]
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
