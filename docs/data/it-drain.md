---
name: Drain
slug: drain
sourceRef: it/seq.go#L26
category: it
subCategory: sequence
signatures:
  - "func Drain[T any](collection iter.Seq[T])"
variantHelpers: []
similarHelpers: []
position: 170
---

Drain consumes an entire sequence.

```go
collection := func(yield func(int) bool) {
    yield(1)
    yield(2)
    yield(3)
    fmt.Println("yielding")
}

it.Drain(collection)
// prints "yielding" three times, sequence is consumed
```