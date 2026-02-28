---
name: WithoutByErr
slug: withoutbyerr
sourceRef: intersect.go#L287
category: core
subCategory: intersect
signatures:
  - "func WithoutByErr[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) (K, error), exclude ...K) (Slice, error)"
variantHelpers:
  - core#intersect#withoutbyerr
similarHelpers:
  - core#intersect#withoutby
  - core#intersect#without
  - core#slice#rejectbyerr
position: 125
---

Filters a slice by excluding elements whose extracted keys match any in the exclude list. Returns an error if the iteratee function fails, stopping iteration immediately.

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

filtered, err := lo.WithoutByErr(users, func(u User) (int, error) {
  if u.ID == 2 {
    return 0, fmt.Errorf("Bob not allowed")
  }
  return u.ID, nil
}, 2, 3)
// []User(nil), error("Bob not allowed")
```

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

filtered, err := lo.WithoutByErr(users, func(u User) (int, error) {
  return u.ID, nil
}, 2, 3)
// []User{{1, "Alice"}}, nil
```
