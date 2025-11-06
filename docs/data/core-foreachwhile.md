---
name: ForEachWhile
slug: foreachwhile
sourceRef: slice.go#L116
category: core
subCategory: slice
signatures:
  - "func ForEachWhile[T any](collection []T, predicate func(item T, index int) bool)"
playUrl: https://go.dev/play/p/dG7h9H4nJQf
variantHelpers:
  - core#slice#foreachwhile
similarHelpers:
  - core#slice#foreach
  - core#slice#filter
  - core#slice#some
  - core#slice#every
  - core#slice#droprightwhile
  - core#slice#dropwhile
  - parallel#slice#foreach
position: 80
---

Iterates over elements of a collection and invokes the predicate for each element until false is returned.

```go
numbers := []int64{1, 2, -9223372036854775808, 4}
lo.ForEachWhile(numbers, func(x int64, _ int) bool {
    if x < 0 {
        return false
    }
    fmt.Println(x)
    return true
})
// Output:
// 1
// 2
```