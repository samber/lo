---
name: ZipX
slug: zipx
sourceRef: it/tuples.go#L14
category: it
subCategory: tuple
signatures:
  - "func Zip2[T1, T2 any](list1 iter.Seq[T1], list2 iter.Seq[T2]) iter.Seq[lo.Tuple2[T1, T2]]"
  - "func Zip3[T1, T2, T3 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3]) iter.Seq[lo.Tuple3[T1, T2, T3]]"
  - "func Zip4[T1, T2, T3, T4 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4]) iter.Seq[lo.Tuple4[T1, T2, T3, T4]]"
  - "func Zip5[T1, T2, T3, T4, T5 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5]) iter.Seq[lo.Tuple5[T1, T2, T3, T4, T5]]"
  - "func Zip6[T1, T2, T3, T4, T5, T6 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5], list6 iter.Seq[T6]) iter.Seq[lo.Tuple6[T1, T2, T3, T4, T5, T6]]"
  - "func Zip7[T1, T2, T3, T4, T5, T6, T7 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5], list6 iter.Seq[T6], list7 iter.Seq[T7]) iter.Seq[lo.Tuple7[T1, T2, T3, T4, T5, T6, T7]]"
  - "func Zip8[T1, T2, T3, T4, T5, T6, T7, T8 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5], list6 iter.Seq[T6], list7 iter.Seq[T7], list8 iter.Seq[T8]) iter.Seq[lo.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]"
  - "func Zip9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5], list6 iter.Seq[T6], list7 iter.Seq[T7], list8 iter.Seq[T8], list9 iter.Seq[T9]) iter.Seq[lo.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]"
playUrl: ""
variantHelpers:
  - it#tuple#zipx
similarHelpers:
  - core#tuple#zipx
  - it#tuple#crossjoinx
  - it#tuple#zipbyx
position: 0
---

Creates a sequence of grouped elements from multiple sequences. The resulting sequence has length equal to the shortest input sequence.

Variants: `Zip2..Zip9`

```go
seq1 := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
seq2 := func(yield func(string) bool) {
    _ = yield("a")
    _ = yield("b")
    _ = yield("c")
}
zipped := it.Zip2(seq1, seq2)
var result []string
for tuple := range zipped {
    result = append(result, fmt.Sprintf("%d%s", tuple.A, tuple.B))
}
// result contains "1a", "2b", "3c"
```