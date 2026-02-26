---
name: Mean
slug: mean
sourceRef: exp/simd/math_avx.go#L352
category: exp
subCategory: simd
similarHelpers:
  - exp#simd#mean
  - exp#simd#meanby
position: 10
signatures:
  - "func MeanInt8x16[T ~int8](collection []T) T"
  - "func MeanInt8x32[T ~int8](collection []T) T"
  - "func MeanInt8x64[T ~int8](collection []T) T"
  - "func MeanInt16x8[T ~int16](collection []T) T"
  - "func MeanInt16x16[T ~int16](collection []T) T"
  - "func MeanInt16x32[T ~int16](collection []T) T"
  - "func MeanInt32x4[T ~int32](collection []T) T"
  - "func MeanInt32x8[T ~int32](collection []T) T"
  - "func MeanInt32x16[T ~int32](collection []T) T"
  - "func MeanInt64x2[T ~int64](collection []T) T"
  - "func MeanInt64x4[T ~int64](collection []T) T"
  - "func MeanInt64x8[T ~int64](collection []T) T"
  - "func MeanUint8x16[T ~uint8](collection []T) T"
  - "func MeanUint8x32[T ~uint8](collection []T) T"
  - "func MeanUint8x64[T ~uint8](collection []T) T"
  - "func MeanUint16x8[T ~uint16](collection []T) T"
  - "func MeanUint16x16[T ~uint16](collection []T) T"
  - "func MeanUint16x32[T ~uint16](collection []T) T"
  - "func MeanUint32x4[T ~uint32](collection []T) T"
  - "func MeanUint32x8[T ~uint32](collection []T) T"
  - "func MeanUint32x16[T ~uint32](collection []T) T"
  - "func MeanUint64x2[T ~uint64](collection []T) T"
  - "func MeanUint64x4[T ~uint64](collection []T) T"
  - "func MeanUint64x8[T ~uint64](collection []T) T"
  - "func MeanFloat32x4[T ~float32](collection []T) T"
  - "func MeanFloat32x8[T ~float32](collection []T) T"
  - "func MeanFloat32x16[T ~float32](collection []T) T"
  - "func MeanFloat64x2[T ~float64](collection []T) T"
  - "func MeanFloat64x4[T ~float64](collection []T) T"
  - "func MeanFloat64x8[T ~float64](collection []T) T"
---

Calculates the arithmetic mean of a collection using SIMD instructions. The suffix (x2, x4, x8, x16, x32, x64) indicates the number of lanes processed simultaneously.

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
mean := simd.MeanInt8x32([]int8{1, 2, 3, 4, 5})
// 3
```

```go
// Using AVX-512 variant (16 lanes at once) - Intel Skylake-X+
mean := simd.MeanFloat32x16([]float32{1.0, 2.0, 3.0, 4.0})
// 2.5
```

```go
// Using AVX variant (8 lanes at once) - works on all amd64
mean := simd.MeanInt16x8([]int16{10, 20, 30, 40})
// 25
```

```go
// Empty collection returns 0
mean := simd.MeanUint32x4([]uint32{})
// 0
```
