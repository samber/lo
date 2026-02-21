---
name: Sliding
slug: sliding
sourceRef: slice.go#L313
category: core
subCategory: slice
variantHelpers:
  - core#slice#sliding
similarHelpers:
  - core#slice#window
  - core#slice#chunk
  - core#slice#partitionby
  - core#slice#flatten
  - it#sequence#sliding
position: 150
signatures:
  - "func Sliding[T any, Slice ~[]T](collection Slice, size, step int) []Slice"
---

Creates a slice of sliding windows of a given size with a given step. If step is equal to size, windows don't overlap (similar to Chunk). If step is less than size, windows overlap.

```go
// Overlapping windows (step < size)
lo.Sliding([]int{1, 2, 3, 4, 5, 6}, 3, 1)
// [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}

// Non-overlapping windows (step == size, like Chunk)
lo.Sliding([]int{1, 2, 3, 4, 5, 6}, 3, 3)
// [][]int{{1, 2, 3}, {4, 5, 6}}

// Step > size (skipping elements)
lo.Sliding([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, 3)
// [][]int{{1, 2}, {4, 5}, {7, 8}}
```

