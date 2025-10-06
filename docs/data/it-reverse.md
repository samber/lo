---
name: Reverse
slug: reverse
sourceRef: it/seq.go#L366
category: it
subCategory: sequence
signatures:
  - "func Reverse[T any, I ~func(func(T) bool)](collection I) I"
playUrl: ""
variantHelpers:
  - it#sequence#reverse
similarHelpers:
  - core#slice#reverse
  - it#sequence#shuffle
position: 90
---

Reverses a sequence so the first element becomes the last and the last element becomes the first.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
}
reversed := it.Reverse(seq)
var result []int
for v := range reversed {
    result = append(result, v)
}
// result contains 4, 3, 2, 1
```