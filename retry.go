package lo

import (
	"sync"
	"time"
)

type Debounce struct {
	after time.Duration
	mu    *sync.Mutex
	timer *time.Timer
	done  bool
}

func (d *Debounce) Add(f func()) *Debounce {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.done {
		return d
	}
	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(d.after, f)
	return d
}

func (d *Debounce) Cancel() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.timer != nil {
		d.timer.Stop()
		d.timer = nil
	}
	d.done = true
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
func NewDebounce(duration time.Duration) *Debounce {
	return &Debounce{
		mu:    new(sync.Mutex),
		after: duration,
	}
}
