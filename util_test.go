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
	result4 := RangeWithSteps[float64](1.0, 4.0, 2.0)
	result5 := RangeWithSteps[float32](-1.0, -4.0, -1.0)
	is.Equal(result1, []int{0, 6, 12, 18})
	is.Equal(result2, []int{})
	is.Equal(result3, []int{})
	is.Equal(result4, []float64{1.0, 3.0})
	is.Equal(result5, []float32{-1.0, -2.0, -3.0})
}
