---
name: PickByKeys
slug: pickbykeys
sourceRef: map.go#L118
category: core
subCategory: map
playUrl: https://go.dev/play/p/R1imbuci9qU
variantHelpers:
  - core#map#pickbykeys
similarHelpers:
  - core#map#pickby
  - core#map#pickbyvalues
  - core#map#omitby
  - core#map#omitbykeys
  - core#map#omitbyvalues
position: 70
signatures:
  - "func PickByKeys[K comparable, V any, Map ~map[K]V](in Map, keys []K) Map"
---

Returns a map of the same type filtered by the provided keys.

```go
m := lo.PickByKeys(
    map[string]int{"foo": 1, "bar": 2, "baz": 3},
    []string{"foo", "baz"},
)
// map[string]int{"foo": 1, "baz": 3}
```


