---
name: FilterErr
slug: filtererr
sourceRef: slice.go#L27
category: core
subCategory: slice
signatures:
  - "func FilterErr[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) (bool, error)) (Slice, error)"
playUrl: https://go.dev/play/p/Apjg3WeSi7K
variantHelpers:
  - core#slice#filtererr
similarHelpers:
  - core#slice#filter
  - core#slice#reject
  - core#slice#filtermap
  - core#slice#filterreject
position: 5
---

Iterates over a collection and returns a slice of all the elements the predicate function returns `true` for. If the predicate returns an error, iteration stops immediately and returns the error.

```go
even, err := lo.FilterErr([]int{1, 2, 3, 4}, func(x int, index int) (bool, error) {
    if x == 3 {
        return false, errors.New("number 3 is not allowed")
    }
    return x%2 == 0, nil
})
// []int(nil), error("number 3 is not allowed")
```

```go
even, err := lo.FilterErr([]int{1, 2, 3, 4}, func(x int, index int) (bool, error) {
    return x%2 == 0, nil
})
// []int{2, 4}, nil
```
