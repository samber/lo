---
name: MinIndex
slug: minindex
sourceRef: it/find.go#L231
category: it
subCategory: find
signatures:
  - "func MinIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int)"
playUrl: "https://go.dev/play/p/4BvOSo-yza"
variantHelpers:
  - it#find#minindex
similarHelpers:
  - core#slice#minindex
position: 460
---

Searches the minimum value of a collection and returns both the value and its index.

Returns (zero value, -1) when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
// Find the minimum value and its index
numbers := it.Slice([]int{5, 2, 8, 1, 9})
value, index := it.MinIndex(numbers)
// value: 1, index: 3

// With empty collection
empty := it.Slice([]int{})
value, index := it.MinIndex(empty)
// value: 0, index: -1
```