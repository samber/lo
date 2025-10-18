---
name: PartitionBy
slug: partitionby
sourceRef: slice.go#L240
category: core
subCategory: slice
playUrl: https://go.dev/play/p/NfQ_nGjkgXW
variantHelpers:
  - core#slice#partitionby
similarHelpers:
  - core#slice#groupby
  - core#slice#groupbymap
  - core#slice#chunk
  - core#map#keyby
position: 150
signatures:
  - "func PartitionBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K) []Slice"
---

Partitions a slice into groups determined by a key computed from each element, preserving original order.


