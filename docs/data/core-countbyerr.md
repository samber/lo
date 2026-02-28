---
name: CountByErr
slug: countbyerr
sourceRef: slice.go#L863
category: core
subCategory: slice
signatures:
  - "func CountByErr[T any](collection []T, predicate func(item T) (bool, error)) (int, error)"
variantHelpers:
  - core#slice#countbyerr
similarHelpers:
  - core#slice#countby
  - core#slice#count
  - core#slice#everybyerr
  - core#slice#somebyerr
position: 5
---

Counts the number of elements for which the predicate is true. Returns an error if the predicate function fails, stopping iteration immediately.

```go
count, err := lo.CountByErr([]int{1, 5, 1}, func(i int) (bool, error) {
    if i == 5 {
        return false, fmt.Errorf("5 not allowed")
    }
    return i < 4, nil
})
// 0, error("5 not allowed")
```

```go
count, err := lo.CountByErr([]int{1, 5, 1}, func(i int) (bool, error) {
    return i < 4, nil
})
// 2, nil
```
