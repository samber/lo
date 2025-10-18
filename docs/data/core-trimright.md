---
name: TrimRight
slug: trimright
sourceRef: slice.go#L873
category: core
subCategory: slice
playUrl: https://go.dev/play/p/MRpAfR6sf0g
variantHelpers:
  - core#slice#trimright
similarHelpers:
  - core#slice#trim
  - core#slice#trimleft
  - core#slice#trimprefix
  - core#slice#trimsuffix
position: 0
signatures:
  - "func TrimRight[T comparable, Slice ~[]T](collection Slice, cutset Slice) Slice"
---

Removes all trailing elements found in the cutset from the collection.

```go
result := lo.TrimRight([]int{0, 1, 2, 0, 3, 0}, []int{0, 3})
// []int{0, 1, 2}

result = lo.TrimRight([]string{"hello", "world", "  "}, []string{" ", ""})
// []string{"hello", "world", ""}
```


