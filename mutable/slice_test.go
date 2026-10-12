package mutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		input := []int{1, 2, 3, 4}
		result := Filter(input, func(x int) bool {
			return x%2 == 0
		})

		is.Equal([]int{2, 4, 3, 4}, input)
		is.Equal([]int{2, 4}, result)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		input := []string{"", "foo", "", "bar", ""}
		result := Filter(input, func(x string) bool {
			return len(x) > 0
		})

		is.Equal([]string{"foo", "bar", "", "bar", ""}, input)
		is.Equal([]string{"foo", "bar"}, result)
	})
}

func TestFilterI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterI([]int{1, 2, 3, 4}, func(x, i int) bool {
		is.Equal(i, x-1)
		return x%2 == 0
	})

	is.Equal([]int{2, 4}, r1)
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	input1 := []int{1, 2, 3, 4}
	r1 := Reject(input1, func(x int) bool {
		return x%2 == 0
	})

	is.Equal([]int{1, 3, 3, 4}, input1)
	is.Equal([]int{1, 3}, r1)

	input2 := []string{"", "foo", "", "bar", ""}
	r2 := Reject(input2, func(x string) bool {
		return len(x) > 0
	})

	is.Equal([]string{"", "", "", "bar", ""}, input2)
	is.Equal([]string{"", "", ""}, r2)
}

func TestRejectI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := RejectI([]int{1, 2, 3, 4}, func(x, i int) bool {
		is.Equal(i, x-1)
		return x%2 == 0
	})

	is.Equal([]int{1, 3}, r1)
}

func TestMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		multiply int
		expected []int
	}{
		{
			name:     "multiply by 2",
			input:    []int{1, 2, 3, 4},
			multiply: 2,
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "multiply by 4",
			input:    []int{1, 2, 3, 4},
			multiply: 4,
			expected: []int{4, 8, 12, 16},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			list := append([]int{}, tt.input...)
			Map(list, func(x int) int {
				return x * tt.multiply
			})
			is.Equal(tt.expected, list)
		})
	}
}

func TestMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list := []int{1, 2, 3, 4}
	MapI(list, func(x, index int) int {
		is.Equal(index, x-1)
		return x * 2
	})
	is.Equal([]int{2, 4, 6, 8}, list)

	list = []int{1, 2, 3, 4}
	MapI(list, func(x, index int) int {
		is.Equal(index, x-1)
		return x * 4
	})
	is.Equal([]int{4, 8, 12, 16}, list)
}

func TestShuffle(t *testing.T) {
	t.Parallel()

	t.Run("non-empty slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		Shuffle(list)
		is.NotEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, list)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		list := []int{}
		Shuffle(list)
		is.Empty(list)
	})
}

func TestReverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "even length",
			input:    []int{0, 1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			input:    []int{0, 1, 2, 3, 4, 5, 6},
			expected: []int{6, 5, 4, 3, 2, 1, 0},
		},
		{
			name:     "empty",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			list := append([]int{}, tt.input...)
			Reverse(list)
			is.Equal(tt.expected, list)
		})
	}

	t.Run("preserves named slice type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		Reverse(allStrings)
		is.IsType(myStrings{"", "foo", "bar"}, allStrings, "type preserved")
	})
}

func TestFill(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []string
		value    string
		expected []string
	}{
		{
			name:     "non-empty slice",
			input:    []string{"a", "0"},
			value:    "b",
			expected: []string{"b", "b"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			value:    "b",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			list := append([]string{}, tt.input...)
			Fill(list, tt.value)
			is.Equal(tt.expected, list)
		})
	}
}
