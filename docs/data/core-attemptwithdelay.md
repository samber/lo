---
name: AttemptWithDelay
slug: attemptwithdelay
sourceRef: retry.go#L172
category: core
subCategory: retry
playUrl: https://go.dev/play/p/tVs6CygC7m1
variantHelpers:
  - core#retry#attemptwithdelay
similarHelpers:
  - core#retry#attempt
  - core#retry#attemptwhile
  - core#retry#attemptwhilewithdelay
position: 10
signatures:
  - "func AttemptWithDelay(maxIteration int, delay time.Duration, f func(index int, duration time.Duration) error) (int, time.Duration, error)"
---

Invokes a function up to N times with a pause between calls until it returns nil. Returns iterations, elapsed, and error.

```go
iter, dur, err := lo.AttemptWithDelay(
    3,
    100*time.Millisecond,
    func(i int, d time.Duration) error {
        if i == 1 {
            return nil
        }
        return errors.New("x")
    },
)
```


