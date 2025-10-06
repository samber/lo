---
name: TryX
slug: tryx
sourceRef: errors.go#L118
category: core
subCategory: error-handling
signatures:
  - "func Try(callback func() error) bool"
  - "func Try0(callback func()) bool"
  - "func Try1(callback func() error) bool"
  - "func Try2[T any](callback func() (T, error)) bool"
  - "func Try3[T, R any](callback func() (T, R, error)) bool"
  - "func Try4[T, R, S any](callback func() (T, R, S, error)) bool"
  - "func Try5[T, R, S, Q any](callback func() (T, R, S, Q, error)) bool"
  - "func Try6[T, R, S, Q, U any](callback func() (T, R, S, Q, U, error)) bool"
playUrl:
variantHelpers:
  - core#error-handling#try
  - core#error-handling#tryx
similarHelpers:
  - core#error-handling#validate
  - core#error-handling#mustx
  - core#error-handling#tryorx
  - core#error-handling#trycatch
  - core#error-handling#trywitherrorvalue
  - core#error-handling#errorsas
  - core#error-handling#assert
position: 20
---

Calls a function and returns false in case of error or panic. Variants cover callbacks returning 0 to 6 values.

```go
ok := lo.Try(func() error {
    // return an error to mark failure
    return fmt.Errorf("boom")
})
// ok == false

ok = lo.Try0(func() {
    // panics are caught and return false
    panic("boom")
})
// ok == false

ok = lo.Try2(func() (int, error) {
    return 42, nil
})
// ok == true
```


