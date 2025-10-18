---
name: RepeatBy
slug: repeatby
sourceRef: slice.go#L362
category: core
subCategory: slice
playUrl: https://go.dev/play/p/ozZLCtX_hNU
variantHelpers:
  - core#slice#repeatby
similarHelpers:
  - core#slice#times
  - core#slice#repeat
position: 220
signatures:
  - "func RepeatBy[T any](count int, predicate func(index int) T) []T"
---

Builds a slice by calling the predicate N times with the current index.

```go
lo.RepeatBy(5, func(i int) string {
    return strconv.Itoa(i * i)
})
// []string{"0", "1", "4", "9", "16"}
```


