---
name: ProductByErr
slug: productbyerr
sourceRef: math.go#L134
category: core
subCategory: math
variantHelpers:
  - core#math#productbyerr
similarHelpers:
  - core#math#productby
  - core#math#product
  - core#math#sumbyerr
  - core#math#meanby
  - core#find#minby
  - core#find#maxby
position: 71
signatures:
  - "func ProductByErr[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) (R, error)) (R, error)"
---

Calculates the product of values computed by a predicate. Returns 1 for nil or empty collections.

If the iteratee returns an error, iteration stops and the error is returned.

```go
strings := []string{"foo", "bar"}
result, err := lo.ProductByErr(strings, func(item string) (int, error) {
    return len(item), nil
})
// 9, <nil>
```

Example with error:

```go
strings := []string{"foo", "bar", "baz"}
result, err := lo.ProductByErr(strings, func(item string) (int, error) {
    if item == "bar" {
        return 0, fmt.Errorf("bar is not allowed")
    }
    return len(item), nil
})
// 3, error("bar is not allowed")
```
