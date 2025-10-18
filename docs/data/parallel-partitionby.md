---
name: PartitionBy
slug: partitionby
sourceRef: parallel/slice.go#L92
category: parallel
subCategory: slice
playUrl: ""
similarHelpers:
  - core#slice#partitionby
  - parallel#slice#groupby
position: 40
signatures:
  - "func PartitionBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K) []Slice"
variantHelpers:
  - parallel#slice#partitionby
---

Returns a slice of groups where contiguous elements sharing the same key are batched together. Groups are created from the results of running each element of the collection through the predicate. The predicate is called in parallel and the order of groups follows their first appearance in the collection.

```go
import (
    lop "github.com/samber/lo/parallel"
)

groups := lop.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
    if x < 0 { return "negative" }
    if x%2 == 0 { return "even" }
    return "odd"
})
// [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}}
```

Works with any comparable key type:

```go
type Bucket int
parts := lop.PartitionBy([]int{1,2,2,3,3,3}, func(x int) Bucket {
    return Bucket(x)
})
// [][]int{{1}, {2,2}, {3,3,3}}
```


