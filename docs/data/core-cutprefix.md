---
name: CutPrefix
slug: cutprefix
sourceRef: slice.go#L800
category: core
subCategory: slice
playUrl: https://go.dev/play/p/7Plak4a1ICl
variantHelpers:
  - core#slice#cutprefix
similarHelpers:
  - core#slice#cutsuffix
  - core#slice#cut
  - core#slice#trimprefix
  - core#slice#trimleft
  - core#slice#slice
  - core#slice#drop
  - core#slice#hasprefix
position: 0
signatures:
  - "func CutPrefix[T comparable, Slice ~[]T](collection Slice, separator Slice) (after Slice, found bool)"
---

Returns the collection without the provided leading prefix and a boolean indicating whether it was present.

```go
right, found := lo.CutPrefix([]string{"a", "b", "c", "d"}, []string{"a", "b", "c"})
// right: []string{"d"}
// found: true

right, found = lo.CutPrefix([]string{"a", "b", "c"}, []string{"b"})
// right: []string{"a", "b", "c"}
// found: false
```


