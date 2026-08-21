---
name: Ternary
slug: ternary
sourceRef: condition.go#L6
category: core
subCategory: condition
playUrl: https://go.dev/play/p/t-D7WBL44h2
variantHelpers:
  - core#condition#ternary
  - core#condition#ternaryf
similarHelpers:
  - core#condition#if
  - core#condition#switch
  - core#condition#validate
  - core#function#partialx
position: 0
signatures:
  - "func Ternary[T any](condition bool, ifOutput T, elseOutput T) T"
  - "func TernaryF[T any](condition bool, ifFunc func() T, elseFunc func() T) T"
---

A single-line if/else that returns `ifOutput` when condition is true, otherwise `elseOutput`.

```go
result := lo.Ternary(true, "a", "b")
// result: "a"

result = lo.Ternary(false, "a", "b")
// result: "b"

// With numbers
age := 25
status := lo.Ternary(age >= 18, "adult", "minor")
// status: "adult"

// With complex expressions
x, y := 10, 20
max := lo.Ternary(x > y, x, y)
// max: 20
```

### TernaryF - Function Form

Function form of Ternary. Lazily evaluates the selected branch to avoid unnecessary work or nil dereferences.

```go
var s *string
str := lo.TernaryF(s == nil, func() string {
    return uuid.New().String()
}, func() string {
    return *s
})
// str: newly generated UUID (since s is nil)

// Example with expensive operations
str = "hello"
result := lo.TernaryF(s != nil, func() string {
    return *s // safe dereference only when s is not nil
}, func() string {
    return "default value"
})
// result: "hello"

// Performance benefit - only executes needed branch
result = lo.TernaryF(false, func() string {
    time.Sleep(1 * time.Second) // this won't execute
    return "slow operation"
}, func() string {
    return "fast operation" // only this executes
})
// result: "fast operation" (returns immediately)
```


