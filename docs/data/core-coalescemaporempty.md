---
name: CoalesceMapOrEmpty
slug: coalescemaporempty
sourceRef: type_manipulation.go#L211
category: core
subCategory: type
signatures:
  - "func CoalesceMapOrEmpty[K comparable, V any](v ...map[K]V) map[K]V"
playUrl: https://go.dev/play/p/Gyo9otyvFHH
variantHelpers:
  - core#type#coalescemaporempty
similarHelpers:
  - core#type#coalesce
  - core#type#coalesceorempty
  - core#type#coalesceslice
  - core#type#coalescesliceorempty
  - core#type#coalescemap
  - core#type#empty
position: 155
---

Returns the first non-empty map from the provided arguments, or an empty map if all arguments are empty.

```go
result := lo.CoalesceMapOrEmpty(map[string]int{}, map[string]int{"a": 1}, map[string]int{"b": 2})
// map[string]int{"a": 1}

result = lo.CoalesceMapOrEmpty(map[string]int{}, map[string]int{})
// map[string]int{}

result = lo.CoalesceMapOrEmpty(map[int]string{}, map[int]string{1: "one"}, map[int]string{2: "two"})
// map[int]string{1: "one"}
```