---
name: Intersect
slug: intersect
sourceRef: intersect.go#L103
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/uuElL9X9e58
variantHelpers:
  - core#intersect#intersect
similarHelpers:
  - core#intersect#difference
  - core#intersect#union
  - core#intersect#without
  - core#slice#uniq
position: 80
signatures:
  - "func Intersect[T comparable, Slice ~[]T](list1 Slice, list2 Slice) Slice"
---

Returns the intersection between two collections.

```go
lo.Intersect([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
// []int{0, 2}
```


