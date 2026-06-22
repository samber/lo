---
name: NthOrEmpty
slug: nthorempty
sourceRef: find.go#L647
category: core
subCategory: find
playUrl: https://go.dev/play/p/sHoh88KWt6B
variantHelpers:
  - core#find#nthorempty
similarHelpers:
  - core#find#nthor
  - core#find#nth
  - core#find#findorelse
  - core#find#firstorempty
position: 340
signatures:
  - "func NthOrEmpty[T any, N constraints.Integer](collection []T, nth N) T"
---

Returns the element at index nth of collection, or the zero value if out of bounds. If nth is negative, returns the nth element from the end.

```go
v := lo.NthOrEmpty([]int{10, 20, 30}, 10)
// v == 0
```


