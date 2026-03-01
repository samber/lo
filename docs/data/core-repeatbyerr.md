---
name: RepeatByErr
slug: repeatbyerr
sourceRef: slice.go#L564
category: core
subCategory: slice
signatures:
  - "func RepeatByErr[T any](count int, callback func(index int) (T, error)) ([]T, error)"
variantHelpers:
  - core#slice#repeatbyerr
similarHelpers:
  - core#slice#repeatby
  - core#slice#times
  - core#slice#repeat
position: 225
---

Builds a slice by calling the callback N times with the current index. The callback can return an error to stop iteration immediately.

```go
result, err := lo.RepeatByErr(5, func(i int) (int, error) {
    return i * i, nil
})
// []int{0, 1, 4, 9, 16}, <nil>
```

Example with error:

```go
result, err := lo.RepeatByErr(5, func(i int) (int, error) {
    if i == 3 {
        return 0, fmt.Errorf("number 3 is not allowed")
    }
    return i * i, nil
})
// []int(nil), error("number 3 is not allowed")
```
