---
name: Take
slug: take
sourceRef: it/seq.go#L682
category: it
subCategory: sequence
signatures:
  - "func Take[T any, I ~func(func(T) bool)](collection I, n int) I"
variantHelpers:
  - it#sequence#take
similarHelpers:
  - core#slice#take
position: 110
---

Takes the first n elements from a sequence.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}
result := it.Take(seq, 2)
var out []int
for v := range result {
    out = append(out, v)
}
// out contains [1, 2]
```
