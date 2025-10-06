---
name: ZipByX
slug: zipbyx
sourceRef: it/tuples.go#L295
category: it
subCategory: tuple
signatures:
  - "func ZipBy2[T1, T2, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], transform func(T1, T2) R) iter.Seq[R]"
  - "func ZipBy3[T1, T2, T3, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], transform func(T1, T2, T3) R) iter.Seq[R]"
  - "func ZipBy4[T1, T2, T3, T4, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], transform func(T1, T2, T3, T4) R) iter.Seq[R]"
  - "func ZipBy5[T1, T2, T3, T4, T5, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], transform func(T1, T2, T3, T4, T5) R) iter.Seq[R]"
  - "func ZipBy6[T1, T2, T3, T4, T5, T6, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], seq6 iter.Seq[T6], transform func(T1, T2, T3, T4, T5, T6) R) iter.Seq[R]"
  - "func ZipBy7[T1, T2, T3, T4, T5, T6, T7, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], seq6 iter.Seq[T6], seq7 iter.Seq[T7], transform func(T1, T2, T3, T4, T5, T6, T7) R) iter.Seq[R]"
  - "func ZipBy8[T1, T2, T3, T4, T5, T6, T7, T8, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], seq6 iter.Seq[T6], seq7 iter.Seq[T7], seq8 iter.Seq[T8], transform func(T1, T2, T3, T4, T5, T6, T7, T8) R) iter.Seq[R]"
  - "func ZipBy9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], seq3 iter.Seq[T3], seq4 iter.Seq[T4], seq5 iter.Seq[T5], seq6 iter.Seq[T6], seq7 iter.Seq[T7], seq8 iter.Seq[T8], seq9 iter.Seq[T9], transform func(T1, T2, T3, T4, T5, T6, T7, T8, T9) R) iter.Seq[R]"
playUrl: ""
variantHelpers:
  - it#tuple#zipbyx
similarHelpers:
  - core#tuple#zipbyx
  - it#tuple#zipx
  - core#slice#map
position: 10
---

Creates a sequence of transformed elements from multiple sequences using a transform function. When sequences are different lengths, shorter sequences are padded with zero values before transformation.

Variants: `ZipBy2..ZipBy9`

```go
seq1 := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
seq2 := func(yield func(string) bool) {
    _ = yield("a")
    _ = yield("b")
}
zipped := it.ZipBy2(seq1, seq2, func(i int, s string) string {
    return fmt.Sprintf("%d-%s", i, s)
})
var result []string
for item := range zipped {
    result = append(result, item)
}
// result contains ["1-a", "2-b", "3-"]
```