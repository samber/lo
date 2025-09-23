package lo

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChannelDispatcher(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	ch := make(chan int, 10)

	ch <- 0
	ch <- 1
	ch <- 2
	ch <- 3

	is.Len(ch, 4)

	children := ChannelDispatcher(ch, 5, 10, DispatchingStrategyRoundRobin[int])
	time.Sleep(10 * time.Millisecond)

	// check channels allocation
	is.Len(children, 5)

	is.Equal(10, cap(children[0]))
	is.Equal(10, cap(children[1]))
	is.Equal(10, cap(children[2]))
	is.Equal(10, cap(children[3]))
	is.Equal(10, cap(children[4]))

	is.Len(children[0], 1)
	is.Len(children[1], 1)
	is.Len(children[2], 1)
	is.Len(children[3], 1)
	is.Empty(children[4])

	// check channels content
	is.Empty(ch)

	msg0, ok0 := <-children[0]
	is.True(ok0)
	is.Zero(msg0)

	msg1, ok1 := <-children[1]
	is.True(ok1)
	is.Equal(1, msg1)

	msg2, ok2 := <-children[2]
	is.True(ok2)
	is.Equal(2, msg2)

	msg3, ok3 := <-children[3]
	is.True(ok3)
	is.Equal(3, msg3)

	// msg4, ok4 := <-children[4]
	// is.False(ok4)
	// is.Zero(msg4)
	// is.Nil(children[4])

	// check it is closed
	close(ch)
	time.Sleep(10 * time.Millisecond)
	is.Panics(func() {
		ch <- 42
	})

	msg0, ok0 = <-children[0]
	is.False(ok0)
	is.Zero(msg0)

	msg1, ok1 = <-children[1]
	is.False(ok1)
	is.Zero(msg1)

	msg2, ok2 = <-children[2]
	is.False(ok2)
	is.Zero(msg2)

	msg3, ok3 = <-children[3]
	is.False(ok3)
	is.Zero(msg3)

	msg4, ok4 := <-children[4]
	is.False(ok4)
	is.Zero(msg4)

	// unbuffered channels
	children = ChannelDispatcher(ch, 5, 0, DispatchingStrategyRoundRobin[int])
	is.Zero(cap(children[0]))
}

func TestDispatchingStrategyRoundRobin(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	children := createChannels[int](3, 2)
	rochildren := channelsToReadOnly(children)
	defer closeChannels(children)

	is.Zero(DispatchingStrategyRoundRobin(42, 0, rochildren))
	is.Equal(1, DispatchingStrategyRoundRobin(42, 1, rochildren))
	is.Equal(2, DispatchingStrategyRoundRobin(42, 2, rochildren))
	is.Zero(DispatchingStrategyRoundRobin(42, 3, rochildren))
}

func TestDispatchingStrategyRandom(t *testing.T) {
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	// with this seed, the order of random channels are: 1 - 0
	rand.Seed(14)

	children := createChannels[int](2, 2)
	rochildren := channelsToReadOnly(children)
	defer closeChannels(children)

	for i := 0; i < 2; i++ {
		children[1] <- i
	}

	is.Zero(DispatchingStrategyRandom(42, 0, rochildren))
}

func TestDispatchingStrategyWeightedRandom(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	children := createChannels[int](2, 2)
	rochildren := channelsToReadOnly(children)
	defer closeChannels(children)

	dispatcher := DispatchingStrategyWeightedRandom[int]([]int{0, 42})

	is.Equal(1, dispatcher(42, 0, rochildren))
	children[0] <- 0
	is.Equal(1, dispatcher(42, 0, rochildren))
	children[1] <- 1
	is.Equal(1, dispatcher(42, 0, rochildren))
}

func TestDispatchingStrategyFirst(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	children := createChannels[int](2, 2)
	rochildren := channelsToReadOnly(children)
	defer closeChannels(children)

	is.Zero(DispatchingStrategyFirst(42, 0, rochildren))
	children[0] <- 0
	is.Zero(DispatchingStrategyFirst(42, 0, rochildren))
	children[0] <- 1
	is.Equal(1, DispatchingStrategyFirst(42, 0, rochildren))
}

func TestDispatchingStrategyLeast(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	children := createChannels[int](2, 2)
	rochildren := channelsToReadOnly(children)
	defer closeChannels(children)

	is.Zero(DispatchingStrategyLeast(42, 0, rochildren))
	children[0] <- 0
	is.Equal(1, DispatchingStrategyLeast(42, 0, rochildren))
	children[1] <- 0
	is.Zero(DispatchingStrategyLeast(42, 0, rochildren))
	children[0] <- 1
	is.Equal(1, DispatchingStrategyLeast(42, 0, rochildren))
	children[1] <- 1
	is.Zero(DispatchingStrategyLeast(42, 0, rochildren))
}

func TestDispatchingStrategyMost(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	children := createChannels[int](2, 2)
	rochildren := channelsToReadOnly(children)
	defer closeChannels(children)

	is.Zero(DispatchingStrategyMost(42, 0, rochildren))
	children[0] <- 0
	is.Zero(DispatchingStrategyMost(42, 0, rochildren))
	children[1] <- 0
	is.Zero(DispatchingStrategyMost(42, 0, rochildren))
	children[0] <- 1
	is.Zero(DispatchingStrategyMost(42, 0, rochildren))
	children[1] <- 1
	is.Zero(DispatchingStrategyMost(42, 0, rochildren))
}

func TestSliceToChannel(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	ch := SliceToChannel(2, []int{1, 2, 3})

	r1, ok1 := <-ch
	r2, ok2 := <-ch
	r3, ok3 := <-ch
	is.True(ok1)
	is.Equal(1, r1)
	is.True(ok2)
	is.Equal(2, r2)
	is.True(ok3)
	is.Equal(3, r3)

	_, ok4 := <-ch
	is.False(ok4)
}

func TestChannelToSlice(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	ch := SliceToChannel(2, []int{1, 2, 3})
	items := ChannelToSlice(ch)

	is.Equal([]int{1, 2, 3}, items)
}

func TestGenerate(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	generator := func(yield func(int)) {
		yield(0)
		yield(1)
		yield(2)
		yield(3)
	}

	i := 0

	for v := range Generator(2, generator) {
		is.Equal(i, v)
		i++
	}

	is.Equal(4, i)
}

func TestBuffer(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	ch := SliceToChannel(2, []int{1, 2, 3})

	items1, length1, _, ok1 := Buffer(ch, 2)
	items2, length2, _, ok2 := Buffer(ch, 2)
	items3, length3, _, ok3 := Buffer(ch, 2)

	is.Equal([]int{1, 2}, items1)
	is.Equal(2, length1)
	is.True(ok1)
	is.Equal([]int{3}, items2)
	is.Equal(1, length2)
	is.False(ok2)
	is.Empty(items3)
	is.Zero(length3)
	is.False(ok3)
}

func TestBufferWithContext(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 200*time.Millisecond)
	is := assert.New(t)

	ch1 := make(chan int, 10)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
		time.Sleep(5 * time.Millisecond)
		cancel()
		ch1 <- 3
		ch1 <- 4
		ch1 <- 5
		close(ch1)
	}()
	items1, length1, _, ok1 := BufferWithContext(ctx, ch1, 20)
	is.Equal([]int{0, 1, 2}, items1)
	is.Equal(3, length1)
	is.True(ok1)

	ch2 := make(chan int, 10)
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	defer close(ch2)
	for i := 0; i < 10; i++ {
		ch2 <- i
	}
	items2, length2, _, ok2 := BufferWithContext(ctx, ch2, 5)
	is.Equal([]int{0, 1, 2, 3, 4}, items2)
	is.Equal(5, length2)
	is.True(ok2)
}

func TestBufferWithTimeout(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 200*time.Millisecond)
	is := assert.New(t)

	generator := func(yield func(int)) {
		for i := 0; i < 5; i++ {
			yield(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
	ch := Generator(0, generator)

	items1, length1, _, ok1 := BufferWithTimeout(ch, 20, 15*time.Millisecond)
	is.Equal([]int{0, 1}, items1)
	is.Equal(2, length1)
	is.True(ok1)

	items2, length2, _, ok2 := BufferWithTimeout(ch, 20, 2*time.Millisecond)
	is.Empty(items2)
	is.Zero(length2)
	is.True(ok2)

	items3, length3, _, ok3 := BufferWithTimeout(ch, 1, 30*time.Millisecond)
	is.Equal([]int{2}, items3)
	is.Equal(1, length3)
	is.True(ok3)

	items4, length4, _, ok4 := BufferWithTimeout(ch, 2, 25*time.Millisecond)
	is.Equal([]int{3, 4}, items4)
	is.Equal(2, length4)
	is.True(ok4)

	items5, length5, _, ok5 := BufferWithTimeout(ch, 3, 25*time.Millisecond)
	is.Empty(items5)
	is.Zero(length5)
	is.False(ok5)
}

func TestFanIn(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	upstreams := createChannels[int](3, 10)
	roupstreams := channelsToReadOnly(upstreams)
	for i := range roupstreams {
		go func(i int) {
			upstreams[i] <- 1
			upstreams[i] <- 1
			close(upstreams[i])
		}(i)
	}
	out := FanIn(10, roupstreams...)
	time.Sleep(10 * time.Millisecond)

	// check input channels
	is.Empty(roupstreams[0])
	is.Empty(roupstreams[1])
	is.Empty(roupstreams[2])

	// check channels allocation
	is.Len(out, 6)
	is.Equal(10, cap(out))

	// check channels content
	for i := 0; i < 6; i++ {
		msg0, ok0 := <-out
		is.True(ok0)
		is.Equal(1, msg0)
	}

	// check it is closed
	time.Sleep(10 * time.Millisecond)
	msg0, ok0 := <-out
	is.False(ok0)
	is.Zero(msg0)
}

func TestFanOut(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	upstream := SliceToChannel(10, []int{0, 1, 2, 3, 4, 5})
	rodownstreams := FanOut(3, 10, upstream)

	time.Sleep(10 * time.Millisecond)

	// check output channels
	is.Len(rodownstreams, 3)

	// check channels allocation
	for i := range rodownstreams {
		is.Len(rodownstreams[i], 6)
		is.Equal(10, cap(rodownstreams[i]))
		is.Equal([]int{0, 1, 2, 3, 4, 5}, ChannelToSlice(rodownstreams[i]))
	}

	// check it is closed
	time.Sleep(10 * time.Millisecond)

	// check channels allocation
	for i := range rodownstreams {
		msg, ok := <-rodownstreams[i]
		is.False(ok)
		is.Zero(msg)
	}
}
