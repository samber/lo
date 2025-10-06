---
name: RangeWithSteps
slug: rangewithsteps
sourceRef: it/math.go#L35
category: it
subCategory: math
signatures:
  - "func RangeWithSteps[T constraints.Integer | constraints.Float](start, end, step T) iter.Seq[T]"
playUrl: ""
variantHelpers:
  - it#math#rangewithsteps
similarHelpers:
  - core#slice#rangewithsteps
  - it#math#range
  - it#math#rangefrom
position: 20
---

Creates a sequence of numbers from start up to (excluding) end with a custom step. Step set to zero will return an empty sequence.

```go
seq := it.RangeWithSteps(0, 10, 3)
var result []int
for item := range seq {
    result = append(result, item)
}
// result contains [0, 3, 6, 9]

seq2 := it.RangeWithSteps(10, 1, -3)
var result2 []int
for item := range seq2 {
    result2 = append(result2, item)
}
// result2 contains [10, 7, 4, 1]

seq3 := it.RangeWithSteps(0, 5, 1.5)
var result3 []float64
for item := range seq3 {
    result3 = append(result3, item)
}
// result3 contains [0, 1.5, 3, 4.5]
```