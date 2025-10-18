---
name: TrimLeft
slug: trimleft
sourceRef: slice.go#L847
category: core
subCategory: slice
playUrl: https://go.dev/play/p/74aqfAYLmyi
variantHelpers:
  - core#slice#trimleft
similarHelpers:
  - core#slice#trim
  - core#slice#trimright
  - core#slice#trimprefix
  - core#slice#trimsuffix
position: 0
signatures:
  - "func TrimLeft[T comparable, Slice ~[]T](collection Slice, cutset Slice) Slice"
---

Removes all leading elements found in the cutset from the collection.

```go
result := lo.TrimLeft([]int{0, 1, 2, 0, 3, 0}, []int{1, 0})
// []int{2, 0, 3, 0}

result = lo.TrimLeft([]string{"hello", "world", " "}, []string{" ", ""})
// []string{"hello", "world", " "}
```


