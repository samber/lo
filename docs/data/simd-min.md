---
name: Min
slug: min
sourceRef: exp/simd/math_avx.go#L833
category: exp
subCategory: simd
similarHelpers:
  - exp#simd#min
position: 20
signatures:
  - "func MinInt8x16[T ~int8](collection []T) T"
  - "func MinInt8x32[T ~int8](collection []T) T"
  - "func MinInt8x64[T ~int8](collection []T) T"
  - "func MinInt16x8[T ~int16](collection []T) T"
  - "func MinInt16x16[T ~int16](collection []T) T"
  - "func MinInt16x32[T ~int16](collection []T) T"
  - "func MinInt32x4[T ~int32](collection []T) T"
  - "func MinInt32x8[T ~int32](collection []T) T"
  - "func MinInt32x16[T ~int32](collection []T) T"
  - "func MinInt64x2[T ~int64](collection []T) T"
  - "func MinInt64x4[T ~int64](collection []T) T"
  - "func MinInt64x8[T ~int64](collection []T) T"
  - "func MinUint8x16[T ~uint8](collection []T) T"
  - "func MinUint8x32[T ~uint8](collection []T) T"
  - "func MinUint8x64[T ~uint8](collection []T) T"
  - "func MinUint16x8[T ~uint16](collection []T) T"
  - "func MinUint16x16[T ~uint16](collection []T) T"
  - "func MinUint16x32[T ~uint16](collection []T) T"
  - "func MinUint32x4[T ~uint32](collection []T) T"
  - "func MinUint32x8[T ~uint32](collection []T) T"
  - "func MinUint32x16[T ~uint32](collection []T) T"
  - "func MinUint64x2[T ~uint64](collection []T) T"
  - "func MinUint64x4[T ~uint64](collection []T) T"
  - "func MinUint64x8[T ~uint64](collection []T) T"
  - "func MinFloat32x4[T ~float32](collection []T) T"
  - "func MinFloat32x8[T ~float32](collection []T) T"
  - "func MinFloat32x16[T ~float32](collection []T) T"
  - "func MinFloat64x2[T ~float64](collection []T) T"
  - "func MinFloat64x4[T ~float64](collection []T) T"
  - "func MinFloat64x8[T ~float64](collection []T) T"
---

Finds the minimum value in a collection using SIMD instructions. The suffix (x2, x4, x8, x16, x32, x64) indicates the number of lanes processed simultaneously.

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
min := simd.MinInt8x32([]int8{5, 2, 8, 1, 9})
// 1
```

```go
// Using AVX-512 variant (16 lanes at once) - Intel Skylake-X+
min := simd.MinFloat32x16([]float32{3.5, 1.2, 4.8, 2.1})
// 1.2
```

```go
// Using AVX variant (4 lanes at once) - works on all amd64
min := simd.MinInt32x4([]int32{100, 50, 200, 75})
// 50
```

```go
// Empty collection returns 0
min := simd.MinUint16x8([]uint16{})
// 0
```
