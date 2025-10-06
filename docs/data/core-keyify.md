---
name: Keyify
slug: keyify
sourceRef: slice.go#L429
category: core
subCategory: slice
playUrl: https://go.dev/play/p/RYhhM_csqIG
variantHelpers:
  - core#slice#keyify
similarHelpers:
  - core#slice#keyby
  - core#slice#uniq
  - core#slice#uniqby
  - core#slice#groupby
position: 270
signatures:
  - "func Keyify[T comparable, Slice ~[]T](collection Slice) map[T]struct{}"
---

Returns a set-like map where each unique element of the slice is a key.

```go
set := lo.Keyify([]int{1, 1, 2, 3, 4})
// map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}
```


