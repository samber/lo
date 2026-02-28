---
name: FlatMapErr
slug: flatmaperr
sourceRef: slice.go#L99
category: core
subCategory: slice
signatures:
  - "func FlatMapErr[T any, R any](collection []T, transform func(item T, index int) ([]R, error)) ([]R, error)"
variantHelpers:
  - core#slice#flatmaperr
similarHelpers:
  - core#slice#flatmap
  - core#slice#maperr
  - core#slice#map
  - core#slice#filtermap
position: 41
---

Manipulates a slice and transforms and flattens it to a slice of another type using a function that can return an error. Stops iteration immediately when an error is encountered.

```go
// Error case - stops on first error
result, err := lo.FlatMapErr([]int64{0, 1, 2, 3}, func(x int64, _ int) ([]string, error) {
    if x == 2 {
        return nil, fmt.Errorf("number 2 is not allowed")
    }
    return []string{strconv.FormatInt(x, 10), strconv.FormatInt(x, 10)}, nil
})
// []string(nil), error("number 2 is not allowed")
```

```go
// Success case
result, err := lo.FlatMapErr([]int64{0, 1, 2}, func(x int64, _ int) ([]string, error) {
    return []string{strconv.FormatInt(x, 10), strconv.FormatInt(x, 10)}, nil
})
// []string{"0", "0", "1", "1", "2", "2"}, nil
```
