---
name: Coalesce
slug: coalesce
sourceRef: type_manipulation.go#L153
category: core
subCategory: type
signatures:
  - "func Coalesce[T comparable](values ...T) (T, bool)"
playUrl: https://go.dev/play/p/Gyo9otyvFHH
variantHelpers:
  - core#type#coalesce
similarHelpers:
  - core#type#coalesceorempty
  - core#type#coalesceslice
  - core#type#coalescesliceorempty
  - core#type#coalescemap
  - core#type#coalescemaporempty
  - core#type#valueor
  - core#type#empty
  - core#type#fromptror
position: 130
---

Returns the first non-zero value from the provided comparable arguments, with a boolean indicating if a non-zero value was found.

```go
// With strings - returns first non-empty string
result, ok := lo.Coalesce("", "foo", "bar")
// result: "foo", ok: true

// All zero values - returns zero value with false
result, ok = lo.Coalesce("", "")
// result: "", ok: false

// With integers - zero is considered zero value
result, ok = lo.Coalesce(0, 42, 100)
// result: 42, ok: true

// With floats - zero is considered zero value
result, ok = lo.Coalesce(0.0, 3.14, 2.71)
// result: 3.14, ok: true

// With pointers - nil is zero value for pointer types
var s *string
str := "hello"
result, ok = lo.Coalesce(nil, &str)
// result: &str, ok: true

// All nil pointers
result, ok = lo.Coalesce[*string](nil, nil, nil)
// result: nil, ok: false
```