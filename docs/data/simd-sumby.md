---
name: SumBy
slug: sumby
sourceRef: exp/simd/math.go#L841
category: exp
subCategory: simd
similarHelpers:
  - exp#simd#sum
  - exp#simd#meanby
position: 20
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
  - "func SumByInt8x16[T any, R ~int8](collection []T, iteratee func(item T) R) R"
  - "func SumByInt8x32[T any, R ~int8](collection []T, iteratee func(item T) R) R"
  - "func SumByInt8x64[T any, R ~int8](collection []T, iteratee func(item T) R) R"
  - "func SumByInt16x8[T any, R ~int16](collection []T, iteratee func(item T) R) R"
  - "func SumByInt16x16[T any, R ~int16](collection []T, iteratee func(item T) R) R"
  - "func SumByInt16x32[T any, R ~int16](collection []T, iteratee func(item T) R) R"
  - "func SumByInt32x4[T any, R ~int32](collection []T, iteratee func(item T) R) R"
  - "func SumByInt32x8[T any, R ~int32](collection []T, iteratee func(item T) R) R"
  - "func SumByInt32x16[T any, R ~int32](collection []T, iteratee func(item T) R) R"
  - "func SumByInt64x2[T any, R ~int64](collection []T, iteratee func(item T) R) R"
  - "func SumByInt64x4[T any, R ~int64](collection []T, iteratee func(item T) R) R"
  - "func SumByInt64x8[T any, R ~int64](collection []T, iteratee func(item T) R) R"
  - "func SumByUint8x16[T any, R ~uint8](collection []T, iteratee func(item T) R) R"
  - "func SumByUint8x32[T any, R ~uint8](collection []T, iteratee func(item T) R) R"
  - "func SumByUint8x64[T any, R ~uint8](collection []T, iteratee func(item T) R) R"
  - "func SumByUint16x8[T any, R ~uint16](collection []T, iteratee func(item T) R) R"
  - "func SumByUint16x16[T any, R ~uint16](collection []T, iteratee func(item T) R) R"
  - "func SumByUint16x32[T any, R ~uint16](collection []T, iteratee func(item T) R) R"
  - "func SumByUint32x4[T any, R ~uint32](collection []T, iteratee func(item T) R) R"
  - "func SumByUint32x8[T any, R ~uint32](collection []T, iteratee func(item T) R) R"
  - "func SumByUint32x16[T any, R ~uint32](collection []T, iteratee func(item T) R) R"
  - "func SumByUint64x2[T any, R ~uint64](collection []T, iteratee func(item T) R) R"
  - "func SumByUint64x4[T any, R ~uint64](collection []T, iteratee func(item T) R) R"
  - "func SumByUint64x8[T any, R ~uint64](collection []T, iteratee func(item T) R) R"
  - "func SumByFloat32x4[T any, R ~float32](collection []T, iteratee func(item T) R) R"
  - "func SumByFloat32x8[T any, R ~float32](collection []T, iteratee func(item T) R) R"
  - "func SumByFloat32x16[T any, R ~float32](collection []T, iteratee func(item T) R) R"
  - "func SumByFloat64x2[T any, R ~float64](collection []T, iteratee func(item T) R) R"
  - "func SumByFloat64x4[T any, R ~float64](collection []T, iteratee func(item T) R) R"
  - "func SumByFloat64x8[T any, R ~float64](collection []T, iteratee func(item T) R) R"
---

SumBy transforms a collection using an iteratee function and sums the result using SIMD instructions. The automatic dispatch functions (e.g., `SumByInt8`) will select the best SIMD variant based on CPU capabilities. The specific variants (e.g., `SumByInt8x32`) use a fixed SIMD instruction set regardless of CPU capabilities.

## Requirements

- **Go 1.26+** with `GOEXPERIMENT=simd`
- **amd64** architecture only

### CPU compatibility

| SIMD variant | Lanes | Required flags | Typical CPUs                   |
| ------------ | ----- | -------------- | ------------------------------ |
| AVX (xN)     | 2-16  | `avx`          | All amd64                      |
| AVX2 (xN)    | 4-32  | `avx2`         | Intel Haswell+, AMD Excavator+ |
| AVX-512 (xN) | 8-64  | `avx512f`      | Intel Skylake-X+, some Xeons   |

> **Note**: The automatic dispatch functions (e.g., `SumByInt8`) will use the best available SIMD variant for the current CPU. Use specific variants (e.g., `SumByInt8x32`) only if you know your target CPU supports that instruction set.

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

// Automatic dispatch - uses best available SIMD
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

// Sum stock value using specific AVX2 variant
sum := simd.SumByFloat32x8(products, func(p Product) float32 {
    return p.Price * float32(p.Stock)
})
// 152.5
```

```go
type Metric struct {
    Value uint16
}

metrics := []Metric{
    {Value: 100},
    {Value: 200},
    {Value: 300},
    {Value: 400},
}

// Using AVX variant - works on all amd64
sum := simd.SumByUint16x8(metrics, func(m Metric) uint16 {
    return m.Value
})
// 1000
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
