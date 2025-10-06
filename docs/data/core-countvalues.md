---
name: CountValues
slug: countvalues
sourceRef: slice.go#L610
category: core
subCategory: slice
playUrl: https://go.dev/play/p/-p-PyLT4dfy
variantHelpers:
  - core#slice#countvalues
similarHelpers:
  - core#slice#count
  - core#slice#countby
  - core#slice#countvaluesby
  - core#slice#groupby
  - core#slice#uniq
position: 0
signatures:
  - "func CountValues[T comparable](collection []T) map[T]int"
---

Counts the number of occurrences of each element in the collection.

```go
lo.CountValues([]int{1, 2, 2})
// map[int]int{1: 1, 2: 2}
```


