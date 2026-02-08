---
name: TakeWhile
slug: takewhile
sourceRef: it/seq.go#L706
category: it
subCategory: sequence
signatures:
  - "func TakeWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I"
variantHelpers:
  - it#sequence#takewhile
similarHelpers:
  - core#slice#takewhile
position: 120
---

Takes elements from the beginning of a sequence while the predicate returns true.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}
result := it.TakeWhile(seq, func(x int) bool {
    return x < 3
})
var out []int
for v := range result {
    out = append(out, v)
}
// out contains [1, 2]
```
