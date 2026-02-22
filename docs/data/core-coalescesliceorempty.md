---
name: CoalesceSliceOrEmpty
slug: coalescesliceorempty
sourceRef: type_manipulation.go#L189
category: core
subCategory: type
signatures:
  - "func CoalesceSliceOrEmpty[T any](v ...[]T) []T"
playUrl: https://go.dev/play/p/Gyo9otyvFHH
variantHelpers:
  - core#type#coalescesliceorempty
similarHelpers:
  - core#type#coalesce
  - core#type#coalesceorempty
  - core#type#coalesceslice
  - core#type#coalescemap
  - core#type#coalescemaporempty
  - core#type#empty
position: 145
---

Returns the first non-empty slice from the provided arguments, or an empty slice if all arguments are empty.

```go
result := lo.CoalesceSliceOrEmpty([]int{}, []int{1, 2, 3}, []int{4, 5})
// []int{1, 2, 3}

result = lo.CoalesceSliceOrEmpty([]int{}, []int{})
// []int{}

result = lo.CoalesceSliceOrEmpty([]string{}, []string{"a", "b"}, []string{"c"})
// []string{"a", "b"}
```