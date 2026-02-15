//go:build go1.26 && goexperiment.simd && amd64

package simd

// MeanInt8x64 calculates the mean of a slice of int8 using AVX-512 SIMD
func MeanInt8x64[T ~int8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt8x64(collection)
	return sum / T(len(collection))
}

// MeanInt16x32 calculates the mean of a slice of int16 using AVX-512 SIMD
func MeanInt16x32[T ~int16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt16x32(collection)
	return sum / T(len(collection))
}

// MeanInt32x16 calculates the mean of a slice of int32 using AVX-512 SIMD
func MeanInt32x16[T ~int32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt32x16(collection)
	return sum / T(len(collection))
}

// MeanInt64x8 calculates the mean of a slice of int64 using AVX-512 SIMD
func MeanInt64x8[T ~int64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt64x8(collection)
	return sum / T(len(collection))
}

// MeanUint8x64 calculates the mean of a slice of uint8 using AVX-512 SIMD
func MeanUint8x64[T ~uint8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint8x64(collection)
	return sum / T(len(collection))
}

// MeanUint16x32 calculates the mean of a slice of uint16 using AVX-512 SIMD
func MeanUint16x32[T ~uint16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint16x32(collection)
	return sum / T(len(collection))
}

// MeanUint32x16 calculates the mean of a slice of uint32 using AVX-512 SIMD
func MeanUint32x16[T ~uint32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint32x16(collection)
	return sum / T(len(collection))
}

// MeanUint64x8 calculates the mean of a slice of uint64 using AVX-512 SIMD
func MeanUint64x8[T ~uint64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint64x8(collection)
	return sum / T(len(collection))
}

// MeanFloat32x16 calculates the mean of a slice of float32 using AVX-512 SIMD
func MeanFloat32x16[T ~float32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumFloat32x16(collection)
	return sum / T(len(collection))
}

// MeanFloat64x8 calculates the mean of a slice of float64 using AVX-512 SIMD
func MeanFloat64x8[T ~float64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumFloat64x8(collection)
	return sum / T(len(collection))
}
