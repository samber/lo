---
name: FromPairs
slug: frompairs
sourceRef: it/map.go#L104
category: it
subCategory: map
signatures:
  - "func FromPairs[K comparable, V any](entries ...iter.Seq2[K, V]) map[K]V"
playUrl: 
variantHelpers:
  - it#map#fromentries
similarHelpers:
  - core#slice#frompairs
  - core#slice#fromentries
position: 20
---

Transforms a sequence of key/value pairs into a map. Alias of FromEntries().

```go
pairs := it.Seq2(func(yield func(string, int) bool) {
    yield("a", 1)
    yield("b", 2)
    yield("c", 3)
})

result := it.FromPairs(pairs)
// map[string]int{"a": 1, "b": 2, "c": 3}
```
