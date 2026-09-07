---
name: ForEachErr
slug: foreacherr
sourceRef: slice.go#L200
category: core
subCategory: slice
signatures:
  - "func ForEachErr[T any](collection []T, callback func(item T, index int) error) error"
variantHelpers:
  - core#slice#foreacherr
similarHelpers:
  - core#slice#foreach
  - core#slice#foreachwhile
  - core#slice#filtererr
  - core#slice#maperr
position: 75
---

Iterates over elements of a collection and invokes the callback for each element. If the callback returns a non-nil error, iteration stops immediately and that error is returned; otherwise the function returns `nil` after the last element.

```go
lo.ForEachErr([]string{"hello", "world"}, func(x string, _ int) error {
    println(x)
    return nil
})
// prints "hello\nworld\n"
// returns nil
```

```go
err := lo.ForEachErr([]int64{1, 2, -42, 4}, func(x int64, _ int) error {
    if x < 0 {
        return fmt.Errorf("%d is not allowed", x)
    }
    println(x)
    return nil
})
// prints "1\n2\n"
// returns error("-42 is not allowed")
```
