---
name: SumBy
slug: sumby
sourceRef: exp/simd/math.go#L214
category: experimental
subCategory: simd
similarHelpers:
  - experimental#simd#sum
  - experimental#simd#meanby
position: 50
signatures:
  - "func SumByInt8[T any, R ~int8](collection []T, iteratee func(item T) R) R"
  - "func SumByInt16[T any, R ~int16](collection []T, iteratee func(item T) R) R"
  - "func SumByInt32[T any, R ~int32](collection []T, iteratee func(item T) R) R"
  - "func SumByInt64[T any, R ~int64](collection []T, iteratee func(item T) R) R"
  - "func SumByUint8[T any, R ~uint8](collection []T, iteratee func(item T) R) R"
  - "func SumByUint16[T any, R ~uint16](collection []T, iteratee func(item T) R) R"
  - "func SumByUint32[T any, R ~uint32](collection []T, iteratee func(item T) R) R"
  - "func SumByUint64[T any, R ~uint64](collection []T, iteratee func(item T) R) R"
  - "func SumByFloat32[T any, R ~float32](collection []T, iteratee func(item T) R) R"
  - "func SumByFloat64[T any, R ~float64](collection []T, iteratee func(item T) R) R"
---

Transforms a collection using an iteratee function and sums the result using SIMD instructions when available, falling back to a plain scalar loop everywhere else.

```go
type Person struct {
    Name string
    Age  int8
}

people := []Person{
    {Name: "Alice", Age: 25},
    {Name: "Bob", Age: 30},
    {Name: "Charlie", Age: 35},
}

sum := simd.SumByInt8(people, func(p Person) int8 {
    return p.Age
})
// 90
```

```go
type Product struct {
    Name  string
    Price float32
    Stock int32
}

products := []Product{
    {Name: "Widget", Price: 10.50, Stock: 5},
    {Name: "Gadget", Price: 20.00, Stock: 3},
    {Name: "Tool", Price: 15.75, Stock: 2},
}

sum := simd.SumByFloat32(products, func(p Product) float32 {
    return p.Price * float32(p.Stock)
})
// 152.5
```

```go
// Empty collection returns 0
type Item struct {
    Count int64
}

sum := simd.SumByInt64([]Item{}, func(i Item) int64 {
    return i.Count
})
// 0
```
