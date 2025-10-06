---
name: LastIndexOf
slug: lastindexof
sourceRef: it/find.go#L34
category: it
subCategory: find
signatures:
  - "func LastIndexOf[T comparable](collection iter.Seq[T], element T) int"
playUrl: ""
variantHelpers:
  - it#find#lastindexof
similarHelpers:
  - core#slice#lastindexof
  - it#find#indexof
position: 10
---

Returns the index at which the last occurrence of a value is found in the sequence, or -1 if the value is not found.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
    _ = yield(20)
}
idx := it.LastIndexOf(seq, 20)
// idx == 3
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
idx := it.LastIndexOf(seq, "orange")
// idx == -1
```