---
name: FirstOrEmpty
slug: firstorempty
sourceRef: find.go#L567
category: core
subCategory: find
playUrl: https://go.dev/play/p/ul45Z0y2EFO
variantHelpers:
  - core#find#firstorempty
similarHelpers:
  - core#find#first
  - core#find#firstor
  - core#find#lastorempty
  - core#find#nthorempty
position: 270
signatures:
  - "func FirstOrEmpty[T any](collection []T) T"
---

Returns the first element of a collection or the zero value if empty.

```go
v := lo.FirstOrEmpty([]int{})
// v == 0
```


