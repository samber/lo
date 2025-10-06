---
name: EveryBy
slug: everyby
sourceRef: it/intersect.go#L43
category: it
subCategory: intersect
signatures:
  - "func EveryBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool"
playUrl: ""
variantHelpers:
  - it#intersect#everyby
similarHelpers:
  - core#slice#everyby
position: 660
---

Returns true if the predicate returns true for all elements in the collection or if the collection is empty.

Will iterate through the entire sequence if predicate never returns false.

Examples:

```go
// Check if all numbers are positive
numbers := it.Slice([]int{1, 3, 5, 7, 9})
allPositive := it.EveryBy(numbers, func(n int) bool { return n > 0 })
// allPositive: true

numbers = it.Slice([]int{1, -3, 5, 7, 9})
allPositive = it.EveryBy(numbers, func(n int) bool { return n > 0 })
// allPositive: false

// Check if all strings have minimum length
words := it.Slice([]string{"hello", "world", "go", "lang"})
allLongEnough := it.EveryBy(words, func(s string) bool { return len(s) >= 2 })
// allLongEnough: true

allVeryLong := it.EveryBy(words, func(s string) bool { return len(s) >= 5 })
// allVeryLong: false

// Check if all people are adults
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
allAdults := it.EveryBy(people, func(p Person) bool { return p.Age >= 18 })
// allAdults: true

minors := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 15},  // Not adult
    {Name: "Charlie", Age: 35},
})
allAdults = it.EveryBy(minors, func(p Person) bool { return p.Age >= 18 })
// allAdults: false

// Check if all numbers are even
numbers = it.Slice([]int{2, 4, 6, 8, 10})
allEven := it.EveryBy(numbers, func(n int) bool { return n%2 == 0 })
// allEven: true

numbers = it.Slice([]int{2, 4, 6, 7, 10})  // 7 is odd
allEven = it.EveryBy(numbers, func(n int) bool { return n%2 == 0 })
// allEven: false

// Check if all strings are lowercase
strings := it.Slice([]string{"hello", "world", "go", "lang"})
allLowercase := it.EveryBy(strings, func(s string) bool { return s == strings.ToLower(s) })
// allLowercase: true

strings = it.Slice([]string{"hello", "World", "go", "lang"})  // "World" has uppercase
allLowercase = it.EveryBy(strings, func(s string) bool { return s == strings.ToLower(s) })
// allLowercase: false

// Empty collection returns true
empty := it.Slice([]int{})
allPositive := it.EveryBy(empty, func(n int) bool { return n > 0 })
// allPositive: true

// Check if all emails are valid
emails := it.Slice([]string{"user@example.com", "test@domain.org", "admin@site.net"})
allValid := it.EveryBy(emails, func(email string) bool {
    return strings.Contains(email, "@") && strings.Contains(email, ".")
})
// allValid: true

emails = it.Slice([]string{"user@example.com", "invalid-email", "test@domain.org"})
allValid = it.EveryBy(emails, func(email string) bool {
    return strings.Contains(email, "@") && strings.Contains(email, ".")
})
// allValid: false
```