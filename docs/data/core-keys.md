---
name: Keys
slug: keys
sourceRef: map.go#L5
category: core
subCategory: map
playUrl: https://go.dev/play/p/Uu11fHASqrU
variantHelpers:
  - core#map#keys
similarHelpers:
  - core#map#values
  - core#map#uniqkeys
  - core#map#entries
  - core#map#topairs
  - core#map#frompairs
  - core#map#filterkeys
position: 0
signatures:
  - "func Keys[K comparable, V any](in ...map[K]V) []K"
---

Creates a slice of the map keys.

Use the UniqKeys variant to deduplicate common keys.

```go
keys := lo.Keys(map[string]int{"foo": 1, "bar": 2})
// []string{"foo", "bar"}

keys := lo.Keys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
// []string{"foo", "bar", "baz"}

keys := lo.Keys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"bar": 3})
// []string{"foo", "bar", "bar"}
```
