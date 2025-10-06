---
name: TryCatch
slug: trycatch
sourceRef: errors.go#L335
category: core
subCategory: error-handling
playUrl: https://go.dev/play/p/PnOON-EqBiU
variantHelpers:
  - core#error-handling#trycatch
similarHelpers:
  - core#error-handling#validate
  - core#error-handling#mustx
  - core#error-handling#tryx
  - core#error-handling#tryorx
  - core#error-handling#trywitherrorvalue
  - core#error-handling#trycatchwitherrorvalue
  - core#error-handling#errorsas
  - core#error-handling#assert
position: 50
signatures:
  - "func TryCatch(callback func() error, catch func())"
---

Calls the catch function when the callback errors or panics.

```go
caught := false
lo.TryCatch(func() error { panic("error") }, func() { caught = true })
// caught == true
```


