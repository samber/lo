---
name: Replace
slug: replace
sourceRef: slice.go#L684
category: core
subCategory: slice
playUrl: https://go.dev/play/p/XfPzmf9gql6
variantHelpers:
  - core#slice#replace
similarHelpers:
  - core#string#replace
  - core#string#replaceall
  - core#slice#fill
  - core#slice#splice
  - core#slice#slice
position: 0
signatures:
  - "func Replace[T comparable, Slice ~[]T](collection Slice, old T, nEw T, n int) Slice"
---

Returns a copy of the slice with the first n non-overlapping instances of old replaced by new.

```go
in := []int{0, 1, 0, 1, 2, 3, 0}
lo.Replace(in, 0, 42, 2)
// []int{42, 1, 42, 1, 2, 3, 0}
```


