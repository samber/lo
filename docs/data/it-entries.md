---
name: Entries
slug: entries
sourceRef: it/map.go#L77
category: it
subCategory: map
signatures:
  - "func Entries[K comparable, V any](in ...map[K]V) iter.Seq2[K, V]"
playUrl: "https://go.dev/play/p/N8RbJ5t6H2k"
variantHelpers:
  - it#map#entries
similarHelpers:
  - core#slice#entries
  - it#map#fromentries
  - it#map#topairs
position: 20
---

Transforms a map into a sequence of key/value pairs. Accepts multiple maps and concatenates their entries.

Examples:

```go
m := map[string]int{
    "apple":  1,
    "banana": 2,
    "cherry": 3,
}
entriesSeq := it.Entries(m)
var keys []string
var values []int
for k, v := range entriesSeq {
    keys = append(keys, k)
    values = append(values, v)
}
// keys contains map keys, values contains corresponding values
```