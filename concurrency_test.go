package lo

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSynchronize(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	// check that callbacks are not executed concurrently
	{
		start := time.Now()

		wg := sync.WaitGroup{}
		wg.Add(10)

		s := Synchronize()

		for i := 0; i < 10; i++ {
			go s.Do(func() {
				time.Sleep(5 * time.Millisecond)
				wg.Done()
			})
		}

		wg.Wait()

		duration := time.Since(start)

		is.Greater(duration, 50*time.Millisecond)
		is.Less(duration, 60*time.Millisecond)
	}

	// check locker is locked
	{
		mu := &sync.Mutex{}
		s := Synchronize(mu)

		s.Do(func() {
			is.False(mu.TryLock())
		})
		is.True(mu.TryLock())

		Try0(func() {
			mu.Unlock()
		})
	}

	// check we don't accept multiple arguments
	{
		is.PanicsWithValue("unexpected arguments", func() {
			mu := &sync.Mutex{}
			Synchronize(mu, mu, mu)
		})
	}
}

func TestAsync(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	sync := make(chan struct{})

	ch := Async(func() int {
		<-sync
		return 10
	})

	sync <- struct{}{}

	select {
	case result := <-ch:
		is.Equal(result, 10)
	case <-time.After(time.Millisecond):
		is.Fail("Async should not block")
	}
}

func TestAsyncX(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	{
		sync := make(chan struct{})

		ch := Async0(func() {
			<-sync
		})

		sync <- struct{}{}

		select {
		case <-ch:
		case <-time.After(time.Millisecond):
			is.Fail("Async0 should not block")
		}
	}

	{
		sync := make(chan struct{})

		ch := Async1(func() int {
			<-sync
			return 10
		})

		sync <- struct{}{}

		select {
		case result := <-ch:
			is.Equal(result, 10)
		case <-time.After(time.Millisecond):
			is.Fail("Async1 should not block")
		}
	}

	{
		sync := make(chan struct{})

		ch := Async2(func() (int, string) {
			<-sync
			return 10, "Hello"
		})

		sync <- struct{}{}

		select {
		case result := <-ch:
			is.Equal(result, Tuple2[int, string]{10, "Hello"})
		case <-time.After(time.Millisecond):
			is.Fail("Async2 should not block")
		}
	}

	{
		sync := make(chan struct{})

		ch := Async3(func() (int, string, bool) {
			<-sync
			return 10, "Hello", true
		})

		sync <- struct{}{}

		select {
		case result := <-ch:
			is.Equal(result, Tuple3[int, string, bool]{10, "Hello", true})
		case <-time.After(time.Millisecond):
			is.Fail("Async3 should not block")
		}
	}

	{
		sync := make(chan struct{})

		ch := Async4(func() (int, string, bool, float64) {
			<-sync
			return 10, "Hello", true, 3.14
		})

		sync <- struct{}{}

		select {
		case result := <-ch:
			is.Equal(result, Tuple4[int, string, bool, float64]{10, "Hello", true, 3.14})
		case <-time.After(time.Millisecond):
			is.Fail("Async4 should not block")
		}
	}

	{
		sync := make(chan struct{})

		ch := Async5(func() (int, string, bool, float64, string) {
			<-sync
			return 10, "Hello", true, 3.14, "World"
		})

		sync <- struct{}{}

		select {
		case result := <-ch:
			is.Equal(result, Tuple5[int, string, bool, float64, string]{10, "Hello", true, 3.14, "World"})
		case <-time.After(time.Millisecond):
			is.Fail("Async5 should not block")
		}
	}

	{
		sync := make(chan struct{})

		ch := Async6(func() (int, string, bool, float64, string, int) {
			<-sync
			return 10, "Hello", true, 3.14, "World", 100
		})

		sync <- struct{}{}

		select {
		case result := <-ch:
			is.Equal(result, Tuple6[int, string, bool, float64, string, int]{10, "Hello", true, 3.14, "World", 100})
		case <-time.After(time.Millisecond):
			is.Fail("Async6 should not block")
		}
	}
}

func TestTimeout(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	err := Timeout(10*time.Millisecond, func(done func()) {
		done()
	})
	is.Nil(err)

	err = Timeout(10*time.Millisecond, func(done func()) {
		time.Sleep(20 * time.Millisecond)
		done()
	})
	is.Error(err)
	is.Equal(err, context.DeadlineExceeded)
}

func TestDeadline(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	err := Deadline(time.Now().Add(10*time.Millisecond), func(done func()) {
		done()
	})
	is.Nil(err)

	err = Deadline(time.Now().Add(10*time.Millisecond), func(done func()) {
		time.Sleep(20 * time.Millisecond)
		done()
	})
	is.Error(err)
	is.Equal(err, context.DeadlineExceeded)
}

func TestRace(t *testing.T) {
	t.Parallel()
	testWithTimeout(t, 100*time.Millisecond)
	is := assert.New(t)

	var wonRace int32

	func1 := func(done func()) {
		time.Sleep(5 * time.Millisecond)
		atomic.CompareAndSwapInt32(&wonRace, 0, 1)
		done()
	}

	func2 := func(done func()) {
		time.Sleep(30 * time.Millisecond)
		atomic.CompareAndSwapInt32(&wonRace, 0, 2)
		done()
	}

	func3 := func(done func()) {
		time.Sleep(50 * time.Millisecond)
		atomic.CompareAndSwapInt32(&wonRace, 0, 3)
		done()
	}

	Race(func1, func2, func3)
	is.EqualValues(1, atomic.LoadInt32(&wonRace))
}
