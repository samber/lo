---
name: Every
slug: every
sourceRef: intersect.go#L29
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/W1EvyqY6t9j
variantHelpers:
  - core#intersect#every
similarHelpers:
  - core#intersect#everyby
  - core#intersect#some
  - core#intersect#someby
  - core#intersect#none
  - core#intersect#noneby
  - core#intersect#contains
  - core#intersect#containsby
  - core#intersect#elementsmatch
  - core#intersect#elementsmatchby
position: 20
signatures:
  - "func Every[T comparable](collection []T, subset []T) bool"
---

Returns true if all elements of a subset are contained in a collection, or if the subset is empty.

```go
ok := lo.Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
// true
```


