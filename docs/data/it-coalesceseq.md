---
name: CoalesceSeq
slug: coalesceseq
sourceRef: it/type_manipulation.go#L74
playUrl: "https://go.dev/play/p/krZ-laaVi2C"
category: iter
subCategory: type
signatures:
  - "func CoalesceSeq[T any](v ...iter.Seq[T]) (iter.Seq[T], bool)"
variantHelpers:
  - iter#type#coalesceseq
similarHelpers:
  - iter#type#coalesceseqorempty
  - core#type#coalesce
  - core#type#coalesceslice
position: 100
---

Returns the first non-empty sequence from the provided arguments, with a boolean indicating if a non-empty sequence was found.

```go
emptySeq := func(yield func(int) bool) {
    // empty sequence
}
nonEmptySeq := it.Range(3)
result, ok := it.CoalesceSeq(emptySeq, nonEmptySeq, emptySeq)
// iter.Seq[int] yielding 0, 1, 2, true

emptyStrSeq := func(yield func(string) bool) {
    // empty sequence
}
strSeq := func(yield func(string) bool) {
    yield("a")
    yield("b")
}
strResult, ok := it.CoalesceSeq(emptyStrSeq, strSeq)
// iter.Seq[string] yielding "a", "b", true

noResult, ok := it.CoalesceSeq(emptySeq, emptySeq)
// nil sequence, false
```