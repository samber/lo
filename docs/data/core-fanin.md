---
name: FanIn
slug: fanin
sourceRef: channel.go#L18
category: core
subCategory: channel
signatures:
  - "func FanIn[T any](channelBufferCap int, upstreams ...<-chan T) <-chan T"
similarHelpers:
  - core#channel#fanout
  - core#channel#channelmerge
  - core#channel#channeldispatcher
position: 254
---

FanIn merges multiple upstream channels into a single downstream channel, reading from all input channels and forwarding their values to one output channel. The channelBufferCap parameter sets the buffer size of the output channel.

```go
ch1 := make(chan int, 2)
ch2 := make(chan int, 2)
ch1 <- 1
ch1 <- 2
ch2 <- 3
ch2 <- 4
close(ch1)
close(ch2)

merged := lo.FanIn(10, ch1, ch2)
var result []int
for item := range merged {
    result = append(result, item)
}
// result: []int{1, 2, 3, 4} (order may vary as goroutine scheduling differs)
```