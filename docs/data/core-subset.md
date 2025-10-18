---
name: Subset
slug: subset
sourceRef: slice.go#L635
category: core
subCategory: slice
playUrl: https://go.dev/play/p/tOQu1GhFcog
variantHelpers:
  - core#slice#subset
similarHelpers:
  - core#slice#slice
  - core#slice#chunk
  - core#slice#drop
  - core#slice#dropright
position: 0
signatures:
  - "func Subset[T any, Slice ~[]T](collection Slice, offset int, length uint) Slice"
---

Returns a copy of a slice from `offset` up to `length` elements. Like `slice[start:start+length]`, but does not panic on overflow.

```go
in := []int{0, 1, 2, 3, 4}
lo.Subset(in, 2, 3)
// []int{2, 3, 4}
```


