---
name: Assert
slug: assert
sourceRef: errors.go#L359
category: core
subCategory: error-handling
playUrl: https://go.dev/play/p/Xv8LLKBMNwI
variantHelpers:
  - core#error-handling#assert
  - core#error-handling#assertf
similarHelpers:
  - core#error-handling#validate
  - core#error-handling#mustx
  - core#error-handling#tryx
  - core#error-handling#tryorx
  - core#error-handling#trycatch
  - core#error-handling#trywitherrorvalue
  - core#error-handling#errorsas
position: 80
signatures:
  - "func Assert(condition bool, message ...string)"
  - "func Assertf(condition bool, format string, args ...any)"
---

Does nothing when condition is true; otherwise panics.
```go
// Base variant with optional message
age := 12
lo.Assert(age >= 15, "user age must be >= 15")
// panics: "user age must be >= 15"

// Without message - panics with default message
x := -1
lo.Assert(x > 0)
// panics: "assertion failed: condition is not true"

// Formatted variant with custom message
age = 12
lo.Assertf(age >= 15, "user age must be >= 15, got %d", age)
// panics: "user age must be >= 15, got 12"

// When condition is true - no panic
age = 20
lo.Assert(age >= 15, "user age must be >= 15")
// continues normally
```

## Custom handler

Replace `lo.Assert` and `lo.Assertf` with your own statement:

```go
lo.Assertf = func(condition bool, format string, args ...any) {
		if !condition {
			  panic(fmt.Errorf("%s: %s", "customErr", fmt.Sprintf(format, args...)))
		}
}
```
