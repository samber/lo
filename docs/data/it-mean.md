---
name: Mean / MeanBy
slug: mean
sourceRef: it/math.go#L80
category: it
subCategory: math
signatures:
  - "func Mean[T constraints.Float | constraints.Integer](collection iter.Seq[T]) T"
  - "func MeanBy[T any, R constraints.Float | constraints.Integer](collection iter.Seq[T], iteratee func(item T) R) R"
playUrl: "https://go.dev/play/p/Lez0CsvVRl_l"
variantHelpers:
  - it#math#mean
  - it#math#meanby
similarHelpers:
  - core#slice#mean
  - core#slice#meanby
position: 30
---

Computes the arithmetic mean. `MeanBy` applies a transform before averaging. Returns 0 for empty sequences.

Examples:

```go
avg := it.Mean(iter.Seq[int](func(y func(int) bool){ _ = y(2); _ = y(3); _ = y(5) }))
// avg == 10/3 == 3 (int division)
```

```go
avg := it.MeanBy(iter.Seq[string](func(y func(string) bool){ _ = y("aa"); _ = y("bbb") }), func(s string) int { return len(s) })
// (2+3)/2 == 2
```


