---
name: MaxIndexByErr
slug: maxindexbyerr
sourceRef: find.go#L591
category: core
subCategory: find
variantHelpers:
  - core#find#maxindexbyerr
similarHelpers:
  - core#find#maxindexby
  - core#find#max
  - core#find#maxby
  - core#find#maxindex
  - core#find#min
  - core#find#minby
  - core#find#minindex
  - core#find#minindexby
position: 231
signatures:
  - "func MaxIndexByErr[T any](collection []T, comparison func(a T, b T) (bool, error)) (T, int, error)"
---

Returns the maximum value and its index using the given comparison function. Returns (zero value, -1, nil) when empty.

If the comparison function returns an error, iteration stops and the error is returned.

```go
type Point struct{ X int }
value, idx, err := lo.MaxIndexByErr([]Point{{1}, {5}, {3}}, func(a, b Point) (bool, error) {
    return a.X > b.X, nil
})
// value == {5}, idx == 1, err == nil
```

Example with error:

```go
type Point struct{ X int }
value, idx, err := lo.MaxIndexByErr([]Point{{1}, {5}, {3}}, func(a, b Point) (bool, error) {
    if a.X == 5 {
        return false, fmt.Errorf("cannot compare with 5")
    }
    return a.X > b.X, nil
})
// value == {1}, idx == 0, error("cannot compare with 5")
```

Note: the comparison function is inconsistent with most languages, since we use the opposite of the usual convention.

See https://github.com/samber/lo/issues/129
