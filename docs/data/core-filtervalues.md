---
name: FilterValues
slug: filtervalues
sourceRef: map.go#L365
category: core
subCategory: map
playUrl: https://go.dev/play/p/YVD5r_h-LX-
variantHelpers:
  - core#map#filtervalues
similarHelpers:
  - core#map#filterkeys
  - core#map#pickbyvalues
  - core#map#omitbyvalues
  - core#map#pickbykeys
  - core#map#omitbykeys
position: 240
signatures:
  - "func FilterValues[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) []V"
---

Returns a slice of values for which the predicate is true.

```go
kv := map[int]string{1:"foo", 2:"bar", 3:"baz"}
result := lo.FilterValues(kv, func(k int, v string) bool {
    return v == "foo"
})
// []string{"foo"}
```


