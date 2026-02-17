//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"github.com/samber/lo"
)

func SumInt8[T ~int8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumInt8x64(collection)
	case simdFeatureAVX2:
		return SumInt8x32(collection)
	case simdFeatureAVX:
		return SumInt8x16(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumInt16[T ~int16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumInt16x32(collection)
	case simdFeatureAVX2:
		return SumInt16x16(collection)
	case simdFeatureAVX:
		return SumInt16x8(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumInt32[T ~int32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumInt32x16(collection)
	case simdFeatureAVX2:
		return SumInt32x8(collection)
	case simdFeatureAVX:
		return SumInt32x4(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumInt64[T ~int64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumInt64x8(collection)
	case simdFeatureAVX2:
		return SumInt64x4(collection)
	case simdFeatureAVX:
		return SumInt64x2(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumUint8[T ~uint8](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumUint8x64(collection)
	case simdFeatureAVX2:
		return SumUint8x32(collection)
	case simdFeatureAVX:
		return SumUint8x16(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumUint16[T ~uint16](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumUint16x32(collection)
	case simdFeatureAVX2:
		return SumUint16x16(collection)
	case simdFeatureAVX:
		return SumUint16x8(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumUint32[T ~uint32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumUint32x16(collection)
	case simdFeatureAVX2:
		return SumUint32x8(collection)
	case simdFeatureAVX:
		return SumUint32x4(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumUint64[T ~uint64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumUint64x8(collection)
	case simdFeatureAVX2:
		return SumUint64x4(collection)
	case simdFeatureAVX:
		return SumUint64x2(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumFloat32[T ~float32](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumFloat32x16(collection)
	case simdFeatureAVX2:
		return SumFloat32x8(collection)
	case simdFeatureAVX:
		return SumFloat32x4(collection)
	default:
		return lo.Sum(collection)
	}
}

func SumFloat64[T ~float64](collection []T) T {
	switch currentSimdFeature {
	case simdFeatureAVX512:
		return SumFloat64x8(collection)
	case simdFeatureAVX2:
		return SumFloat64x4(collection)
	case simdFeatureAVX:
		return SumFloat64x2(collection)
	default:
		return lo.Sum(collection)
	}
}
