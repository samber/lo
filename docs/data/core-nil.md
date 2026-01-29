---
name: Nil
slug: nil
sourceRef: type_manipulation.go#L37
category: core
subCategory: type
playUrl: https://go.dev/play/p/P2sD0PMXw4F
variantHelpers:
  - core#type#nil
similarHelpers:
  - core#type#isnil
  - core#type#isnotnil
  - core#type#toptr
  - core#type#fromptr
position: 1
signatures:
  - "func Nil[T any]() *T"
---

Returns a nil pointer of type.

```go
lo.Nil[string]()
// (*string)(nil)

lo.Nil[int]()
// (*int)(nil)
```

Useful when you need a nil pointer of a specific type without declaring a variable first.