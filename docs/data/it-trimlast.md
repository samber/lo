---
name: TrimLast
slug: trimlast
sourceRef: it/seq.go#L778
category: it
subCategory: string
signatures:
  - "func TrimLast[T comparable, I ~func(func(T) bool)](collection I, cutset ...T) I"
variantHelpers:
  - it#string#trim
  - it#string#trimfirst
similarHelpers: []
position: 264
---

TrimLast removes all the trailing cutset from the collection.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(0)
    yield(0)
}

trimmed := it.TrimLast(collection, 0)
var result []int
for item := range trimmed {
    result = append(result, item)
}
// result contains [1, 2, 3]
```