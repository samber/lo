---
name: Empty
slug: empty
sourceRef: type_manipulation.go#L140
playUrl: "https://go.dev/play/p/P2sD0PMXw4F"
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
strResult := lo.Empty[string]()
// "" (zero value for string)

intResult := lo.Empty[int]()
// 0 (zero value for int)

sliceResult := lo.Empty[[]int]()
// []int(nil) (zero value for slice)

mapResult := lo.Empty[map[string]int]()
// map[string]int(nil) (zero value for map)

ptrResult := lo.Empty[*int]()
// nil (zero value for pointer)
```