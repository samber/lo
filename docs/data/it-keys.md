---
name: Keys
slug: keys
sourceRef: it/map.go#L11
category: iter
subCategory: map
signatures:
  - "func Keys[K comparable, V any](in ...map[K]V) iter.Seq[K]"
playUrl: "https://go.dev/play/p/Fu7h-eW18QM"
variantHelpers:
  - iter#map#keys
similarHelpers:
  - core#slice#keys
  - iter#map#values
  - iter#map#uniqkeys
position: 0
---

Creates a sequence of the map keys. Accepts multiple maps and concatenates their keys.

Examples:

```go
m1 := map[string]int{
    "apple":  1,
    "banana": 2,
}
m2 := map[string]int{
    "cherry": 3,
    "date":   4,
}
keysSeq := it.Keys(m1, m2)
var result []string
for k := range keysSeq {
    result = append(result, k)
}
// result contains keys from both maps
```