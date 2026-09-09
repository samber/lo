//go:build goexperiment.simd

package simd

import "unsafe"

// asInt8 reinterprets a []T (T ~int8) as []int8. The two have identical size, alignment
// and representation. unsafe.SliceData is used instead of &s[0] because it is defined for
// empty and nil slices too.
func asInt8[T ~int8](s []T) []int8 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*int8)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asInt16 is asInt8 for int16.
func asInt16[T ~int16](s []T) []int16 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*int16)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asInt32 is asInt8 for int32.
func asInt32[T ~int32](s []T) []int32 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*int32)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asInt64 is asInt8 for int64.
func asInt64[T ~int64](s []T) []int64 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*int64)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asUint8 is asInt8 for uint8.
func asUint8[T ~uint8](s []T) []uint8 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*uint8)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asUint16 is asInt8 for uint16.
func asUint16[T ~uint16](s []T) []uint16 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*uint16)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asUint32 is asInt8 for uint32.
func asUint32[T ~uint32](s []T) []uint32 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*uint32)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asUint64 is asInt8 for uint64.
func asUint64[T ~uint64](s []T) []uint64 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*uint64)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asFloat32 is asInt8 for float32.
func asFloat32[T ~float32](s []T) []float32 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*float32)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

// asFloat64 is asInt8 for float64.
func asFloat64[T ~float64](s []T) []float64 {
	// bearer:disable go_gosec_unsafe_unsafe
	return unsafe.Slice((*float64)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}
