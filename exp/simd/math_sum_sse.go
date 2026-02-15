//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"
	"unsafe"
)

// SSE (128-bit) SIMD sum functions - 16/8/4/2 lanes

// SumInt8x16 sums a slice of int8 using SSE SIMD (Int8x16, 16 lanes)
func SumInt8x16[T ~int8](collection []T) T {
	length := len(collection)
	lanes := 16

	var acc archsimd.Int8x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt8x16Slice(s)
		acc = acc.Add(v)
	}

	var buf [16]int8
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

// SumInt16x8 sums a slice of int16 using SSE SIMD (Int16x8, 8 lanes)
func SumInt16x8[T ~int16](collection []T) T {
	length := len(collection)
	lanes := 8

	var acc archsimd.Int16x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt16x8Slice(s)
		acc = acc.Add(v)
	}

	var buf [8]int16
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

// SumInt32x4 sums a slice of int32 using SSE SIMD (Int32x4, 4 lanes)
func SumInt32x4[T ~int32](collection []T) T {
	length := len(collection)
	lanes := 4

	var acc archsimd.Int32x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt32x4Slice(s)
		acc = acc.Add(v)
	}

	var buf [4]int32
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

// SumInt64x2 sums a slice of int64 using SSE SIMD (Int64x2, 2 lanes)
func SumInt64x2[T ~int64](collection []T) T {
	length := len(collection)
	lanes := 2

	var acc archsimd.Int64x2

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*int64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadInt64x2Slice(s)
		acc = acc.Add(v)
	}

	var buf [2]int64
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

// SumUint8x16 sums a slice of uint8 using SSE SIMD (Uint8x16, 16 lanes)
func SumUint8x16[T ~uint8](collection []T) T {
	length := len(collection)
	lanes := 16

	var acc archsimd.Uint8x16

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint8)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint8x16Slice(s)
		acc = acc.Add(v)
	}

	var buf [16]uint8
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

// SumUint16x8 sums a slice of uint16 using SSE SIMD (Uint16x8, 8 lanes)
func SumUint16x8[T ~uint16](collection []T) T {
	length := len(collection)
	lanes := 8

	var acc archsimd.Uint16x8

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint16)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint16x8Slice(s)
		acc = acc.Add(v)
	}

	var buf [8]uint16
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

// SumUint32x4 sums a slice of uint32 using SSE SIMD (Uint32x4, 4 lanes)
func SumUint32x4[T ~uint32](collection []T) T {
	length := len(collection)
	lanes := 4

	var acc archsimd.Uint32x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint32x4Slice(s)
		acc = acc.Add(v)
	}

	var buf [4]uint32
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

// SumUint64x2 sums a slice of uint64 using SSE SIMD (Uint64x2, 2 lanes)
func SumUint64x2[T ~uint64](collection []T) T {
	length := len(collection)
	lanes := 2

	var acc archsimd.Uint64x2

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*uint64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadUint64x2Slice(s)
		acc = acc.Add(v)
	}

	var buf [2]uint64
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

// SumFloat32x4 sums a slice of float32 using SSE SIMD (Float32x4, 4 lanes)
func SumFloat32x4[T ~float32](collection []T) T {
	length := len(collection)
	lanes := 4

	var acc archsimd.Float32x4

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float32)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat32x4Slice(s)
		acc = acc.Add(v)
	}

	var buf [4]float32
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

// SumFloat64x2 sums a slice of float64 using SSE SIMD (Float64x2, 2 lanes)
func SumFloat64x2[T ~float64](collection []T) T {
	length := len(collection)
	lanes := 2

	var acc archsimd.Float64x2

	i := 0
	for ; i+lanes <= length; i += lanes {
		s := unsafe.Slice((*float64)(unsafe.Pointer(&collection[i])), lanes)
		v := archsimd.LoadFloat64x2Slice(s)
		acc = acc.Add(v)
	}

	var buf [2]float64
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
