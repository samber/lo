---
name: OmitByValues
slug: omitbyvalues
sourceRef: map.go#L167
category: core
subCategory: map
playUrl: https://go.dev/play/p/9UYZi-hrs8j
variantHelpers:
  - core#map#omitbyvalues
similarHelpers:
  - core#map#omitby
  - core#map#omitbykeys
  - core#map#pickby
  - core#map#pickbykeys
  - core#map#pickbyvalues
position: 110
signatures:
  - "func OmitByValues[K comparable, V comparable, Map ~map[K]V](in Map, values []V) Map"
---

Returns a map of the same type excluding the provided values.

```go
m := lo.OmitByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})
// map[string]int{"bar": 2}
```


