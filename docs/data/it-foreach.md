---
name: ForEach
slug: foreach
sourceRef: it/seq.go#L166
category: iter
subCategory: sequence
signatures:
  - "func ForEach[T any](collection iter.Seq[T], callback func(item T))"
playUrl: "https://go.dev/play/p/agIsKpG-S-P"
variantHelpers:
  - iter#sequence#foreach
  - iter#sequence#foreachi
similarHelpers:
  - core#slice#foreach
  - iter#sequence#map
position: 40
---

Iterates over elements and invokes a callback function for each element.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
}
var result []int
it.ForEach(seq, func(item int) {
    result = append(result, item*2)
})
// result contains 2, 4, 6
```

```go
seq := func(yield func(string) bool) {
    _ = yield("hello")
    _ = yield("world")
}
it.ForEach(seq, func(item string) {
    fmt.Println("Item:", item)
})
// Prints: Item: hello
//        Item: world
```