---
name: Reject
slug: reject
sourceRef: slice.go#L532
category: core
subCategory: slice
playUrl: https://go.dev/play/p/pFCF5WVB225
variantHelpers:
  - core#slice#reject
similarHelpers:
  - core#slice#filter
  - core#slice#filterreject
  - core#slice#rejectmap
  - core#slice#filterreject
  - mutable#slice#reject
position: 260
signatures:
  - "func Reject[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice"
---

Returns the elements for which the predicate returns false (opposite of Filter).

```go
lo.Reject(
    []int{1, 2, 3, 4},
    func(x int, _ int) bool {
        return x%2 == 0
    },
)
// []int{1, 3}
```


