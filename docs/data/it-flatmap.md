---
name: FlatMap
slug: flatmap
sourceRef: it/seq.go#L109
category: it
subCategory: sequence
signatures:
  - "func FlatMap[T, R any](collection iter.Seq[T], transform func(item T) iter.Seq[R]) iter.Seq[R]"
  - "func FlatMapI[T, R any](collection iter.Seq[T], transform func(item T, index int) iter.Seq[R]) iter.Seq[R]"
playUrl: https://go.dev/play/p/1YsRLPl-wx
variantHelpers:
  - it#sequence#flatmap
  - it#sequence#flatmapi
similarHelpers:
  - core#slice#flatmap
  - it#sequence#map
position: 120
---

Transforms and flattens a sequence to another type. Each element is transformed into a sequence, then all sequences are concatenated.

Examples:

```go
seq := func(yield func([]int) bool) {
    _ = yield([]int{1, 2})
    _ = yield([]int{3, 4})
    _ = yield([]int{5})
}
flattened := it.FlatMap(seq, func(arr []int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for _, v := range arr {
            if !yield(v * 2) {
                return
            }
        }
    }
})
var result []int
for v := range flattened {
    result = append(result, v)
}
// result contains 2, 4, 6, 8, 10
```

### FlatMapI

Transforms and flattens a sequence to another type. Each element is transformed into a sequence with access to the element's index, then all sequences are concatenated.

```go
seq := func(yield func(string) bool) {
    _ = yield("a")
    _ = yield("b")
    _ = yield("c")
}
flattened := it.FlatMapI(seq, func(s string, index int) iter.Seq[string] {
    return func(yield func(string) bool) {
        for i := 0; i <= index; i++ {
            if !yield(fmt.Sprintf("%s-%d", s, i)) {
                return
            }
        }
    }
})
var result []string
for v := range flattened {
    result = append(result, v)
}
// result contains "a-0", "b-0", "b-1", "c-0", "c-1", "c-2"
```