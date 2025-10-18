---
name: Trim
slug: trim
sourceRef: it/seq.go#L778
category: it
subCategory: string
signatures:
  - "func Trim[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I"
variantHelpers:
  - it#string#trimfirst
  - it#string#trimlast
similarHelpers:
  - core#string#trim
position: 262
---

Trim removes all the leading and trailing cutset from the collection.

```go
collection := func(yield func(int) bool) {
    yield(0)
    yield(0)
    yield(1)
    yield(2)
    yield(3)
    yield(0)
    yield(0)
}

trimmed := it.Trim(collection, 0)
var result []int
for item := range trimmed {
    result = append(result, item)
}
// result contains [1, 2, 3]
```