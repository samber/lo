---
name: Shuffle
slug: shuffle
sourceRef: slice.go#L322
category: core
subCategory: slice
playUrl: https://go.dev/play/p/ZTGG7OUCdnp
variantHelpers:
  - core#slice#shuffle
similarHelpers:
  - mutable#slice#shuffle
  - core#slice#sample
  - core#slice#samples
  - core#slice#sampleby
  - core#slice#samplesby
position: 180
signatures:
  - "func Shuffle[T any, Slice ~[]T](collection Slice) Slice"
---

Returns a slice of shuffled values (Fisher–Yates). Deprecated: use `mutable.Shuffle`.


