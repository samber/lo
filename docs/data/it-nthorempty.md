---
name: NthOrEmpty
slug: nthorempty
sourceRef: it/find.go#L484
category: iter
subCategory: find
signatures:
  - "func NthOrEmpty[T any, N constraints.Integer](collection iter.Seq[T], nth N) T"
playUrl: https://go.dev/play/p/pC0Zhu3EUhe
variantHelpers:
  - iter#find#nthorempty
similarHelpers:
  - core#find#nthorempty
position: 600
---

Returns the element at index `nth` of collection. If `nth` is out of bounds, returns the zero value (empty value) for that type.

Will iterate n times through the sequence.

Examples:

```go
// Get element at specific index
numbers := slices.Values([]int{5, 2, 8, 1, 9})
element := it.NthOrEmpty(numbers, 2)
// element: 8

// Get first element (index 0)
first := it.NthOrEmpty(numbers, 0)
// first: 5

// Get last element
last := it.NthOrEmpty(numbers, 4)
// last: 9

// Out of bounds - negative, returns zero value
element = it.NthOrEmpty(numbers, -1)
// element: 0 (zero value for int)

// Out of bounds - too large, returns zero value
element = it.NthOrEmpty(numbers, 10)
// element: 0 (zero value for int)

// With strings
words := slices.Values([]string{"hello", "world", "go", "lang"})
elementStr := it.NthOrEmpty(words, 1)
// elementStr: "world"

// Out of bounds with string - returns empty string
elementStr = it.NthOrEmpty(words, 10)
// elementStr: "" (zero value for string)

// With structs
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
elementPerson := it.NthOrEmpty(people, 1)
// elementPerson: {Name: "Bob", Age: 25}

// Out of bounds with struct - returns zero value
elementPerson = it.NthOrEmpty(people, 5)
// elementPerson: {Name: "", Age: 0} (zero value for Person)

// With pointers - returns nil when out of bounds
values := slices.Values([]*string{lo.ToPtr("hello"), lo.ToPtr("world")})
elementPtr := it.NthOrEmpty(values, 1)
// elementPtr: pointer to "world"

// Out of bounds with pointer - returns nil
elementPtr = it.NthOrEmpty(values, 5)
// elementPtr: nil (zero value for *string)

// With different integer types
numbers = slices.Values([]int{1, 2, 3, 4, 5})
element = it.NthOrEmpty(numbers, int8(3))
// element: 4
```