---
name: FromAnySeq
slug: fromanyseq
sourceRef: it/type_manipulation.go#L11
category: it
subCategory: type
signatures:
  - "func FromAnySeq[T any](collection iter.Seq[any]) iter.Seq[T]"
variantHelpers:
  - it#type#toanyseq
similarHelpers:
  - core#type#fromany
position: 244
---

FromAnySeq returns a sequence with all elements mapped to a type.
Panics on type conversion failure.

```go
collection := func(yield func(any) bool) {
    yield(1)
    yield(2)
    yield("three") // This will cause panic
}

intSeq := it.FromAnySeq[int](collection)
// This will panic when trying to convert "three" to int
```