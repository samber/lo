---
name: Mean
slug: mean
sourceRef: exp/simd/math.go#L123
category: experimental
subCategory: simd
similarHelpers:
  - experimental#simd#mean
  - experimental#simd#meanby
position: 10
signatures:
  - "func MeanInt8[T ~int8](collection []T) T"
  - "func MeanInt16[T ~int16](collection []T) T"
  - "func MeanInt32[T ~int32](collection []T) T"
  - "func MeanInt64[T ~int64](collection []T) T"
  - "func MeanUint8[T ~uint8](collection []T) T"
  - "func MeanUint16[T ~uint16](collection []T) T"
  - "func MeanUint32[T ~uint32](collection []T) T"
  - "func MeanUint64[T ~uint64](collection []T) T"
  - "func MeanFloat32[T ~float32](collection []T) T"
  - "func MeanFloat64[T ~float64](collection []T) T"
---

Calculates the arithmetic mean of a collection using SIMD instructions when available, falling back to a plain scalar loop everywhere else.

```go
mean := simd.MeanInt16([]int16{10, 20, 30, 40})
// 25
```

```go
mean := simd.MeanFloat32([]float32{1.0, 2.0, 3.0, 4.0})
// 2.5
```

```go
// Empty collection returns 0
mean := simd.MeanUint32([]uint32{})
// 0
```
