---
name: GroupByMapErr
slug: groupbymaperr
sourceRef: slice.go#L311
category: core
subCategory: slice
signatures:
  - "func GroupByMapErr[T any, K comparable, V any](collection []T, transform func(item T) (K, V, error)) (map[K][]V, error)"
variantHelpers:
  - core#slice#groupbymaperr
similarHelpers:
  - core#slice#groupbymap
  - core#slice#groupby
  - core#slice#groupbyerr
  - core#slice#partitionby
  - core#slice#keyby
  - core#map#associate
  - parallel#slice#groupby
position: 131
---

Groups items by a key computed from each element and maps each element to a value using a transform function that can return an error. Stops iteration immediately when an error is encountered.

```go
// Error case - stops on first error
result, err := lo.GroupByMapErr([]int{0, 1, 2, 3, 4, 5}, func(i int) (int, int, error) {
    if i == 3 {
        return 0, 0, fmt.Errorf("number 3 is not allowed")
    }
    return i % 3, i * 2, nil
})
// map[int][]int(nil), error("number 3 is not allowed")
```

```go
// Success case
result, err := lo.GroupByMapErr([]int{0, 1, 2, 3, 4, 5}, func(i int) (int, int, error) {
    return i % 3, i * 2, nil
})
// map[int][]int{0: {0, 6}, 1: {2, 8}, 2: {4, 10}}, nil
```
