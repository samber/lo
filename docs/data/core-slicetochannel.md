---
name: SliceToChannel
slug: slicetochannel
sourceRef: channel.go#L18
category: core
subCategory: channel
signatures:
  - "func SliceToChannel[T any](bufferSize int, collection []T) <-chan T"
similarHelpers:
  - core#channel#channeltoslice
  - core#channel#buffer
  - it#channel#seqtochannel
  - it#channel#channeltoseq
position: 251
---

SliceToChannel converts a slice to a channel with specified buffer size.

```go
items := []int{1, 2, 3, 4, 5}
ch := lo.SliceToChannel(10, items)

for item := range ch {
    fmt.Println(item)
}
// Prints 1, 2, 3, 4, 5
```