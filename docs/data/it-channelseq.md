---
name: ChannelToSeq
slug: channeltoseq
sourceRef: it/channel.go#L39
category: it
subCategory: channel
signatures:
  - "func ChannelToSeq[T any](ch <-chan T) iter.Seq[T]"
playUrl: ""
variantHelpers:
  - it#channel#channeltoseq
similarHelpers:
  - it#channel#seqtochannel
position: 10
---

Builds an `iter.Seq` from a channel. The returned sequence yields each value received from the channel in order. Iteration blocks until the channel is closed.

Examples:

```go
ch := make(chan int, 3)
ch <- 1; ch <- 2; ch <- 3
close(ch)

seq := it.ChannelToSeq(ch)
var got []int
for v := range seq {
    got = append(got, v)
}
// got == []int{1, 2, 3}
```


