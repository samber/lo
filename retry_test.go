package lo

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAttempt(t *testing.T) {
	t.Parallel()

	err := errors.New("failed")

	tests := []struct {
		name         string
		maxIteration int
		fn           func(i int) error
		expectedIter int
		expectErr    bool
	}{
		{
			name:         "always succeeds",
			maxIteration: 42,
			fn: func(i int) error {
				return nil
			},
			expectedIter: 1,
			expectErr:    false,
		},
		{
			name:         "succeeds after some attempts",
			maxIteration: 42,
			fn: func(i int) error {
				if i == 5 {
					return nil
				}

				return err
			},
			expectedIter: 6,
			expectErr:    false,
		},
		{
			name:         "exhausts max iterations before success",
			maxIteration: 2,
			fn: func(i int) error {
				if i == 5 {
					return nil
				}

				return err
			},
			expectedIter: 2,
			expectErr:    true,
		},
		{
			name:         "unlimited iterations until success",
			maxIteration: 0,
			fn: func(i int) error {
				if i < 42 {
					return err
				}

				return nil
			},
			expectedIter: 43,
			expectErr:    false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			iter, gotErr := Attempt(tt.maxIteration, tt.fn)
			is.Equal(tt.expectedIter, iter)
			if tt.expectErr {
				is.ErrorIs(gotErr, err)
			} else {
				is.NoError(gotErr)
			}
		})
	}
}

func TestAttemptWithDelay(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	err := errors.New("failed")

	tests := []struct {
		name          string
		maxIteration  int
		fn            func(i int, d time.Duration) error
		expectedIter  int
		expectedDelta time.Duration
		deltaEpsilon  time.Duration
		expectErr     bool
	}{
		{
			name:         "always succeeds",
			maxIteration: 42,
			fn: func(i int, d time.Duration) error {
				return nil
			},
			expectedIter:  1,
			expectedDelta: 1 * time.Microsecond,
			deltaEpsilon:  1 * time.Millisecond,
			expectErr:     false,
		},
		{
			name:         "succeeds after some attempts",
			maxIteration: 42,
			fn: func(i int, d time.Duration) error {
				if i == 3 {
					return nil
				}

				return err
			},
			expectedIter:  4,
			expectedDelta: 30 * time.Millisecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     false,
		},
		{
			name:         "exhausts max iterations before success",
			maxIteration: 2,
			fn: func(i int, d time.Duration) error {
				if i == 3 {
					return nil
				}

				return err
			},
			expectedIter:  2,
			expectedDelta: 10 * time.Millisecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     true,
		},
		{
			name:         "unlimited iterations until success",
			maxIteration: 0,
			fn: func(i int, d time.Duration) error {
				if i < 10 {
					return err
				}

				return nil
			},
			expectedIter:  11,
			expectedDelta: 100 * time.Millisecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     false,
		},
	}

	for _, tt := range tests { //nolint:paralleltest
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			is := assert.New(t)

			iter, dur, gotErr := AttemptWithDelay(tt.maxIteration, 10*time.Millisecond, tt.fn)
			is.Equal(tt.expectedIter, iter)
			is.InDelta(tt.expectedDelta, dur, float64(tt.deltaEpsilon))
			if tt.expectErr {
				is.ErrorIs(gotErr, err)
			} else {
				is.NoError(gotErr)
			}
		})
	}
}

func TestAttemptWhile(t *testing.T) {
	t.Parallel()

	err := errors.New("failed")

	tests := []struct {
		name         string
		maxIteration int
		fn           func(i int) (error, bool)
		expectedIter int
		expectErr    bool
	}{
		{
			name:         "always succeeds",
			maxIteration: 42,
			fn: func(i int) (error, bool) {
				return nil, true
			},
			expectedIter: 1,
			expectErr:    false,
		},
		{
			name:         "succeeds after some attempts",
			maxIteration: 42,
			fn: func(i int) (error, bool) {
				if i == 5 {
					return nil, true
				}

				return err, true
			},
			expectedIter: 6,
			expectErr:    false,
		},
		{
			name:         "exhausts max iterations before success",
			maxIteration: 2,
			fn: func(i int) (error, bool) {
				if i == 5 {
					return nil, true
				}

				return err, true
			},
			expectedIter: 2,
			expectErr:    true,
		},
		{
			name:         "unlimited iterations until success",
			maxIteration: 0,
			fn: func(i int) (error, bool) {
				if i < 42 {
					return err, true
				}

				return nil, true
			},
			expectedIter: 43,
			expectErr:    false,
		},
		{
			name:         "stops early with no error",
			maxIteration: 0,
			fn: func(i int) (error, bool) {
				if i == 5 {
					return nil, false
				}

				return err, true
			},
			expectedIter: 6,
			expectErr:    false,
		},
		{
			name:         "stops on first iteration",
			maxIteration: 0,
			fn: func(i int) (error, bool) {
				return nil, false
			},
			expectedIter: 1,
			expectErr:    false,
		},
		{
			name:         "stops right before max iteration reached",
			maxIteration: 42,
			fn: func(i int) (error, bool) {
				if i == 42 {
					return nil, false
				}
				if i < 41 {
					return err, true
				}

				return nil, true
			},
			expectedIter: 42,
			expectErr:    false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			iter, gotErr := AttemptWhile(tt.maxIteration, tt.fn)
			is.Equal(tt.expectedIter, iter)
			if tt.expectErr {
				is.ErrorIs(gotErr, err)
			} else {
				is.NoError(gotErr)
			}
		})
	}
}

func TestAttemptWhileWithDelay(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	err := errors.New("failed")

	tests := []struct {
		name          string
		maxIteration  int
		fn            func(i int, d time.Duration) (error, bool)
		expectedIter  int
		expectedDelta time.Duration
		deltaEpsilon  time.Duration
		expectErr     bool
	}{
		{
			name:         "always succeeds",
			maxIteration: 42,
			fn: func(i int, d time.Duration) (error, bool) {
				return nil, true
			},
			expectedIter:  1,
			expectedDelta: 1 * time.Microsecond,
			deltaEpsilon:  3 * time.Millisecond,
			expectErr:     false,
		},
		{
			name:         "succeeds after some attempts",
			maxIteration: 42,
			fn: func(i int, d time.Duration) (error, bool) {
				if i == 3 {
					return nil, true
				}

				return err, true
			},
			expectedIter:  4,
			expectedDelta: 30 * time.Millisecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     false,
		},
		{
			name:         "exhausts max iterations before success",
			maxIteration: 2,
			fn: func(i int, d time.Duration) (error, bool) {
				if i == 5 {
					return nil, true
				}

				return err, true
			},
			expectedIter:  2,
			expectedDelta: 10 * time.Millisecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     true,
		},
		{
			name:         "unlimited iterations until success",
			maxIteration: 0,
			fn: func(i int, d time.Duration) (error, bool) {
				if i < 10 {
					return err, true
				}

				return nil, true
			},
			expectedIter:  11,
			expectedDelta: 100 * time.Millisecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     false,
		},
		{
			name:         "stops early with no error",
			maxIteration: 0,
			fn: func(i int, d time.Duration) (error, bool) {
				if i == 3 {
					return nil, false
				}

				return err, true
			},
			expectedIter:  4,
			expectedDelta: 30 * time.Millisecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     false,
		},
		{
			name:         "stops on first iteration",
			maxIteration: 0,
			fn: func(i int, d time.Duration) (error, bool) {
				return nil, false
			},
			expectedIter:  1,
			expectedDelta: 1 * time.Microsecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     false,
		},
		{
			name:         "stops right before max iteration reached",
			maxIteration: 42,
			fn: func(i int, d time.Duration) (error, bool) {
				if i == 42 {
					return nil, false
				}
				if i < 41 {
					return err, true
				}

				return nil, true
			},
			expectedIter:  42,
			expectedDelta: 410 * time.Millisecond,
			deltaEpsilon:  5 * time.Millisecond,
			expectErr:     false,
		},
	}

	for _, tt := range tests { //nolint:paralleltest
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			is := assert.New(t)

			iter, dur, gotErr := AttemptWhileWithDelay(tt.maxIteration, 10*time.Millisecond, tt.fn)
			is.Equal(tt.expectedIter, iter)
			is.InDelta(tt.expectedDelta, dur, float64(tt.deltaEpsilon))
			if tt.expectErr {
				is.ErrorIs(gotErr, err)
			} else {
				is.NoError(gotErr)
			}
		})
	}
}

func TestDebounce(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	t.Run("repeated bursts each trigger a call", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		f1 := func() {
			println("1. Called once after 10ms when func stopped invoking!")
		}

		d1, _ := NewDebounce(100*time.Millisecond, f1)

		// execute 3 times
		for i := 0; i < 3; i++ {
			for j := 0; j < 10; j++ {
				d1()
			}
			time.Sleep(200 * time.Millisecond)
		}
	})

	t.Run("only last invoke within the window is worked", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		f2 := func() {
			println("2. Called once after 10ms when func stopped invoking!")
		}

		d2, _ := NewDebounce(100*time.Millisecond, f2)

		// execute once because it is always invoked and only last invoke is worked after 100ms
		for i := 0; i < 3; i++ {
			for j := 0; j < 5; j++ {
				d2()
			}
			time.Sleep(50 * time.Millisecond)
		}

		time.Sleep(100 * time.Millisecond)
	})

	t.Run("canceled debounce stops invoking", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()

		f3 := func() {
			println("3. Called once after 10ms when func stopped invoking!")
		}

		// execute once because it is canceled after 200ms.
		d3, cancel := NewDebounce(100*time.Millisecond, f3)
		for i := 0; i < 3; i++ {
			for j := 0; j < 10; j++ {
				d3()
			}
			time.Sleep(200 * time.Millisecond)
			if i == 0 {
				cancel()
			}
		}
	})
}

func TestDebounceBy(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	mu := sync.Mutex{}
	output := map[int]int{0: 0, 1: 0, 2: 0}

	// NOTE: these subtests are intentionally sequential (not parallel) and share
	// the `output` map above: each one accumulates on top of the previous one's
	// result, and the expected totals below reflect that cumulative state.

	t.Run("repeated bursts each trigger a call", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		f1 := func(key, count int) {
			mu.Lock()
			output[key] += count
			mu.Unlock()
			// fmt.Printf("[key=%d] 1. Called once after 10ms when func stopped invoking!\n", key)
		}

		d1, _ := NewDebounceBy(100*time.Millisecond, f1)

		// execute 3 times
		for i := 0; i < 3; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 3; k++ {
					d1(k)
				}
			}
			time.Sleep(200 * time.Millisecond)
		}

		// Wait for debounced calls to complete
		time.Sleep(150 * time.Millisecond)

		mu.Lock()
		is.Equal(30, output[0])
		is.Equal(30, output[1])
		is.Equal(30, output[2])
		mu.Unlock()
	})

	t.Run("only last invoke within the window is worked", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		f2 := func(key, count int) {
			mu.Lock()
			output[key] += count
			mu.Unlock()
			// fmt.Printf("[key=%d] 2. Called once after 10ms when func stopped invoking!\n", key)
		}

		d2, _ := NewDebounceBy(100*time.Millisecond, f2)

		// execute once because it is always invoked and only last invoke is worked after 100ms
		for i := 0; i < 3; i++ {
			for j := 0; j < 5; j++ {
				for k := 0; k < 3; k++ {
					d2(k)
				}
			}
			time.Sleep(50 * time.Millisecond)
		}

		// Wait for debounced calls to complete
		time.Sleep(150 * time.Millisecond)

		mu.Lock()
		is.Equal(45, output[0])
		is.Equal(45, output[1])
		is.Equal(45, output[2])
		mu.Unlock()
	})

	t.Run("canceled debounce stops invoking", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		f3 := func(key, count int) {
			mu.Lock()
			output[key] += count
			mu.Unlock()
			// fmt.Printf("[key=%d] 3. Called once after 10ms when func stopped invoking!\n", key)
		}

		// execute once because it is canceled after 200ms.
		d3, cancel := NewDebounceBy(100*time.Millisecond, f3)
		for i := 0; i < 3; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 3; k++ {
					d3(k)
				}
			}

			time.Sleep(200 * time.Millisecond)
			if i == 0 {
				for k := 0; k < 3; k++ {
					cancel(k)
				}
			}
		}

		// Wait for debounced calls to complete
		time.Sleep(150 * time.Millisecond)

		mu.Lock()
		is.Equal(75, output[0])
		is.Equal(75, output[1])
		is.Equal(75, output[2])
		mu.Unlock()
	})
}

func TestTransaction(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		transaction := NewTransaction[int]().
			Then(
				func(state int) (int, error) {
					return state + 100, nil
				},
				func(state int) int {
					return state - 100
				},
			).
			Then(
				func(state int) (int, error) {
					return state + 21, nil
				},
				func(state int) int {
					return state - 21
				},
			)

		state, err := transaction.Process(21)
		is.Equal(142, state)
		is.NoError(err)
	})

	t.Run("with error", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		transaction := NewTransaction[int]().
			Then(
				func(state int) (int, error) {
					return state + 100, nil
				},
				func(state int) int {
					return state - 100
				},
			).
			Then(
				func(state int) (int, error) {
					return state, assert.AnError
				},
				func(state int) int {
					return state - 21
				},
			).
			Then(
				func(state int) (int, error) {
					return state + 42, nil
				},
				func(state int) int {
					return state - 42
				},
			)

		state, err := transaction.Process(21)
		is.Equal(21, state)
		is.ErrorIs(err, assert.AnError)
	})

	t.Run("with error and update value", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		transaction := NewTransaction[int]().
			Then(
				func(state int) (int, error) {
					return state + 100, nil
				},
				func(state int) int {
					return state - 100
				},
			).
			Then(
				func(state int) (int, error) {
					return state + 21, assert.AnError
				},
				func(state int) int {
					return state - 21
				},
			).
			Then(
				func(state int) (int, error) {
					return state + 42, nil
				},
				func(state int) int {
					return state - 42
				},
			)

		state, err := transaction.Process(21)
		is.Equal(42, state)
		is.ErrorIs(err, assert.AnError)
	})
}

func TestNewThrottle(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
	callCount := 0
	f1 := func() {
		callCount++
	}
	th, reset := NewThrottle(100*time.Millisecond, f1)

	// NOTE: these subtests are intentionally sequential (not parallel) and share
	// `callCount`, `th` and `reset` above: each one builds on the previous one's
	// state, and the expected counts below reflect that cumulative progression.

	t.Run("only the first call within the window is worked", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		is.Zero(callCount)
		for j := 0; j < 100; j++ {
			th()
		}
		is.Equal(1, callCount)
	})

	t.Run("a new window allows another call", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		time.Sleep(150 * time.Millisecond)

		for j := 0; j < 100; j++ {
			th()
		}

		is.Equal(2, callCount)
	})

	t.Run("reset allows an immediate call", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		reset()
		th()
		is.Equal(3, callCount)
	})
}

func TestNewThrottleWithCount(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
	callCount := 0
	f1 := func() {
		callCount++
	}
	th, reset := NewThrottleWithCount(100*time.Millisecond, 3, f1)

	// NOTE: these subtests are intentionally sequential (not parallel) and share
	// `callCount`, `th` and `reset` above: each one builds on the previous one's
	// state, and the expected counts below reflect that cumulative progression.

	t.Run("does not throttle for initial count number", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		for i := 0; i < 20; i++ {
			th()
		}
		is.Equal(3, callCount)
	})

	t.Run("a new window allows another burst up to count", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		time.Sleep(150 * time.Millisecond)

		for i := 0; i < 20; i++ {
			th()
		}

		is.Equal(6, callCount)
	})

	t.Run("reset allows an immediate burst up to count", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		reset()
		for i := 0; i < 20; i++ {
			th()
		}

		is.Equal(9, callCount)
	})
}

func TestNewThrottleBy(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
	callCountA := 0
	callCountB := 0
	f1 := func(key string) {
		if key == "a" {
			callCountA++
		} else {
			callCountB++
		}
	}
	th, reset := NewThrottleBy(100*time.Millisecond, f1)

	// NOTE: these subtests are intentionally sequential (not parallel) and share
	// `callCountA`, `callCountB`, `th` and `reset` above: each one builds on the
	// previous one's state, and the expected counts below reflect that
	// cumulative progression.

	t.Run("only the first call per key within the window is worked", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		is.Zero(callCountA)
		is.Zero(callCountB)
		for j := 0; j < 100; j++ {
			th("a")
			th("b")
		}
		is.Equal(1, callCountA)
		is.Equal(1, callCountB)
	})

	t.Run("a new window allows another call per key", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		time.Sleep(150 * time.Millisecond)

		for j := 0; j < 100; j++ {
			th("a")
			th("b")
		}

		is.Equal(2, callCountA)
		is.Equal(2, callCountB)
	})

	t.Run("reset allows an immediate call for the invoked key only", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		reset()
		th("a")
		is.Equal(3, callCountA)
		is.Equal(2, callCountB)
	})
}

func TestNewThrottleByWithCount(t *testing.T) { //nolint:paralleltest
	// t.Parallel()

	callCountA := 0
	callCountB := 0
	f1 := func(key string) {
		if key == "a" {
			callCountA++
		} else {
			callCountB++
		}
	}
	th, reset := NewThrottleByWithCount(100*time.Millisecond, 3, f1)

	// NOTE: these subtests are intentionally sequential (not parallel) and share
	// `callCountA`, `callCountB`, `th` and `reset` above: each one builds on the
	// previous one's state, and the expected counts below reflect that
	// cumulative progression.

	t.Run("does not throttle for initial count number per key", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		for i := 0; i < 20; i++ {
			th("a")
			th("b")
		}
		is.Equal(3, callCountA)
		is.Equal(3, callCountB)
	})

	t.Run("a new window allows another burst up to count per key", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		time.Sleep(150 * time.Millisecond)

		for i := 0; i < 20; i++ {
			th("a")
			th("b")
		}

		is.Equal(6, callCountA)
		is.Equal(6, callCountB)
	})

	t.Run("reset allows an immediate burst up to count for the invoked key only", func(t *testing.T) { //nolint:paralleltest
		// t.Parallel()
		is := assert.New(t)

		reset()
		for i := 0; i < 20; i++ {
			th("a")
		}

		is.Equal(9, callCountA)
		is.Equal(6, callCountB)
	})
}
