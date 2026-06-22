---
name: ValueOr
slug: valueor
sourceRef: map.go#L97
category: core
subCategory: map
playUrl: https://go.dev/play/p/bAq9mHErB4V
variantHelpers:
  - core#map#valueor
similarHelpers:
  - core#map#haskey
  - core#map#keys
  - core#map#values
position: 50
signatures:
  - "func ValueOr[K comparable, V any](in map[K]V, key K, fallback V) V"
---

Returns the value for a key or a fallback if the key is not present.

```go
value := lo.ValueOr(map[string]int{"foo": 1, "bar": 2}, "foo", 42)
// 1

value = lo.ValueOr(map[string]int{"foo": 1, "bar": 2}, "baz", 42)
// 42
```


