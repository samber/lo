//go:build go1.26 && goexperiment.simd && amd64

package simd

import "simd/archsimd"

const (
	simdLanes2  = uint(2)
	simdLanes4  = uint(4)
	simdLanes8  = uint(8)
	simdLanes16 = uint(16)
	simdLanes32 = uint(32)
	simdLanes64 = uint(64)
)

// simdFeature represents the highest available SIMD instruction set
type simdFeature int

const (
	simdFeatureNone simdFeature = iota
	simdFeatureAVX
	simdFeatureAVX2
	simdFeatureAVX512
)

// currentSimdFeature is cached at package init to avoid repeated CPU feature checks
var currentSimdFeature simdFeature

func init() {
	if archsimd.X86.AVX512() {
		currentSimdFeature = simdFeatureAVX512
	} else if archsimd.X86.AVX2() {
		currentSimdFeature = simdFeatureAVX2
	} else if archsimd.X86.AVX() {
		currentSimdFeature = simdFeatureAVX
	}
}
