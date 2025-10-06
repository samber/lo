---
name: MapKeys
slug: mapkeys
sourceRef: map.go#L283
category: core
subCategory: map
playUrl: https://go.dev/play/p/9_4WPIqOetJ
variantHelpers:
  - core#map#mapkeys
similarHelpers:
  - core#map#mapvalues
  - core#map#mapentries
  - core#map#keyby
  - core#slice#map
position: 180
signatures:
  - "func MapKeys[K comparable, V any, R comparable](in map[K]V, iteratee func(value V, key K) R) map[R]V"
---

Transforms map keys using a predicate while keeping values.

```go
in := map[int]int{1:1, 2:2}
out := lo.MapKeys(in, func(v int, _ int) string {
    return strconv.Itoa(v)
})
// map[string]int{"1":1, "2":2}
```


