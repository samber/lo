---
name: ToAnySlice
slug: toanyslice
sourceRef: type_manipulation.go#L111
category: core
subCategory: type
signatures:
  - "func ToAnySlice[T any](collection []T) []any"
variantHelpers:
  - core#type#toanyslice
similarHelpers:
  - core#type#fromanyslice
  - core#type#tosliceptr
  - core#type#fromsliceptr
position: 126
---

Converts a typed slice to a slice of empty interface values (any). This is useful when working with heterogeneous data.

```go
ints := []int{1, 2, 3}
result := lo.ToAnySlice(ints)
// []any{1, 2, 3}

strings := []string{"a", "b", "c"}
result = lo.ToAnySlice(strings)
// []any{"a", "b", "c"}

custom := []struct{ Name string }{{Name: "Alice"}, {Name: "Bob"}}
result = lo.ToAnySlice(custom)
// []any{struct{ Name string }{Name: "Alice"}, struct{ Name string }{Name: "Bob"}}
```