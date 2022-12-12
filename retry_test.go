package lo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAttempt(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	err := fmt.Errorf("failed")

	iter1, err1 := Attempt(42, func(i int) error {
		return nil
	})
	iter2, err2 := Attempt(42, func(i int) error {
		if i == 5 {
			return nil
		}

		return err
	})
	iter3, err3 := Attempt(2, func(i int) error {
		if i == 5 {
			return nil
		}

		return err
	})
	iter4, err4 := Attempt(0, func(i int) error {
		if i < 42 {
			return err
		}

		return nil
	})

	is.Equal(iter1, 1)
	is.Equal(err1, nil)
	is.Equal(iter2, 6)
	is.Equal(err2, nil)
	is.Equal(iter3, 2)
	is.Equal(err3, err)
	is.Equal(iter4, 43)
	is.Equal(err4, nil)
}

func TestAttemptWithDelay(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	err := fmt.Errorf("failed")

	iter1, dur1, err1 := AttemptWithDelay(42, 10*time.Millisecond, func(i int, d time.Duration) error {
		return nil
	})
	iter2, dur2, err2 := AttemptWithDelay(42, 10*time.Millisecond, func(i int, d time.Duration) error {
		if i == 5 {
			return nil
		}

		return err
	})
	iter3, dur3, err3 := AttemptWithDelay(2, 10*time.Millisecond, func(i int, d time.Duration) error {
		if i == 5 {
			return nil
		}

		return err
	})
	iter4, dur4, err4 := AttemptWithDelay(0, 10*time.Millisecond, func(i int, d time.Duration) error {
		if i < 10 {
			return err
		}

		return nil
	})

	is.Equal(iter1, 1)
	is.Greater(dur1, 0*time.Millisecond)
	is.Less(dur1, 1*time.Millisecond)
	is.Equal(err1, nil)
	is.Equal(iter2, 6)
	is.Greater(dur2, 50*time.Millisecond)
	is.Less(dur2, 60*time.Millisecond)
	is.Equal(err2, nil)
	is.Equal(iter3, 2)
	is.Greater(dur3, 10*time.Millisecond)
	is.Less(dur3, 20*time.Millisecond)
	is.Equal(err3, err)
	is.Equal(iter4, 11)
	is.Greater(dur4, 100*time.Millisecond)
	is.Less(dur4, 115*time.Millisecond)
	is.Equal(err4, nil)
}

func TestAttemptWhile(t *testing.T) {
	is := assert.New(t)

	err := fmt.Errorf("failed")

	iter1, err1 := AttemptWhile(42, func(i int) (error, bool) {
		return nil, true
	})

	is.Equal(iter1, 1)
	is.Nil(err1)

	iter2, err2 := AttemptWhile(42, func(i int) (error, bool) {
		if i == 5 {
			return nil, true
		}

		return err, true
	})

	is.Equal(iter2, 6)
	is.Nil(err2)

	iter3, err3 := AttemptWhile(2, func(i int) (error, bool) {
		if i == 5 {
			return nil, true
		}

		return err, true
	})

	is.Equal(iter3, 2)
	is.Equal(err3, err)

	iter4, err4 := AttemptWhile(0, func(i int) (error, bool) {
		if i < 42 {
			return err, true
		}

		return nil, true
	})

	is.Equal(iter4, 43)
	is.Nil(err4)

	iter5, err5 := AttemptWhile(0, func(i int) (error, bool) {
		if i == 5 {
			return nil, false
		}

		return err, true
	})

	is.Equal(iter5, 6)
	is.Nil(err5)

	iter6, err6 := AttemptWhile(0, func(i int) (error, bool) {
		return nil, false
	})

	is.Equal(iter6, 1)
	is.Nil(err6)

	iter7, err7 := AttemptWhile(42, func(i int) (error, bool) {
		if i == 42 {
			return nil, false
		}
		if i < 41 {
			return err, true
		}

		return nil, true
	})

	is.Equal(iter7, 42)
	is.Nil(err7)
}

func TestAttemptWhileWithDelay(t *testing.T) {
	is := assert.New(t)

	err := fmt.Errorf("failed")

	iter1, dur1, err1 := AttemptWhileWithDelay(42, 10*time.Millisecond, func(i int, d time.Duration) (error, bool) {
		return nil, true
	})

	is.Equal(iter1, 1)
	is.Greater(dur1, 0*time.Millisecond)
	is.Less(dur1, 1*time.Millisecond)
	is.Nil(err1)

	iter2, dur2, err2 := AttemptWhileWithDelay(42, 10*time.Millisecond, func(i int, d time.Duration) (error, bool) {
		if i == 5 {
			return nil, true
		}

		return err, true
	})

	is.Equal(iter2, 6)
	is.Greater(dur2, 50*time.Millisecond)
	is.Less(dur2, 60*time.Millisecond)
	is.Nil(err2)

	iter3, dur3, err3 := AttemptWhileWithDelay(2, 10*time.Millisecond, func(i int, d time.Duration) (error, bool) {
		if i == 5 {
			return nil, true
		}

		return err, true
	})

	is.Equal(iter3, 2)
	is.Greater(dur3, 10*time.Millisecond)
	is.Less(dur3, 20*time.Millisecond)
	is.Equal(err3, err)

	iter4, dur4, err4 := AttemptWhileWithDelay(0, 10*time.Millisecond, func(i int, d time.Duration) (error, bool) {
		if i < 10 {
			return err, true
		}

		return nil, true
	})

	is.Equal(iter4, 11)
	is.Greater(dur4, 100*time.Millisecond)
	is.Less(dur4, 115*time.Millisecond)
	is.Nil(err4)

	iter5, dur5, err5 := AttemptWhileWithDelay(0, 10*time.Millisecond, func(i int, d time.Duration) (error, bool) {
		if i == 5 {
			return nil, false
		}

		return err, true
	})

	is.Equal(iter5, 6)
	is.Greater(dur5, 10*time.Millisecond)
	is.Less(dur5, 115*time.Millisecond)
	is.Nil(err5)

	iter6, dur6, err6 := AttemptWhileWithDelay(0, 10*time.Millisecond, func(i int, d time.Duration) (error, bool) {
		return nil, false
	})

	is.Equal(iter6, 1)
	is.Less(dur6, 10*time.Millisecond)
	is.Less(dur6, 115*time.Millisecond)
	is.Nil(err6)

	iter7, dur7, err7 := AttemptWhileWithDelay(42, 10*time.Millisecond, func(i int, d time.Duration) (error, bool) {
		if i == 42 {
			return nil, false
		}
		if i < 41 {
			return err, true
		}

		return nil, true
	})

	is.Equal(iter7, 42)
	is.Less(dur7, 500*time.Millisecond)
	is.Nil(err7)
}

func TestDebounce(t *testing.T) {
	t.Parallel()

	f1 := func() {
		println("1. Called once after 10ms when func stopped invoking!")
	}
	f2 := func() {
		println("2. Called once after 10ms when func stopped invoking!")
	}
	f3 := func() {
		println("3. Called once after 10ms when func stopped invoking!")
	}

	d1, _ := NewDebounce(10*time.Millisecond, f1)

	// execute 3 times
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			d1()
		}
		time.Sleep(20 * time.Millisecond)
	}

	d2, _ := NewDebounce(10*time.Millisecond, f2)

	// execute once because it is always invoked and only last invoke is worked after 100ms
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			d2()
		}
		time.Sleep(5 * time.Millisecond)
	}

	time.Sleep(10 * time.Millisecond)

	// execute once because it is canceled after 200ms.
	d3, cancel := NewDebounce(10*time.Millisecond, f3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			d3()
		}
		time.Sleep(20 * time.Millisecond)
		if i == 0 {
			cancel()
		}
	}
}

func TestTransation(t *testing.T) {
	is := assert.New(t)

	// no error
	{
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
		is.Equal(nil, err)
	}

	// with error
	{
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
		is.Equal(assert.AnError, err)
	}

	// with error + update value
	{
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
		is.Equal(assert.AnError, err)
	}
}
