---
name: ReduceRight
slug: reduceright
sourceRef: slice.go#L97
category: core
subCategory: slice
playUrl: https://go.dev/play/p/Fq3W70l7wXF
variantHelpers:
  - core#slice#reduceright
similarHelpers:
  - core#slice#reduce
  - core#slice#sum
  - core#slice#product
  - core#slice#mean
  - core#slice#max
  - core#slice#min
position: 60
signatures:
  - "func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R"
---

Like Reduce except it iterates from right to left and accumulates into a single value.

```go
result := lo.ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, _ int) []int {
    return append(agg, item...)
}, []int{})
// []int{4, 5, 2, 3, 0, 1}
```


