---
name: NewDebounce
slug: newdebounce
sourceRef: retry.go#L54
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/mz32VMK2nqe
variantHelpers:
  - core#concurrency#newdebounce
similarHelpers:
  - core#concurrency#newdebounceby
  - core#concurrency#newthrottle
  - core#concurrency#newthrottleby
  - core#concurrency#newtransaction
  - core#concurrency#synchronize
position: 0
signatures:
  - "func NewDebounce(duration time.Duration, f ...func()) (func(), func())"
---

Creates a debounced function that delays invoking the callbacks until after the wait duration has elapsed since the last call. Returns the debounced function and a cancel function.

```go
debounce, cancel := lo.NewDebounce(100 * time.Millisecond, func() { println("Called once after debounce!") })
for i := 0; i < 10; i++ { debounce() }
time.Sleep(200 * time.Millisecond)
cancel()
```


