---
name: FlatMap
slug: flatmap
sourceRef: slice.go#L74
category: core
subCategory: slice
playUrl: https://go.dev/play/p/pFCF5WVB225
variantHelpers:
  - core#slice#flatmap
similarHelpers:
  - core#slice#map
  - parallel#slice#map
  - mutable#slice#map
  - core#slice#filtermap
position: 40
signatures:
  - "func FlatMap[T any, R any](collection []T, transform func(item T, index int) []R) []R"
---

Manipulates a slice and transforms and flattens it to a slice of another type. The transform function can either return a slice or a `nil`, and in the `nil` case no value is added to the final slice.

```go
out := lo.FlatMap([]int64{0, 1, 2}, func(x int64, _ int) []string {
    return []string{strconv.FormatInt(x, 10), strconv.FormatInt(x, 10)}
})
// []string{"0", "0", "1", "1", "2", "2"}
```


