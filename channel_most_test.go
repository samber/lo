package lo_test

import (
	"testing"
	"time"

	"github.com/samber/lo"
)

func TestDispatchingStrategyMost_fullChannels(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name       string
		capacities []int
		lengths    []int
		want       int
	}{
		{"fullFirst", []int{2, 2}, []int{2, 1}, 1},
		{"fullFirstEmptySecond", []int{2, 2}, []int{2, 0}, 1},
		{"differentCapacities", []int{3, 2, 5}, []int{3, 1, 2}, 2},
		{"fullMiddle", []int{3, 1, 3}, []int{1, 1, 2}, 2},
		{"availableTie", []int{1, 3, 3}, []int{1, 2, 2}, 1},
		{"unbuffered", []int{1, 0}, []int{1, 0}, 1},
		{"ordinary", []int{3, 3}, []int{1, 2}, 1},
		{"allFull", []int{2, 2}, []int{2, 2}, 0},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			channels := make([]<-chan int, len(tc.capacities))
			for i, capacity := range tc.capacities {
				ch := make(chan int, capacity)
				for j := 0; j < tc.lengths[i]; j++ {
					ch <- j
				}
				channels[i] = ch
			}
			if got := lo.DispatchingStrategyMost(42, 0, channels); got != tc.want {
				t.Fatalf("got channel %d, want %d", got, tc.want)
			}
		})
	}
}

func TestChannelDispatcher_mostSkipsFull(t *testing.T) {
	t.Parallel()

	stream := make(chan int, 3)
	stream <- 1
	stream <- 2
	stream <- 3
	close(stream)
	children := lo.ChannelDispatcher(stream, 2, 2, lo.DispatchingStrategyMost[int])
	// Draining after the assertion also releases the baseline's blocked send.
	defer func() {
		for _, child := range children {
			lo.ChannelToSlice(child)
		}
	}()

	select {
	case got := <-children[1]:
		if got != 3 {
			t.Fatalf("got %d, want third message 3", got)
		}
	case <-time.After(time.Second):
		t.Fatal("dispatcher blocked on full first channel while second had room")
	}
}
