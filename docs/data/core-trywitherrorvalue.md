---
name: TryWithErrorValue
slug: trywitherrorvalue
sourceRef: errors.go#L314
category: core
subCategory: error-handling
playUrl: https://go.dev/play/p/Kc7afQIT2Fs
variantHelpers:
  - core#error-handling#trywitherrorvalue
similarHelpers:
  - core#error-handling#validate
  - core#error-handling#mustx
  - core#error-handling#tryx
  - core#error-handling#tryorx
  - core#error-handling#trycatch
  - core#error-handling#trycatchwitherrorvalue
  - core#error-handling#errorsas
  - core#error-handling#assert
position: 40
signatures:
  - "func TryWithErrorValue(callback func() error) (errorValue any, ok bool)"
---

Runs a function and returns the error value (panic or error) along with a boolean indicating success.

```go
err, ok := lo.TryWithErrorValue(func() error { panic("error") })
// err == "error", ok == false
```


