---
name: ChannelDispatcher
slug: channeldispatcher
sourceRef: channel.go#L18
category: core
subCategory: channel
signatures:
  - "func ChannelDispatcher[T any](stream <-chan T, count, channelBufferCap int, strategy DispatchingStrategy[T]) []<-chan T"
similarHelpers:
  - core#channel#fanin
  - core#channel#fanout
  - it#channel#seqtochannel
  - it#channel#channeltoseq
position: 250
---

ChannelDispatcher distributes messages from a stream to multiple channels based on a strategy.

```go
stream := make(chan int, 100)
for i := 0; i < 100; i++ {
    stream <- i
}
close(stream)

channels := lo.ChannelDispatcher(stream, 3, 10, lo.DispatchingStrategyRoundRobin[int])
// Returns 3 channels with round-robin distribution
```