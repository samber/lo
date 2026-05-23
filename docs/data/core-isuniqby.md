---
name: IsUniqBy
slug: isuniqby
sourceRef: slice.go#L307
category: core
subCategory: slice
variantHelpers:
  - core#slice#isuniqby
similarHelpers:
  - core#slice#isuniq
  - core#slice#uniqby
position: 130
signatures:
  - "func IsUniqBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) bool"
---

Returns true if all elements in the slice are unique based on a custom criterion. The `iteratee` is invoked for each element to generate the key by which uniqueness is computed. Returns true for nil and empty slices.

```go
type User struct {
    ID   int
    Name string
}

lo.IsUniqBy([]User{{ID: 1}, {ID: 2}, {ID: 3}}, func(u User) int {
    return u.ID
})
// true

lo.IsUniqBy([]User{{ID: 1}, {ID: 2}, {ID: 1}}, func(u User) int {
    return u.ID
})
// false
```


