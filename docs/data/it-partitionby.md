---
name: PartitionBy
slug: partitionby
sourceRef: it/seq.go#L26
category: it
subCategory: sequence
signatures:
  - "func PartitionBy[T any, K comparable](collection iter.Seq[T], transform func(item T) K) [][]T"
variantHelpers: []
similarHelpers:
  - core#slice#partitionby
position: 171
---

PartitionBy returns a sequence of elements split into groups. The order of grouped values is
determined by the order they occur in collection. The grouping is generated from the results
of running each element of collection through transform.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
    yield(6)
}

result := it.PartitionBy(collection, func(x int) int {
    return x % 3
})
// result contains [[1, 4], [2, 5], [3, 6]]
```