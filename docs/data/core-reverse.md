---
name: Reverse
slug: reverse
sourceRef: slice.go#L331
category: core
subCategory: slice
playUrl: https://go.dev/play/p/iv2e9jslfBM
variantHelpers:
  - core#slice#reverse
similarHelpers:
  - mutable#slice#reverse
  - core#slice#drop
  - core#slice#dropright
  - core#slice#slice
  - core#slice#flatten
position: 190
signatures:
  - "func Reverse[T any, Slice ~[]T](collection Slice) Slice"
---

Reverses a slice in place. Deprecated: use `mutable.Reverse`.


