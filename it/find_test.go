//go:build go1.23

package it

import (
	"iter"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/samber/lo/internal/xrand"
)

func TestIndexOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{name: "found", input: []int{0, 1, 2, 1, 2, 3}, target: 2, expected: 2},
		{name: "not found", input: []int{0, 1, 2, 1, 2, 3}, target: 6, expected: -1},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := IndexOf(values(tt.input...), tt.target)
			is.Equal(tt.expected, result)
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{name: "found", input: []int{0, 1, 2, 1, 2, 3}, target: 2, expected: 4},
		{name: "not found", input: []int{0, 1, 2, 1, 2, 3}, target: 6, expected: -1},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := LastIndexOf(values(tt.input...), tt.target)
			is.Equal(tt.expected, result)
		})
	}
}

func TestHasPrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		prefix   []int
		expected bool
	}{
		{name: "full prefix match", input: []int{1, 2, 3, 4}, prefix: []int{1, 2, 3, 4}, expected: true},
		{name: "partial prefix match", input: []int{1, 2, 3, 4}, prefix: []int{1, 2}, expected: true},
		{name: "no match", input: []int{1, 2, 3, 4}, prefix: []int{42}, expected: false},
		{name: "prefix longer than input", input: []int{1, 2}, prefix: []int{1, 2, 3, 4}, expected: false},
		{name: "empty prefix", input: []int{1, 2, 3, 4}, prefix: []int{}, expected: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := HasPrefix(values(tt.input...), tt.prefix...)
			is.Equal(tt.expected, result)
		})
	}
}

func TestHasSuffix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		suffix   []int
		expected bool
	}{
		{name: "full suffix match", input: []int{1, 2, 3, 4}, suffix: []int{1, 2, 3, 4}, expected: true},
		{name: "partial suffix match", input: []int{1, 2, 3, 4}, suffix: []int{3, 4}, expected: true},
		{name: "partial suffix match longer input", input: []int{1, 2, 3, 4, 5}, suffix: []int{3, 4, 5}, expected: true},
		{name: "no match", input: []int{1, 2, 3, 4}, suffix: []int{42}, expected: false},
		{name: "suffix longer than input", input: []int{1, 2}, suffix: []int{1, 2, 3, 4}, expected: false},
		{name: "empty suffix", input: []int{1, 2, 3, 4}, suffix: []int{}, expected: true},
		{name: "suffix longer than single-element input", input: []int{0}, suffix: []int{0, 0}, expected: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := HasSuffix(values(tt.input...), tt.suffix...)
			is.Equal(tt.expected, result)
		})
	}
}

func TestFind(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []string
		expectedResult string
		expectedOk     bool
	}{
		{name: "finds matching element", input: []string{"a", "b", "c", "d"}, expectedResult: "b", expectedOk: true},
		{name: "element not found", input: []string{"foobar"}, expectedResult: "", expectedOk: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			index := 0
			result, ok := Find(values(tt.input...), func(item string) bool {
				is.Equal(tt.input[index], item)
				index++
				return item == "b"
			})

			is.Equal(tt.expectedOk, ok)
			is.Equal(tt.expectedResult, result)
		})
	}
}

func TestFindIndexOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         []string
		expectedItem  string
		expectedIndex int
		expectedOk    bool
	}{
		{name: "finds matching element", input: []string{"a", "b", "c", "d", "b"}, expectedItem: "b", expectedIndex: 1, expectedOk: true},
		{name: "element not found", input: []string{"foobar"}, expectedItem: "", expectedIndex: -1, expectedOk: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			index := 0
			item, idx, ok := FindIndexOf(values(tt.input...), func(item string) bool {
				is.Equal(tt.input[index], item)
				index++
				return item == "b"
			})

			is.Equal(tt.expectedItem, item)
			is.Equal(tt.expectedOk, ok)
			is.Equal(tt.expectedIndex, idx)
		})
	}
}

func TestFindLastIndexOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         []string
		expectedItem  string
		expectedIndex int
		expectedOk    bool
	}{
		{name: "finds last matching element", input: []string{"a", "b", "c", "d", "b"}, expectedItem: "b", expectedIndex: 4, expectedOk: true},
		{name: "element not found", input: []string{"foobar"}, expectedItem: "", expectedIndex: -1, expectedOk: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			item, idx, ok := FindLastIndexOf(values(tt.input...), func(item string) bool {
				return item == "b"
			})

			is.Equal(tt.expectedItem, item)
			is.Equal(tt.expectedOk, ok)
			is.Equal(tt.expectedIndex, idx)
		})
	}
}

func TestFindOrElse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{name: "finds matching element", input: []string{"a", "b", "c", "d"}, expected: "b"},
		{name: "falls back when not found", input: []string{"foobar"}, expected: "x"},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			index := 0
			result := FindOrElse(values(tt.input...), "x", func(item string) bool {
				is.Equal(tt.input[index], item)
				index++
				return item == "b"
			})

			is.Equal(tt.expected, result)
		})
	}
}

func TestFindUniques(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "all unique", input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "some duplicates", input: []int{1, 2, 2, 3, 1, 2}, expected: []int{3}},
		{name: "all duplicates", input: []int{1, 2, 2, 1}, expected: nil},
		{name: "empty", input: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := FindUniques(values(tt.input...))
			is.Equal(tt.expected, slices.Collect(result))
		})
	}

	t.Run("preserves iterator type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := FindUniques(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestFindUniquesBy(t *testing.T) {
	t.Parallel()

	mod3 := func(i int) int { return i % 3 }

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "all unique keys", input: []int{0, 1, 2}, expected: []int{0, 1, 2}},
		{name: "one unique key", input: []int{0, 1, 2, 3, 4}, expected: []int{2}},
		{name: "no unique key", input: []int{0, 1, 2, 3, 4, 5}, expected: nil},
		{name: "empty", input: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := FindUniquesBy(values(tt.input...), mod3)
			is.Equal(tt.expected, slices.Collect(result))
		})
	}

	t.Run("preserves iterator type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := FindUniquesBy(allStrings, func(i string) string {
			return i
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestFindDuplicates(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "has duplicates", input: []int{1, 2, 2, 1, 2, 3}, expected: []int{2, 1}},
		{name: "no duplicates", input: []int{1, 2, 3}, expected: nil},
		{name: "empty", input: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := FindDuplicates(values(tt.input...))
			is.Equal(tt.expected, slices.Collect(result))
		})
	}

	t.Run("preserves iterator type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := FindDuplicates(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestFindDuplicatesBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		keyFn    func(int) int
		expected []int
	}{
		{name: "has duplicate keys", input: []int{3, 4, 5, 6, 7}, keyFn: func(i int) int { return i % 3 }, expected: []int{3, 4}},
		{name: "no duplicate keys", input: []int{0, 1, 2, 3, 4}, keyFn: func(i int) int { return i % 5 }, expected: nil},
		{name: "empty", input: []int{}, keyFn: func(i int) int { return i % 3 }, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := FindDuplicatesBy(values(tt.input...), tt.keyFn)
			is.Equal(tt.expected, slices.Collect(result))
		})
	}

	t.Run("preserves iterator type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := FindDuplicatesBy(allStrings, func(i string) string {
			return i
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestMin(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []int
			expected int
		}{
			{name: "ascending", input: []int{1, 2, 3}, expected: 1},
			{name: "descending", input: []int{3, 2, 1}, expected: 1},
			{name: "empty", input: []int{}, expected: 0},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				result := Min(values(tt.input...))
				is.Equal(tt.expected, result)
			})
		}
	})

	t.Run("durations", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Min(values(time.Second, time.Minute, time.Hour))
		is.Equal(time.Second, result)
	})
}

func TestMinIndex(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name          string
			input         []int
			expected      int
			expectedIndex int
		}{
			{name: "ascending", input: []int{1, 2, 3}, expected: 1, expectedIndex: 0},
			{name: "descending", input: []int{3, 2, 1}, expected: 1, expectedIndex: 2},
			{name: "empty", input: []int{}, expected: 0, expectedIndex: -1},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				result, index := MinIndex(values(tt.input...))
				is.Equal(tt.expected, result)
				is.Equal(tt.expectedIndex, index)
			})
		}
	})

	t.Run("durations", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result, index := MinIndex(values(time.Second, time.Minute, time.Hour))
		is.Equal(time.Second, result)
		is.Zero(index)
	})
}

func TestMinBy(t *testing.T) {
	t.Parallel()

	shorter := func(item, mIn string) bool {
		return len(item) < len(mIn)
	}

	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{name: "shortest first", input: []string{"s1", "string2", "s3"}, expected: "s1"},
		{name: "shortest last", input: []string{"string1", "string2", "s3"}, expected: "s3"},
		{name: "empty", input: []string{}, expected: ""},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := MinBy(values(tt.input...), shorter)
			is.Equal(tt.expected, result)
		})
	}
}

func TestMinIndexBy(t *testing.T) {
	t.Parallel()

	shorter := func(item, mIn string) bool {
		return len(item) < len(mIn)
	}

	tests := []struct {
		name          string
		input         []string
		expected      string
		expectedIndex int
	}{
		{name: "shortest first", input: []string{"s1", "string2", "s3"}, expected: "s1", expectedIndex: 0},
		{name: "shortest last", input: []string{"string1", "string2", "s3"}, expected: "s3", expectedIndex: 2},
		{name: "empty", input: []string{}, expected: "", expectedIndex: -1},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result, index := MinIndexBy(values(tt.input...), shorter)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedIndex, index)
		})
	}
}

func TestEarliest(t *testing.T) {
	t.Parallel()

	a := time.Now()
	b := a.Add(time.Hour)

	tests := []struct {
		name     string
		input    []time.Time
		expected time.Time
	}{
		{name: "non-empty", input: []time.Time{a, b}, expected: a},
		{name: "empty", input: []time.Time{}, expected: time.Time{}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := Earliest(values(tt.input...))
			is.Equal(tt.expected, result)
		})
	}
}

func TestEarliestBy(t *testing.T) {
	t.Parallel()

	type foo struct {
		bar time.Time
	}

	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	t3 := t1.Add(-time.Hour)
	extractBar := func(i foo) time.Time { return i.bar }

	tests := []struct {
		name     string
		input    []foo
		expected foo
	}{
		{name: "multiple items", input: []foo{{t1}, {t2}, {t3}}, expected: foo{t3}},
		{name: "single item", input: []foo{{t1}}, expected: foo{t1}},
		{name: "empty", input: []foo{}, expected: foo{}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := EarliestBy(values(tt.input...), extractBar)
			is.Equal(tt.expected, result)
		})
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []int
			expected int
		}{
			{name: "ascending", input: []int{1, 2, 3}, expected: 3},
			{name: "descending", input: []int{3, 2, 1}, expected: 3},
			{name: "empty", input: []int{}, expected: 0},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				result := Max(values(tt.input...))
				is.Equal(tt.expected, result)
			})
		}
	})

	t.Run("durations", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Max(values(time.Second, time.Minute, time.Hour))
		is.Equal(time.Hour, result)
	})
}

func TestMaxIndex(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name          string
			input         []int
			expected      int
			expectedIndex int
		}{
			{name: "ascending", input: []int{1, 2, 3}, expected: 3, expectedIndex: 2},
			{name: "descending", input: []int{3, 2, 1}, expected: 3, expectedIndex: 0},
			{name: "empty", input: []int{}, expected: 0, expectedIndex: -1},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)
				result, index := MaxIndex(values(tt.input...))
				is.Equal(tt.expected, result)
				is.Equal(tt.expectedIndex, index)
			})
		}
	})

	t.Run("durations", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result, index := MaxIndex(values(time.Second, time.Minute, time.Hour))
		is.Equal(time.Hour, result)
		is.Equal(2, index)
	})
}

func TestMaxBy(t *testing.T) {
	t.Parallel()

	longer := func(item, mAx string) bool {
		return len(item) > len(mAx)
	}

	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{name: "longest last", input: []string{"s1", "string2", "s3"}, expected: "string2"},
		{name: "longest first", input: []string{"string1", "string2", "s3"}, expected: "string1"},
		{name: "empty", input: []string{}, expected: ""},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := MaxBy(values(tt.input...), longer)
			is.Equal(tt.expected, result)
		})
	}
}

func TestMaxIndexBy(t *testing.T) {
	t.Parallel()

	longer := func(item, mAx string) bool {
		return len(item) > len(mAx)
	}

	tests := []struct {
		name          string
		input         []string
		expected      string
		expectedIndex int
	}{
		{name: "longest last", input: []string{"s1", "string2", "s3"}, expected: "string2", expectedIndex: 1},
		{name: "longest first", input: []string{"string1", "string2", "s3"}, expected: "string1", expectedIndex: 0},
		{name: "empty", input: []string{}, expected: "", expectedIndex: -1},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result, index := MaxIndexBy(values(tt.input...), longer)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedIndex, index)
		})
	}
}

func TestLatest(t *testing.T) {
	t.Parallel()

	a := time.Now()
	b := a.Add(time.Hour)

	tests := []struct {
		name     string
		input    []time.Time
		expected time.Time
	}{
		{name: "non-empty", input: []time.Time{a, b}, expected: b},
		{name: "empty", input: []time.Time{}, expected: time.Time{}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := Latest(values(tt.input...))
			is.Equal(tt.expected, result)
		})
	}
}

func TestLatestBy(t *testing.T) {
	t.Parallel()

	type foo struct {
		bar time.Time
	}

	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	t3 := t1.Add(-time.Hour)
	extractBar := func(i foo) time.Time { return i.bar }

	tests := []struct {
		name     string
		input    []foo
		expected foo
	}{
		{name: "multiple items", input: []foo{{t1}, {t2}, {t3}}, expected: foo{t2}},
		{name: "single item", input: []foo{{t1}}, expected: foo{t1}},
		{name: "empty", input: []foo{}, expected: foo{}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := LatestBy(values(tt.input...), extractBar)
			is.Equal(tt.expected, result)
		})
	}
}

func TestFirst(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		input      []int
		expected   int
		expectedOk bool
	}{
		{name: "non-empty", input: []int{1, 2, 3}, expected: 1, expectedOk: true},
		{name: "empty", input: []int{}, expected: 0, expectedOk: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result, ok := First(values(tt.input...))
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedOk, ok)
		})
	}
}

func TestFirstOrEmpty(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []int
			expected int
		}{
			{name: "non-empty", input: []int{1, 2, 3}, expected: 1},
			{name: "empty", input: []int{}, expected: 0},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				result := FirstOrEmpty(values(tt.input...))
				is.Equal(tt.expected, result)
			})
		}
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := FirstOrEmpty(values[string]())
		is.Empty(result)
	})
}

func TestFirstOr(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []int
			fallback int
			expected int
		}{
			{name: "non-empty", input: []int{1, 2, 3}, fallback: 63, expected: 1},
			{name: "empty", input: []int{}, fallback: 23, expected: 23},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				result := FirstOr(values(tt.input...), tt.fallback)
				is.Equal(tt.expected, result)
			})
		}
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := FirstOr(values[string](), "test")
		is.Equal("test", result)
	})
}

func TestLast(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		input      []int
		expected   int
		expectedOk bool
	}{
		{name: "non-empty", input: []int{1, 2, 3}, expected: 3, expectedOk: true},
		{name: "empty", input: []int{}, expected: 0, expectedOk: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result, ok := Last(values(tt.input...))
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedOk, ok)
		})
	}
}

func TestLastOrEmpty(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []int
			expected int
		}{
			{name: "non-empty", input: []int{1, 2, 3}, expected: 3},
			{name: "empty", input: []int{}, expected: 0},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				result := LastOrEmpty(values(tt.input...))
				is.Equal(tt.expected, result)
			})
		}
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := LastOrEmpty(values[string]())
		is.Empty(result)
	})
}

func TestLastOr(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		tests := []struct {
			name     string
			input    []int
			fallback int
			expected int
		}{
			{name: "non-empty", input: []int{1, 2, 3}, fallback: 63, expected: 3},
			{name: "empty", input: []int{}, fallback: 23, expected: 23},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				result := LastOr(values(tt.input...), tt.fallback)
				is.Equal(tt.expected, result)
			})
		}
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := LastOr(values[string](), "test")
		is.Equal("test", result)
	})
}

func TestNth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       []int
		n           int
		expected    int
		expectedErr string
	}{
		{name: "positive index in range", input: []int{0, 1, 2, 3}, n: 2, expected: 2},
		{name: "negative index", input: []int{0, 1, 2, 3}, n: -2, expected: 0, expectedErr: "nth: -2 out of bounds"},
		{name: "index out of range", input: []int{0, 1, 2, 3}, n: 42, expected: 0, expectedErr: "nth: 42 out of bounds"},
		{name: "empty collection", input: []int{}, n: 0, expected: 0, expectedErr: "nth: 0 out of bounds"},
		{name: "single element in range", input: []int{42}, n: 0, expected: 42},
		{name: "single element out of range", input: []int{42}, n: -1, expected: 0, expectedErr: "nth: -1 out of bounds"},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result, err := Nth(values(tt.input...), tt.n)
			is.Equal(tt.expected, result)
			if tt.expectedErr == "" {
				is.NoError(err)
			} else {
				is.EqualError(err, tt.expectedErr)
			}
		})
	}
}

func TestNthOr(t *testing.T) {
	t.Parallel()

	t.Run("Integers", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		const defaultValue = -1
		ints := values(10, 20, 30, 40, 50)

		is.Equal(30, NthOr(ints, 2, defaultValue))
		is.Equal(defaultValue, NthOr(ints, -1, defaultValue))
		is.Equal(defaultValue, NthOr(ints, 5, defaultValue))
	})

	t.Run("Strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		const defaultValue = "none"
		strs := values("apple", "banana", "cherry", "date")

		is.Equal("banana", NthOr(strs, 1, defaultValue))      // Index 1, expected "banana"
		is.Equal(defaultValue, NthOr(strs, -2, defaultValue)) // Negative index -2, expected "cherry"
		is.Equal(defaultValue, NthOr(strs, 10, defaultValue)) // Out of bounds, fallback "none"
	})

	t.Run("Structs", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type User struct {
			ID   int
			Name string
		}
		users := values(
			User{ID: 1, Name: "Alice"},
			User{ID: 2, Name: "Bob"},
			User{ID: 3, Name: "Charlie"},
		)
		defaultValue := User{ID: 0, Name: "Unknown"}

		is.Equal(User{ID: 1, Name: "Alice"}, NthOr(users, 0, defaultValue))
		is.Equal(defaultValue, NthOr(users, -1, defaultValue))
		is.Equal(defaultValue, NthOr(users, 10, defaultValue))
	})
}

func TestNthOrEmpty(t *testing.T) {
	t.Parallel()

	t.Run("Integers", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		ints := values(10, 20, 30, 40, 50)

		is.Equal(30, NthOrEmpty(ints, 2))
		is.Zero(NthOrEmpty(ints, -1))
		is.Zero(NthOrEmpty(ints, 10))
	})

	t.Run("Strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		strs := values("apple", "banana", "cherry", "date")

		is.Equal("banana", NthOrEmpty(strs, 1))
		is.Empty(NthOrEmpty(strs, -2))
		is.Empty(NthOrEmpty(strs, 10))
	})

	t.Run("Structs", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type User struct {
			ID   int
			Name string
		}
		users := values(
			User{ID: 1, Name: "Alice"},
			User{ID: 2, Name: "Bob"},
			User{ID: 3, Name: "Charlie"},
		)

		is.Equal(User{ID: 1, Name: "Alice"}, NthOrEmpty(users, 0))
		is.Zero(NthOrEmpty(users, -1))
		is.Zero(NthOrEmpty(users, 10))
	})
}

func TestSample(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       []string
		expectEmpty bool
	}{
		{name: "non-empty", input: []string{"a", "b", "c"}, expectEmpty: false},
		{name: "empty", input: []string{}, expectEmpty: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := Sample(values(tt.input...))
			if tt.expectEmpty {
				is.Empty(result)
			} else {
				is.True(Contains(values(tt.input...), result))
			}
		})
	}
}

func TestSampleBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       []string
		expectEmpty bool
	}{
		{name: "non-empty", input: []string{"a", "b", "c"}, expectEmpty: false},
		{name: "empty", input: []string{}, expectEmpty: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := SampleBy(values(tt.input...), xrand.IntN)
			if tt.expectEmpty {
				is.Empty(result)
			} else {
				is.True(Contains(values(tt.input...), result))
			}
		})
	}
}

func TestSamples(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       []string
		expectEmpty bool
	}{
		{name: "non-empty", input: []string{"a", "b", "c"}, expectEmpty: false},
		{name: "empty", input: []string{}, expectEmpty: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := slices.Collect(Samples(values(tt.input...), 3))
			if tt.expectEmpty {
				is.Empty(result)
			} else {
				is.ElementsMatch(result, tt.input)
			}
		})
	}

	t.Run("preserves iterator type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Samples(allStrings, 2)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestSamplesBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []string
		n     int
		keyFn func(int) int
		mode  string // "elementsMatch", "equal", "empty"
		exact []string
	}{
		{name: "random selection", input: []string{"a", "b", "c"}, n: 3, keyFn: xrand.IntN, mode: "elementsMatch"},
		{name: "empty input", input: []string{}, n: 3, keyFn: xrand.IntN, mode: "empty"},
		{name: "reverse order key", input: []string{"a", "b", "c"}, n: 3, keyFn: func(n int) int { return n - 1 }, mode: "equal", exact: []string{"c", "b", "a"}},
		{name: "constant zero key", input: []string{"a", "b", "c"}, n: 3, keyFn: func(int) int { return 0 }, mode: "equal", exact: []string{"a", "c", "b"}},
		{name: "zero count", input: []string{"a", "b", "c"}, n: 0, keyFn: func(int) int { return 1 }, mode: "empty"},
		{name: "negative count with nil keyFn", input: []string{"a", "b", "c"}, n: -1, keyFn: nil, mode: "empty"},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			result := slices.Collect(SamplesBy(values(tt.input...), tt.n, tt.keyFn))
			switch tt.mode {
			case "elementsMatch":
				is.ElementsMatch(result, tt.input)
			case "equal":
				is.Equal(tt.exact, result)
			case "empty":
				is.Empty(result)
			}
		})
	}

	t.Run("panics on out of range key", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		// index out of range [1] with length 1
		is.Panics(func() {
			SamplesBy(values("a", "b", "c"), 3, func(int) int { return 1 })
		})
	})

	t.Run("preserves iterator type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := SamplesBy(allStrings, 2, xrand.IntN)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}
