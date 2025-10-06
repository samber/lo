---
name: Interleave
slug: interleave
sourceRef: it/seq.go#L26
category: it
subCategory: sequence
signatures:
  - "func Interleave[T any](collections ...iter.Seq[T]) iter.Seq[T]"
variantHelpers: []
similarHelpers:
  - core#slice#interleave
position: 173
---

Interleave round-robin alternating input sequences and sequentially appending value at index into result.

```go
seq1 := func(yield func(int) bool) {
    yield(1)
    yield(3)
}
seq2 := func(yield func(int) bool) {
    yield(2)
    yield(4)
}
seq3 := func(yield func(int) bool) {
    yield(5)
    yield(6)
}

interleaved := it.Interleave(seq1, seq2, seq3)
var result []int
for item := range interleaved {
    result = append(result, item)
}
// result contains [1, 2, 5, 3, 4, 6]
```