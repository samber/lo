---
name: Filter
slug: filter
sourceRef: mutable/slice.go#L11
category: mutable
subCategory: slice
playUrl: https://go.dev/play/p/0jY3Z0B7O_5
similarHelpers:
  - core#slice#filter
  - core#slice#filterreject
  - core#slice#filtermap
  - parallel#slice#filter
  - it#chunkstring
variantHelpers:
  - "mutable#slice#filter"
  - "mutable#slice#filteri"
position: 0
signatures:
  - "func Filter[T any, Slice ~[]T](collection Slice, predicate func(item T) bool) Slice"
  - "func FilterI[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice"
---

Modifies the input slice in place, keeping only the elements for which the predicate returns true. Order is preserved. Returns the shortened slice view backed by the original array.

Variants: `FilterI` accepts an index-aware predicate `(item T, index int) bool`.

```go
import lom "github.com/samber/lo/mutable"

list := []int{1, 2, 3, 4}
kept := lom.Filter(list, func(x int) bool {
    return x%2 == 0
})

// kept -> []int{2, 4}
// list is modified in place (backed by same array)
```

Another example with a struct type:

```go
type User struct { Name string; Active bool }

users := []User{{"Alex", true}, {"Bob", false}, {"Carol", true}}
active := lom.Filter(users, func(u User) bool {
    return u.Active
})

// active -> []User{{"Alex", true}, {"Carol", true}}
// users underlying storage is reused
```

Index-aware variant (FilterI):

```go
// keep even-indexed items whose length >= 2
list2 := []string{"a", "bb", "ccc", "dddd"}
kept2 := lom.FilterI(list2, func(s string, i int) bool {
    return i%2 == 0 && len(s) >= 2
})
// kept2 -> []string{"ccc"}

nums := []int{10, 11, 12, 13, 14}
evenPos := lom.FilterI(nums, func(_ int, i int) bool {
    return i%2 == 0
})
// evenPos -> []int{10, 12, 14}
```

