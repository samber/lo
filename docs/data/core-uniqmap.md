---
name: UniqMap
slug: uniqmap
sourceRef: slice.go#L38
category: core
subCategory: slice
playUrl: https://go.dev/play/p/fygzLBhvUdB
variantHelpers:
  - core#slice#uniqmap
similarHelpers:
  - core#slice#uniq
  - core#slice#uniqby
  - core#slice#map
  - core#slice#filtermap
position: 20
signatures:
  - "func UniqMap[T any, R comparable](collection []T, iteratee func(item T, index int) R) []R"
---

Manipulates a slice and transforms it to a slice of another type with unique values.

```go
type User struct {
    Name string
    Age  int
}

users := []User{{Name: "Alex", Age: 10}, {Name: "Alex", Age: 12}, {Name: "Bob", Age: 11}, {Name: "Alice", Age: 20}}

names := lo.UniqMap(users, func(u User, index int) string {
    return u.Name
})
// []string{"Alex", "Bob", "Alice"}
```


