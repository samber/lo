---
name: NewThrottle
slug: newthrottle
sourceRef: retry.go#L349
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/qQn3fm8Z7jS
variantHelpers:
  - core#concurrency#newthrottle
similarHelpers:
  - core#concurrency#newthrottleby
  - core#concurrency#newthrottlewithcount
  - core#concurrency#newdebounce
  - core#concurrency#newdebounceby
position: 70
signatures:
  - "func NewThrottle(interval time.Duration, f ...func()) (throttle func(), reset func())"
---

Creates a throttled function that invokes callbacks at most once per interval. Returns the throttled function and a reset function.

```go
throttle, reset := lo.NewThrottle(
    100*time.Millisecond,
    func() {
        println("tick")
    },
)

for i := 0; i < 10; i++ {
    throttle();
    time.Sleep(30*time.Millisecond)
}

reset()
```


