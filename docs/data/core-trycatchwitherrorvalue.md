---
name: TryCatchWithErrorValue
slug: trycatchwitherrorvalue
sourceRef: errors.go#L343
category: core
subCategory: error-handling
playUrl: https://go.dev/play/p/8Pc9gwX_GZO
variantHelpers:
  - core#error-handling#trycatchwitherrorvalue
similarHelpers:
  - core#error-handling#validate
  - core#error-handling#mustx
  - core#error-handling#tryx
  - core#error-handling#tryorx
  - core#error-handling#trywitherrorvalue
  - core#error-handling#trycatch
  - core#error-handling#errorsas
  - core#error-handling#assert
position: 60
signatures:
  - "func TryCatchWithErrorValue(callback func() error, catch func(any))"
---

Calls the catch function with the error value when the callback errors or panics.

```go
lo.TryCatchWithErrorValue(
    func() error {
        panic("error")
    },
    func(val any) {
        fmt.Println(val)
    },
)
```

