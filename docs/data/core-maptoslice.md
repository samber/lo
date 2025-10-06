---
name: MapToSlice
slug: maptoslice
sourceRef: map.go#L320
category: core
subCategory: map
playUrl: https://go.dev/play/p/ZuiCZpDt6LD
variantHelpers:
  - core#map#maptoslice
similarHelpers:
  - core#map#mapentries
  - core#map#entries
  - core#slice#map
  - core#slice#mapentries
position: 210
signatures:
  - "func MapToSlice[K comparable, V any, R any](in map[K]V, iteratee func(key K, value V) R) []R"
---

Transforms a map into a slice by applying an predicate to each key/value pair.

```go
m := map[int]int64{1:4, 2:5, 3:6}
s := lo.MapToSlice(m, func(k int, v int64) string {
    return fmt.Sprintf("%d_%d", k, v)
})
// []string{"1_4", "2_5", "3_6"}
```


