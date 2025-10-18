---
name: Assign
slug: assign
sourceRef: map.go#L234
category: core
subCategory: map
playUrl: https://go.dev/play/p/VhwfJOyxf5o
variantHelpers:
  - core#map#assign
similarHelpers:
  - core#map#entries
  - core#map#keys
  - core#map#values
position: 160
signatures:
  - "func Assign[K comparable, V any, Map ~map[K]V](maps ...Map) Map"
---

Merges multiple maps from left to right. Later maps overwrite earlier keys.

```go
merged := lo.Assign(
    map[string]int{"a": 1, "b": 2},
    map[string]int{"b": 3, "c": 4},
)
// map[string]int{"a": 1, "b": 3, "c": 4}
```


