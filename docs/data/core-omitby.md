---
name: OmitBy
slug: omitby
sourceRef: map.go#L142
category: core
subCategory: map
playUrl: https://go.dev/play/p/EtBsR43bdsd
variantHelpers:
  - core#map#omitby
similarHelpers:
  - core#map#pickby
  - core#map#omitbykeys
  - core#map#omitbyvalues
  - core#map#pickbykeys
  - core#map#pickbyvalues
  - core#map#filterkeys
  - core#map#filtervalues
position: 90
signatures:
  - "func OmitBy[K comparable, V any, Map ~map[K]V](in Map, predicate func(key K, value V) bool) Map"
---

Returns a map of the same type excluding entries that match the predicate.

```go
m := lo.OmitBy(
    map[string]int{"foo": 1, "bar": 2, "baz": 3},
    func(key string, value int) bool {
        return value%2 == 1
    },
)
// map[string]int{"bar": 2}
```


