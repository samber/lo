---
name: First
slug: first
sourceRef: find.go#L554
category: core
subCategory: find
playUrl: https://go.dev/play/p/ul45Z0y2EFO
variantHelpers:
  - core#find#first
similarHelpers:
  - core#find#firstor
  - core#find#firstorempty
  - core#find#last
  - core#find#lastor
  - core#find#lastorempty
  - core#find#nth
position: 260
signatures:
  - "func First[T any](collection []T) (T, bool)"
---

Returns the first element of a collection and whether it exists.

```go
v, ok := lo.First([]int{1, 2, 3})
// v == 1, ok == true
```


