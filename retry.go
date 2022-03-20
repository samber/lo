package lo

import (
	"sync"
	"time"
)

type debounce struct {
	after     time.Duration
	mu        *sync.Mutex
	timer     *time.Timer
	done      bool
	callbacks []func()
}

func (d *debounce) reset() *debounce {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.done {
		return d
	}

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.after, func() {
		for _, f := range d.callbacks {
			f()
		}
	})
	return d
}

func (d *debounce) cancel() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
		d.timer = nil
	}

	d.done = true
}

// NewDebounce creates a debounced instance that delays invoking functions given until after wait milliseconds have elapsed.
func NewDebounce(duration time.Duration, f ...func()) (func(), func()) {
	d := &debounce{
		after:     duration,
		mu:        new(sync.Mutex),
		timer:     nil,
		done:      false,
		callbacks: f,
	}

	return func() {
		d.reset()
	}, d.cancel
}

// Attempt invokes a function N times until it returns valid output. Returning either the caught error or nil. When first argument is less than `1`, the function runs until a sucessfull response is returned.
func Attempt(maxIteration int, f func(int) error) (int, error) {
	var err error

	for i := 0; maxIteration <= 0 || i < maxIteration; i++ {
		// for retries >= 0 {
		err = f(i)
		if err == nil {
			return i + 1, nil
		}
	}

	return maxIteration, err
}

// throttle ?
