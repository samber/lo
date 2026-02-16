//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// ContainsInt8x64 checks if collection contains target using AVX-512 SIMD
func ContainsInt8x64[T ~int8](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 64
	targetVec := archsimd.BroadcastInt8x64(int8(target))

	base := unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt8x64Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsInt16x32 checks if collection contains target using AVX-512 SIMD
func ContainsInt16x32[T ~int16](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 32
	targetVec := archsimd.BroadcastInt16x32(int16(target))

	base := unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt16x32Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsInt32x16 checks if collection contains target using AVX-512 SIMD
func ContainsInt32x16[T ~int32](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 16
	targetVec := archsimd.BroadcastInt32x16(int32(target))

	base := unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt32x16Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsInt64x8 checks if collection contains target using AVX-512 SIMD
func ContainsInt64x8[T ~int64](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 8
	targetVec := archsimd.BroadcastInt64x8(int64(target))

	base := unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt64x8Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsUint8x64 checks if collection contains target using AVX-512 SIMD
func ContainsUint8x64[T ~uint8](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 64
	targetVec := archsimd.BroadcastUint8x64(uint8(target))

	base := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint8x64Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsUint16x32 checks if collection contains target using AVX-512 SIMD
func ContainsUint16x32[T ~uint16](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 32
	targetVec := archsimd.BroadcastUint16x32(uint16(target))

	base := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint16x32Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsUint32x16 checks if collection contains target using AVX-512 SIMD
func ContainsUint32x16[T ~uint32](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 16
	targetVec := archsimd.BroadcastUint32x16(uint32(target))

	base := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint32x16Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsUint64x8 checks if collection contains target using AVX-512 SIMD
func ContainsUint64x8[T ~uint64](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 8
	targetVec := archsimd.BroadcastUint64x8(uint64(target))

	base := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint64x8Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsFloat32x16 checks if collection contains target using AVX-512 SIMD
func ContainsFloat32x16[T ~float32](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 16
	targetVec := archsimd.BroadcastFloat32x16(float32(target))

	base := unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadFloat32x16Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsFloat64x8 checks if collection contains target using AVX-512 SIMD
func ContainsFloat64x8[T ~float64](collection []T, target T) bool {
	length := len(collection)
	if length == 0 {
		return false
	}

	lanes := 8
	targetVec := archsimd.BroadcastFloat64x8(float64(target))

	base := unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), length)

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadFloat64x8Slice(s)

		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}
