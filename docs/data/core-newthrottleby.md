---
name: NewThrottleBy
slug: newthrottleby
sourceRef: retry.go#L371
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/0Wv6oX7dHdC
variantHelpers:
  - core#concurrency#newthrottleby
similarHelpers:
  - core#concurrency#newthrottle
  - core#concurrency#newthrottlebywithcount
  - core#concurrency#newdebounce
  - core#concurrency#newdebounceby
position: 90
signatures:
  - "func NewThrottleBy[T comparable](interval time.Duration, f ...func(key T)) (throttle func(key T), reset func())"
---

Creates a throttled function per key.

```go
throttle, reset := lo.NewThrottleBy[string](100*time.Millisecond, func(key string) { println(key) })
for i := 0; i < 10; i++ { throttle("foo"); time.Sleep(30*time.Millisecond) }
reset()
```


