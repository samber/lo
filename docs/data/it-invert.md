---
name: Invert
slug: invert
sourceRef: it/map.go#L111
category: it
subCategory: map
signatures:
  - "func Invert[K, V comparable](in iter.Seq2[K, V]) iter.Seq2[V, K]"
playUrl: ""
variantHelpers:
  - it#map#invert
similarHelpers:
  - core#slice#invert
  - it#map#entries
position: 40
---

Creates a sequence composed of inverted keys and values from a sequence of key/value pairs.

Examples:

```go
entries := func(yield func(string, int) bool) {
    _ = yield("apple", 1)
    _ = yield("banana", 2)
    _ = yield("cherry", 3)
}
inverted := it.Invert(entries)
var keys []int
var values []string
for k, v := range inverted {
    keys = append(keys, k)
    values = append(values, v)
}
// keys contains 1, 2, 3 and values contains "apple", "banana", "cherry"
```