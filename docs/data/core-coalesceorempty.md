---
name: CoalesceOrEmpty
slug: coalesceorempty
sourceRef: type_manipulation.go#L170
category: core
subCategory: type
signatures:
  - "func CoalesceOrEmpty[T comparable](v ...T) T"
playUrl: https://go.dev/play/p/Gyo9otyvFHH
variantHelpers:
  - core#type#coalesceorempty
similarHelpers:
  - core#type#coalesce
  - core#type#coalesceslice
  - core#type#coalescesliceorempty
  - core#type#coalescemap
  - core#type#coalescemaporempty
  - core#type#valueor
  - core#type#empty
position: 135
---

Returns the first non-zero value from the provided comparable arguments, or the zero value if all arguments are zero.

```go
result := lo.CoalesceOrEmpty("", "foo", "bar")
// "foo"

result = lo.CoalesceOrEmpty("", "")
// ""

result = lo.CoalesceOrEmpty(0, 42, 100)
// 42

result = lo.CoalesceOrEmpty(0.0, 0.0, 0.0)
// 0.0
```