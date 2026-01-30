---
name: Tee
slug: tee
sourceRef: channel.go#L369
category: core
subCategory: channel
signatures:
  - "func Tee[T any](count, channelsBufferCap int, upstream <-chan T) []<-chan T"
similarHelpers:
  - core#channel#fanout
  - core#channel#channeldispatcher
position: 257
---

Tee duplicates values into multiple channels. If an output channel is not ready, the value is dropped for that channel.

```go
upstream := lo.SliceToChannel(10, []int{0, 1, 2, 3})
downstreams := lo.Tee(2, 2, upstream)
// Each downstream receives a best-effort copy of the stream.
```
