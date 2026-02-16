//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"simd/archsimd"

	"github.com/samber/lo"
)

func SumInt8[T ~int8](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumInt8x32(collection)
	} else if archsimd.X86.AVX512() {
		return SumInt8x64(collection)
	}
	return lo.Sum(collection)
}

func SumInt16[T ~int16](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumInt16x16(collection)
	} else if archsimd.X86.AVX512() {
		return SumInt16x32(collection)
	}
	return lo.Sum(collection)
}

func SumInt32[T ~int32](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumInt32x8(collection)
	} else if archsimd.X86.AVX512() {
		return SumInt32x16(collection)
	}
	return lo.Sum(collection)
}

func SumInt64[T ~int64](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumInt64x4(collection)
	} else if archsimd.X86.AVX512() {
		return SumInt64x8(collection)
	}
	return lo.Sum(collection)
}

func SumUint8[T ~uint8](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumUint8x32(collection)
	} else if archsimd.X86.AVX512() {
		return SumUint8x64(collection)
	}
	return lo.Sum(collection)
}

func SumUint16[T ~uint16](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumUint16x16(collection)
	} else if archsimd.X86.AVX512() {
		return SumUint16x32(collection)
	}
	return lo.Sum(collection)
}

func SumUint32[T ~uint32](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumUint32x8(collection)
	} else if archsimd.X86.AVX512() {
		return SumUint32x16(collection)
	}
	return lo.Sum(collection)
}

func SumUint64[T ~uint64](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumUint64x4(collection)
	} else if archsimd.X86.AVX512() {
		return SumUint64x8(collection)
	}
	return lo.Sum(collection)
}

func SumFloat32[T ~float32](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumFloat32x8(collection)
	} else if archsimd.X86.AVX512() {
		return SumFloat32x16(collection)
	}
	return lo.Sum(collection)
}

func SumFloat64[T ~float64](collection []T) T {
	if archsimd.X86.AVX2() {
		return SumFloat64x4(collection)
	} else if archsimd.X86.AVX512() {
		return SumFloat64x8(collection)
	}
	return lo.Sum(collection)
}
