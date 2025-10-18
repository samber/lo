---
name: DropLastWhile
slug: droplastwhile
sourceRef: it/seq.go#L180
category: it
subCategory: sequence
signatures:
  - "func DropLastWhile[T any, I ~func(func(T) bool)](collection I, predicate func(item T) bool) I"
variantHelpers: []
similarHelpers:
  - core#slice#droplastwhile
position: 163
---

DropLastWhile drops elements from the end of a sequence while the predicate returns true.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}

filtered := it.DropLastWhile(collection, func(x int) bool {
    return x > 3
})
var result []int
for item := range filtered {
    result = append(result, item)
}
// result contains [1, 2, 3]
```