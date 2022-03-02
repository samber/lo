package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	is := assert.New(t)

	result1 := Contains[int]([]int{0, 1, 2, 3, 4, 5}, 5)
	result2 := Contains[int]([]int{0, 1, 2, 3, 4, 5}, 6)

	is.Equal(result1, true)
	is.Equal(result2, false)
}

func TestSome(t *testing.T) {
	is := assert.New(t)

	result1 := Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})

	is.Equal(result1, true)
	is.Equal(result2, true)
	is.Equal(result3, false)
}

func TestEvery(t *testing.T) {
	is := assert.New(t)

	result1 := Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})

	is.Equal(result1, true)
	is.Equal(result2, false)
	is.Equal(result3, false)
}

func TestIntersect(t *testing.T) {
	is := assert.New(t)

	result1 := Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := Intersect[int]([]int{0, 6}, []int{0, 1, 2, 3, 4, 5})
	result5 := Intersect[int]([]int{0, 6, 0}, []int{0, 1, 2, 3, 4, 5})

	is.Equal(result1, []int{0, 2})
	is.Equal(result2, []int{0})
	is.Equal(result3, []int{})
	is.Equal(result4, []int{0})
	is.Equal(result5, []int{0})
}

func TestDifference(t *testing.T) {
	is := assert.New(t)

	left1, right1 := Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	is.Equal(left1, []int{1, 3, 4, 5})
	is.Equal(right1, []int{6})

	left2, right2 := Difference[int]([]int{1, 2, 3, 4, 5}, []int{0, 6})
	is.Equal(left2, []int{1, 2, 3, 4, 5})
	is.Equal(right2, []int{0, 6})

	left3, right3 := Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5})
	is.Equal(left3, []int{})
	is.Equal(right3, []int{})
}
