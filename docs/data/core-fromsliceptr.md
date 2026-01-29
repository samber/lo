---
name: FromSlicePtr
slug: fromsliceptr
sourceRef: type_manipulation.go#L88
category: core
subCategory: type
signatures:
  - "func FromSlicePtr[T any](collection []*T) []T"
variantHelpers:
  - core#type#fromsliceptr
similarHelpers:
  - core#type#toptr
  - core#type#fromptr
  - core#type#fromptror
  - core#type#emptyabletoptr
  - core#type#tosliceptr
position: 115
---

Converts a slice of pointers to a slice of values. Nil pointers are converted to zero values.

```go
a, b, c := 1, 2, 3
ptrs := []*int{&a, &b, &c}
slice := lo.FromSlicePtr(ptrs)
// []int{1, 2, 3}

a, b = "hello", "world"
ptrs = []*string{&a, nil, &b}
slice = lo.FromSlicePtr(ptrs)
// []string{"hello", "", "world"} (nil pointer becomes zero value)

ptrs = []*int{}
slice = lo.FromSlicePtr(ptrs)
// []int{}
```