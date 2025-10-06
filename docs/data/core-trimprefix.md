---
name: TrimPrefix
slug: trimprefix
sourceRef: slice.go#L858
category: core
subCategory: slice
playUrl: https://go.dev/play/p/SHO6X-YegPg
variantHelpers:
  - core#slice#trimprefix
similarHelpers:
  - core#slice#trim
  - core#slice#trimleft
  - core#slice#trimright
  - core#slice#trimsuffix
position: 0
signatures:
  - "func TrimPrefix[T comparable, Slice ~[]T](collection Slice, prefix Slice) Slice"
---

Removes all leading occurrences of the given prefix from the collection.

```go
result := lo.TrimPrefix([]int{1, 2, 1, 2, 3, 1, 2, 4}, []int{1, 2})
// []int{3, 1, 2, 4}

result = lo.TrimPrefix([]string{"hello", "world", "hello", "test"}, []string{"hello"})
// []string{"world", "hello", "test"}
```


