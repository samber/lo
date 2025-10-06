---
name: CutSuffix
slug: cutsuffix
sourceRef: slice.go#L821
category: core
subCategory: slice
playUrl: https://go.dev/play/p/7FKfBFvPTaT
variantHelpers:
  - core#slice#cutsuffix
similarHelpers:
  - core#slice#cutprefix
  - core#slice#cut
  - core#slice#trimsuffix
  - core#slice#trimright
  - core#slice#slice
  - core#slice#dropright
  - core#slice#hassuffix
position: 0
signatures:
  - "func CutSuffix[T comparable, Slice ~[]T](collection Slice, separator Slice) (before Slice, found bool)"
---

Returns the collection without the provided trailing suffix and a boolean indicating whether it was present.

```go
left, found := lo.CutSuffix([]string{"a", "b", "c", "d", "e", "f", "g"}, []string{"f", "g"})
// left: []string{"a", "b", "c", "d", "e"}
// found: true

left, found = lo.CutSuffix([]string{"a", "b", "c"}, []string{"b"})
// left: []string{"a", "b", "c"}
// found: false
```


