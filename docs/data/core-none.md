---
name: None
slug: none
sourceRef: intersect.go#L79
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/fye7JsmxzPV
variantHelpers:
  - core#intersect#none
similarHelpers:
  - core#intersect#noneby
  - core#intersect#contains
  - core#intersect#every
  - core#intersect#some
  - core#intersect#containsby
position: 60
signatures:
  - "func None[T comparable](collection []T, subset []T) bool"
---

Returns true if no element of a subset is contained in a collection, or if the subset is empty.

```go
ok := lo.None([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
// true
```


