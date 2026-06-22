---
name: EveryBy
slug: everyby
sourceRef: intersect.go#L41
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/dn1-vhHsq9x
variantHelpers:
  - core#intersect#everyby
similarHelpers:
  - core#intersect#every
  - core#intersect#someby
  - core#intersect#some
  - core#intersect#noneby
  - core#intersect#none
  - core#intersect#containsby
  - core#intersect#contains
  - core#intersect#elementsmatchby
  - core#intersect#elementsmatch
position: 30
signatures:
  - "func EveryBy[T any](collection []T, predicate func(item T) bool) bool"
---

Returns true if the predicate returns true for all elements in the collection, or if the collection is empty.

```go
ok := lo.EveryBy(
    []int{1, 2, 3, 4},
    func(x int) bool {
        return x < 5
    },
)
// true
```


