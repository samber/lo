---
name: CoalesceSlice
slug: coalesceslice
sourceRef: type_manipulation.go#L178
category: core
subCategory: type
signatures:
  - "func CoalesceSlice[T any](v ...[]T) ([]T, bool)"
playUrl: https://go.dev/play/p/Gyo9otyvFHH
variantHelpers:
  - core#type#coalesceslice
similarHelpers:
  - core#type#coalesce
  - core#type#coalesceorempty
  - core#type#coalescesliceorempty
  - core#type#coalescemap
  - core#type#coalescemaporempty
  - core#type#empty
position: 140
---

Returns the first non-empty slice from the provided arguments, with a boolean indicating if a non-empty slice was found.

```go
result, ok := lo.CoalesceSlice([]int{}, []int{1, 2, 3}, []int{4, 5})
// []int{1, 2, 3}, true

result, ok = lo.CoalesceSlice([]int{}, []int{})
// []int{}, false

result, ok = lo.CoalesceSlice([]string{}, []string{"a", "b"}, []string{"c"})
// []string{"a", "b"}, true
```