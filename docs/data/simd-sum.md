---
name: Sum
slug: sum
sourceRef: exp/simd/math_avx.go#L14
category: exp
subCategory: simd
similarHelpers:
  - exp#simd#sum
  - exp#simd#sumby
position: 0
signatures:
  - "func SumInt8x16[T ~int8](collection []T) T"
  - "func SumInt8x32[T ~int8](collection []T) T"
  - "func SumInt8x64[T ~int8](collection []T) T"
  - "func SumInt16x8[T ~int16](collection []T) T"
  - "func SumInt16x16[T ~int16](collection []T) T"
  - "func SumInt16x32[T ~int16](collection []T) T"
  - "func SumInt32x4[T ~int32](collection []T) T"
  - "func SumInt32x8[T ~int32](collection []T) T"
  - "func SumInt32x16[T ~int32](collection []T) T"
  - "func SumInt64x2[T ~int64](collection []T) T"
  - "func SumInt64x4[T ~int64](collection []T) T"
  - "func SumInt64x8[T ~int64](collection []T) T"
  - "func SumUint8x16[T ~uint8](collection []T) T"
  - "func SumUint8x32[T ~uint8](collection []T) T"
  - "func SumUint8x64[T ~uint8](collection []T) T"
  - "func SumUint16x8[T ~uint16](collection []T) T"
  - "func SumUint16x16[T ~uint16](collection []T) T"
  - "func SumUint16x32[T ~uint16](collection []T) T"
  - "func SumUint32x4[T ~uint32](collection []T) T"
  - "func SumUint32x8[T ~uint32](collection []T) T"
  - "func SumUint32x16[T ~uint32](collection []T) T"
  - "func SumUint64x2[T ~uint64](collection []T) T"
  - "func SumUint64x4[T ~uint64](collection []T) T"
  - "func SumUint64x8[T ~uint64](collection []T) T"
  - "func SumFloat32x4[T ~float32](collection []T) T"
  - "func SumFloat32x8[T ~float32](collection []T) T"
  - "func SumFloat32x16[T ~float32](collection []T) T"
  - "func SumFloat64x2[T ~float64](collection []T) T"
  - "func SumFloat64x4[T ~float64](collection []T) T"
  - "func SumFloat64x8[T ~float64](collection []T) T"
---

Sums the values in a collection using SIMD instructions. The suffix (x2, x4, x8, x16, x32, x64) indicates the number of lanes processed simultaneously.

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
sum := simd.SumInt8x32([]int8{1, 2, 3, 4, 5})
// 15
```

```go
// Using AVX-512 variant (16 lanes at once) - Intel Skylake-X+
sum := simd.SumFloat32x16([]float32{1.1, 2.2, 3.3, 4.4})
// 11
```

```go
// Using AVX variant (4 lanes at once) - works on all amd64
sum := simd.SumInt32x4([]int32{1000000, 2000000, 3000000})
// 6000000
```

```go
// Empty collection returns 0
sum := simd.SumUint16x16([]uint16{})
// 0
```
