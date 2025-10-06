---
name: Map
slug: map
sourceRef: parallel/slice.go#L8
category: parallel
subCategory: slice
playUrl: https://go.dev/play/p/sCJaB3quRMC
similarHelpers:
  - core#slice#map
  - mutable#slice#map
  - parallel#slice#foreach
position: 0
signatures:
  - "func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R"
variantHelpers:
  - parallel#slice#map
---

Manipulates a slice and transforms it into a slice of another type. The predicate is called in parallel and results are written back in the original order.

```go
import (
    "strconv"
    lop "github.com/samber/lo/parallel"
)

out := lop.Map([]int64{1, 2, 3, 4}, func(x int64, i int) string {
    return strconv.FormatInt(x, 10)
})
// []string{"1", "2", "3", "4"}
```

Parallel execution is useful when the predicate is slow or I/O-bound:

```go
import (
    "net/http"
    "io"
    lop "github.com/samber/lo/parallel"
)

urls := []string{"https://example.com/a", "https://example.com/b"}
pages := lop.Map(urls, func(u string, _ int) string {
    resp, _ := http.Get(u)
    defer resp.Body.Close()
    b, _ := io.ReadAll(resp.Body)
    return string(b)
})
// pages keeps the same order as urls
```


