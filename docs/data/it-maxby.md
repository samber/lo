---
name: MaxBy
slug: maxby
sourceRef: it/find.go#L310
category: it
subCategory: find
signatures:
  - "func MaxBy[T any](collection iter.Seq[T], comparison func(a, b T) bool) T"
playUrl: "https://go.dev/play/p/1IcVZ5-zh"
variantHelpers:
  - it#find#maxby
similarHelpers:
  - core#slice#maxby
  - it#find#minby
  - it#find#max
position: 150
---

Searches maximum value using a custom comparison function. The comparison function should return true if the first argument is "greater than" the second.

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
oldest := it.MaxBy(seq, func(a, b Person) bool {
    return a.Age > b.Age
})
// oldest == Person{"Charlie", 35}
```