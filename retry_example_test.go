//go:build !race
// +build !race

package lo

import (
	"fmt"
	"sync"
	"time"
)

func ExampleNewDebounce() {
	i := 0
	calls := []int{}
	mu := sync.Mutex{}

	debounce, cancel := NewDebounce(time.Millisecond, func() {
		mu.Lock()
		defer mu.Unlock()
		calls = append(calls, i)
	})

	debounce()
	i++

	time.Sleep(5 * time.Millisecond)

	debounce()
	i++
	debounce()
	i++
	debounce()
	i++

	time.Sleep(5 * time.Millisecond)

	cancel()

	fmt.Printf("%v", calls)
	// Output: [1 4]
}

func ExampleAttempt() {
	count1, err1 := Attempt(2, func(i int) error {
		if i == 0 {
			return fmt.Errorf("error")
		}

		return nil
	})

	count2, err2 := Attempt(2, func(i int) error {
		if i < 10 {
			return fmt.Errorf("error")
		}

		return nil
	})

	fmt.Printf("%v %v\n", count1, err1)
	fmt.Printf("%v %v\n", count2, err2)
	// Output:
	// 2 <nil>
	// 2 error
}

func ExampleAttemptWithDelay() {
	count1, time1, err1 := AttemptWithDelay(2, time.Millisecond, func(i int, _ time.Duration) error {
		if i == 0 {
			return fmt.Errorf("error")
		}

		return nil
	})

	count2, time2, err2 := AttemptWithDelay(2, time.Millisecond, func(i int, _ time.Duration) error {
		if i < 10 {
			return fmt.Errorf("error")
		}

		return nil
	})

	fmt.Printf("%v %v %v\n", count1, time1.Truncate(time.Millisecond), err1)
	fmt.Printf("%v %v %v\n", count2, time2.Truncate(time.Millisecond), err2)
	// Output:
	// 2 1ms <nil>
	// 2 1ms error
}
