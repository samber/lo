---
name: ErrorsAs
slug: errorsas
sourceRef: errors.go#L351
category: core
subCategory: error-handling
playUrl: https://go.dev/play/p/8wk5rH8UfrE
variantHelpers:
  - core#error-handling#errorsas
similarHelpers:
  - core#error-handling#validate
  - core#error-handling#mustx
  - core#error-handling#tryx
  - core#error-handling#tryorx
  - core#error-handling#trycatch
  - core#error-handling#trywitherrorvalue
  - core#error-handling#trycatchwitherrorvalue
  - core#error-handling#assert
position: 70
signatures:
  - "func ErrorsAs[T error](err error) (T, bool)"
---

A generic wrapper around `errors.As` that returns the typed error and a boolean indicating success.

```go
type RateLimitError struct {
    RetryAfter time.Duration
}

func (e *RateLimitError) Error() string {
    return fmt.Sprintf("rate limited, retry after %s", e.RetryAfter)
}

err := fmt.Errorf("request failed: %w", &RateLimitError{RetryAfter: 5 * time.Second})

if rateLimitErr, ok := lo.ErrorsAs[*RateLimitError](err); ok {
    fmt.Println(rateLimitErr.RetryAfter)
    // 5s
}
```


