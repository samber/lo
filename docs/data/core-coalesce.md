---
name: Coalesce
slug: coalesce
sourceRef: type_manipulation.go#L161
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
  - core#map#valueor
  - core#type#empty
  - core#type#fromptror
position: 130
---

Returns the first non-zero value from the provided comparable arguments, with a boolean indicating if a non-zero value was found.

```go
// With strings - returns first non-empty string
strResult, ok := lo.Coalesce("", "foo", "bar")
// strResult: "foo", ok: true

// All zero values - returns zero value with false
strResult, ok = lo.Coalesce("", "")
// strResult: "", ok: false

// With integers - zero is considered zero value
intResult, ok := lo.Coalesce(0, 42, 100)
// intResult: 42, ok: true

// With floats - zero is considered zero value
floatResult, ok := lo.Coalesce(0.0, 3.14, 2.71)
// floatResult: 3.14, ok: true

// With pointers - nil is zero value for pointer types
str := "hello"
ptrResult, ok := lo.Coalesce[*string](nil, &str)
// ptrResult: &str, ok: true

// All nil pointers
ptrResult, ok = lo.Coalesce[*string](nil, nil, nil)
// ptrResult: nil, ok: false
```