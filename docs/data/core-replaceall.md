---
name: ReplaceAll
slug: replaceall
sourceRef: slice.go#L700
category: core
subCategory: slice
playUrl: https://go.dev/play/p/a9xZFUHfYcV
variantHelpers:
  - core#slice#replaceall
similarHelpers:
  - core#slice#replace
  - core#slice#map
  - core#slice#filtermap
position: 0
signatures:
  - "func ReplaceAll[T comparable, Slice ~[]T](collection Slice, old T, nEw T) Slice"
---

Returns a copy of the slice with all non-overlapping instances of old replaced by new.

```go
in := []int{0, 1, 0, 1, 2, 3, 0}
lo.ReplaceAll(in, 0, 42)
// []int{42, 1, 42, 1, 2, 3, 42}
```


