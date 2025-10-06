---
name: Interleave
slug: interleave
sourceRef: slice.go#L282
category: core
subCategory: slice
playUrl: https://go.dev/play/p/-RJkTLQEDVt
variantHelpers:
  - core#slice#interleave
similarHelpers:
  - core#slice#flatten
  - core#slice#chunk
  - core#slice#slice
  - core#slice#shuffle
position: 170
signatures:
  - "func Interleave[T any, Slice ~[]T](collections ...Slice) Slice"
---

Round-robins input slices by index, appending values sequentially into the result.

```go
lo.Interleave([]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9})
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

lo.Interleave([]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10})
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```


