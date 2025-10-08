---
name: IndexOf
slug: indexof
sourceRef: it/find.go#L19
category: it
subCategory: find
signatures:
  - "func IndexOf[T comparable](collection iter.Seq[T], element T) int"
playUrl: https://go.dev/play/p/1OZHU2yfb-m
variantHelpers:
  - it#find#indexof
similarHelpers:
  - core#slice#indexof
  - it#find#lastindexof
position: 0
---

Returns the index at which the first occurrence of a value is found in the sequence, or -1 if the value is not found.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
    _ = yield(20)
}
idx := it.IndexOf(seq, 20)
// idx == 1
```

```go
seq := func(yield func(string) bool) {
    _ = yield("apple")
    _ = yield("banana")
    _ = yield("cherry")
}
idx := it.IndexOf(seq, "orange")
// idx == -1
```