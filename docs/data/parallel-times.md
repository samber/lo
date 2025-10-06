---
name: Times
slug: times
sourceRef: parallel/slice.go#L49
category: parallel
subCategory: slice
playUrl: ""
similarHelpers:
  - core#slice#times
position: 20
signatures:
  - "func Times[T any](count int, iteratee func(index int) T) []T"
variantHelpers:
  - parallel#slice#times
---

Invokes the predicate count times, returning a slice of the results of each invocation. The predicate is called in parallel with the index as argument.

```go
import (
    "strconv"
    lop "github.com/samber/lo/parallel"
)

nums := lop.Times(5, func(i int) string {
    return strconv.Itoa(i)
})
// []string{"0", "1", "2", "3", "4"}
```

Great for generating data concurrently:

```go
ids := lop.Times(10, func(i int) string {
    return fmt.Sprintf("item-%d", i)
})
```


