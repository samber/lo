---
name: Invert
slug: invert
sourceRef: map.go#L222
category: core
subCategory: map
playUrl: https://go.dev/play/p/rFQ4rak6iA1
variantHelpers:
  - core#map#invert
similarHelpers:
  - core#map#entries
  - core#map#topairs
  - core#map#frompairs
position: 150
signatures:
  - "func Invert[K comparable, V comparable](in map[K]V) map[V]K"
---

Creates a map with keys and values swapped. If values are duplicated, later keys overwrite earlier ones.

```go
lo.Invert(map[string]int{"a": 1, "b": 2})
// map[int]string{1: "a", 2: "b"}
```


