---
name: Splice
slug: splice
sourceRef: slice.go#L748
category: core
subCategory: slice
playUrl: https://go.dev/play/p/G5_GhkeSUBA
variantHelpers:
  - core#slice#splice
similarHelpers:
  - core#slice#slice
  - core#slice#drop
  - core#slice#dropright
  - core#slice#insert
position: 0
signatures:
  - "func Splice[T any, Slice ~[]T](collection Slice, i int, elements ...T) Slice"
---

Inserts multiple elements at the specified index, with support for negative indices and automatic bounds handling. Negative indices count from the end (-1 means before last element), and indices beyond the slice length append to the end.

```go
// Basic insertion at position 1
result := lo.Splice([]string{"a", "b"}, 1, "1", "2")
// result: []string{"a", "1", "2", "b"}

// Negative index: -1 means before the last element
result = lo.Splice([]string{"a", "b"}, -1, "1", "2")
// result: []string{"a", "1", "2", "b"}

// Index overflow: when index > len(slice), elements are appended
result = lo.Splice([]string{"a", "b"}, 42, "1", "2")
// result: []string{"a", "b", "1", "2"}

// Insert at beginning (index 0)
result = lo.Splice([]int{3, 4, 5}, 0, 1, 2)
// result: []int{1, 2, 3, 4, 5}

// Insert before last element with negative index
result = lo.Splice([]int{1, 2, 3}, -2, 99)
// result: []int{1, 99, 2, 3}

// No elements to insert returns original slice
result = lo.Splice([]string{"a", "b"}, 1)
// result: []string{"a", "b"}
```


