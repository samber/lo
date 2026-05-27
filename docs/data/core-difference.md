---
name: Difference
slug: difference
sourceRef: intersect.go#L210
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
  - "func Difference[T comparable, Slice ~[]T](left, right Slice) (notInRight, notInLeft Slice)"
---

Returns the difference between two collections. The first slice contains elements from left absent from right; the second contains elements from right absent from left.

```go
left := []int{0, 1, 2, 3, 4, 5}
right := []int{0, 2, 6}

notInRight, notInLeft := lo.Difference(left, right)
// []int{1, 3, 4, 5}, []int{6}
```

