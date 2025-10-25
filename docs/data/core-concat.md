---
name: Concat
slug: concat
sourceRef: slice.go#L282
category: core
subCategory: slice
playUrl:
variantHelpers:
  - core#slice#concat
similarHelpers:
  - it#sequence#concat
  - core#slice#flatten
  - core#intersection#union
position: 160
signatures:
  - "func Concat[T any, Slice ~[]T](collections ...Slice) Slice"
---

Returns a new slice containing all the elements in collections. Concat conserves the order of the elements.

```go
list1 := []int{0, 1}
list2 := []int{2, 3, 4, 5}
flat := lo.Flatten(list1, list2)
// []int{0, 1, 2, 3, 4, 5}
```
