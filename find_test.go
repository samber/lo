package lo

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(2, result1)
	is.Equal(-1, result2)
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(4, result1)
	is.Equal(-1, result2)
}

func TestHasPrefix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(HasPrefix([]int{1, 2, 3, 4}, []int{1, 2}))
	is.False(HasPrefix([]int{1, 2, 3, 4}, []int{42}))
	is.True(HasPrefix([]int{1, 2, 3, 4}, nil))
}

func TestHasSuffix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(HasSuffix([]int{1, 2, 3, 4}, []int{3, 4}))
	is.False(HasSuffix([]int{1, 2, 3, 4}, []int{42}))
	is.True(HasSuffix([]int{1, 2, 3, 4}, nil))
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1, ok1 := Find([]string{"a", "b", "c", "d"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})

	result2, ok2 := Find([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.True(ok1)
	is.Equal("b", result1)
	is.False(ok2)
	is.Empty(result2)
}

func TestFindErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := assert.AnError

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "finds matching element",
			input:    []string{"a", "b", "c", "d"},
			expected: "b",
		},
		{
			name:     "element not found",
			input:    []string{"foobar"},
			expected: "",
		},
		{
			name:     "empty collection",
			input:    []string{},
			expected: "",
		},
		{
			name:     "single element found",
			input:    []string{"b"},
			expected: "b",
		},
		{
			name:     "single element not found",
			input:    []string{"a"},
			expected: "",
		},
		{
			name:     "finds first match",
			input:    []string{"a", "b", "c", "b"},
			expected: "b", // first "b"
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := FindErr(tt.input, func(item string) (bool, error) {
				return item == "b", nil
			})
			is.NoError(err)
			is.Equal(tt.expected, result)
		})
	}

	// Test error cases - table driven with callback count verification
	errorTests := []struct {
		name          string
		input         []string
		errorAt       string
		expectedCalls int
	}{
		{
			name:          "error at first element",
			input:         []string{"b", "c", "d"},
			errorAt:       "b",
			expectedCalls: 1,
		},
		{
			name:          "error at second element",
			input:         []string{"a", "b", "c"},
			errorAt:       "b",
			expectedCalls: 2,
		},
		{
			name:          "error at third element",
			input:         []string{"a", "c", "b"},
			errorAt:       "b",
			expectedCalls: 3,
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			result, err := FindErr(tt.input, func(item string) (bool, error) {
				callbackCount++
				if item == tt.errorAt {
					return false, testErr
				}
				return item == "b", nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedCalls, callbackCount, "callback count mismatch - iteration didn't stop early")
			is.Empty(result, "zero value should be returned on error")
		})
	}
}

func TestFindIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d", "b"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal("b", item1)
	is.True(ok1)
	is.Equal(1, index1)
	is.Empty(item2)
	is.False(ok2)
	is.Equal(-1, index2)
}

func TestFindLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindLastIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"b", "d", "c", "b", "a"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindLastIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal("b", item1)
	is.True(ok1)
	is.Equal(4, index1)
	is.Empty(item2)
	is.False(ok2)
	is.Equal(-1, index2)
}

func TestFindOrElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1 := FindOrElse([]string{"a", "b", "c", "d"}, "x", func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})
	result2 := FindOrElse([]string{"foobar"}, "x", func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal("b", result1)
	is.Equal("x", result2)
}

func TestFindKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 2)
	is.Equal("bar", result1)
	is.True(ok1)

	result2, ok2 := FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 42)
	is.Empty(result2)
	is.False(ok2)

	type test struct {
		foobar string
	}

	result3, ok3 := FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"foo"})
	is.Equal("foo", result3)
	is.True(ok3)

	result4, ok4 := FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"hello world"})
	is.Empty(result4)
	is.False(ok4)
}

func TestFindKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return k == "foo"
	})
	is.Equal("foo", result1)
	is.True(ok1)

	result2, ok2 := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Empty(result2)
	is.False(ok2)
}

// TestFindUniquesSmallScan exercises the small-scan path (all collections
// here are <= findSmallThreshold). See TestFindUniquesLarge for the
// map-based path.
func TestFindUniquesSmallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniques([]int{1, 2, 3})
	is.Equal([]int{1, 2, 3}, result1)

	result2 := FindUniques([]int{1, 2, 2, 3, 1, 2})
	is.Equal([]int{3}, result2)

	result3 := FindUniques([]int{1, 2, 2, 1})
	is.Empty(result3)

	result4 := FindUniques([]int{})
	is.Empty(result4)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := FindUniques(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

// FindUniques dispatches on len(collection) <= findSmallThreshold (8): a
// collection of 12 elements forces the findUniquesLarge path, which the
// table above never exercises (its collections are all <= 6 elements).
func TestFindUniquesLarge(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10}
	is.Greater(len(collection), findSmallThreshold, "sanity check: collection must exceed findSmallThreshold")
	is.Equal([]int{30, 50, 60, 70, 80, 90}, FindUniques(collection))

	allDup := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	is.Greater(len(allDup), findSmallThreshold, "sanity check: allDup must exceed findSmallThreshold")
	is.Empty(FindUniques(allDup))

	type myInts []int
	allUnique := myInts{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	is.Greater(len(allUnique), findSmallThreshold, "sanity check: allUnique must exceed findSmallThreshold")
	nonempty := FindUniques(allUnique)
	is.Equal(myInts{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nonempty)
	is.IsType(nonempty, allUnique, "type preserved")
}

// TestFindUniquesBySmallScan exercises the small-scan path (all collections
// here are <= findSmallThreshold). See TestFindUniquesByLarge for the
// map-based path.
func TestFindUniquesBySmallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniquesBy([]int{0, 1, 2}, func(i int) int {
		return i % 3
	})
	is.Equal([]int{0, 1, 2}, result1)

	result2 := FindUniquesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 3
	})
	is.Equal([]int{2}, result2)

	result3 := FindUniquesBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	is.Empty(result3)

	result4 := FindUniquesBy([]int{}, func(i int) int {
		return i % 3
	})
	is.Empty(result4)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := FindUniquesBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

// FindUniquesBy dispatches on len(collection) <= findSmallThreshold (8): a
// collection of 12 elements forces the findUniquesByLarge path, which the
// table above never exercises (its collections are all <= 6 elements).
func TestFindUniquesByLarge(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10}
	is.Greater(len(collection), findSmallThreshold, "sanity check: collection must exceed findSmallThreshold")
	byTen := func(v int) int { return v / 10 }
	is.Equal([]int{30, 50, 60, 70, 80, 90}, FindUniquesBy(collection, byTen))

	allDup := []int{10, 11, 20, 21, 30, 31, 40, 41, 50, 51}
	is.Greater(len(allDup), findSmallThreshold, "sanity check: allDup must exceed findSmallThreshold")
	is.Empty(FindUniquesBy(allDup, byTen))

	type myStrings []string
	allStrings := myStrings{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiiiii"}
	is.Greater(len(allStrings), findSmallThreshold, "sanity check: allStrings must exceed findSmallThreshold")
	nonempty := FindUniquesBy(allStrings, func(s string) int { return len(s) })
	is.IsType(nonempty, allStrings, "type preserved")
}

// TestFindDuplicatesSmallScan exercises the small-scan path (all
// collections here are <= findSmallThreshold). See
// TestFindDuplicatesLarge for the map-based path.
func TestFindDuplicatesSmallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicates([]int{1, 2, 2, 1, 2, 3})
	is.Equal([]int{1, 2}, result1)

	result2 := FindDuplicates([]int{1, 2, 3})
	is.Empty(result2)

	result3 := FindDuplicates([]int{})
	is.Empty(result3)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := FindDuplicates(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

// FindDuplicates dispatches on len(collection) <= findSmallThreshold (8): a
// collection of 12 elements forces the findDuplicatesLarge path, which the
// table above never exercises (its collections are all <= 6 elements).
func TestFindDuplicatesLarge(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10}
	is.Greater(len(collection), findSmallThreshold, "sanity check: collection must exceed findSmallThreshold")
	is.Equal([]int{10, 20, 40}, FindDuplicates(collection))

	noDup := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	is.Greater(len(noDup), findSmallThreshold, "sanity check: noDup must exceed findSmallThreshold")
	is.Empty(FindDuplicates(noDup))

	type myStrings []string
	allStrings := myStrings{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	is.Greater(len(allStrings), findSmallThreshold, "sanity check: allStrings must exceed findSmallThreshold")
	nonempty := FindDuplicates(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

// TestFindDuplicatesBySmallScan exercises the small-scan path (all
// collections here are <= findSmallThreshold). See
// TestFindDuplicatesByLarge for the map-based path.
func TestFindDuplicatesBySmallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicatesBy([]int{3, 4, 5, 6, 7}, func(i int) int {
		return i % 3
	})
	is.Equal([]int{3, 4}, result1)

	result2 := FindDuplicatesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 5
	})
	is.Empty(result2)

	result3 := FindDuplicatesBy([]int{}, func(i int) int {
		return i % 3
	})
	is.Empty(result3)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := FindDuplicatesBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

// FindDuplicatesBy dispatches on len(collection) <= findSmallThreshold (8): a
// collection of 12 elements forces the findDuplicatesByLarge path, which the
// table above never exercises (its collections are all <= 5 elements).
func TestFindDuplicatesByLarge(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10}
	is.Greater(len(collection), findSmallThreshold, "sanity check: collection must exceed findSmallThreshold")
	byTen := func(v int) int { return v / 10 }
	is.Equal([]int{10, 20, 40}, FindDuplicatesBy(collection, byTen))

	noDup := []int{10, 21, 32, 43, 54, 65, 76, 87, 98, 109}
	is.Greater(len(noDup), findSmallThreshold, "sanity check: noDup must exceed findSmallThreshold")
	is.Empty(FindDuplicatesBy(noDup, byTen))

	type myStrings []string
	allStrings := myStrings{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiiiii"}
	is.Greater(len(allStrings), findSmallThreshold, "sanity check: allStrings must exceed findSmallThreshold")
	nonempty := FindDuplicatesBy(allStrings, func(s string) int { return len(s) })
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFindDuplicatesByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Table-driven tests for normal operation
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "finds duplicates by key",
			input:    []int{3, 4, 5, 6, 7},
			expected: []int{3, 4},
		},
		{
			name:     "no duplicates",
			input:    []int{0, 1, 2, 3, 4},
			expected: []int{0, 1},
		},
		{
			name:     "empty collection",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "all duplicates",
			input:    []int{0, 3, 6, 9},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := FindDuplicatesByErr(tt.input, func(i int) (int, error) {
				return i % 3, nil
			})
			is.NoError(err)
			is.Equal(tt.expected, result)
		})
	}

	// Table-driven tests with callback count verification for early return
	testErr := errors.New("test error")

	errorTests := []struct {
		name          string
		input         []int
		errorAt       int
		expectedCalls int
	}{
		{
			name:          "error in first pass at element 0",
			input:         []int{3, 4, 5},
			errorAt:       0,
			expectedCalls: 1,
		},
		{
			name:          "error in first pass at element 2",
			input:         []int{3, 4, 5},
			errorAt:       2,
			expectedCalls: 3,
		},
		{
			name:          "error in second pass at first duplicate",
			input:         []int{3, 4, 5, 6},
			errorAt:       3,
			expectedCalls: 4, // First pass completes (4 items), error at first item of second pass
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			result, err := FindDuplicatesByErr(tt.input, func(i int) (int, error) {
				callbackCount++
				if i == tt.input[tt.errorAt] {
					return 0, testErr
				}
				return i % 3, nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedCalls, callbackCount, "callback count mismatch - iteration didn't stop early")
			is.Nil(result, "nil should be returned on error")
		})
	}

	// Test type preservation
	type myStrings []string
	allStrings := myStrings{"a", "b", "a", "c", "b"}
	result, err := FindDuplicatesByErr(allStrings, func(s string) (string, error) {
		return s, nil
	})
	is.NoError(err)
	is.IsType(result, allStrings, "type preserved")
}

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Min([]int{1, 2, 3})
	result2 := Min([]int{3, 2, 1})
	result3 := Min([]time.Duration{time.Second, time.Minute, time.Hour})
	result4 := Min([]int{})

	is.Equal(1, result1)
	is.Equal(1, result2)
	is.Equal(time.Second, result3)
	is.Zero(result4)
}

func TestMinIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MinIndex([]int{1, 2, 3})
	result2, index2 := MinIndex([]int{3, 2, 1})
	result3, index3 := MinIndex([]time.Duration{time.Second, time.Minute, time.Hour})
	result4, index4 := MinIndex([]int{})

	is.Equal(1, result1)
	is.Zero(index1)

	is.Equal(1, result2)
	is.Equal(2, index2)

	is.Equal(time.Second, result3)
	is.Zero(index3)

	is.Zero(result4)
	is.Equal(-1, index4)
}

func TestMinBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MinBy([]string{"s1", "string2", "s3"}, func(item, mIn string) bool {
		return len(item) < len(mIn)
	})
	result2 := MinBy([]string{"string1", "string2", "s3"}, func(item, mIn string) bool {
		return len(item) < len(mIn)
	})
	result3 := MinBy([]string{}, func(item, mIn string) bool {
		return len(item) < len(mIn)
	})

	is.Equal("s1", result1)
	is.Equal("s3", result2)
	is.Empty(result3)
}

func TestMinByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := assert.AnError

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "finds min by length - first match",
			input:    []string{"s1", "string2", "s3"},
			expected: "s1",
		},
		{
			name:     "finds min by length - third match",
			input:    []string{"string1", "string2", "s3"},
			expected: "s3",
		},
		{
			name:     "empty collection",
			input:    []string{},
			expected: "",
		},
		{
			name:     "single element",
			input:    []string{"single"},
			expected: "single",
		},
		{
			name:     "all equal length",
			input:    []string{"a", "b", "c"},
			expected: "a", // first minimal value
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := MinByErr(tt.input, func(item, mIn string) (bool, error) {
				return len(item) < len(mIn), nil
			})
			is.NoError(err)
			is.Equal(tt.expected, result)
		})
	}

	// Test error cases - table driven
	errorTests := []struct {
		name          string
		input         []string
		errorAt       string
		expectedCalls int
	}{
		{
			name:          "error at second comparison",
			input:         []string{"a", "bb", "ccc"},
			errorAt:       "bb",
			expectedCalls: 1, // Only first comparison (initial element vs second)
		},
		{
			name:          "error at third comparison",
			input:         []string{"a", "bb", "ccc"},
			errorAt:       "ccc",
			expectedCalls: 2, // First two comparisons
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			result, err := MinByErr(tt.input, func(item, mIn string) (bool, error) {
				callbackCount++
				if item == tt.errorAt {
					return false, testErr
				}
				return len(item) < len(mIn), nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedCalls, callbackCount)
			is.Empty(result) // Zero value on error
		})
	}
}

func TestMinIndexBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MinIndexBy([]string{"s1", "string2", "s3"}, func(item, mIn string) bool {
		return len(item) < len(mIn)
	})
	result2, index2 := MinIndexBy([]string{"string1", "string2", "s3"}, func(item, mIn string) bool {
		return len(item) < len(mIn)
	})
	result3, index3 := MinIndexBy([]string{}, func(item, mIn string) bool {
		return len(item) < len(mIn)
	})

	is.Equal("s1", result1)
	is.Zero(index1)

	is.Equal("s3", result2)
	is.Equal(2, index2)

	is.Empty(result3)
	is.Equal(-1, index3)
}

func TestMinIndexByErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		input                 []string
		less                  func(a, b string) (bool, error)
		wantValue             string
		wantIndex             int
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:                  "empty slice",
			input:                 []string{},
			less:                  func(a, b string) (bool, error) { return len(a) < len(b), nil },
			wantValue:             "",
			wantIndex:             -1,
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:                  "success case",
			input:                 []string{"s1", "string2", "s3"},
			less:                  func(a, b string) (bool, error) { return len(a) < len(b), nil },
			wantValue:             "s1",
			wantIndex:             0,
			wantErr:               false,
			expectedCallbackCount: 2,
		},
		{
			name:  "error on first comparison",
			input: []string{"s1", "string2", "s3"},
			less: func(a, b string) (bool, error) {
				return false, errors.New("comparison error")
			},
			wantValue:             "",
			wantIndex:             -1,
			wantErr:               true,
			errMsg:                "comparison error",
			expectedCallbackCount: 1,
		},
		{
			name:  "error on second comparison",
			input: []string{"a", "bb", "ccc", "error", "e"},
			less: func(a, b string) (bool, error) {
				if a == "error" || b == "error" {
					return false, errors.New("error value encountered")
				}
				return len(a) < len(b), nil
			},
			wantValue:             "",
			wantIndex:             -1,
			wantErr:               true,
			errMsg:                "error value encountered",
			expectedCallbackCount: 3,
		},
		{
			name:                  "single element",
			input:                 []string{"single"},
			less:                  func(a, b string) (bool, error) { return len(a) < len(b), nil },
			wantValue:             "single",
			wantIndex:             0,
			wantErr:               false,
			expectedCallbackCount: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			wrappedLess := func(a, b string) (bool, error) {
				callbackCount++
				return tt.less(a, b)
			}

			value, index, err := MinIndexByErr(tt.input, wrappedLess)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Equal(t, tt.errMsg, err.Error())
				}
				assert.Empty(t, value)
				assert.Equal(t, -1, index)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantValue, value)
				assert.Equal(t, tt.wantIndex, index)
			}

			assert.Equal(t, tt.expectedCallbackCount, callbackCount, "callback count mismatch")
		})
	}
}

func TestEarliest(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	a := time.Now()
	b := a.Add(time.Hour)
	result1 := Earliest(a, b)
	result2 := Earliest()

	is.Equal(a, result1)
	is.Zero(result2)
}

func TestEarliestBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type foo struct {
		bar time.Time
	}

	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	t3 := t1.Add(-time.Hour)
	result1 := EarliestBy([]foo{{t1}, {t2}, {t3}}, func(i foo) time.Time {
		return i.bar
	})
	result2 := EarliestBy([]foo{{t1}}, func(i foo) time.Time {
		return i.bar
	})
	result3 := EarliestBy([]foo{}, func(i foo) time.Time {
		return i.bar
	})

	is.Equal(foo{t3}, result1)
	is.Equal(foo{t1}, result2)
	is.Zero(result3)
}

func TestEarliestByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := assert.AnError

	type foo struct {
		bar time.Time
	}

	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	t3 := t1.Add(-time.Hour)

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    []foo
		expected foo
	}{
		{
			name:     "finds earliest time",
			input:    []foo{{t1}, {t2}, {t3}},
			expected: foo{t3},
		},
		{
			name:     "single element",
			input:    []foo{{t1}},
			expected: foo{t1},
		},
		{
			name:     "empty collection",
			input:    []foo{},
			expected: foo{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := EarliestByErr(tt.input, func(i foo) (time.Time, error) {
				return i.bar, nil
			})
			is.NoError(err)
			is.Equal(tt.expected, result)
		})
	}

	// Test error cases - table driven
	errorTests := []struct {
		name          string
		input         []foo
		errorAt       int
		expectedCalls int
	}{
		{
			name:          "error at first element",
			input:         []foo{{t1}, {t2}, {t3}},
			errorAt:       0,
			expectedCalls: 1, // Only first callback
		},
		{
			name:          "error at second element",
			input:         []foo{{t1}, {t2}, {t3}},
			errorAt:       1,
			expectedCalls: 2, // First two callbacks
		},
		{
			name:          "error at third element",
			input:         []foo{{t1}, {t2}, {t3}},
			errorAt:       2,
			expectedCalls: 3, // All three callbacks
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			_, err := EarliestByErr(tt.input, func(i foo) (time.Time, error) {
				callbackCount++
				if len(tt.input) > 0 && i == tt.input[tt.errorAt] {
					return time.Time{}, testErr
				}
				return i.bar, nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedCalls, callbackCount)
		})
	}
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Max([]int{1, 2, 3})
	result2 := Max([]int{3, 2, 1})
	result3 := Max([]time.Duration{time.Second, time.Minute, time.Hour})
	result4 := Max([]int{})

	is.Equal(3, result1)
	is.Equal(3, result2)
	is.Equal(time.Hour, result3)
	is.Zero(result4)
}

func TestMaxIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MaxIndex([]int{1, 2, 3})
	result2, index2 := MaxIndex([]int{3, 2, 1})
	result3, index3 := MaxIndex([]time.Duration{time.Second, time.Minute, time.Hour})
	result4, index4 := MaxIndex([]int{})

	is.Equal(3, result1)
	is.Equal(2, index1)

	is.Equal(3, result2)
	is.Zero(index2)

	is.Equal(time.Hour, result3)
	is.Equal(2, index3)

	is.Zero(result4)
	is.Equal(-1, index4)
}

func TestMaxBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MaxBy([]string{"s1", "string2", "s3"}, func(item, mAx string) bool {
		return len(item) > len(mAx)
	})
	result2 := MaxBy([]string{"string1", "string2", "s3"}, func(item, mAx string) bool {
		return len(item) > len(mAx)
	})
	result3 := MaxBy([]string{}, func(item, mAx string) bool {
		return len(item) > len(mAx)
	})

	is.Equal("string2", result1)
	is.Equal("string1", result2)
	is.Empty(result3)
}

func TestMaxByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := assert.AnError

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "finds max by length - second match",
			input:    []string{"s1", "string2", "s3"},
			expected: "string2",
		},
		{
			name:     "finds max by length - first match",
			input:    []string{"string1", "string2", "s3"},
			expected: "string1",
		},
		{
			name:     "empty collection",
			input:    []string{},
			expected: "",
		},
		{
			name:     "single element",
			input:    []string{"single"},
			expected: "single",
		},
		{
			name:     "all equal length",
			input:    []string{"a", "b", "c"},
			expected: "a", // first maximal value
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := MaxByErr(tt.input, func(item, mAx string) (bool, error) {
				return len(item) > len(mAx), nil
			})
			is.NoError(err)
			is.Equal(tt.expected, result)
		})
	}

	// Test error cases - table driven
	errorTests := []struct {
		name          string
		input         []string
		errorAt       string
		expectedCalls int
	}{
		{
			name:          "error at second comparison",
			input:         []string{"a", "bb", "ccc"},
			errorAt:       "bb",
			expectedCalls: 1, // Only first comparison (initial element vs second)
		},
		{
			name:          "error at third comparison",
			input:         []string{"a", "bb", "ccc"},
			errorAt:       "ccc",
			expectedCalls: 2, // First two comparisons
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			result, err := MaxByErr(tt.input, func(item, mAx string) (bool, error) {
				callbackCount++
				if item == tt.errorAt {
					return false, testErr
				}
				return len(item) > len(mAx), nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedCalls, callbackCount)
			// Result should be the current max at the time of error
			if tt.expectedCalls == 1 {
				is.Equal("a", result) // Still the first element
			} else {
				is.Equal("bb", result) // "bb" became max after second comparison
			}
		})
	}
}

func TestMaxIndexBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MaxIndexBy([]string{"s1", "string2", "s3"}, func(item, mAx string) bool {
		return len(item) > len(mAx)
	})
	result2, index2 := MaxIndexBy([]string{"string1", "string2", "s3"}, func(item, mAx string) bool {
		return len(item) > len(mAx)
	})
	result3, index3 := MaxIndexBy([]string{}, func(item, mAx string) bool {
		return len(item) > len(mAx)
	})

	is.Equal("string2", result1)
	is.Equal(1, index1)

	is.Equal("string1", result2)
	is.Zero(index2)

	is.Empty(result3)
	is.Equal(-1, index3)
}

func TestMaxIndexByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := assert.AnError

	// Test normal operation (no error) - table driven
	tests := []struct {
		name           string
		input          []string
		expectedResult string
		expectedIndex  int
	}{
		{
			name:           "finds max by length - second match",
			input:          []string{"s1", "string2", "s3"},
			expectedResult: "string2",
			expectedIndex:  1,
		},
		{
			name:           "finds max by length - first match",
			input:          []string{"string1", "string2", "s3"},
			expectedResult: "string1",
			expectedIndex:  0,
		},
		{
			name:           "empty collection",
			input:          []string{},
			expectedResult: "",
			expectedIndex:  -1,
		},
		{
			name:           "single element",
			input:          []string{"single"},
			expectedResult: "single",
			expectedIndex:  0,
		},
		{
			name:           "all equal length",
			input:          []string{"a", "b", "c"},
			expectedResult: "a", // first maximal value
			expectedIndex:  0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, index, err := MaxIndexByErr(tt.input, func(item, mAx string) (bool, error) {
				return len(item) > len(mAx), nil
			})
			is.NoError(err)
			is.Equal(tt.expectedResult, result)
			is.Equal(tt.expectedIndex, index)
		})
	}

	// Test error cases - table driven
	errorTests := []struct {
		name           string
		input          []string
		errorAt        string
		expectedCalls  int
		expectedResult string
		expectedIndex  int
	}{
		{
			name:           "error at second comparison",
			input:          []string{"a", "bb", "ccc"},
			errorAt:        "bb",
			expectedCalls:  1, // Only first comparison (initial element vs second)
			expectedResult: "",
			expectedIndex:  -1,
		},
		{
			name:           "error at third comparison",
			input:          []string{"a", "bb", "ccc"},
			errorAt:        "ccc",
			expectedCalls:  2, // First two comparisons
			expectedResult: "",
			expectedIndex:  -1,
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			result, index, err := MaxIndexByErr(tt.input, func(item, mAx string) (bool, error) {
				callbackCount++
				if item == tt.errorAt {
					return false, testErr
				}
				return len(item) > len(mAx), nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedCalls, callbackCount)
			is.Equal(tt.expectedResult, result)
			is.Equal(tt.expectedIndex, index)
		})
	}
}

func TestLatest(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	a := time.Now()
	b := a.Add(time.Hour)
	result1 := Latest(a, b)
	result2 := Latest()

	is.Equal(b, result1)
	is.Zero(result2)
}

func TestLatestBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type foo struct {
		bar time.Time
	}

	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	t3 := t1.Add(-time.Hour)
	result1 := LatestBy([]foo{{t1}, {t2}, {t3}}, func(i foo) time.Time {
		return i.bar
	})
	result2 := LatestBy([]foo{{t1}}, func(i foo) time.Time {
		return i.bar
	})
	result3 := LatestBy([]foo{}, func(i foo) time.Time {
		return i.bar
	})

	is.Equal(foo{t2}, result1)
	is.Equal(foo{t1}, result2)
	is.Zero(result3)
}

func TestLatestByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := assert.AnError

	type foo struct {
		bar time.Time
	}

	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	t3 := t1.Add(-time.Hour)

	// Test normal operation (no error) - table driven
	tests := []struct {
		name     string
		input    []foo
		expected foo
	}{
		{
			name:     "finds latest time",
			input:    []foo{{t1}, {t2}, {t3}},
			expected: foo{t2},
		},
		{
			name:     "single element",
			input:    []foo{{t1}},
			expected: foo{t1},
		},
		{
			name:     "empty collection",
			input:    []foo{},
			expected: foo{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := LatestByErr(tt.input, func(i foo) (time.Time, error) {
				return i.bar, nil
			})
			is.NoError(err)
			is.Equal(tt.expected, result)
		})
	}

	// Test error cases - table driven
	errorTests := []struct {
		name          string
		input         []foo
		errorAt       int
		expectedCalls int
	}{
		{
			name:          "error at first element",
			input:         []foo{{t1}, {t2}, {t3}},
			errorAt:       0,
			expectedCalls: 1, // Only first callback
		},
		{
			name:          "error at second element",
			input:         []foo{{t1}, {t2}, {t3}},
			errorAt:       1,
			expectedCalls: 2, // First two callbacks
		},
		{
			name:          "error at third element",
			input:         []foo{{t1}, {t2}, {t3}},
			errorAt:       2,
			expectedCalls: 3, // All three callbacks
		},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			_, err := LatestByErr(tt.input, func(i foo) (time.Time, error) {
				callbackCount++
				if len(tt.input) > 0 && i == tt.input[tt.errorAt] {
					return time.Time{}, testErr
				}
				return i.bar, nil
			})
			is.ErrorIs(err, testErr)
			is.Equal(tt.expectedCalls, callbackCount)
		})
	}
}

func TestFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := First([]int{1, 2, 3})
	result2, ok2 := First([]int{})

	is.Equal(1, result1)
	is.True(ok1)
	is.Zero(result2)
	is.False(ok2)
}

func TestFirstOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FirstOrEmpty([]int{1, 2, 3})
	result2 := FirstOrEmpty([]int{})
	result3 := FirstOrEmpty([]string{})

	is.Equal(1, result1)
	is.Zero(result2)
	is.Empty(result3)
}

func TestFirstOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FirstOr([]int{1, 2, 3}, 63)
	result2 := FirstOr([]int{}, 23)
	result3 := FirstOr([]string{}, "test")

	is.Equal(1, result1)
	is.Equal(23, result2)
	is.Equal("test", result3)
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := Last([]int{1, 2, 3})
	result2, ok2 := Last([]int{})

	is.Equal(3, result1)
	is.True(ok1)
	is.Zero(result2)
	is.False(ok2)
}

func TestLastOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastOrEmpty([]int{1, 2, 3})
	result2 := LastOrEmpty([]int{})
	result3 := LastOrEmpty([]string{})

	is.Equal(3, result1)
	is.Zero(result2)
	is.Empty(result3)
}

func TestLastOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastOr([]int{1, 2, 3}, 63)
	result2 := LastOr([]int{}, 23)
	result3 := LastOr([]string{}, "test")

	is.Equal(3, result1)
	is.Equal(23, result2)
	is.Equal("test", result3)
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := Nth([]int{0, 1, 2, 3}, 2)
	result2, err2 := Nth([]int{0, 1, 2, 3}, -2)
	result3, err3 := Nth([]int{0, 1, 2, 3}, 42)
	result4, err4 := Nth([]int{}, 0)
	result5, err5 := Nth([]int{42}, 0)
	result6, err6 := Nth([]int{42}, -1)

	is.Equal(2, result1)
	is.NoError(err1)
	is.Equal(2, result2)
	is.NoError(err2)
	is.Zero(result3)
	is.EqualError(err3, "nth: 42 out of slice bounds")
	is.Zero(result4)
	is.EqualError(err4, "nth: 0 out of slice bounds")
	is.Equal(42, result5)
	is.NoError(err5)
	is.Equal(42, result6)
	is.NoError(err6)
}

func TestNthOr(t *testing.T) {
	t.Parallel()

	t.Run("Integers", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		const defaultValue = -1
		intSlice := []int{10, 20, 30, 40, 50}

		is.Equal(30, NthOr(intSlice, 2, defaultValue))
		is.Equal(50, NthOr(intSlice, -1, defaultValue))
		is.Equal(defaultValue, NthOr(intSlice, 5, defaultValue))
	})

	t.Run("Strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		const defaultValue = "none"
		strSlice := []string{"apple", "banana", "cherry", "date"}

		is.Equal("banana", NthOr(strSlice, 1, defaultValue))      // Index 1, expected "banana"
		is.Equal("cherry", NthOr(strSlice, -2, defaultValue))     // Negative index -2, expected "cherry"
		is.Equal(defaultValue, NthOr(strSlice, 10, defaultValue)) // Out of bounds, fallback "none"
	})

	t.Run("Structs", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type User struct {
			ID   int
			Name string
		}
		userSlice := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}

		expectedUser := User{ID: 1, Name: "Alice"}
		is.Equal(expectedUser, NthOr(userSlice, 0, User{ID: 0, Name: "Unknown"}))

		expectedUser = User{ID: 3, Name: "Charlie"}
		is.Equal(expectedUser, NthOr(userSlice, -1, User{ID: 0, Name: "Unknown"}))

		expectedUser = User{ID: 0, Name: "Unknown"}
		is.Equal(expectedUser, NthOr(userSlice, 10, User{ID: 0, Name: "Unknown"}))
	})
}

func TestNthOrEmpty(t *testing.T) {
	t.Parallel()

	t.Run("Integers", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		intSlice := []int{10, 20, 30, 40, 50}

		is.Equal(30, NthOrEmpty(intSlice, 2))
		is.Equal(50, NthOrEmpty(intSlice, -1))
		is.Zero(NthOrEmpty(intSlice, 10))
	})

	t.Run("Strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		strSlice := []string{"apple", "banana", "cherry", "date"}

		is.Equal("banana", NthOrEmpty(strSlice, 1))
		is.Equal("cherry", NthOrEmpty(strSlice, -2))
		is.Empty(NthOrEmpty(strSlice, 10))
	})

	t.Run("Structs", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type User struct {
			ID   int
			Name string
		}
		userSlice := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}

		expectedUser := User{ID: 1, Name: "Alice"}
		is.Equal(expectedUser, NthOrEmpty(userSlice, 0))

		expectedUser = User{ID: 3, Name: "Charlie"}
		is.Equal(expectedUser, NthOrEmpty(userSlice, -1))

		is.Zero(NthOrEmpty(userSlice, 10))
	})
}

func TestSample(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Sample([]string{"a", "b", "c"})
	result2 := Sample([]string{})

	is.True(Contains([]string{"a", "b", "c"}, result1))
	is.Empty(result2)
}

func TestSampleBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := rand.New(rand.NewSource(42))

	result1 := SampleBy([]string{"a", "b", "c"}, r.Intn)
	result2 := SampleBy([]string{}, rand.Intn)

	is.True(Contains([]string{"a", "b", "c"}, result1))
	is.Empty(result2)
}

func TestSamples(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Samples([]string{"a", "b", "c"}, 3)
	result2 := Samples([]string{}, 3)

	is.ElementsMatch(result1, []string{"a", "b", "c"})
	is.Empty(result2)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Samples(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestSamplesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := rand.New(rand.NewSource(42))

	result1 := SamplesBy([]string{"a", "b", "c"}, 3, r.Intn)
	result2 := SamplesBy([]string{}, 3, r.Intn)
	result3 := SamplesBy([]string{"a", "b", "c"}, 3, func(n int) int { return n - 1 })
	result4 := SamplesBy([]string{"a", "b", "c"}, 3, func(int) int { return 0 })
	result5 := SamplesBy([]string{"a", "b", "c"}, 0, func(int) int { return 1 })
	result6 := SamplesBy([]string{"a", "b", "c"}, -1, nil)

	// index out of range [1] with length 1
	is.Panics(func() {
		SamplesBy([]string{"a", "b", "c"}, 3, func(int) int { return 1 })
	})

	is.ElementsMatch(result1, []string{"a", "b", "c"})
	is.Empty(result2)
	is.Equal([]string{"c", "b", "a"}, result3)
	is.Equal([]string{"a", "c", "b"}, result4)
	is.Empty(result5)
	is.Empty(result6)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := SamplesBy(allStrings, 2, r.Intn)
	is.IsType(nonempty, allStrings, "type preserved")
}

// SamplesBy switches between two different algorithms depending on the
// count/size ratio: a map-based "sparse" selection (count <= size/16) and an
// index-slice-based "dense" selection (count > size/16). A collection small
// enough to fit the sparse branch's threshold with a non-trivial count
// (e.g. the {"a", "b", "c"} slices used above) never exercises the sparse
// branch at all, so it needs its own coverage on a large collection.
func TestSamplesBySparse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := Range(10_000)
	threshold := len(collection) / 16

	for _, count := range []int{1, 2, 10, threshold} {
		is.LessOrEqual(count, threshold, "sanity check: count must land in the sparse branch")

		for seed := int64(0); seed < 10; seed++ {
			r := rand.New(rand.NewSource(seed))
			result := SamplesBy(collection, count, r.Intn)

			is.Len(result, count)
			is.Len(Uniq(result), count, "sparse branch must not return duplicate elements")
			for _, v := range result {
				is.True(v >= 0 && v < len(collection), "sampled value must belong to the collection")
			}
		}
	}
}

// The sparse (map-based) and dense (index-slice-based) branches are two
// independent implementations of the same abstract algorithm: draw
// randomIntGenerator(n) for n = size, size-1, size-2, ... and swap the pick
// out of the remaining pool. Because both branches consume the generator
// with the exact same sequence of n values regardless of which one runs,
// seeding two generators identically and requesting a smaller sparse count
// and a larger dense count from the same collection must produce the same
// prefix of picks. This pins down the equivalence promised by the comment in
// find.go and would catch a regression in either branch that the two
// single-branch tests above could miss (e.g. one branch drifting out of
// sync with the other after an edit to just one of them).
func TestSamplesBySparseDenseEquivalence(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := Range(1_000)
	threshold := len(collection) / 16

	sparseCount := threshold     // <= threshold: takes the sparse branch
	denseCount := threshold + 10 // > threshold: takes the dense branch

	for seed := int64(0); seed < 10; seed++ {
		sparse := SamplesBy(collection, sparseCount, rand.New(rand.NewSource(seed)).Intn)
		dense := SamplesBy(collection, denseCount, rand.New(rand.NewSource(seed)).Intn)

		is.Equal(sparse, dense[:sparseCount])
	}
}

// Exercises the boundary of the count <= size/16 condition itself: one count
// value on each side of the threshold, on the very same collection, so a
// future change to the condition (e.g. an off-by-one on the comparison
// operator or the division) is caught even if each branch is otherwise
// individually correct.
func TestSamplesBySparseBoundary(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := Range(1_000)
	threshold := len(collection) / 16

	for _, count := range []int{threshold, threshold + 1} {
		result := SamplesBy(collection, count, rand.New(rand.NewSource(7)).Intn)

		is.Len(result, count)
		is.Len(Uniq(result), count)
	}
}
