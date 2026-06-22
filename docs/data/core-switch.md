---
name: Switch
slug: switch
sourceRef: condition.go#L101
category: core
subCategory: condition
playUrl: https://go.dev/play/p/TGbKUMAeRUd
variantHelpers:
  - core#condition#switch
  - core#condition#case
  - core#condition#casef
  - core#condition#default
  - core#condition#defaultf
similarHelpers:
  - core#condition#if
  - core#condition#ternary
  - core#condition#validate
position: 20
signatures:
  - "func Switch[T comparable, R any](predicate T) *switchCase[T, R]"
  - "func (s *switchCase[T, R]) Case(val T, result R) *switchCase[T, R]"
  - "func (s *switchCase[T, R]) CaseF(val T, callback func() R) *switchCase[T, R]"
  - "func (s *switchCase[T, R]) Default(result R) R"
  - "func (s *switchCase[T, R]) DefaultF(callback func() R) R"
---

Starts a functional switch/case/default chain using a predicate value.

```go
result := lo.Switch(2).Case(1, "1").Case(2, "2").Default("3")
// "2"
```

### Case

Adds a Case branch to a Switch chain returning a constant result.

```go
result := lo.Switch(1).Case(1, "1").Default("?")
// "1"
```

### CaseF

Adds a Case branch that lazily computes its result.

```go
result := lo.Switch(2).CaseF(2, func() string {
    return "2"
}).Default("?")
// "2"
```

### Default

Completes the Switch chain by providing a default result when no Case matched.

```go
result := lo.Switch(42).Default("none")
// "none"
```

### DefaultF

Function form of Default. Lazily computes the default result when no Case matched.

```go
result := lo.Switch(0).Case(1, "1").DefaultF(func() string {
    return "3"
})
// "3"
```


