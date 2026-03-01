---
name: FindDuplicatesByErr
slug: findduplicatesbyerr
sourceRef: find.go#L296
category: core
subCategory: find
signatures:
  - "func FindDuplicatesByErr[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) (U, error)) (Slice, error)"
variantHelpers:
  - core#find#findduplicatesbyerr
similarHelpers:
  - core#find#findduplicatesby
  - core#find#findduplicates
  - core#find#finduniques
  - core#find#finduniquesby
position: 135
---

Returns a slice with the first occurrence of each duplicated element by key, preserving order. The iteratee can return an error to stop iteration immediately.

```go
result, err := lo.FindDuplicatesByErr([]int{3, 4, 5, 6, 7}, func(i int) (int, error) {
    return i % 3, nil
})
// []int{3, 4}, <nil>
```

Example with error:

```go
result, err := lo.FindDuplicatesByErr([]int{3, 4, 5, 6, 7}, func(i int) (int, error) {
    if i == 5 {
        return 0, fmt.Errorf("number 5 is not allowed")
    }
    return i % 3, nil
})
// []int(nil), error("number 5 is not allowed")
```
