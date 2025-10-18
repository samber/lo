---
name: Last
slug: last
sourceRef: find.go#L585
category: core
subCategory: find
playUrl: https://go.dev/play/p/ul45Z0y2EFO
variantHelpers:
  - core#find#last
similarHelpers:
  - core#find#first
  - core#find#nth
  - core#find#lastor
  - core#find#lastorempty
position: 290
signatures:
  - "func Last[T any](collection []T) (T, bool)"
---

Returns the last element of a collection and whether it exists.

```go
v, ok := lo.Last([]int{1, 2, 3})
// v == 3, ok == true
```


