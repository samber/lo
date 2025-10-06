---
name: Count
slug: count
sourceRef: slice.go#L582
category: core
subCategory: slice
playUrl: https://go.dev/play/p/Y3FlK54yveC
variantHelpers:
  - core#slice#count
similarHelpers:
  - core#slice#countby
  - core#slice#countvalues
  - core#slice#every
  - core#slice#some
position: 0
signatures:
  - "func Count[T comparable](collection []T, value T) int"
---

Counts the number of elements in the collection that equal a given value.

```go
lo.Count([]int{1, 5, 1}, 1)
// 2
```


