---
name: Map
slug: map
sourceRef: mutable/slice.go#L41
category: mutable
subCategory: slice
playUrl: https://go.dev/play/p/0jY3Z0B7O_5
variantHelpers:
  - "mutable#slice#map"
  - "mutable#slice#mapi"
similarHelpers:
  - core#slice#map
  - core#slice#mapkeys
  - core#slice#mapvalues
  - parallel#slice#map
position: 10
signatures:
  - "func Map[T any, Slice ~[]T](collection Slice, fn func(item T) T)"
  - "func MapI[T any, Slice ~[]T](collection Slice, fn func(item T, index int) T)"
---

Transforms each element in the slice by applying the mapper function in place. The length remains unchanged; values are overwritten in the same backing array.

Variants: `MapI` accepts an index-aware mapper `(item T, index int) T`.

```go
import lom "github.com/samber/lo/mutable"

list := []int{1, 2, 3, 4}
lom.Map(list, func(x int) int {
    return x * 2
})
// list -> []int{2, 4, 6, 8}
```

Mapping strings:

```go
words := []string{"go", "LoDash", "lo"}
lom.Map(words, func(s string) string {
    return strings.ToUpper(s)
})
// words -> []string{"GO", "LODASH", "LO"}
```

Index-aware variant (MapI):

```go
nums := []int{10, 11, 12}
// add index to each number
lom.MapI(nums, func(x int, i int) int {
    return x + i
})
// nums -> []int{10, 12, 14}

vals := []string{"a", "b", "c", "d"}
lom.MapI(vals, func(s string, i int) string {
    if i%2 == 0 { return s + "!" }
    return s
})
// vals -> []string{"a!", "b", "c!", "d"}
```
