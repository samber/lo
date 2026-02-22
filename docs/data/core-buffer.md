---
name: Buffer
slug: buffer
sourceRef: channel.go#L214
category: core
subCategory: channel
signatures:
  - "func Buffer[T any](ch <-chan T, size int) (collection []T, length int, readTime time.Duration, ok bool)"
  - "func BufferWithContext[T any](ctx context.Context, ch <-chan T, size int) (collection []T, length int, readTime time.Duration, ok bool)"
variantHelpers:
  - core#channel#buffer
  - core#channel#bufferwithcontext
similarHelpers:
  - core#channel#slicetochannel
  - core#channel#channeltoslice
position: 260
---

Buffer reads up to size items from a channel and returns them as a slice.

```go
ch := make(chan int, 10)
for i := 1; i <= 5; i++ {
    ch <- i
}
close(ch)

items, length, readTime, ok := lo.Buffer(ch, 3)
// items: []int{1, 2, 3}
// length: 3
// readTime: ~0s (immediate read from buffered channel)
// ok: true (channel was closed)
```

### BufferWithContext

BufferWithContext reads up to size items from a channel with context cancellation.

```go
ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
defer cancel()

ch := make(chan int)
go func() {
    time.Sleep(50 * time.Millisecond)
    ch <- 1
    time.Sleep(100 * time.Millisecond)
    ch <- 2
}()

items, length, readTime, ok := lo.BufferWithContext(ctx, ch, 5)
// items: []int{1} (only first item received before timeout)
// length: 1
// readTime: ~100ms (context timeout)
// ok: false (context cancelled)
```