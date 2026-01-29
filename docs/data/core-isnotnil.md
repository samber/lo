---
name: IsNotNil
slug: isnil
sourceRef: type_manipulation.go#L25
category: core
subCategory: type
signatures:
  - "func IsNotNil(x any) bool"
variantHelpers:
  - core#type#isnotnil
similarHelpers:
  - core#type#isnil
  - core#type#empty
  - core#type#isempty
  - core#type#isnotempty
position: 87
---

Returns true if the input is not nil and does not point to a nil value. This works with pointers, interfaces, maps, slices, channels, and functions.

```go
result := lo.IsNotNil(nil)
// false

var ptr *int
result = lo.IsNotNil(ptr)
// false

result = lo.IsNotNil(42)
// true

result = lo.IsNotNil("hello")
// true

var iface interface{}
result = lo.IsNotNil(iface)
// false

iface = 42
result = lo.IsNotNil(iface)
// true
```