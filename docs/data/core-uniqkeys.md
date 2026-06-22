---
name: UniqKeys
slug: uniqkeys
sourceRef: map.go#L23
category: core
subCategory: map
playUrl: https://go.dev/play/p/TPKAb6ILdHk
variantHelpers:
  - core#map#uniqkeys
similarHelpers:
  - core#map#keys
  - core#map#values
  - core#map#uniqvalues
position: 10
signatures:
  - "func UniqKeys[K comparable, V any](in ...map[K]V) []K"
---

Creates a slice of unique map keys across one or more maps.

```go
keys := lo.UniqKeys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"bar": 3})
// []string{"foo", "bar"}
```


