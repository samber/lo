---
name: Union
slug: union
sourceRef: intersect.go#L157
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/DI9RVEB_qMK
variantHelpers:
  - core#intersect#union
similarHelpers:
  - core#intersect#intersect
  - core#intersect#difference
  - core#intersect#without
  - core#slice#uniq
  - core#slice#uniqby
position: 100
signatures:
  - "func Union[T comparable, Slice ~[]T](lists ...Slice) Slice"
---

Returns all distinct elements from given collections while preserving relative order.

```go
lo.Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2}, []int{0, 10})
// []int{0, 1, 2, 3, 4, 5, 10}
```


