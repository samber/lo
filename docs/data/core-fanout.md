---
name: FanOut
slug: fanout
sourceRef: channel.go#L18
category: core
subCategory: channel
signatures:
  - "func FanOut[T any](count, channelsBufferCap int, upstream <-chan T) []<-chan T"
similarHelpers:
  - core#channel#fanin
  - core#channel#channeldispatcher
  - it#channel#channelseq
  - core#channel#tee
position: 256
---

FanOut splits a single channel into multiple channels.

```go
upstream := make(chan int, 6)
for i := 0; i < 6; i++ {
    upstream <- i
}
close(upstream)

downstreams := lo.FanOut(3, 10, upstream)
// Returns 3 channels, each receiving 2 items
```