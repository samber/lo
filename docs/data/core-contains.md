---
name: Contains
slug: contains
sourceRef: intersect.go#L5
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/W1EvyqY6t9j
variantHelpers:
  - core#intersect#contains
similarHelpers:
  - core#intersect#containsby
  - core#intersect#every
  - core#intersect#some
  - core#intersect#none
  - core#slice#find
position: 0
signatures:
  - "func Contains[T comparable](collection []T, element T) bool"
---

Returns true if an element is present in a collection.

```go
present := lo.Contains([]int{0, 1, 2, 3, 4, 5}, 5)
// true
```


