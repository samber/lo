---
name: CountValues
slug: countvalues
sourceRef: it/seq.go#L720
category: it
subCategory: slice
signatures:
  - "func CountValues[T comparable](collection iter.Seq[T]) map[T]int"
variantHelpers: []
similarHelpers:
  - core#slice#countvalues
position: 203
---

CountValues counts the number of each element in the collection.

```go
collection := func(yield func(string) bool) {
    yield("apple")
    yield("banana")
    yield("apple")
    yield("cherry")
    yield("banana")
    yield("apple")
}

counts := it.CountValues(collection)
// counts contains {"apple": 3, "banana": 2, "cherry": 1}
```