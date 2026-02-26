---
name: Max
slug: max
sourceRef: exp/simd/math_avx.go#L1279
category: exp
subCategory: simd
similarHelpers:
  - exp#simd#max
position: 30
signatures:
  - "func MaxInt8x16[T ~int8](collection []T) T"
  - "func MaxInt8x32[T ~int8](collection []T) T"
  - "func MaxInt8x64[T ~int8](collection []T) T"
  - "func MaxInt16x8[T ~int16](collection []T) T"
  - "func MaxInt16x16[T ~int16](collection []T) T"
  - "func MaxInt16x32[T ~int16](collection []T) T"
  - "func MaxInt32x4[T ~int32](collection []T) T"
  - "func MaxInt32x8[T ~int32](collection []T) T"
  - "func MaxInt32x16[T ~int32](collection []T) T"
  - "func MaxInt64x2[T ~int64](collection []T) T"
  - "func MaxInt64x4[T ~int64](collection []T) T"
  - "func MaxInt64x8[T ~int64](collection []T) T"
  - "func MaxUint8x16[T ~uint8](collection []T) T"
  - "func MaxUint8x32[T ~uint8](collection []T) T"
  - "func MaxUint8x64[T ~uint8](collection []T) T"
  - "func MaxUint16x8[T ~uint16](collection []T) T"
  - "func MaxUint16x16[T ~uint16](collection []T) T"
  - "func MaxUint16x32[T ~uint16](collection []T) T"
  - "func MaxUint32x4[T ~uint32](collection []T) T"
  - "func MaxUint32x8[T ~uint32](collection []T) T"
  - "func MaxUint32x16[T ~uint32](collection []T) T"
  - "func MaxUint64x2[T ~uint64](collection []T) T"
  - "func MaxUint64x4[T ~uint64](collection []T) T"
  - "func MaxUint64x8[T ~uint64](collection []T) T"
  - "func MaxFloat32x4[T ~float32](collection []T) T"
  - "func MaxFloat32x8[T ~float32](collection []T) T"
  - "func MaxFloat32x16[T ~float32](collection []T) T"
  - "func MaxFloat64x2[T ~float64](collection []T) T"
  - "func MaxFloat64x4[T ~float64](collection []T) T"
  - "func MaxFloat64x8[T ~float64](collection []T) T"
---

Finds the maximum value in a collection using SIMD instructions. The suffix (x2, x4, x8, x16, x32, x64) indicates the number of lanes processed simultaneously.

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
max := simd.MaxInt8x32([]int8{5, 2, 8, 1, 9})
// 9
```

```go
// Using AVX-512 variant (16 lanes at once) - Intel Skylake-X+
max := simd.MaxFloat32x16([]float32{3.5, 1.2, 4.8, 2.1})
// 4.8
```

```go
// Using AVX variant (4 lanes at once) - works on all amd64
max := simd.MaxInt32x4([]int32{100, 50, 200, 75})
// 200
```

```go
// Empty collection returns 0
max := simd.MaxUint16x8([]uint16{})
// 0
```
