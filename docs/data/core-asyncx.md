---
name: AsyncX
slug: asyncx
sourceRef: concurrency.go#L35
category: core
subCategory: concurrency
signatures:
  - "func Async[A any](f func() A) <-chan A"
  - "func Async0(f func()) <-chan struct{}"
  - "func Async1[A any](f func() A) <-chan A"
  - "func Async2[A, B any](f func() (A, B)) <-chan Tuple2[A, B]"
  - "func Async3[A, B, C any](f func() (A, B, C)) <-chan Tuple3[A, B, C]"
  - "func Async4[A, B, C, D any](f func() (A, B, C, D)) <-chan Tuple4[A, B, C, D]"
  - "func Async5[A, B, C, D, E any](f func() (A, B, C, D, E)) <-chan Tuple5[A, B, C, D, E]"
  - "func Async6[A, B, C, D, E, F any](f func() (A, B, C, D, E, F)) <-chan Tuple6[A, B, C, D, E, F]"
playUrl: https://go.dev/play/p/uo35gosuTLw
variantHelpers:
  - core#concurrency#async
  - core#concurrency#asyncx
similarHelpers:
  - core#concurrency#synchronize
  - core#concurrency#waitfor
  - core#retry#newtransaction
  - core#channel#channelseq
position: 10
---

Runs a function asynchronously and returns results via channels. Variants support 0 to 6 return values, using tuple types for multi-value results.

Variants: `Async`, `Async0..Async6`

```go
ch := lo.Async(func() int {
    time.Sleep(10 * time.Millisecond)
    return 42
})
v := <-ch

done := lo.Async0(func() {
    time.Sleep(5 * time.Millisecond)
})
<-done
```


