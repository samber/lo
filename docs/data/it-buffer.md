---
name: Buffer
slug: buffer
sourceRef: it/seq.go#L1162
category: it
subCategory: sequence
signatures:
  - "func Buffer[T any](seq iter.Seq[T], size int) iter.Seq[[]T]"
playUrl: https://go.dev/play/p/zDZdcCA20ut
variantHelpers:
  - it#sequence#buffer
similarHelpers:
  - it#sequence#chunk
  - it#sequence#sliding
  - it#sequence#window
position: 65
---

Returns a sequence of slices, each containing up to size items read from the sequence.
The last slice may be smaller if the sequence closes before filling the buffer.

Examples:

```go
seq := func(yield func(int) bool) {
    _ = yield(1)
    _ = yield(2)
    _ = yield(3)
    _ = yield(4)
    _ = yield(5)
    _ = yield(6)
    _ = yield(7)
}
buffers := it.Buffer(seq, 3)
var result [][]int
for buffer := range buffers {
    result = append(result, buffer)
}
// result contains [[1 2 3] [4 5 6] [7]]
```
