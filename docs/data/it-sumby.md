---
name: SumBy
slug: sumby
sourceRef: it/math.go#L74
category: iter
subCategory: math
signatures:
  - "func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T], transform func(item T) R) R"
playUrl: https://go.dev/play/p/ZNiqXNMu5QP
variantHelpers:
  - iter#math#sumby
similarHelpers:
  - core#slice#sumby
  - core#slice#sum
position: 65
---

Returns the sum of values in the collection using the given transform function.

```go
type Person struct {
    Name string
    Age  int
}

people := it.Slice([]Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 35},
})

result := it.SumBy(people, func(p Person) int {
    return p.Age
})
// 90
```