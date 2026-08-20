---
name: Take
slug: take
sourceRef: it/seq.go#L711
playUrl: "https://go.dev/play/p/VTmRfEQQjij"
category: iter
subCategory: sequence
signatures:
  - "func Take[T any, I ~func(func(T) bool)](collection I, n int) I"
variantHelpers:
  - iter#sequence#take
similarHelpers:
  - core#slice#take
position: 110
---

Takes the first n elements from a sequence.

```go
seq := func(yield func(int) bool) {
    for _, v := range []int{1, 2, 3, 4} {
        if !yield(v) {
            return
        }
    }
}
result := it.Take(seq, 2)
var out []int
for v := range result {
    out = append(out, v)
}
// out contains [1, 2]
```
