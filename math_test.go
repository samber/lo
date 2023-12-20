package lo_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Range(4)
	result2 := lo.Range(-4)
	result3 := lo.Range(0)
	is.Equal(result1, []int{0, 1, 2, 3})
	is.Equal(result2, []int{0, -1, -2, -3})
	is.Equal(result3, []int{})
}

func TestRangeFrom(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.RangeFrom(1, 5)
	result2 := lo.RangeFrom(-1, -5)
	result3 := lo.RangeFrom(10, 0)
	result4 := lo.RangeFrom(2.0, 3)
	result5 := lo.RangeFrom(-2.0, -3)
	is.Equal(result1, []int{1, 2, 3, 4, 5})
	is.Equal(result2, []int{-1, -2, -3, -4, -5})
	is.Equal(result3, []int{})
	is.Equal(result4, []float64{2.0, 3.0, 4.0})
	is.Equal(result5, []float64{-2.0, -3.0, -4.0})
}

func TestRangeClose(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.RangeWithSteps(0, 20, 6)
	result2 := lo.RangeWithSteps(0, 3, -5)
	result3 := lo.RangeWithSteps(1, 1, 0)
	result4 := lo.RangeWithSteps(3, 2, 1)
	result5 := lo.RangeWithSteps(1.0, 4.0, 2.0)
	result6 := lo.RangeWithSteps[float32](-1.0, -4.0, -1.0)
	is.Equal([]int{0, 6, 12, 18}, result1)
	is.Equal([]int{}, result2)
	is.Equal([]int{}, result3)
	is.Equal([]int{}, result4)
	is.Equal([]float64{1.0, 3.0}, result5)
	is.Equal([]float32{-1.0, -2.0, -3.0}, result6)
}

func TestClamp(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Clamp(0, -10, 10)
	result2 := lo.Clamp(-42, -10, 10)
	result3 := lo.Clamp(42, -10, 10)

	is.Equal(result1, 0)
	is.Equal(result2, -10)
	is.Equal(result3, 10)
}

func TestSum(t *testing.T) {
	is := assert.New(t)

	result1 := lo.Sum([]float32{2.3, 3.3, 4, 5.3})
	result2 := lo.Sum([]int32{2, 3, 4, 5})
	result3 := lo.Sum([]uint32{2, 3, 4, 5})
	result4 := lo.Sum([]uint32{})
	result5 := lo.Sum([]complex128{4_4, 2_2})

	is.Equal(result1, float32(14.900001))
	is.Equal(result2, int32(14))
	is.Equal(result3, uint32(14))
	is.Equal(result4, uint32(0))
	is.Equal(result5, complex128(6_6))
}

func TestSumBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.SumBy([]float32{2.3, 3.3, 4, 5.3}, func(n float32) float32 { return n })
	result2 := lo.SumBy([]int32{2, 3, 4, 5}, func(n int32) int32 { return n })
	result3 := lo.SumBy([]uint32{2, 3, 4, 5}, func(n uint32) uint32 { return n })
	result4 := lo.SumBy([]uint32{}, func(n uint32) uint32 { return n })
	result5 := lo.SumBy([]complex128{4_4, 2_2}, func(n complex128) complex128 { return n })

	is.Equal(result1, float32(14.900001))
	is.Equal(result2, int32(14))
	is.Equal(result3, uint32(14))
	is.Equal(result4, uint32(0))
	is.Equal(result5, complex128(6_6))
}
