---
name: IsEmpty
slug: isempty
sourceRef: type_manipulation.go#L143
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

Returns true if the value is empty (zero value) for comparable types. This works with strings, numbers, slices, maps, pointers, etc.

```go
result := lo.IsEmpty("")
// true (empty string)

result = lo.IsEmpty("hello")
// false

result = lo.IsEmpty(0)
// true (zero value for int)

result = lo.IsEmpty(42)
// false

result = lo.IsEmpty([]int{})
// true (empty slice)

result = lo.IsEmpty([]int{1, 2, 3})
// false

result = lo.IsEmpty(map[string]int{})
// true (empty map)

var ptr *int
result = lo.IsEmpty(ptr)
// true (nil pointer)
```