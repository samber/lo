package lo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
	is := assert.New(t)
	testWithTimeout(t, 200*time.Millisecond)

	result := Duration(func() { time.Sleep(100 * time.Millisecond) })
	is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
}

func TestDurationX(t *testing.T) { //nolint:paralleltest
	// t.Parallel()
	is := assert.New(t)
	testWithTimeout(t, 1500*time.Millisecond)

	t.Run("Duration0", func(t *testing.T) { //nolint:paralleltest
		result := Duration0(func() { time.Sleep(100 * time.Millisecond) })
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
	})

	t.Run("Duration1", func(t *testing.T) { //nolint:paralleltest
		a, result := Duration1(func() string { time.Sleep(100 * time.Millisecond); return "a" })
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
	})

	t.Run("Duration2", func(t *testing.T) { //nolint:paralleltest
		a, b, result := Duration2(func() (string, string) { time.Sleep(100 * time.Millisecond); return "a", "b" })
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
	})

	t.Run("Duration3", func(t *testing.T) { //nolint:paralleltest
		a, b, c, result := Duration3(func() (string, string, string) { time.Sleep(100 * time.Millisecond); return "a", "b", "c" })
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
	})

	t.Run("Duration4", func(t *testing.T) { //nolint:paralleltest
		a, b, c, d, result := Duration4(func() (string, string, string, string) {
			time.Sleep(100 * time.Millisecond)
			return "a", "b", "c", "d"
		})
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
	})

	t.Run("Duration5", func(t *testing.T) { //nolint:paralleltest
		a, b, c, d, e, result := Duration5(func() (string, string, string, string, string) {
			time.Sleep(100 * time.Millisecond)
			return "a", "b", "c", "d", "e"
		})
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
	})

	t.Run("Duration6", func(t *testing.T) { //nolint:paralleltest
		a, b, c, d, e, f, result := Duration6(func() (string, string, string, string, string, string) {
			time.Sleep(100 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f"
		})
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
	})

	t.Run("Duration7", func(t *testing.T) { //nolint:paralleltest
		a, b, c, d, e, f, g, result := Duration7(func() (string, string, string, string, string, string, string) {
			time.Sleep(100 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f", "g"
		})
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
		is.Equal("g", g)
	})

	t.Run("Duration8", func(t *testing.T) { //nolint:paralleltest
		a, b, c, d, e, f, g, h, result := Duration8(func() (string, string, string, string, string, string, string, string) {
			time.Sleep(100 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f", "g", "h"
		})
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
		is.Equal("g", g)
		is.Equal("h", h)
	})

	t.Run("Duration9", func(t *testing.T) { //nolint:paralleltest
		a, b, c, d, e, f, g, h, i, result := Duration9(func() (string, string, string, string, string, string, string, string, string) {
			time.Sleep(100 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f", "g", "h", "i"
		})
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
		is.Equal("g", g)
		is.Equal("h", h)
		is.Equal("i", i)
	})

	t.Run("Duration10", func(t *testing.T) { //nolint:paralleltest
		a, b, c, d, e, f, g, h, i, j, result := Duration10(func() (string, string, string, string, string, string, string, string, string, string) {
			time.Sleep(100 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f", "g", "h", "i", "j"
		})
		is.InDelta(100*time.Millisecond, result, float64(20*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
		is.Equal("g", g)
		is.Equal("h", h)
		is.Equal("i", i)
		is.Equal("j", j)
	})
}
