---
name: Trim
slug: trim
sourceRef: slice.go#L842
category: core
subCategory: slice
playUrl: https://go.dev/play/p/1an9mxLdRG5
variantHelpers:
  - core#slice#trim
similarHelpers:
  - core#slice#trimleft
  - core#slice#trimright
  - core#slice#trimprefix
  - core#slice#trimsuffix
position: 0
signatures:
  - "func Trim[T comparable, Slice ~[]T](collection Slice, cutset Slice) Slice"
---

Removes all leading and trailing elements in the cutset from the collection.

```go
result := lo.Trim([]int{0, 1, 2, 0, 3, 0}, []int{1, 0})
// []int{2, 0, 3}

result = lo.Trim([]string{"hello", "world", " "}, []string{" ", ""})
// []string{"hello", "world"}
```


