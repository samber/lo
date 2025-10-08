---
name: Reduce
slug: reduce
sourceRef: it/seq.go#L133
category: it
subCategory: sequence
signatures:
  - "func Reduce[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R"
  - "func ReduceLast[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R"
  - "func ReduceLastI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R"
playUrl: https://go.dev/play/p/FmkVUf39ZP_Y
variantHelpers:
  - it#sequence#reduce
  - it#sequence#reducei
  - it#sequence#reducelast
  - it#sequence#reducelasti
similarHelpers:
  - core#slice#reduce
  - core#slice#reduceright
position: 30
---

Reduces a collection to a single accumulated value by applying an accumulator function to each element starting with an initial value.

Examples:

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

ReduceLast is like Reduce except that it iterates over elements of collection in reverse.

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
}
sum := it.ReduceLast(seq, func(acc int, item int) int {
    return acc + item
}, 0)
// sum == 10 (4 + 3 + 2 + 1 + 0)

result := it.ReduceLast(seq, func(agg string, item int) string {
    return fmt.Sprintf("%d-%s", item, agg)
}, "end")
// result == "4-3-2-1-end"
```

ReduceLastI is like Reduce except that it iterates over elements of collection in reverse, with index.

```go
seq := func(yield func(string) bool) {
    _ = yield("a")
    _ = yield("b")
    _ = yield("c")
}
result := it.ReduceLastI(seq, func(agg string, item string, index int) string {
    return fmt.Sprintf("%s:%d:%s", agg, index, item)
}, "start")
// result == "start:2:c:1:b:0:a"
```