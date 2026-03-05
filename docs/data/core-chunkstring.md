---
name: ChunkString
slug: chunkstring
sourceRef: string.go#L130
category: core
subCategory: string
playUrl: https://go.dev/play/p/__FLTuJVz54
variantHelpers:
  - core#string#chunkstring
similarHelpers:
  - core#slice#chunk
  - core#string#substring
  - core#string#words
  - core#string#runelength
  - core#string#split
position: 20
signatures:
  - "func ChunkString[T ~string](str T, size int) []T"
---

Splits a string into chunks of the given size. The last chunk may be shorter. Returns an empty slice for empty input.

```go
lo.ChunkString("1234567", 2)
// []string{"12", "34", "56", "7"}
```


