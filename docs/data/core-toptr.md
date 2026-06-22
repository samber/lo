---
name: ToPtr
slug: toptr
sourceRef: type_manipulation.go#L28
category: core
subCategory: type
signatures:
  - "func ToPtr[T any](x T) *T"
variantHelpers:
  - core#type#toptr
similarHelpers:
  - core#type#fromptr
  - core#type#fromptror
  - core#type#emptyabletoptr
  - core#type#tosliceptr
  - core#type#fromsliceptr
position: 90
---

Returns a pointer to the provided value.

```go
ptr := lo.ToPtr(42)
// *int pointing to 42

ptr = lo.ToPtr("hello")
// *string pointing to "hello"

ptr = lo.ToPtr([]int{1, 2, 3})
// *[]int pointing to []int{1, 2, 3}
```