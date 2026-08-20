---
name: TakeWhile
slug: takewhile
sourceRef: it/seq.go#L732
playUrl: "https://go.dev/play/p/gs5wsl2R3h6"
category: iter
subCategory: sequence
signatures:
  - "func TakeWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I"
variantHelpers:
  - iter#sequence#takewhile
similarHelpers:
  - core#slice#takewhile
position: 120
---

Takes elements from the beginning of a sequence while the predicate returns true.

```go
seq := func(yield func(int) bool) {
    for _, v := range []int{1, 2, 3, 4} {
        if !yield(v) {
            return
        }
    }
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
