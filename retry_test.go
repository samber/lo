package lo

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAttempt(t *testing.T) {
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
			return fmt.Errorf("failed")
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

func TestDebounce(t *testing.T) {
	is := assert.New(t)
	var (
		counter1 uint64
		counter2 uint64
	)
	f1 := func() {
		atomic.AddUint64(&counter1, 1)
	}
	f2 := func() {
		atomic.AddUint64(&counter2, 1)
	}

	debounced := NewDebounce(100 * time.Millisecond)

	for i := 0; i < 3; i++ {
		// just execute one time in 100 milliseconds
		for j := 0; j < 10; j++ {
			debounced(f1)
		}
		time.Sleep(200 * time.Millisecond)
	}
	for i := 0; i < 3; i++ {
		debounced(f2, f2, f2, f2, f2, f2)
		time.Sleep(200 * time.Millisecond)
	}
	result1 := int(atomic.LoadUint64(&counter1))
	result2 := int(atomic.LoadUint64(&counter2))
	is.Equal(result1, 3)
	is.Equal(result2, 3)
}
