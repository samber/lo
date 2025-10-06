---
name: ChunkEntries
slug: chunkentries
sourceRef: map.go#L253
category: core
subCategory: map
playUrl: https://go.dev/play/p/X_YQL6mmoD-
variantHelpers:
  - core#map#chunkentries
similarHelpers:
  - core#slice#chunk
  - core#map#mapentries
  - core#map#keyby
  - core#map#groupby
  - core#map#values
position: 170
signatures:
  - "func ChunkEntries[K comparable, V any](m map[K]V, size int) []map[K]V"
---

Splits a map into maps of at most the given size. The last chunk may be smaller.

```go
chunks := lo.ChunkEntries(map[string]int{"a":1, "b":2, "c":3, "d":4, "e":5}, 3)
// []map[string]int{ {"a":1, "b":2, "c":3}, {"d":4, "e":5} }
```


