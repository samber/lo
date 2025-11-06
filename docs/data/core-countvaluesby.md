---
name: CountValuesBy
slug: countvaluesby
sourceRef: slice.go#L623
category: core
subCategory: slice
playUrl: https://go.dev/play/p/2U0dG1SnOmS
variantHelpers:
  - core#slice#countvaluesby
similarHelpers:
  - core#slice#countvalues
  - core#slice#groupby
  - core#slice#map
  - core#slice#mapvalues
position: 0
signatures:
  - "func CountValuesBy[T any, U comparable](collection []T, transform func(item T) U) map[U]int"
---

Counts the number of each transformed value (equivalent to Map followed by CountValues).

```go
isEven := func(v int) bool {
    return v%2 == 0
}
lo.CountValuesBy([]int{1, 2, 2}, isEven)
// map[bool]int{false: 1, true: 2}
```


