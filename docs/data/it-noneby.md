---
name: NoneBy
slug: noneby
sourceRef: it/intersect.go#L69
category: it
subCategory: intersect
signatures:
  - "func NoneBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool"
playUrl: ""
variantHelpers:
  - it#intersect#noneby
similarHelpers:
  - core#slice#noneby
position: 690
---

Returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.

Will iterate through the entire sequence if predicate never returns true.

Examples:

```go
// Check if collection has no even numbers
numbers := it.Slice([]int{1, 3, 5, 7, 9})
hasNoEvens := it.NoneBy(numbers, func(n int) bool { return n%2 == 0 })
// hasNoEvens: true

numbers = it.Slice([]int{1, 3, 5, 8, 9})
hasNoEvens = it.NoneBy(numbers, func(n int) bool { return n%2 == 0 })
// hasNoEvens: false (8 is even)

// Check if collection has no strings with specific prefix
words := it.Slice([]string{"hello", "world", "go", "lang"})
hasNoGoPrefix := it.NoneBy(words, func(s string) bool { return strings.HasPrefix(s, "go") })
// hasNoGoPrefix: false ("go" has go prefix)

hasNoPythonPrefix := it.NoneBy(words, func(s string) bool { return strings.HasPrefix(s, "python") })
// hasNoPythonPrefix: true

// Check if collection has no minors
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
hasNoMinors := it.NoneBy(people, func(p Person) bool { return p.Age < 18 })
// hasNoMinors: true

withMinor := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 16},  // Minor
    {Name: "Charlie", Age: 35},
})
hasNoMinors = it.NoneBy(withMinor, func(p Person) bool { return p.Age < 18 })
// hasNoMinors: false

// Check if collection has no negative numbers
numbers = it.Slice([]int{1, 3, 5, 7, 9})
hasNoNegatives := it.NoneBy(numbers, func(n int) bool { return n < 0 })
// hasNoNegatives: true

numbers = it.Slice([]int{1, -3, 5, 7, 9})
hasNoNegatives = it.NoneBy(numbers, func(n int) bool { return n < 0 })
// hasNoNegatives: false (-3 is negative)

// Check if collection has no uppercase strings
strings := it.Slice([]string{"hello", "world", "go", "lang"})
hasNoUppercase := it.NoneBy(strings, func(s string) bool { return s != strings.ToLower(s) })
// hasNoUppercase: true

strings = it.Slice([]string{"hello", "World", "go", "lang"})  // "World" has uppercase
hasNoUppercase = it.NoneBy(strings, func(s string) bool { return s != strings.ToLower(s) })
// hasNoUppercase: false

// Empty collection returns true
empty := it.Slice([]int{})
hasNoEvens := it.NoneBy(empty, func(n int) bool { return n%2 == 0 })
// hasNoEvens: true

// Check if collection has no invalid emails
emails := it.Slice([]string{"user@example.com", "test@domain.org", "admin@site.net"})
hasNoInvalid := it.NoneBy(emails, func(email string) bool {
    return !strings.Contains(email, "@") || !strings.Contains(email, ".")
})
// hasNoInvalid: true

emails = it.Slice([]string{"user@example.com", "invalid-email", "test@domain.org"})
hasNoInvalid = it.NoneBy(emails, func(email string) bool {
    return !strings.Contains(email, "@") || !strings.Contains(email, ".")
})
// hasNoInvalid: false

// Check if collection has no numbers greater than 100
numbers = it.Slice([]int{1, 3, 5, 7, 9})
hasNoLargeNumbers := it.NoneBy(numbers, func(n int) bool { return n > 100 })
// hasNoLargeNumbers: true

numbers = it.Slice([]int{1, 3, 5, 150, 9})
hasNoLargeNumbers = it.NoneBy(numbers, func(n int) bool { return n > 100 })
// hasNoLargeNumbers: false (150 > 100)

// Check if collection has no strings shorter than 3 characters
words := it.Slice([]string{"hello", "world", "go", "lang"})
hasNoShortWords := it.NoneBy(words, func(s string) bool { return len(s) < 3 })
// hasNoShortWords: false ("go" has length 2)
```