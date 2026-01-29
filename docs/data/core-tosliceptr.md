---
name: ToSlicePtr
slug: tosliceptr
sourceRef: type_manipulation.go#L76
category: core
subCategory: type
signatures:
  - "func ToSlicePtr[T any](collection []T) []*T"
variantHelpers:
  - core#type#tosliceptr
similarHelpers:
  - core#type#toptr
  - core#type#fromptr
  - core#type#fromptror
  - core#type#emptyabletoptr
  - core#type#fromsliceptr
position: 110
---

Converts a slice of values to a slice of pointers to those values.

```go
slice := []int{1, 2, 3}
ptrs := lo.ToSlicePtr(slice)
// []*int{&1, &2, &3}

slice = []string{"a", "b", "c"}
ptrs = lo.ToSlicePtr(slice)
// []*string{&"a", &"b", &"c"}

slice = []int{}
ptrs = lo.ToSlicePtr(slice)
// []*int{}
```