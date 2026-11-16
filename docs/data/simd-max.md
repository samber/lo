---
name: Max
slug: max
sourceRef: exp/simd/find.go#L122
category: experimental
subCategory: simd
similarHelpers:
  - experimental#simd#max
position: 30
signatures:
  - "func MaxInt8[T ~int8](collection []T) T"
  - "func MaxInt16[T ~int16](collection []T) T"
  - "func MaxInt32[T ~int32](collection []T) T"
  - "func MaxInt64[T ~int64](collection []T) T"
  - "func MaxUint8[T ~uint8](collection []T) T"
  - "func MaxUint16[T ~uint16](collection []T) T"
  - "func MaxUint32[T ~uint32](collection []T) T"
  - "func MaxUint64[T ~uint64](collection []T) T"
  - "func MaxFloat32[T ~float32](collection []T) T"
  - "func MaxFloat64[T ~float64](collection []T) T"
---

Finds the maximum value in a collection using SIMD instructions when available, falling back to a plain scalar loop everywhere else.

```go
max := simd.MaxInt32([]int32{100, 50, 200, 75})
// 200
```

```go
max := simd.MaxFloat32([]float32{3.5, 1.2, 4.8, 2.1})
// 4.8
```

```go
// Empty collection returns 0
max := simd.MaxUint16([]uint16{})
// 0
```
