---
name: IndexOf
slug: indexof
sourceRef: find.go#L14
category: core
subCategory: find
playUrl: https://go.dev/play/p/Eo7W0lvKTky
variantHelpers:
  - core#find#indexof
similarHelpers:
  - core#find#findindexof
  - core#find#findlastindexof
  - core#find#lastindexof
  - core#find#find
position: 0
signatures:
  - "func IndexOf[T comparable](collection []T, element T) int"
---

Returns the index of the first occurrence of a value in a slice, or -1 if not found.

```go
idx := lo.IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
// 2

idx = lo.IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)
// -1
```


