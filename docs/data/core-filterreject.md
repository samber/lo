---
name: FilterReject
slug: filterreject
sourceRef: slice.go#L565
category: core
subCategory: slice
playUrl: https://go.dev/play/p/lHSEGSznJjB
variantHelpers:
  - core#slice#filterreject
similarHelpers:
  - core#slice#filter
  - core#slice#reject
  - core#slice#partitionby
  - mutable#slice#filterreject
position: 280
signatures:
  - "func FilterReject[T any, Slice ~[]T](collection Slice, predicate func(T, int) bool) (kept Slice, rejected Slice)"
---

Returns two slices: elements kept (predicate true) and elements rejected (predicate false).

```go
kept, rejected := lo.FilterReject(
    []int{1, 2, 3, 4},
    func(x int, _ int) bool {
        return x%2 == 0
    },
)
// kept: []int{2, 4}
// rejected: []int{1, 3}
```


