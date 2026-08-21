---
name: ContainsBy
slug: containsby
sourceRef: intersect.go#L17
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/W1EvyqY6t9j
variantHelpers:
  - core#intersect#containsby
similarHelpers:
  - core#intersect#contains
  - core#intersect#some
  - core#intersect#every
  - core#intersect#none
  - core#slice#some
  - core#slice#every
  - core#slice#findby
  - core#slice#filter
position: 10
signatures:
  - "func ContainsBy[T any](collection []T, predicate func(item T) bool) bool"
---

Returns true if the predicate returns true for any element in the collection.

```go
exists := lo.ContainsBy(
    []int{0, 1, 2, 3},
    func(x int) bool {
        return x == 3
    },
)
// true
```


