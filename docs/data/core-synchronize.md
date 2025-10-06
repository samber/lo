---
name: Synchronize
slug: synchronize
sourceRef: concurrency.go#L21
category: core
subCategory: concurrency
playUrl: https://go.dev/play/p/X3cqROSpQmu
variantHelpers:
  - core#concurrency#synchronize
similarHelpers:
  - core#retry#newtransaction
  - core#concurrency#asyncx
  - core#concurrency#waitfor
position: 0
signatures:
  - "func Synchronize(opt ...sync.Locker) *synchronize"
---

Wraps a callback in a mutex to ensure sequential execution. Optionally accepts a custom locker.

```go
s := lo.Synchronize()
for i := 0; i < 10; i++ {
    go s.Do(func() { println("sequential") })
}
```


