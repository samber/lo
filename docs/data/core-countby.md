---
name: CountBy
slug: countby
sourceRef: slice.go#L596
category: core
subCategory: slice
playUrl: https://go.dev/play/p/ByQbNYQQi4X
variantHelpers:
  - core#slice#countby
similarHelpers:
  - core#slice#count
  - core#slice#every
  - core#slice#some
  - core#slice#filter
  - core#slice#find
position: 0
signatures:
  - "func CountBy[T any](collection []T, predicate func(item T) bool) int"
---

Counts the number of elements for which the predicate is true.

```go
lo.CountBy([]int{1, 5, 1}, func(i int) bool {
    return i < 4
})
// 2
```


