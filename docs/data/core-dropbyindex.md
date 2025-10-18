---
name: DropByIndex
slug: dropbyindex
sourceRef: slice.go#L501
category: core
subCategory: slice
playUrl: https://go.dev/play/p/bPIH4npZRxS
variantHelpers:
  - core#slice#dropbyindex
similarHelpers:
  - core#slice#drop
  - core#slice#dropright
  - core#slice#dropwhile
  - core#slice#droprightwhile
  - core#slice#slice
  - core#slice#withoutnth
  - core#slice#splice
position: 210
signatures:
  - "func DropByIndex[T any, Slice ~[]T](collection Slice, indexes ...int) Slice"
---

Drops elements from a slice by index. Negative indexes count from the end.

```go
lo.DropByIndex([]int{0, 1, 2, 3, 4, 5}, 2, 4, -1)
// []int{0, 1, 3}
```


