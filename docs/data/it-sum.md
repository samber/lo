---
name: Sum / SumBy
slug: sum
sourceRef: it/math.go#L59
category: it
subCategory: math
signatures:
  - "func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T]) T"
  - "func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection iter.Seq[T], iteratee func(item T) R) R"
playUrl: ""
variantHelpers:
  - it#math#sum
  - it#math#sumby
similarHelpers:
  - core#slice#sum
  - core#slice#sumby
  - core#slice#product
  - core#slice#productby
position: 10
---

Sums values from a sequence. `SumBy` applies a transform and sums the results. Returns 0 for empty sequences.

Examples:

```go
sum := it.Sum(it.RangeFrom(1, 5))
// 1 + 2 + 3 + 4 + 5 == 15
```

```go
type User struct { Name string; Score int }
users := func(yield func(User) bool) {
    _ = yield(User{"a", 3})
    _ = yield(User{"b", 7})
}
total := it.SumBy(iter.Seq[User](users), func(u User) int { return u.Score })
// total == 10
```


