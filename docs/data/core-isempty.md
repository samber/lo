---
name: IsEmpty
slug: isempty
sourceRef: type_manipulation.go#L147
playUrl: "https://go.dev/play/p/P2sD0PMXw4F"
category: core
subCategory: type
signatures:
  - "func IsEmpty[T comparable](v T) bool"
variantHelpers:
  - core#type#isempty
similarHelpers:
  - core#type#isnotempty
  - core#type#empty
  - core#type#isnil
  - core#type#isnotnil
position: 122
---

Returns true if the value is empty (zero value) for comparable types. This works with strings, numbers, pointers, structs, etc. Slices and maps are not comparable in Go, so they cannot be used with `IsEmpty`.

```go
result := lo.IsEmpty("")
// true (empty string)

result = lo.IsEmpty("hello")
// false

result = lo.IsEmpty(0)
// true (zero value for int)

result = lo.IsEmpty(42)
// false

var ptr *int
result = lo.IsEmpty(ptr)
// true (nil pointer)
```