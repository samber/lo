---
name: UniqBy
slug: uniqby
sourceRef: it/seq.go#L225
category: iter
subCategory: sequence
signatures:
  - "func UniqBy[T any, U comparable, I ~func(func(T) bool)](collection I, transform func(item T) U) I"
playUrl: https://go.dev/play/p/HKrt3AvwMTR
variantHelpers:
  - iter#slice#uniq
similarHelpers:
  - core#slice#uniqby
  - core#slice#uniq
position: 45
---

Returns a sequence with duplicate elements removed based on a transform function.

```go
result := it.UniqBy(it.Range(1, 7), func(item int) int {
    return item % 3
})
// [1, 2, 3]
```