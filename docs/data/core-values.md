---
name: Values
slug: values
sourceRef: map.go#L54
category: core
subCategory: map
playUrl: https://go.dev/play/p/nnRTQkzQfF6
variantHelpers:
  - core#map#values
similarHelpers:
  - core#map#keys
  - core#map#entries
  - core#map#topairs
  - core#map#frompairs
  - core#map#uniqvalues
position: 30
signatures:
  - "func Values[K comparable, V any](in ...map[K]V) []V"
---

Creates a slice of the map values across one or more maps.

```go
values := lo.Values(map[string]int{"foo": 1, "bar": 2})
// []int{1, 2}

values = lo.Values(map[string]int{"foo": 1}, map[string]int{"bar": 2})
// []int{1, 2}
```


