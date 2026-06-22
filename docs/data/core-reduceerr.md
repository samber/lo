---
name: ReduceErr
slug: reduceerr
sourceRef: slice.go#L128
category: core
subCategory: slice
signatures:
  - "func ReduceErr[T any, R any](collection []T, accumulator func(agg R, item T, index int) (R, error), initial R) (R, error)"
variantHelpers:
  - core#slice#reduceerr
similarHelpers:
  - core#slice#reduce
  - core#slice#reduceright
  - core#slice#sum
  - core#slice#sumby
  - core#slice#sumbyerr
position: 51
---

Reduces a collection to a single value by accumulating results of an accumulator function that can return an error. Stops iteration immediately when an error is encountered.

```go
// Error case - stops on first error
result, err := lo.ReduceErr([]int{1, 2, 3, 4}, func(agg int, item int, _ int) (int, error) {
    if item == 3 {
        return 0, fmt.Errorf("number 3 is not allowed")
    }
    return agg + item, nil
}, 0)
// 0, error("number 3 is not allowed")
```

```go
// Success case
result, err := lo.ReduceErr([]int{1, 2, 3, 4}, func(agg int, item int, _ int) (int, error) {
    return agg + item, nil
}, 0)
// 10, nil
```
