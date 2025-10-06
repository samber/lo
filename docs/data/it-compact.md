---
name: Compact
slug: compact
sourceRef: it/seq.go#L699
category: it
subCategory: slice
signatures:
  - "func Compact[T comparable, I ~func(func(T) bool)](collection I) I"
variantHelpers:
  - it#slice#compact
similarHelpers:
  - core#slice#compact
position: 192
---

Compact returns a sequence of all non-zero elements.

```go
collection := func(yield func(int) bool) {
    yield(0)
    yield(1)
    yield(0)
    yield(2)
    yield(3)
    yield(0)
}

compacted := it.Compact(collection)
var result []int
for item := range compacted {
    result = append(result, item)
}
// result contains [1, 2, 3]
```