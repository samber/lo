---
name: FirstOr
slug: firstor
sourceRef: it/find.go#L377
category: it
subCategory: find
signatures:
  - "func FirstOr[T any](collection iter.Seq[T], fallback T) T"
playUrl: "https://go.dev/play/p/7OiBF1-zn"
variantHelpers:
  - it#find#firstor
similarHelpers:
  - core#slice#firstor
position: 550
---

Returns the first element of a collection or the fallback value if empty.

Will iterate at most once.

Examples:

```go
// Get the first element or fallback value
numbers := it.Slice([]int{5, 2, 8, 1, 9})
first := it.FirstOr(numbers, 42)
// first: 5

// With empty collection
empty := it.Slice([]int{})
first := it.FirstOr(empty, 42)
// first: 42 (fallback value)

// With strings
words := it.Slice([]string{"hello", "world", "go"})
first := it.FirstOr(words, "fallback")
// first: "hello"

emptyWords := it.Slice([]string{})
first := it.FirstOr(emptyWords, "fallback")
// first: "fallback"

// With structs
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
first := it.FirstOr(people, Person{Name: "Default", Age: 0})
// first: {Name: "Alice", Age: 30}

emptyPeople := it.Slice([]Person{})
first := it.FirstOr(emptyPeople, Person{Name: "Default", Age: 0})
// first: {Name: "Default", Age: 0} (fallback value)

// Using with pointers
pointers := it.Slice([]*int{ptr(5), ptr(10), ptr(15)})
first := it.FirstOr(pointers, nil)
// first: pointer to 5

emptyPointers := it.Slice([]*int{})
first := it.FirstOr(emptyPointers, nil)
// first: nil (fallback value)
```