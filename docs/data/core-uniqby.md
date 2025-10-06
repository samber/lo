---
name: UniqBy
slug: uniqby
sourceRef: slice.go#L160
category: core
subCategory: slice
playUrl: https://go.dev/play/p/g42Z3QSb53u
variantHelpers:
  - core#slice#uniqby
similarHelpers:
  - core#slice#uniq
  - core#slice#uniqmap
  - core#slice#partitionby
position: 110
signatures:
  - "func UniqBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) Slice"
---

Returns a duplicate-free version of a slice based on a computed key. Keeps only the first element for each unique key.

```go
lo.UniqBy(
    []int{0, 1, 2, 3, 4, 5},
    func(i int) int {
        return i % 3
    },
)
// []int{0, 1, 2}
```


