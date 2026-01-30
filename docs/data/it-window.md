---
name: Window
slug: window
sourceRef: it/seq.go#L318
category: it
subCategory: sequence
signatures:
  - "func Window[T any](collection iter.Seq[T], size int) iter.Seq[[]T]"
variantHelpers:
  - it#sequence#window
similarHelpers:
  - core#slice#window
position: 70
---

Creates a sequence of sliding windows of a given size. Each window overlaps with the previous one by size-1 elements.

```go
seq := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}
windows := it.Window(seq, 3)
var result [][]int
for w := range windows {
    result = append(result, w)
}
// result contains [1 2 3], [2 3 4], [3 4 5]
```
