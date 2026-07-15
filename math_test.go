package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		elemCount int
		expected  []int
	}{
		{name: "positive count", elemCount: 4, expected: []int{0, 1, 2, 3}},
		{name: "negative count", elemCount: -4, expected: []int{0, -1, -2, -3}},
		{name: "zero count", elemCount: 0, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			result := Range(tt.elemCount)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}
}

func TestRangeFrom(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name      string
			start     int
			elemCount int
			expected  []int
		}{
			{name: "ascending", start: 1, elemCount: 5, expected: []int{1, 2, 3, 4, 5}},
			{name: "descending", start: -1, elemCount: -5, expected: []int{-1, -2, -3, -4, -5}},
			{name: "zero count", start: 10, elemCount: 0, expected: nil},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				result := RangeFrom(tt.start, tt.elemCount)
				if tt.expected == nil {
					is.Empty(result)
				} else {
					is.Equal(tt.expected, result)
				}
			})
		}
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name      string
			start     float64
			elemCount int
			expected  []float64
		}{
			{name: "ascending integer step", start: 2.0, elemCount: 3, expected: []float64{2.0, 3.0, 4.0}},
			{name: "descending integer step", start: -2.0, elemCount: -3, expected: []float64{-2.0, -3.0, -4.0}},
			{name: "ascending fractional", start: 2.5, elemCount: 3, expected: []float64{2.5, 3.5, 4.5}},
			{name: "descending fractional", start: -2.5, elemCount: -3, expected: []float64{-2.5, -3.5, -4.5}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				is.Equal(tt.expected, RangeFrom(tt.start, tt.elemCount))
			})
		}
	})
}

func TestRangeWithSteps(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			start    int
			end      int
			step     int
			expected []int
		}{
			{name: "positive step", start: 0, end: 20, step: 6, expected: []int{0, 6, 12, 18}},
			{name: "negative step on ascending range", start: 0, end: 3, step: -5, expected: nil},
			{name: "zero step", start: 1, end: 1, step: 0, expected: nil},
			{name: "descending range with positive step", start: 3, end: 2, step: 1, expected: nil},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				result := RangeWithSteps(tt.start, tt.end, tt.step)
				if tt.expected == nil {
					is.Empty(result)
				} else {
					is.Equal(tt.expected, result)
				}
			})
		}
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			start    float64
			end      float64
			step     float64
			expected []float64
		}{
			{name: "step 2", start: 1.0, end: 4.0, step: 2.0, expected: []float64{1.0, 3.0}},
			{name: "single element result", start: 0.0, end: 0.5, step: 1.0, expected: []float64{0.0}},
			{name: "fractional step", start: 0.0, end: 0.3, step: 0.1, expected: []float64{0.0, 0.1, 0.2}},
			{name: "step larger than remaining range", start: 0.0, end: 5.5, step: 2.5, expected: []float64{0.0, 2.5, 5.0}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				is.Equal(tt.expected, RangeWithSteps(tt.start, tt.end, tt.step))
			})
		}
	})

	t.Run("float32", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, []float32{-1.0, -2.0, -3.0}, RangeWithSteps[float32](-1.0, -4.0, -1.0))
	})

	t.Run("named float64 type", func(t *testing.T) {
		t.Parallel()

		type f64 float64
		result := RangeWithSteps[f64](0.0, 0.3, 0.1)
		assert.Equal(t, []f64{0.0, 0.1, 0.2}, result)
	})
}

func TestClamp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    int
		min      int
		max      int
		expected int
	}{
		{name: "within range", value: 0, min: -10, max: 10, expected: 0},
		{name: "below min", value: -42, min: -10, max: 10, expected: -10},
		{name: "above max", value: 42, min: -10, max: 10, expected: 10},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, Clamp(tt.value, tt.min, tt.max))
		})
	}
}

func TestSum(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		assert.InEpsilon(t, 14.9, Sum([]float32{2.3, 3.3, 4, 5.3}), 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, int32(14), Sum([]int32{2, 3, 4, 5}))
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, uint32(14), Sum([]uint32{2, 3, 4, 5}))
	})

	t.Run("uint32 empty", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, uint32(0), Sum([]uint32{}))
	})

	t.Run("complex128", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, complex128(6_6), Sum([]complex128{4_4, 2_2}))
	})
}

func TestSumBy(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		result := SumBy([]float32{2.3, 3.3, 4, 5.3}, func(n float32) float32 { return n })
		assert.InEpsilon(t, 14.9, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		result := SumBy([]int32{2, 3, 4, 5}, func(n int32) int32 { return n })
		assert.Equal(t, int32(14), result)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		result := SumBy([]uint32{2, 3, 4, 5}, func(n uint32) uint32 { return n })
		assert.Equal(t, uint32(14), result)
	})

	t.Run("uint32 empty", func(t *testing.T) {
		t.Parallel()
		result := SumBy([]uint32{}, func(n uint32) uint32 { return n })
		assert.Equal(t, uint32(0), result)
	})

	t.Run("complex128", func(t *testing.T) {
		t.Parallel()
		result := SumBy([]complex128{4_4, 2_2}, func(n complex128) complex128 { return n })
		assert.Equal(t, complex128(6_6), result)
	})
}

func TestSumByErr(t *testing.T) {
	t.Parallel()

	t.Run("float32 no error", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		result, err := SumByErr([]float32{2.3, 3.3, 4, 5.3}, func(n float32) (float32, error) { return n, nil })
		is.NoError(err)
		is.InEpsilon(14.9, result, 1e-7)
	})

	t.Run("int32 no error", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		result, err := SumByErr([]int32{2, 3, 4, 5}, func(n int32) (int32, error) { return n, nil })
		is.NoError(err)
		is.Equal(int32(14), result)
	})

	t.Run("uint32 no error", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		result, err := SumByErr([]uint32{2, 3, 4, 5}, func(n uint32) (uint32, error) { return n, nil })
		is.NoError(err)
		is.Equal(uint32(14), result)
	})

	t.Run("complex128 no error", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		result, err := SumByErr([]complex128{4_4, 2_2}, func(n complex128) (complex128, error) { return n, nil })
		is.NoError(err)
		is.Equal(complex128(6_6), result)
	})

	t.Run("empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		result, err := SumByErr([]uint32{}, func(n uint32) (uint32, error) { return n, nil })
		is.NoError(err)
		is.Equal(uint32(0), result)
	})

	// Test error cases - table driven with callback count verification for early return
	testErr := assert.AnError

	errorTests := []struct {
		name          string
		input         []int32
		errorAt       int32
		alwaysError   bool
		expectedSum   int32
		expectedCalls int
	}{
		{
			// Sum up to 1+2 = 3; with 5 elements and error at 3rd, only 3 callbacks should be made.
			name:          "error at third element of five",
			input:         []int32{1, 2, 3, 4, 5},
			errorAt:       3,
			expectedSum:   3,
			expectedCalls: 3,
		},
		{
			name:          "error at first element",
			input:         []int32{1, 2, 3},
			alwaysError:   true,
			expectedSum:   0,
			expectedCalls: 1,
		},
		{
			// All 3 callbacks before error at last element.
			name:          "error at last element",
			input:         []int32{1, 2, 3},
			errorAt:       3,
			expectedSum:   3,
			expectedCalls: 3,
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			callbackCount := 0
			result, err := SumByErr(tt.input, func(n int32) (int32, error) {
				callbackCount++
				if tt.alwaysError || n == tt.errorAt {
					return 0, testErr
				}
				return n, nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedSum, result)
			is.Equal(tt.expectedCalls, callbackCount)
		})
	}
}

func TestProduct(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		assert.InEpsilon(t, 160.908, Product([]float32{2.3, 3.3, 4, 5.3}), 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			input    []int32
			expected int32
		}{
			{name: "basic", input: []int32{2, 3, 4, 5}, expected: 120},
			{name: "with zero", input: []int32{7, 8, 9, 0}, expected: 0},
			{name: "with negative", input: []int32{7, -1, 9, 2}, expected: -126},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				is.Equal(tt.expected, Product(tt.input))
			})
		}
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			input    []uint32
			expected uint32
		}{
			{name: "basic", input: []uint32{2, 3, 4, 5}, expected: 120},
			{name: "empty", input: []uint32{}, expected: 1},
			{name: "nil", input: nil, expected: 1},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				is.Equal(tt.expected, Product(tt.input))
			})
		}
	})

	t.Run("complex128", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, complex128(96_8), Product([]complex128{4_4, 2_2}))
	})
}

func TestProductBy(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		result := ProductBy([]float32{2.3, 3.3, 4, 5.3}, func(n float32) float32 { return n })
		assert.InEpsilon(t, 160.908, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			input    []int32
			expected int32
		}{
			{name: "basic", input: []int32{2, 3, 4, 5}, expected: 120},
			{name: "with zero", input: []int32{7, 8, 9, 0}, expected: 0},
			{name: "with negative", input: []int32{7, -1, 9, 2}, expected: -126},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				result := ProductBy(tt.input, func(n int32) int32 { return n })
				is.Equal(tt.expected, result)
			})
		}
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			input    []uint32
			expected uint32
		}{
			{name: "basic", input: []uint32{2, 3, 4, 5}, expected: 120},
			{name: "empty", input: []uint32{}, expected: 1},
			{name: "nil", input: nil, expected: 1},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				result := ProductBy(tt.input, func(n uint32) uint32 { return n })
				is.Equal(tt.expected, result)
			})
		}
	})

	t.Run("complex128", func(t *testing.T) {
		t.Parallel()
		result := ProductBy([]complex128{4_4, 2_2}, func(n complex128) complex128 { return n })
		assert.Equal(t, complex128(96_8), result)
	})
}

//nolint:errcheck,forcetypeassert
func TestProductByErr(t *testing.T) {
	t.Parallel()

	testErr := assert.AnError

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    any
		expected any
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
			is := assert.New(t)

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
			is := assert.New(t)

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

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		assert.InEpsilon(t, 3.725, Mean([]float32{2.3, 3.3, 4, 5.3}), 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, int32(3), Mean([]int32{2, 3, 4, 5}))
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, uint32(3), Mean([]uint32{2, 3, 4, 5}))
	})

	t.Run("uint32 empty", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, uint32(0), Mean([]uint32{}))
	})
}

func TestMeanBy(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		result := MeanBy([]float32{2.3, 3.3, 4, 5.3}, func(n float32) float32 { return n })
		assert.InEpsilon(t, 3.725, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		result := MeanBy([]int32{2, 3, 4, 5}, func(n int32) int32 { return n })
		assert.Equal(t, int32(3), result)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		result := MeanBy([]uint32{2, 3, 4, 5}, func(n uint32) uint32 { return n })
		assert.Equal(t, uint32(3), result)
	})

	t.Run("uint32 empty", func(t *testing.T) {
		t.Parallel()
		result := MeanBy([]uint32{}, func(n uint32) uint32 { return n })
		assert.Equal(t, uint32(0), result)
	})
}

//nolint:errcheck,forcetypeassert
func TestMeanByErr(t *testing.T) {
	t.Parallel()

	testErr := assert.AnError

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    any
		expected any
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
			is := assert.New(t)

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
			is := assert.New(t)

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

// TestMode_small exercises the small-scan path (all collections here are
// <= smallModeThreshold). See TestMode_large for the map-based path.
func TestMode_small(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, []float32{3.3}, Mode([]float32{2.3, 3.3, 3.3, 5.3}))
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, []int32{2}, Mode([]int32{2, 2, 3, 4}))
	})

	t.Run("uint32 multiple modes", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, []uint32{2, 3}, Mode([]uint32{2, 2, 3, 3}))
	})

	t.Run("uint32 empty", func(t *testing.T) {
		t.Parallel()
		assert.Empty(t, Mode([]uint32{}))
	})
}

// Mode dispatches on len(collection) <= smallModeThreshold (8): a collection
// of 9+ elements forces the modeLarge path, which the table above never
// exercises (its collections are all <= 4 elements).
func TestMode_large(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		collection []int
		expected   []int
	}{
		{name: "all unique elements", collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "two elements tied for mode", collection: []int{1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9}, expected: []int{1, 2}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, Mode(tt.collection))
		})
	}
}

func TestMode_capacityConsistency(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	arr := []int{1, 1, 2, 2, 3, 3, 3}

	result := Mode(arr)

	is.Equal([]int{3}, result, "Mode should return correct mode value")
	is.Equal(len(result), cap(result), "Mode slice capacity should match its length")
}
