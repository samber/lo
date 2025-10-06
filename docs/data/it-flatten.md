---
name: Flatten
slug: flatten
sourceRef: it/seq.go#L26
category: it
subCategory: sequence
signatures:
  - "func Flatten[T any, I ~func(func(T) bool)](collection []I) I"
variantHelpers: []
similarHelpers:
  - core#slice#flatten
position: 172
---

Flatten returns a sequence a single level deep.

```go
seq1 := func(yield func(int) bool) {
    yield(1)
    yield(2)
}
seq2 := func(yield func(int) bool) {
    yield(3)
    yield(4)
}

flattened := it.Flatten([]iter.Seq[int]{seq1, seq2})
var result []int
for item := range flattened {
    result = append(result, item)
}
// result contains [1, 2, 3, 4]
```