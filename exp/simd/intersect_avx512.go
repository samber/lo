//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
)

// ContainsInt8x16 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsInt8x16[T ~int8](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes16
	targetVec := archsimd.BroadcastInt8x16(int8(target))

	base := unsafeSliceInt8(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt8x16Slice(s)

		// Compare for equality; Equal returns a mask, ToBits() its bitmask.
		cmp := v.Equal(targetVec)
		if cmp.ToBits() != 0 {
			return true
		}
	}

	// Handle remaining elements
	for ; i < length; i++ {
		if collection[i] == target {
			return true
		}
	}

	return false
}

// ContainsInt16x8 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsInt16x8[T ~int16](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes8
	targetVec := archsimd.BroadcastInt16x8(int16(target))

	base := unsafeSliceInt16(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt16x8Slice(s)

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

// ContainsInt32x4 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsInt32x4[T ~int32](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes4
	targetVec := archsimd.BroadcastInt32x4(int32(target))

	base := unsafeSliceInt32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt32x4Slice(s)

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

// ContainsInt64x2 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsInt64x2[T ~int64](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes2
	targetVec := archsimd.BroadcastInt64x2(int64(target))

	base := unsafeSliceInt64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt64x2Slice(s)

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

// ContainsUint8x16 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsUint8x16[T ~uint8](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes16
	targetVec := archsimd.BroadcastUint8x16(uint8(target))

	base := unsafeSliceUint8(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint8x16Slice(s)

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

// ContainsUint16x8 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsUint16x8[T ~uint16](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes8
	targetVec := archsimd.BroadcastUint16x8(uint16(target))

	base := unsafeSliceUint16(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint16x8Slice(s)

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

// ContainsUint32x4 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsUint32x4[T ~uint32](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes4
	targetVec := archsimd.BroadcastUint32x4(uint32(target))

	base := unsafeSliceUint32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint32x4Slice(s)

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

// ContainsUint64x2 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsUint64x2[T ~uint64](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes2
	targetVec := archsimd.BroadcastUint64x2(uint64(target))

	base := unsafeSliceUint64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint64x2Slice(s)

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

// ContainsFloat32x4 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsFloat32x4[T ~float32](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes4
	targetVec := archsimd.BroadcastFloat32x4(float32(target))

	base := unsafeSliceFloat32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadFloat32x4Slice(s)

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

// ContainsFloat64x2 checks if collection contains target using AVX SIMD and AVX-512 SIMD
func ContainsFloat64x2[T ~float64](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes2
	targetVec := archsimd.BroadcastFloat64x2(float64(target))

	base := unsafeSliceFloat64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadFloat64x2Slice(s)

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

// ContainsInt8x32 checks if collection contains target using AVX2 SIMD
func ContainsInt8x32[T ~int8](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes32
	targetVec := archsimd.BroadcastInt8x32(int8(target))

	base := unsafeSliceInt8(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt8x32Slice(s)

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

// ContainsInt16x16 checks if collection contains target using AVX2 SIMD
func ContainsInt16x16[T ~int16](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes16
	targetVec := archsimd.BroadcastInt16x16(int16(target))

	base := unsafeSliceInt16(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt16x16Slice(s)

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

// ContainsInt32x8 checks if collection contains target using AVX2 SIMD
func ContainsInt32x8[T ~int32](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes8
	targetVec := archsimd.BroadcastInt32x8(int32(target))

	base := unsafeSliceInt32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt32x8Slice(s)

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

// ContainsInt64x4 checks if collection contains target using AVX2 SIMD
func ContainsInt64x4[T ~int64](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes4
	targetVec := archsimd.BroadcastInt64x4(int64(target))

	base := unsafeSliceInt64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadInt64x4Slice(s)

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

// ContainsUint8x32 checks if collection contains target using AVX2 SIMD
func ContainsUint8x32[T ~uint8](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes32
	targetVec := archsimd.BroadcastUint8x32(uint8(target))

	base := unsafeSliceUint8(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint8x32Slice(s)

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

// ContainsUint16x16 checks if collection contains target using AVX2 SIMD
func ContainsUint16x16[T ~uint16](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes16
	targetVec := archsimd.BroadcastUint16x16(uint16(target))

	base := unsafeSliceUint16(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint16x16Slice(s)

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

// ContainsUint32x8 checks if collection contains target using AVX2 SIMD
func ContainsUint32x8[T ~uint32](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes8
	targetVec := archsimd.BroadcastUint32x8(uint32(target))

	base := unsafeSliceUint32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint32x8Slice(s)

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

// ContainsUint64x4 checks if collection contains target using AVX2 SIMD
func ContainsUint64x4[T ~uint64](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes4
	targetVec := archsimd.BroadcastUint64x4(uint64(target))

	base := unsafeSliceUint64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadUint64x4Slice(s)

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

// ContainsFloat32x8 checks if collection contains target using AVX2 SIMD
func ContainsFloat32x8[T ~float32](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes8
	targetVec := archsimd.BroadcastFloat32x8(float32(target))

	base := unsafeSliceFloat32(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadFloat32x8Slice(s)

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

// ContainsFloat64x4 checks if collection contains target using AVX2 SIMD
func ContainsFloat64x4[T ~float64](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes4
	targetVec := archsimd.BroadcastFloat64x4(float64(target))

	base := unsafeSliceFloat64(collection, length)

	i := uint(0)
	for ; i+lanes <= length; i += lanes {
		s := base[i : i+lanes]
		v := archsimd.LoadFloat64x4Slice(s)

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

// ContainsInt8x64 checks if collection contains target using AVX-512 SIMD
func ContainsInt8x64[T ~int8](collection []T, target T) bool {
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes64
	targetVec := archsimd.BroadcastInt8x64(int8(target))

	base := unsafeSliceInt8(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes32
	targetVec := archsimd.BroadcastInt16x32(int16(target))

	base := unsafeSliceInt16(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes16
	targetVec := archsimd.BroadcastInt32x16(int32(target))

	base := unsafeSliceInt32(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes8
	targetVec := archsimd.BroadcastInt64x8(int64(target))

	base := unsafeSliceInt64(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes64
	targetVec := archsimd.BroadcastUint8x64(uint8(target))

	base := unsafeSliceUint8(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes32
	targetVec := archsimd.BroadcastUint16x32(uint16(target))

	base := unsafeSliceUint16(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes16
	targetVec := archsimd.BroadcastUint32x16(uint32(target))

	base := unsafeSliceUint32(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes8
	targetVec := archsimd.BroadcastUint64x8(uint64(target))

	base := unsafeSliceUint64(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes16
	targetVec := archsimd.BroadcastFloat32x16(float32(target))

	base := unsafeSliceFloat32(collection, length)

	i := uint(0)
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
	length := uint(len(collection))
	if length == 0 {
		return false
	}

	const lanes = simdLanes8
	targetVec := archsimd.BroadcastFloat64x8(float64(target))

	base := unsafeSliceFloat64(collection, length)

	i := uint(0)
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
