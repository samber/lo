---
name: SumBy
slug: sumby
sourceRef: math.go#L80
category: core
subCategory: math
playUrl: https://go.dev/play/p/Dz_a_7jN_ca
variantHelpers:
  - core#math#sumby
similarHelpers:
  - core#math#sum
  - core#math#sumbyerr
  - core#math#productby
  - core#math#meanby
  - core#math#product
  - core#math#mean
position: 50
signatures:
  - "func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R"
---

Sums the values computed by a predicate across a collection. Returns 0 for an empty collection.

```go
strings := []string{"foo", "bar"}
lo.SumBy(strings, func(item string) int {
    return len(item)
})
// 6
```

```go
// See also: SumByErr for error handling
strings := []string{"foo", "bar", "baz"}
sum, err := lo.SumByErr(strings, func(item string) (int, error) {
    if item == "bar" {
        return 0, fmt.Errorf("invalid item: %s", item)
    }
    return len(item), nil
})
// sum: 3, err: invalid item: bar
```
