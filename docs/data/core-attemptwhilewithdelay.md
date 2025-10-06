---
name: AttemptWhileWithDelay
slug: attemptwhilewithdelay
sourceRef: retry.go#L223
category: core
subCategory: retry
playUrl: https://go.dev/play/p/mhufUjJfLEF
variantHelpers:
  - core#retry#attemptwhilewithdelay
similarHelpers:
  - core#retry#attempt
  - core#retry#attemptwithdelay
  - core#retry#attemptwhile
position: 30
signatures:
  - "func AttemptWhileWithDelay(maxIteration int, delay time.Duration, f func(int, time.Duration) (error, bool)) (int, time.Duration, error)"
---

Like AttemptWhile, but pauses between attempts and returns elapsed time.

```go
count, dur, err := lo.AttemptWhileWithDelay(
    5,
    time.Millisecond,
    func(i int, d time.Duration) (error, bool) {
        return errors.New("x"), i < 1
    },
)
```


