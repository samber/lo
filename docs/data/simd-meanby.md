---
name: MeanBy
slug: meanby
sourceRef: exp/simd/math.go#L325
category: experimental
subCategory: simd
similarHelpers:
  - experimental#simd#mean
  - experimental#simd#sumby
position: 60
signatures:
  - "func MeanByInt8[T any, R ~int8](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt16[T any, R ~int16](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt32[T any, R ~int32](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt64[T any, R ~int64](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint8[T any, R ~uint8](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint16[T any, R ~uint16](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint32[T any, R ~uint32](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint64[T any, R ~uint64](collection []T, iteratee func(item T) R) R"
  - "func MeanByFloat32[T any, R ~float32](collection []T, iteratee func(item T) R) R"
  - "func MeanByFloat64[T any, R ~float64](collection []T, iteratee func(item T) R) R"
---

Transforms a collection using an iteratee function and calculates the arithmetic mean of the result using SIMD instructions when available, falling back to a plain scalar loop everywhere else.

```go
type Person struct {
    Name string
    Age  int8
}

people := []Person{
    {Name: "Alice", Age: 20},
    {Name: "Bob", Age: 30},
    {Name: "Charlie", Age: 40},
}

mean := simd.MeanByInt8(people, func(p Person) int8 {
    return p.Age
})
// 30
```

```go
type Product struct {
    Name  string
    Price float32
}

products := []Product{
    {Name: "Widget", Price: 10.50},
    {Name: "Gadget", Price: 20.00},
    {Name: "Tool", Price: 15.75},
}

mean := simd.MeanByFloat32(products, func(p Product) float32 {
    return p.Price
})
// 15.4167
```

```go
// Empty collection returns 0
type Item struct {
    Count int64
}

mean := simd.MeanByInt64([]Item{}, func(i Item) int64 {
    return i.Count
})
// 0
```
