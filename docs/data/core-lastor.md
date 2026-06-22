---
name: LastOr
slug: lastor
sourceRef: find.go#L605
category: core
subCategory: find
playUrl: https://go.dev/play/p/ul45Z0y2EFO
variantHelpers:
  - core#find#lastor
similarHelpers:
  - core#find#firstor
  - core#find#last
  - core#find#lastorempty
  - core#find#nthor
position: 310
signatures:
  - "func LastOr[T any](collection []T, fallback T) T"
---

Returns the last element of a collection or the fallback value if empty.

```go
v := lo.LastOr([]int{}, -1)
// v == -1
```


