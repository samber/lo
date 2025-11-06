---
name: Times
slug: times
sourceRef: it/seq.go#L202
category: it
subCategory: sequence
signatures:
  - "func Times[T any](count int, callback func(index int) T) iter.Seq[T]"
playUrl: "https://go.dev/play/p/9QkDH3-zp"
variantHelpers:
  - it#sequence#times
similarHelpers:
  - core#slice#times
  - it#math#range
position: 70
---

Invokes a callback function n times and returns a sequence of the results.

Examples:

```go
seq := it.Times(5, func(index int) int {
    return index * 2
})
var result []int
for v := range seq {
    result = append(result, v)
}
// result contains 0, 2, 4, 6, 8
```

```go
seq := it.Times(3, func(index int) string {
    return fmt.Sprintf("item-%d", index+1)
})
var result []string
for v := range seq {
    result = append(result, v)
}
// result contains "item-1", "item-2", "item-3"
```