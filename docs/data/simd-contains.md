---
name: Contains
slug: contains
sourceRef: exp/simd/intersect.go#L9
category: experimental
subCategory: simd
similarHelpers:
  - experimental#simd#contains
position: 70
signatures:
  - "func ContainsInt8[T ~int8](collection []T, target T) bool"
  - "func ContainsInt16[T ~int16](collection []T, target T) bool"
  - "func ContainsInt32[T ~int32](collection []T, target T) bool"
  - "func ContainsInt64[T ~int64](collection []T, target T) bool"
  - "func ContainsUint8[T ~uint8](collection []T, target T) bool"
  - "func ContainsUint16[T ~uint16](collection []T, target T) bool"
  - "func ContainsUint32[T ~uint32](collection []T, target T) bool"
  - "func ContainsUint64[T ~uint64](collection []T, target T) bool"
  - "func ContainsFloat32[T ~float32](collection []T, target T) bool"
  - "func ContainsFloat64[T ~float64](collection []T, target T) bool"
---

Checks if a target value is present in a collection using SIMD instructions when available, falling back to a plain scalar loop everywhere else.

```go
found := simd.ContainsInt32([]int32{1, 2, 3, 4, 5}, 3)
// true
```

```go
found := simd.ContainsFloat32([]float32{1.1, 2.2, 3.3, 4.4}, 3.3)
// true
```

```go
// Empty collection returns false
found := simd.ContainsInt16([]int16{}, 5)
// false
```
