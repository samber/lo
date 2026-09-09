//go:build goexperiment.simd

package simd

import "github.com/samber/lo"

// ContainsInt8 reports whether target is present in collection, using SIMD instructions
// when available.
func ContainsInt8[T ~int8](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsInt8(asInt8(collection), int8(target))
}

// ContainsInt16 is ContainsInt8 for int16.
func ContainsInt16[T ~int16](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsInt16(asInt16(collection), int16(target))
}

// ContainsInt32 is ContainsInt8 for int32.
func ContainsInt32[T ~int32](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsInt32(asInt32(collection), int32(target))
}

// ContainsInt64 is ContainsInt8 for int64.
func ContainsInt64[T ~int64](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsInt64(asInt64(collection), int64(target))
}

// ContainsUint8 is ContainsInt8 for uint8.
func ContainsUint8[T ~uint8](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsUint8(asUint8(collection), uint8(target))
}

// ContainsUint16 is ContainsInt8 for uint16.
func ContainsUint16[T ~uint16](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsUint16(asUint16(collection), uint16(target))
}

// ContainsUint32 is ContainsInt8 for uint32.
func ContainsUint32[T ~uint32](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsUint32(asUint32(collection), uint32(target))
}

// ContainsUint64 is ContainsInt8 for uint64.
func ContainsUint64[T ~uint64](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsUint64(asUint64(collection), uint64(target))
}

// ContainsFloat32 reports whether target is present in collection, using SIMD instructions
// when available. The comparison is exact (==), like lo.Contains: NaN never matches, and
// +0.0/-0.0 always match.
func ContainsFloat32[T ~float32](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsFloat32(asFloat32(collection), float32(target))
}

// ContainsFloat64 is ContainsFloat32 for float64.
func ContainsFloat64[T ~float64](collection []T, target T) bool {
	if len(collection) == 0 {
		return false
	}
	if !useSIMD {
		return lo.Contains(collection, target)
	}
	return containsFloat64(asFloat64(collection), float64(target))
}
