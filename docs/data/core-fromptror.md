---
name: FromPtrOr
slug: fromptror
sourceRef: type_manipulation.go#L66
category: core
subCategory: type
signatures:
  - "func FromPtrOr[T any](x *T, fallback T) T"
variantHelpers:
  - core#type#fromptror
similarHelpers:
  - core#type#toptr
  - core#type#fromptr
  - core#type#emptyabletoptr
  - core#type#tosliceptr
  - core#type#fromsliceptr
  - core#type#valueor
position: 100
---

Returns the value pointed to by the pointer, or the fallback value if the pointer is nil.

```go
ptr := lo.ToPtr(42)
value := lo.FromPtrOr(ptr, 0)
// 42

value = lo.FromPtrOr[string](nil, "default")
// "default"

value = lo.FromPtrOr[int](nil, -1)
// -1

ptr = nil
value = lo.FromPtrOr(ptr, 999)
// 999
```