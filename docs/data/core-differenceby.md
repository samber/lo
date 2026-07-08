---
name: DifferenceBy
slug: differenceby
sourceRef: intersect.go#L235
category: core
subCategory: intersect
variantHelpers:
  - core#intersect#differenceby
similarHelpers:
  - core#intersect#difference
  - core#intersect#intersectby
  - core#intersect#withoutby
  - core#slice#uniqby
position: 95
signatures:
  - "func DifferenceBy[T any, R comparable, Slice ~[]T](left, right Slice, iteratee func(item T) R) (notInRight, notInLeft Slice)"
---

Returns the difference between two collections using a custom key selector function. The first slice contains elements from left whose keys are absent from right; the second contains elements from right whose keys are absent from left.

```go
type User struct {
  ID int
  Name string
}

left := []User{
  {ID: 1, Name: "Alice"},
  {ID: 2, Name: "Bob"},
  {ID: 3, Name: "Charlie"},
}

right := []User{
  {ID: 2, Name: "Robert"},
  {ID: 4, Name: "David"},
}

notInRight, notInLeft := lo.DifferenceBy(left, right, func(user User) int {
  return user.ID
})
// []User{{ID: 1, Name: "Alice"}, {ID: 3, Name: "Charlie"}}, []User{{ID: 4, Name: "David"}}
```
