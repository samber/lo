---
name: TryOrX
slug: tryorx
sourceRef: errors.go#L197
category: core
subCategory: error-handling
signatures:
  - "func TryOr[A any](callback func() (A, error), fallbackA A) (A, bool)"
  - "func TryOr1[A any](callback func() (A, error), fallbackA A) (A, bool)"
  - "func TryOr2[A, B any](callback func() (A, B, error), fallbackA A, fallbackB B) (A, B, bool)"
  - "func TryOr3[A, B, C any](callback func() (A, B, C, error), fallbackA A, fallbackB B, fallbackC C) (A, B, C, bool)"
  - "func TryOr4[A, B, C, D any](callback func() (A, B, C, D, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D) (A, B, C, D, bool)"
  - "func TryOr5[A, B, C, D, E any](callback func() (A, B, C, D, E, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E) (A, B, C, D, E, bool)"
  - "func TryOr6[A, B, C, D, E, F any](callback func() (A, B, C, D, E, F, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E, fallbackF F) (A, B, C, D, E, F, bool)"
playUrl: https://go.dev/play/p/B4F7Wg2Zh9X
variantHelpers:
  - core#error-handling#tryor
  - core#error-handling#tryorx
similarHelpers:
  - core#error-handling#validate
  - core#error-handling#mustx
  - core#error-handling#tryx
  - core#error-handling#trycatch
  - core#error-handling#trywitherrorvalue
  - core#error-handling#errorsas
  - core#error-handling#assert
position: 30
---

Like Try, but returns provided fallback values in case of error; also returns a success flag. Variants cover callbacks returning 1 to 6 values.

Variants:

- TryOr: `func TryOr[A any](callback func() (A, error), fallbackA A) (A, bool)`
- TryOr1: `func TryOr1[A any](callback func() (A, error), fallbackA A) (A, bool)`
- TryOr2: `func TryOr2[A, B any](callback func() (A, B, error), fallbackA A, fallbackB B) (A, B, bool)`
- TryOr3: `func TryOr3[A, B, C any](callback func() (A, B, C, error), fallbackA A, fallbackB B, fallbackC C) (A, B, C, bool)`
- TryOr4: `func TryOr4[A, B, C, D any](callback func() (A, B, C, D, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D) (A, B, C, D, bool)`
- TryOr5: `func TryOr5[A, B, C, D, E any](callback func() (A, B, C, D, E, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E) (A, B, C, D, E, bool)`
- TryOr6: `func TryOr6[A, B, C, D, E, F any](callback func() (A, B, C, D, E, F, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E, fallbackF F) (A, B, C, D, E, F, bool)`

```go
value, ok := lo.TryOr(func() (int, error) {
    return 0, fmt.Errorf("boom")
}, 123)
// value == 123, ok == false

value, ok = lo.TryOr(func() (int, error) {
    return 42, nil
}, 0)
// value == 42, ok == true
```


