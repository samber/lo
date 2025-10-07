---
name: MeanBy
slug: meanby
sourceRef: it/math.go#L106
category: it
subCategory: math
signatures:
  - "func MeanBy[T any, R constraints.Float | constraints.Integer](collection iter.Seq[T], transform func(item T) R) R"
playUrl:
variantHelpers:
  - it#math#mean
similarHelpers:
  - core#slice#meanby
  - core#slice#mean
position: 70
---

Returns the mean value of the collection using the given transform function.

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

result := it.MeanBy(people, func(p Person) int {
    return p.Age
})
// 30.0
```