---
name: PartitionByErr
slug: partitionbyerr
sourceRef: slice.go#L385
category: core
subCategory: slice
signatures:
  - "func PartitionByErr[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) (K, error)) ([]Slice, error)"
variantHelpers:
  - core#slice#partitionbyerr
similarHelpers:
  - core#slice#partitionby
  - core#slice#groupby
  - core#slice#groupbyerr
  - core#slice#groupbymap
  - core#slice#chunk
  - core#map#keyby
position: 151
---

Partitions a slice into groups determined by a key computed from each element using an iteratee that can return an error. Stops iteration immediately when an error is encountered. Preserves original order.

```go
// Error case - stops on first error
result, err := lo.PartitionByErr([]int{-2, -1, 0, 1, 2, 3}, func(x int) (string, error) {
    if x == 0 {
        return "", fmt.Errorf("zero is not allowed")
    }
    if x < 0 {
        return "negative", nil
    } else if x%2 == 0 {
        return "even", nil
    }
    return "odd", nil
})
// [][]int(nil), error("zero is not allowed")
```

```go
// Success case
result, err := lo.PartitionByErr([]int{-2, -1, 0, 1, 2}, func(x int) (string, error) {
    if x < 0 {
        return "negative", nil
    } else if x%2 == 0 {
        return "even", nil
    }
    return "odd", nil
})
// [][]int{{-2, -1}, {0, 2}, {1}}, nil
```
