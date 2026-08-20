---
name: Shuffle
slug: shuffle
sourceRef: slice.go#L573
category: core
subCategory: slice
playUrl: https://go.dev/play/p/whgrQWwOy-j
variantHelpers:
  - core#slice#shuffle
similarHelpers:
  - mutable#slice#shuffle
  - core#find#sample
  - core#find#samples
  - core#find#sampleby
  - core#find#samplesby
position: 180
signatures:
  - "func Shuffle[T any, Slice ~[]T](collection Slice) Slice"
---

Returns a slice of shuffled values (Fisher–Yates). Deprecated: use `mutable.Shuffle`.


