---
name: Empty
slug: sequencestate
sourceRef: it/type_manipulation.go#L43
category: it
subCategory: type
signatures:
  - "func Empty[T any]() iter.Seq[T]"
  - "func IsEmpty[T any](collection iter.Seq[T]) bool"
  - "func IsNotEmpty[T any](collection iter.Seq[T]) bool"
  - "func CoalesceSeq[T any](v ...iter.Seq[T]) (iter.Seq[T], bool)"
  - "func CoalesceSeqOrEmpty[T any](v ...iter.Seq[T]) iter.Seq[T]"
variantHelpers:
  - it#type#empty
  - it#type#isempty
  - it#type#isnotempty
  - it#type#coalesceseq
  - it#type#coalesceseqorempty
similarHelpers:
  - core#type#isempty
  - core#type#isnotempty
  - core#condition#coalesce
position: 250
---

Empty returns an empty sequence.

```go
empty := it.Empty[int]()
var result []int
for item := range empty {
    result = append(result, item)
}
// result is empty []
```

IsEmpty returns true if the sequence is empty.

```go
emptySeq := it.Empty[int]()
notEmptySeq := func(yield func(int) bool) {
    yield(1)
}

fmt.Println(it.IsEmpty(emptySeq))    // true
fmt.Println(it.IsEmpty(notEmptySeq)) // false
```

IsNotEmpty returns true if the sequence is not empty.

```go
emptySeq := it.Empty[int]()
notEmptySeq := func(yield func(int) bool) {
    yield(1)
}

fmt.Println(it.IsNotEmpty(emptySeq))    // false
fmt.Println(it.IsNotEmpty(notEmptySeq)) // true
```

CoalesceSeq returns the first non-empty sequence.

```go
empty1 := it.Empty[int]()
empty2 := it.Empty[int]()
notEmpty := func(yield func(int) bool) {
    yield(1)
    yield(2)
}

result, found := it.CoalesceSeq(empty1, empty2, notEmpty)
// found is true, result is the notEmpty sequence

result2, found2 := it.CoalesceSeq(empty1, empty2)
// found2 is false, result2 is an empty sequence
```

CoalesceSeqOrEmpty returns the first non-empty sequence.

```go
empty1 := it.Empty[int]()
empty2 := it.Empty[int]()
notEmpty := func(yield func(int) bool) {
    yield(1)
    yield(2)
}

result := it.CoalesceSeqOrEmpty(empty1, empty2, notEmpty)
// result is the notEmpty sequence

result2 := it.CoalesceSeqOrEmpty(empty1, empty2)
// result2 is an empty sequence
```