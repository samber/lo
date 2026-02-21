//go:build go1.26 && goexperiment.simd && amd64

package simd

import "simd/archsimd"

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
