---
name: SeqToChannel2
slug: channeltoseq
sourceRef: it/channel.go#L26
category: it
subCategory: channel
signatures:
  - "func SeqToChannel2[K, V any](bufferSize int, collection iter.Seq2[K, V]) <-chan lo.Tuple2[K, V]"
  - "func ChannelToSeq[T any](ch <-chan T) iter.Seq[T]"
variantHelpers:
  - it#channel#seqtochannel
  - it#channel#seqtochannel2
  - it#channel#channeltoseq
similarHelpers:
  - core#channel#channelseq
position: 10
---

SeqToChannel2 returns a read-only channel of key-value tuple elements from a sequence.

```go
collection := func(yield func(int, string) bool) {
    yield(1, "a")
    yield(2, "b")
}

ch := it.SeqToChannel2(10, collection)
for tuple := range ch {
    fmt.Printf("%d: %s\n", tuple.A, tuple.B)
}
// 1: a
// 2: b
```

ChannelToSeq returns a sequence built from channel items. Blocks until channel closes.

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

seq := it.ChannelToSeq(ch)
for item := range seq {
    fmt.Println(item)
}
// 1
// 2
// 3
```