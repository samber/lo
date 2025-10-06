---
name: Reject
slug: reject
sourceRef: it/seq.go#L586
category: it
subCategory: sequence
signatures:
  - "func Reject[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I"
  - "func RejectI[T any, I ~func(func(T) bool)](collection I, predicate func(item T, index int) bool) I"
  - "func RejectMap[T, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R]"
  - "func RejectMapI[T, R any](collection iter.Seq[T], callback func(item T, index int) (R, bool)) iter.Seq[R]"
variantHelpers:
  - it#sequence#reject
  - it#sequence#rejecti
  - it#sequence#rejectmap
  - it#sequence#rejectmapi
similarHelpers:
  - core#slice#reject
  - core#slice#filter
  - it#sequence#filter
  - it#sequence#filtermap
position: 180
---

Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return true for.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}

filtered := it.Reject(collection, func(x int) bool {
    return x%2 == 0
})
var result []int
for item := range filtered {
    result = append(result, item)
}
// result contains [1, 3]
```

RejectI is the opposite of Filter, this method returns the elements of collection that predicate does not return true for, with index.

```go
collection := func(yield func(string) bool) {
    yield("a")
    yield("b")
    yield("c")
}

filtered := it.RejectI(collection, func(item string, index int) bool {
    return index == 1
})
var result []string
for item := range filtered {
    result = append(result, item)
}
// result contains ["a", "c"]
```

RejectMap returns a sequence obtained after both filtering and mapping using the given callback function.
The callback function should return two values: the result of the mapping operation and whether the result element should be included or not.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
}

filtered := it.RejectMap(collection, func(x int) (string, bool) {
    return fmt.Sprintf("item-%d", x), x%2 == 0
})
var result []string
for item := range filtered {
    result = append(result, item)
}
// result contains ["item-1", "item-3"]
```