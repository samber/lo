---
name: FromPtr
slug: fromptr
sourceRef: type_manipulation.go#L57
category: core
subCategory: type
signatures:
  - "func FromPtr[T any](x *T) T"
variantHelpers:
  - core#type#fromptr
similarHelpers:
  - core#type#toptr
  - core#type#fromptror
  - core#type#emptyabletoptr
  - core#type#tosliceptr
  - core#type#fromsliceptr
position: 95
---

Dereferences a pointer and returns the underlying value. If the pointer is nil, returns the zero value for the type. This is a safe way to extract values from optional pointers without risking panics.

```go
ptr := lo.ToPtr(42)
value := lo.FromPtr(ptr)
// value: 42

value = lo.FromPtr[string](nil)
// value: "" (zero value for string)

value = lo.FromPtr[int](nil)
// value: 0 (zero value for int)

// Working with structs
type Person struct {
    Name string
    Age  int
}
var personPtr *Person
person := lo.FromPtr(personPtr)
// person: Person{Name: "", Age: 0} (zero value for Person)
```