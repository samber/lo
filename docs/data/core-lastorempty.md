---
name: LastOrEmpty
slug: lastorempty
sourceRef: find.go#L598
category: core
subCategory: find
playUrl: https://go.dev/play/p/ul45Z0y2EFO
variantHelpers:
  - core#find#lastorempty
similarHelpers:
  - core#find#firstorempty
  - core#find#last
  - core#find#lastor
  - core#find#nthorempty
position: 300
signatures:
  - "func LastOrEmpty[T any](collection []T) T"
---

Returns the last element of a collection or the zero value if empty.

```go
v := lo.LastOrEmpty([]int{})
// v == 0
```


