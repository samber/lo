package lo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result := Duration(func() { time.Sleep(10 * time.Millisecond) })
	is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
}

func TestDurationX(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	{
		result := Duration0(func() { time.Sleep(10 * time.Millisecond) })
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
	}

	{
		a, result := Duration1(func() string { time.Sleep(10 * time.Millisecond); return "a" })
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
	}

	{
		a, b, result := Duration2(func() (string, string) { time.Sleep(10 * time.Millisecond); return "a", "b" })
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
	}

	{
		a, b, c, result := Duration3(func() (string, string, string) { time.Sleep(10 * time.Millisecond); return "a", "b", "c" })
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
	}

	{
		a, b, c, d, result := Duration4(func() (string, string, string, string) { time.Sleep(10 * time.Millisecond); return "a", "b", "c", "d" })
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
	}

	{
		a, b, c, d, e, result := Duration5(func() (string, string, string, string, string) {
			time.Sleep(10 * time.Millisecond)
			return "a", "b", "c", "d", "e"
		})
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
	}

	{
		a, b, c, d, e, f, result := Duration6(func() (string, string, string, string, string, string) {
			time.Sleep(10 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f"
		})
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
	}

	{
		a, b, c, d, e, f, g, result := Duration7(func() (string, string, string, string, string, string, string) {
			time.Sleep(10 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f", "g"
		})
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
		is.Equal("g", g)
	}

	{
		a, b, c, d, e, f, g, h, result := Duration8(func() (string, string, string, string, string, string, string, string) {
			time.Sleep(10 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f", "g", "h"
		})
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
		is.Equal("g", g)
		is.Equal("h", h)
	}

	{
		a, b, c, d, e, f, g, h, i, result := Duration9(func() (string, string, string, string, string, string, string, string, string) {
			time.Sleep(10 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f", "g", "h", "i"
		})
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
		is.Equal("a", a)
		is.Equal("b", b)
		is.Equal("c", c)
		is.Equal("d", d)
		is.Equal("e", e)
		is.Equal("f", f)
		is.Equal("g", g)
		is.Equal("h", h)
		is.Equal("i", i)
	}

	{
		a, b, c, d, e, f, g, h, i, j, result := Duration10(func() (string, string, string, string, string, string, string, string, string, string) {
			time.Sleep(10 * time.Millisecond)
			return "a", "b", "c", "d", "e", "f", "g", "h", "i", "j"
		})
		is.InEpsilon(10*time.Millisecond, result, float64(2*time.Millisecond))
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
	}
}
