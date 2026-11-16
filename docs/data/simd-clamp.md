---
name: Clamp
slug: clamp
sourceRef: exp/simd/math.go#L418
category: experimental
subCategory: simd
similarHelpers:
  - experimental#simd#clamp
position: 40
signatures:
  - "func ClampInt8[T ~int8, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampInt16[T ~int16, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampInt32[T ~int32, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampInt64[T ~int64, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampUint8[T ~uint8, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampUint16[T ~uint16, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampUint32[T ~uint32, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampUint64[T ~uint64, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampFloat32[T ~float32, Slice ~[]T](collection Slice, mn, mx T) Slice"
  - "func ClampFloat64[T ~float64, Slice ~[]T](collection Slice, mn, mx T) Slice"
---

Clamps each element in a collection between mn and mx using SIMD instructions when available, falling back to a plain scalar loop everywhere else. If mn > mx, every element clamps to mn.

```go
result := simd.ClampInt8([]int8{1, 5, 10, 15, 20}, 5, 15)
// []int8{5, 5, 10, 15, 15}
```

```go
result := simd.ClampFloat32([]float32{0.5, 1.5, 2.5, 3.5}, 1.0, 3.0)
// []float32{1.0, 1.5, 2.5, 3.0}
```

```go
// Empty collection returns an empty collection
result := simd.ClampUint32([]uint32{}, 10, 100)
// []uint32{}
```
