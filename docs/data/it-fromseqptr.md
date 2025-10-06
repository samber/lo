---
name: FromSeqPtr
slug: fromseqptr
sourceRef: it/type_manipulation.go#L11
category: it
subCategory: type
signatures:
  - "func FromSeqPtr[T any](collection iter.Seq[*T]) iter.Seq[T]"
variantHelpers:
  - it#type#fromseqptror
similarHelpers:
  - core#type#fromptr
position: 241
---

FromSeqPtr returns a sequence with the pointer values.
Returns a zero value in case of a nil pointer element.

```go
one := 1
two := 2
var three *int = nil

collection := func(yield func(*int) bool) {
    yield(&one)
    yield(&two)
    yield(three)
}

values := it.FromSeqPtr(collection)
var result []int
for val := range values {
    result = append(result, val)
}
// result contains [1, 2, 0]
```