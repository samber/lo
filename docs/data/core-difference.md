---
name: Difference
slug: difference
sourceRef: intersect.go#L124
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/pKE-JgzqRpz
variantHelpers:
  - core#intersect#difference
similarHelpers:
  - core#intersect#intersect
  - core#intersect#union
  - core#intersect#without
  - core#intersect#withoutby
  - core#slice#uniq
  - core#slice#uniqby
position: 90
signatures:
  - "func Difference[T comparable, Slice ~[]T](list1 Slice, list2 Slice) (Slice, Slice)"
---

Returns the difference between two collections. The first slice contains elements absent from list2; the second contains elements absent from list1.

```go
left, right := lo.Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
// []int{1, 3, 4, 5}, []int{6}
```


