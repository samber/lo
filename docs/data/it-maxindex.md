---
name: MaxIndex
slug: maxindex
sourceRef: it/find.go#L299
category: it
subCategory: find
signatures:
  - "func MaxIndex[T constraints.Ordered](collection iter.Seq[T]) (T, int)"
playUrl: "https://go.dev/play/p/0HbUY4-zg"
variantHelpers:
  - it#find#maxindex
similarHelpers:
  - core#slice#maxindex
position: 480
---

Searches the maximum value of a collection and returns both the value and its index.

Returns (zero value, -1) when the collection is empty.
Will iterate through the entire sequence.

Examples:

```go
// Find the maximum value and its index
numbers := it.Slice([]int{5, 2, 8, 1, 9})
value, index := it.MaxIndex(numbers)
// value: 9, index: 4

// With empty collection
empty := it.Slice([]int{})
value, index := it.MaxIndex(empty)
// value: 0, index: -1

// Find the maximum string alphabetically and its index
words := it.Slice([]string{"apple", "zebra", "banana", "xylophone"})
value, index := it.MaxIndex(words)
// value: "zebra", index: 1
```