---
name: Sum
slug: sum
sourceRef: exp/simd/math.go#L11
category: experimental
subCategory: simd
similarHelpers:
  - experimental#simd#sum
  - experimental#simd#sumby
position: 0
signatures:
  - "func SumInt8[T ~int8](collection []T) T"
  - "func SumInt16[T ~int16](collection []T) T"
  - "func SumInt32[T ~int32](collection []T) T"
  - "func SumInt64[T ~int64](collection []T) T"
  - "func SumUint8[T ~uint8](collection []T) T"
  - "func SumUint16[T ~uint16](collection []T) T"
  - "func SumUint32[T ~uint32](collection []T) T"
  - "func SumUint64[T ~uint64](collection []T) T"
  - "func SumFloat32[T ~float32](collection []T) T"
  - "func SumFloat64[T ~float64](collection []T) T"
---

Sums the values in a collection using SIMD instructions when available (AVX/AVX2/AVX512 on amd64, NEON on arm64, SIMD128 on wasm), falling back to a plain scalar loop everywhere else.

```go
sum := simd.SumInt32([]int32{1000000, 2000000, 3000000})
// 6000000
```

```go
sum := simd.SumFloat32([]float32{1.1, 2.2, 3.3, 4.4})
// 11
```

```go
// Empty collection returns 0
sum := simd.SumUint16([]uint16{})
// 0
```
