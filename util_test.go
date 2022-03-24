package lo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
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
	result4 := RangeFrom[float64](2.0, 3)
	result5 := RangeFrom[float64](-2.0, -3)
	is.Equal(result1, []int{1, 2, 3, 4, 5})
	is.Equal(result2, []int{-1, -2, -3, -4, -5})
	is.Equal(result3, []int{})
	is.Equal(result4, []float64{2.0, 3.0, 4.0})
	is.Equal(result5, []float64{-2.0, -3.0, -4.0})
}

func TestRangeClose(t *testing.T) {
	is := assert.New(t)
	result1 := RangeWithSteps(0, 20, 6)
	result2 := RangeWithSteps(0, 3, -5)
	result3 := RangeWithSteps(1, 1, 0)
	result4 := RangeWithSteps(3, 2, 1)
	result5 := RangeWithSteps[float64](1.0, 4.0, 2.0)
	result6 := RangeWithSteps[float32](-1.0, -4.0, -1.0)
	is.Equal([]int{0, 6, 12, 18}, result1)
	is.Equal([]int{}, result2)
	is.Equal([]int{}, result3)
	is.Equal([]int{}, result4)
	is.Equal([]float64{1.0, 3.0}, result5)
	is.Equal([]float32{-1.0, -2.0, -3.0}, result6)
}

func TestMust(t *testing.T) {
	is := assert.New(t)
	is.Equal("foo", Must("foo", nil))
	is.Panics(func() {
		Must("", errors.New("something went wrong"))
	})
}
