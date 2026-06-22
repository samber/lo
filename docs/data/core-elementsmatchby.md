---
name: ElementsMatchBy
slug: elementsmatchby
sourceRef: intersect.go#L255
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/XWSEM4Ic_t0
variantHelpers:
  - core#intersect#elementsmatchby
similarHelpers:
  - core#intersect#elementsmatch
  - core#intersect#containsby
  - core#intersect#contains
  - core#intersect#intersect
  - core#intersect#difference
  - core#intersect#union
  - core#intersect#everyby
  - core#intersect#every
  - core#intersect#someby
  - core#intersect#some
  - core#intersect#noneby
  - core#intersect#none
position: 160
signatures:
  - "func ElementsMatchBy[T any, K comparable](list1 []T, list2 []T, iteratee func(item T) K) bool"
---

Returns true if lists contain the same set of keys computed by the predicate, with matching multiplicities. Order is not checked.

```go
type Item struct{
    ID string
}

lo.ElementsMatchBy(
    []Item{{"a"}, {"b"}},
    []Item{{"b"}, {"a"}},
    func(i Item) string {
        return i.ID
    },
)
// true
```


