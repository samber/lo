---
name: If/Else
slug: if-else
sourceRef: condition.go#L31
category: core
subCategory: condition
playUrl: https://go.dev/play/p/WSw3ApMxhyW
variantHelpers:
  - core#condition#if
  - core#condition#iff
similarHelpers:
  - core#condition#switch
  - core#condition#ternary
  - core#condition#validate
position: 10
signatures:
  - "func If[T any](condition bool, result T) *ifElse[T]"
  - "func IfF[T any](condition bool, resultF func() T) *ifElse[T]"
  - "func (i *ifElse[T]) ElseIf(condition bool, result T) *ifElse[T]"
  - "func (i *ifElse[T]) ElseIfF(condition bool, resultF func() T) *ifElse[T]"
  - "func (i *ifElse[T]) Else(result T) T"
  - "func (i *ifElse[T]) ElseF(resultF func() T) T"
---

A fluent conditional builder that allows chaining If/ElseIf/Else conditions.

### If

Starts a fluent If/ElseIf/Else chain. Returns a builder that can be completed with `ElseIf`, `Else`, etc.

```go
result := lo.If(true, 1).Else(3)
// 1
```

### IfF

Function form of If. Lazily computes the initial result when the condition is true.

```go
result := lo.IfF(true, func() int {
    return 1
}).Else(3)
// 1
```

### ElseIf

Adds an ElseIf branch to an If/Else chain.

```go
result := lo.If(false, 1).ElseIf(true, 2).Else(3)
// 2
```

### ElseIfF

Function form of ElseIf. Lazily computes the branch result when the condition is true.

```go
result := lo.If(false, 1).ElseIfF(true, func() int {
    return 2
}).Else(3)
// 2
```

### Else

Completes the If/Else chain by returning the chosen result or the default provided here.

```go
result := lo.If(false, 1).ElseIf(false, 2).Else(3)
// 3
```

### ElseF

Function form of Else. Lazily computes the default result if no previous branch matched.

```go
result := lo.If(false, 1).ElseIf(false, 2).ElseF(func() int {
    return 3
})
// 3
```