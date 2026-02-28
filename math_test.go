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

	is.Equal([]int{0, 6, 12, 18}, result1)
	is.Empty(result2)
	is.Empty(result3)
	is.Empty(result4)
	is.Equal([]float64{1.0, 3.0}, result5)
	is.Equal([]float32{-1.0, -2.0, -3.0}, result6)
	is.Equal([]float64{0.0}, result7)
	is.Equal([]float64{0.0, 0.1, 0.2}, result8)
	is.Equal([]float64{0.0, 2.5, 5.0}, result9)
	is.Equal([]f64{0.0, 0.1, 0.2}, result10)
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

func TestSumByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Test normal operation (no error)
	result1, err1 := SumByErr([]float32{2.3, 3.3, 4, 5.3}, func(n float32) (float32, error) { return n, nil })
	result2, err2 := SumByErr([]int32{2, 3, 4, 5}, func(n int32) (int32, error) { return n, nil })
	result3, err3 := SumByErr([]uint32{2, 3, 4, 5}, func(n uint32) (uint32, error) { return n, nil })
	result4, err4 := SumByErr([]complex128{4_4, 2_2}, func(n complex128) (complex128, error) { return n, nil })

	is.NoError(err1)
	is.InEpsilon(14.9, result1, 1e-7)
	is.NoError(err2)
	is.Equal(int32(14), result2)
	is.NoError(err3)
	is.Equal(uint32(14), result3)
	is.NoError(err4)
	is.Equal(complex128(6_6), result4)

	// Test empty collection
	result5, err5 := SumByErr([]uint32{}, func(n uint32) (uint32, error) { return n, nil })
	is.NoError(err5)
	is.Equal(uint32(0), result5)

	// Test error - iteratee returns error
	testErr := assert.AnError
	result6, err6 := SumByErr([]int32{1, 2, 3, 4, 5}, func(n int32) (int32, error) {
		if n == 3 {
			return 0, testErr
		}
		return n, nil
	})
	is.ErrorIs(err6, testErr)
	// Early return: sum up to 1+2 = 3
	is.Equal(int32(3), result6)

	// Test early return - callback count verification
	// With 5 elements and error at 3rd, only 3 callbacks should be made
	items := []int32{1, 2, 3, 4, 5}
	callbackCount := 0
	_, err7 := SumByErr(items, func(n int32) (int32, error) {
		callbackCount++
		if n == 3 {
			return 0, testErr
		}
		return n, nil
	})
	is.ErrorIs(err7, testErr)
	is.Equal(3, callbackCount) // Only 3 callbacks before error

	// Test error at first element
	result8, err8 := SumByErr([]int32{1, 2, 3}, func(n int32) (int32, error) {
		return 0, testErr
	})
	is.ErrorIs(err8, testErr)
	is.Equal(int32(0), result8)

	// Test error at last element
	callbackCount2 := 0
	result9, err9 := SumByErr([]int32{1, 2, 3}, func(n int32) (int32, error) {
		callbackCount2++
		if n == 3 {
			return 0, testErr
		}
		return n, nil
	})
	_ = result9 // unused due to error
	is.ErrorIs(err9, testErr)
	is.Equal(3, callbackCount2) // All 3 callbacks before error at last element
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

//nolint:errcheck,forcetypeassert
func TestProductByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := assert.AnError

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "float32 slice",
			input:    []float32{2.3, 3.3, 4, 5.3},
			expected: float32(160.908),
		},
		{
			name:     "int32 slice",
			input:    []int32{2, 3, 4, 5},
			expected: int32(120),
		},
		{
			name:     "int32 slice with zero",
			input:    []int32{7, 8, 9, 0},
			expected: int32(0),
		},
		{
			name:     "int32 slice with negative",
			input:    []int32{7, -1, 9, 2},
			expected: int32(-126),
		},
		{
			name:     "uint32 slice",
			input:    []uint32{2, 3, 4, 5},
			expected: uint32(120),
		},
		{
			name:     "empty uint32 slice",
			input:    []uint32{},
			expected: uint32(1),
		},
		{
			name:     "complex128 slice",
			input:    []complex128{4 + 4i, 2 + 2i},
			expected: complex128(0 + 16i),
		},
		{
			name:     "nil int32 slice",
			input:    ([]int32)(nil),
			expected: int32(1),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			switch input := tt.input.(type) {
			case []float32:
				result, err := ProductByErr(input, func(n float32) (float32, error) { return n, nil })
				is.NoError(err)
				is.InEpsilon(tt.expected.(float32), result, 1e-7)
			case []int32:
				result, err := ProductByErr(input, func(n int32) (int32, error) { return n, nil })
				is.NoError(err)
				is.Equal(tt.expected.(int32), result)
			case []uint32:
				result, err := ProductByErr(input, func(n uint32) (uint32, error) { return n, nil })
				is.NoError(err)
				is.Equal(tt.expected.(uint32), result)
			case []complex128:
				result, err := ProductByErr(input, func(n complex128) (complex128, error) { return n, nil })
				is.NoError(err)
				is.Equal(tt.expected.(complex128), result)
			}
		})
	}

	// Test error cases - table driven
	errorTests := []struct {
		name          string
		input         []int32
		errorAt       int32
		expectedProd  int32
		expectedCalls int
	}{
		{
			name:          "error at third element",
			input:         []int32{1, 2, 3, 4, 5},
			errorAt:       3,
			expectedProd:  2, // 1 * 2
			expectedCalls: 3,
		},
		{
			name:          "error at first element",
			input:         []int32{1, 2, 3},
			errorAt:       1,
			expectedProd:  1, // initial value
			expectedCalls: 1,
		},
		{
			name:          "error at last element",
			input:         []int32{1, 2, 3},
			errorAt:       3,
			expectedProd:  2, // 1 * 2
			expectedCalls: 3,
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			result, err := ProductByErr(tt.input, func(n int32) (int32, error) {
				callbackCount++
				if n == tt.errorAt {
					return 0, testErr
				}
				return n, nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedProd, result)
			is.Equal(tt.expectedCalls, callbackCount)
		})
	}
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

//nolint:errcheck,forcetypeassert
func TestMeanByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := assert.AnError

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "float32 slice",
			input:    []float32{2.3, 3.3, 4, 5.3},
			expected: float32(3.725),
		},
		{
			name:     "int32 slice",
			input:    []int32{2, 3, 4, 5},
			expected: int32(3),
		},
		{
			name:     "uint32 slice",
			input:    []uint32{2, 3, 4, 5},
			expected: uint32(3),
		},
		{
			name:     "empty uint32 slice",
			input:    []uint32{},
			expected: uint32(0),
		},
		{
			name:     "nil int32 slice",
			input:    ([]int32)(nil),
			expected: int32(0),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			switch input := tt.input.(type) {
			case []float32:
				result, err := MeanByErr(input, func(n float32) (float32, error) { return n, nil })
				is.NoError(err)
				is.InEpsilon(tt.expected.(float32), result, 1e-7)
			case []int32:
				result, err := MeanByErr(input, func(n int32) (int32, error) { return n, nil })
				is.NoError(err)
				is.Equal(tt.expected.(int32), result)
			case []uint32:
				result, err := MeanByErr(input, func(n uint32) (uint32, error) { return n, nil })
				is.NoError(err)
				is.Equal(tt.expected.(uint32), result)
			}
		})
	}

	// Test error cases - table driven
	errorTests := []struct {
		name          string
		input         []int32
		errorAt       int32
		expectedCalls int
	}{
		{
			name:          "error at third element",
			input:         []int32{1, 2, 3, 4, 5},
			errorAt:       3,
			expectedCalls: 3,
		},
		{
			name:          "error at first element",
			input:         []int32{1, 2, 3},
			errorAt:       1,
			expectedCalls: 1,
		},
		{
			name:          "error at last element",
			input:         []int32{1, 2, 3},
			errorAt:       3,
			expectedCalls: 3,
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			_, err := MeanByErr(tt.input, func(n int32) (int32, error) {
				callbackCount++
				if n == tt.errorAt {
					return 0, testErr
				}
				return n, nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedCalls, callbackCount)
		})
	}
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
