package lo

import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestThrottleCallbackReset(t *testing.T) {
	t.Parallel()
	for _, count := range []int{1, 3} {
		count := count
		t.Run(strconv.Itoa(count), func(t *testing.T) {
			t.Parallel()
			calls := 0
			var reset func()
			throttle, resetFn := NewThrottleWithCount(time.Hour, count, func() {
				calls++
				reset()
			})
			reset = resetFn
			done := make(chan struct{})
			go func() {
				throttle()
				throttle()
				close(done)
			}()
			select {
			case <-done:
				reset()
				if calls != 2 {
					t.Fatalf("got %d calls, want 2 after callback reset", calls)
				}
			case <-time.After(time.Second):
				t.Fatal("callback reset deadlocked")
			}
		})
	}
}

func TestThrottleCallbacksDoNotBlockOtherKeys(t *testing.T) {
	t.Parallel()
	entered := make(chan struct{})
	release := make(chan struct{})
	firstDone := make(chan struct{})
	var calls int32
	throttle, reset := NewThrottleBy(time.Hour, func(key string) {
		atomic.AddInt32(&calls, 1)
		if key == "slow" {
			close(entered)
			<-release
		}
	})
	go func() {
		throttle("slow")
		close(firstDone)
	}()
	<-entered
	secondDone := make(chan struct{})
	go func() {
		throttle("other")
		throttle("other")
		close(secondDone)
	}()
	select {
	case <-secondDone:
	case <-time.After(time.Second):
		close(release)
		<-firstDone
		t.Fatal("a blocked callback prevented another key from running")
	}
	close(release)
	<-firstDone
	reset()
	if got := atomic.LoadInt32(&calls); got != 2 {
		t.Fatalf("got %d calls, want 2", got)
	}
}

func TestThrottleConcurrentCountLimit(t *testing.T) {
	t.Parallel()
	const count = 3
	var calls int32
	throttle, reset := NewThrottleByWithCount(time.Hour, count, func(string) { atomic.AddInt32(&calls, 1) })
	defer reset()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			throttle("same")
		}()
	}
	wg.Wait()
	if got := atomic.LoadInt32(&calls); got != count {
		t.Fatalf("got %d calls, want %d", got, count)
	}
	reset()
	throttle("same")
	if got := atomic.LoadInt32(&calls); got != count+1 {
		t.Fatalf("reset did not restore eligibility: got %d", got)
	}
}

func TestThrottleCallbackReentry(t *testing.T) {
	t.Parallel()
	var throttle func(string)
	calls := 0
	throttle, reset := NewThrottleBy(time.Hour, func(key string) {
		calls++
		throttle(key)
	})
	done := make(chan struct{})
	go func() {
		throttle("same")
		close(done)
	}()
	select {
	case <-done:
		reset()
		if calls != 1 {
			t.Fatalf("got %d calls, want 1", calls)
		}
	case <-time.After(time.Second):
		t.Fatal("callback reentry deadlocked")
	}
}
