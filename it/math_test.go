//go:build go1.23

package it

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Range(4)
	result2 := Range(-4)
	result3 := Range(0)
	is.Equal([]int{0, 1, 2, 3}, slices.Collect(result1))
	is.Equal([]int{0, -1, -2, -3}, slices.Collect(result2))
	is.Empty(slices.Collect(result3))
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
	is.Equal([]int{1, 2, 3, 4, 5}, slices.Collect(result1))
	is.Equal([]int{-1, -2, -3, -4, -5}, slices.Collect(result2))
	is.Empty(slices.Collect(result3))
	is.Equal([]float64{2.0, 3.0, 4.0}, slices.Collect(result4))
	is.Equal([]float64{-2.0, -3.0, -4.0}, slices.Collect(result5))
	is.Equal([]float64{2.5, 3.5, 4.5}, slices.Collect(result6))
	is.Equal([]float64{-2.5, -3.5, -4.5}, slices.Collect(result7))
}

func TestRangeWithSteps(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := RangeWithSteps(0, 20, 6)
	result2 := RangeWithSteps(0, 3, -5)
	result3 := RangeWithSteps(1, 1, 0)
	result4 := RangeWithSteps(3, 2, 1)
	result5 := RangeWithSteps(1.0, 4.0, 2.0)
	result6 := RangeWithSteps[float32](-1.0, -4.0, -1.0)
	result7 := RangeWithSteps(0.0, 0.5, 1.0)
	result8 := RangeWithSteps(0.0, 0.3, 0.1)
	result9 := RangeWithSteps(0.0, 5.5, 2.5)

	type f64 float64
	result10 := RangeWithSteps[f64](0.0, 0.3, 0.1)

	is.Equal([]int{0, 6, 12, 18}, slices.Collect(result1))
	is.Empty(slices.Collect(result2))
	is.Empty(slices.Collect(result3))
	is.Empty(slices.Collect(result4))
	is.Equal([]float64{1.0, 3.0}, slices.Collect(result5))
	is.Equal([]float32{-1.0, -2.0, -3.0}, slices.Collect(result6))
	is.Equal([]float64{0.0}, slices.Collect(result7))
	is.Equal([]float64{0.0, 0.1, 0.2}, slices.Collect(result8))
	is.Equal([]float64{0.0, 2.5, 5.0}, slices.Collect(result9))
	is.Equal([]f64{0.0, 0.1, 0.2}, slices.Collect(result10))
}

func TestSum(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Sum(values[float32](2.3, 3.3, 4, 5.3))
	result2 := Sum(values[int32](2, 3, 4, 5))
	result3 := Sum(values[uint32](2, 3, 4, 5))
	result4 := Sum(values[uint32]())
	result5 := Sum(values[complex128](4_4, 2_2))

	is.InEpsilon(14.9, result1, 1e-7)
	is.Equal(int32(14), result2)
	is.Equal(uint32(14), result3)
	is.Equal(uint32(0), result4)
	is.Equal(complex128(6_6), result5)
}

func TestSumBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SumBy(values[float32](2.3, 3.3, 4, 5.3), func(n float32) float32 { return n })
	result2 := SumBy(values[int32](2, 3, 4, 5), func(n int32) int32 { return n })
	result3 := SumBy(values[uint32](2, 3, 4, 5), func(n uint32) uint32 { return n })
	result4 := SumBy(values[uint32](), func(n uint32) uint32 { return n })
	result5 := SumBy(values[complex128](4_4, 2_2), func(n complex128) complex128 { return n })

	is.InEpsilon(14.9, result1, 1e-7)
	is.Equal(int32(14), result2)
	is.Equal(uint32(14), result3)
	is.Equal(uint32(0), result4)
	is.Equal(complex128(6_6), result5)
}

func TestProduct(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Product(values[float32](2.3, 3.3, 4, 5.3))
	result2 := Product(values[int32](2, 3, 4, 5))
	result3 := Product(values[int32](7, 8, 9, 0))
	result4 := Product(values[int32](7, -1, 9, 2))
	result5 := Product(values[uint32](2, 3, 4, 5))
	result6 := Product(values[uint32]())
	result7 := Product(values[complex128](4_4, 2_2))

	is.InEpsilon(160.908, result1, 1e-7)
	is.Equal(int32(120), result2)
	is.Equal(int32(0), result3)
	is.Equal(int32(-126), result4)
	is.Equal(uint32(120), result5)
	is.Equal(uint32(1), result6)
	is.Equal(complex128(96_8), result7)
}

func TestProductBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ProductBy(values[float32](2.3, 3.3, 4, 5.3), func(n float32) float32 { return n })
	result2 := ProductBy(values[int32](2, 3, 4, 5), func(n int32) int32 { return n })
	result3 := ProductBy(values[int32](7, 8, 9, 0), func(n int32) int32 { return n })
	result4 := ProductBy(values[int32](7, -1, 9, 2), func(n int32) int32 { return n })
	result5 := ProductBy(values[uint32](2, 3, 4, 5), func(n uint32) uint32 { return n })
	result6 := ProductBy(values[uint32](), func(n uint32) uint32 { return n })
	result7 := ProductBy(values[complex128](4_4, 2_2), func(n complex128) complex128 { return n })

	is.InEpsilon(160.908, result1, 1e-7)
	is.Equal(int32(120), result2)
	is.Equal(int32(0), result3)
	is.Equal(int32(-126), result4)
	is.Equal(uint32(120), result5)
	is.Equal(uint32(1), result6)
	is.Equal(complex128(96_8), result7)
}

func TestMean(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Mean(values[float32](2.3, 3.3, 4, 5.3))
	result2 := Mean(values[int32](2, 3, 4, 5))
	result3 := Mean(values[uint32](2, 3, 4, 5))
	result4 := Mean(values[uint32]())

	is.InEpsilon(3.725, result1, 1e-7)
	is.Equal(int32(3), result2)
	is.Equal(uint32(3), result3)
	is.Equal(uint32(0), result4)
}

func TestMeanBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MeanBy(values[float32](2.3, 3.3, 4, 5.3), func(n float32) float32 { return n })
	result2 := MeanBy(values[int32](2, 3, 4, 5), func(n int32) int32 { return n })
	result3 := MeanBy(values[uint32](2, 3, 4, 5), func(n uint32) uint32 { return n })
	result4 := MeanBy(values[uint32](), func(n uint32) uint32 { return n })

	is.InEpsilon(3.725, result1, 1e-7)
	is.Equal(int32(3), result2)
	is.Equal(uint32(3), result3)
	is.Equal(uint32(0), result4)
}

func TestMode(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Mode(values[float32](2.3, 3.3, 3.3, 5.3))
	result2 := Mode(values[int32](2, 2, 3, 4))
	result3 := Mode(values[uint32](2, 2, 3, 3))
	result4 := Mode(values[uint32]())
	result5 := Mode(values(1, 2, 3, 4, 5, 6, 7, 8, 9))

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

	result := Mode(values(arr...))

	is.Equal([]int{3}, result, "Mode should return correct mode value")
	is.Equal(len(result), cap(result), "Mode slice capacity should match its length")
}
