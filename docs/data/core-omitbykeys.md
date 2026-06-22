---
name: OmitByKeys
slug: omitbykeys
sourceRef: map.go#L154
category: core
subCategory: map
playUrl: https://go.dev/play/p/t1QjCrs-ysk
variantHelpers:
  - core#map#omitbykeys
similarHelpers:
  - core#map#omitby
  - core#map#omitbyvalues
  - core#map#pickby
  - core#map#pickbykeys
  - core#map#pickbyvalues
position: 100
signatures:
  - "func OmitByKeys[K comparable, V any, Map ~map[K]V](in Map, keys []K) Map"
---

Returns a map of the same type excluding the provided keys.

```go
m := lo.OmitByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
// map[string]int{"bar": 2}
```


