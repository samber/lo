---
name: MapEntries
slug: mapentries
sourceRef: map.go#L307
category: core
subCategory: map
playUrl: https://go.dev/play/p/VuvNQzxKimT
variantHelpers:
  - core#map#mapentries
similarHelpers:
  - core#map#mapkeys
  - core#map#mapvalues
  - core#map#maptoslice
  - core#slice#map
position: 200
signatures:
  - "func MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](in map[K1]V1, iteratee func(key K1, value V1) (K2, V2)) map[K2]V2"
---

Transforms both keys and values using an predicate function.

```go
in := map[string]int{"foo":1, "bar":2}
out := lo.MapEntries(in, func(k string, v int) (int, string) {
    return v, k
})
// map[int]string{1:"foo", 2:"bar"}
```


