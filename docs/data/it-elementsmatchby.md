---
name: ElementsMatchBy
slug: elementsmatchby
sourceRef: it/intersect.go#L183
category: it
subCategory: intersect
signatures:
  - "func ElementsMatchBy[T any, K comparable](list1, list2 iter.Seq[T], transform func(item T) K) bool"
playUrl: ""
variantHelpers:
  - it#intersect#elementsmatchby
similarHelpers:
  - core#slice#elementsmatchby
position: 730
---

Returns true if lists contain the same set of elements' keys (including empty set).

If there are duplicate keys, the number of occurrences in each list should match.
The order of elements is not checked.
Will iterate through each sequence before returning and allocate a map large enough to hold all distinct transformed elements.
Long heterogeneous input sequences can cause excessive memory usage.

Examples:

```go
// Match people by age (ignoring names)
type Person struct {
    Name string
    Age  int
}
people1 := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
})
people2 := it.Slice([]Person{
    {Name: "David", Age: 35},
    {Name: "Eve", Age: 25},
    {Name: "Frank", Age: 30},
})
match := it.ElementsMatchBy(people1, people2, func(p Person) int { return p.Age })
// match: true (both have ages 25, 30, 35)

// Match by string length
words1 := it.Slice([]string{"hello", "world", "go", "lang"})
words2 := it.Slice([]string{"short", "longer", "golang", "python"})
match = it.ElementsMatchBy(words1, words2, func(s string) int { return len(s) })
// match: false (lengths: 5,5,2,4 vs 6,6,6,6)

// Match by first character
items1 := it.Slice([]string{"apple", "apricot", "banana", "blueberry"})
items2 := it.Slice([]string{"ant", "anchor", "boat", "berry"})
match = it.ElementsMatchBy(items1, items2, func(s string) byte { return s[0] })
// match: true (both start with a, a, b, b)

// Match emails by domain
type Email struct {
    Address string
}
emails1 := it.Slice([]Email{
    {Address: "user1@example.com"},
    {Address: "user2@gmail.com"},
    {Address: "user3@example.com"},
})
emails2 := it.Slice([]Email{
    {Address: "different@gmail.com"},
    {Address: "another@example.com"},
    {Address: "third@example.com"},
})
match = it.ElementsMatchBy(emails1, emails2, func(e Email) string {
    parts := strings.Split(e.Address, "@")
    if len(parts) > 1 {
        return parts[1]
    }
    return ""
})
// match: true (both have domains: example.com, gmail.com, example.com)

// Match by modulo operation
numbers1 := it.Slice([]int{1, 2, 3, 4, 5, 6})
numbers2 := it.Slice([]int{7, 8, 9, 10, 11, 12})
match = it.ElementsMatchBy(numbers1, numbers2, func(n int) int { return n % 3 })
// match: true (remainders: 1,2,0,1,2,0 vs 1,2,0,1,2,0)

// Match by case-insensitive strings
strings1 := it.Slice([]string{"Hello", "World", "GO"})
strings2 := it.Slice([]string{"hello", "world", "go"})
match = it.ElementsMatchBy(strings1, strings2, func(s string) string { return strings.ToLower(s) })
// match: true

// Match orders by customer ID (ignoring order details)
type Order struct {
    ID         string
    CustomerID string
    ProductID  string
}
orders1 := it.Slice([]Order{
    {ID: "1", CustomerID: "A", ProductID: "X"},
    {ID: "2", CustomerID: "B", ProductID: "Y"},
    {ID: "3", CustomerID: "A", ProductID: "Z"},
})
orders2 := it.Slice([]Order{
    {ID: "4", CustomerID: "B", ProductID: "W"},
    {ID: "5", CustomerID: "A", ProductID: "V"},
    {ID: "6", CustomerID: "A", ProductID: "U"},
})
match = it.ElementsMatchBy(orders1, orders2, func(o Order) string { return o.CustomerID })
// match: true (both have customer IDs: A, B, A)

// Match dates by year-month (ignoring day)
import "time"
dates1 := it.Slice([]time.Time{
    time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 2, 20, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 1, 25, 0, 0, 0, 0, time.UTC),
})
dates2 := it.Slice([]time.Time{
    time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 2, 10, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 1, 30, 0, 0, 0, 0, time.UTC),
})
match = it.ElementsMatchBy(dates1, dates2, func(t time.Time) string {
    return fmt.Sprintf("%d-%02d", t.Year(), t.Month())
})
// match: true (both have: 2023-01, 2023-02, 2023-01)

// Match by category function
type Product struct {
    Name  string
    Price float64
}
products1 := it.Slice([]Product{
    {Name: "Book", Price: 19.99},
    {Name: "Pen", Price: 1.99},
    {Name: "Laptop", Price: 999.99},
})
products2 := it.Slice([]Product{
    {Name: "Pencil", Price: 0.99},
    {Name: "Phone", Price: 699.99},
    {Name: "Desk", Price: 199.99},
})
match = it.ElementsMatchBy(products1, products2, func(p Product) string {
    if p.Price < 10 {
        return "cheap"
    } else if p.Price < 100 {
        return "medium"
    }
    return "expensive"
})
// match: true (both have: medium, cheap, expensive)

// Match by custom classification
type Student struct {
    Name string
    Age  int
}
students1 := it.Slice([]Student{
    {Name: "Alice", Age: 8},
    {Name: "Bob", Age: 15},
    {Name: "Charlie", Age: 20},
})
students2 := it.Slice([]Student{
    {Name: "Diana", Age: 25},
    {Name: "Eve", Age: 10},
    {Name: "Frank", Age: 12},
})
match = it.ElementsMatchBy(students1, students2, func(s Student) string {
    if s.Age < 12 {
        return "child"
    } else if s.Age < 18 {
        return "teen"
    }
    return "adult"
})
// match: false (students1: child, teen, adult vs students2: adult, child, child)
```