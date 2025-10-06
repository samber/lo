---
name: ChannelToSlice
slug: channeltoslice
sourceRef: channel.go#L18
category: core
subCategory: channel
signatures:
  - "func ChannelToSlice[T any](ch <-chan T) []T"
similarHelpers:
  - core#channel#slicetochannel
  - core#channel#buffer
  - it#channel#seqtochannel
  - it#channel#channeltoseq
position: 252
---

ChannelToSlice converts a channel to a slice.

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

slice := lo.ChannelToSlice(ch)
// slice contains [1, 2, 3]
```