---
name: Empty
slug: empty
sourceRef: type_manipulation.go#L135
category: core
subCategory: type
signatures:
  - "func Empty[T any]() T"
variantHelpers:
  - core#type#empty
similarHelpers:
  - core#type#isempty
  - core#type#isnotempty
  - core#type#coalesceorempty
  - core#type#coalescesliceorempty
  - core#type#coalescemaporempty
position: 120
---

Returns the zero value for the specified type. This is useful when you need an empty value of a specific type.

```go
result := lo.Empty[string]()
// "" (zero value for string)

result = lo.Empty[int]()
// 0 (zero value for int)

result = lo.Empty[[]int]()
// []int{} (zero value for slice)

result = lo.Empty[map[string]int]()
// map[string]int{} (zero value for map)

result = lo.Empty[*int]()
// nil (zero value for pointer)
```