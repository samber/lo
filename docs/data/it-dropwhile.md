---
name: DropWhile
slug: dropwhile
sourceRef: it/seq.go#L180
category: iter
subCategory: sequence
signatures:
  - "func DropWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I"
variantHelpers: []
similarHelpers:
  - core#slice#dropwhile
  - iter#sequence#drop
position: 162
---

DropWhile drops elements from the beginning of a sequence while the predicate returns true.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}

filtered := it.DropWhile(collection, func(x int) bool {
    return x < 3
})
var result []int
for item := range filtered {
    result = append(result, item)
}
// result contains [3, 4, 5]
```