---
name: ToSeqPtr
slug: toseqptr
sourceRef: it/type_manipulation.go#L11
category: it
subCategory: type
signatures:
  - "func ToSeqPtr[T any](collection iter.Seq[T]) iter.Seq[*T]"
variantHelpers: []
similarHelpers:
  - core#type#toptr
position: 240
---

ToSeqPtr returns a sequence of pointers to each value.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
}

ptrs := it.ToSeqPtr(collection)
var result []*int
for ptr := range ptrs {
    result = append(result, ptr)
}
// result contains pointers to 1, 2, 3
```