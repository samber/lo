---
name: FilterMapToSliceErr
slug: filtermaptosliceerr
sourceRef: map.go#L443
category: core
subCategory: map
signatures:
  - "func FilterMapToSliceErr[K comparable, V any, R any](in map[K]V, iteratee func(key K, value V) (R, bool, error)) ([]R, error)"
playUrl: https://go.dev/play/p/YjFEORLBWvk
variantHelpers:
  - core#map#filtermaptosliceerr
similarHelpers:
  - core#map#filtermaptoslice
  - core#map#maptosliceerr
  - core#slice#filtermap
  - core#slice#filtermaperr
position: 225
---

Transforms a map into a slice using a predicate that returns a value, a boolean to include it, and an error. Stops iteration immediately on error.

```go
kv := map[int]int64{1:1, 2:2, 3:3, 4:4}
result, err := lo.FilterMapToSliceErr(kv, func(k int, v int64) (string, bool, error) {
    if k == 3 {
        return "", false, fmt.Errorf("key 3 not allowed")
    }
    return fmt.Sprintf("%d_%d", k, v), k%2 == 0, nil
})
// []string(nil), error("key 3 not allowed")
```

```go
kv := map[int]int64{1:1, 2:2, 3:3, 4:4}
result, err := lo.FilterMapToSliceErr(kv, func(k int, v int64) (string, bool, error) {
    return fmt.Sprintf("%d_%d", k, v), k%2 == 0, nil
})
// []string{"2_2", "4_4"}, nil
```
