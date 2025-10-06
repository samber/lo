---
name: ForEach
slug: foreach
sourceRef: parallel/slice.go#L32
category: parallel
subCategory: slice
playUrl: https://go.dev/play/p/sCJaB3quRMC
similarHelpers:
  - core#slice#foreach
  - parallel#slice#map
position: 10
signatures:
  - "func ForEach[T any](collection []T, iteratee func(item T, index int))"
variantHelpers:
  - parallel#slice#foreach
---

Iterates over elements of a collection and invokes the predicate for each element. The predicate is called in parallel.

```go
import (
    "fmt"
    lop "github.com/samber/lo/parallel"
)

lop.ForEach([]string{"hello", "world"}, func(x string, _ int) {
    fmt.Println(x)
})
// prints lines in any order depending on scheduling
```

Useful for fire-and-forget work like publishing events or independent side effects:

```go
type Job struct{ ID int }

jobs := []Job{{1}, {2}, {3}}
lop.ForEach(jobs, func(j Job, _ int) {
    // process each job concurrently
    // send(j)
})
```


