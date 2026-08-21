---
name: SomeBy
slug: someby
sourceRef: it/intersect.go#L82
category: iter
subCategory: intersect
signatures:
  - "func SomeBy[T any](collection iter.Seq[T], predicate func(item T) bool) bool"
playUrl: https://go.dev/play/p/PDT6dWCl7Md
variantHelpers:
  - iter#intersect#someby
similarHelpers:
  - core#intersect#someby
position: 670
---

Returns true if the predicate returns true for any of the elements in the collection.

If the collection is empty SomeBy returns false.
Will iterate through the entire sequence if predicate never returns true.

Examples:

```go
// Check if any number is even
numbers := slices.Values([]int{1, 3, 5, 7, 9})
hasEven := it.SomeBy(numbers, func(n int) bool { return n%2 == 0 })
// hasEven: false

numbers = slices.Values([]int{1, 3, 5, 8, 9})
hasEven = it.SomeBy(numbers, func(n int) bool { return n%2 == 0 })
// hasEven: true

// Check if any string starts with specific prefix
words := slices.Values([]string{"hello", "world", "go", "lang"})
hasGoPrefix := it.SomeBy(words, func(s string) bool { return strings.HasPrefix(s, "go") })
// hasGoPrefix: true

hasPythonPrefix := it.SomeBy(words, func(s string) bool { return strings.HasPrefix(s, "python") })
// hasPythonPrefix: false

// Check if any person is a teenager
type Person struct {
    Name string
    Age  int
}
people := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
hasTeenager := it.SomeBy(people, func(p Person) bool { return p.Age >= 13 && p.Age <= 19 })
// hasTeenager: false

teenagers := slices.Values([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 16},  // Teenager
    {Name: "Charlie", Age: 35},
})
hasTeenager = it.SomeBy(teenagers, func(p Person) bool { return p.Age >= 13 && p.Age <= 19 })
// hasTeenager: true

// Check if any number is greater than 100
numbers = slices.Values([]int{1, 3, 5, 7, 9})
hasLargeNumber := it.SomeBy(numbers, func(n int) bool { return n > 100 })
// hasLargeNumber: false

numbers = slices.Values([]int{1, 3, 5, 150, 9})
hasLargeNumber = it.SomeBy(numbers, func(n int) bool { return n > 100 })
// hasLargeNumber: true

// Check if any string contains a substring
strs := slices.Values([]string{"hello", "world", "go", "lang"})
hasWorld := it.SomeBy(strs, func(s string) bool { return strings.Contains(s, "world") })
// hasWorld: true

hasPython := it.SomeBy(strs, func(s string) bool { return strings.Contains(s, "python") })
// hasPython: false

// Empty collection returns false
empty := slices.Values([]int{})
hasAny := it.SomeBy(empty, func(n int) bool { return n > 0 })
// hasAny: false

// Check if any email is from specific domain
emails := slices.Values([]string{"user@example.com", "test@gmail.com", "admin@site.net"})
hasGmail := it.SomeBy(emails, func(email string) bool { return strings.HasSuffix(email, "@gmail.com") })
// hasGmail: true

hasYahoo := it.SomeBy(emails, func(email string) bool { return strings.HasSuffix(email, "@yahoo.com") })
// hasYahoo: false

// Check if any string is palindrome
reverseString := func(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
words = slices.Values([]string{"level", "hello", "world", "radar"})
hasPalindrome := it.SomeBy(words, func(s string) bool {
    return s == reverseString(s)
})
// hasPalindrome: true ("level" and "radar" are palindromes)
```