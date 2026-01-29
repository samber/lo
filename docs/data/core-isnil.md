---
name: IsNil
slug: isnil
sourceRef: type_manipulation.go#L11
category: core
subCategory: type
signatures:
  - "func IsNil(x any) bool"
variantHelpers:
  - core#type#isnil
similarHelpers:
  - core#type#isnotnil
  - core#type#empty
  - core#type#isempty
  - core#type#isnotempty
position: 85
---

Returns true if the input is nil or points to a nil value. This works with pointers, interfaces, maps, slices, channels, and functions.

```go
result := lo.IsNil(nil)
// true

var ptr *int
result = lo.IsNil(ptr)
// true

result = lo.IsNil(42)
// false

result = lo.IsNil("hello")
// false

var iface interface{}
result = lo.IsNil(iface)
// true

iface = 42
result = lo.IsNil(iface)
// false
```