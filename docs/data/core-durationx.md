---
name: DurationX
slug: durationx
sourceRef: time.go#L7
category: core
subCategory: time
signatures:
  - "func Duration(cb func()) time.Duration"
  - "func Duration0(cb func()) time.Duration"
  - "func Duration1[T any](cb func() T) (T, time.Duration)"
  - "func Duration2[T, U any](cb func() (T, U)) (T, U, time.Duration)"
  - "func Duration3[T, U, V any](cb func() (T, U, V)) (T, U, V, time.Duration)"
  - "func Duration4[T, U, V, W any](cb func() (T, U, V, W)) (T, U, V, W, time.Duration)"
  - "func Duration5[T, U, V, W, X any](cb func() (T, U, V, W, X)) (T, U, V, W, X, time.Duration)"
  - "func Duration6[T, U, V, W, X, Y any](cb func() (T, U, V, W, X, Y)) (T, U, V, W, X, Y, time.Duration)"
  - "func Duration7[T, U, V, W, X, Y, Z any](cb func() (T, U, V, W, X, Y, Z)) (T, U, V, W, X, Y, Z, time.Duration)"
  - "func Duration8[T, U, V, W, X, Y, Z, A any](cb func() (T, U, V, W, X, Y, Z, A)) (T, U, V, W, X, Y, Z, A, time.Duration)"
  - "func Duration9[T, U, V, W, X, Y, Z, A, B any](cb func() (T, U, V, W, X, Y, Z, A, B)) (T, U, V, W, X, Y, Z, A, B, time.Duration)"
  - "func Duration10[T, U, V, W, X, Y, Z, A, B, C any](cb func() (T, U, V, W, X, Y, Z, A, B, C)) (T, U, V, W, X, Y, Z, A, B, C, time.Duration)"
playUrl: https://go.dev/play/p/HQfbBbAXaFP
variantHelpers:
  - core#time#duration
  - core#time#duration0
  - core#time#duration1
  - core#time#duration2
  - core#time#duration3
  - core#time#duration4
  - core#time#duration5
  - core#time#duration6
  - core#time#duration7
  - core#time#duration8
  - core#time#duration9
  - core#time#duration10
similarHelpers:
  - core#concurrency#waitfor
  - core#retry#attemptwithdelay
  - core#retry#attemptwhilewithdelay
  - core#retry#newdebounce
  - core#retry#newdebounceby
  - core#retry#newthrottle
  - core#retry#newthrottleby
  - core#retry#newthrottlebywithcount
position: 0
---

Measures execution time of a function. Variants return the elapsed duration alongside 0 to 10 returned values from the function.

```go
// Base variant (no return values): Duration
elapsedOnly := lo.Duration(func() {
    time.Sleep(3 * time.Millisecond)
})
_ = elapsedOnly

// Zero-return variant: Duration0
elapsed := lo.Duration0(func() {
    time.Sleep(10 * time.Millisecond)
})
_ = elapsed

// One-return variant: Duration1
v, dur := lo.Duration1(func() int {
    time.Sleep(5 * time.Millisecond)
    return 123
})
_ = v
_ = dur

// Two-return variant: Duration2
a, b, elapsed2 := lo.Duration2(func() (int, string) {
    time.Sleep(2 * time.Millisecond)
    return 7, "x"
})
_ = a
_ = b
_ = elapsed2
```


