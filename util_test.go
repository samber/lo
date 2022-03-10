package lo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	is := assert.New(t)
	result1 := Range(4)
	result2 := Range(-4)
	result3 := Range(0)
	is.Equal(result1, []int{0, 1, 2, 3})
	is.Equal(result2, []int{0, -1, -2, -3})
	is.Equal(result3, []int{})
}

func TestRangeFrom(t *testing.T) {
	is := assert.New(t)
	result1 := RangeFrom(1, 5)
	result2 := RangeFrom(-1, -5)
	result3 := RangeFrom(10, 0)
	is.Equal(result1, []int{1, 2, 3, 4, 5})
	is.Equal(result2, []int{-1, -2, -3, -4, -5})
	is.Equal(result3, []int{})
}

func TestRangeClose(t *testing.T) {
	is := assert.New(t)
	result1 := RangeOpen(0, 20, 6)
	result2 := RangeOpen(0, 3, -5)
	result3 := RangeOpen(0, -4, -1)
	result4 := RangeOpen(1, 4, 0)
	result5 := RangeOpen(1, 1, 0)
	is.Equal(result1, []int{0, 6, 12, 18})
	is.Equal(result2, []int{0, 1, 2})
	is.Equal(result3, []int{0, -1, -2, -3})
	is.Equal(result4, []int{1, 2, 3})
	is.Equal(len(result5), 0)
}
