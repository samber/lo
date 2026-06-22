---
name: SumByErr
slug: sumbyerr
sourceRef: math.go#L102
category: core
subCategory: math
variantHelpers:
  - core#math#sumbyerr
similarHelpers:
  - core#math#sumby
  - core#math#sum
  - core#math#productby
  - core#math#meanby
position: 50
signatures:
  - "func SumByErr[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) (R, error)) (R, error)"
---

Sums the values computed by a predicate across a collection, stopping early and returning an error if the predicate returns one. Returns 0 for an empty collection.

```go
strings := []string{"foo", "bar", "baz"}
sum, err := lo.SumByErr(strings, func(item string) (int, error) {
    if item == "bar" {
        return 0, fmt.Errorf("invalid item: %s", item)
    }
    return len(item), nil
})
// sum: 3, err: invalid item: bar
```

```go
strings := []string{"foo", "bar"}
sum, err := lo.SumByErr(strings, func(item string) (int, error) {
    return len(item), nil
})
// sum: 6, err: nil
```

```go
strings := []string{}
sum, err := lo.SumByErr(strings, func(item string) (int, error) {
    return len(item), nil
})
// sum: 0, err: nil
```
