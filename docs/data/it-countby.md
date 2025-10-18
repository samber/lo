---
name: CountBy
slug: countby
sourceRef: it/seq.go#L325
category: it
subCategory: find
signatures:
  - "func CountBy[T any](collection iter.Seq[T], predicate func(item T) bool) int"
playUrl:
variantHelpers:
  - it#find#count
similarHelpers:
  - core#slice#countby
  - core#slice#count
position: 35
---

Counts the number of elements in the collection that satisfy the predicate.

```go
result := it.CountBy(it.Range(1, 11), func(item int) bool {
    return item%2 == 0
})
// 5
```