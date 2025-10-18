---
name: Replace
slug: replace
sourceRef: it/seq.go#L699
category: it
subCategory: slice
signatures:
  - "func Replace[T comparable, I ~func(func(T) bool)](collection I, old, nEw T, n int) I"
variantHelpers:
  - it#slice#replace
similarHelpers:
  - core#slice#replace
position: 190
---

Replace returns a sequence with the first n non-overlapping instances of old replaced by new.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(2)
    yield(3)
    yield(2)
    yield(4)
}

replaced := it.Replace(collection, 2, 99, 2)
var result []int
for item := range replaced {
    result = append(result, item)
}
// result contains [1, 99, 99, 3, 2, 4]
```