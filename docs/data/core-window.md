---
name: Window
slug: window
sourceRef: slice.go#L289
category: core
subCategory: slice
variantHelpers:
  - core#slice#window
similarHelpers:
  - core#slice#sliding
  - core#slice#chunk
  - core#slice#partitionby
  - core#slice#flatten
  - it#sequence#window
position: 145
signatures:
  - "func Window[T any, Slice ~[]T](collection Slice, size int) []Slice"
---

Creates a slice of sliding windows of a given size. Each window overlaps with the previous one by size-1 elements. This is equivalent to `Sliding(collection, size, 1)`.

```go
lo.Window([]int{1, 2, 3, 4, 5}, 3)
// [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}

lo.Window([]float64{20, 22, 21, 23, 24}, 3)
// [][]float64{{20, 22, 21}, {22, 21, 23}, {21, 23, 24}}
```

