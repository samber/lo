---
name: FromAnySlice
slug: fromanyslice
sourceRef: type_manipulation.go#L118
category: core
subCategory: type
signatures:
  - "func FromAnySlice[T any](in []any) ([]T, bool)"
variantHelpers:
  - core#type#fromanyslice
similarHelpers:
  - core#type#toanyslice
  - core#type#tosliceptr
  - core#type#fromsliceptr
position: 128
---

Converts a slice of empty interface values back to a typed slice. Returns the converted slice and a boolean indicating success. All elements must be of the target type.

```go
data := []any{1, 2, 3}
result, ok := lo.FromAnySlice[int](data)
// []int{1, 2, 3}, true

data = []any{"a", "b", "c"}
result, ok = lo.FromAnySlice[string](data)
// []string{"a", "b", "c"}, true

data = []any{1, "b", 3} // mixed types
result, ok = lo.FromAnySlice[int](data)
// []int{}, false (conversion failed due to string element)

data = []any{} // empty slice
result, ok = lo.FromAnySlice[int](data)
// []int{}, true (empty slice always succeeds)
```