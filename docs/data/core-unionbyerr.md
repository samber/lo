---
name: UnionByErr
slug: unionbyerr
sourceRef: intersect.go#L285
category: core
subCategory: intersect
playUrl:
variantHelpers:
  - core#intersect#unionby
  - core#intersect#unionbyerr
similarHelpers:
  - core#intersect#unionby
  - core#intersect#union
  - core#intersect#intersect
  - core#intersect#intersectby
  - core#slice#uniq
  - core#slice#uniqby
position: 111
signatures:
  - "func UnionByErr[T any, V comparable, Slice ~[]T](iteratee func(item T) (V, error), lists ...Slice) (Slice, error)"
---

Returns all distinct elements from multiple collections based on a key function that can return an error. The result maintains the relative order of first occurrences. Returns the first error encountered from the iteratee.

```go
lo.UnionByErr(func(i int) (int, error) { return i / 2, nil }, []int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
// []int{0, 2, 4, 10}, nil
```

```go
lo.UnionByErr(func(s string) (string, error) { return s[:1], nil }, []string{"foo", "bar"}, []string{"baz"})
// []string{"foo", "baz"}, nil
```

```go
lo.UnionByErr(func(i int) (int, error) {
    if i == 42 {
        return 0, errors.New("invalid value")
    }
    return i / 2, nil
}, []int{0, 1, 2}, []int{42})
// []int{0, 1}, error("invalid value")
```
