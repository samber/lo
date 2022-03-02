package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	is := assert.New(t)

	result1 := IndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := IndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}

func TestLastIndexOf(t *testing.T) {
	is := assert.New(t)

	result1 := LastIndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := LastIndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 4)
	is.Equal(result2, -1)
}

func TestFind(t *testing.T) {
	is := assert.New(t)

	result1, ok1 := Find[string]([]string{"a", "b", "c", "d"}, func(i string) bool {
		return i == "b"
	})
	result2, ok2 := Find[string]([]string{"foobar"}, func(i string) bool {
		return i == "b"
	})

	is.Equal(ok1, true)
	is.Equal(result1, "b")
	is.Equal(ok2, false)
	is.Equal(result2, "")
}

func TestMin(t *testing.T) {
	is := assert.New(t)

	result1 := Min[int]([]int{1, 2, 3})
	result2 := Min[int]([]int{3, 2, 1})
	result3 := Min[int]([]int{})

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 0)
}

func TestMax(t *testing.T) {
	is := assert.New(t)

	result1 := Max[int]([]int{1, 2, 3})
	result2 := Max[int]([]int{3, 2, 1})
	result3 := Max[int]([]int{})

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, 0)
}

func TestLast(t *testing.T) {
	is := assert.New(t)

	result1 := Last[int]([]int{1, 2, 3})

	is.Equal(result1, 3)
}

func TestNth(t *testing.T) {
	is := assert.New(t)

	result1 := Nth[int]([]int{0, 1, 2, 3}, 2)
	result2 := Nth[int]([]int{0, 1, 2, 3}, -2)

	is.Equal(result1, 2)
	is.Equal(result2, 2)
}
