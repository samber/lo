---
name: Drop
slug: drop
sourceRef: slice.go#L441
category: core
subCategory: slice
playUrl: https://go.dev/play/p/JswS7vXRJP2
variantHelpers:
  - core#slice#drop
similarHelpers:
  - core#slice#dropright
  - core#slice#dropwhile
  - core#slice#dropbyindex
  - core#slice#slice
  - core#slice#droprightwhile
  - core#slice#cutprefix
  - core#slice#take
position: 170
signatures:
  - "func Drop[T any, Slice ~[]T](collection Slice, n int) Slice"
---

Drops n elements from the beginning of a slice.

```go
lo.Drop([]int{0, 1, 2, 3, 4, 5}, 2)
// []int{2, 3, 4, 5}
```


