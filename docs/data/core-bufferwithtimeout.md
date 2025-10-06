---
name: BufferWithTimeout
slug: bufferwithtimeout
sourceRef: channel.go#L214
category: core
subCategory: channel
signatures:
  - "func BufferWithTimeout[T any](ch <-chan T, size int, timeout time.Duration) (collection []T, length int, readTime time.Duration, ok bool)"
  - "func BatchWithTimeout[T any](ch <-chan T, size int, timeout time.Duration) (collection []T, length int, readTime time.Duration, ok bool)"
variantHelpers:
  - core#channel#bufferwithtimeout
  - core#channel#batchwithtimeout
similarHelpers:
  - core#channel#buffer
  - core#channel#bufferwithcontext
  - core#channel#slicetochannel
  - core#channel#channeltoslice
position: 263
---

BufferWithTimeout reads up to size items from a channel with timeout.

```go
ch := make(chan int)
go func() {
    time.Sleep(200 * time.Millisecond)
    ch <- 1
}()

items, length, readTime, ok := lo.BufferWithTimeout(ch, 5, 100*time.Millisecond)
// Returns empty slice due to timeout
```

### BatchWithTimeout (Deprecated)

BatchWithTimeout is an alias for BufferWithTimeout.

```go
ch := make(chan float64)
go func() {
    time.Sleep(150 * time.Millisecond)
    ch <- 3.14
}()

batch, length, readTime, ok := lo.BatchWithTimeout(ch, 3, 100*time.Millisecond)
// Returns empty slice due to timeout
```