---
name: Shuffle
slug: shuffle
sourceRef: mutable/slice.go#L57
category: mutable
subCategory: slice
playUrl: https://go.dev/play/p/2xb3WdLjeSJ
variantHelpers:
  - "mutable#slice#shuffle"
similarHelpers:
  - core#slice#shuffle
  - core#slice#sample
  - core#slice#samples
position: 20
signatures:
  - "func Shuffle[T any, Slice ~[]T](collection Slice)"
---

Shuffles the slice in place using the Fisher–Yates algorithm. The operation mutates the original slice order.

```go
import lom "github.com/samber/lo/mutable"

list := []int{0, 1, 2, 3, 4, 5}
lom.Shuffle(list)
// list order is randomized, e.g., []int{1, 4, 0, 3, 5, 2}
```

With strings:

```go
names := []string{"alice", "bob", "carol"}
lom.Shuffle(names)
// names order is randomized
```

