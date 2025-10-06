---
name: CrossJoinX
slug: crossjoinx
sourceRef: it/tuples.go#L370
category: it
subCategory: tuple
signatures:
  - "func CrossJoin2[T1, T2 any](list1 iter.Seq[T1], list2 iter.Seq[T2]) iter.Seq[lo.Tuple2[T1, T2]]"
  - "func CrossJoin3[T1, T2, T3 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3]) iter.Seq[lo.Tuple3[T1, T2, T3]]"
  - "func CrossJoin4[T1, T2, T3, T4 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4]) iter.Seq[lo.Tuple4[T1, T2, T3, T4]]"
  - "func CrossJoin5[T1, T2, T3, T4, T5 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5]) iter.Seq[lo.Tuple5[T1, T2, T3, T4, T5]]"
  - "func CrossJoin6[T1, T2, T3, T4, T5, T6 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5], list6 iter.Seq[T6]) iter.Seq[lo.Tuple6[T1, T2, T3, T4, T5, T6]]"
  - "func CrossJoin7[T1, T2, T3, T4, T5, T6, T7 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5], list6 iter.Seq[T6], list7 iter.Seq[T7]) iter.Seq[lo.Tuple7[T1, T2, T3, T4, T5, T6, T7]]"
  - "func CrossJoin8[T1, T2, T3, T4, T5, T6, T7, T8 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5], list6 iter.Seq[T6], list7 iter.Seq[T7], list8 iter.Seq[T8]) iter.Seq[lo.Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]]"
  - "func CrossJoin9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](list1 iter.Seq[T1], list2 iter.Seq[T2], list3 iter.Seq[T3], list4 iter.Seq[T4], list5 iter.Seq[T5], list6 iter.Seq[T6], list7 iter.Seq[T7], list8 iter.Seq[T8], list9 iter.Seq[T9]) iter.Seq[lo.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]"
playUrl: ""
variantHelpers:
  - it#tuple#crossjoinx
similarHelpers:
  - core#tuple#crossjoinx
  - it#tuple#zipx
  - it#tuple#crossjoinbyx
position: 10
---

Combines every item from multiple lists (cartesian product). The resulting sequence contains all possible combinations of elements from each input sequence.

The cartesian product means every element from the first sequence is paired with every element from the second sequence, then those pairs are combined with every element from the third sequence, and so on.

```go
seq1 := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
}
seq2 := func(yield func(string) bool) {
    _ = yield("a")
    _ = yield("b")
}
cross := it.CrossJoin2(seq1, seq2)
var result []string
for tuple := range cross {
    result = append(result, fmt.Sprintf("%d%s", tuple.A, tuple.B))
}
// result contains ["1a", "1b", "2a", "2b"] (all 4 combinations in lexical order)
```

Example with 3 sequences:
```go
numbers := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
}
letters := func(yield func(string) bool) {
    _ = yield("a")
    _ = yield("b")
}
colors := func(yield func(string) bool) {
    _ = yield("red")
    _ = yield("blue")
}
cross := it.CrossJoin3(numbers, letters, colors)
var result []string
for tuple := range cross {
    result = append(result, fmt.Sprintf("%d%s%s", tuple.A, tuple.B, tuple.C))
}
// result contains 8 combinations: ["1ared", "1ablue", "1bred", "1bblue", "2ared", "2ablue", "2bred", "2bblue"]
```