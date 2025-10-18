---
name: HasSuffix
slug: hassuffix
sourceRef: find.go#L57
category: core
subCategory: find
playUrl: https://go.dev/play/p/bJeLetQNAON
variantHelpers:
  - core#find#hassuffix
similarHelpers:
  - core#find#hasprefix
  - core#find#contains
  - core#slice#trimright
  - core#slice#trimsuffix
position: 30
signatures:
  - "func HasSuffix[T comparable](collection []T, suffix []T) bool"
---

Returns true if a collection ends with the given suffix slice.

```go
lo.HasSuffix([]int{1, 2, 3, 4}, []int{3, 4})
// true
```


