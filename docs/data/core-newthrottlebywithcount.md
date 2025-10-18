---
name: NewThrottleByWithCount
slug: newthrottlebywithcount
sourceRef: retry.go#L377
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/vQk3ECH7_EW
variantHelpers:
  - core#concurrency#newthrottlebywithcount
similarHelpers:
  - core#concurrency#newthrottle
  - core#concurrency#newthrottleby
  - core#concurrency#newthrottlewithcount
  - core#concurrency#newdebounce
  - core#concurrency#newdebounceby
position: 100
signatures:
  - "func NewThrottleByWithCount[T comparable](interval time.Duration, count int, f ...func(key T)) (throttle func(key T), reset func())"
---

Creates a throttled function per key with a per-interval invocation limit.

```go
throttle, reset := lo.NewThrottleByWithCount[string](
    100*time.Millisecond,
    3,
    func(key string) {
        println(key)
    },
)

for i := 0; i < 10; i++ {
    throttle("foo")
}

reset()
```


