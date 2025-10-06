---
name: NoneBy
slug: noneby
sourceRef: intersect.go#L91
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/O64WZ32H58S
variantHelpers:
  - core#intersect#noneby
similarHelpers:
  - core#intersect#none
  - core#intersect#everyby
  - core#intersect#every
  - core#intersect#someby
  - core#intersect#some
  - core#intersect#containsby
  - core#intersect#contains
position: 70
signatures:
  - "func NoneBy[T any](collection []T, predicate func(item T) bool) bool"
---

Returns true if the predicate returns true for none of the elements in the collection, or if the collection is empty.

```go
ok := lo.NoneBy(
    []int{1, 2, 3, 4},
    func(x int) bool {
        return x < 0
    },
)
// true
```


