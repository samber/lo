---
name: MapToSliceErr
slug: maptosliceerr
sourceRef: map.go#L379
category: core
subCategory: map
signatures:
  - "func MapToSliceErr[K comparable, V any, R any](in map[K]V, iteratee func(key K, value V) (R, error)) ([]R, error)"
variantHelpers:
  - core#map#maptosliceerr
similarHelpers:
  - core#map#maptoslice
  - core#map#mapentrieserr
  - core#slice#maperr
position: 215
---

Transforms a map into a slice by applying an predicate to each key/value pair. Returns an error if the iteratee function fails, stopping iteration immediately.

```go
m := map[int]int64{1: 4, 2: 5, 3: 6}
s, err := lo.MapToSliceErr(m, func(k int, v int64) (string, error) {
    if k == 2 {
        return "", fmt.Errorf("key 2 not allowed")
    }
    return fmt.Sprintf("%d_%d", k, v), nil
})
// []string(nil), error("key 2 not allowed")
```

```go
m := map[int]int64{1:4, 2:5, 3:6}
s, err := lo.MapToSliceErr(m, func(k int, v int64) (string, error) {
    return fmt.Sprintf("%d_%d", k, v), nil
})
// []string{"1_4", "2_5", "3_6"}, nil
```
