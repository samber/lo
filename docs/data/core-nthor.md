---
name: NthOr
slug: nthor
sourceRef: find.go#L635
category: core
subCategory: find
playUrl: https://go.dev/play/p/sHoh88KWt6B
variantHelpers:
  - core#find#nthor
similarHelpers:
  - core#find#nthorempty
  - core#find#nth
  - core#find#findorelse
  - core#find#firstor
position: 330
signatures:
  - "func NthOr[T any, N constraints.Integer](collection []T, nth N, fallback T) T"
---

Returns the element at index nth of collection, or the fallback if out of bounds. If nth is negative, returns the nth element from the end.

```go
v := lo.NthOr([]int{10, 20, 30}, 10, -1)
// v == -1
```


