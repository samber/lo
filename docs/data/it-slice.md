---
name: Slice
slug: slice
sourceRef: it/seq.go#L680
category: it
subCategory: sequence
signatures:
  - "func Slice[T any, I ~func(func(T) bool)](collection I, start, end int) I"
playUrl: 
variantHelpers:
  - it#slice#drop
similarHelpers:
  - core#slice#slice
  - core#slice#chunk
position: 80
---

Returns a sub-sequence from start index to end index (exclusive).

```go
result := it.Slice(it.Range(1, 10), 2, 5)
// [3, 4, 5]
```