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

func ExampleTransaction() {
	transaction := NewTransaction[int]().
		Then(
			func(state int) (int, error) {
				fmt.Println("step 1")
				return state + 10, nil
			},
			func(state int) int {
				fmt.Println("rollback 1")
				return state - 10
			},
		).
		Then(
			func(state int) (int, error) {
				fmt.Println("step 2")
				return state + 15, nil
			},
			func(state int) int {
				fmt.Println("rollback 2")
				return state - 15
			},
		).
		Then(
			func(state int) (int, error) {
				fmt.Println("step 3")

				if true {
					return state, fmt.Errorf("error")
				}

				return state + 42, nil
			},
			func(state int) int {
				fmt.Println("rollback 3")
				return state - 42
			},
		)

	_, _ = transaction.Process(-5)

	// Output:
	// step 1
	// step 2
	// step 3
	// rollback 2
	// rollback 1
}

func ExampleTransaction_ok() {
	transaction := NewTransaction[int]().
		Then(
			func(state int) (int, error) {
				return state + 10, nil
			},
			func(state int) int {
				return state - 10
			},
		).
		Then(
			func(state int) (int, error) {
				return state + 15, nil
			},
			func(state int) int {
				return state - 15
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

	state, err := transaction.Process(-5)

	fmt.Println(state)
	fmt.Println(err)
	// Output:
	// 62
	// <nil>
}

func ExampleTransaction_error() {
	transaction := NewTransaction[int]().
		Then(
			func(state int) (int, error) {
				return state + 10, nil
			},
			func(state int) int {
				return state - 10
			},
		).
		Then(
			func(state int) (int, error) {
				return state, fmt.Errorf("error")
			},
			func(state int) int {
				return state - 15
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

	state, err := transaction.Process(-5)

	fmt.Println(state)
	fmt.Println(err)
	// Output:
	// -5
	// error
}
