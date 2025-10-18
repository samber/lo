---
name: ForEach
slug: foreach
sourceRef: slice.go#L107
category: core
subCategory: slice
playUrl: https://go.dev/play/p/oofyiUPRf8t
variantHelpers:
  - core#slice#foreach
similarHelpers:
  - core#slice#times
  - core#slice#map
  - core#slice#foreachwhile
  - parallel#slice#foreach
position: 70
signatures:
  - "func ForEach[T any](collection []T, iteratee func(item T, index int))"
---

Iterates over elements of a collection and invokes the callback for each element.

```go
lo.ForEach([]string{"hello", "world"}, func(x string, _ int) {
    println(x)
})
// prints "hello\nworld\n"
```


