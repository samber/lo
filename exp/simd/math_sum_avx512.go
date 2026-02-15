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
	lanes := 64

	var acc archsimd.Int8x64

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt8x64Slice(s)
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
	lanes := 32

	var acc archsimd.Int16x32

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt16x32Slice(s)
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
	lanes := 16

	var acc archsimd.Int32x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt32x16Slice(s)
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
	lanes := 8

	var acc archsimd.Int64x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt64x8Slice(s)
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
	lanes := 64

	var acc archsimd.Uint8x64

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint8x64Slice(s)
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
	lanes := 32

	var acc archsimd.Uint16x32

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint16x32Slice(s)
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
	lanes := 16

	var acc archsimd.Uint32x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint32x16Slice(s)
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
	lanes := 8

	var acc archsimd.Uint64x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint64x8Slice(s)
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
	lanes := 16

	var acc archsimd.Float32x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat32x16Slice(s)
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
	lanes := 8

	var acc archsimd.Float64x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat64x8Slice(s)
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
