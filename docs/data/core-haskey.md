---
name: HasKey
slug: haskey
sourceRef: map.go#L47
category: core
subCategory: map
playUrl: https://go.dev/play/p/aVwubIvECqS
variantHelpers:
  - core#map#haskey
similarHelpers:
  - core#map#valueor
  - core#map#keys
  - core#map#values
position: 20
signatures:
  - "func HasKey[K comparable, V any](in map[K]V, key K) bool"
---

Returns whether the given key exists in the map.

```go
exists := lo.HasKey(map[string]int{"foo": 1, "bar": 2}, "foo")
// true

exists = lo.HasKey(map[string]int{"foo": 1, "bar": 2}, "baz")
// false
```


