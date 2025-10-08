---
name: Attempt
slug: attempt
sourceRef: retry.go#L153
category: core
subCategory: retry
playUrl: https://go.dev/play/p/3ggJZ2ZKcMj
variantHelpers:
  - core#retry#attempt
similarHelpers:
  - core#retry#attemptwithdelay
  - core#retry#attemptwhile
  - core#retry#attemptwhilewithdelay
position: 0
signatures:
  - "func Attempt(maxIteration int, f func(index int) error) (int, error)"
---

Invokes a function up to N times until it returns nil. Returns the number of iterations attempted (1-based) and the last error. Useful for retrying operations that might fail temporarily.

```go
// Success after 3 attempts
iter, err := lo.Attempt(5, func(i int) error {
    if i == 2 {
        return nil // succeeds on 3rd attempt (index 2)
    }
    return errors.New("failed")
})
// iter: 3, err: nil
```

```go
// All attempts fail - returns last error
iter, err = lo.Attempt(3, func(i int) error {
    return fmt.Errorf("attempt %d failed", i)
})
// iter: 3, err: "attempt 2 failed" (last error from index 2)
```

```go
// Immediate success on first attempt
iter, err = lo.Attempt(5, func(i int) error {
    return nil // succeeds immediately
})
// iter: 1, err: nil
```

```go
// Zero attempts - returns error without calling function
iter, err = lo.Attempt(0, func(i int) error {
    return errors.New("should not be called")
})
// iter: 0, err: "maxIteration must be greater than 0"
```


