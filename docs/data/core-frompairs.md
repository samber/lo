---
name: FromPairs
slug: frompairs
sourceRef: map.go#L214
category: core
subCategory: map
playUrl: https://go.dev/play/p/oIr5KHFGCEN
variantHelpers:
  - core#map#frompairs
similarHelpers:
  - core#map#fromentries
  - core#map#entries
  - core#map#topairs
  - core#map#keys
  - core#map#values
position: 140
signatures:
  - "func FromPairs[K comparable, V any](entries []Entry[K, V]) map[K]V"
---

Transforms a slice of key/value pairs into a map. Alias of `FromEntries`.

```go
m := lo.FromPairs([]lo.Entry[string, int]{{Key: "foo", Value: 1}})
// map[string]int{"foo": 1}
```


