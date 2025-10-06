---
name: TrimPrefix
slug: trimprefix
sourceRef: it/seq.go#L778
category: it
subCategory: string
signatures:
  - "func TrimPrefix[T comparable, I ~func(func(T) bool)](collection I, prefix []T) I"
variantHelpers: []
similarHelpers:
  - core#string#trimprefix
position: 265
---

TrimPrefix removes all the leading prefix from the collection.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(1)
    yield(2)
    yield(3)
}

trimmed := it.TrimPrefix(collection, []int{1, 2})
var result []int
for item := range trimmed {
    result = append(result, item)
}
// result contains [1, 2, 3]
```