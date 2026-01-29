---
name: CoalesceMap
slug: coalescemap
sourceRef: type_manipulation.go#L199
category: core
subCategory: type
signatures:
  - "func CoalesceMap[K comparable, V any](v ...map[K]V) (map[K]V, bool)"
playUrl: https://go.dev/play/p/Gyo9otyvFHH
variantHelpers:
  - core#type#coalescemap
similarHelpers:
  - core#type#coalesce
  - core#type#coalesceorempty
  - core#type#coalesceslice
  - core#type#coalescesliceorempty
  - core#type#coalescemaporempty
  - core#type#empty
position: 150
---

Returns the first non-empty map from the provided arguments, with a boolean indicating if a non-empty map was found.

```go
result, ok := lo.CoalesceMap(map[string]int{}, map[string]int{"a": 1}, map[string]int{"b": 2})
// map[string]int{"a": 1}, true

result, ok = lo.CoalesceMap(map[string]int{}, map[string]int{})
// map[string]int{}, false

result, ok = lo.CoalesceMap(map[int]string{}, map[int]string{1: "one"}, map[int]string{2: "two"})
// map[int]string{1: "one"}, true
```