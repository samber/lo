---
name: Chunk
slug: chunk
sourceRef: slice.go#L209
category: core
subCategory: slice
playUrl: https://go.dev/play/p/kEMkFbdu85g
variantHelpers:
  - core#slice#chunk
similarHelpers:
  - core#slice#slice
  - core#slice#partitionby
  - core#slice#drop
  - core#slice#flatten
  - core#slice#window
  - core#slice#sliding
  - core#map#chunkentries
position: 140
signatures:
  - "func Chunk[T any, Slice ~[]T](collection Slice, size int) []Slice"
---

Splits a slice into chunks of the given size. The final chunk may be smaller.

```go
lo.Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
// [][]int{{0, 1}, {2, 3}, {4, 5}}

lo.Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)
// [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}
```

## Note

`lo.ChunkString` and `lo.Chunk` functions behave inconsistently for empty input: `lo.ChunkString("", n)` returns `[""]` instead of `[]`.

See https://github.com/samber/lo/issues/788
