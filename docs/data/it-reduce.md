---
name: Reduce
slug: reduce
sourceRef: it/seq.go#L133
category: it
subCategory: sequence
signatures:
  - "func Reduce[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R"
  - "func ReduceI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R"
playUrl: ""
variantHelpers:
  - it#sequence#reduce
  - it#sequence#reducei
similarHelpers:
  - core#slice#reduce
  - core#slice#reducei
  - core#slice#reduceright
position: 30
---

Reduces a collection to a single accumulated value by applying an accumulator function to each element starting with an initial value.

### Reduce

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
}
sum := it.Reduce(seq, func(acc int, item int) int {
    return acc + item
}, 0)
// sum == 10
```

```go
seq := func(yield func(string) bool) {
    _ = yield("hello")
    _ = yield("world")
}
concat := it.Reduce(seq, func(acc string, item string) string {
    return acc + " " + item
}, "")
// concat == " hello world"
```

### ReduceI

Reduces a collection to a single value by iterating through elements and applying an accumulator function that includes the index.

```go
result := it.ReduceI(it.Range(1, 5), func(agg int, item int, index int) int {
    return agg + item*index
}, 0)
// 20 (0*0 + 1*1 + 2*2 + 3*3)
```