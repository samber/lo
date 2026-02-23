---
name: TakeFilter
slug: takefilter
sourceRef: it/seq.go#L725
category: it
subCategory: sequence
signatures:
  - "func TakeFilter[T any, I ~func(func(T) bool)](collection I, n int, predicate func(item T) bool) I"
  - "func TakeFilterI[T any, I ~func(func(T) bool)](collection I, n int, predicate func(item T, index int) bool) I"
variantHelpers:
  - it#sequence#takefilter
  - it#sequence#takefilteri
similarHelpers:
  - core#slice#takefilter
position: 130
---

Filters elements and takes the first n matches. Stops once n matches are found.

### TakeFilter

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
    yield(6)
}
result := it.TakeFilter(seq, 2, func(x int) bool {
    return x%2 == 0
})
var out []int
for v := range result {
    out = append(out, v)
}
// out contains [2, 4]
```

### TakeFilterI

```go
result := it.TakeFilterI(seq, 2, func(x, index int) bool {
    return x%2 == 0 && index < 4
})
// out contains [2, 4]
```
