---
name: Clone
slug: clone
sourceRef: slice.go#L741
category: core
subCategory: slice
playUrl:
variantHelpers:
  - core#slice#clone
similarHelpers:
  - core#slice#repeat
  - core#slice#fill
position: 160
signatures:
  - "func Clone[T any, Slice ~[]T](collection Slice) Slice"
---

Returns a shallow copy of the collection.

```go
in := []int{1, 2, 3, 4, 5}
cloned := lo.Clone(in)
// Verify it's a different slice by checking that modifying one doesn't affect the other
in[0] = 99
// cloned is []int{1, 2, 3, 4, 5}
```
