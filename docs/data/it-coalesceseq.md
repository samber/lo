---
name: CoalesceSeq
slug: coalesceseq
sourceRef: it/type_manipulation.go#L65
category: it
subCategory: type
signatures:
  - "func CoalesceSeq[T any](v ...iter.Seq[T]) (iter.Seq[T], bool)"
variantHelpers:
  - it#type#coalesceseq
similarHelpers:
  - it#type#coalesceseqorempty
  - core#type#coalesce
  - core#type#coalesceslice
position: 100
---

Returns the first non-empty sequence from the provided arguments, with a boolean indicating if a non-empty sequence was found.

```go
emptySeq := func(yield func(int) bool) bool {
    return false // empty sequence
}
nonEmptySeq := it.Range(3)
result, ok := it.CoalesceSeq(emptySeq, nonEmptySeq, emptySeq)
// iter.Seq[int] yielding 0, 1, 2, true

emptyStrSeq := func(yield func(string) bool) bool {
    return false // empty sequence
}
strSeq := func(yield func(string) bool) bool {
    yield("a")
    yield("b")
    return true
}
result, ok = it.CoalesceSeq(emptyStrSeq, strSeq)
// iter.Seq[string] yielding "a", "b", true

result, ok = it.CoalesceSeq(emptySeq, emptyStrSeq)
// nil sequence, false
```