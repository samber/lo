---
name: PickBy
slug: pickby
sourceRef: map.go#L106
category: core
subCategory: map
playUrl: https://go.dev/play/p/kdg8GR_QMmf
variantHelpers:
  - core#map#pickby
similarHelpers:
  - core#map#omitby
  - core#map#omitbykeys
  - core#map#omitbyvalues
  - core#map#pickbykeys
  - core#map#pickbyvalues
  - core#map#filterkeys
  - core#map#filtervalues
position: 60
signatures:
  - "func PickBy[K comparable, V any, Map ~map[K]V](in Map, predicate func(key K, value V) bool) Map"
---

Returns a map of the same type filtered by a key/value predicate.

```go
m := lo.PickBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
    return value%2 == 1
})
// map[string]int{"foo": 1, "baz": 3}
```


