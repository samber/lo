---
name: CoalesceSeqOrEmpty
slug: coalesceseqorempty
sourceRef: it/type_manipulation.go#L86
playUrl: "https://go.dev/play/p/wep4z2KJLCO"
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
emptySeq := func(yield func(int) bool) {
    // empty sequence
}
nonEmptySeq := it.Range(3)
result := it.CoalesceSeqOrEmpty(emptySeq, nonEmptySeq, emptySeq)
// iter.Seq[int] yielding 0, 1, 2

emptyStrSeq := func(yield func(string) bool) {
    // empty sequence
}
strSeq := func(yield func(string) bool) {
    yield("a")
    yield("b")
}
strResult := it.CoalesceSeqOrEmpty(emptyStrSeq, strSeq)
// iter.Seq[string] yielding "a", "b"

noResult := it.CoalesceSeqOrEmpty(emptySeq, emptySeq)
// empty sequence (yields nothing)
```