---
name: Some
slug: some
sourceRef: intersect.go#L54
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/Lj4ceFkeT9V
variantHelpers:
  - core#intersect#some
similarHelpers:
  - core#intersect#someby
  - core#intersect#every
  - core#intersect#everyby
  - core#intersect#none
  - core#intersect#noneby
  - core#intersect#contains
  - core#intersect#containsby
position: 40
signatures:
  - "func Some[T comparable](collection []T, subset []T) bool"
---

Returns true if at least one element of a subset is contained in a collection. Returns false for an empty subset.

```go
ok := lo.Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
// true
```


