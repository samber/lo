---
name: ElementsMatch
slug: elementsmatch
sourceRef: intersect.go#L247
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/XWSEM4Ic_t0
variantHelpers:
  - core#intersect#elementsmatch
similarHelpers:
  - core#intersect#elementsmatchby
  - core#intersect#contains
  - core#intersect#containsby
  - core#intersect#intersect
  - core#intersect#difference
  - core#intersect#union
  - core#intersect#every
  - core#intersect#everyby
  - core#intersect#some
  - core#intersect#someby
  - core#intersect#none
  - core#intersect#noneby
position: 150
signatures:
  - "func ElementsMatch[T comparable, Slice ~[]T](list1 Slice, list2 Slice) bool"
---

Returns true if lists contain the same set of elements (including empty set). If there are duplicate elements, occurrences must match. Order is not checked.

```go
lo.ElementsMatch([]int{1, 1, 2}, []int{2, 1, 1})
// true
```


