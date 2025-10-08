---
name: SeqToChannel
slug: seqtochannel
sourceRef: it/channel.go#L12
category: it
subCategory: channel
signatures:
  - "func SeqToChannel[T any](bufferSize int, collection iter.Seq[T]) <-chan T"
  - "func SeqToChannel2[K, V any](bufferSize int, collection iter.Seq2[K, V]) <-chan Tuple2[K, V]"
playUrl: "https://go.dev/play/p/id3jqJPffT6"
variantHelpers:
  - it#channel#seqtochannel
  - it#channel#seqtochannel2
similarHelpers:
  - it#channel#channeltoseq
position: 0
---

Converts an `iter.Seq` (or `iter.Seq2`) into a read-only channel. Items are sent on a buffered channel and the channel is closed when the sequence ends.

Examples:

```go
// SeqToChannel: stream ints from a sequence
seq := it.Range(5) // 0..4
ch := it.SeqToChannel(2, seq)
var got []int
for v := range ch {
    got = append(got, v)
}
// got == []int{0, 1, 2, 3, 4}
```

```go
// SeqToChannel2: stream key/value pairs as Tuple2
m := map[string]int{"a": 1, "b": 2}
kv := it.Entries(m)
ch := it.SeqToChannel2(1, kv)
for pair := range ch {
    // pair.A is key, pair.B is value
}
```


