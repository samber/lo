---
name: FanOut
slug: fanout
sourceRef: channel.go#L287
playUrl: "https://go.dev/play/p/2LHxcjKX23L"
category: core
subCategory: channel
signatures:
  - "func FanOut[T any](count, channelsBufferCap int, upstream <-chan T) []<-chan T"
similarHelpers:
  - core#channel#fanin
  - core#channel#channeldispatcher
  - iter#channel#channelseq
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