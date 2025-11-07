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
  - core#intersect#intersectby
  - it#intersect#intersect
  - it#intersect#intersectby
  - core#intersect#difference
  - core#intersect#union
  - core#intersect#without
  - core#slice#uniq
position: 80
signatures:
  - "func Intersect[T comparable, Slice ~[]T](lists ...Slice) Slice"
---

Returns the intersection between collections.

```go
lo.Intersect([]int{0, 3, 5, 7}, []int{3, 5}, []int{0, 1, 2, 0, 3, 0})
// []int{3}
```
