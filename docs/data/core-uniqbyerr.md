---
name: UniqByErr
slug: uniqbyerr
sourceRef: slice.go#L240
category: core
subCategory: slice
signatures:
  - "func UniqByErr[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) (U, error)) (Slice, error)"
variantHelpers:
  - core#slice#uniqbyerr
similarHelpers:
  - core#slice#uniqby
  - core#slice#uniq
  - core#slice#uniqmap
  - core#slice#partitionby
position: 111
---

Returns a duplicate-free version of a slice based on a computed key using an iteratee that can return an error. Stops iteration immediately when an error is encountered. Keeps only the first element for each unique key.

```go
// Error case - stops on first error
result, err := lo.UniqByErr([]int{0, 1, 2, 3, 4, 5}, func(i int) (int, error) {
    if i == 3 {
        return 0, fmt.Errorf("number 3 is not allowed")
    }
    return i % 3, nil
})
// []int(nil), error("number 3 is not allowed")
```

```go
// Success case
result, err := lo.UniqByErr([]int{0, 1, 2, 3, 4, 5}, func(i int) (int, error) {
    return i % 3, nil
})
// []int{0, 1, 2}, nil
```
