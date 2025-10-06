---
name: RejectMap
slug: rejectmap
sourceRef: slice.go#L550
category: core
subCategory: slice
playUrl: https://go.dev/play/p/W9Ug9r0QFk
variantHelpers:
  - core#slice#rejectmap
similarHelpers:
  - core#slice#filtermap
  - core#slice#map
  - core#slice#filter
  - core#slice#reject
position: 270
signatures:
  - "func RejectMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R"
---

Opposite of FilterMap: maps each item and includes results where the predicate returned false.

```go
items := lo.RejectMap([]int{1, 2, 3, 4}, func(x int, _ int) (int, bool) {
    return x * 10, x%2 == 0
})
// []int{10, 30}
```


