---
name: ToAnySeq
slug: toanyseq
sourceRef: it/type_manipulation.go#L11
category: iter
subCategory: type
signatures:
  - "func ToAnySeq[T any](collection iter.Seq[T]) iter.Seq[any]"
variantHelpers:
  - iter#type#fromanyseq
playUrl: "https://go.dev/play/p/ktE4IMXDMxv"
similarHelpers:
  - core#type#toany
position: 243
---

ToAnySeq returns a sequence with all elements mapped to `any` type.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
}

anySeq := it.ToAnySeq(collection)
var result []any
for item := range anySeq {
    result = append(result, item)
}
// result contains [1, 2, 3] as any type
```