---
name: GroupBy
slug: groupby
sourceRef: slice.go#L180
category: core
subCategory: slice
playUrl: https://go.dev/play/p/XnQBd_v6brd
variantHelpers:
  - core#slice#groupby
similarHelpers:
  - core#slice#groupbymap
  - core#slice#partitionby
  - core#slice#keyby
  - parallel#slice#groupby
position: 120
signatures:
  - "func GroupBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) map[U]Slice"
---

Groups elements by a key computed from each element. The result is a map keyed by the group key with slices of original elements.

```go
groups := lo.GroupBy(
    []int{0, 1, 2, 3, 4, 5},
    func(i int) int {
        return i % 3
    },
)
// map[int][]int{0: {0, 3}, 1: {1, 4}, 2: {2, 5}}
```


