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

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{name: "positive count", n: 4, expected: []int{0, 1, 2, 3}},
		{name: "negative count", n: -4, expected: []int{0, -1, -2, -3}},
		{name: "zero count", n: 0, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := slices.Collect(Range(tt.n))
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
		is := assert.New(t)

		tests := []struct {
			name       string
			start      int
			elementNum int
			expected   []int
		}{
			{name: "ascending", start: 1, elementNum: 5, expected: []int{1, 2, 3, 4, 5}},
			{name: "descending", start: -1, elementNum: -5, expected: []int{-1, -2, -3, -4, -5}},
			{name: "zero elements", start: 10, elementNum: 0, expected: nil},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				result := slices.Collect(RangeFrom(tt.start, tt.elementNum))
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
		is := assert.New(t)

		tests := []struct {
			name       string
			start      float64
			elementNum int
			expected   []float64
		}{
			{name: "positive integer step", start: 2.0, elementNum: 3, expected: []float64{2.0, 3.0, 4.0}},
			{name: "negative integer step", start: -2.0, elementNum: -3, expected: []float64{-2.0, -3.0, -4.0}},
			{name: "positive fractional step", start: 2.5, elementNum: 3, expected: []float64{2.5, 3.5, 4.5}},
			{name: "negative fractional step", start: -2.5, elementNum: -3, expected: []float64{-2.5, -3.5, -4.5}},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, slices.Collect(RangeFrom(tt.start, tt.elementNum)))
			})
		}
	})
}

func TestRangeWithSteps(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name             string
			start, end, step int
			expected         []int
		}{
			{name: "positive step", start: 0, end: 20, step: 6, expected: []int{0, 6, 12, 18}},
			{name: "negative step on ascending range", start: 0, end: 3, step: -5, expected: nil},
			{name: "zero step", start: 1, end: 1, step: 0, expected: nil},
			{name: "positive step on descending range", start: 3, end: 2, step: 1, expected: nil},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				result := slices.Collect(RangeWithSteps(tt.start, tt.end, tt.step))
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
		is := assert.New(t)

		tests := []struct {
			name             string
			start, end, step float64
			expected         []float64
		}{
			{name: "integer step", start: 1.0, end: 4.0, step: 2.0, expected: []float64{1.0, 3.0}},
			{name: "step larger than range", start: 0.0, end: 0.5, step: 1.0, expected: []float64{0.0}},
			{name: "fractional step", start: 0.0, end: 0.3, step: 0.1, expected: []float64{0.0, 0.1, 0.2}},
			{name: "fractional step, larger range", start: 0.0, end: 5.5, step: 2.5, expected: []float64{0.0, 2.5, 5.0}},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, slices.Collect(RangeWithSteps(tt.start, tt.end, tt.step)))
			})
		}
	})

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := RangeWithSteps[float32](-1.0, -4.0, -1.0)
		is.Equal([]float32{-1.0, -2.0, -3.0}, slices.Collect(result))
	})

	t.Run("custom float64-based type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type f64 float64

		result := RangeWithSteps[f64](0.0, 0.3, 0.1)
		is.Equal([]f64{0.0, 0.1, 0.2}, slices.Collect(result))
	})
}

func TestSum(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Sum(values[float32](2.3, 3.3, 4, 5.3))
		is.InEpsilon(14.9, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Sum(values[int32](2, 3, 4, 5))
		is.Equal(int32(14), result)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []uint32
			expected uint32
		}{
			{name: "non-empty", input: []uint32{2, 3, 4, 5}, expected: 14},
			{name: "empty", input: []uint32{}, expected: 0},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, Sum(values(tt.input...)))
			})
		}
	})

	t.Run("complex128", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Sum(values[complex128](4_4, 2_2))
		is.Equal(complex128(6_6), result)
	})
}

func TestSumBy(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := SumBy(values[float32](2.3, 3.3, 4, 5.3), func(n float32) float32 { return n })
		is.InEpsilon(14.9, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := SumBy(values[int32](2, 3, 4, 5), func(n int32) int32 { return n })
		is.Equal(int32(14), result)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []uint32
			expected uint32
		}{
			{name: "non-empty", input: []uint32{2, 3, 4, 5}, expected: 14},
			{name: "empty", input: []uint32{}, expected: 0},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, SumBy(values(tt.input...), func(n uint32) uint32 { return n }))
			})
		}
	})

	t.Run("complex128", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := SumBy(values[complex128](4_4, 2_2), func(n complex128) complex128 { return n })
		is.Equal(complex128(6_6), result)
	})
}

func TestProduct(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Product(values[float32](2.3, 3.3, 4, 5.3))
		is.InEpsilon(160.908, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []int32
			expected int32
		}{
			{name: "positive values", input: []int32{2, 3, 4, 5}, expected: 120},
			{name: "contains zero", input: []int32{7, 8, 9, 0}, expected: 0},
			{name: "contains negative", input: []int32{7, -1, 9, 2}, expected: -126},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, Product(values(tt.input...)))
			})
		}
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []uint32
			expected uint32
		}{
			{name: "non-empty", input: []uint32{2, 3, 4, 5}, expected: 120},
			{name: "empty", input: []uint32{}, expected: 1},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, Product(values(tt.input...)))
			})
		}
	})

	t.Run("complex128", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Product(values[complex128](4_4, 2_2))
		is.Equal(complex128(96_8), result)
	})
}

func TestProductBy(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := ProductBy(values[float32](2.3, 3.3, 4, 5.3), func(n float32) float32 { return n })
		is.InEpsilon(160.908, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []int32
			expected int32
		}{
			{name: "positive values", input: []int32{2, 3, 4, 5}, expected: 120},
			{name: "contains zero", input: []int32{7, 8, 9, 0}, expected: 0},
			{name: "contains negative", input: []int32{7, -1, 9, 2}, expected: -126},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, ProductBy(values(tt.input...), func(n int32) int32 { return n }))
			})
		}
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []uint32
			expected uint32
		}{
			{name: "non-empty", input: []uint32{2, 3, 4, 5}, expected: 120},
			{name: "empty", input: []uint32{}, expected: 1},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, ProductBy(values(tt.input...), func(n uint32) uint32 { return n }))
			})
		}
	})

	t.Run("complex128", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := ProductBy(values[complex128](4_4, 2_2), func(n complex128) complex128 { return n })
		is.Equal(complex128(96_8), result)
	})
}

func TestMean(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Mean(values[float32](2.3, 3.3, 4, 5.3))
		is.InEpsilon(3.725, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Mean(values[int32](2, 3, 4, 5))
		is.Equal(int32(3), result)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []uint32
			expected uint32
		}{
			{name: "non-empty", input: []uint32{2, 3, 4, 5}, expected: 3},
			{name: "empty", input: []uint32{}, expected: 0},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, Mean(values(tt.input...)))
			})
		}
	})
}

func TestMeanBy(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := MeanBy(values[float32](2.3, 3.3, 4, 5.3), func(n float32) float32 { return n })
		is.InEpsilon(3.725, result, 1e-7)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := MeanBy(values[int32](2, 3, 4, 5), func(n int32) int32 { return n })
		is.Equal(int32(3), result)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []uint32
			expected uint32
		}{
			{name: "non-empty", input: []uint32{2, 3, 4, 5}, expected: 3},
			{name: "empty", input: []uint32{}, expected: 0},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				is.Equal(tt.expected, MeanBy(values(tt.input...), func(n uint32) uint32 { return n }))
			})
		}
	})
}

func TestMode(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Mode(values[float32](2.3, 3.3, 3.3, 5.3))
		is.Equal([]float32{3.3}, result)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Mode(values[int32](2, 2, 3, 4))
		is.Equal([]int32{2}, result)
	})

	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []uint32
			expected []uint32
		}{
			{name: "two-way tie", input: []uint32{2, 2, 3, 3}, expected: []uint32{2, 3}},
			{name: "empty", input: []uint32{}, expected: nil},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				result := Mode(values(tt.input...))
				if tt.expected == nil {
					is.Empty(result)
				} else {
					is.Equal(tt.expected, result)
				}
			})
		}
	})

	t.Run("int, all unique", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Mode(values(1, 2, 3, 4, 5, 6, 7, 8, 9))
		is.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestMode_capacityConsistency(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	arr := []int{1, 1, 2, 2, 3, 3, 3}

	result := Mode(values(arr...))

	is.Equal([]int{3}, result, "Mode should return correct mode value")
	is.Equal(len(result), cap(result), "Mode slice capacity should match its length")
}
