---
name: IsUniq
slug: isuniq
sourceRef: slice.go#L290
category: core
subCategory: slice
playUrl:
variantHelpers:
  - core#slice#isuniq
  - core#slice#isuniqby
similarHelpers:
  - core#slice#uniq
  - core#slice#uniqby
  - core#slice#finduniques
  - core#find#findduplicates
position: 112
signatures:
  - "func IsUniq[T comparable, Slice ~[]T](collection Slice) bool"
  - "func IsUniqBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) bool"
---

Checks whether all elements in a slice are unique. Returns `true` for nil and empty slices.

### IsUniq

```go
lo.IsUniq([]int{1, 2, 3})
// true

lo.IsUniq([]int{1, 2, 1})
// false
```

### IsUniqBy

Checks uniqueness based on a key computed by the iteratee function.

```go
type User struct {
    ID  int
    Name string
}

users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}, {ID: 1, Name: "Charlie"}}

lo.IsUniqBy(users, func(u User) int { return u.ID })
// false

lo.IsUniqBy(users, func(u User) string { return u.Name })
// true
```
