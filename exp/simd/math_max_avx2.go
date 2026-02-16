//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// MaxInt8x32 finds the maximum value in a collection of int8 using AVX2 SIMD
func MaxInt8x32[T ~int8](collection []T) T {
	length := len(collection)
	if length == 0 {
		return 0
	}

	lanes := 32

	var maxVec archsimd.Int8x32
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt8x32Slice(s)

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

	var maxVec archsimd.Int16x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt16x16Slice(s)

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

	var maxVec archsimd.Int32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt32x8Slice(s)

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

	var maxVec archsimd.Int64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt64x4Slice(s)

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

	var maxVec archsimd.Uint8x32
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint8x32Slice(s)

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

	var maxVec archsimd.Uint16x16
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint16x16Slice(s)

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

	var maxVec archsimd.Uint32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint32x8Slice(s)

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

	var maxVec archsimd.Uint64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint64x4Slice(s)

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

	var maxVec archsimd.Float32x8
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat32x8Slice(s)

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

	var maxVec archsimd.Float64x4
	firstInitialized := false

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat64x4Slice(s)

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
