package lo

import (
	"strconv"
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

func TestCompose(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	filterEven := func(ns []int) []int {
		return Filter(ns, func(x int, i int) bool { return x%2 == 0 })
	}
	sum := func(ns []int) int {
		return ReduceRight(ns, func(sum int, x int, i int) int { return sum + x }, 0)
	}
	result := Compose(filterEven, sum)([]int{1, 2, 3, 4, 5})
	is.Equal("6", strconv.Itoa(result))
}

func TestCompose2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	filterEven := func(ns []int) []int {
		return Filter(ns, func(x int, i int) bool { return x%2 == 0 })
	}
	sum := func(ns []int) int {
		return ReduceRight(ns, func(sum int, x int, i int) int { return sum + x }, 0)
	}
	result := Compose2(filterEven, sum)([]int{1, 2, 3, 4, 5})
	is.Equal("6", strconv.Itoa(result))
}

func TestCompose3(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	filterEven := func(ns []int) []int {
		return Filter(ns, func(x int, i int) bool { return x%2 == 0 })
	}
	mulEachBy10 := func(ns []int) []int {
		return Map(ns, func(x int, i int) int { return x * 10 })
	}
	sum := func(ns []int) int {
		return ReduceRight(ns, func(sum int, x int, i int) int { return sum + x }, 0)
	}
	result := Compose3(
		filterEven,
		mulEachBy10,
		sum,
	)([]int{1, 2, 3, 4, 5})
	is.Equal("60", strconv.Itoa(result))
}
