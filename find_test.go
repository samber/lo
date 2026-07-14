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

	tests := []struct {
		name       string
		collection []int
		element    int
		expected   int
	}{
		{name: "element present", collection: []int{0, 1, 2, 1, 2, 3}, element: 2, expected: 2},
		{name: "element absent", collection: []int{0, 1, 2, 1, 2, 3}, element: 6, expected: -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, IndexOf(tt.collection, tt.element))
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		element    int
		expected   int
	}{
		{name: "element present", collection: []int{0, 1, 2, 1, 2, 3}, element: 2, expected: 4},
		{name: "element absent", collection: []int{0, 1, 2, 1, 2, 3}, element: 6, expected: -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, LastIndexOf(tt.collection, tt.element))
		})
	}
}

func TestHasPrefix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		prefix     []int
		expected   bool
	}{
		{name: "matching prefix", collection: []int{1, 2, 3, 4}, prefix: []int{1, 2}, expected: true},
		{name: "non-matching prefix", collection: []int{1, 2, 3, 4}, prefix: []int{42}, expected: false},
		{name: "nil prefix", collection: []int{1, 2, 3, 4}, prefix: nil, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, HasPrefix(tt.collection, tt.prefix))
		})
	}
}

func TestHasSuffix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		suffix     []int
		expected   bool
	}{
		{name: "matching suffix", collection: []int{1, 2, 3, 4}, suffix: []int{3, 4}, expected: true},
		{name: "non-matching suffix", collection: []int{1, 2, 3, 4}, suffix: []int{42}, expected: false},
		{name: "nil suffix", collection: []int{1, 2, 3, 4}, suffix: nil, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, HasSuffix(tt.collection, tt.suffix))
		})
	}
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name         string
		collection   []string
		wantSequence []string // items expected to be passed to the predicate, in order
		wantResult   string
		wantOk       bool
	}{
		{
			name:         "found in middle",
			collection:   []string{"a", "b", "c", "d"},
			wantSequence: []string{"a", "b"},
			wantResult:   "b",
			wantOk:       true,
		},
		{
			name:         "not found",
			collection:   []string{"foobar"},
			wantSequence: []string{"foobar"},
			wantResult:   "",
			wantOk:       false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			index := 0
			result, ok := Find(tt.collection, func(item string) bool {
				is.Equal(tt.wantSequence[index], item)
				index++
				return item == "b"
			})

			is.Equal(tt.wantOk, ok)
			is.Equal(tt.wantResult, result)
		})
	}
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

	tests := []struct {
		name         string
		collection   []string
		wantSequence []string // items expected to be passed to the predicate, in order
		wantItem     string
		wantIndex    int
		wantOk       bool
	}{
		{
			name:         "found in middle",
			collection:   []string{"a", "b", "c", "d", "b"},
			wantSequence: []string{"a", "b"},
			wantItem:     "b",
			wantIndex:    1,
			wantOk:       true,
		},
		{
			name:         "not found",
			collection:   []string{"foobar"},
			wantSequence: []string{"foobar"},
			wantItem:     "",
			wantIndex:    -1,
			wantOk:       false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			index := 0
			item, itemIndex, ok := FindIndexOf(tt.collection, func(item string) bool {
				is.Equal(tt.wantSequence[index], item)
				index++
				return item == "b"
			})

			is.Equal(tt.wantItem, item)
			is.Equal(tt.wantOk, ok)
			is.Equal(tt.wantIndex, itemIndex)
		})
	}
}

func TestFindLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name         string
		collection   []string
		wantSequence []string // items expected to be passed to the predicate, in reverse order
		wantItem     string
		wantIndex    int
		wantOk       bool
	}{
		{
			name:         "found scanning from the end",
			collection:   []string{"a", "b", "c", "d", "b"},
			wantSequence: []string{"b", "d", "c", "b", "a"},
			wantItem:     "b",
			wantIndex:    4,
			wantOk:       true,
		},
		{
			name:         "not found",
			collection:   []string{"foobar"},
			wantSequence: []string{"foobar"},
			wantItem:     "",
			wantIndex:    -1,
			wantOk:       false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			index := 0
			item, itemIndex, ok := FindLastIndexOf(tt.collection, func(item string) bool {
				is.Equal(tt.wantSequence[index], item)
				index++
				return item == "b"
			})

			is.Equal(tt.wantItem, item)
			is.Equal(tt.wantOk, ok)
			is.Equal(tt.wantIndex, itemIndex)
		})
	}
}

func TestFindOrElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name         string
		collection   []string
		fallback     string
		wantSequence []string // items expected to be passed to the predicate, in order
		expected     string
	}{
		{
			name:         "found in middle",
			collection:   []string{"a", "b", "c", "d"},
			fallback:     "x",
			wantSequence: []string{"a", "b"},
			expected:     "b",
		},
		{
			name:         "not found falls back",
			collection:   []string{"foobar"},
			fallback:     "x",
			wantSequence: []string{"foobar"},
			expected:     "x",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			index := 0
			result := FindOrElse(tt.collection, tt.fallback, func(item string) bool {
				is.Equal(tt.wantSequence[index], item)
				index++
				return item == "b"
			})

			is.Equal(tt.expected, result)
		})
	}
}

func TestFindKey(t *testing.T) {
	t.Parallel()

	t.Run("int values", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name       string
			collection map[string]int
			value      int
			wantKey    string
			wantOk     bool
		}{
			{name: "value present", collection: map[string]int{"foo": 1, "bar": 2, "baz": 3}, value: 2, wantKey: "bar", wantOk: true},
			{name: "value absent", collection: map[string]int{"foo": 1, "bar": 2, "baz": 3}, value: 42, wantKey: "", wantOk: false},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				result, ok := FindKey(tt.collection, tt.value)
				is.Equal(tt.wantKey, result)
				is.Equal(tt.wantOk, ok)
			})
		}
	})

	t.Run("struct values", func(t *testing.T) {
		t.Parallel()

		type test struct {
			foobar string
		}

		tests := []struct {
			name       string
			collection map[string]test
			value      test
			wantKey    string
			wantOk     bool
		}{
			{name: "value present", collection: map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, value: test{"foo"}, wantKey: "foo", wantOk: true},
			{name: "value absent", collection: map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, value: test{"hello world"}, wantKey: "", wantOk: false},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				result, ok := FindKey(tt.collection, tt.value)
				is.Equal(tt.wantKey, result)
				is.Equal(tt.wantOk, ok)
			})
		}
	})
}

func TestFindKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		predicate func(k string, v int) bool
		wantKey   string
		wantOk    bool
	}{
		{
			name:      "predicate matches",
			predicate: func(k string, v int) bool { return k == "foo" },
			wantKey:   "foo",
			wantOk:    true,
		},
		{
			name:      "predicate matches nothing",
			predicate: func(k string, v int) bool { return false },
			wantKey:   "",
			wantOk:    false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, ok := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, tt.predicate)
			is.Equal(tt.wantKey, result)
			is.Equal(tt.wantOk, ok)
		})
	}
}

// TestFindUniques_smallScan exercises the small-scan path (all collections
// here are <= findSmallThreshold). See TestFindUniques_large for the
// map-based path.
func TestFindUniques_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		expected   []int
	}{
		{name: "all unique", collection: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "one unique among duplicates", collection: []int{1, 2, 2, 3, 1, 2}, expected: []int{3}},
		{name: "no unique", collection: []int{1, 2, 2, 1}, expected: nil},
		{name: "empty collection", collection: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindUniques(tt.collection)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := FindUniques(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// FindUniques dispatches on len(collection) <= findSmallThreshold (8): a
// collection of 12 elements forces the findUniquesLarge path, which the
// table above never exercises (its collections are all <= 6 elements).
func TestFindUniques_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		expected   []int
	}{
		{
			name:       "some unique",
			collection: []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10},
			expected:   []int{30, 50, 60, 70, 80, 90},
		},
		{
			name:       "no unique",
			collection: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
			expected:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Greater(len(tt.collection), findSmallThreshold, "sanity check: collection must exceed findSmallThreshold")

			result := FindUniques(tt.collection)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()

		type myInts []int
		allUnique := myInts{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		is.Greater(len(allUnique), findSmallThreshold, "sanity check: allUnique must exceed findSmallThreshold")
		nonempty := FindUniques(allUnique)
		is.Equal(myInts{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nonempty)
		is.IsType(nonempty, allUnique, "type preserved")
	})
}

// TestFindUniquesBy_smallScan exercises the small-scan path (all collections
// here are <= findSmallThreshold). See TestFindUniquesBy_large for the
// map-based path.
func TestFindUniquesBy_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	mod3 := func(i int) int { return i % 3 }

	tests := []struct {
		name       string
		collection []int
		expected   []int
	}{
		{name: "all unique", collection: []int{0, 1, 2}, expected: []int{0, 1, 2}},
		{name: "one unique among duplicates", collection: []int{0, 1, 2, 3, 4}, expected: []int{2}},
		{name: "no unique", collection: []int{0, 1, 2, 3, 4, 5}, expected: nil},
		{name: "empty collection", collection: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindUniquesBy(tt.collection, mod3)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := FindUniquesBy(allStrings, func(i string) string {
			return i
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// FindUniquesBy dispatches on len(collection) <= findSmallThreshold (8): a
// collection of 12 elements forces the findUniquesByLarge path, which the
// table above never exercises (its collections are all <= 6 elements).
func TestFindUniquesBy_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	byTen := func(v int) int { return v / 10 }

	tests := []struct {
		name       string
		collection []int
		expected   []int
	}{
		{
			name:       "some unique",
			collection: []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10},
			expected:   []int{30, 50, 60, 70, 80, 90},
		},
		{
			name:       "no unique",
			collection: []int{10, 11, 20, 21, 30, 31, 40, 41, 50, 51},
			expected:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Greater(len(tt.collection), findSmallThreshold, "sanity check: collection must exceed findSmallThreshold")

			result := FindUniquesBy(tt.collection, byTen)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()

		type myStrings []string
		allStrings := myStrings{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiiiii"}
		is.Greater(len(allStrings), findSmallThreshold, "sanity check: allStrings must exceed findSmallThreshold")
		nonempty := FindUniquesBy(allStrings, func(s string) int { return len(s) })
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// TestFindDuplicates_smallScan exercises the small-scan path (all
// collections here are <= findSmallThreshold). See
// TestFindDuplicates_large for the map-based path.
func TestFindDuplicates_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		expected   []int
	}{
		{name: "some duplicates", collection: []int{1, 2, 2, 1, 2, 3}, expected: []int{1, 2}},
		{name: "no duplicates", collection: []int{1, 2, 3}, expected: nil},
		{name: "empty collection", collection: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindDuplicates(tt.collection)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := FindDuplicates(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// FindDuplicates dispatches on len(collection) <= findSmallThreshold (8): a
// collection of 12 elements forces the findDuplicatesLarge path, which the
// table above never exercises (its collections are all <= 6 elements).
func TestFindDuplicates_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		expected   []int
	}{
		{
			name:       "some duplicates",
			collection: []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10},
			expected:   []int{10, 20, 40},
		},
		{
			name:       "no duplicates",
			collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Greater(len(tt.collection), findSmallThreshold, "sanity check: collection must exceed findSmallThreshold")

			result := FindDuplicates(tt.collection)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()

		type myStrings []string
		allStrings := myStrings{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
		is.Greater(len(allStrings), findSmallThreshold, "sanity check: allStrings must exceed findSmallThreshold")
		nonempty := FindDuplicates(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// TestFindDuplicatesBy_smallScan exercises the small-scan path (all
// collections here are <= findSmallThreshold). See
// TestFindDuplicatesBy_large for the map-based path.
func TestFindDuplicatesBy_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		by         func(i int) int
		expected   []int
	}{
		{name: "some duplicates", collection: []int{3, 4, 5, 6, 7}, by: func(i int) int { return i % 3 }, expected: []int{3, 4}},
		{name: "no duplicates", collection: []int{0, 1, 2, 3, 4}, by: func(i int) int { return i % 5 }, expected: nil},
		{name: "empty collection", collection: []int{}, by: func(i int) int { return i % 3 }, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindDuplicatesBy(tt.collection, tt.by)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := FindDuplicatesBy(allStrings, func(i string) string {
			return i
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// FindDuplicatesBy dispatches on len(collection) <= findSmallThreshold (8): a
// collection of 12 elements forces the findDuplicatesByLarge path, which the
// table above never exercises (its collections are all <= 5 elements).
func TestFindDuplicatesBy_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	byTen := func(v int) int { return v / 10 }

	tests := []struct {
		name       string
		collection []int
		expected   []int
	}{
		{
			name:       "some duplicates",
			collection: []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10},
			expected:   []int{10, 20, 40},
		},
		{
			name:       "no duplicates",
			collection: []int{10, 21, 32, 43, 54, 65, 76, 87, 98, 109},
			expected:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Greater(len(tt.collection), findSmallThreshold, "sanity check: collection must exceed findSmallThreshold")

			result := FindDuplicatesBy(tt.collection, byTen)
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()

		type myStrings []string
		allStrings := myStrings{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiiiii"}
		is.Greater(len(allStrings), findSmallThreshold, "sanity check: allStrings must exceed findSmallThreshold")
		nonempty := FindDuplicatesBy(allStrings, func(s string) int { return len(s) })
		is.IsType(nonempty, allStrings, "type preserved")
	})
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

	tests := []struct {
		name       string
		collection []int
		expected   int
	}{
		{name: "ascending", collection: []int{1, 2, 3}, expected: 1},
		{name: "descending", collection: []int{3, 2, 1}, expected: 1},
		{name: "empty collection", collection: []int{}, expected: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, Min(tt.collection))
		})
	}

	t.Run("time.Duration", func(t *testing.T) {
		t.Parallel()

		is.Equal(time.Second, Min([]time.Duration{time.Second, time.Minute, time.Hour}))
	})
}

func TestMinIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name          string
		collection    []int
		expected      int
		expectedIndex int
	}{
		{name: "ascending", collection: []int{1, 2, 3}, expected: 1, expectedIndex: 0},
		{name: "descending", collection: []int{3, 2, 1}, expected: 1, expectedIndex: 2},
		{name: "empty collection", collection: []int{}, expected: 0, expectedIndex: -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, index := MinIndex(tt.collection)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedIndex, index)
		})
	}

	t.Run("time.Duration", func(t *testing.T) {
		t.Parallel()

		result, index := MinIndex([]time.Duration{time.Second, time.Minute, time.Hour})
		is.Equal(time.Second, result)
		is.Zero(index)
	})
}

func TestMinBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	less := func(item, mIn string) bool { return len(item) < len(mIn) }

	tests := []struct {
		name       string
		collection []string
		expected   string
	}{
		{name: "first is shortest", collection: []string{"s1", "string2", "s3"}, expected: "s1"},
		{name: "last is shortest", collection: []string{"string1", "string2", "s3"}, expected: "s3"},
		{name: "empty collection", collection: []string{}, expected: ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, MinBy(tt.collection, less))
		})
	}
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

	less := func(item, mIn string) bool { return len(item) < len(mIn) }

	tests := []struct {
		name          string
		collection    []string
		expected      string
		expectedIndex int
	}{
		{name: "first is shortest", collection: []string{"s1", "string2", "s3"}, expected: "s1", expectedIndex: 0},
		{name: "last is shortest", collection: []string{"string1", "string2", "s3"}, expected: "s3", expectedIndex: 2},
		{name: "empty collection", collection: []string{}, expected: "", expectedIndex: -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, index := MinIndexBy(tt.collection, less)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedIndex, index)
		})
	}
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

	tests := []struct {
		name     string
		items    []time.Time
		expected time.Time
	}{
		{name: "two items", items: []time.Time{a, b}, expected: a},
		{name: "no items", items: nil, expected: time.Time{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, Earliest(tt.items...))
		})
	}
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
	by := func(i foo) time.Time { return i.bar }

	tests := []struct {
		name       string
		collection []foo
		expected   foo
	}{
		{name: "earliest in middle", collection: []foo{{t1}, {t2}, {t3}}, expected: foo{t3}},
		{name: "single element", collection: []foo{{t1}}, expected: foo{t1}},
		{name: "empty collection", collection: []foo{}, expected: foo{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, EarliestBy(tt.collection, by))
		})
	}
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

	tests := []struct {
		name       string
		collection []int
		expected   int
	}{
		{name: "ascending", collection: []int{1, 2, 3}, expected: 3},
		{name: "descending", collection: []int{3, 2, 1}, expected: 3},
		{name: "empty collection", collection: []int{}, expected: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, Max(tt.collection))
		})
	}

	t.Run("time.Duration", func(t *testing.T) {
		t.Parallel()

		is.Equal(time.Hour, Max([]time.Duration{time.Second, time.Minute, time.Hour}))
	})
}

func TestMaxIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name          string
		collection    []int
		expected      int
		expectedIndex int
	}{
		{name: "ascending", collection: []int{1, 2, 3}, expected: 3, expectedIndex: 2},
		{name: "descending", collection: []int{3, 2, 1}, expected: 3, expectedIndex: 0},
		{name: "empty collection", collection: []int{}, expected: 0, expectedIndex: -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, index := MaxIndex(tt.collection)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedIndex, index)
		})
	}

	t.Run("time.Duration", func(t *testing.T) {
		t.Parallel()

		result, index := MaxIndex([]time.Duration{time.Second, time.Minute, time.Hour})
		is.Equal(time.Hour, result)
		is.Equal(2, index)
	})
}

func TestMaxBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	greater := func(item, mAx string) bool { return len(item) > len(mAx) }

	tests := []struct {
		name       string
		collection []string
		expected   string
	}{
		{name: "second is longest", collection: []string{"s1", "string2", "s3"}, expected: "string2"},
		{name: "first is longest", collection: []string{"string1", "string2", "s3"}, expected: "string1"},
		{name: "empty collection", collection: []string{}, expected: ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, MaxBy(tt.collection, greater))
		})
	}
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

	greater := func(item, mAx string) bool { return len(item) > len(mAx) }

	tests := []struct {
		name          string
		collection    []string
		expected      string
		expectedIndex int
	}{
		{name: "second is longest", collection: []string{"s1", "string2", "s3"}, expected: "string2", expectedIndex: 1},
		{name: "first is longest", collection: []string{"string1", "string2", "s3"}, expected: "string1", expectedIndex: 0},
		{name: "empty collection", collection: []string{}, expected: "", expectedIndex: -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, index := MaxIndexBy(tt.collection, greater)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedIndex, index)
		})
	}
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

	tests := []struct {
		name     string
		items    []time.Time
		expected time.Time
	}{
		{name: "two items", items: []time.Time{a, b}, expected: b},
		{name: "no items", items: nil, expected: time.Time{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, Latest(tt.items...))
		})
	}
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
	by := func(i foo) time.Time { return i.bar }

	tests := []struct {
		name       string
		collection []foo
		expected   foo
	}{
		{name: "latest in middle", collection: []foo{{t1}, {t2}, {t3}}, expected: foo{t2}},
		{name: "single element", collection: []foo{{t1}}, expected: foo{t1}},
		{name: "empty collection", collection: []foo{}, expected: foo{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, LatestBy(tt.collection, by))
		})
	}
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

	tests := []struct {
		name       string
		collection []int
		expected   int
		expectedOk bool
	}{
		{name: "non-empty collection", collection: []int{1, 2, 3}, expected: 1, expectedOk: true},
		{name: "empty collection", collection: []int{}, expected: 0, expectedOk: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, ok := First(tt.collection)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedOk, ok)
		})
	}
}

func TestFirstOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		expected   int
	}{
		{name: "non-empty collection", collection: []int{1, 2, 3}, expected: 1},
		{name: "empty collection", collection: []int{}, expected: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, FirstOrEmpty(tt.collection))
		})
	}

	t.Run("empty string collection", func(t *testing.T) {
		t.Parallel()

		is.Empty(FirstOrEmpty([]string{}))
	})
}

func TestFirstOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		fallback   int
		expected   int
	}{
		{name: "non-empty collection", collection: []int{1, 2, 3}, fallback: 63, expected: 1},
		{name: "empty collection", collection: []int{}, fallback: 23, expected: 23},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, FirstOr(tt.collection, tt.fallback))
		})
	}

	t.Run("empty string collection", func(t *testing.T) {
		t.Parallel()

		is.Equal("test", FirstOr([]string{}, "test"))
	})
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		expected   int
		expectedOk bool
	}{
		{name: "non-empty collection", collection: []int{1, 2, 3}, expected: 3, expectedOk: true},
		{name: "empty collection", collection: []int{}, expected: 0, expectedOk: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, ok := Last(tt.collection)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedOk, ok)
		})
	}
}

func TestLastOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		expected   int
	}{
		{name: "non-empty collection", collection: []int{1, 2, 3}, expected: 3},
		{name: "empty collection", collection: []int{}, expected: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, LastOrEmpty(tt.collection))
		})
	}

	t.Run("empty string collection", func(t *testing.T) {
		t.Parallel()

		is.Empty(LastOrEmpty([]string{}))
	})
}

func TestLastOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		fallback   int
		expected   int
	}{
		{name: "non-empty collection", collection: []int{1, 2, 3}, fallback: 63, expected: 3},
		{name: "empty collection", collection: []int{}, fallback: 23, expected: 23},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, LastOr(tt.collection, tt.fallback))
		})
	}

	t.Run("empty string collection", func(t *testing.T) {
		t.Parallel()

		is.Equal("test", LastOr([]string{}, "test"))
	})
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name        string
		collection  []int
		nth         int
		expected    int
		expectedErr string
	}{
		{name: "positive index", collection: []int{0, 1, 2, 3}, nth: 2, expected: 2, expectedErr: ""},
		{name: "negative index", collection: []int{0, 1, 2, 3}, nth: -2, expected: 2, expectedErr: ""},
		{name: "positive index out of bounds", collection: []int{0, 1, 2, 3}, nth: 42, expected: 0, expectedErr: "nth: 42 out of slice bounds"},
		{name: "empty collection", collection: []int{}, nth: 0, expected: 0, expectedErr: "nth: 0 out of slice bounds"},
		{name: "single element, index zero", collection: []int{42}, nth: 0, expected: 42, expectedErr: ""},
		{name: "single element, negative index", collection: []int{42}, nth: -1, expected: 42, expectedErr: ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := Nth(tt.collection, tt.nth)
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

	t.Run("non-empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Sample([]string{"a", "b", "c"})
		is.True(Contains([]string{"a", "b", "c"}, result))
	})

	t.Run("empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Sample([]string{})
		is.Empty(result)
	})
}

func TestSampleBy(t *testing.T) {
	t.Parallel()

	t.Run("non-empty collection with custom random source", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := rand.New(rand.NewSource(42))
		result := SampleBy([]string{"a", "b", "c"}, r.Intn)
		is.True(Contains([]string{"a", "b", "c"}, result))
	})

	t.Run("empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := SampleBy([]string{}, rand.Intn)
		is.Empty(result)
	})
}

func TestSamples(t *testing.T) {
	t.Parallel()

	t.Run("non-empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Samples([]string{"a", "b", "c"}, 3)
		is.ElementsMatch(result, []string{"a", "b", "c"})
	})

	t.Run("empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Samples([]string{}, 3)
		is.Empty(result)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := Samples(allStrings, 2)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestSamplesBy(t *testing.T) {
	t.Parallel()

	t.Run("non-empty collection with custom random source", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := rand.New(rand.NewSource(42))
		result := SamplesBy([]string{"a", "b", "c"}, 3, r.Intn)
		is.ElementsMatch(result, []string{"a", "b", "c"})
	})

	t.Run("empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := rand.New(rand.NewSource(42))
		result := SamplesBy([]string{}, 3, r.Intn)
		is.Empty(result)
	})

	t.Run("generator returning offset index", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := SamplesBy([]string{"a", "b", "c"}, 3, func(n int) int { return n - 1 })
		is.Equal([]string{"c", "b", "a"}, result)
	})

	t.Run("generator always returning zero", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := SamplesBy([]string{"a", "b", "c"}, 3, func(int) int { return 0 })
		is.Equal([]string{"a", "c", "b"}, result)
	})

	t.Run("zero count", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := SamplesBy([]string{"a", "b", "c"}, 0, func(int) int { return 1 })
		is.Empty(result)
	})

	t.Run("negative count with nil generator", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := SamplesBy([]string{"a", "b", "c"}, -1, nil)
		is.Empty(result)
	})

	t.Run("panics on out-of-range index", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		// index out of range [1] with length 1
		is.Panics(func() {
			SamplesBy([]string{"a", "b", "c"}, 3, func(int) int { return 1 })
		})
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := rand.New(rand.NewSource(42))
		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := SamplesBy(allStrings, 2, r.Intn)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// SamplesBy switches between two different algorithms depending on the
// count/size ratio: a map-based "sparse" selection (count <= size/16) and an
// index-slice-based "dense" selection (count > size/16). A collection small
// enough to fit the sparse branch's threshold with a non-trivial count
// (e.g. the {"a", "b", "c"} slices used above) never exercises the sparse
// branch at all, so it needs its own coverage on a large collection.
func TestSamplesBy_sparse(t *testing.T) {
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
func TestSamplesBy_sparseDenseEquivalence(t *testing.T) {
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
func TestSamplesBy_sparseBoundary(t *testing.T) {
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
