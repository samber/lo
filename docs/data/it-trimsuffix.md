---
name: TrimSuffix
slug: trimsuffix
sourceRef: it/seq.go#L778
category: it
subCategory: string
signatures:
  - "func TrimSuffix[T comparable, I ~func(func(T) bool)](collection I, suffix []T) I"
variantHelpers: []
similarHelpers:
  - core#string#trimsuffix
position: 266
---

TrimSuffix removes all the trailing suffix from the collection.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(3)
    yield(4)
}

trimmed := it.TrimSuffix(collection, []int{3, 4})
var result []int
for item := range trimmed {
    result = append(result, item)
}
// result contains [1, 2]
```