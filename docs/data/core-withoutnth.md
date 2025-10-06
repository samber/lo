---
name: WithoutNth
slug: withoutnth
sourceRef: intersect.go#L223
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/5g3F9R2H1xL
variantHelpers:
  - core#intersect#withoutnth
similarHelpers:
  - core#intersect#without
  - core#intersect#withoutby
  - core#slice#dropbyindex
position: 140
signatures:
  - "func WithoutNth[T comparable, Slice ~[]T](collection Slice, nths ...int) Slice"
---

Returns a slice excluding the elements at the given indexes.

```go
lo.WithoutNth([]int{-2, -1, 0, 1, 2}, 3)
// []int{-2, -1, 0, 2}
```


