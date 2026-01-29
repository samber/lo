---
name: EmptyableToPtr
slug: emptyabletoptr
sourceRef: type_manipulation.go#L45
category: core
subCategory: type
signatures:
  - "func EmptyableToPtr[T any](x T) *T"
variantHelpers:
  - core#type#emptyabletoptr
similarHelpers:
  - core#type#toptr
  - core#type#fromptr
  - core#type#fromptror
  - core#type#tosliceptr
  - core#type#fromsliceptr
position: 105
---

Returns a pointer to the provided value, or nil if the value is empty (zero value). This is useful for avoiding pointers to empty values.

```go
ptr := lo.EmptyableToPtr("")
// nil (because empty string is zero value)

ptr = lo.EmptyableToPtr("hello")
// *string pointing to "hello"

ptr = lo.EmptyableToPtr(0)
// nil (because 0 is zero value for int)

ptr = lo.EmptyableToPtr(42)
// *int pointing to 42
```