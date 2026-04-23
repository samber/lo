---
name: MinBy
slug: minby
sourceRef: it/find.go#L242
category: iter
subCategory: find
signatures:
  - "func MinBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) T"
playUrl: https://go.dev/play/p/J5koo8khN-g
variantHelpers:
  - iter#find#minby
similarHelpers:
  - core#slice#minby
  - iter#find#maxby
  - iter#find#min
position: 140
---

Searches minimum value using a custom comparison function. The comparison function should return true if the first argument is "less than" the second.

Examples:

```go
type Person struct {
    Name string
    Age  int
}
seq := func(yield func(Person) bool) {
    _ = yield(Person{"Alice", 30})
    _ = yield(Person{"Bob", 25})
    _ = yield(Person{"Charlie", 35})
}
youngest := it.MinBy(seq, func(a, b Person) bool {
    return a.Age < b.Age
})
// youngest == Person{"Bob", 25}
```