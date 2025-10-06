---
name: ChunkEntries
slug: chunkentries
sourceRef: it/map.go#L138
category: it
subCategory: map
signatures:
  - "func ChunkEntries[K comparable, V any](m map[K]V, size int) iter.Seq[map[K]V]"
variantHelpers:
  - it#map#chunkentries
similarHelpers:
  - core#map#chunkentries
  - it#sequence#chunk
  - it#map#keys
  - it#map#values
position: 60
---

Chunks a map into smaller maps of the specified size. Returns a sequence of maps, each containing up to the specified number of entries.

```go
originalMap := map[string]int{
    "a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
}
result := lo.ChunkEntries(originalMap, 2)
// iter.Seq[map[string]int] yielding:
// map[string]int{"a": 1, "b": 2}
// map[string]int{"c": 3, "d": 4}
// map[string]int{"e": 5}

smallMap := map[int]string{1: "one", 2: "two"}
result = lo.ChunkEntries(smallMap, 5)
// iter.Seq[map[int]string] yielding:
// map[int]string{1: "one", 2: "two"}

largeMap := make(map[int]bool)
for i := 0; i < 10; i++ {
    largeMap[i] = true
}
result = lo.ChunkEntries(largeMap, 3)
// iter.Seq[map[int]bool] yielding 4 maps with 3, 3, 3, and 1 entries respectively
```