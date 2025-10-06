---
name: IsSortedBy
slug: issortedby
sourceRef: it/seq.go#L720
category: it
subCategory: slice
signatures:
  - "func IsSortedBy[T any, K constraints.Ordered](collection iter.Seq[T], transform func(item T) K) bool"
variantHelpers: []
similarHelpers:
  - core#slice#issortedby
position: 201
---

IsSortedBy checks if a sequence is sorted by transform.

```go
collection := func(yield func(string) bool) {
    yield("apple")
    yield("banana")
    yield("cherry")
}

sortedByLength := it.IsSortedBy(collection, func(s string) int {
    return len(s)
})
// true (5, 6, 6 is sorted)
```