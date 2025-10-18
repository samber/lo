---
name: MapValues
slug: mapvalues
sourceRef: map.go#L295
category: core
subCategory: map
playUrl: https://go.dev/play/p/T_8xAfvcf0W
variantHelpers:
  - core#map#mapvalues
similarHelpers:
  - core#map#mapkeys
  - core#map#mapentries
  - core#map#groupby
  - core#slice#map
position: 190
signatures:
  - "func MapValues[K comparable, V any, R any](in map[K]V, iteratee func(value V, key K) R) map[K]R"
---

Transforms map values using a predicate while keeping keys.

```go
in := map[int]int64{1:1, 2:2}
out := lo.MapValues(in, func(v int64, _ int) string {
    return strconv.FormatInt(v, 10)
})
// map[int]string{1:"1", 2:"2"}
```


