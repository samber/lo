//go:build go1.26 && goexperiment.simd && amd64

package simd

// MeanInt8x16 calculates the mean of a slice of int8 using SSE SIMD
func MeanInt8x16[T ~int8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt8x16(collection)
	return sum / T(len(collection))
}

// MeanInt16x8 calculates the mean of a slice of int16 using SSE SIMD
func MeanInt16x8[T ~int16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt16x8(collection)
	return sum / T(len(collection))
}

// MeanInt32x4 calculates the mean of a slice of int32 using SSE SIMD
func MeanInt32x4[T ~int32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt32x4(collection)
	return sum / T(len(collection))
}

// MeanInt64x2 calculates the mean of a slice of int64 using SSE SIMD
func MeanInt64x2[T ~int64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumInt64x2(collection)
	return sum / T(len(collection))
}

// MeanUint8x16 calculates the mean of a slice of uint8 using SSE SIMD
func MeanUint8x16[T ~uint8](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint8x16(collection)
	return sum / T(len(collection))
}

// MeanUint16x8 calculates the mean of a slice of uint16 using SSE SIMD
func MeanUint16x8[T ~uint16](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint16x8(collection)
	return sum / T(len(collection))
}

// MeanUint32x4 calculates the mean of a slice of uint32 using SSE SIMD
func MeanUint32x4[T ~uint32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint32x4(collection)
	return sum / T(len(collection))
}

// MeanUint64x2 calculates the mean of a slice of uint64 using SSE SIMD
func MeanUint64x2[T ~uint64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumUint64x2(collection)
	return sum / T(len(collection))
}

// MeanFloat32x4 calculates the mean of a slice of float32 using SSE SIMD
func MeanFloat32x4[T ~float32](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumFloat32x4(collection)
	return sum / T(len(collection))
}

// MeanFloat64x2 calculates the mean of a slice of float64 using SSE SIMD
func MeanFloat64x2[T ~float64](collection []T) T {
	if len(collection) == 0 {
		return 0
	}
	sum := SumFloat64x2(collection)
	return sum / T(len(collection))
}
