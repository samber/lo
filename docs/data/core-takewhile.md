---
name: TakeWhile
slug: takewhile
sourceRef: slice.go#L605
category: core
subCategory: slice
variantHelpers:
  - core#slice#takewhile
similarHelpers:
  - core#slice#take
  - core#slice#dropwhile
  - core#slice#droprightwhile
  - core#slice#filter
  - core#slice#takefilter
  - core#slice#first
  - it#sequence#takewhile
position: 195
signatures:
  - "func TakeWhile[T any, Slice ~[]T](collection Slice, predicate func(item T) bool) Slice"
---

Takes elements from the beginning while the predicate returns true.

```go
lo.TakeWhile([]int{0, 1, 2, 3, 4, 5}, func(val int) bool {
    return val < 3
})
// []int{0, 1, 2}

lo.TakeWhile([]string{"a", "aa", "aaa", "aa"}, func(val string) bool {
    return len(val) <= 2
})
// []string{"a", "aa"}
```

