---
name: ForEachWhile
slug: foreachwhile
sourceRef: it/seq.go#L180
category: it
subCategory: sequence
signatures:
  - "func ForEachWhile[T any](collection iter.Seq[T], predicate func(item T) bool)"
  - "func ForEachWhileI[T any](collection iter.Seq[T], predicate func(item T, index int) bool)"
variantHelpers:
  - it#sequence#foreachwhile
  - it#sequence#foreachwhilei
similarHelpers:
  - it#sequence#foreach
position: 160
---

ForEachWhile iterates over elements of collection and invokes predicate for each element.
The predicate return value decides to continue or break, like do while().

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    yield(4)
    yield(5)
}

called := 0
it.ForEachWhile(collection, func(item int) bool {
    called++
    return item < 3
})
// called is 3 (elements 1, 2, 3 were processed)
```

ForEachWhileI iterates over elements of collection and invokes predicate for each element with index.
The predicate return value decides to continue or break, like do while().

```go
collection := func(yield func(string) bool) {
    yield("a")
    yield("b")
    yield("c")
    yield("d")
}

called := 0
it.ForEachWhileI(collection, func(item string, index int) bool {
    called++
    return index < 2
})
// called is 3 (elements at indices 0, 1, 2 were processed)
```