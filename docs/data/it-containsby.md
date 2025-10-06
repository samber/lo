---
name: ContainsBy
slug: containsby
sourceRef: it/intersect.go#L17
category: it
subCategory: intersect
signatures:
  - "func ContainsBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool"
playUrl: ""
variantHelpers:
  - it#intersect#containsby
similarHelpers:
  - core#slice#containsby
position: 650
---

Returns true if predicate function returns true for any element in the collection.

Will iterate through the entire sequence if predicate never returns true.

Examples:

```go
// Check if collection contains an even number
numbers := it.Slice([]int{1, 3, 5, 7, 9})
hasEven := it.ContainsBy(numbers, func(n int) bool { return n%2 == 0 })
// hasEven: false

numbers = it.Slice([]int{1, 3, 5, 8, 9})
hasEven = it.ContainsBy(numbers, func(n int) bool { return n%2 == 0 })
// hasEven: true

// Check if collection contains a string with specific prefix
words := it.Slice([]string{"hello", "world", "go", "lang"})
hasPrefix := it.ContainsBy(words, func(s string) bool { return strings.HasPrefix(s, "go") })
// hasPrefix: true

// Check if collection contains a person with specific age
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
hasAge30 := it.ContainsBy(people, func(p Person) bool { return p.Age == 30 })
// hasAge30: true

hasAge40 := it.ContainsBy(people, func(p Person) bool { return p.Age == 40 })
// hasAge40: false

// Check if collection contains an element with specific property
strings := it.Slice([]string{"apple", "banana", "cherry"})
hasLongString := it.ContainsBy(strings, func(s string) bool { return len(s) > 5 })
// hasLongString: true

// Check if collection contains negative numbers
numbers = it.Slice([]int{1, -2, 3, 4, -5})
hasNegative := it.ContainsBy(numbers, func(n int) bool { return n < 0 })
// hasNegative: true

// Check if collection contains valid email
emails := it.Slice([]string{"user@example.com", "invalid-email", "test@domain.org"})
hasValidEmail := it.ContainsBy(emails, func(email string) bool {
    return strings.Contains(email, "@") && strings.Contains(email, ".")
})
// hasValidEmail: true

// Check empty collection
empty := it.Slice([]int{})
hasAny := it.ContainsBy(empty, func(n int) bool { return n > 0 })
// hasAny: false

// Check for nil pointers (with pointer slice)
ptrs := it.Slice([]*int{ptr(5), nil, ptr(10)})
hasNil := it.ContainsBy(ptrs, func(p *int) bool { return p == nil })
// hasNil: true
```