//go:build go1.26 && goexperiment.simd && amd64

package simd

// Empty file to satisfy the build constraint for non-supported architectures.

const (
	simdLanes2  = uint(2)
	simdLanes4  = uint(4)
	simdLanes8  = uint(8)
	simdLanes16 = uint(16)
	simdLanes32 = uint(32)
	simdLanes64 = uint(64)
)
