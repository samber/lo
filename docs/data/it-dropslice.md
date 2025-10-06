---
name: DropLast
slug: dropslice
sourceRef: it/seq.go#L510
category: it
subCategory: slice
signatures:
  - "func DropLast[T any, I ~func(func(T) bool)](collection I, n int) I"
  - "func DropByIndex[T any, I ~func(func(T) bool)](collection I, indexes ...int) I"
  - "func Subset[T any, I ~func(func(T) bool)](collection I, offset, length int) I"
  - "func Slice[T any, I ~func(func(T) bool)](collection I, start, end int) I"
variantHelpers:
  - it#slice#droplast
  - it#slice#dropbyindex
  - it#slice#subset
  - it#slice#slice
similarHelpers:
  - core#slice#drop
  - core#slice#droplast
  - core#slice#dropbyindex
  - core#slice#subset
  - core#slice#slice
position: 150
---

DropLast drops n elements from the end of a sequence.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}

filtered := it.DropLast(collection, 2)
var result []int
for item := range filtered {
    result = append(result, item)
}
// result contains [1, 2, 3]
```

DropByIndex drops elements from a sequence by the index.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}

filtered := it.DropByIndex(collection, 1, 3)
var result []int
for item := range filtered {
    result = append(result, item)
}
// result contains [1, 3, 5]
```

Subset returns a subset of a sequence from `offset` up to `length` elements.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
    yield(6)
}

subset := it.Subset(collection, 2, 3)
var result []int
for item := range subset {
    result = append(result, item)
}
// result contains [3, 4, 5]
```

Slice returns a subset of a sequence from `start` up to, but not including `end`.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}

sliced := it.Slice(collection, 1, 4)
var result []int
for item := range sliced {
    result = append(result, item)
}
// result contains [2, 3, 4]
```