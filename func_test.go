package lo

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartial(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := Partial(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}

func TestPartial1(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := Partial1(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}

func TestPartial2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int) string {
		return strconv.Itoa(int(x) + y + z)
	}
	f := Partial2(add, 5)
	is.Equal("24", f(10, 9))
	is.Equal("8", f(-5, 8))
}

func TestPartial3(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32) string {
		return strconv.Itoa(int(x) + y + z + int(a))
	}
	f := Partial3(add, 5)
	is.Equal("21", f(10, 9, -3))
	is.Equal("15", f(-5, 8, 7))
}

func TestPartial4(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32, b int32) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b))
	}
	f := Partial4(add, 5)
	is.Equal("21", f(10, 9, -3, 0))
	is.Equal("14", f(-5, 8, 7, -1))
}

func TestPartial5(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32, b int32, c int) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b) + c)
	}
	f := Partial5(add, 5)
	is.Equal("26", f(10, 9, -3, 0, 5))
	is.Equal("21", f(-5, 8, 7, -1, 7))
}

func TestPartialRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := PartialRight(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}

func TestPartialRight1(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := PartialRight1(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}

func TestPartialRight2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int) string {
		return strconv.Itoa(int(x) + y + z)
	}
	f := PartialRight2(add, 5)
	is.Equal("24", f(10, 9))
	is.Equal("8", f(-5, 8))
}

func TestPartialRight3(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32) string {
		return strconv.Itoa(int(x) + y + z + int(a))
	}
	f := PartialRight3(add, 5)
	is.Equal("21", f(10, 9, -3))
	is.Equal("15", f(-5, 8, 7))
}

func TestPartialRight4(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32, b int32) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b))
	}
	f := PartialRight4(add, 5)
	is.Equal("21", f(10, 9, -3, 0))
	is.Equal("14", f(-5, 8, 7, -1))
}

func TestPartialRight5(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32, b int32, c int) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b) + c)
	}
	f := PartialRight5(add, 5)
	is.Equal("26", f(10, 9, -3, 0, 5))
	is.Equal("21", f(-5, 8, 7, -1, 7))
}

func TestNot(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	isOdd := func(i int) bool {
		return i%2 == 1
	}

	is.False(Not(isOdd)(2))
	is.True(Not(isOdd)(3))
}

func TestNot2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.False(Not2(strings.Contains)("test", "t"))
	is.True(Not2(strings.Contains)("test", "p"))
}
