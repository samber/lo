---
name: TrimSuffix
slug: trimsuffix
sourceRef: slice.go#L884
category: core
subCategory: slice
playUrl: https://go.dev/play/p/IjEUrV0iofq
variantHelpers:
  - core#slice#trimsuffix
similarHelpers:
  - core#slice#trim
  - core#slice#trimleft
  - core#slice#trimright
  - core#slice#trimprefix
position: 0
signatures:
  - "func TrimSuffix[T comparable, Slice ~[]T](collection Slice, suffix Slice) Slice"
---

Removes all trailing occurrences of the given suffix from the collection.

```go
result := lo.TrimSuffix([]int{1, 2, 3, 1, 2, 4, 2, 4, 2, 4}, []int{2, 4})
// []int{1, 2, 3, 1}

result = lo.TrimSuffix([]string{"hello", "world", "hello", "test"}, []string{"test"})
// []string{"hello", "world", "hello"}
```


