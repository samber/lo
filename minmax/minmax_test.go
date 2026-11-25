package minmax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	is := assert.New(t)

	result1 := Max([]int{1, 2, 3}...)
	result2 := Max([]int{3, 2, 1}...)
	result3 := Max([]int{}...)
	result4 := Max(5, 1)
	result5 := Max(3)
	result6 := Max(-123, 23, 0)

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, 0)
	is.Equal(result4, 5)
	is.Equal(result5, 3)
	is.Equal(result6, 23)
}

func TestMin(t *testing.T) {
	is := assert.New(t)

	result1 := Min([]int{1, 2, 3}...)
	result2 := Min([]int{3, 2, 1}...)
	result3 := Min([]int{}...)
	result4 := Min(1, 2)
	result5 := Min(1)
	result6 := Min(5, 2, 0)

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 0)
	is.Equal(result4, 1)
	is.Equal(result5, 1)
	is.Equal(result6, 0)
}
