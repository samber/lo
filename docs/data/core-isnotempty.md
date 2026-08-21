---
name: IsNotEmpty
slug: isnotempty
sourceRef: type_manipulation.go#L154
playUrl: "https://go.dev/play/p/P2sD0PMXw4F"
category: core
subCategory: type
signatures:
  - "func IsNotEmpty[T comparable](v T) bool"
variantHelpers:
  - core#type#isnotempty
similarHelpers:
  - core#type#isempty
  - core#type#empty
  - core#type#isnil
  - core#type#isnotnil
position: 124
---

Returns true if the value is not empty (not zero value) for comparable types. This is the opposite of IsEmpty.

```go
result := lo.IsNotEmpty("")
// false (empty string)

result = lo.IsNotEmpty("hello")
// true

result = lo.IsNotEmpty(0)
// false (zero value for int)

result = lo.IsNotEmpty(42)
// true

result = lo.IsNotEmpty([3]int{})
// false (empty/zero-value array)

result = lo.IsNotEmpty([3]int{1, 2, 3})
// true

result = lo.IsNotEmpty(0.0)
// false (zero value for float64)

var ptr *int
result = lo.IsNotEmpty(ptr)
// false (nil pointer)
```