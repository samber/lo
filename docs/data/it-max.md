---
name: Max
slug: max
sourceRef: it/find.go#L295
category: it
subCategory: find
signatures:
  - "func Max[T constraints.Ordered](collection iter.Seq[T]) T"
playUrl: ""
variantHelpers:
  - it#find#max
similarHelpers:
  - core#slice#max
  - it#find#min
  - it#find#maxby
position: 110
---

Searches the maximum value of a collection. Returns the largest element found.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(5)
    _ = yield(2)
    _ = yield(8)
    _ = yield(1)
    _ = yield(9)
}
max := it.Max(seq)
// max == 9
```

```go
seq := func(yield func(string) bool) {
    _ = yield("zebra")
    _ = yield("apple")
    _ = yield("banana")
}
max := it.Max(seq)
// max == "zebra" (lexicographically largest)
```