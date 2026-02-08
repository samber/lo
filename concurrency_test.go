package lo

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSynchronize(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
	testWithTimeout(t, 1000*time.Millisecond)
	is := assert.New(t)

	// check that callbacks are not executed concurrently
	{
		start := time.Now()

		wg := sync.WaitGroup{}
		wg.Add(10)

		s := Synchronize()

		for i := 0; i < 10; i++ {
			go s.Do(func() {
				time.Sleep(50 * time.Millisecond)
				wg.Done()
			})
		}

		wg.Wait()

		duration := time.Since(start)
		is.InDelta(500*time.Millisecond, duration, float64(40*time.Millisecond))
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
		is.PanicsWithValue("lo.Synchronize: unexpected arguments", func() {
			mu := &sync.Mutex{}
			Synchronize(mu, mu, mu)
		})
	}
}

func TestAsync(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
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
		is.Equal(10, result)
	case <-time.After(time.Millisecond):
		is.Fail("Async should not block")
	}
}

func TestAsyncX(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
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
			is.Equal(10, result)
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
			is.Equal(Tuple2[int, string]{10, "Hello"}, result)
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
			is.Equal(Tuple3[int, string, bool]{10, "Hello", true}, result)
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
			is.Equal(Tuple4[int, string, bool, float64]{10, "Hello", true, 3.14}, result)
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
			is.Equal(Tuple5[int, string, bool, float64, string]{10, "Hello", true, 3.14, "World"}, result)
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
			is.Equal(Tuple6[int, string, bool, float64, string, int]{10, "Hello", true, 3.14, "World", 100}, result)
		case <-time.After(time.Millisecond):
			is.Fail("Async6 should not block")
		}
	}
}

func TestWaitFor(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	t.Run("exist condition works", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 300*time.Millisecond)
		is := assert.New(t)

		laterTrue := func(i int) bool {
			return i >= 5
		}

		iter, duration, ok := WaitFor(laterTrue, 200*time.Millisecond, 10*time.Millisecond)
		is.Equal(6, iter, "unexpected iteration count")
		is.InDelta(60*time.Millisecond, duration, float64(5*time.Millisecond))
		is.True(ok)
	})

	t.Run("counter is incremented", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 100*time.Millisecond)
		is := assert.New(t)

		counter := 0
		alwaysFalse := func(i int) bool {
			is.Equal(counter, i)
			counter++
			return false
		}

		iter, duration, ok := WaitFor(alwaysFalse, 40*time.Millisecond, 10*time.Millisecond)
		is.Equal(counter, iter, "unexpected iteration count")
		is.InDelta(40*time.Millisecond, duration, float64(5*time.Millisecond))
		is.False(ok)
	})

	alwaysTrue := func(_ int) bool { return true }
	alwaysFalse := func(_ int) bool { return false }

	t.Run("timeout works", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 200*time.Millisecond)
		is := assert.New(t)

		iter, duration, ok := WaitFor(alwaysFalse, 50*time.Millisecond, 100*time.Millisecond)
		is.Zero(iter, "unexpected iteration count")
		is.InDelta(50*time.Millisecond, duration, float64(10*time.Millisecond))
		is.False(ok)
	})

	t.Run("exist on first condition", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 200*time.Millisecond)
		is := assert.New(t)

		iter, duration, ok := WaitFor(alwaysTrue, 100*time.Millisecond, 30*time.Millisecond)
		is.Equal(1, iter, "unexpected iteration count")
		is.InDelta(30*time.Millisecond, duration, float64(10*time.Millisecond))
		is.True(ok)
	})
}

func TestWaitForWithContext(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	t.Run("exist condition works", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 200*time.Millisecond)
		is := assert.New(t)

		laterTrue := func(_ context.Context, i int) bool {
			return i >= 5
		}

		iter, duration, ok := WaitForWithContext(context.Background(), laterTrue, 200*time.Millisecond, 10*time.Millisecond)
		is.Equal(6, iter, "unexpected iteration count")
		is.InDelta(60*time.Millisecond, duration, float64(5*time.Millisecond))
		is.True(ok)
	})

	t.Run("counter is incremented", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 150*time.Millisecond)
		is := assert.New(t)

		counter := 0
		alwaysFalse := func(_ context.Context, i int) bool {
			is.Equal(counter, i)
			counter++
			return false
		}

		iter, duration, ok := WaitForWithContext(context.Background(), alwaysFalse, 80*time.Millisecond, 20*time.Millisecond)
		is.Equal(counter, iter, "unexpected iteration count")
		is.InDelta(80*time.Millisecond, duration, float64(10*time.Millisecond))
		is.False(ok)
	})

	alwaysTrue := func(_ context.Context, _ int) bool { return true }
	alwaysFalse := func(_ context.Context, _ int) bool { return false }

	t.Run("timeout works", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 200*time.Millisecond)
		is := assert.New(t)

		iter, duration, ok := WaitForWithContext(context.Background(), alwaysFalse, 50*time.Millisecond, 100*time.Millisecond)
		is.Zero(iter, "unexpected iteration count")
		is.InDelta(50*time.Millisecond, duration, float64(10*time.Millisecond))
		is.False(ok)
	})

	t.Run("exist on first condition", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 200*time.Millisecond)
		is := assert.New(t)

		iter, duration, ok := WaitForWithContext(context.Background(), alwaysTrue, 100*time.Millisecond, 10*time.Millisecond)
		is.Equal(1, iter, "unexpected iteration count")
		is.InDelta(10*time.Millisecond, duration, float64(5*time.Millisecond))
		is.True(ok)
	})

	t.Run("context cancellation stops everything", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 100*time.Millisecond)
		is := assert.New(t)

		expiringCtx, clean := context.WithTimeout(context.Background(), 45*time.Millisecond)
		t.Cleanup(func() {
			clean()
		})

		iter, duration, ok := WaitForWithContext(expiringCtx, alwaysFalse, 100*time.Millisecond, 30*time.Millisecond)
		is.Equal(1, iter, "unexpected iteration count")
		is.InDelta(45*time.Millisecond, duration, float64(10*time.Millisecond))
		is.False(ok)
	})

	t.Run("canceled context stops everything", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		testWithTimeout(t, 200*time.Millisecond)
		is := assert.New(t)

		canceledCtx, cancel := context.WithCancel(context.Background())
		cancel()

		iter, duration, ok := WaitForWithContext(canceledCtx, alwaysFalse, 100*time.Millisecond, 30*time.Millisecond)
		is.Zero(iter, "unexpected iteration count")
		is.InDelta(1*time.Millisecond, duration, float64(1*time.Millisecond))
		is.False(ok)
	})
}

func TestBatchConcurrentProcess(t *testing.T) {
	is := assert.New(t)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	respArr := make([]int, len(arr))
	ch := BatchConcurrentProcess(ctx, 20, arr, func(index int, item int) {
		item = item * 2
		// suppose some network call delay
		time.Sleep(10 * time.Millisecond)
		respArr[index] = item
	})
	<-ch
	is.Equal(respArr, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40})
}

func TestBatchConcurrentProcess_MaxConcurrency(t *testing.T) {
	is := assert.New(t)

	arr := make([]int, 20)
	ctx := context.Background()

	var (
		mu          sync.Mutex
		current     int
		maxObserved int
	)

	done := BatchConcurrentProcess(ctx, 5, arr, func(index int, item int) {
		mu.Lock()
		current++
		if current > maxObserved {
			maxObserved = current
		}
		mu.Unlock()

		time.Sleep(20 * time.Millisecond)

		mu.Lock()
		current--
		mu.Unlock()
	})
	<-done
	is.LessOrEqual(maxObserved, 5, "concurrency exceeded limit")
}
