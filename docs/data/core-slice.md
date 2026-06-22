---
name: Slice
slug: slice
sourceRef: slice.go#L658
category: core
subCategory: slice
playUrl: https://go.dev/play/p/8XWYhfMMA1h
variantHelpers:
  - core#slice#slice
similarHelpers:
  - core#slice#subset
  - core#slice#drop
  - core#slice#dropright
  - core#slice#splice
  - core#slice#replace
position: 0
signatures:
  - "func Slice[T any, Slice ~[]T](collection Slice, start int, end int) Slice"
---

Returns a copy of a slice from `start` up to, but not including, `end`. Like `slice[start:end]`, but does not panic on overflow.

```go
in := []int{0, 1, 2, 3, 4}
lo.Slice(in, 2, 6)
// []int{2, 3, 4}
```


