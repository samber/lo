---
name: CoalesceSeqOrEmpty
slug: coalesceseqorempty
sourceRef: it/type_manipulation.go#L76
category: it
subCategory: type
signatures:
  - "func CoalesceSeqOrEmpty[T any](v ...iter.Seq[T]) iter.Seq[T]"
variantHelpers:
  - it#type#coalesceseqorempty
similarHelpers:
  - it#type#coalesceseq
  - core#type#coalesceorempty
  - core#type#coalescesliceorempty
position: 102
---

Returns the first non-empty sequence from the provided arguments, or an empty sequence if all arguments are empty.

```go
emptySeq := func(yield func(int) bool) bool {
    return false // empty sequence
}
nonEmptySeq := it.Range(3)
result := it.CoalesceSeqOrEmpty(emptySeq, nonEmptySeq, emptySeq)
// iter.Seq[int] yielding 0, 1, 2

emptyStrSeq := func(yield func(string) bool) bool {
    return false // empty sequence
}
strSeq := func(yield func(string) bool) bool {
    yield("a")
    yield("b")
    return true
}
result = it.CoalesceSeqOrEmpty(emptyStrSeq, strSeq)
// iter.Seq[string] yielding "a", "b"

result = it.CoalesceSeqOrEmpty(emptySeq, emptyStrSeq)
// empty sequence (yields nothing)
```