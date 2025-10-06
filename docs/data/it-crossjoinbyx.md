---
name: CrossJoinByX
slug: crossjoinbyx
sourceRef: it/tuples.go#L439
category: it
subCategory: tuple
signatures:
  - "func CrossJoinBy2[T1, T2, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], project func(T1, T2) R) iter.Seq[R]"
  - "func CrossJoinBy3[T1, T2, T3, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], project func(T1, T2, T3) R) iter.Seq[R]"
  - "func CrossJoinBy4[T1, T2, T3, T4, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], project func(T1, T2, T3, T4) R) iter.Seq[R]"
  - "func CrossJoinBy5[T1, T2, T3, T4, T5, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], project func(T1, T2, T3, T4, T5) R) iter.Seq[R]"
  - "func CrossJoinBy6[T1, T2, T3, T4, T5, T6, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], seq6 iter.Seq[T6], project func(T1, T2, T3, T4, T5, T6) R) iter.Seq[R]"
  - "func CrossJoinBy7[T1, T2, T3, T4, T5, T6, T7, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], seq6 iter.Seq[T6], seq7 iter.Seq[T7], project func(T1, T2, T3, T4, T5, T6, T7) R) iter.Seq[R]"
  - "func CrossJoinBy8[T1, T2, T3, T4, T5, T6, T7, T8, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], seq6 iter.Seq[T6], seq7 iter.Seq[T7], seq8 iter.Seq[T8], project func(T1, T2, T3, T4, T5, T6, T7, T8) R) iter.Seq[R]"
  - "func CrossJoinBy9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], seq6 iter.Seq[T6], seq7 iter.Seq[T7], seq8 iter.Seq[T8], seq9 iter.Seq[T9], project func(T1, T2, T3, T4, T5, T6, T7, T8, T9) R) iter.Seq[R]"
playUrl: ""
variantHelpers:
  - it#tuple#crossjoinbyx
similarHelpers:
  - core#tuple#crossjoinbyx
  - it#tuple#crossjoinx
  - core#slice#map
position: 20
---

Combines every item from multiple lists (cartesian product) using a project function to transform the results. Returns an empty sequence if any input sequence is empty.

Variants: `CrossJoinBy2..CrossJoinBy9`

```go
seq1 := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
}
seq2 := func(yield func(string) bool) {
    _ = yield("a")
    _ = yield("b")
}
result := it.CrossJoinBy2(seq1, seq2, func(i int, s string) string {
    return fmt.Sprintf("%d-%s", i, s)
})
var output []string
for item := range result {
    output = append(output, item)
}
// output contains ["1-a", "1-b", "2-a", "2-b"]
```