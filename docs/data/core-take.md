---
name: Take
slug: take
sourceRef: slice.go#L587
category: core
subCategory: slice
variantHelpers:
  - core#slice#take
similarHelpers:
  - core#slice#takewhile
  - core#slice#drop
  - core#slice#dropright
  - core#slice#dropwhile
  - core#slice#first
  - core#slice#filtermap
  - core#slice#takefilter
  - it#sequence#take
position: 175
signatures:
  - "func Take[T any, Slice ~[]T](collection Slice, n int) Slice"
---

Takes the first n elements from a slice.

```go
lo.Take([]int{0, 1, 2, 3, 4, 5}, 3)
// []int{0, 1, 2}

lo.Take([]int{0, 1, 2}, 5)
// []int{0, 1, 2}
```

