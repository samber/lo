---
name: IntersectBy
slug: intersectby
sourceRef: intersect.go#L174
category: core
subCategory: intersect
playUrl:
variantHelpers:
  - core#intersect#intersectby
similarHelpers:
  - core#intersect#intersect
  - it#intersect#intersect
  - it#intersect#intersectby
  - core#intersect#difference
  - core#intersect#union
  - core#intersect#without
  - core#slice#uniq
position: 80
signatures:
  - "func IntersectBy[T any, K comparable, Slice ~[]T](transform func(T) K, lists ...Slice) Slice"
---

Returns the intersection between two collections using a custom key selector function.

```go
transform := func(v int) string {
  return strconv.Itoa(v)
}

lo.IntersectBy(transform, []int{0, 3, 5, 7}, []int{3, 5}, []int{0, 1, 2, 0, 3, 0})
// []int{3}
```
