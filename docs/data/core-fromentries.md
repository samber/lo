---
name: FromEntries
slug: fromentries
sourceRef: map.go#L201
category: core
subCategory: map
playUrl: https://go.dev/play/p/oIr5KHFGCEN
variantHelpers:
  - core#map#fromentries
similarHelpers:
  - core#map#frompairs
  - core#map#entries
  - core#map#topairs
  - core#map#keys
  - core#map#values
position: 130
signatures:
  - "func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V"
---

Transforms a slice of key/value pairs into a map.

```go
m := lo.FromEntries([]lo.Entry[string, int]{
    {Key: "foo", Value: 1},
    {Key: "bar", Value: 2},
})
// map[string]int{"foo": 1, "bar": 2}
```


