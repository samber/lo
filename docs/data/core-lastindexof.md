---
name: LastIndexOf
slug: lastindexof
sourceRef: find.go#L27
category: core
subCategory: find
playUrl: https://go.dev/play/p/Eo7W0lvKTky
variantHelpers:
  - core#find#lastindexof
similarHelpers:
  - core#find#indexof
  - core#find#findkey
  - core#find#findlastindexof
position: 10
signatures:
  - "func LastIndexOf[T comparable](collection []T, element T) int"
---

Returns the index of the last occurrence of a value in a slice, or -1 if not found.

```go
idx := lo.LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
// 4

idx = lo.LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)
// -1
```


