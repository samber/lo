---
name: FilterMap
slug: filtermap
sourceRef: slice.go#L58
category: core
subCategory: slice
playUrl: https://go.dev/play/p/CgHYNUpOd1I
variantHelpers:
  - core#slice#filtermap
similarHelpers:
  - core#slice#map
  - core#slice#filter
  - core#slice#uniqmap
  - core#slice#rejectmap
  - core#slice#filtermaptoslice
  - core#slice#filtertake
  - parallel#slice#filtermap
position: 30
signatures:
  - "func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R"
---

Returns a slice obtained after both filtering and mapping using the given callback function.

The callback function should return two values: the result of the mapping operation and whether the result element should be included or not.

```go
matching := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
    if strings.HasSuffix(x, "pu") {
        return "xpu", true
    }
    return "", false
})
// []string{"xpu", "xpu"}
```
