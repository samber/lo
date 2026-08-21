---
name: Fill
slug: fill
sourceRef: it/seq.go#L492
category: iter
subCategory: sequence
signatures:
  - "func Fill[T lo.Clonable[T], I ~func(func(T) bool)](collection I, initial T) I"
variantHelpers: []
playUrl: https://go.dev/play/p/mHShWq5ezMc
similarHelpers:
  - core#slice#fill
position: 174
---

Fill replaces elements of a sequence with `initial` value.

```go
type item struct{ value int }

func (i item) Clone() item {
    return item{i.value}
}

collection := func(yield func(item) bool) {
    yield(item{1})
    yield(item{2})
    yield(item{3})
}

filled := it.Fill(collection, item{99})
var result []item
for v := range filled {
    result = append(result, v)
}
// result contains [{99} {99} {99}]
```