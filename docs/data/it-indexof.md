---
name: IndexOf
slug: indexof
sourceRef: it/find.go#L19
category: it
subCategory: find
signatures:
  - "func IndexOf[T comparable](collection iter.Seq[T], element T) int"
playUrl: ""
variantHelpers:
  - it#find#indexof
similarHelpers:
  - core#slice#indexof
  - it#find#lastindexof
position: 0
---

Returns the index at which the first occurrence of a value is found in the sequence, or -1 if the value is not found. Scans the sequence from the beginning and returns the position of the first matching element.

```go
// Find existing element - returns first occurrence
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
    _ = yield(20)
}
idx := it.IndexOf(seq, 20)
// idx: 1 (first occurrence of 20)
```

```go
// Element not found - returns -1
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
idx := it.IndexOf(seq, "orange")
// idx: -1 (orange not found in sequence)
```

```go
// Empty sequence - returns -1
emptySeq := func(yield func(string) bool) {
    // no elements yielded
}
idx := it.IndexOf(emptySeq, "anything")
// idx: -1 (sequence is empty)
```