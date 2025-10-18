---
name: NewDebounceBy
slug: newdebounceby
sourceRef: retry.go#L137
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/d3Vpt6pxhY8
variantHelpers:
  - core#concurrency#newdebounceby
similarHelpers:
  - core#concurrency#newdebounce
  - core#concurrency#newthrottle
  - core#concurrency#newthrottleby
  - core#concurrency#newdebounceby
position: 10
signatures:
  - "func NewDebounceBy[T comparable](duration time.Duration, f ...func(key T, count int)) (func(key T), func(key T))"
---

Creates a debounced function per key that delays invoking callbacks until after the wait duration has elapsed for that key. Returns a per-key debounced function and a per-key cancel function.

```go
debounce, cancel := lo.NewDebounceBy[string](
    100*time.Millisecond,
    func(key string, count int) {
        println(key, count)
    },
)

for i := 0; i < 10; i++ {
    debounce("first")
}

time.Sleep(200 * time.Millisecond)
cancel("first")
```


