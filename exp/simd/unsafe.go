//go:build go1.26 && goexperiment.simd && amd64

package simd

import "unsafe"

// unsafeSliceInt8 converts a []T (where T ~int8) to []int8 via unsafe operations.
// This helper reduces code duplication and the risk of copy-paste errors.
//
//go:nosplit
func unsafeSliceInt8[T ~int8](collection []T, length uint) []int8 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*int8)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceInt16 converts a []T (where T ~int16) to []int16 via unsafe operations.
//
//go:nosplit
func unsafeSliceInt16[T ~int16](collection []T, length uint) []int16 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*int16)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceInt32 converts a []T (where T ~int32) to []int32 via unsafe operations.
//
//go:nosplit
func unsafeSliceInt32[T ~int32](collection []T, length uint) []int32 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*int32)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceInt64 converts a []T (where T ~int64) to []int64 via unsafe operations.
//
//go:nosplit
func unsafeSliceInt64[T ~int64](collection []T, length uint) []int64 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*int64)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceUint8 converts a []T (where T ~uint8) to []uint8 via unsafe operations.
//
//go:nosplit
func unsafeSliceUint8[T ~uint8](collection []T, length uint) []uint8 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*uint8)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceUint16 converts a []T (where T ~uint16) to []uint16 via unsafe operations.
//
//go:nosplit
func unsafeSliceUint16[T ~uint16](collection []T, length uint) []uint16 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*uint16)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceUint32 converts a []T (where T ~uint32) to []uint32 via unsafe operations.
//
//go:nosplit
func unsafeSliceUint32[T ~uint32](collection []T, length uint) []uint32 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*uint32)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceUint64 converts a []T (where T ~uint64) to []uint64 via unsafe operations.
//
//go:nosplit
func unsafeSliceUint64[T ~uint64](collection []T, length uint) []uint64 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*uint64)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceFloat32 converts a []T (where T ~float32) to []float32 via unsafe operations.
//
//go:nosplit
func unsafeSliceFloat32[T ~float32](collection []T, length uint) []float32 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*float32)(unsafe.Pointer(&collection[0])), length)
}

// unsafeSliceFloat64 converts a []T (where T ~float64) to []float64 via unsafe operations.
//
//go:nosplit
func unsafeSliceFloat64[T ~float64](collection []T, length uint) []float64 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*float64)(unsafe.Pointer(&collection[0])), length)
}
