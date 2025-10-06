---
name: Map
slug: map
sourceRef: slice.go#L26
category: core
subCategory: slice
playUrl: https://go.dev/play/p/OkPcYAhBo0D
similarHelpers:
  - core#slice#filtermap
  - core#slice#flatmap
  - core#slice#uniqmap
  - core#slice#rejectmap
  - core#slice#mapkeys
  - core#slice#mapvalues
  - core#slice#mapentries
  - core#slice#maptoslice
  - core#slice#filtermaptoslice
  - parallel#slice#map
  - mutable#slice#map
variantHelpers:
  - core#slice#map
position: 10
signatures:
  - "func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R"
---

Manipulates a slice and transforms it to a slice of another type.

```go
transformed := lo.Map([]int64{1, 2, 3, 4}, func(x int64, index int) string {
    return strconv.FormatInt(x, 10)
})
// []string{"1", "2", "3", "4"}
```
