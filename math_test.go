package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Range(4)
	result2 := Range(-4)
	result3 := Range(0)
	is.Equal([]int{0, 1, 2, 3}, result1)
	is.Equal([]int{0, -1, -2, -3}, result2)
	is.Empty(result3)
}

func TestRangeFrom(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := RangeFrom(1, 5)
	result2 := RangeFrom(-1, -5)
	result3 := RangeFrom(10, 0)
	result4 := RangeFrom(2.0, 3)
	result5 := RangeFrom(-2.0, -3)
	result6 := RangeFrom(2.5, 3)
	result7 := RangeFrom(-2.5, -3)
	is.Equal([]int{1, 2, 3, 4, 5}, result1)
	is.Equal([]int{-1, -2, -3, -4, -5}, result2)
	is.Empty(result3)
	is.Equal([]float64{2.0, 3.0, 4.0}, result4)
	is.Equal([]float64{-2.0, -3.0, -4.0}, result5)
	is.Equal([]float64{2.5, 3.5, 4.5}, result6)
	is.Equal([]float64{-2.5, -3.5, -4.5}, result7)
}

func TestRangeClose(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := RangeWithSteps(0, 20, 6)
	result2 := RangeWithSteps(0, 3, -5)
	result3 := RangeWithSteps(1, 1, 0)
	result4 := RangeWithSteps(3, 2, 1)
	result5 := RangeWithSteps(1.0, 4.0, 2.0)
	result6 := RangeWithSteps[float32](-1.0, -4.0, -1.0)
	is.Equal([]int{0, 6, 12, 18}, result1)
	is.Empty(result2)
	is.Empty(result3)
	is.Empty(result4)
	is.Equal([]float64{1.0, 3.0}, result5)
	is.Equal([]float32{-1.0, -2.0, -3.0}, result6)
}

func TestClamp(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Clamp(0, -10, 10)
	result2 := Clamp(-42, -10, 10)
	result3 := Clamp(42, -10, 10)

	is.Zero(result1)
	is.Equal(-10, result2)
	is.Equal(10, result3)
}

func TestSum(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Sum([]float32{2.3, 3.3, 4, 5.3})
	result2 := Sum([]int32{2, 3, 4, 5})
	result3 := Sum([]uint32{2, 3, 4, 5})
	result4 := Sum([]uint32{})
	result5 := Sum([]complex128{4_4, 2_2})

	is.InEpsilon(14.9, result1, 1e-7)
	is.Equal(int32(14), result2)
	is.Equal(uint32(14), result3)
	is.Equal(uint32(0), result4)
	is.Equal(complex128(6_6), result5)
}

func TestSumBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SumBy([]float32{2.3, 3.3, 4, 5.3}, func(n float32) float32 { return n })
	result2 := SumBy([]int32{2, 3, 4, 5}, func(n int32) int32 { return n })
	result3 := SumBy([]uint32{2, 3, 4, 5}, func(n uint32) uint32 { return n })
	result4 := SumBy([]uint32{}, func(n uint32) uint32 { return n })
	result5 := SumBy([]complex128{4_4, 2_2}, func(n complex128) complex128 { return n })

	is.InEpsilon(14.9, result1, 1e-7)
	is.Equal(int32(14), result2)
	is.Equal(uint32(14), result3)
	is.Equal(uint32(0), result4)
	is.Equal(complex128(6_6), result5)
}

func TestProduct(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Product([]float32{2.3, 3.3, 4, 5.3})
	result2 := Product([]int32{2, 3, 4, 5})
	result3 := Product([]int32{7, 8, 9, 0})
	result4 := Product([]int32{7, -1, 9, 2})
	result5 := Product([]uint32{2, 3, 4, 5})
	result6 := Product([]uint32{})
	result7 := Product([]complex128{4_4, 2_2})
	result8 := Product[uint32](nil)

	is.InEpsilon(160.908, result1, 1e-7)
	is.Equal(int32(120), result2)
	is.Equal(int32(0), result3)
	is.Equal(int32(-126), result4)
	is.Equal(uint32(120), result5)
	is.Equal(uint32(1), result6)
	is.Equal(complex128(96_8), result7)
	is.Equal(uint32(1), result8)
}

func TestProductBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ProductBy([]float32{2.3, 3.3, 4, 5.3}, func(n float32) float32 { return n })
	result2 := ProductBy([]int32{2, 3, 4, 5}, func(n int32) int32 { return n })
	result3 := ProductBy([]int32{7, 8, 9, 0}, func(n int32) int32 { return n })
	result4 := ProductBy([]int32{7, -1, 9, 2}, func(n int32) int32 { return n })
	result5 := ProductBy([]uint32{2, 3, 4, 5}, func(n uint32) uint32 { return n })
	result6 := ProductBy([]uint32{}, func(n uint32) uint32 { return n })
	result7 := ProductBy([]complex128{4_4, 2_2}, func(n complex128) complex128 { return n })
	result8 := ProductBy(nil, func(n uint32) uint32 { return n })

	is.InEpsilon(160.908, result1, 1e-7)
	is.Equal(int32(120), result2)
	is.Equal(int32(0), result3)
	is.Equal(int32(-126), result4)
	is.Equal(uint32(120), result5)
	is.Equal(uint32(1), result6)
	is.Equal(complex128(96_8), result7)
	is.Equal(uint32(1), result8)
}

func TestMean(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Mean([]float32{2.3, 3.3, 4, 5.3})
	result2 := Mean([]int32{2, 3, 4, 5})
	result3 := Mean([]uint32{2, 3, 4, 5})
	result4 := Mean([]uint32{})

	is.InEpsilon(3.725, result1, 1e-7)
	is.Equal(int32(3), result2)
	is.Equal(uint32(3), result3)
	is.Equal(uint32(0), result4)
}

func TestMeanBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MeanBy([]float32{2.3, 3.3, 4, 5.3}, func(n float32) float32 { return n })
	result2 := MeanBy([]int32{2, 3, 4, 5}, func(n int32) int32 { return n })
	result3 := MeanBy([]uint32{2, 3, 4, 5}, func(n uint32) uint32 { return n })
	result4 := MeanBy([]uint32{}, func(n uint32) uint32 { return n })

	is.InEpsilon(3.725, result1, 1e-7)
	is.Equal(int32(3), result2)
	is.Equal(uint32(3), result3)
	is.Equal(uint32(0), result4)
}

func TestMode(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Mode([]float32{2.3, 3.3, 3.3, 5.3})
	result2 := Mode([]int32{2, 2, 3, 4})
	result3 := Mode([]uint32{2, 2, 3, 3})
	result4 := Mode([]uint32{})
	result5 := Mode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	is.Equal([]float32{3.3}, result1)
	is.Equal([]int32{2}, result2)
	is.Equal([]uint32{2, 3}, result3)
	is.Empty(result4)
	is.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, result5)
}

func TestModeCapacityConsistency(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	arr := []int{1, 1, 2, 2, 3, 3, 3}

	result := Mode(arr)

	is.Equal([]int{3}, result, "Mode should return correct mode value")
	is.Equal(len(result), cap(result), "Mode slice capacity should match its length")
}
