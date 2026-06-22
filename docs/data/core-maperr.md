---
name: MapErr
slug: maperr
sourceRef: slice.go#L36
category: core
subCategory: slice
signatures:
  - "func MapErr[T any, R any](collection []T, transform func(item T, index int) (R, error)) ([]R, error)"
variantHelpers:
  - core#slice#maperr
similarHelpers:
  - core#slice#map
  - core#slice#filtermap
  - core#slice#flatmap
  - core#slice#rejectmap
  - parallel#slice#map
position: 11
---

Transforms each element in a slice to a new type using a function that can return an error. Stops iteration immediately when an error is encountered.

```go
// Error case - stops on first error
result, err := lo.MapErr([]int{1, 2, 3, 4}, func(x int, _ int) (string, error) {
    if x == 3 {
        return "", fmt.Errorf("number 3 is not allowed")
    }
    return strconv.Itoa(x), nil
})
// []string(nil), error("number 3 is not allowed")
```

```go
// Success case
result, err := lo.MapErr([]int{1, 2, 3, 4}, func(x int, _ int) (string, error) {
    return strconv.Itoa(x), nil
})
// []string{"1", "2", "3", "4"}, nil
```
