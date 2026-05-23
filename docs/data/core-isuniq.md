---
name: IsUniq
slug: isuniq
sourceRef: slice.go#L290
category: core
subCategory: slice
variantHelpers:
  - core#slice#isuniq
similarHelpers:
  - core#slice#isuniqby
  - core#slice#uniq
position: 120
signatures:
  - "func IsUniq[T comparable, Slice ~[]T](collection Slice) bool"
---

Returns true if all elements in the slice are unique, false otherwise. Returns true for nil and empty slices.

```go
lo.IsUniq([]int{1, 2, 3})
// true

lo.IsUniq([]int{1, 2, 1})
// false
```


