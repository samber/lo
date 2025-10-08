---
name: Drop
slug: drop
sourceRef: it/seq.go#L498
category: it
subCategory: sequence
signatures:
  - "func Drop[T any, I ~func(func(T) bool)](collection I, n int) I"
playUrl: "https://go.dev/play/p/1SmFJ5-zr"
variantHelpers:
  - it#sequence#drop
similarHelpers:
  - core#slice#drop
  - it#sequence#droplast
position: 100
---

Drops n elements from the beginning of a sequence.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
}
dropped := it.Drop(seq, 2)
var result []int
for v := range dropped {
    result = append(result, v)
}
// result contains 3, 4, 5
```