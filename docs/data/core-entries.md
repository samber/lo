---
name: Entries
slug: entries
sourceRef: map.go#L179
category: core
subCategory: map
playUrl: https://go.dev/play/p/_t4Xe34-Nl5
variantHelpers:
  - core#map#entries
similarHelpers:
  - core#map#fromentries
  - core#map#keys
  - core#map#values
  - core#map#toentries
position: 120
signatures:
  - "func Entries[K comparable, V any](in map[K]V) []Entry[K, V]"
---

Transforms a map into a slice of key/value pairs.

```go
entries := lo.Entries(map[string]int{"foo": 1, "bar": 2})
// []lo.Entry[string, int]{ {Key: "foo", Value: 1}, {Key: "bar", Value: 2} }
```


