---
name: Associate
slug: associate
sourceRef: slice.go#L389
category: core
subCategory: slice
playUrl: https://go.dev/play/p/WHa2CfMO3Lr
variantHelpers:
  - core#slice#associate
  - core#slice#associatei
  - core#slice#slicetomap
similarHelpers:
  - core#slice#keyby
  - core#slice#groupby
  - core#slice#filterslicetomap
  - core#slice#slicetomap
position: 240
signatures:
  - "func Associate[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V"
  - "func AssociateI[T any, K comparable, V any](collection []T, transform func(item T, index int) (K, V)) map[K]V"
  - "func SliceToMap[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V"
---

Builds a map from a slice using a transform that yields key/value for each item.

```go
type foo struct {
    baz string
    bar int
}

in := []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}}

m := lo.Associate(in, func(f *foo) (string, int) {
    return f.baz, f.bar
})
// map[string]int{"apple": 1, "banana": 2}
```

### With index

```go
type User struct {
    Name string
    Age  int
}

users := []User{
    {Name: "Alice", Age: 25},
    {Name: "Bob", Age: 30},
}

result := lo.AssociateI(users, func(user User, index int) (string, int) {
    return fmt.Sprintf("%s-%d", user.Name, index), user.Age
})
// map[string]int{"Alice-0": 25, "Bob-1": 30}
```
