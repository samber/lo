---
name: ContainsBy
slug: containsby
sourceRef: it/intersect.go#L20
category: iter
subCategory: intersect
signatures:
  - "func ContainsBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool"
playUrl: https://go.dev/play/p/m86Cpsoyv8k
variantHelpers:
  - iter#intersect#containsby
similarHelpers:
  - core#intersect#containsby
position: 650
---

Returns true if predicate function returns true for any element in the collection.

Will iterate through the entire sequence if predicate never returns true.

Examples:

```go
// Check if collection contains an even number
numbers := slices.Values([]int{1, 3, 5, 7, 9})
hasEven := it.ContainsBy(numbers, func(n int) bool { return n%2 == 0 })
// hasEven: false

numbers = slices.Values([]int{1, 3, 5, 8, 9})
hasEven = it.ContainsBy(numbers, func(n int) bool { return n%2 == 0 })
// hasEven: true

// Check if collection contains a string with specific prefix
words := slices.Values([]string{"hello", "world", "go", "lang"})
hasPrefix := it.ContainsBy(words, func(s string) bool { return strings.HasPrefix(s, "go") })
// hasPrefix: true

// Check if collection contains a person with specific age
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
hasAge30 := it.ContainsBy(people, func(p Person) bool { return p.Age == 30 })
// hasAge30: true

hasAge40 := it.ContainsBy(people, func(p Person) bool { return p.Age == 40 })
// hasAge40: false

// Check if collection contains an element with specific property
strs := slices.Values([]string{"apple", "banana", "cherry"})
hasLongString := it.ContainsBy(strs, func(s string) bool { return len(s) > 5 })
// hasLongString: true

// Check if collection contains negative numbers
numbers = slices.Values([]int{1, -2, 3, 4, -5})
hasNegative := it.ContainsBy(numbers, func(n int) bool { return n < 0 })
// hasNegative: true

// Check if collection contains valid email
emails := slices.Values([]string{"user@example.com", "invalid-email", "test@domain.org"})
hasValidEmail := it.ContainsBy(emails, func(email string) bool {
    return strings.Contains(email, "@") && strings.Contains(email, ".")
})
// hasValidEmail: true

// Check empty collection
empty := slices.Values([]int{})
hasAny := it.ContainsBy(empty, func(n int) bool { return n > 0 })
// hasAny: false

// Check for nil pointers (with pointer slice)
ptrs := slices.Values([]*int{lo.ToPtr(5), nil, lo.ToPtr(10)})
hasNil := it.ContainsBy(ptrs, func(p *int) bool { return p == nil })
// hasNil: true
```