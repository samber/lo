---
name: MaxByErr
slug: maxbyerr
sourceRef: find.go#L528
category: core
subCategory: find
playUrl: https://go.dev/play/p/s-63-6_9zqM
variantHelpers:
  - core#find#maxbyerr
similarHelpers:
  - core#find#maxby
  - core#find#max
  - core#find#maxindex
  - core#find#maxindexby
  - core#find#min
  - core#find#minby
  - core#find#minindex
  - core#find#minindexby
position: 221
signatures:
  - "func MaxByErr[T any](collection []T, comparison func(a T, b T) (bool, error)) (T, error)"
---

Searches the maximum value of a collection using the given comparison function. Returns zero value and nil error when empty.

If the comparison function returns an error, iteration stops and the error is returned.

```go
type Point struct{ X int }
max, err := lo.MaxByErr([]Point{{1}, {5}, {3}}, func(a, b Point) (bool, error) {
    return a.X > b.X, nil
})
// {5}, <nil>
```

Example with error:

```go
type Point struct{ X int }
max, err := lo.MaxByErr([]Point{{1}, {5}, {3}}, func(a, b Point) (bool, error) {
    if a.X == 5 {
        return false, fmt.Errorf("cannot compare with 5")
    }
    return a.X > b.X, nil
})
// {1}, error("cannot compare with 5")
```

Note: the comparison function is inconsistent with most languages, since we use the opposite of the usual convention.

See https://github.com/samber/lo/issues/129
