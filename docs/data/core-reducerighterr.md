---
name: ReduceRightErr
slug: reducerighterr
sourceRef: slice.go#L153
category: core
subCategory: slice
signatures:
  - "func ReduceRightErr[T any, R any](collection []T, accumulator func(agg R, item T, index int) (R, error), initial R) (R, error)"
variantHelpers:
  - core#slice#reducerighterr
similarHelpers:
  - core#slice#reduceright
  - core#slice#reduce
  - core#slice#reduceerr
position: 61
---

Like Reduce but iterates from right to left and accumulates into a single value using an accumulator function that can return an error. Stops iteration immediately when an error is encountered.

```go
// Error case - stops on first error (from right to left)
result, err := lo.ReduceRightErr([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, _ int) ([]int, error) {
    if len(item) > 0 && item[0] == 4 {
        return nil, fmt.Errorf("element starting with 4 is not allowed")
    }
    return append(agg, item...), nil
}, []int{})
// []int(nil), error("element starting with 4 is not allowed")
```

```go
// Success case
result, err := lo.ReduceRightErr([]int{1, 2, 3, 4}, func(agg int, item int, _ int) (int, error) {
    return agg + item, nil
}, 0)
// 10, nil
```
