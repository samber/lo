package parallel

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Map([]int{1, 2, 3, 4}, func(x, _ int) string {
		return "Hello"
	})
	result2 := Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal([]string{"Hello", "Hello", "Hello", "Hello"}, result1)
	is.Equal([]string{"1", "2", "3", "4"}, result2)
}

func TestMapWithConcurrency(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var maxActive int64
	var active int64

	result := Map([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x, _ int) int {
		cur := atomic.AddInt64(&active, 1)
		for {
			old := atomic.LoadInt64(&maxActive)
			if cur <= old || atomic.CompareAndSwapInt64(&maxActive, old, cur) {
				break
			}
		}
		time.Sleep(50 * time.Millisecond)
		atomic.AddInt64(&active, -1)
		return x * 2
	}, WithConcurrency(3))

	is.Equal([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, result)
	is.LessOrEqual(atomic.LoadInt64(&maxActive), int64(3))
}

func TestMapErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result, err := MapErr([]int{1, 2, 3, 4}, func(x, _ int) (string, error) {
		return strconv.Itoa(x * 10), nil
	})
	is.NoError(err)
	is.Equal([]string{"10", "20", "30", "40"}, result)
}

func TestMapErrWithError(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	_, err := MapErr([]int{1, 2, 3, 4}, func(x, _ int) (int, error) {
		if x == 3 {
			return 0, fmt.Errorf("item %d failed", x)
		}
		return x, nil
	}, WithConcurrency(1))
	is.Error(err)
	is.Equal("item 3 failed", err.Error())
}

func TestMapErrWithConcurrency(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var maxActive int64
	var active int64

	result, err := MapErr([]int{1, 2, 3, 4, 5}, func(x, _ int) (int, error) {
		cur := atomic.AddInt64(&active, 1)
		for {
			old := atomic.LoadInt64(&maxActive)
			if cur <= old || atomic.CompareAndSwapInt64(&maxActive, old, cur) {
				break
			}
		}
		time.Sleep(50 * time.Millisecond)
		atomic.AddInt64(&active, -1)
		return x * 2, nil
	}, WithConcurrency(2))

	is.NoError(err)
	is.Equal([]int{2, 4, 6, 8, 10}, result)
	is.LessOrEqual(atomic.LoadInt64(&maxActive), int64(2))
}

func TestMapErrWithContext(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Cancel after 100ms. With concurrency=2 and items sleeping 50ms,
	// only a few items should be processed before cancellation kicks in.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var processed int64
	_, err := MapErr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x, _ int) (int, error) {
		atomic.AddInt64(&processed, 1)
		time.Sleep(50 * time.Millisecond)
		return x, nil
	}, WithConcurrency(2), WithContext(ctx))

	is.Error(err)
	// Context cancellation should prevent all 10 items from being processed
	is.Less(atomic.LoadInt64(&processed), int64(10))
}

func TestMapErrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result, err := MapErr([]int{}, func(x, _ int) (int, error) {
		return x, nil
	})
	is.NoError(err)
	is.Equal([]int{}, result)
}

func TestForEach(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var counter uint64
	collection := []int{1, 2, 3, 4}
	ForEach(collection, func(x, i int) {
		atomic.AddUint64(&counter, 1)
	})

	is.Equal(uint64(4), atomic.LoadUint64(&counter))
}

func TestForEachWithConcurrency(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var maxActive int64
	var active int64
	var counter int64

	ForEach([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(x, _ int) {
		cur := atomic.AddInt64(&active, 1)
		for {
			old := atomic.LoadInt64(&maxActive)
			if cur <= old || atomic.CompareAndSwapInt64(&maxActive, old, cur) {
				break
			}
		}
		time.Sleep(50 * time.Millisecond)
		atomic.AddInt64(&active, -1)
		atomic.AddInt64(&counter, 1)
	}, WithConcurrency(2))

	is.Equal(int64(8), atomic.LoadInt64(&counter))
	is.LessOrEqual(atomic.LoadInt64(&maxActive), int64(2))
}

func TestForEachErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var counter int64
	err := ForEachErr([]int{1, 2, 3, 4}, func(x, _ int) error {
		atomic.AddInt64(&counter, 1)
		return nil
	}, WithConcurrency(2))
	is.NoError(err)
	is.Equal(int64(4), atomic.LoadInt64(&counter))
}

func TestForEachErrWithError(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	err := ForEachErr([]int{1, 2, 3, 4, 5}, func(x, _ int) error {
		if x == 3 {
			return fmt.Errorf("item %d failed", x)
		}
		return nil
	}, WithConcurrency(1))
	is.Error(err)
	is.Equal("item 3 failed", err.Error())
}

func TestForEachErrWithContext(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var processed int64
	err := ForEachErr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x, _ int) error {
		atomic.AddInt64(&processed, 1)
		time.Sleep(50 * time.Millisecond)
		return nil
	}, WithConcurrency(2), WithContext(ctx))

	is.Error(err)
	is.Less(atomic.LoadInt64(&processed), int64(10))
}

func TestForEachErrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	err := ForEachErr([]int{}, func(x, _ int) error {
		return errors.New("should not be called")
	})
	is.NoError(err)
}

func TestTimes(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal([]string{"0", "1", "2"}, result1)
}

func TestTimesWithConcurrency(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var maxActive int64
	var active int64

	result := Times(6, func(i int) int {
		cur := atomic.AddInt64(&active, 1)
		for {
			old := atomic.LoadInt64(&maxActive)
			if cur <= old || atomic.CompareAndSwapInt64(&maxActive, old, cur) {
				break
			}
		}
		time.Sleep(50 * time.Millisecond)
		atomic.AddInt64(&active, -1)
		return i * 2
	}, WithConcurrency(2))

	is.Equal([]int{0, 2, 4, 6, 8, 10}, result)
	is.LessOrEqual(atomic.LoadInt64(&maxActive), int64(2))
}

func TestConcurrencyEdgeCases(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	items := []int{1, 2, 3, 4, 5}

	// Negative concurrency — treated as unbounded.
	result := Map(items, func(x, _ int) int { return x * 2 }, WithConcurrency(-1))
	is.Equal([]int{2, 4, 6, 8, 10}, result)

	// Zero concurrency — treated as unbounded.
	result = Map(items, func(x, _ int) int { return x * 2 }, WithConcurrency(0))
	is.Equal([]int{2, 4, 6, 8, 10}, result)

	// Concurrency larger than item count.
	result = Map(items, func(x, _ int) int { return x * 2 }, WithConcurrency(100))
	is.Equal([]int{2, 4, 6, 8, 10}, result)

	// Concurrency of 1 — sequential.
	result = Map(items, func(x, _ int) int { return x * 2 }, WithConcurrency(1))
	is.Equal([]int{2, 4, 6, 8, 10}, result)

	// Single element.
	result = Map([]int{42}, func(x, _ int) int { return x * 2 }, WithConcurrency(3))
	is.Equal([]int{84}, result)

	// Err variants with same edge cases.
	r, err := MapErr(items, func(x, _ int) (int, error) { return x * 2, nil }, WithConcurrency(-1))
	is.NoError(err)
	is.Equal([]int{2, 4, 6, 8, 10}, r)

	r, err = MapErr(items, func(x, _ int) (int, error) { return x * 2, nil }, WithConcurrency(0))
	is.NoError(err)
	is.Equal([]int{2, 4, 6, 8, 10}, r)

	r, err = MapErr(items, func(x, _ int) (int, error) { return x * 2, nil }, WithConcurrency(100))
	is.NoError(err)
	is.Equal([]int{2, 4, 6, 8, 10}, r)
}

// TestErrWorkersDrainChannel verifies that the worker pool shuts down cleanly
// after an error: the sender stops dispatching, workers drain remaining items,
// and the function returns promptly.
func TestErrWorkersDrainChannel(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	done := make(chan struct{})
	go func() {
		defer close(done)
		err := ForEachErr([]int{1, 2, 3, 4, 5}, func(x, _ int) error {
			time.Sleep(10 * time.Millisecond)
			if x == 1 {
				return fmt.Errorf("fail on %d", x)
			}
			return nil
		}, WithConcurrency(2))
		is.Error(err)
	}()

	select {
	case <-done:
		// completed without deadlock
	case <-time.After(2 * time.Second):
		t.Fatal("deadlock: worker pool did not shut down after error")
	}
}

func TestMapErrRaceOnReturn(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Reproduces a data race: the sender breaks on hasErr and returns firstErr
	// before wg.Wait() ensures the worker has finished writing firstErr.
	// Run with -race to detect.
	for i := 0; i < 100; i++ {
		_, err := MapErr([]int{1}, func(x, _ int) (int, error) {
			time.Sleep(time.Millisecond)
			return 0, errors.New("fail")
		}, WithConcurrency(1))
		is.Error(err)
		is.Equal("fail", err.Error())
	}
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	// order
	for x := range result1 {
		sort.Ints(result1[x])
	}

	is.Equal(map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	}, result1)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := GroupBy(allStrings, func(i string) int {
		return 42
	})
	is.IsType(nonempty[42], allStrings, "type preserved")
}

func TestPartitionBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	oddEven := func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	}

	result1 := PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)
	result2 := PartitionBy([]int{}, oddEven)

	// order
	sort.Slice(result1, func(i, j int) bool {
		return result1[i][0] < result1[j][0]
	})
	for x := range result1 {
		sort.Ints(result1[x])
	}

	is.ElementsMatch(result1, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Empty(result2)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := PartitionBy(allStrings, func(item string) int {
		return len(item)
	})
	is.IsType(nonempty[0], allStrings, "type preserved")
}
