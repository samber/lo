---
name: FindDuplicates
slug: findduplicates
sourceRef: find.go#L207
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#findduplicates
similarHelpers:
  - core#find#findduplicatesby
  - core#slice#uniq
  - core#slice#finduniques
position: 120
signatures:
  - "func FindDuplicates[T comparable, Slice ~[]T](collection Slice) Slice"
---

Returns a slice with the first occurrence of each duplicated element in the collection, preserving order.

```go
lo.FindDuplicates([]int{1, 2, 2, 1, 2, 3})
// []int{1, 2}
```


