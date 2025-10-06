---
name: PickByValues
slug: pickbyvalues
sourceRef: map.go#L130
category: core
subCategory: map
playUrl: https://go.dev/play/p/1zdzSvbfsJc
variantHelpers:
  - core#map#pickbyvalues
similarHelpers:
  - core#map#pickby
  - core#map#pickbykeys
  - core#map#omitby
  - core#map#omitbykeys
  - core#map#omitbyvalues
position: 80
signatures:
  - "func PickByValues[K comparable, V comparable, Map ~map[K]V](in Map, values []V) Map"
---

Returns a map of the same type filtered by the provided values.

```go
m := lo.PickByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})
// map[string]int{"foo": 1, "baz": 3}
```


