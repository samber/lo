---
name: UniqValues
slug: uniqvalues
sourceRef: map.go#L72
category: core
subCategory: map
playUrl: https://go.dev/play/p/nf6bXMh7rM3
variantHelpers:
  - core#map#uniqvalues
similarHelpers:
  - core#map#values
  - core#map#uniqkeys
  - core#slice#uniq
position: 40
signatures:
  - "func UniqValues[K comparable, V comparable](in ...map[K]V) []V"
---

Creates a slice of unique map values across one or more maps.

```go
values := lo.UniqValues(map[string]int{"foo": 1, "bar": 2}, map[string]int{"bar": 2})
// []int{1, 2}
```


