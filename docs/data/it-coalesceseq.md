---
name: CoalesceSeq
slug: coalesceseq
sourceRef: it/type_manipulation.go#L74
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
emptySeq := it.Empty[int]()
nonEmptySeq := it.Range(3)
result, ok := it.CoalesceSeq(emptySeq, nonEmptySeq, emptySeq)
fmt.Println(slices.Collect(result), ok)
// [0 1 2] true

strSeq := slices.Values([]string{"a", "b"})
strResult, ok := it.CoalesceSeq(it.Empty[string](), strSeq)
fmt.Println(slices.Collect(strResult), ok)
// [a b] true

result, ok = it.CoalesceSeq(emptySeq, emptySeq)
fmt.Println(slices.Collect(result), ok)
// [] false
```