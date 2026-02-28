---
name: GroupByErr
slug: groupbyerr
sourceRef: slice.go#L279
category: core
subCategory: slice
signatures:
  - "func GroupByErr[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) (U, error)) (map[U]Slice, error)"
variantHelpers:
  - core#slice#groupbyerr
similarHelpers:
  - core#slice#groupby
  - core#slice#groupbymap
  - core#slice#partitionby
  - core#slice#keyby
  - parallel#slice#groupby
position: 121
---

Groups elements by a key computed from each element using an iteratee that can return an error. Stops iteration immediately when an error is encountered. The result is a map keyed by the group key with slices of original elements.

```go
// Error case - stops on first error
result, err := lo.GroupByErr([]int{0, 1, 2, 3, 4, 5}, func(i int) (int, error) {
    if i == 3 {
        return 0, fmt.Errorf("number 3 is not allowed")
    }
    return i % 3, nil
})
// map[int][]int(nil), error("number 3 is not allowed")
```

```go
// Success case
result, err := lo.GroupByErr([]int{0, 1, 2, 3, 4, 5}, func(i int) (int, error) {
    return i % 3, nil
})
// map[int][]int{0: {0, 3}, 1: {1, 4}, 2: {2, 5}}, nil
```
