---
name: ChannelMerge
slug: channelmerge
sourceRef: channel.go#L18
category: core
subCategory: channel
signatures:
  - "func ChannelMerge[T any](channelBufferCap int, upstreams ...<-chan T) <-chan T"
similarHelpers:
  - core#channel#fanin
  - core#channel#fanout
  - core#channel#channeldispatcher
position: 255
---

ChannelMerge merges multiple channels into a single channel.

```go
ch1 := make(chan int, 2)
ch2 := make(chan int, 2)
ch1 <- 1
ch1 <- 2
ch2 <- 3
ch2 <- 4
close(ch1)
close(ch2)

merged := lo.ChannelMerge(10, ch1, ch2)
var result []int
for item := range merged {
    result = append(result, item)
}
// result contains [1, 2, 3, 4] (order may vary)
```