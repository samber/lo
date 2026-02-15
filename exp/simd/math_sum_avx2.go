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
	lanes := 32

	var acc archsimd.Int8x32

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt8x32Slice(s)
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
	lanes := 16

	var acc archsimd.Int16x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt16x16Slice(s)
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
	lanes := 8

	var acc archsimd.Int32x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt32x8Slice(s)
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
	lanes := 4

	var acc archsimd.Int64x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt64x4Slice(s)
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
	lanes := 32

	var acc archsimd.Uint8x32

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint8x32Slice(s)
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
	lanes := 16

	var acc archsimd.Uint16x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint16x16Slice(s)
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
	lanes := 8

	var acc archsimd.Uint32x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint32x8Slice(s)
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
	lanes := 4

	var acc archsimd.Uint64x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint64x4Slice(s)
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
	lanes := 8

	var acc archsimd.Float32x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat32x8Slice(s)
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
	lanes := 4

	var acc archsimd.Float64x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat64x4Slice(s)
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
