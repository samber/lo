---
name: Chunk
slug: chunk
sourceRef: it/seq.go#L264
category: iter
subCategory: sequence
signatures:
  - "func Chunk[T any](collection iter.Seq[T], size int) iter.Seq[[]T]"
playUrl: https://go.dev/play/p/qo8esZ_L60Q
variantHelpers:
  - iter#sequence#chunk
similarHelpers:
  - core#slice#chunk
  - iter#sequence#partitionby
position: 60
---

Returns a sequence of elements split into groups of length size. The last chunk may be smaller than size.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
}
chunks := it.Chunk(seq, 2)
var result [][]int
for chunk := range chunks {
    result = append(result, chunk)
}
// result contains [1, 2], [3, 4], [5]
```

## Note

`it.ChunkString` and `it.Chunk` functions behave inconsistently for empty input: `it.ChunkString("", n)` returns `[""]` instead of `[]`.

See https://github.com/samber/lo/issues/788
