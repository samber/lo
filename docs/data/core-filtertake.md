---
name: FilterTake
slug: filtertake
sourceRef: slice.go#L652
category: core
subCategory: slice
variantHelpers:
  - core#slice#filtertake
similarHelpers:
  - core#slice#filter
  - core#slice#take
  - core#slice#takewhile
  - core#slice#filtermap
  - core#slice#filterreject
position: 125
signatures:
  - "func FilterTake[T any, Slice ~[]T](collection Slice, n int, predicate func(item T, index int) bool) Slice"
---

Filters elements and takes the first n elements that match the predicate. This is more efficient than chaining Filter and Take, as it stops after finding n matches.

```go
lo.FilterTake([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, func(val int, index int) bool {
    return val%2 == 0
})
// []int{2, 4, 6}

lo.FilterTake([]string{"a", "aa", "aaa", "aaaa"}, 2, func(val string, index int) bool {
    return len(val) > 1
})
// []string{"aa", "aaa"}
```

