---
name: MinByErr
slug: minbyerr
sourceRef: find.go#L349
category: core
subCategory: find
variantHelpers:
  - core#find#minbyerr
similarHelpers:
  - core#find#minby
  - core#find#min
  - core#find#minindex
  - core#find#minindexby
  - core#find#max
  - core#find#maxby
  - core#find#maxindex
  - core#find#maxindexby
position: 161
signatures:
  - "func MinByErr[T any](collection []T, comparison func(a T, b T) (bool, error)) (T, error)"
---

Searches the minimum value of a collection using the given comparison function. Returns the first minimal value; zero value and nil error when empty.

If the comparison function returns an error, iteration stops and the error is returned.

```go
type Point struct{ X int }
min, err := lo.MinByErr([]Point{{1}, {5}, {3}}, func(a, b Point) (bool, error) {
    return a.X < b.X, nil
})
// {1}, <nil>
```

Example with error:

```go
type Point struct{ X int }
min, err := lo.MinByErr([]Point{{1}, {5}, {3}}, func(a, b Point) (bool, error) {
    if a.X == 5 {
        return false, fmt.Errorf("cannot compare with 5")
    }
    return a.X < b.X, nil
})
// {1}, error("cannot compare with 5")
```
