---
name: Clamp
slug: clamp
sourceRef: exp/simd/math_avx.go#L453
category: exp
subCategory: simd
similarHelpers:
  - exp#simd#clamp
position: 40
signatures:
  - "func ClampInt8x16[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt8x32[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt8x64[T ~int8, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt16x8[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt16x16[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt16x32[T ~int16, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt32x4[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt32x8[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt32x16[T ~int32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt64x2[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt64x4[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampInt64x8[T ~int64, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint8x16[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint8x32[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint8x64[T ~uint8, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint16x8[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint16x16[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint16x32[T ~uint16, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint32x4[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint32x8[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint32x16[T ~uint32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint64x2[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint64x4[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampUint64x8[T ~uint64, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampFloat32x4[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampFloat32x8[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampFloat32x16[T ~float32, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampFloat64x2[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampFloat64x4[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice"
  - "func ClampFloat64x8[T ~float64, Slice ~[]T](collection Slice, min, max T) Slice"
---

Clamps each element in a collection between min and max values using SIMD instructions. The suffix (x2, x4, x8, x16, x32, x64) indicates the number of lanes processed simultaneously.

## Requirements

- **Go 1.26+** with `GOEXPERIMENT=simd`
- **amd64** architecture only

### CPU compatibility

| SIMD variant | Lanes | Required flags | Typical CPUs                   |
| ------------ | ----- | -------------- | ------------------------------ |
| AVX (xN)     | 2-16  | `avx`          | All amd64                      |
| AVX2 (xN)    | 4-32  | `avx2`         | Intel Haswell+, AMD Excavator+ |
| AVX-512 (xN) | 8-64  | `avx512f`      | Intel Skylake-X+, some Xeons   |

> **Note**: Choose the variant matching your CPU's capabilities. Higher lane counts provide better performance but require newer CPU support.

```go
// Using AVX2 variant (32 lanes at once) - Intel Haswell+ / AMD Excavator+
result := simd.ClampInt8x32([]int8{1, 5, 10, 15, 20}, 5, 15)
// []int8{5, 5, 10, 15, 15}
```

```go
// Using AVX-512 variant (16 lanes at once) - Intel Skylake-X+
result := simd.ClampFloat32x16([]float32{0.5, 1.5, 2.5, 3.5}, 1.0, 3.0)
// []float32{1.0, 1.5, 2.5, 3.0}
```

```go
// Using AVX variant (8 lanes at once) - works on all amd64
result := simd.ClampInt16x8([]int16{100, 150, 200, 250}, 120, 220)
// []int16{120, 150, 200, 220}
```

```go
// Empty collection returns empty collection
result := simd.ClampUint32x4([]uint32{}, 10, 100)
// []uint32{}
```
