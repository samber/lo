---
name: WithoutEmpty
slug: withoutempty
sourceRef: intersect.go#L218
category: core
subCategory: intersect
playUrl: 
variantHelpers:
  - core#intersect#withoutempty
similarHelpers:
  - core#slice#compact
  - core#intersect#without
  - core#intersect#withoutby
position: 130
signatures:
  - "func WithoutEmpty[T comparable, Slice ~[]T](collection Slice) Slice"
---

Returns a slice excluding zero values. Deprecated: use `Compact` instead.

```go
lo.WithoutEmpty([]int{0, 2, 10})
// []int{2, 10}
```


