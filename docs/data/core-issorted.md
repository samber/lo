---
name: IsSorted
slug: issorted
sourceRef: slice.go#L722
category: core
subCategory: slice
playUrl: https://go.dev/play/p/mc3qR-t4mcx
variantHelpers:
  - core#slice#issorted
similarHelpers:
  - core#slice#issortedby
  - core#slice#min
  - core#slice#max
  - core#slice#reverse
position: 0
signatures:
  - "func IsSorted[T constraints.Ordered](collection []T) bool"
---

Checks if a slice is sorted in ascending order.

```go
lo.IsSorted([]int{0, 1, 2, 3, 4, 5})
// true
```


