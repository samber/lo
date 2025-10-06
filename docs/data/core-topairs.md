---
name: ToPairs
slug: topairs
sourceRef: map.go#L195
category: core
subCategory: map
playUrl: https://go.dev/play/p/3Dhgx46gawJ
variantHelpers:
  - core#map#topairs
similarHelpers:
  - core#map#entries
  - core#map#fromentries
  - core#map#frompairs
position: 130
signatures:
  - "func ToPairs[K comparable, V any](in map[K]V) []Entry[K, V]"
---

Transforms a map into a slice of key/value pairs. Alias of `Entries`.

```go
pairs := lo.ToPairs(map[string]int{"foo": 1, "bar": 2})
// []lo.Entry[string, int]{...}
```


