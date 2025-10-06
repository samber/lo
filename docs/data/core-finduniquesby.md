---
name: FindUniquesBy
slug: finduniquesby
sourceRef: find.go#L178
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#finduniquesby
similarHelpers:
  - core#find#finduniques
  - core#find#findduplicatesby
  - core#find#uniq
  - core#find#uniqby
position: 110
signatures:
  - "func FindUniquesBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) Slice"
---

Returns a slice of elements that are unique by a computed key, preserving original order.

```go
lo.FindUniquesBy([]int{3, 4, 5, 6, 7}, func(i int) int {
    return i % 3
})
// []int{5}
```


