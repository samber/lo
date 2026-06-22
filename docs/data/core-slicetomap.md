---
name: SliceToMap
slug: slicetomap
sourceRef: slice.go#L405
category: core
subCategory: slice
playUrl: https://go.dev/play/p/WHa2CfMO3Lr
variantHelpers:
  - core#slice#slicetomap
  - core#slice#associate
  - core#slice#associatei
similarHelpers:
  - core#slice#filterslicetomap
  - core#slice#keyby
  - core#slice#groupby
  - core#slice#groupbymap
position: 250
signatures:
  - "func SliceToMap[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V"
  - "func Associate[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V"
  - "func AssociateI[T any, K comparable, V any](collection []T, transform func(item T, index int) (K, V)) map[K]V"
---

Alias of Associate: transforms a slice into a map using a key/value transform function.

```go
type foo struct {
    baz string
    bar int
}

in := []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}}

m := lo.SliceToMap(in, func(f *foo) (string, int) {
    return f.baz, f.bar
})
// map[string]int{"apple": 1, "banana": 2}
```


