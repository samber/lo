---
name: Generator
slug: generator
sourceRef: channel.go#L18
category: core
subCategory: channel
signatures:
  - "func Generator[T any](bufferSize int, generator func(yield func(T))) <-chan T"
similarHelpers:
  - core#channel#slicetochannel
  - core#channel#fanin
  - it#channel#seqtochannel
position: 253
---

Generator creates a channel from a generator function.

```go
gen := lo.Generator(10, func(yield func(int)) {
    for i := 0; i < 10; i++ {
        yield(i * 2)
    }
})

for item := range gen {
    fmt.Println(item)
}
// Prints even numbers 0, 2, 4, 6, 8, 10, 12, 14, 16, 18
```