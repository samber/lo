---
name: ChunkString
slug: chunkstring
sourceRef: it/string.go#L130
category: it
subCategory: string
playUrl: ""
variantHelpers:
  - it#string#chunkstring
similarHelpers:
  - core#string#chunkstring
position: 0
signatures:
  - "func ChunkString[T ~string](str T, size int) iter.Seq[T]"
---

Returns a sequence of chunks of length `size` from the input string. If the string length is not a multiple of `size`, the final chunk contains the remaining characters. Panics if `size <= 0`.

Examples:

```go
// Even split
seq := it.ChunkString("123456", 2)
var out []string
for s := range seq { out = append(out, s) }
// out == []string{"12", "34", "56"}
```

```go
// Remainder chunk
seq := it.ChunkString("1234567", 2)
var out []string
for s := range seq { out = append(out, s) }
// out == []string{"12", "34", "56", "7"}
```

```go
// Empty and small inputs
seq1 := it.ChunkString("", 2)
seq2 := it.ChunkString("1", 2)
// seq1 yields ""
// seq2 yields "1"
```


