---
name: DispatchingStrategy
slug: dispatchingstrategy
sourceRef: channel.go#L12
playUrl: "https://go.dev/play/p/UZGu2wVg3J2"
category: core
subCategory: channel
signatures:
  - "func DispatchingStrategyRoundRobin[T any](msg T, index uint64, channels []<-chan T) int"
  - "func DispatchingStrategyRandom[T any](msg T, index uint64, channels []<-chan T) int"
  - "func DispatchingStrategyWeightedRandom[T any](weights []int) DispatchingStrategy[T]"
  - "func DispatchingStrategyFirst[T any](msg T, index uint64, channels []<-chan T) int"
  - "func DispatchingStrategyLeast[T any](msg T, index uint64, channels []<-chan T) int"
  - "func DispatchingStrategyMost[T any](msg T, index uint64, channels []<-chan T) int"
variantHelpers:
  - core#channel#dispatchingstrategy
  - core#channel#dispatchingstrategy
  - core#channel#dispatchingstrategy
  - core#channel#dispatchingstrategy
  - core#channel#dispatchingstrategy
  - core#channel#dispatchingstrategy
similarHelpers:
  - core#channel#channeldispatcher
position: 270
---

DispatchingStrategyRoundRobin distributes messages to channels in round-robin order.

```go
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
ch3 := make(chan int, 1)
strategy := lo.DispatchingStrategyRoundRobin[int]
index := strategy(42, 0, []<-chan int{ch1, ch2, ch3})
// Returns 0, then 1, then 2, then 0, cycling through channels
```

DispatchingStrategyRandom distributes messages to a random channel.

```go
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
ch3 := make(chan int, 1)
strategy := lo.DispatchingStrategyRandom[int]
index := strategy(42, 0, []<-chan int{ch1, ch2, ch3})
// Returns a random channel index 0, 1, or 2
```

DispatchingStrategyWeightedRandom distributes messages to channels based on weights.

```go
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
ch3 := make(chan int, 1)
weights := []int{1, 3, 6} // Channel 0: 10%, Channel 1: 30%, Channel 2: 60%
strategy := lo.DispatchingStrategyWeightedRandom[int](weights)
index := strategy(42, 0, []<-chan int{ch1, ch2, ch3})
// Returns 2 most often, 1 sometimes, 0 rarely
```

DispatchingStrategyFirst distributes messages to the first non-full channel.

```go
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
ch3 := make(chan int, 1)
strategy := lo.DispatchingStrategyFirst[int]
index := strategy(42, 0, []<-chan int{ch1, ch2, ch3})
// Always returns 0 if ch1 is not full
```

DispatchingStrategyLeast distributes messages to the channel with the fewest items.

```go
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
ch3 := make(chan int, 1)
strategy := lo.DispatchingStrategyLeast[int]
index := strategy(42, 0, []<-chan int{ch1, ch2, ch3})
// Returns the index of the channel with the smallest buffer size
```

DispatchingStrategyMost distributes messages to the channel with the most items.

```go
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
ch3 := make(chan int, 1)
strategy := lo.DispatchingStrategyMost[int]
index := strategy(42, 0, []<-chan int{ch1, ch2, ch3})
// Returns the index of the channel with the largest buffer size
```