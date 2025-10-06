---
name: DropRight
slug: dropright
sourceRef: slice.go#L457
category: core
subCategory: slice
playUrl: https://go.dev/play/p/GG0nXkSJJa3
variantHelpers:
  - core#slice#dropright
similarHelpers:
  - core#slice#drop
  - core#slice#dropwhile
  - core#slice#droprightwhile
  - core#slice#dropbyindex
  - core#slice#slice
  - core#slice#cutsuffix
  - core#slice#trimright
position: 180
signatures:
  - "func DropRight[T any, Slice ~[]T](collection Slice, n int) Slice"
---

Drops n elements from the end of a slice.

```go
lo.DropRight([]int{0, 1, 2, 3, 4, 5}, 2)
// []int{0, 1, 2, 3}
```


