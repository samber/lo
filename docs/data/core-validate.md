---
name: Validate
slug: validate
sourceRef: errors.go#L13
category: core
subCategory: error-handling
playUrl: https://go.dev/play/p/vPyh51XpCBt
variantHelpers:
  - core#error-handling#validate
similarHelpers:
  - core#error-handling#mustx
  - core#error-handling#tryx
  - core#error-handling#tryorx
  - core#error-handling#assert
position: 0
signatures:
  - "func Validate(ok bool, format string, args ...any) error"
---

Creates an error when a condition is not met; returns nil when it is.

```go
slice := []string{"a"}

err := lo.Validate(len(slice) == 0, "Slice should be empty")
// error("Slice should be empty")

err := lo.Validate(len(slice) == 0, "Slice should be empty but contains %v", slice)
// error("Slice should be empty but contains [a]")
```


