---
name: FindUniques
slug: finduniques
sourceRef: find.go#L152
category: core
subCategory: find
playUrl: 
variantHelpers:
  - core#find#finduniques
similarHelpers:
  - core#find#finduniquesby
  - core#find#findduplicates
  - core#find#findduplicatesby
position: 100
signatures:
  - "func FindUniques[T comparable, Slice ~[]T](collection Slice) Slice"
---

Returns a slice with elements that appear only once in the collection, preserving original order.

```go
lo.FindUniques([]int{1, 2, 2, 1, 2, 3})
// []int{3}
```


