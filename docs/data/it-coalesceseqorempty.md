---
name: CoalesceSeqOrEmpty
slug: coalesceseqorempty
sourceRef: it/type_manipulation.go#L86
category: iter
subCategory: type
signatures:
  - "func CoalesceSeqOrEmpty[T any](v ...iter.Seq[T]) iter.Seq[T]"
variantHelpers:
  - iter#type#coalesceseqorempty
similarHelpers:
  - iter#type#coalesceseq
  - core#type#coalesceorempty
  - core#type#coalescesliceorempty
position: 102
---

Returns the first non-empty sequence from the provided arguments, or an empty sequence if all arguments are empty.

```go
emptySeq := it.Empty[int]()
nonEmptySeq := it.Range(3)
result := it.CoalesceSeqOrEmpty(emptySeq, nonEmptySeq, emptySeq)
fmt.Println(slices.Collect(result))
// [0 1 2]

strSeq := slices.Values([]string{"a", "b"})
strResult := it.CoalesceSeqOrEmpty(it.Empty[string](), strSeq)
fmt.Println(slices.Collect(strResult))
// [a b]

result = it.CoalesceSeqOrEmpty(emptySeq, emptySeq)
fmt.Println(slices.Collect(result))
// []
```