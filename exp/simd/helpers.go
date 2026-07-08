//go:build go1.26 && goexperiment.simd && amd64

package simd

type number interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// sum16 sums a vector of 16 elements.
func sum16[T number](vec *[16]T) T {
	return 0 +
		vec[0] + vec[1] + vec[2] + vec[3] +
		vec[4] + vec[5] + vec[6] + vec[7] +
		vec[8] + vec[9] + vec[10] + vec[11] +
		vec[12] + vec[13] + vec[14] + vec[15]
}

// sum8 sums a vector of 8 elements.
func sum8[T number](vec *[8]T) T {
	return 0 +
		vec[0] + vec[1] + vec[2] + vec[3] +
		vec[4] + vec[5] + vec[6] + vec[7]
}

// sum4 sums a vector of 4 elements.
func sum4[T number](vec *[4]T) T {
	return vec[0] + vec[1] + vec[2] + vec[3]
}

// sum2 sums a vector of 2 elements.
func sum2[T number](vec *[2]T) T {
	return vec[0] + vec[1]
}
