---
name: NewThrottleWithCount
slug: newthrottlewithcount
sourceRef: retry.go#L355
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/w5nc0MgWtjC
variantHelpers:
  - core#concurrency#newthrottlewithcount
similarHelpers:
  - core#concurrency#newthrottle
  - core#concurrency#newthrottleby
  - core#concurrency#newthrottlebywithcount
  - core#concurrency#newdebounce
  - core#concurrency#newdebounceby
position: 80
signatures:
  - "func NewThrottleWithCount(interval time.Duration, count int, f ...func()) (throttle func(), reset func())"
---

Creates a throttled function with a per-interval invocation limit.

```go
throttle, reset := lo.NewThrottleWithCount(100*time.Millisecond, 3, func() { println("tick") })
for i := 0; i < 10; i++ { throttle(); time.Sleep(30*time.Millisecond) }
reset()
```


