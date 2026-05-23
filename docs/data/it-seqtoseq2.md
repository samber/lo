---
name: SeqToSeq2
slug: seqtoseq2
sourceRef: it/seq.go#L1167
category: iter
subCategory: sequence
signatures:
  - "func SeqToSeq2[T any](in iter.Seq[T]) iter.Seq2[int, T]"
playUrl: https://go.dev/play/p/V5wL9xY8nQr
variantHelpers:
  - iter#sequence#seqtoseq2
similarHelpers:
  - iter#map#seq2keyseq
  - iter#map#seq2valueseq
position: 185
---

Converts a sequence into an indexed sequence of key-value pairs. The first element of `iter.Seq2` is the index (starting from 0, incrementing by 1 for each item).

```go
seq := slices.Values([]string{"foo", "bar", "baz"})
for i, v := range it.SeqToSeq2(seq) {
    fmt.Printf("%d:%s ", i, v)
}
// 0:foo 1:bar 2:baz
```
