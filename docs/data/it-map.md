---
name: Map
slug: map
sourceRef: it/seq.go#L51
category: it
subCategory: sequence
signatures:
  - "func Map[T, R any](collection iter.Seq[T], transform func(item T) R) iter.Seq[R]"
  - "func MapI[T, R any](collection iter.Seq[T], transform func(item T, index int) R) iter.Seq[R]"
playUrl: ""
variantHelpers:
  - it#sequence#map
  - it#sequence#mapi
similarHelpers:
  - core#slice#map
  - it#sequence#filtermap
  - it#sequence#flatmap
position: 20
---

Transforms a sequence to another type by applying a transform function to each element.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
mapped := it.Map(seq, func(x int) string {
    return fmt.Sprintf("item-%d", x)
})
var result []string
for v := range mapped {
    result = append(result, v)
}
// result contains "item-1", "item-2", "item-3"
```

### MapI

Transforms a sequence to another type by applying a transform function to each element and its index.

```go
seq := func(yield func(int) bool) {
    _ = yield(10)
    _ = yield(20)
    _ = yield(30)
}
mapped := it.MapI(seq, func(x int, index int) string {
    return fmt.Sprintf("item-%d-%d", x, index)
})
var result []string
for v := range mapped {
    result = append(result, v)
}
// result contains "item-10-0", "item-20-1", "item-30-2"
```
