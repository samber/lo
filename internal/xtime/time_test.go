package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFakeClock(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	base := time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)
	c := NewFakeClockAt(base)

	// Now returns the fixed time, and does not advance on its own.
	is.Equal(base, c.Now())
	is.Equal(base, c.Now())

	// Since is measured against the fake's current time.
	is.Equal(2*time.Hour, c.Since(base.Add(-2*time.Hour)))

	// Until is measured towards the fake's current time.
	is.Equal(3*time.Hour, c.Until(base.Add(3*time.Hour)))

	// Sleep advances the fake clock rather than blocking.
	c.Sleep(90 * time.Minute)
	is.Equal(base.Add(90*time.Minute), c.Now())
	is.Equal(90*time.Minute, c.Since(base))
}

func TestNewFakeClock(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	before := time.Now()
	c := NewFakeClock()
	after := time.Now()

	// NewFakeClock seeds with the current time, so Now must fall within
	// the window observed around its construction.
	now := c.Now()
	is.False(now.Before(before))
	is.False(now.After(after))
}

func TestRealClock(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	c := NewRealClock()

	// Now tracks wall-clock time.
	before := time.Now()
	got := c.Now()
	after := time.Now()
	is.False(got.Before(before))
	is.False(got.After(after))

	// Since/Until stay consistent with the standard library within a
	// generous tolerance to avoid flakiness on slow CI.
	past := time.Now().Add(-time.Hour)
	is.InDelta(time.Hour, c.Since(past), float64(time.Second))

	future := time.Now().Add(time.Hour)
	is.InDelta(time.Hour, c.Until(future), float64(time.Second))

	// Sleep actually waits at least the requested duration.
	start := time.Now()
	c.Sleep(5 * time.Millisecond)
	is.GreaterOrEqual(time.Since(start), 5*time.Millisecond)
}

// Not parallel: mutates the package-global clock via SetClock.
func TestClockPackageFuncs(t *testing.T) { //nolint:paralleltest
	is := assert.New(t)

	// The package-level functions delegate to the currently installed clock.
	// Swap in a fake, then restore the real clock so other tests are unaffected.
	original := clock
	t.Cleanup(func() { SetClock(original) })

	base := time.Date(2024, time.June, 15, 8, 30, 0, 0, time.UTC)
	SetClock(NewFakeClockAt(base))

	is.Equal(base, Now())
	is.Equal(time.Hour, Since(base.Add(-time.Hour)))
	is.Equal(2*time.Hour, Until(base.Add(2*time.Hour)))

	Sleep(45 * time.Minute)
	is.Equal(base.Add(45*time.Minute), Now())
}

// Not parallel: mutates the package-global clock via SetClock.
func TestSetClock(t *testing.T) { //nolint:paralleltest
	is := assert.New(t)

	original := clock
	t.Cleanup(func() { SetClock(original) })

	fake := NewFakeClockAt(time.Unix(0, 0).UTC())
	SetClock(fake)
	is.Equal(time.Unix(0, 0).UTC(), Now())

	SetClock(NewRealClock())
	is.WithinDuration(time.Now(), Now(), time.Second)
}
