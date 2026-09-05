//go:build go1.26 && goexperiment.simd && amd64

package simd

import "unsafe"

// unsafeIndexVec provides direct pointer-based access to SIMD lanes without slice bounds checking.
// This is a replacement for slice-based access like `base[i : i+lanes]` that eliminates
// runtime.panicBounds checks in critical SIMD loops.
//
// Parameters:
//
//	V - The SIMD vector type (e.g., [4]float64, [2]int64, [16]uint8])
//	T - The element type of the collection
//	collection - The input slice
//	i - The starting vector index (0-based, multiplied by lanes internally)
//
// Usage example:
//
//		const lanes = 4
//		i := uint(0)
//		for ; i+lanes <= uint(len(collection)); i += lanes {
//			// Old way (with bounds checks):
//			archsimd.LoadInt64x2Slice(base[i : i+lanes])
//
//			// New way (no bounds checks):
//			archsimd.LoadInt64x2(unsafeIndexVec[[lanes]int64](collection, i))
//		}
//	}
//
// Benefits over slice-based access:
//   - Eliminates runtime.panicBounds checks
//   - Smaller generated code size (~96 bytes / ~12 instructions less)
//   - Better cache locality (no slice header overhead)
//   - More efficient for tight SIMD loops
//
//bearer:disable go_gosec_unsafe_unsafe
//go:nosplit
func unsafeIndexVec[V any, T any](collection []T, i uint) *V {
	data := unsafe.SliceData(collection)
	size := unsafe.Sizeof(*data)
	return (*V)(unsafe.Add(unsafe.Pointer(data), uintptr(i)*size))
}
