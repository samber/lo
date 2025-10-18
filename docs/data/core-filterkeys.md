---
name: FilterKeys
slug: filterkeys
sourceRef: map.go#L350
category: core
subCategory: map
playUrl: https://go.dev/play/p/OFlKXlPrBAe
variantHelpers:
  - core#map#filterkeys
similarHelpers:
  - core#map#filtervalues
  - core#map#pickbykeys
  - core#map#omitbykeys
  - core#map#pickbyvalues
  - core#map#omitbyvalues
position: 230
signatures:
  - "func FilterKeys[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) []K"
---

Returns a slice of keys for which the predicate is true.

```go
kv := map[int]string{1:"foo", 2:"bar", 3:"baz"}
result := lo.FilterKeys(kv, func(k int, v string) bool {
    return v == "foo"
})
// []int{1}
```


