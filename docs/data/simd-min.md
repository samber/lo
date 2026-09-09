---
name: Min
slug: min
sourceRef: exp/simd/find.go#L9
category: experimental
subCategory: simd
similarHelpers:
  - experimental#simd#min
position: 20
signatures:
  - "func MinInt8[T ~int8](collection []T) T"
  - "func MinInt16[T ~int16](collection []T) T"
  - "func MinInt32[T ~int32](collection []T) T"
  - "func MinInt64[T ~int64](collection []T) T"
  - "func MinUint8[T ~uint8](collection []T) T"
  - "func MinUint16[T ~uint16](collection []T) T"
  - "func MinUint32[T ~uint32](collection []T) T"
  - "func MinUint64[T ~uint64](collection []T) T"
  - "func MinFloat32[T ~float32](collection []T) T"
  - "func MinFloat64[T ~float64](collection []T) T"
---

Finds the minimum value in a collection using SIMD instructions when available, falling back to a plain scalar loop everywhere else.

```go
min := simd.MinInt32([]int32{100, 50, 200, 75})
// 50
```

```go
min := simd.MinFloat32([]float32{3.5, 1.2, 4.8, 2.1})
// 1.2
```

```go
// Empty collection returns 0
min := simd.MinUint16([]uint16{})
// 0
```
