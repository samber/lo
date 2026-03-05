---
name: RejectErr
slug: rejecterr
sourceRef: slice.go#L897
category: core
subCategory: slice
signatures:
  - "func RejectErr[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) (bool, error)) (Slice, error)"
playUrl: https://go.dev/play/p/pFCF5WVB225
variantHelpers:
  - core#slice#rejecterr
similarHelpers:
  - core#slice#reject
  - core#slice#filtererr
  - core#slice#rejectmap
  - core#slice#filterreject
position: 265
---

The opposite of FilterErr. Returns the elements for which the predicate returns false. If the predicate returns an error, iteration stops immediately and returns the error.

```go
odd, err := lo.RejectErr([]int{1, 2, 3, 4}, func(x int, index int) (bool, error) {
    if x == 3 {
        return false, errors.New("number 3 is not allowed")
    }
    return x%2 == 0, nil
})
// []int(nil), error("number 3 is not allowed")
```

```go
odd, err := lo.RejectErr([]int{1, 2, 3, 4}, func(x int, index int) (bool, error) {
    return x%2 == 0, nil
})
// []int{1, 3}, nil
```
