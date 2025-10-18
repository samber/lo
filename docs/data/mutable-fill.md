---
name: Fill
slug: fill
sourceRef: mutable/slice_example_test.go#L72
category: mutable
subCategory: slice
signatures:
  - "func Fill[T any, Slice ~[]T](collection Slice, initial T)"
playUrl: https://go.dev/play/p/VwR34GzqEub
variantHelpers:
  - mutable#slice#fill
similarHelpers:
  - core#slice#fill
  - core#slice#repeat
  - core#slice#times
position: 60
---

Fills all elements of a slice with the specified initial value. The operation modifies the slice in place.

```go
slice := make([]int, 5)
lo.Fill(slice, 42)
// []int{42, 42, 42, 42, 42}

slice = make([]string, 3)
lo.Fill(slice, "default")
// []string{"default", "default", "default"}

slice = make([]bool, 4)
lo.Fill(slice, true)
// []bool{true, true, true, true}

slice = []int{1, 2, 3, 4, 5}
lo.Fill(slice, 0)
// []int{0, 0, 0, 0, 0}
```