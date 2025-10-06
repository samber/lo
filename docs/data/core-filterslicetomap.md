---
name: FilterSliceToMap
slug: filterslicetomap
sourceRef: slice.go#L414
category: core
subCategory: slice
playUrl: https://go.dev/play/p/2z0rDz2ZSGU
variantHelpers:
  - core#slice#filterslicetomap
similarHelpers:
  - core#slice#filter
  - core#slice#map
  - core#slice#groupby
  - core#slice#keyby
position: 260
signatures:
  - "func FilterSliceToMap[T any, K comparable, V any](collection []T, transform func(item T) (K, V, bool)) map[K]V"
---

Transforms elements to key/value pairs and includes them in the result only when the transform's boolean is true.

```go
list := []string{"a", "aa", "aaa"}

m := lo.FilterSliceToMap(list, func(str string) (string, int, bool) {
    return str, len(str), len(str) > 1
})
// map[string]int{"aa": 2, "aaa": 3}
```


