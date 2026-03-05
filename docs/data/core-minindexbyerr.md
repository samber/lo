---
name: MinIndexByErr
slug: minindexbyerr
sourceRef: find.go#L404
category: core
subCategory: find
playUrl: https://go.dev/play/p/MUqi_NvTKM1
signatures:
  - "func MinIndexByErr[T any](collection []T, comparison func(a T, b T) (bool, error)) (T, int, error)"
variantHelpers:
  - core#find#minindexbyerr
similarHelpers:
  - core#find#minindexby
  - core#find#min
  - core#find#minby
  - core#find#minbyerr
  - core#find#minindex
  - core#find#max
  - core#find#maxby
  - core#find#maxbyerr
  - core#find#maxindex
  - core#find#maxindexby
  - core#find#maxindexbyerr
  - core#math#sum
  - core#math#sumbyerr
  - core#math#mean
  - core#math#meanbyerr
  - core#math#product
  - core#math#productbyerr
  - core#math#mode
position: 171
---

Searches the minimum value using a comparison function and returns the value and its index. Returns (zero value, -1, nil) when empty. Stops iteration immediately when an error is encountered.

```go
type Point struct{ X int }
value, idx, err := lo.MinIndexByErr([]Point{{1}, {5}, {3}}, func(a, b Point) (bool, error) {
    return a.X < b.X, nil
})
// value == {1}, idx == 0, err == nil
```

```go
// Error case - stops on first error
_, _, err := lo.MinIndexByErr([]Point{{1}, {5}, {0}}, func(a, b Point) (bool, error) {
    if a.X == 0 || b.X == 0 {
        return false, fmt.Errorf("zero value not allowed")
    }
    return a.X < b.X, nil
})
// error("zero value not allowed")
```

```go
// Error case on first comparison
_, _, err := lo.MinIndexByErr([]Point{{1}, {5}}, func(a, b Point) (bool, error) {
    return false, fmt.Errorf("comparison error")
})
// error("comparison error")
```

