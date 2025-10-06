---
name: FilterMapToSlice
slug: filtermaptoslice
sourceRef: map.go#L335
category: core
subCategory: map
playUrl: https://go.dev/play/p/jgsD_Kil9pV
variantHelpers:
  - core#map#filtermaptoslice
similarHelpers:
  - core#map#maptoslice
  - core#slice#filtermap
  - core#slice#filterreject
  - core#map#filterkeys
  - core#map#filtervalues
position: 220
signatures:
  - "func FilterMapToSlice[K comparable, V any, R any](in map[K]V, iteratee func(key K, value V) (R, bool)) []R"
---

Transforms a map into a slice using an predicate that returns a value and a boolean to include it.

```go
kv := map[int]int64{1:1, 2:2, 3:3, 4:4}
result := lo.FilterMapToSlice(kv, func(k int, v int64) (string, bool) {
    return fmt.Sprintf("%d_%d", k, v), k%2 == 0
})
// []string{"2_2", "4_4"}
```


