---
name: IsSortedByKey
slug: issortedbykey
sourceRef: slice.go#L733
category: core
subCategory: slice
playUrl: https://go.dev/play/p/wiG6XyBBu49
variantHelpers:
  - core#slice#issortedbykey
similarHelpers:
  - core#slice#issorted
  - core#slice#minby
  - core#slice#maxby
  - core#slice#reverse
position: 0
signatures:
  - "func IsSortedByKey[T any, K constraints.Ordered](collection []T, iteratee func(item T) K) bool"
---

Checks if a slice is sorted based on a key computed for each element.

```go
ok := lo.IsSortedByKey([]string{"a", "bb", "ccc"}, func(s string) int {
    return len(s)
})
// true
```
