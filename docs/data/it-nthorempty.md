---
name: NthOrEmpty
slug: nthorempty
sourceRef: it/find.go#L444
category: it
subCategory: find
signatures:
  - "func NthOrEmpty[T any, N constraints.Integer](collection iter.Seq[T], nth N) T"
playUrl: ""
variantHelpers:
  - it#find#nthorempty
similarHelpers:
  - core#slice#nthorempty
position: 600
---

Returns the element at index `nth` of collection. If `nth` is out of bounds, returns the zero value (empty value) for that type.

Will iterate n times through the sequence.

Examples:

```go
// Get element at specific index
numbers := it.Slice([]int{5, 2, 8, 1, 9})
element := it.NthOrEmpty(numbers, 2)
// element: 8

// Get first element (index 0)
first := it.NthOrEmpty(numbers, 0)
// first: 5

// Get last element
last := it.NthOrEmpty(numbers, 4)
// last: 9

// Out of bounds - negative, returns zero value
element := it.NthOrEmpty(numbers, -1)
// element: 0 (zero value for int)

// Out of bounds - too large, returns zero value
element := it.NthOrEmpty(numbers, 10)
// element: 0 (zero value for int)

// With strings
words := it.Slice([]string{"hello", "world", "go", "lang"})
element := it.NthOrEmpty(words, 1)
// element: "world"

// Out of bounds with string - returns empty string
element := it.NthOrEmpty(words, 10)
// element: "" (zero value for string)

// With structs
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
element := it.NthOrEmpty(people, 1)
// element: {Name: "Bob", Age: 25}

// Out of bounds with struct - returns zero value
element := it.NthOrEmpty(people, 5)
// element: {Name: "", Age: 0} (zero value for Person)

// With pointers - returns nil when out of bounds
values := it.Slice([]*string{ptr("hello"), ptr("world")})
element := it.NthOrEmpty(values, 1)
// element: pointer to "world"

// Out of bounds with pointer - returns nil
element := it.NthOrEmpty(values, 5)
// element: nil (zero value for *string)

// With different integer types
numbers := it.Slice([]int{1, 2, 3, 4, 5})
element := it.NthOrEmpty(numbers, int8(3))
// element: 4
```