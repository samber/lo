package lo

import (
	"sync"
	"time"
)

type debounce struct {
	after time.Duration
	mu    sync.Mutex
	timer *time.Timer
}

func (d *debounce) register(f func()) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(d.after, f)
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

// NewDebounce creates a debounced instance that delays invoking func given until after wait milliseconds have elapsed.
func NewDebounce(duration time.Duration) func(f ...func()) {
	d := debounce{
		after: duration,
	}
	return func(f ...func()) {
		if len(f) == 0 {
			return
		}
		for i := 0; i < len(f); i++ {
			d.register(f[i])
		}
	}
}
