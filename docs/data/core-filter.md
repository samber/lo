---
name: Filter
slug: filter
sourceRef: slice.go#L12
category: core
subCategory: slice
playUrl: https://go.dev/play/p/Apjg3WeSi7K
similarHelpers:
  - core#slice#reject
  - core#slice#filtermap
  - core#slice#filterreject
  - core#slice#rejectmap
  - core#slice#filtertake
  - parallel#slice#filter
  - mutable#slice#filter
variantHelpers:
  - core#slice#filter
position: 0
signatures:
  - "func Filter[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice"
---

Iterates over a collection and returns a slice of all the elements the predicate function returns `true` for.

```go
even := lo.Filter([]int{1, 2, 3, 4}, func(x int, index int) bool {
    return x%2 == 0
})
// []int{2, 4}
```
