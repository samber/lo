---
name: NthOr
slug: nthor
sourceRef: it/find.go#L433
category: it
subCategory: find
signatures:
  - "func NthOr[T any, N constraints.Integer](collection iter.Seq[T], nth N, fallback T) T"
playUrl: "https://go.dev/play/p/2TnGK6-zs"
variantHelpers:
  - it#find#nthor
similarHelpers:
  - core#slice#nthor
position: 590
---

Returns the element at index `nth` of collection. If `nth` is out of bounds, returns the fallback value instead of an error.

Will iterate n times through the sequence.

Examples:

```go
// Get element at specific index
numbers := it.Slice([]int{5, 2, 8, 1, 9})
element := it.NthOr(numbers, 2, 42)
// element: 8

// Get first element (index 0)
first := it.NthOr(numbers, 0, 42)
// first: 5

// Get last element
last := it.NthOr(numbers, 4, 42)
// last: 9

// Out of bounds - negative, returns fallback
element := it.NthOr(numbers, -1, 42)
// element: 42 (fallback)

// Out of bounds - too large, returns fallback
element := it.NthOr(numbers, 10, 42)
// element: 42 (fallback)

// With strings
words := it.Slice([]string{"hello", "world", "go", "lang"})
element := it.NthOr(words, 1, "fallback")
// element: "world"

// Out of bounds with string fallback
element := it.NthOr(words, 10, "fallback")
// element: "fallback"

// With structs
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
fallback := Person{Name: "Default", Age: 0}
element := it.NthOr(people, 1, fallback)
// element: {Name: "Bob", Age: 25}

// Out of bounds with struct fallback
element := it.NthOr(people, 5, fallback)
// element: {Name: "Default", Age: 0}

// With different integer types
numbers := it.Slice([]int{1, 2, 3, 4, 5})
element := it.NthOr(numbers, int8(3), 99)
// element: 4
```