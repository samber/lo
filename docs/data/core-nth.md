---
name: Nth
slug: nth
sourceRef: find.go#L617
category: core
subCategory: find
playUrl: https://go.dev/play/p/sHoh88KWt6B
variantHelpers:
  - core#find#nth
similarHelpers:
  - core#find#first
  - core#find#last
  - core#find#nthor
  - core#find#nthorempty
  - core#slice#indexof
  - core#slice#drop
position: 320
signatures:
  - "func Nth[T any, N constraints.Integer](collection []T, nth N) (T, error)"
---

Returns the element at index nth of collection. If nth is negative, returns the nth element from the end. Returns an error when nth is out of slice bounds.

```go
v, _ := lo.Nth([]int{10, 20, 30}, 1)
// v == 20
```


