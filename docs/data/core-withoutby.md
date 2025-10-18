---
name: WithoutBy
slug: withoutby
sourceRef: intersect.go#L199
category: core
subCategory: intersect
playUrl: https://go.dev/play/p/VgWJOF01NbJ
variantHelpers:
  - core#intersect#withoutby
similarHelpers:
  - core#intersect#without
  - core#intersect#difference
  - core#slice#rejectby
position: 120
signatures:
  - "func WithoutBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K, exclude ...K) Slice"
---

Filters a slice by excluding elements whose extracted keys match any in the exclude list.

```go
type User struct {
  ID int
  Name string
}

users := []User{
  {1, "Alice"},
  {2, "Bob"},
  {3, "Charlie"},
}

filtered := lo.WithoutBy(users, func(u User) int {
    return u.ID
}, 2, 3)
// []User{{1, "Alice"}}
```


