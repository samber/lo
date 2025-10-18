---
name: Cut
slug: cut
sourceRef: slice.go#L774
category: core
subCategory: slice
playUrl: https://go.dev/play/p/GiL3qhpIP3f
variantHelpers:
  - core#slice#cut
similarHelpers:
  - core#slice#cutprefix
  - core#slice#cutsuffix
  - core#slice#slice
  - core#slice#trim
  - core#slice#trimprefix
  - core#slice#trimsuffix
  - core#slice#partitionby
position: 0
signatures:
  - "func Cut[T comparable, Slice ~[]T](collection Slice, separator Slice) (before Slice, after Slice, found bool)"
---

Slices the collection around the first instance of the separator, returning before, after, and whether it was found.

```go
left, right, found := lo.Cut([]string{"a", "b", "c", "d", "e", "f", "g"}, []string{"b", "c", "d"})
// left: []string{"a"}
// right: []string{"e", "f", "g"}
// found: true

left, right, found = lo.Cut([]string{"a", "b", "c"}, []string{"z"})
// left: []string{"a", "b", "c"}
// right: []string{}
// found: false
```


