package lo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChannelDispatcher2(t *testing.T) {
	ch := make(chan int, 42)
	for i := 0; i <= 10; i++ {
		ch <- i
	}

	children := ChannelDispatcher(ch, 5, 10, DispatchingStrategyRoundRobin[int])
	// []<-chan int{...}

	consumer := func(c <-chan int) {
		for {
			msg, ok := <-c
			if !ok {
				println("closed")
				break
			}

			println(msg)
		}
	}

	for i := range children {
		go consumer(children[i])
	}
}

func TestChannelDispatcher(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	ch := make(chan int, 10)

	ch <- 0
	ch <- 1
	ch <- 2
	ch <- 3

	is.Equal(4, len(ch))

	children := ChannelDispatcher(ch, 5, 10, DispatchingStrategyRoundRobin[int])
	time.Sleep(10 * time.Millisecond)

	// check channels allocation
	is.Equal(5, len(children))

	is.Equal(10, cap(children[0]))
	is.Equal(10, cap(children[1]))
	is.Equal(10, cap(children[2]))
	is.Equal(10, cap(children[3]))
	is.Equal(10, cap(children[4]))

	is.Equal(1, len(children[0]))
	is.Equal(1, len(children[1]))
	is.Equal(1, len(children[2]))
	is.Equal(1, len(children[3]))
	is.Equal(0, len(children[4]))

	// check channels content
	is.Equal(0, len(ch))

	msg0, ok0 := <-children[0]
	is.Equal(ok0, true)
	is.Equal(msg0, 0)

	msg1, ok1 := <-children[1]
	is.Equal(ok1, true)
	is.Equal(msg1, 1)

	msg2, ok2 := <-children[2]
	is.Equal(ok2, true)
	is.Equal(msg2, 2)

	msg3, ok3 := <-children[3]
	is.Equal(ok3, true)
	is.Equal(msg3, 3)

	// msg4, ok4 := <-children[4]
	// is.Equal(ok4, false)
	// is.Equal(msg4, 0)
	// is.Nil(children[4])

	// check it is closed
	close(ch)
	time.Sleep(10 * time.Millisecond)
	is.Panics(func() {
		ch <- 42
	})

	msg0, ok0 = <-children[0]
	is.Equal(ok0, false)
	is.Equal(msg0, 0)

	msg1, ok1 = <-children[1]
	is.Equal(ok1, false)
	is.Equal(msg1, 0)

	msg2, ok2 = <-children[2]
	is.Equal(ok2, false)
	is.Equal(msg2, 0)

	msg3, ok3 = <-children[3]
	is.Equal(ok3, false)
	is.Equal(msg3, 0)

	msg4, ok4 := <-children[4]
	is.Equal(ok4, false)
	is.Equal(msg4, 0)

	// unbuffered channels
	children = ChannelDispatcher(ch, 5, 0, DispatchingStrategyRoundRobin[int])
	is.Equal(0, cap(children[0]))
}

func TestDispatchingStrategyRoundRobin(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	children := createChannels[int](3, 2)
	rochildren := channelsToReadOnly(children)
	defer closeChannels(children)

	is.Equal(0, DispatchingStrategyRoundRobin(42, 0, rochildren))
	is.Equal(1, DispatchingStrategyRoundRobin(42, 1, rochildren))
	is.Equal(2, DispatchingStrategyRoundRobin(42, 2, rochildren))
	is.Equal(0, DispatchingStrategyRoundRobin(42, 3, rochildren))
}

func TestDispatchingStrategyRandom(t *testing.T) {
	// @TODO
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

	is.Equal(0, DispatchingStrategyFirst(42, 0, rochildren))
	children[0] <- 0
	is.Equal(0, DispatchingStrategyFirst(42, 0, rochildren))
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

	is.Equal(0, DispatchingStrategyLeast(42, 0, rochildren))
	children[0] <- 0
	is.Equal(1, DispatchingStrategyLeast(42, 0, rochildren))
	children[1] <- 0
	is.Equal(0, DispatchingStrategyLeast(42, 0, rochildren))
	children[0] <- 1
	is.Equal(1, DispatchingStrategyLeast(42, 0, rochildren))
	children[1] <- 1
	is.Equal(0, DispatchingStrategyLeast(42, 0, rochildren))
}

func TestDispatchingStrategyMost(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 10*time.Millisecond)
	is := assert.New(t)

	children := createChannels[int](2, 2)
	rochildren := channelsToReadOnly(children)
	defer closeChannels(children)

	is.Equal(0, DispatchingStrategyMost(42, 0, rochildren))
	children[0] <- 0
	is.Equal(0, DispatchingStrategyMost(42, 0, rochildren))
	children[1] <- 0
	is.Equal(0, DispatchingStrategyMost(42, 0, rochildren))
	children[0] <- 1
	is.Equal(0, DispatchingStrategyMost(42, 0, rochildren))
	children[1] <- 1
	is.Equal(0, DispatchingStrategyMost(42, 0, rochildren))
}
