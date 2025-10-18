---
name: MustX
slug: mustx
sourceRef: errors.go#L65
category: core
subCategory: error-handling
signatures:
  - "func Must[T any](val T, err any, messageArgs ...any) T"
  - "func Must0(err any, messageArgs ...any)"
  - "func Must1[T any](val T, err any, messageArgs ...any) T"
  - "func Must2[T1, T2 any](val1 T1, val2 T2, err any, messageArgs ...any) (T1, T2)"
  - "func Must3[T1, T2, T3 any](val1 T1, val2 T2, val3 T3, err any, messageArgs ...any) (T1, T2, T3)"
  - "func Must4[T1, T2, T3, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, err any, messageArgs ...any) (T1, T2, T3, T4)"
  - "func Must5[T1, T2, T3, T4, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, err any, messageArgs ...any) (T1, T2, T3, T4, T5)"
  - "func Must6[T1, T2, T3, T4, T5, T6 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, val6 T6, err any, messageArgs ...any) (T1, T2, T3, T4, T5, T6)"
playUrl: https://go.dev/play/p/TMoWrRp3DyC
variantHelpers:
  - core#error-handling#must
  - core#error-handling#mustx
similarHelpers:
  - core#error-handling#validate
  - core#error-handling#tryx
  - core#error-handling#tryorx
  - core#error-handling#trycatch
  - core#error-handling#trywitherrorvalue
  - core#error-handling#errorsas
  - core#error-handling#assert
position: 10
---

Panics if err is an error or false, returning successful values otherwise. Variants cover functions returning 0 to 6 values.


```go
// returns 10, panics if err is not nil
v := lo.Must(strconv.Atoi("10"))

// panics with custom message
lo.Must0(fmt.Errorf("boom"), "failed to parse")

// panics if myFunc returns an error
func myFunc() (int, string, float64, bool, error) { ... }
a, b, c, d := lo.Must4(myFunc())
```
