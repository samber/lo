---
name: CountValues
slug: countvalues
sourceRef: it/seq.go#L720
category: iter
subCategory: slice
signatures:
  - "func CountValues[T comparable](collection iter.Seq[T]) map[T]int"
playUrl: https://go.dev/play/p/PPBT4Fp-V3B
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