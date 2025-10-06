---
name: PartialX
slug: partialx
sourceRef: func.go#L5
category: core
subCategory: function
signatures:
  - "func Partial[T1, T2, R any](f func(a T1, b T2) R, arg1 T1) func(T2) R"
  - "func Partial1[T1, T2, R any](f func(T1, T2) R, arg1 T1) func(T2) R"
  - "func Partial2[T1, T2, T3, R any](f func(T1, T2, T3) R, arg1 T1) func(T2, T3) R"
  - "func Partial3[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, arg1 T1) func(T2, T3, T4) R"
  - "func Partial4[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) R, arg1 T1) func(T2, T3, T4, T5) R"
  - "func Partial5[T1, T2, T3, T4, T5, T6, R any](f func(T1, T2, T3, T4, T5, T6) R, arg1 T1) func(T2, T3, T4, T5, T6) R"
playUrl: https://go.dev/play/p/Sy1gAQiQZ3v
variantHelpers:
  - core#function#partial
  - core#function#partial1
  - core#function#partial2
  - core#function#partial3
  - core#function#partial4
  - core#function#partial5
similarHelpers:
  - core#condition#if
  - core#condition#ternary
  - core#condition#switch
  - core#retry#attempt
  - core#retry#attemptwithdelay
position: 0
---

Pre-binds the first argument of a function. Variants support functions taking from 2 up to 6 input parameters.

```go
add := func(x, y int) int {
    return x + y
}
add10 := lo.Partial(add, 10)
sum := add10(5)
// 15
```


