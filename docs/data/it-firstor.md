---
name: FirstOr
slug: firstor
sourceRef: it/find.go#L404
category: iter
subCategory: find
signatures:
  - "func FirstOr[T any](collection iter.Seq[T], fallback T) T"
playUrl: https://go.dev/play/p/wGFXI5NHkE2
variantHelpers:
  - iter#find#firstor
similarHelpers:
  - core#find#firstor
position: 550
---

Returns the first element of a collection or the fallback value if empty.

Will iterate at most once.

Examples:

```go
// Get the first element or fallback value
numbers := slices.Values([]int{5, 2, 8, 1, 9})
first := it.FirstOr(numbers, 42)
// first: 5

// With empty collection
empty := slices.Values([]int{})
first = it.FirstOr(empty, 42)
// first: 42 (fallback value)

// With strings
words := slices.Values([]string{"hello", "world", "go"})
firstStr := it.FirstOr(words, "fallback")
// firstStr: "hello"

emptyWords := slices.Values([]string{})
firstStr = it.FirstOr(emptyWords, "fallback")
// firstStr: "fallback"

// With structs
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
})
firstPerson := it.FirstOr(people, Person{Name: "Default", Age: 0})
// firstPerson: {Name: "Alice", Age: 30}

emptyPeople := slices.Values([]Person{})
firstPerson = it.FirstOr(emptyPeople, Person{Name: "Default", Age: 0})
// firstPerson: {Name: "Default", Age: 0} (fallback value)

// Using with pointers
pointers := slices.Values([]*int{lo.ToPtr(5), lo.ToPtr(10), lo.ToPtr(15)})
firstPtr := it.FirstOr(pointers, nil)
// firstPtr: pointer to 5

emptyPointers := slices.Values([]*int{})
firstPtr = it.FirstOr(emptyPointers, nil)
// firstPtr: nil (fallback value)
```