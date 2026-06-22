---
name: Compact
slug: compact
sourceRef: slice.go#L706
category: core
subCategory: slice
playUrl: https://go.dev/play/p/tXiy-iK6PAc
variantHelpers:
  - core#slice#compact
similarHelpers:
  - core#slice#filter
  - core#slice#reject
  - core#slice#without
  - core#slice#withoutempty
position: 270
signatures:
  - "func Compact[T comparable, Slice ~[]T](collection Slice) Slice"
---

Returns a slice of all non-zero elements.

```go
lo.Compact([]string{"", "foo", "", "bar", ""})
// []string{"foo", "bar"}
```


