---
name: ComparePtr
slug: compareptr
sourceRef: type_manipulation.go#L229
category: core
subCategory: type
signatures:
  - "func ComparePtr[T constraints.Ordered](a, b *T) bool"
variantHelpers:
  - core#type#compareptr
similarHelpers:
  - core#type#equalptr
  - core#type#fromptr
  - core#type#toptr
position: 110
---

ComparePtr compares two pointers and returns:                                                                                                                                             
- 0 if both are nil or *a == *b                                                                                                                                                         
- -1 if a is nil (and b is not) or *a < *b                                                                                                                                              
- 1 if b is nil (and a is not) or *a > *b                                                                                                                                               
                                                                                                                                                                                          
Nil is treated as less than any non-nil value.

```go
ptr42 := lo.ToPtr(42)
ptr42_2 := lo.ToPtr(42)
ptr100 := lo.ToPtr(100)
ptrNil := (*int)(nil)

equal1 := lo.ComparePtr(ptr42, ptr42_2)
// equal1: 0 (42 == 42)

equal2 := lo.ComparePtr(ptr42, ptr100)
// equal2: -1 (42 < 100)

equal3 := lo.ComparePtr(ptr100, ptr42)
// equal3: 1 (100 > 42)

equal4 := lo.ComparePtr(ptr42, ptrNil)
// equal4: 1 (42 > nil)

equal5 := lo.ComparePtr(ptrNil, ptr42)
// equal5: -1 (nil < 42)

equal6 := lo.ComparePtr(ptrNil, ptrNil)
// equal6: 0 (nil == nil)
```