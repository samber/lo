---
name: EqualPtr
slug: equalptr
sourceRef: type_manipulation.go#L219
category: core
subCategory: type
signatures:
  - "func EqualPtr[T comparable](a, b *T) bool"
variantHelpers:
  - core#type#equalptr
similarHelpers:
  - core#type#compareptr
  - core#type#fromptr
  - core#type#toptr
position: 100
---

Compares two pointers: returns true if both are nil or both point to equal values, false if only one is nil or values differ.

```go
ptr42 := lo.ToPtr(42)
ptr42_2 := lo.ToPtr(42)
ptr100 := lo.ToPtr(100)
ptrNil := (*int)(nil)

equal1 := lo.EqualPtr(ptr42, ptr42_2)
// equal1: true (both point to 42)

equal2 := lo.EqualPtr(ptr42, ptr100)
// equal2: false (42 != 100)

equal3 := lo.EqualPtr(ptr42, ptrNil)
// equal3: false (one is nil)

equal4 := lo.EqualPtr(ptrNil, ptrNil)
// equal4: true (both are nil)
```