---
name: TakeFilter
slug: takefilter
sourceRef: it/seq.go#L729
category: it
subCategory: sequence
signatures:
  - "func TakeFilter[T any, I ~func(func(T) bool)](collection I, n int, predicate func(item T, index int) bool) I"
variantHelpers:
  - it#sequence#takefilter
similarHelpers:
  - core#slice#takefilter
position: 130
---

Filters elements and takes the first n matches. Stops once n matches are found.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
    yield(6)
}
result := it.TakeFilter(seq, 2, func(x, _ int) bool {
    return x%2 == 0
})
var out []int
for v := range result {
    out = append(out, v)
}
// out contains [2, 4]
```
