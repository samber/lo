---
name: Flatten
slug: flatten
sourceRef: slice.go#L266
category: core
subCategory: slice
playUrl: https://go.dev/play/p/rbp9ORaMpjw
variantHelpers:
  - core#slice#flatten
similarHelpers:
  - core#slice#chunk
  - core#slice#interleave
  - core#slice#slice
position: 160
signatures:
  - "func Flatten[T any, Slice ~[]T](collection []Slice) Slice"
---

Flattens a slice of slices by one level.

```go
flat := lo.Flatten([][]int{{0, 1}, {2, 3, 4, 5}})
// []int{0, 1, 2, 3, 4, 5}
```
