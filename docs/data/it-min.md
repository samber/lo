---
name: Min
slug: min
sourceRef: it/find.go#L227
category: it
subCategory: find
signatures:
  - "func Min[T constraints.Ordered](collection iter.Seq[T]) T"
playUrl: "https://go.dev/play/p/3AuTNRn-yz"
variantHelpers:
  - it#find#min
similarHelpers:
  - core#slice#min
  - it#find#max
  - it#find#minby
position: 100
---

Searches the minimum value of a collection. Returns the smallest element found.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(5)
    _ = yield(2)
    _ = yield(8)
    _ = yield(1)
    _ = yield(9)
}
min := it.Min(seq)
// min == 1
```

```go
seq := func(yield func(string) bool) {
    _ = yield("zebra")
    _ = yield("apple")
    _ = yield("banana")
}
min := it.Min(seq)
// min == "apple" (lexicographically smallest)
```