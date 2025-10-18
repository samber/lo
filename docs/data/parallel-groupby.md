---
name: GroupBy
slug: groupby
sourceRef: parallel/slice.go#L73
category: parallel
subCategory: slice
playUrl: ""
similarHelpers:
  - core#slice#groupby
  - core#slice#groupbymap
  - parallel#slice#partitionby
position: 30
signatures:
  - "func GroupBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) map[U]Slice"
variantHelpers:
  - parallel#slice#groupby
---

Returns a map composed of keys generated from the results of running each element of the collection through the predicate. The predicate is called in parallel. Values keep the input order within each group.

```go
import (
    lop "github.com/samber/lo/parallel"
)

groups := lop.GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
    return i % 3
})
// map[int][]int{0: {0, 3}, 1: {1, 4}, 2: {2, 5}}
```

Custom key types work as long as they are comparable:

```go
type Kind string
groups2 := lop.GroupBy([]string{"go", "rust", "java"}, func(s string) Kind {
    if len(s) <= 2 { return "short" }
    return "long"
})
// map[Kind][]string{"short": {"go"}, "long": {"rust", "java"}}
```


