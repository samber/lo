---
name: WithoutBy
slug: withoutby
sourceRef: it/intersect.go#L159
category: it
subCategory: intersect
signatures:
  - "func WithoutBy[T any, K comparable, I ~func(func(T) bool)](collection I, transform func(item T) K, exclude ...K) I"
playUrl: ""
variantHelpers:
  - it#intersect#withoutby
similarHelpers:
  - core#slice#withoutby
position: 700
---

Filters a sequence by excluding elements whose extracted keys match any in the exclude list.

Returns a sequence containing only the elements whose keys are not in the exclude list.
Will allocate a map large enough to hold all distinct excludes.

Examples:

```go
// Exclude people by specific ages
type Person struct {
    Name string
    Age  int
}
people := it.Slice([]Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
    {Name: "Diana", Age: 30},  // Same age as Alice
})

filtered := it.WithoutBy(people, func(p Person) int { return p.Age }, 30)
// filtered: sequence with Bob (age 25) and Charlie (age 35)

// Exclude strings by their length
words := it.Slice([]string{"hello", "world", "hi", "go", "bye"})
filtered := it.WithoutBy(words, func(s string) int { return len(s) }, 2)
// filtered: sequence with "hello" (5), "world" (5), "bye" (3)
// excludes "hi" and "go" (both length 2)

// Exclude items by first letter
items := it.Slice([]string{"apple", "apricot", "banana", "blueberry", "cherry"})
filtered := it.WithoutBy(items, func(s string) byte { return s[0] }, 'b')
// filtered: sequence with "apple", "apricot", "cherry"
// excludes "banana" and "blueberry" (both start with 'b')

// Exclude numbers by modulo
numbers := it.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
filtered := it.WithoutBy(numbers, func(n int) int { return n % 3 }, 1)
// filtered: sequence with numbers where n % 3 != 1
// excludes 1, 4, 7, 10 (all have remainder 1)

// Exclude emails by domain
type Email struct {
    Address string
}
emails := it.Slice([]Email{
    {Address: "user1@example.com"},
    {Address: "user2@gmail.com"},
    {Address: "user3@example.com"},
    {Address: "user4@yahoo.com"},
})
filtered := it.WithoutBy(emails, func(e Email) string {
    parts := strings.Split(e.Address, "@")
    if len(parts) > 1 {
        return parts[1]
    }
    return ""
}, "example.com")
// filtered: sequence with gmail.com and yahoo.com emails

// Exclude orders by customer ID
type Order struct {
    ID         string
    CustomerID string
    ProductID  string
}
orders := it.Slice([]Order{
    {ID: "1", CustomerID: "A", ProductID: "X"},
    {ID: "2", CustomerID: "B", ProductID: "Y"},
    {ID: "3", CustomerID: "A", ProductID: "Z"},
    {ID: "4", CustomerID: "C", ProductID: "W"},
})
filtered := it.WithoutBy(orders, func(o Order) string { return o.CustomerID }, "A")
// filtered: sequence with orders from customers B and C

// Exclude strings by case-insensitive value
words := it.Slice([]string{"Hello", "hello", "WORLD", "world", "Go"})
filtered := it.WithoutBy(words, func(s string) string { return strings.ToLower(s) }, "hello")
// filtered: sequence with "WORLD", "world", "Go"
// excludes both "Hello" and "hello" (both become "hello" when lowercased)

// Exclude dates by month
import "time"
dates := it.Slice([]time.Time{
    time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 2, 20, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 1, 25, 0, 0, 0, 0, time.UTC),
    time.Date(2023, 3, 10, 0, 0, 0, 0, time.UTC),
})
filtered := it.WithoutBy(dates, func(t time.Time) time.Month { return t.Month() }, time.January)
// filtered: sequence with February and March dates

// Exclude multiple values
numbers := it.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
filtered := it.WithoutBy(numbers, func(n int) int { return n % 2 }, 0, 1)
// filtered: empty sequence (all numbers are either even (0) or odd (1))

// Exclude by custom function result
type Product struct {
    Name  string
    Price float64
}
products := it.Slice([]Product{
    {Name: "Book", Price: 19.99},
    {Name: "Pen", Price: 1.99},
    {Name: "Laptop", Price: 999.99},
    {Name: "Pencil", Price: 0.99},
})
// Exclude products with price < $10
filtered := it.WithoutBy(products, func(p Product) string {
    if p.Price < 10 {
        return "cheap"
    }
    return "expensive"
}, "cheap")
// filtered: sequence with "Book" and "Laptop"
```