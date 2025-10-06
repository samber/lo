---
name: FromEntries
slug: fromentries
sourceRef: it/map.go#L96
category: it
subCategory: map
signatures:
  - "func FromEntries[K comparable, V any](entries ...iter.Seq2[K, V]) map[K]V"
playUrl: ""
variantHelpers:
  - it#map#fromentries
  - it#map#frompairs
similarHelpers:
  - core#slice#fromentries
  - it#map#entries
position: 30
---

Transforms a sequence of key/value pairs into a map. Accepts multiple sequences and merges them.

Examples:

```go
entries := func(yield func(string, int) bool) {
    _ = yield("apple", 1)
    _ = yield("banana", 2)
    _ = yield("cherry", 3)
}
m := it.FromEntries(entries)
// m == map[string]int{"apple": 1, "banana": 2, "cherry": 3}
```

```go
entries1 := func(yield func(string, int) bool) {
    _ = yield("a", 1)
    _ = yield("b", 2)
}
entries2 := func(yield func(string, int) bool) {
    _ = yield("c", 3)
    _ = yield("d", 4)
}
m := it.FromEntries(entries1, entries2)
// m contains all entries from both sequences
```