---
name: FirstOr
slug: firstor
sourceRef: find.go#L574
category: core
subCategory: find
playUrl: https://go.dev/play/p/ul45Z0y2EFO
variantHelpers:
  - core#find#firstor
similarHelpers:
  - core#find#first
  - core#find#firstorempty
  - core#find#lastor
  - core#find#nthor
position: 280
signatures:
  - "func FirstOr[T any](collection []T, fallback T) T"
---

Returns the first element of a collection or the fallback value if empty.

```go
v := lo.FirstOr([]int{}, -1)
// v == -1
```


