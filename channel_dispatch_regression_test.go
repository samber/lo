package lo_test

import (
	"sync"
	"testing"
	"time"

	"github.com/samber/lo"
)

func TestDispatchingStrategyMost_nonFull(t *testing.T) {
	for _, tc := range []struct {
		name string
		caps []int
		fill []int
		want int
	}{
		{"first_full_second_empty", []int{1, 1}, []int{1, 0}, 1},
		{"first_full_second_partial", []int{2, 2}, []int{2, 1}, 1},
		{"different_capacities", []int{3, 2, 4}, []int{3, 1, 2}, 2},
		{"later_full", []int{2, 1}, []int{1, 1}, 0},
		{"nonFull_tie", []int{3, 3}, []int{1, 1}, 0},
		{"all_full_fallback", []int{1, 2}, []int{1, 2}, 0},
		{"all_empty", []int{2, 2}, []int{0, 0}, 0},
		{"unbuffered_after_full", []int{1, 0}, []int{1, 0}, 1},
		{"all_unbuffered", []int{0, 0}, []int{0, 0}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			channels := make([]<-chan int, len(tc.caps))
			for i, capacity := range tc.caps {
				ch := make(chan int, capacity)
				defer close(ch)
				for j := 0; j < tc.fill[i]; j++ {
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

func TestChannelDispatcher_mostUsesAvailableChannel(t *testing.T) {
	input := make(chan int, 2)
	input <- 1
	input <- 2
	close(input)
	children := lo.ChannelDispatcher(input, 2, 1, lo.DispatchingStrategyMost[int])
	defer func() {
		var wg sync.WaitGroup
		for _, ch := range children {
			wg.Add(1)
			go func(ch <-chan int) {
				defer wg.Done()
				for range ch {
				}
			}(ch)
		}
		wg.Wait()
	}()
	select {
	case got := <-children[1]:
		if got != 2 {
			t.Fatalf("got %d, want 2", got)
		}
	case <-time.After(time.Second):
		t.Fatal("dispatcher blocked on full first child while second child was empty")
	}
}
