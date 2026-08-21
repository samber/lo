---
name: Mode
slug: mode
sourceRef: it/math.go#L134
category: iter
subCategory: math
signatures:
  - "func Mode[T constraints.Integer | constraints.Float](collection iter.Seq[T]) []T"
playUrl: "https://go.dev/play/p/c_cmMMA5EhH"
variantHelpers:
  - iter#math#mode
similarHelpers:
  - core#math#mode
position: 40
---

Returns the mode (most frequent value) of a collection. If multiple values have the same highest frequency, then multiple values are returned. If the collection is empty, then the zero value of T is returned.

Will iterate through the entire sequence and allocate a map large enough to hold all distinct elements. Long heterogeneous input sequences can cause excessive memory usage.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(2)
    _ = yield(3)
    _ = yield(3)
    _ = yield(3)
}
mode := it.Mode(seq)
// mode == []int{3}
```

```go
// Multiple modes
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(1)
    _ = yield(2)
}
mode := it.Mode(seq)
// mode contains both 1 and 2 (order may vary)
```