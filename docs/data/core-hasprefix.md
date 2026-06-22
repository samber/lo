---
name: HasPrefix
slug: hasprefix
sourceRef: find.go#L41
category: core
subCategory: find
playUrl: https://go.dev/play/p/SrljzVDpMQM
variantHelpers:
  - core#find#hasprefix
similarHelpers:
  - core#find#hassuffix
  - core#find#contains
  - core#slice#trimleft
  - core#slice#trimprefix
position: 20
signatures:
  - "func HasPrefix[T comparable](collection []T, prefix []T) bool"
---

Returns true if a collection starts with the given prefix slice.

```go
lo.HasPrefix([]int{1, 2, 3, 4}, []int{1, 2})
// true
```


