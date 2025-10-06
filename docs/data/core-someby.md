---
name: SomeBy
slug: someby
sourceRef: intersect.go#L67
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/DXF-TORBudx
variantHelpers:
  - core#intersect#someby
similarHelpers:
  - core#intersect#some
  - core#intersect#everyby
  - core#intersect#every
  - core#intersect#noneby
  - core#intersect#none
  - core#intersect#containsby
  - core#intersect#contains
position: 50
signatures:
  - "func SomeBy[T any](collection []T, predicate func(item T) bool) bool"
---

Returns true if the predicate returns true for any element in the collection. Returns false for an empty collection.

```go
ok := lo.SomeBy(
    []int{1, 2, 3, 4},
    func(x int) bool {
        return x < 3
    },
)
// true
```


