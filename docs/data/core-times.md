---
name: Times
slug: times
sourceRef: slice.go#L127
category: core
subCategory: slice
playUrl: https://go.dev/play/p/vgQj3Glr6lT
variantHelpers:
  - core#slice#times
similarHelpers:
  - core#slice#foreach
  - core#slice#repeat
  - core#slice#repeatby
  - parallel#slice#times
position: 90
signatures:
  - "func Times[T any](count int, iteratee func(index int) T) []T"
---

Invokes the predicate n times, returning a slice of results. The predicate receives the index on each call.

```go
lo.Times(3, func(i int) string {
    return strconv.Itoa(i)
})
// []string{"0", "1", "2"}
```


