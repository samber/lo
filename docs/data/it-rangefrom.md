---
name: RangeFrom
slug: rangefrom
sourceRef: it/math.go#L25
category: it
subCategory: math
signatures:
  - "func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) iter.Seq[T]"
playUrl: ""
variantHelpers:
  - it#math#rangefrom
similarHelpers:
  - core#slice#rangefrom
  - it#math#range
  - it#math#rangewithsteps
position: 10
---

Creates a sequence of numbers from start with specified length. Yields `elementNum` values starting from `start`, stepping by ±1 depending on sign.

```go
seq := it.RangeFrom(5, 4)
var result []int
for item := range seq {
    result = append(result, item)
}
// result contains [5, 6, 7, 8]

seq2 := it.RangeFrom(10.5, 3)
var result2 []float64
for item := range seq2 {
    result2 = append(result2, item)
}
// result2 contains [10.5, 11.5, 12.5]
```