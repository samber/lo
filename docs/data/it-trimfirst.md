---
name: TrimFirst
slug: trimfirst
sourceRef: it/seq.go#L778
category: it
subCategory: string
signatures:
  - "func TrimFirst[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I"
variantHelpers:
  - it#string#trim
  - it#string#trimlast
similarHelpers: []
position: 263
---

TrimFirst removes all the leading cutset from the collection.

```go
collection := func(yield func(int) bool) {
    yield(0)
    yield(0)
    yield(1)
    yield(2)
    yield(3)
}

trimmed := it.TrimFirst(collection, 0)
var result []int
for item := range trimmed {
    result = append(result, item)
}
// result contains [1, 2, 3]
```