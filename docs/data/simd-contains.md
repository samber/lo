---
name: Contains
slug: contains
sourceRef: exp/simd/intersect_avx512.go#L9
category: exp
subCategory: simd
similarHelpers:
  - exp#simd#contains
position: 0
signatures:
  - "func ContainsInt8x16[T ~int8](collection []T, target T) bool"
  - "func ContainsInt8x32[T ~int8](collection []T, target T) bool"
  - "func ContainsInt8x64[T ~int8](collection []T, target T) bool"
  - "func ContainsInt16x8[T ~int16](collection []T, target T) bool"
  - "func ContainsInt16x16[T ~int16](collection []T, target T) bool"
  - "func ContainsInt16x32[T ~int16](collection []T, target T) bool"
  - "func ContainsInt32x4[T ~int32](collection []T, target T) bool"
  - "func ContainsInt32x8[T ~int32](collection []T, target T) bool"
  - "func ContainsInt32x16[T ~int32](collection []T, target T) bool"
  - "func ContainsInt64x2[T ~int64](collection []T, target T) bool"
  - "func ContainsInt64x4[T ~int64](collection []T, target T) bool"
  - "func ContainsInt64x8[T ~int64](collection []T, target T) bool"
  - "func ContainsUint8x16[T ~uint8](collection []T, target T) bool"
  - "func ContainsUint8x32[T ~uint8](collection []T, target T) bool"
  - "func ContainsUint8x64[T ~uint8](collection []T, target T) bool"
  - "func ContainsUint16x8[T ~uint16](collection []T, target T) bool"
  - "func ContainsUint16x16[T ~uint16](collection []T, target T) bool"
  - "func ContainsUint16x32[T ~uint16](collection []T, target T) bool"
  - "func ContainsUint32x4[T ~uint32](collection []T, target T) bool"
  - "func ContainsUint32x8[T ~uint32](collection []T, target T) bool"
  - "func ContainsUint32x16[T ~uint32](collection []T, target T) bool"
  - "func ContainsUint64x2[T ~uint64](collection []T, target T) bool"
  - "func ContainsUint64x4[T ~uint64](collection []T, target T) bool"
  - "func ContainsUint64x8[T ~uint64](collection []T, target T) bool"
  - "func ContainsFloat32x4[T ~float32](collection []T, target T) bool"
  - "func ContainsFloat32x8[T ~float32](collection []T, target T) bool"
  - "func ContainsFloat32x16[T ~float32](collection []T, target T) bool"
  - "func ContainsFloat64x2[T ~float64](collection []T, target T) bool"
  - "func ContainsFloat64x4[T ~float64](collection []T, target T) bool"
  - "func ContainsFloat64x8[T ~float64](collection []T, target T) bool"
---

Checks if a target value is present in a collection using SIMD instructions. The suffix (x4, x8, x16, x32, x64) indicates the number of lanes processed simultaneously.

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
found := simd.ContainsInt8x32([]int8{1, 2, 3, 4, 5}, 3)
// true
```

```go
// Using AVX variant (16 lanes at once) - works on all amd64
found := simd.ContainsInt64x2([]int64{1000000, 2000000, 3000000}, 2000000)
// true
```

```go
// Using AVX-512 variant (64 lanes at once) - Intel Skylake-X+
found := simd.ContainsUint8x64([]uint8{10, 20, 30, 40, 50}, 30)
// true
```

```go
// Float32 with AVX2 (8 lanes at once)
found := simd.ContainsFloat32x8([]float32{1.1, 2.2, 3.3, 4.4}, 3.3)
// true
```

```go
// Empty collection returns false
found := simd.ContainsInt16x16([]int16{}, 5)
// false
```
