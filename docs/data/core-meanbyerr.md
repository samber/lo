---
name: MeanByErr
slug: meanbyerr
sourceRef: math.go#L172
category: core
subCategory: math
variantHelpers:
  - core#math#meanbyerr
similarHelpers:
  - core#math#meanby
  - core#math#mean
  - core#math#mode
  - core#math#sum
  - core#math#sumbyerr
  - core#math#product
  - core#math#productby
  - core#find#min
  - core#find#max
  - core#find#minby
  - core#find#maxby
position: 91
signatures:
  - "func MeanByErr[T any, R constraints.Float | constraints.Integer](collection []T, iteratee func(item T) (R, error)) (R, error)"
---

Calculates the mean of values computed by a predicate. Returns 0 for an empty collection.

If the iteratee returns an error, iteration stops and the error is returned.

```go
list := []string{"aa", "bbb", "cccc", "ddddd"}
result, err := lo.MeanByErr(list, func(item string) (float64, error) {
    return float64(len(item)), nil
})
// 3.5, <nil>
```

Example with error:

```go
list := []string{"aa", "bbb", "cccc", "ddddd"}
result, err := lo.MeanByErr(list, func(item string) (float64, error) {
    if item == "cccc" {
        return 0, fmt.Errorf("cccc is not allowed")
    }
    return float64(len(item)), nil
})
// 0, error("cccc is not allowed")
```
