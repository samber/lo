---
name: WaitFor
slug: waitfor
sourceRef: concurrency.go#L112
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/t_wTDmubbK3
variantHelpers:
  - core#concurrency#waitfor
  - core#concurrency#waitforwithcontext
similarHelpers:
  - core#retry#attemptwithdelay
  - core#retry#attemptwhilewithdelay
  - core#time#durationx
position: 20
signatures:
  - "func WaitFor(condition func(i int) bool, timeout time.Duration, heartbeatDelay time.Duration) (totalIterations int, elapsed time.Duration, conditionFound bool)"
  - "func WaitForWithContext(ctx context.Context, condition func(ctx context.Context, i int) bool, timeout time.Duration, heartbeatDelay time.Duration) (totalIterations int, elapsed time.Duration, conditionFound bool)"
---

Runs periodically until a condition is validated. Use WaitFor for simple predicates, or WaitForWithContext when you need context cancellation/timeout inside the predicate. Both return total iterations, elapsed time, and whether the condition became true.

```go
iterations, elapsed, ok := lo.WaitFor(
    func(i int) bool {
        return i > 5
    },
    10*time.Millisecond,
    time.Millisecond,
)
```

With context:

```go
iterations, elapsed, ok := lo.WaitForWithContext(
    context.Background(),
    func(_ context.Context, i int) bool {
        return i >= 5
    },
    10*time.Millisecond,
    time.Millisecond,
)
```


