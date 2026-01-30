---
name: Sliding
slug: sliding
sourceRef: it/seq.go#L329
category: it
subCategory: sequence
signatures:
  - "func Sliding[T any](collection iter.Seq[T], size, step int) iter.Seq[[]T]"
variantHelpers:
  - it#sequence#sliding
similarHelpers:
  - core#slice#sliding
position: 80
---

Creates a sequence of sliding windows of a given size with a given step. If step equals size, windows don't overlap.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
    yield(6)
    yield(7)
    yield(8)
}
windows := it.Sliding(seq, 2, 3)
var result [][]int
for w := range windows {
    result = append(result, w)
}
// result contains [1 2], [4 5], [7 8]
```
