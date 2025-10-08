---
name: ReduceLast
slug: reducelast
sourceRef: it/seq.go#L153
category: it
subCategory: sequence
signatures:
  - "func ReduceLast[T, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R"
  - "func ReduceLastI[T, R any](collection iter.Seq[T], accumulator func(agg R, item T, index int) R, initial R) R"
playUrl:
variantHelpers:
  - it#sequence#reduce
  - it#sequence#reducei
  - it#sequence#reducelast
  - it#sequence#reducelasti
similarHelpers:
  - core#slice#reducelast
  - core#slice#reducelasti
  - core#slice#reduce
position: 54
---

Reduces a collection from right to left, returning a single value.

### ReduceLast

```go
result := it.ReduceLast(it.Range(1, 5), func(agg int, item int) int {
    return agg - item
}, 0)
// -10 (0 - 4 - 3 - 2 - 1)
```

### ReduceLastI

Reduces a collection from right to left, returning a single value. The accumulator function includes the index.

```go
result := it.ReduceLastI(it.Range(1, 5), func(agg int, item int, index int) int {
    return agg - item*index
}, 0)
// -20 (0 - 4*3 - 3*2 - 2*1 - 1*0)
```