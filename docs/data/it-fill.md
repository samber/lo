---
name: Fill
slug: fill
sourceRef: it/seq.go#L26
category: it
subCategory: sequence
signatures:
  - "func Fill[T lo.Clonable[T], I ~func(func(T) bool)](collection I, initial T) I"
variantHelpers: []
similarHelpers:
  - core#slice#fill
position: 174
---

Fill replaces elements of a sequence with `initial` value.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
}

filled := it.Fill(collection, 99)
var result []int
for item := range filled {
    result = append(result, item)
}
// result contains [99, 99, 99]
```