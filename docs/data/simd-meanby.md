---
name: MeanBy
slug: meanby
sourceRef: exp/simd/math.go#L1006
category: exp
subCategory: simd
similarHelpers:
  - exp#simd#mean
  - exp#simd#sumby
position: 30
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
  - "func MeanByInt8x16[T any, R ~int8](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt8x32[T any, R ~int8](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt8x64[T any, R ~int8](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt16x8[T any, R ~int16](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt16x16[T any, R ~int16](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt16x32[T any, R ~int16](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt32x4[T any, R ~int32](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt32x8[T any, R ~int32](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt32x16[T any, R ~int32](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt64x2[T any, R ~int64](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt64x4[T any, R ~int64](collection []T, iteratee func(item T) R) R"
  - "func MeanByInt64x8[T any, R ~int64](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint8x16[T any, R ~uint8](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint8x32[T any, R ~uint8](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint8x64[T any, R ~uint8](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint16x8[T any, R ~uint16](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint16x16[T any, R ~uint16](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint16x32[T any, R ~uint16](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint32x4[T any, R ~uint32](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint32x8[T any, R ~uint32](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint32x16[T any, R ~uint32](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint64x2[T any, R ~uint64](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint64x4[T any, R ~uint64](collection []T, iteratee func(item T) R) R"
  - "func MeanByUint64x8[T any, R ~uint64](collection []T, iteratee func(item T) R) R"
  - "func MeanByFloat32x4[T any, R ~float32](collection []T, iteratee func(item T) R) R"
  - "func MeanByFloat32x8[T any, R ~float32](collection []T, iteratee func(item T) R) R"
  - "func MeanByFloat32x16[T any, R ~float32](collection []T, iteratee func(item T) R) R"
  - "func MeanByFloat64x2[T any, R ~float64](collection []T, iteratee func(item T) R) R"
  - "func MeanByFloat64x4[T any, R ~float64](collection []T, iteratee func(item T) R) R"
  - "func MeanByFloat64x8[T any, R ~float64](collection []T, iteratee func(item T) R) R"
---

MeanBy transforms a collection using an iteratee function and calculates the arithmetic mean of the result using SIMD instructions. The automatic dispatch functions (e.g., `MeanByInt8`) will select the best SIMD variant based on CPU capabilities. The specific variants (e.g., `MeanByInt8x32`) use a fixed SIMD instruction set regardless of CPU capabilities.

## Requirements

- **Go 1.26+** with `GOEXPERIMENT=simd`
- **amd64** architecture only

### CPU compatibility

| SIMD variant | Lanes | Required flags | Typical CPUs                   |
| ------------ | ----- | -------------- | ------------------------------ |
| AVX (xN)     | 2-16  | `avx`          | All amd64                      |
| AVX2 (xN)    | 4-32  | `avx2`         | Intel Haswell+, AMD Excavator+ |
| AVX-512 (xN) | 8-64  | `avx512f`      | Intel Skylake-X+, some Xeons   |

> **Note**: The automatic dispatch functions (e.g., `MeanByInt8`) will use the best available SIMD variant for the current CPU. Use specific variants (e.g., `MeanByInt8x32`) only if you know your target CPU supports that instruction set.

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

// Automatic dispatch - uses best available SIMD
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

// Mean price using specific AVX2 variant
mean := simd.MeanByFloat32x8(products, func(p Product) float32 {
    return p.Price
})
// 15.4167
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
mean := simd.MeanByUint16x8(metrics, func(m Metric) uint16 {
    return m.Value
})
// 250
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
