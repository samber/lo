---
name: FindDuplicatesBy
slug: findduplicatesby
sourceRef: find.go#L234
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#findduplicatesby
similarHelpers:
  - core#find#findduplicates
  - core#find#finduniques
  - core#find#finduniquesby
position: 130
signatures:
  - "func FindDuplicatesBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) Slice"
---

Returns a slice with the first occurrence of each duplicated element by key, preserving order.

```go
lo.FindDuplicatesBy([]int{3, 4, 5, 6, 7}, func(i int) int {
    return i % 3
})
// []int{3, 4}
```


