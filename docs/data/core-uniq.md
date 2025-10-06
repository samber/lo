---
name: Uniq
slug: uniq
sourceRef: slice.go#L140
category: core
subCategory: slice
playUrl: https://go.dev/play/p/DTzbeXZ6iEN
variantHelpers:
  - core#slice#uniq
similarHelpers:
  - core#slice#uniqby
  - core#slice#uniqmap
  - core#slice#uniqkeys
  - core#slice#uniqvalues
position: 100
signatures:
  - "func Uniq[T comparable, Slice ~[]T](collection Slice) Slice"
---

Returns a duplicate-free version of a slice, keeping only the first occurrence of each value. Order is preserved.

```go
lo.Uniq([]int{1, 2, 2, 1})
// []int{1, 2}
```


