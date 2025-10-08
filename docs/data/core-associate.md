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

Builds a map from a slice using a transform function that yields key/value pairs for each item. Perfect for converting collections to lookup maps.

### Associate

Transforms each element into a key-value pair. Later items with the same key will overwrite earlier ones.

```go
type foo struct {
    baz string
    bar int
}

in := []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}}

m := lo.Associate(in, func(f *foo) (string, int) {
    return f.baz, f.bar
})
// m: map[string]int{"apple": 1, "banana": 2}
```

### AssociateI

Variant that includes the element index in the transform function, useful when you need the position in the original slice.

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
// result: map[string]int{"Alice-0": 25, "Bob-1": 30}
```

### SliceToMap

Alias for Associate - provides the same functionality with a more explicit name.

```go
products := []string{"apple", "banana", "cherry"}

result := lo.SliceToMap(products, func(product string) (string, int) {
    return product, len(product)
})
// result: map[string]int{"apple": 5, "banana": 6, "cherry": 6}
```
