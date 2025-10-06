---
name: Without
slug: without
sourceRef: intersect.go#L181
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/5j30Ux8TaD0
variantHelpers:
  - core#intersect#without
similarHelpers:
  - core#intersect#withoutby
  - core#intersect#intersect
  - core#intersect#difference
  - core#intersect#union
  - core#slice#reject
position: 110
signatures:
  - "func Without[T comparable, Slice ~[]T](collection Slice, exclude ...T) Slice"
---

Returns a slice excluding all given values.

```go
lo.Without([]int{0, 2, 10}, 2)
// []int{0, 10}
```


