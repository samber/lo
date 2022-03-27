package lo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrop(t *testing.T) {
	is := assert.New(t)

	is.Equal([]int{1, 2, 3, 4}, Drop([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{2, 3, 4}, Drop([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{3, 4}, Drop([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{4}, Drop([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, Drop([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, Drop([]int{0, 1, 2, 3, 4}, 6))
}

func TestDropRight(t *testing.T) {
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, DropRight([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{0, 1, 2}, DropRight([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{0, 1}, DropRight([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{0}, DropRight([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, DropRight([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, DropRight([]int{0, 1, 2, 3, 4}, 6))
}

func TestDropWhile(t *testing.T) {
	is := assert.New(t)

	is.Equal([]int{4, 5, 6}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 4
	}))

	is.Equal([]int{}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return true
	}))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t == 10
	}))
}

func TestDropRightWhile(t *testing.T) {
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 3
	}))

	is.Equal([]int{0, 1}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 1
	}))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t == 10
	}))

	is.Equal([]int{}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 10
	}))
}
