---
name: Fill
slug: fill
sourceRef: mutable/slice.go#L78
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
import lom "github.com/samber/lo/mutable"

slice := make([]int, 5)
lom.Fill(slice, 42)
// []int{42, 42, 42, 42, 42}

slice = make([]string, 3)
lom.Fill(slice, "default")
// []string{"default", "default", "default"}

slice = make([]bool, 4)
lom.Fill(slice, true)
// []bool{true, true, true, true}

slice = []int{1, 2, 3, 4, 5}
lom.Fill(slice, 0)
// []int{0, 0, 0, 0, 0}
```