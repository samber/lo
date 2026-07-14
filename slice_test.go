package lo

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		r1 := Filter([]int{1, 2, 3, 4}, func(x, _ int) bool {
			return x%2 == 0
		})
		assert.Equal(t, []int{2, 4}, r1)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		r2 := Filter([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
			return len(x) > 0
		})
		assert.Equal(t, []string{"foo", "bar"}, r2)
	})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Filter(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFilterErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     []int
		predicate func(item, index int) (bool, error)
		want      []int
		wantErr   string
		callbacks int // Number of predicates called before error/finish
	}{
		{
			name:  "filter even numbers",
			input: []int{1, 2, 3, 4},
			predicate: func(x, _ int) (bool, error) {
				return x%2 == 0, nil
			},
			want:      []int{2, 4},
			callbacks: 4,
		},
		{
			name:  "empty slice",
			input: []int{},
			predicate: func(x, _ int) (bool, error) {
				return true, nil
			},
			want:      []int{},
			callbacks: 0,
		},
		{
			name:  "filter all out",
			input: []int{1, 2, 3, 4},
			predicate: func(x, _ int) (bool, error) {
				return false, nil
			},
			want:      []int{},
			callbacks: 4,
		},
		{
			name:  "filter all in",
			input: []int{1, 2, 3, 4},
			predicate: func(x, _ int) (bool, error) {
				return true, nil
			},
			want:      []int{1, 2, 3, 4},
			callbacks: 4,
		},
		{
			name:  "error on specific index",
			input: []int{1, 2, 3, 4},
			predicate: func(x, _ int) (bool, error) {
				if x == 3 {
					return false, errors.New("number 3 is not allowed")
				}
				return x%2 == 0, nil
			},
			callbacks: 3,
			wantErr:   "number 3 is not allowed",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var callbacks int
			wrappedPredicate := func(item, index int) (bool, error) {
				callbacks++
				return tt.predicate(item, index)
			}

			got, err := FilterErr(tt.input, wrappedPredicate)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
				is.Equal(tt.callbacks, callbacks, "callback count should match expected early return")
			} else {
				is.NoError(err)
				is.Equal(tt.want, got)
				is.Equal(tt.callbacks, callbacks)
			}
		})
	}

	// Test type preservation
	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty, err := FilterErr(allStrings, func(x string, _ int) (bool, error) {
		return len(x) > 0, nil
	})
	is.NoError(err)
	is.IsType(nonempty, allStrings, "type preserved")
	is.Equal(myStrings{"foo", "bar"}, nonempty)
}

func TestMap(t *testing.T) {
	t.Parallel()

	t.Run("int to string", func(t *testing.T) {
		t.Parallel()
		result1 := Map([]int{1, 2, 3, 4}, func(x, _ int) string {
			return "Hello"
		})
		assert.Equal(t, []string{"Hello", "Hello", "Hello", "Hello"}, result1)
	})

	t.Run("int64 to string", func(t *testing.T) {
		t.Parallel()
		result2 := Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
			return strconv.FormatInt(x, 10)
		})
		assert.Equal(t, []string{"1", "2", "3", "4"}, result2)
	})
}

func TestMapErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name                  string
		input                 []int
		transform             func(item, index int) (string, error)
		wantResult            []string
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:  "successful transformation",
			input: []int{1, 2, 3, 4},
			transform: func(x, _ int) (string, error) {
				return strconv.Itoa(x), nil
			},
			wantResult:            []string{"1", "2", "3", "4"},
			wantErr:               false,
			expectedCallbackCount: 4,
		},
		{
			name:  "error at third element stops iteration",
			input: []int{1, 2, 3, 4},
			transform: func(x, _ int) (string, error) {
				if x == 3 {
					return "", errors.New("number 3 is not allowed")
				}
				return strconv.Itoa(x), nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 3 is not allowed",
			expectedCallbackCount: 3,
		},
		{
			name:  "error at first element stops iteration immediately",
			input: []int{1, 2, 3, 4},
			transform: func(x, _ int) (string, error) {
				if x == 1 {
					return "", errors.New("number 1 is not allowed")
				}
				return strconv.Itoa(x), nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 1 is not allowed",
			expectedCallbackCount: 1,
		},
		{
			name:  "error at last element",
			input: []int{1, 2, 3, 4},
			transform: func(x, _ int) (string, error) {
				if x == 4 {
					return "", errors.New("number 4 is not allowed")
				}
				return strconv.Itoa(x), nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 4 is not allowed",
			expectedCallbackCount: 4,
		},
		{
			name:  "empty input slice",
			input: []int{},
			transform: func(x, _ int) (string, error) {
				return strconv.Itoa(x), nil
			},
			wantResult:            []string{},
			wantErr:               false,
			expectedCallbackCount: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to test early return
			callbackCount := 0
			wrappedTransform := func(item, index int) (string, error) {
				callbackCount++
				return tt.transform(item, index)
			}

			result, err := MapErr(tt.input, wrappedTransform)

			if tt.wantErr {
				is.Error(err)
				is.Equal(tt.errMsg, err.Error())
				is.Nil(result)
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestUniqMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type User struct {
		Name string
		age  int
	}

	users := []User{{Name: "Alice", age: 20}, {Name: "Alex", age: 21}, {Name: "Alex", age: 22}}
	result := UniqMap(users, func(item User, index int) string {
		return item.Name
	})

	sort.Strings(result)

	is.Equal([]string{"Alex", "Alice"}, result)
}

func TestFilterMap(t *testing.T) {
	t.Parallel()

	t.Run("int64 slice", func(t *testing.T) {
		t.Parallel()
		r1 := FilterMap([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, bool) {
			if x%2 == 0 {
				return strconv.FormatInt(x, 10), true
			}
			return "", false
		})
		assert.Equal(t, []string{"2", "4"}, r1)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		r2 := FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
			if strings.HasSuffix(x, "pu") {
				return "xpu", true
			}
			return "", false
		})
		assert.Equal(t, []string{"xpu", "xpu"}, r2)
	})
}

func TestFlatMap(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		result1 := FlatMap([]int{0, 1, 2, 3, 4}, func(x, _ int) []string {
			return []string{"Hello"}
		})
		assert.Equal(t, []string{"Hello", "Hello", "Hello", "Hello", "Hello"}, result1)
	})

	t.Run("int64 slice", func(t *testing.T) {
		t.Parallel()
		result2 := FlatMap([]int64{0, 1, 2, 3, 4}, func(x int64, _ int) []string {
			result := make([]string, 0, x)
			for i := int64(0); i < x; i++ {
				result = append(result, strconv.FormatInt(x, 10))
			}
			return result
		})
		assert.Equal(t, []string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"}, result2)
	})
}

func TestFlatMapErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name                  string
		input                 []int64
		transform             func(item int64, index int) ([]string, error)
		wantResult            []string
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:  "successful transformation",
			input: []int64{0, 1, 2},
			transform: func(x int64, _ int) ([]string, error) {
				return []string{strconv.FormatInt(x, 10), strconv.FormatInt(x, 10)}, nil
			},
			wantResult:            []string{"0", "0", "1", "1", "2", "2"},
			wantErr:               false,
			expectedCallbackCount: 3,
		},
		{
			name:  "error at second element stops iteration",
			input: []int64{0, 1, 2, 3},
			transform: func(x int64, _ int) ([]string, error) {
				if x == 1 {
					return nil, errors.New("number 1 is not allowed")
				}
				return []string{strconv.FormatInt(x, 10)}, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 1 is not allowed",
			expectedCallbackCount: 2,
		},
		{
			name:  "error at first element stops iteration immediately",
			input: []int64{0, 1, 2, 3},
			transform: func(x int64, _ int) ([]string, error) {
				if x == 0 {
					return nil, errors.New("number 0 is not allowed")
				}
				return []string{strconv.FormatInt(x, 10)}, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 0 is not allowed",
			expectedCallbackCount: 1,
		},
		{
			name:  "error at last element",
			input: []int64{0, 1, 2},
			transform: func(x int64, _ int) ([]string, error) {
				if x == 2 {
					return nil, errors.New("number 2 is not allowed")
				}
				return []string{strconv.FormatInt(x, 10)}, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 2 is not allowed",
			expectedCallbackCount: 3,
		},
		{
			name:  "empty input slice",
			input: []int64{},
			transform: func(x int64, _ int) ([]string, error) {
				return []string{strconv.FormatInt(x, 10)}, nil
			},
			wantResult:            []string{},
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:  "returns empty slice for each element",
			input: []int64{1, 2, 3},
			transform: func(x int64, _ int) ([]string, error) {
				return []string{}, nil
			},
			wantResult:            []string{},
			wantErr:               false,
			expectedCallbackCount: 3,
		},
		{
			name:  "returns nil for some elements",
			input: []int64{1, 2, 3},
			transform: func(x int64, _ int) ([]string, error) {
				if x == 2 {
					return nil, nil
				}
				return []string{strconv.FormatInt(x, 10)}, nil
			},
			wantResult:            []string{"1", "3"},
			wantErr:               false,
			expectedCallbackCount: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to test early return
			callbackCount := 0
			wrappedTransform := func(item int64, index int) ([]string, error) {
				callbackCount++
				return tt.transform(item, index)
			}

			result, err := FlatMapErr(tt.input, wrappedTransform)

			if tt.wantErr {
				is.Error(err)
				is.Equal(tt.errMsg, err.Error())
				is.Nil(result)
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestTimes(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	is.Equal([]string{"0", "1", "2"}, result1)
}

func TestReduce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		initial  int
		expected int
	}{
		{name: "initial 0", initial: 0, expected: 10},
		{name: "initial 10", initial: 10, expected: 20},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Reduce([]int{1, 2, 3, 4}, func(agg, item, _ int) int {
				return agg + item
			}, tt.initial)

			is.Equal(tt.expected, result)
		})
	}
}

func TestReduceErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name                  string
		input                 []int
		accumulator           func(agg, item, index int) (int, error)
		initial               int
		wantResult            int
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:  "successful reduction",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				return agg + item, nil
			},
			initial:               0,
			wantResult:            10,
			wantErr:               false,
			expectedCallbackCount: 4,
		},
		{
			name:  "error at third element stops iteration",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				if item == 3 {
					return 0, errors.New("number 3 is not allowed")
				}
				return agg + item, nil
			},
			initial:               0,
			wantResult:            0,
			wantErr:               true,
			errMsg:                "number 3 is not allowed",
			expectedCallbackCount: 3,
		},
		{
			name:  "error at first element stops iteration immediately",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				if item == 1 {
					return 0, errors.New("number 1 is not allowed")
				}
				return agg + item, nil
			},
			initial:               0,
			wantResult:            0,
			wantErr:               true,
			errMsg:                "number 1 is not allowed",
			expectedCallbackCount: 1,
		},
		{
			name:  "error at last element",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				if item == 4 {
					return 0, errors.New("number 4 is not allowed")
				}
				return agg + item, nil
			},
			initial:               0,
			wantResult:            0,
			wantErr:               true,
			errMsg:                "number 4 is not allowed",
			expectedCallbackCount: 4,
		},
		{
			name:  "empty input slice",
			input: []int{},
			accumulator: func(agg, item, _ int) (int, error) {
				return agg + item, nil
			},
			initial:               10,
			wantResult:            10,
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:  "with non-zero initial value",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				return agg + item, nil
			},
			initial:               10,
			wantResult:            20,
			wantErr:               false,
			expectedCallbackCount: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to test early return
			callbackCount := 0
			wrappedAccumulator := func(agg, item, index int) (int, error) {
				callbackCount++
				return tt.accumulator(agg, item, index)
			}

			result, err := ReduceErr(tt.input, wrappedAccumulator, tt.initial)

			if tt.wantErr {
				is.Error(err)
				is.Equal(tt.errMsg, err.Error())
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestReduceRight(t *testing.T) {
	t.Parallel()

	t.Run("slice of slices", func(t *testing.T) {
		t.Parallel()
		result1 := ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg, item []int, _ int) []int {
			return append(agg, item...)
		}, []int{})
		assert.Equal(t, []int{4, 5, 2, 3, 0, 1}, result1)
	})

	t.Run("named collection type", func(t *testing.T) {
		t.Parallel()
		type collection []int
		result3 := ReduceRight(collection{1, 2, 3, 4}, func(agg, item, _ int) int {
			return agg + item
		}, 10)
		assert.Equal(t, 20, result3)
	})
}

func TestReduceRightErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name                  string
		input                 []int
		accumulator           func(agg, item, index int) (int, error)
		initial               int
		wantResult            int
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:  "successful reduction",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				return agg + item, nil
			},
			initial:               0,
			wantResult:            10,
			wantErr:               false,
			expectedCallbackCount: 4,
		},
		{
			name:  "error at second element from right stops iteration",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				if item == 3 {
					return 0, errors.New("number 3 is not allowed")
				}
				return agg + item, nil
			},
			initial:               0,
			wantResult:            0,
			wantErr:               true,
			errMsg:                "number 3 is not allowed",
			expectedCallbackCount: 2,
		},
		{
			name:  "error at first element from right stops iteration immediately",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				if item == 4 {
					return 0, errors.New("number 4 is not allowed")
				}
				return agg + item, nil
			},
			initial:               0,
			wantResult:            0,
			wantErr:               true,
			errMsg:                "number 4 is not allowed",
			expectedCallbackCount: 1,
		},
		{
			name:  "error at last element from left",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				if item == 1 {
					return 0, errors.New("number 1 is not allowed")
				}
				return agg + item, nil
			},
			initial:               0,
			wantResult:            0,
			wantErr:               true,
			errMsg:                "number 1 is not allowed",
			expectedCallbackCount: 4,
		},
		{
			name:  "empty input slice",
			input: []int{},
			accumulator: func(agg, item, _ int) (int, error) {
				return agg + item, nil
			},
			initial:               10,
			wantResult:            10,
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:  "with non-zero initial value",
			input: []int{1, 2, 3, 4},
			accumulator: func(agg, item, _ int) (int, error) {
				return agg + item, nil
			},
			initial:               10,
			wantResult:            20,
			wantErr:               false,
			expectedCallbackCount: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to test early return
			callbackCount := 0
			wrappedAccumulator := func(agg, item, index int) (int, error) {
				callbackCount++
				return tt.accumulator(agg, item, index)
			}

			result, err := ReduceRightErr(tt.input, wrappedAccumulator, tt.initial)

			if tt.wantErr {
				is.Error(err)
				is.Equal(tt.errMsg, err.Error())
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestForEach(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	ForEach([]string{"a", "b", "c"}, func(item string, i int) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.Equal([]string{"a", "b", "c"}, callParams1)
	is.Equal([]int{0, 1, 2}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestForEachWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	var callParams1 []string
	var callParams2 []int

	ForEachWhile([]string{"a", "b", "c"}, func(item string, i int) bool {
		if item == "c" {
			return false
		}
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
		return true
	})

	is.Equal([]string{"a", "b"}, callParams1)
	is.Equal([]int{0, 1}, callParams2)
	is.IsIncreasing(callParams2)
}

// TestUniq_small exercises the small-scan path (all collections here are
// <= uniqSmallInputThreshold). See TestUniq_large for the map-based path.
func TestUniq_small(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Uniq([]int{1, 2, 2, 1})
	is.Equal([]int{1, 2}, result1)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Uniq(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

// Uniq dispatches on len(collection) <= uniqSmallInputThreshold (8): a
// collection of 12 elements forces the uniqLarge path, which the table
// above never exercises (its collections are all <= 4 elements).
func TestUniq_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := []int{1, 2, 2, 3, 4, 4, 5, 6, 7, 8, 9, 1}
	is.Greater(len(collection), uniqSmallInputThreshold, "sanity check: collection must exceed uniqSmallInputThreshold")
	is.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, Uniq(collection))

	type myStrings []string
	allStrings := myStrings{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	is.Greater(len(allStrings), uniqSmallInputThreshold, "sanity check: allStrings must exceed uniqSmallInputThreshold")
	nonempty := Uniq(allStrings)
	is.Equal(allStrings, nonempty)
	is.IsType(nonempty, allStrings, "type preserved")
}

// TestUniqBy_small exercises the small-scan path (all collections here are
// <= uniqSmallInputThreshold). See TestUniqBy_large for the map-based path.
func TestUniqBy_small(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := UniqBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	is.Equal([]int{0, 1, 2}, result1)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := UniqBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

// UniqBy dispatches on len(collection) <= uniqSmallInputThreshold (8): a
// collection of 12 elements forces the uniqByLarge path, which the table
// above never exercises (its collections are all <= 6 elements).
func TestUniqBy_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := []int{10, 20, 30, 20, 40, 50, 60, 70, 80, 90, 40, 10}
	is.Greater(len(collection), uniqSmallInputThreshold, "sanity check: collection must exceed uniqSmallInputThreshold")
	byTen := func(v int) int { return v / 10 }
	is.Equal([]int{10, 20, 30, 40, 50, 60, 70, 80, 90}, UniqBy(collection, byTen))

	type myStrings []string
	allStrings := myStrings{"a", "bb", "ccc", "dddd", "eeeee", "ffffff", "ggggggg", "hhhhhhhh", "iiiiiiiii"}
	is.Greater(len(allStrings), uniqSmallInputThreshold, "sanity check: allStrings must exceed uniqSmallInputThreshold")
	nonempty := UniqBy(allStrings, func(s string) int { return len(s) })
	is.Equal(allStrings, nonempty)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestIsUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "nil slice",
			input: nil,
			want:  true,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  true,
		},
		{
			name:  "single item",
			input: []int{1},
			want:  true,
		},
		{
			name:  "unique",
			input: []int{1, 2, 3},
			want:  true,
		},
		{
			name:  "non unique",
			input: []int{1, 2, 1},
			want:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.want, IsUniq(tt.input))
		})
	}
}

func TestIsUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "nil slice",
			input: nil,
			want:  true,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  true,
		},
		{
			name:  "single item",
			input: []int{1},
			want:  true,
		},
		{
			name:  "unique",
			input: []int{1, 2, 3},
			want:  true,
		},
		{
			name:  "non unique",
			input: []int{1, 2, 4},
			want:  false,
		},
	}

	iteratee := func(i int) int { return i % 3 }

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.want, IsUniqBy(tt.input, iteratee))
		})
	}
}

func TestUniqByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name                  string
		input                 []int
		iteratee              func(item int) (int, error)
		wantResult            []int
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:  "successful uniq",
			input: []int{0, 1, 2, 3, 4, 5},
			iteratee: func(i int) (int, error) {
				return i % 3, nil
			},
			wantResult:            []int{0, 1, 2},
			wantErr:               false,
			expectedCallbackCount: 6,
		},
		{
			name:  "error at fourth element stops iteration",
			input: []int{0, 1, 2, 3, 4, 5},
			iteratee: func(i int) (int, error) {
				if i == 3 {
					return 0, errors.New("number 3 is not allowed")
				}
				return i % 3, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 3 is not allowed",
			expectedCallbackCount: 4,
		},
		{
			name:  "error at first element stops iteration immediately",
			input: []int{0, 1, 2, 3, 4, 5},
			iteratee: func(i int) (int, error) {
				if i == 0 {
					return 0, errors.New("number 0 is not allowed")
				}
				return i % 3, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 0 is not allowed",
			expectedCallbackCount: 1,
		},
		{
			name:  "error at last element",
			input: []int{0, 1, 2, 3, 4, 5},
			iteratee: func(i int) (int, error) {
				if i == 5 {
					return 0, errors.New("number 5 is not allowed")
				}
				return i % 3, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 5 is not allowed",
			expectedCallbackCount: 6,
		},
		{
			name:  "empty input slice",
			input: []int{},
			iteratee: func(i int) (int, error) {
				return i % 3, nil
			},
			wantResult:            []int{},
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:  "all duplicates",
			input: []int{1, 1, 1, 1},
			iteratee: func(i int) (int, error) {
				return i % 3, nil
			},
			wantResult:            []int{1},
			wantErr:               false,
			expectedCallbackCount: 4,
		},
		{
			name:  "no duplicates",
			input: []int{0, 1, 2, 3},
			iteratee: func(i int) (int, error) {
				return i, nil
			},
			wantResult:            []int{0, 1, 2, 3},
			wantErr:               false,
			expectedCallbackCount: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to test early return
			callbackCount := 0
			wrappedIteratee := func(item int) (int, error) {
				callbackCount++
				return tt.iteratee(item)
			}

			result, err := UniqByErr(tt.input, wrappedIteratee)

			if tt.wantErr {
				is.Error(err)
				is.Equal(tt.errMsg, err.Error())
				is.Nil(result)
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	is.Equal(map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	}, result1)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := GroupBy(allStrings, func(i string) int {
		return 42
	})
	is.IsType(nonempty[42], allStrings, "type preserved")
}

func TestGroupByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name                  string
		input                 []int
		iteratee              func(item int) (int, error)
		wantResult            map[int][]int
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:  "successful grouping",
			input: []int{0, 1, 2, 3, 4, 5},
			iteratee: func(i int) (int, error) {
				return i % 3, nil
			},
			wantResult: map[int][]int{
				0: {0, 3},
				1: {1, 4},
				2: {2, 5},
			},
			wantErr:               false,
			expectedCallbackCount: 6,
		},
		{
			name:  "error at fourth element stops iteration",
			input: []int{0, 1, 2, 3, 4, 5},
			iteratee: func(i int) (int, error) {
				if i == 3 {
					return 0, errors.New("number 3 is not allowed")
				}
				return i % 3, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 3 is not allowed",
			expectedCallbackCount: 4,
		},
		{
			name:  "error at first element stops iteration immediately",
			input: []int{0, 1, 2, 3, 4, 5},
			iteratee: func(i int) (int, error) {
				if i == 0 {
					return 0, errors.New("number 0 is not allowed")
				}
				return i % 3, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 0 is not allowed",
			expectedCallbackCount: 1,
		},
		{
			name:  "error at last element",
			input: []int{0, 1, 2, 3, 4, 5},
			iteratee: func(i int) (int, error) {
				if i == 5 {
					return 0, errors.New("number 5 is not allowed")
				}
				return i % 3, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 5 is not allowed",
			expectedCallbackCount: 6,
		},
		{
			name:  "empty input slice",
			input: []int{},
			iteratee: func(i int) (int, error) {
				return i % 3, nil
			},
			wantResult:            map[int][]int{},
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:  "all elements in same group",
			input: []int{3, 6, 9, 12},
			iteratee: func(i int) (int, error) {
				return 0, nil
			},
			wantResult: map[int][]int{
				0: {3, 6, 9, 12},
			},
			wantErr:               false,
			expectedCallbackCount: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to test early return
			callbackCount := 0
			wrappedIteratee := func(item int) (int, error) {
				callbackCount++
				return tt.iteratee(item)
			}

			result, err := GroupByErr(tt.input, wrappedIteratee)

			if tt.wantErr {
				is.Error(err)
				is.Equal(tt.errMsg, err.Error())
				is.Nil(result)
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestGroupByMap(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		result1 := GroupByMap([]int{0, 1, 2, 3, 4, 5}, func(i int) (int, string) {
			return i % 3, strconv.Itoa(i)
		})
		assert.Equal(t, map[int][]string{
			0: {"0", "3"},
			1: {"1", "4"},
			2: {"2", "5"},
		}, result1)
	})

	t.Run("named int slice", func(t *testing.T) {
		t.Parallel()
		type myInt int
		type myInts []myInt
		result2 := GroupByMap(myInts{1, 0, 2, 3, 4, 5}, func(i myInt) (int, string) {
			return int(i % 3), strconv.Itoa(int(i))
		})
		assert.Equal(t, map[int][]string{
			0: {"0", "3"},
			1: {"1", "4"},
			2: {"2", "5"},
		}, result2)
	})

	t.Run("struct slice", func(t *testing.T) {
		t.Parallel()
		type product struct {
			ID         int64
			CategoryID int64
		}
		products := []product{
			{ID: 1, CategoryID: 1},
			{ID: 2, CategoryID: 1},
			{ID: 3, CategoryID: 2},
			{ID: 4, CategoryID: 3},
			{ID: 5, CategoryID: 3},
		}
		result3 := GroupByMap(products, func(item product) (int64, string) {
			return item.CategoryID, "Product " + strconv.FormatInt(item.ID, 10)
		})
		assert.Equal(t, map[int64][]string{
			1: {"Product 1", "Product 2"},
			2: {"Product 3"},
			3: {"Product 4", "Product 5"},
		}, result3)
	})
}

func TestGroupByMapErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name                  string
		input                 []int
		transform             func(item int) (int, int, error)
		wantResult            map[int][]int
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:  "successful grouping",
			input: []int{0, 1, 2, 3, 4, 5},
			transform: func(i int) (int, int, error) {
				return i % 3, i * 2, nil
			},
			wantResult: map[int][]int{
				0: {0, 6},
				1: {2, 8},
				2: {4, 10},
			},
			wantErr:               false,
			expectedCallbackCount: 6,
		},
		{
			name:  "error at fourth element stops iteration",
			input: []int{0, 1, 2, 3, 4, 5},
			transform: func(i int) (int, int, error) {
				if i == 3 {
					return 0, 0, errors.New("number 3 is not allowed")
				}
				return i % 3, i * 2, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 3 is not allowed",
			expectedCallbackCount: 4,
		},
		{
			name:  "error at first element stops iteration immediately",
			input: []int{0, 1, 2, 3, 4, 5},
			transform: func(i int) (int, int, error) {
				if i == 0 {
					return 0, 0, errors.New("number 0 is not allowed")
				}
				return i % 3, i * 2, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 0 is not allowed",
			expectedCallbackCount: 1,
		},
		{
			name:  "error at last element",
			input: []int{0, 1, 2, 3, 4, 5},
			transform: func(i int) (int, int, error) {
				if i == 5 {
					return 0, 0, errors.New("number 5 is not allowed")
				}
				return i % 3, i * 2, nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 5 is not allowed",
			expectedCallbackCount: 6,
		},
		{
			name:  "empty input slice",
			input: []int{},
			transform: func(i int) (int, int, error) {
				return i % 3, i * 2, nil
			},
			wantResult:            map[int][]int{},
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:  "all elements in same group",
			input: []int{3, 6, 9, 12},
			transform: func(i int) (int, int, error) {
				return 0, i, nil
			},
			wantResult: map[int][]int{
				0: {3, 6, 9, 12},
			},
			wantErr:               false,
			expectedCallbackCount: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to test early return
			callbackCount := 0
			wrappedTransform := func(item int) (int, int, error) {
				callbackCount++
				return tt.transform(item)
			}

			result, err := GroupByMapErr(tt.input, wrappedTransform)

			if tt.wantErr {
				is.Error(err)
				is.Equal(tt.errMsg, err.Error())
				is.Nil(result)
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestChunk(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		size     int
		expected [][]int
	}{
		{name: "even split", input: []int{0, 1, 2, 3, 4, 5}, size: 2, expected: [][]int{{0, 1}, {2, 3}, {4, 5}}},
		{name: "remainder", input: []int{0, 1, 2, 3, 4, 5, 6}, size: 2, expected: [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}},
		{name: "empty input", input: []int{}, size: 2, expected: nil},
		{name: "single element", input: []int{0}, size: 2, expected: [][]int{{0}}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Chunk(tt.input, tt.size)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	is.PanicsWithValue("lo.Chunk: size must be greater than 0", func() {
		Chunk([]int{0}, 0)
	})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Chunk(allStrings, 2)
	is.IsType(nonempty[0], allStrings, "type preserved")

	// appending to a chunk should not affect original slice
	original := []int{0, 1, 2, 3, 4, 5}
	result5 := Chunk(original, 2)
	result5[0] = append(result5[0], 6)
	is.Equal([]int{0, 1, 2, 3, 4, 5}, original)
}

func TestWindow(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		size     int
		expected [][]int
	}{
		{name: "size 3 exact", input: []int{1, 2, 3, 4, 5}, size: 3, expected: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}},
		{name: "size 3 with remainder", input: []int{1, 2, 3, 4, 5, 6}, size: 3, expected: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}},
		{name: "input smaller than size", input: []int{1, 2}, size: 3, expected: nil},
		{name: "input equal to size", input: []int{1, 2, 3}, size: 3, expected: [][]int{{1, 2, 3}}},
		{name: "size 1", input: []int{1, 2, 3, 4}, size: 1, expected: [][]int{{1}, {2}, {3}, {4}}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Window(tt.input, tt.size)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	is.PanicsWithValue("lo.Window: size must be greater than 0", func() {
		Window([]int{1, 2, 3}, 0)
	})

	is.PanicsWithValue("lo.Window: size must be greater than 0", func() {
		Window([]int{1, 2, 3}, -1)
	})

	type myStrings []string
	allStrings := myStrings{"a", "b", "c", "d"}
	windows := Window(allStrings, 2)
	is.IsType(windows[0], allStrings, "type preserved")
	is.Equal(myStrings{"a", "b"}, windows[0])

	// appending to a window should not affect original slice
	original := []int{1, 2, 3, 4, 5}
	windows2 := Window(original, 3)
	windows2[0] = append(windows2[0], 6)
	is.Equal([]int{1, 2, 3, 4, 5}, original)
}

func TestSliding(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		size     int
		step     int
		expected [][]int
	}{
		{name: "overlapping windows (step < size)", input: []int{1, 2, 3, 4, 5, 6}, size: 3, step: 1, expected: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}},
		{name: "non-overlapping windows (step == size, like Chunk)", input: []int{1, 2, 3, 4, 5, 6}, size: 3, step: 3, expected: [][]int{{1, 2, 3}, {4, 5, 6}}},
		{name: "step > size (skipping elements)", input: []int{1, 2, 3, 4, 5, 6, 7, 8}, size: 2, step: 3, expected: [][]int{{1, 2}, {4, 5}, {7, 8}}},
		{name: "single element windows", input: []int{1, 2, 3, 4}, size: 1, step: 1, expected: [][]int{{1}, {2}, {3}, {4}}},
		{name: "empty result when collection is too small", input: []int{1, 2}, size: 3, step: 1, expected: nil},
		{name: "step 2, size 2", input: []int{1, 2, 3, 4, 5, 6}, size: 2, step: 2, expected: [][]int{{1, 2}, {3, 4}, {5, 6}}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Sliding(tt.input, tt.size, tt.step)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	is.PanicsWithValue("lo.Sliding: size must be greater than 0", func() {
		Sliding([]int{1, 2, 3}, 0, 1)
	})

	is.PanicsWithValue("lo.Sliding: step must be greater than 0", func() {
		Sliding([]int{1, 2, 3}, 2, 0)
	})

	is.PanicsWithValue("lo.Sliding: step must be greater than 0", func() {
		Sliding([]int{1, 2, 3}, 2, -1)
	})

	type myStrings []string
	allStrings := myStrings{"a", "b", "c", "d", "e"}
	windows := Sliding(allStrings, 2, 2)
	is.IsType(windows[0], allStrings, "type preserved")
	is.Equal(myStrings{"a", "b"}, windows[0])

	// appending to a window should not affect original slice
	original := []int{1, 2, 3, 4, 5, 6}
	windows2 := Sliding(original, 2, 2)
	windows2[0] = append(windows2[0], 7)
	is.Equal([]int{1, 2, 3, 4, 5, 6}, original)
}

func TestPartitionBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	oddEven := func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	}

	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{name: "mixed values", input: []int{-2, -1, 0, 1, 2, 3, 4, 5}, expected: [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}}},
		{name: "empty input", input: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := PartitionBy(tt.input, oddEven)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := PartitionBy(allStrings, func(item string) int {
		return len(item)
	})
	is.IsType(nonempty[0], allStrings, "type preserved")
}

func TestPartitionByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	oddEven := func(x int) (string, error) {
		if x < 0 {
			return "negative", nil
		} else if x%2 == 0 {
			return "even", nil
		}
		return "odd", nil
	}

	tests := []struct {
		name                  string
		input                 []int
		iteratee              func(item int) (string, error)
		wantResult            [][]int
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:                  "successful partition",
			input:                 []int{-2, -1, 0, 1, 2, 3, 4, 5},
			iteratee:              oddEven,
			wantResult:            [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}},
			wantErr:               false,
			expectedCallbackCount: 8,
		},
		{
			name:  "error at fifth element stops iteration",
			input: []int{-2, -1, 0, 1, 2, 3},
			iteratee: func(x int) (string, error) {
				if x == 2 {
					return "", errors.New("number 2 is not allowed")
				}
				return oddEven(x)
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 2 is not allowed",
			expectedCallbackCount: 5,
		},
		{
			name:  "error at first element stops iteration immediately",
			input: []int{-2, -1, 0, 1},
			iteratee: func(x int) (string, error) {
				if x == -2 {
					return "", errors.New("number -2 is not allowed")
				}
				return oddEven(x)
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number -2 is not allowed",
			expectedCallbackCount: 1,
		},
		{
			name:  "error at last element",
			input: []int{-2, -1, 0, 1, 2},
			iteratee: func(x int) (string, error) {
				if x == 2 {
					return "", errors.New("number 2 is not allowed")
				}
				return oddEven(x)
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "number 2 is not allowed",
			expectedCallbackCount: 5,
		},
		{
			name:                  "empty input slice",
			input:                 []int{},
			iteratee:              oddEven,
			wantResult:            [][]int{},
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:  "all elements in same partition",
			input: []int{1, 3, 5},
			iteratee: func(x int) (string, error) {
				return "odd", nil
			},
			wantResult:            [][]int{{1, 3, 5}},
			wantErr:               false,
			expectedCallbackCount: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to test early return
			callbackCount := 0
			wrappedIteratee := func(item int) (string, error) {
				callbackCount++
				return tt.iteratee(item)
			}

			result, err := PartitionByErr(tt.input, wrappedIteratee)

			if tt.wantErr {
				is.Error(err)
				is.Equal(tt.errMsg, err.Error())
				is.Nil(result)
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Flatten([][]int{{0, 1}, {2, 3, 4, 5}})

	is.Equal([]int{0, 1, 2, 3, 4, 5}, result1)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Flatten([]myStrings{allStrings})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestConcat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Concat([][]int{{0, 1}, {2, 3, 4, 5}}...)

	is.Equal([]int{0, 1, 2, 3, 4, 5}, result1)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Concat([]myStrings{allStrings}...)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestInterleave(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testCases := []struct {
		name string
		in   [][]int
		want []int
	}{
		{
			"nil",
			[][]int{nil},
			[]int{},
		},
		{
			"empty",
			[][]int{},
			[]int{},
		},
		{
			"empties",
			[][]int{{}, {}},
			[]int{},
		},
		{
			"same length",
			[][]int{{1, 3, 5}, {2, 4, 6}},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"different length",
			[][]int{{1, 3, 5, 6}, {2, 4}},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"many slices",
			[][]int{{1}, {2, 5, 8}, {3, 6}, {4, 7, 9, 10}},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, Interleave(tc.in...))
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Interleave(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("non-empty slice", func(t *testing.T) {
		t.Parallel()
		result1 := Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		assert.NotEqual(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, result1)
		assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, result1)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		result2 := Shuffle([]int{})
		assert.Empty(t, result2)
	})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Shuffle(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "even length", input: []int{0, 1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1, 0}},
		{name: "odd length", input: []int{0, 1, 2, 3, 4, 5, 6}, expected: []int{6, 5, 4, 3, 2, 1, 0}},
		{name: "empty", input: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Reverse(tt.input)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Reverse(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFill(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []foo
		value    foo
		expected []foo
	}{
		{name: "non-empty slice", input: []foo{{"a"}, {"a"}}, value: foo{"b"}, expected: []foo{{"b"}, {"b"}}},
		{name: "empty slice", input: []foo{}, value: foo{"a"}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Fill(tt.input, tt.value)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		count    int
		value    foo
		expected []foo
	}{
		{name: "count 2", count: 2, value: foo{"a"}, expected: []foo{{"a"}, {"a"}}},
		{name: "count 0", count: 0, value: foo{"a"}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Repeat(tt.count, tt.value)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}
}

func TestRepeatBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	cb := func(i int) int {
		return int(math.Pow(float64(i), 2))
	}

	tests := []struct {
		name     string
		count    int
		expected []int
	}{
		{name: "count 0", count: 0, expected: nil},
		{name: "count 2", count: 2, expected: []int{0, 1}},
		{name: "count 5", count: 5, expected: []int{0, 1, 4, 9, 16}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := RepeatBy(tt.count, cb)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}
}

func TestRepeatByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testErr := errors.New("test error")

	// Table-driven tests
	tests := []struct {
		name                  string
		count                 int
		callback              func(index int) (int, error)
		wantResult            []int
		wantErr               bool
		expectedCallbackCount int
	}{
		{
			name:  "successful completion",
			count: 5,
			callback: func(i int) (int, error) {
				return i * i, nil
			},
			wantResult:            []int{0, 1, 4, 9, 16},
			wantErr:               false,
			expectedCallbackCount: 5,
		},
		{
			name:  "error at first iteration",
			count: 5,
			callback: func(i int) (int, error) {
				if i == 0 {
					return 0, testErr
				}
				return i * i, nil
			},
			wantResult:            nil,
			wantErr:               true,
			expectedCallbackCount: 1,
		},
		{
			name:  "error at third iteration",
			count: 5,
			callback: func(i int) (int, error) {
				if i == 2 {
					return 0, testErr
				}
				return i * i, nil
			},
			wantResult:            nil,
			wantErr:               true,
			expectedCallbackCount: 3,
		},
		{
			name:  "error at last iteration",
			count: 5,
			callback: func(i int) (int, error) {
				if i == 4 {
					return 0, testErr
				}
				return i * i, nil
			},
			wantResult:            nil,
			wantErr:               true,
			expectedCallbackCount: 5,
		},
		{
			name:  "zero count - empty result",
			count: 0,
			callback: func(i int) (int, error) {
				return i * i, nil
			},
			wantResult:            []int{},
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:  "single item success",
			count: 1,
			callback: func(i int) (int, error) {
				return 42, nil
			},
			wantResult:            []int{42},
			wantErr:               false,
			expectedCallbackCount: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Track callback count to verify early return
			callbackCount := 0
			wrappedCallback := func(i int) (int, error) {
				callbackCount++
				return tt.callback(i)
			}

			result, err := RepeatByErr(tt.count, wrappedCallback)

			if tt.wantErr {
				is.ErrorIs(err, testErr)
				is.Nil(result)
			} else {
				is.NoError(err)
				is.Equal(tt.wantResult, result)
			}

			// Verify callback count matches expected (tests early return)
			is.Equal(tt.expectedCallbackCount, callbackCount, "callback count should match expected")
		})
	}
}

func TestKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := KeyBy([]string{"a", "aa", "aaa"}, func(str string) int {
		return len(str)
	})

	is.Equal(map[int]string{1: "a", 2: "aa", 3: "aaa"}, result1)
}

func TestKeyByErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		input                 []string
		iteratee              func(item string) (int, error)
		wantResult            map[int]string
		wantErr               bool
		errMsg                string
		expectedCallbackCount int
	}{
		{
			name:                  "empty slice",
			input:                 []string{},
			iteratee:              func(s string) (int, error) { return len(s), nil },
			wantResult:            map[int]string{},
			wantErr:               false,
			expectedCallbackCount: 0,
		},
		{
			name:                  "success case",
			input:                 []string{"a", "aa", "aaa"},
			iteratee:              func(s string) (int, error) { return len(s), nil },
			wantResult:            map[int]string{1: "a", 2: "aa", 3: "aaa"},
			wantErr:               false,
			expectedCallbackCount: 3,
		},
		{
			name:  "error stops iteration - first item",
			input: []string{"a", "aa", "aaa"},
			iteratee: func(s string) (int, error) {
				return 0, fmt.Errorf("error on %s", s)
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "error on a",
			expectedCallbackCount: 1,
		},
		{
			name:  "error stops iteration - middle item",
			input: []string{"a", "aa", "aaa"},
			iteratee: func(s string) (int, error) {
				if s == "aa" {
					return 0, errors.New("middle error")
				}
				return len(s), nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "middle error",
			expectedCallbackCount: 2,
		},
		{
			name:  "error stops iteration - last item",
			input: []string{"a", "aa", "aaa"},
			iteratee: func(s string) (int, error) {
				if s == "aaa" {
					return 0, errors.New("last error")
				}
				return len(s), nil
			},
			wantResult:            nil,
			wantErr:               true,
			errMsg:                "last error",
			expectedCallbackCount: 3,
		},
		{
			name:                  "duplicate keys",
			input:                 []string{"a", "b", "c"},
			iteratee:              func(s string) (int, error) { return 1, nil },
			wantResult:            map[int]string{1: "c"},
			wantErr:               false,
			expectedCallbackCount: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callbackCount := 0
			wrappedIteratee := func(s string) (int, error) {
				callbackCount++
				return tt.iteratee(s)
			}

			result, err := KeyByErr(tt.input, wrappedIteratee)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Equal(t, tt.errMsg, err.Error())
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResult, result)
			}

			assert.Equal(t, tt.expectedCallbackCount, callbackCount, "callback count mismatch")
		})
	}
}

func TestAssociate(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int) {
		return f.baz, f.bar
	}
	testCases := []struct {
		in   []*foo
		want map[string]int
	}{
		{
			in:   []*foo{{baz: "apple", bar: 1}},
			want: map[string]int{"apple": 1},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			want: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			want: map[string]int{"apple": 2},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, Associate(tc.in, transform))
		})
	}
}

func TestAssociateI(t *testing.T) {
	t.Parallel()

	transform := func(s string, i int) (int, string) {
		return i % 2, s
	}
	testCases := []struct {
		in   []string
		want map[int]string
	}{
		{
			in:   []string{"zero"},
			want: map[int]string{0: "zero"},
		},
		{
			in:   []string{"zero", "one"},
			want: map[int]string{0: "zero", 1: "one"},
		},
		{
			in:   []string{"two", "one", "zero"},
			want: map[int]string{0: "zero", 1: "one"},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, AssociateI(tc.in, transform))
		})
	}
}

func TestSliceToMap(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int) {
		return f.baz, f.bar
	}
	testCases := []struct {
		in   []*foo
		want map[string]int
	}{
		{
			in:   []*foo{{baz: "apple", bar: 1}},
			want: map[string]int{"apple": 1},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			want: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			want: map[string]int{"apple": 2},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, SliceToMap(tc.in, transform))
		})
	}
}

func TestSliceToMapI(t *testing.T) {
	t.Parallel()

	transform := func(s string, i int) (int, string) {
		return i % 2, s
	}
	testCases := []struct {
		in   []string
		want map[int]string
	}{
		{
			in:   []string{"zero"},
			want: map[int]string{0: "zero"},
		},
		{
			in:   []string{"zero", "one"},
			want: map[int]string{0: "zero", 1: "one"},
		},
		{
			in:   []string{"two", "one", "zero"},
			want: map[int]string{0: "zero", 1: "one"},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, SliceToMapI(tc.in, transform))
		})
	}
}

func TestFilterSliceToMap(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int, bool) {
		return f.baz, f.bar, f.bar > 1
	}
	testCases := []struct {
		in   []*foo
		want map[string]int
	}{
		{
			in:   []*foo{{baz: "apple", bar: 1}},
			want: map[string]int{},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			want: map[string]int{"banana": 2},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			want: map[string]int{"apple": 2},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, FilterSliceToMap(tc.in, transform))
		})
	}
}

func TestFilterSliceToMapI(t *testing.T) {
	t.Parallel()

	transform := func(s string, i int) (int, string, bool) {
		return i % 5, s, i%2 == 0
	}
	testCases := []struct {
		in   []string
		want map[int]string
	}{
		{
			in:   []string{"zero"},
			want: map[int]string{0: "zero"},
		},
		{
			in:   []string{"zero", "one", "two", "three", "four"},
			want: map[int]string{0: "zero", 2: "two", 4: "four"},
		},
		{
			in:   []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
			want: map[int]string{0: "ten", 1: "six", 2: "two", 3: "eight", 4: "four"},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, FilterSliceToMapI(tc.in, transform))
		})
	}
}

func TestKeyify(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		expected map[int]struct{}
	}{
		{name: "distinct values", input: []int{1, 2, 3, 4}, expected: map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}},
		{name: "duplicate values", input: []int{1, 1, 1, 2}, expected: map[int]struct{}{1: {}, 2: {}}},
		{name: "empty", input: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Keyify(tt.input)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}
}

func TestDrop(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{name: "drop 0", n: 0, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop 1", n: 1, expected: []int{1, 2, 3, 4}},
		{name: "drop 2", n: 2, expected: []int{2, 3, 4}},
		{name: "drop 3", n: 3, expected: []int{3, 4}},
		{name: "drop 4", n: 4, expected: []int{4}},
		{name: "drop 5", n: 5, expected: nil},
		{name: "drop 6", n: 6, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Drop([]int{0, 1, 2, 3, 4}, tt.n)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	is.PanicsWithValue("lo.Drop: n must not be negative", func() {
		Drop([]int{0, 1, 2, 3, 4}, -1)
	})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Drop(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{name: "drop 0", n: 0, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop 1", n: 1, expected: []int{0, 1, 2, 3}},
		{name: "drop 2", n: 2, expected: []int{0, 1, 2}},
		{name: "drop 3", n: 3, expected: []int{0, 1}},
		{name: "drop 4", n: 4, expected: []int{0}},
		{name: "drop 5", n: 5, expected: nil},
		{name: "drop 6", n: 6, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := DropRight([]int{0, 1, 2, 3, 4}, tt.n)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	is.PanicsWithValue("lo.DropRight: n must not be negative", func() {
		DropRight([]int{0, 1, 2, 3, 4}, -1)
	})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := DropRight(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		predicate func(t int) bool
		expected  []int
	}{
		{name: "drop until 4", predicate: func(t int) bool { return t != 4 }, expected: []int{4, 5, 6}},
		{name: "drop all", predicate: func(t int) bool { return true }, expected: nil},
		{name: "drop none", predicate: func(t int) bool { return t == 10 }, expected: []int{0, 1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, tt.predicate)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := DropWhile(allStrings, func(t string) bool {
		return t != "foo"
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropRightWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		predicate func(t int) bool
		expected  []int
	}{
		{name: "drop right until 3", predicate: func(t int) bool { return t != 3 }, expected: []int{0, 1, 2, 3}},
		{name: "drop right until 1", predicate: func(t int) bool { return t != 1 }, expected: []int{0, 1}},
		{name: "drop none", predicate: func(t int) bool { return t == 10 }, expected: []int{0, 1, 2, 3, 4, 5, 6}},
		{name: "drop all", predicate: func(t int) bool { return t != 10 }, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, tt.predicate)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := DropRightWhile(allStrings, func(t string) bool {
		return t != "foo"
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestTake(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{name: "take 3", n: 3, expected: []int{0, 1, 2}},
		{name: "take 2", n: 2, expected: []int{0, 1}},
		{name: "take 1", n: 1, expected: []int{0}},
		{name: "take 0", n: 0, expected: nil},
		{name: "take exactly len", n: 5, expected: []int{0, 1, 2, 3, 4}},
		{name: "take more than len", n: 10, expected: []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Take([]int{0, 1, 2, 3, 4}, tt.n)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	is.PanicsWithValue("lo.Take: n must not be negative", func() {
		Take([]int{0, 1, 2, 3, 4}, -1)
	})

	type myStrings []string
	allStrings := myStrings{"foo", "bar", "baz"}
	taken := Take(allStrings, 2)
	is.IsType(taken, allStrings, "type preserved")
	is.Equal(myStrings{"foo", "bar"}, taken)
}

func TestTakeWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		predicate func(t int) bool
		expected  []int
	}{
		{name: "take while < 4", predicate: func(t int) bool { return t < 4 }, expected: []int{0, 1, 2, 3}},
		{name: "take all", predicate: func(t int) bool { return t < 10 }, expected: []int{0, 1, 2, 3, 4, 5, 6}},
		{name: "take none", predicate: func(t int) bool { return t < 0 }, expected: nil},
		{name: "take while != 3", predicate: func(t int) bool { return t != 3 }, expected: []int{0, 1, 2}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := TakeWhile([]int{0, 1, 2, 3, 4, 5, 6}, tt.predicate)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"foo", "bar", "baz", "qux"}
	taken := TakeWhile(allStrings, func(t string) bool {
		return t != "baz"
	})
	is.IsType(taken, allStrings, "type preserved")
	is.Equal(myStrings{"foo", "bar"}, taken)
}

func TestTakeFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	even := func(item, index int) bool {
		return item%2 == 0
	}

	tests := []struct {
		name      string
		input     []int
		n         int
		predicate func(item, index int) bool
		expected  []int
	}{
		{name: "take 2 even", input: []int{1, 2, 3, 4, 5, 6}, n: 2, predicate: even, expected: []int{2, 4}},
		{name: "take more than available", input: []int{1, 2, 3, 4, 5, 6}, n: 10, predicate: even, expected: []int{2, 4, 6}},
		{name: "take 0", input: []int{1, 2, 3, 4, 5, 6}, n: 0, predicate: even, expected: nil},
		{name: "no matches", input: []int{1, 3, 5}, n: 2, predicate: even, expected: nil},
		{name: "take odd", input: []int{1, 2, 3, 4, 5}, n: 1, predicate: func(item, index int) bool { return item%2 != 0 }, expected: []int{1}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := TakeFilter(tt.input, tt.n, tt.predicate)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	is.PanicsWithValue("lo.TakeFilter: n must not be negative", func() {
		TakeFilter([]int{1, 2, 3}, -1, func(item, index int) bool { return true })
	})

	type myStrings []string
	allStrings := myStrings{"foo", "bar", "baz", "qux"}
	filtered := TakeFilter(allStrings, 2, func(item string, index int) bool {
		return len(item) == 3
	})
	is.IsType(filtered, allStrings, "type preserved")
	is.Equal(myStrings{"foo", "bar"}, filtered)
}

func TestDropByIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		indexes  []int
		expected []int
	}{
		{name: "drop index 0", input: []int{0, 1, 2, 3, 4}, indexes: []int{0}, expected: []int{1, 2, 3, 4}},
		{name: "drop indexes 0,1,2", input: []int{0, 1, 2, 3, 4}, indexes: []int{0, 1, 2}, expected: []int{3, 4}},
		{name: "drop negative indexes -4,-2,-3", input: []int{0, 1, 2, 3, 4}, indexes: []int{-4, -2, -3}, expected: []int{0, 4}},
		{name: "drop duplicate negative index -4,-4", input: []int{0, 1, 2, 3, 4}, indexes: []int{-4, -4}, expected: []int{0, 2, 3, 4}},
		{name: "drop indexes 3,1,0", input: []int{0, 1, 2, 3, 4}, indexes: []int{3, 1, 0}, expected: []int{2, 4}},
		{name: "drop index 2", input: []int{0, 1, 2, 3, 4}, indexes: []int{2}, expected: []int{0, 1, 3, 4}},
		{name: "drop index 4", input: []int{0, 1, 2, 3, 4}, indexes: []int{4}, expected: []int{0, 1, 2, 3}},
		{name: "no indexes", input: []int{0, 1, 2, 3, 4}, indexes: nil, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop out of range index 5", input: []int{0, 1, 2, 3, 4}, indexes: []int{5}, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop out of range index 100", input: []int{0, 1, 2, 3, 4}, indexes: []int{100}, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop out of range index -100", input: []int{0, 1, 2, 3, 4}, indexes: []int{-100}, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop index -1", input: []int{0, 1, 2, 3, 4}, indexes: []int{-1}, expected: []int{0, 1, 2, 3}},
		{name: "drop indexes -1,4", input: []int{0, 1, 2, 3, 4}, indexes: []int{-1, 4}, expected: []int{0, 1, 2, 3}},
		{name: "drop indexes -100,4", input: []int{0, 1, 2, 3, 4}, indexes: []int{-100, 4}, expected: []int{0, 1, 2, 3}},
		{name: "empty input, drop 0,1", input: []int{}, indexes: []int{0, 1}, expected: nil},
		{name: "single element, drop 0,1", input: []int{42}, indexes: []int{0, 1}, expected: nil},
		{name: "single element, drop 1,0", input: []int{42}, indexes: []int{1, 0}, expected: nil},
		{name: "empty input, drop 1", input: []int{}, indexes: []int{1}, expected: nil},
		{name: "single element, drop 0", input: []int{1}, indexes: []int{0}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := DropByIndex(tt.input, tt.indexes...)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := DropByIndex(allStrings, 0)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		r1 := Reject([]int{1, 2, 3, 4}, func(x, _ int) bool {
			return x%2 == 0
		})
		assert.Equal(t, []int{1, 3}, r1)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		r2 := Reject([]string{"Smith", "foo", "Domin", "bar", "Olivia"}, func(x string, _ int) bool {
			return len(x) > 3
		})
		assert.Equal(t, []string{"foo", "bar"}, r2)
	})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Reject(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestRejectErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     []int
		predicate func(item, index int) (bool, error)
		want      []int
		wantErr   string
		callbacks int // Number of predicates called before error/finish
	}{
		{
			name:  "reject even numbers",
			input: []int{1, 2, 3, 4},
			predicate: func(x, _ int) (bool, error) {
				return x%2 == 0, nil
			},
			want:      []int{1, 3},
			callbacks: 4,
		},
		{
			name:  "empty slice",
			input: []int{},
			predicate: func(x, _ int) (bool, error) {
				return true, nil
			},
			want:      []int{},
			callbacks: 0,
		},
		{
			name:  "reject all out",
			input: []int{1, 2, 3, 4},
			predicate: func(x, _ int) (bool, error) {
				return false, nil
			},
			want:      []int{1, 2, 3, 4},
			callbacks: 4,
		},
		{
			name:  "reject all in",
			input: []int{1, 2, 3, 4},
			predicate: func(x, _ int) (bool, error) {
				return true, nil
			},
			want:      []int{},
			callbacks: 4,
		},
		{
			name:  "error on specific index",
			input: []int{1, 2, 3, 4},
			predicate: func(x, _ int) (bool, error) {
				if x == 3 {
					return false, errors.New("number 3 is not allowed")
				}
				return x%2 == 0, nil
			},
			callbacks: 3,
			wantErr:   "number 3 is not allowed",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var callbacks int
			wrappedPredicate := func(item, index int) (bool, error) {
				callbacks++
				return tt.predicate(item, index)
			}

			got, err := RejectErr(tt.input, wrappedPredicate)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
				is.Equal(tt.callbacks, callbacks, "callback count should match expected early return")
			} else {
				is.NoError(err)
				is.Equal(tt.want, got)
				is.Equal(tt.callbacks, callbacks)
			}
		})
	}

	// Test type preservation
	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty, err := RejectErr(allStrings, func(x string, _ int) (bool, error) {
		return len(x) > 0, nil
	})
	is.NoError(err)
	is.IsType(nonempty, allStrings, "type preserved")
	is.Equal(myStrings{""}, nonempty)
}

func TestRejectMap(t *testing.T) {
	t.Parallel()

	t.Run("int64 slice", func(t *testing.T) {
		t.Parallel()
		r1 := RejectMap([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, bool) {
			if x%2 == 0 {
				return strconv.FormatInt(x, 10), false
			}
			return "", true
		})
		assert.Equal(t, []string{"2", "4"}, r1)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		r2 := RejectMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
			if strings.HasSuffix(x, "pu") {
				return "xpu", false
			}
			return "", true
		})
		assert.Equal(t, []string{"xpu", "xpu"}, r2)
	})
}

func TestFilterReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		left1, right1 := FilterReject([]int{1, 2, 3, 4}, func(x, _ int) bool {
			return x%2 == 0
		})
		assert.Equal(t, []int{2, 4}, left1)
		assert.Equal(t, []int{1, 3}, right1)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		left2, right2 := FilterReject([]string{"Smith", "foo", "Domin", "bar", "Olivia"}, func(x string, _ int) bool {
			return len(x) > 3
		})
		assert.Equal(t, []string{"Smith", "Domin", "Olivia"}, left2)
		assert.Equal(t, []string{"foo", "bar"}, right2)
	})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	a, b := FilterReject(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.IsType(a, allStrings, "type preserved")
	is.IsType(b, allStrings, "type preserved")
}

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		value    int
		expected int
	}{
		{name: "value present twice", input: []int{1, 2, 1}, value: 1, expected: 2},
		{name: "value absent", input: []int{1, 2, 1}, value: 3, expected: 0},
		{name: "empty input", input: []int{}, value: 1, expected: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Count(tt.input, tt.value)

			is.Equal(tt.expected, result)
		})
	}
}

func TestCountBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     []int
		predicate func(i int) bool
		expected  int
	}{
		{name: "less than 2", input: []int{1, 2, 1}, predicate: func(i int) bool { return i < 2 }, expected: 2},
		{name: "greater than 2", input: []int{1, 2, 1}, predicate: func(i int) bool { return i > 2 }, expected: 0},
		{name: "empty input", input: []int{}, predicate: func(i int) bool { return i <= 2 }, expected: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CountBy(tt.input, tt.predicate)

			is.Equal(tt.expected, result)
		})
	}
}

func TestCountByErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         []int
		predicate     func(int) (bool, error)
		want          int
		wantErr       string
		wantCallCount int
	}{
		{
			name:  "count elements less than 2",
			input: []int{1, 2, 1},
			predicate: func(i int) (bool, error) {
				return i < 2, nil
			},
			want:          2,
			wantErr:       "",
			wantCallCount: 3,
		},
		{
			name:  "count elements greater than 2",
			input: []int{1, 2, 1},
			predicate: func(i int) (bool, error) {
				return i > 2, nil
			},
			want:          0,
			wantErr:       "",
			wantCallCount: 3,
		},
		{
			name:  "empty slice",
			input: []int{},
			predicate: func(i int) (bool, error) {
				return i <= 2, nil
			},
			want:          0,
			wantErr:       "",
			wantCallCount: 0,
		},
		{
			name:  "error on third element",
			input: []int{1, 2, 3, 4, 5},
			predicate: func(i int) (bool, error) {
				if i == 3 {
					return false, fmt.Errorf("error at %d", i)
				}
				return i < 3, nil
			},
			want:          0,
			wantErr:       "error at 3",
			wantCallCount: 3, // stops early at error
		},
		{
			name:  "error on first element",
			input: []int{1, 2, 3},
			predicate: func(i int) (bool, error) {
				return false, errors.New("first element error")
			},
			want:          0,
			wantErr:       "first element error",
			wantCallCount: 1,
		},
		{
			name:  "all match",
			input: []int{1, 2, 3},
			predicate: func(i int) (bool, error) {
				return i > 0, nil
			},
			want:          3,
			wantErr:       "",
			wantCallCount: 3,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			callCount := 0
			wrappedPredicate := func(i int) (bool, error) {
				callCount++
				return tt.predicate(i)
			}

			got, err := CountByErr(tt.input, wrappedPredicate)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Equal(tt.want, got)
				if tt.wantCallCount > 0 {
					is.Equal(tt.wantCallCount, callCount, "should stop early on error")
				}
			} else {
				is.NoError(err)
				is.Equal(tt.want, got)
				is.Equal(tt.wantCallCount, callCount)
			}
		})
	}
}

func TestCountValues(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Empty(CountValues([]int{}))
		is.Equal(map[int]int{1: 1, 2: 1}, CountValues([]int{1, 2}))
		is.Equal(map[int]int{1: 1, 2: 2}, CountValues([]int{1, 2, 2}))
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Equal(map[string]int{"": 1, "foo": 1, "bar": 1}, CountValues([]string{"foo", "bar", ""}))
		is.Equal(map[string]int{"foo": 1, "bar": 2}, CountValues([]string{"foo", "bar", "bar"}))
	})
}

func TestCountValuesBy(t *testing.T) {
	t.Parallel()

	oddEven := func(v int) bool {
		return v%2 == 0
	}
	length := func(v string) int {
		return len(v)
	}

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Empty(CountValuesBy([]int{}, oddEven))
		is.Equal(map[bool]int{true: 1, false: 1}, CountValuesBy([]int{1, 2}, oddEven))
		is.Equal(map[bool]int{true: 2, false: 1}, CountValuesBy([]int{1, 2, 2}, oddEven))
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Equal(map[int]int{0: 1, 3: 2}, CountValuesBy([]string{"foo", "bar", ""}, length))
		is.Equal(map[int]int{3: 3}, CountValuesBy([]string{"foo", "bar", "bar"}, length))
	})
}

func TestSubset(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 2, 3, 4}

	tests := []struct {
		name     string
		offset   int
		length   uint
		expected []int
	}{
		{name: "offset 0 length 0", offset: 0, length: 0, expected: nil},
		{name: "offset 10 length 2", offset: 10, length: 2, expected: nil},
		{name: "offset -10 length 2", offset: -10, length: 2, expected: []int{0, 1}},
		{name: "offset 0 length 10", offset: 0, length: 10, expected: []int{0, 1, 2, 3, 4}},
		{name: "offset 0 length 2", offset: 0, length: 2, expected: []int{0, 1}},
		{name: "offset 2 length 2", offset: 2, length: 2, expected: []int{2, 3}},
		{name: "offset 2 length 5", offset: 2, length: 5, expected: []int{2, 3, 4}},
		{name: "offset 2 length 3", offset: 2, length: 3, expected: []int{2, 3, 4}},
		{name: "offset 2 length 4", offset: 2, length: 4, expected: []int{2, 3, 4}},
		{name: "offset -2 length 4", offset: -2, length: 4, expected: []int{3, 4}},
		{name: "offset -4 length 1", offset: -4, length: 1, expected: []int{1}},
		{name: "offset -4 length MaxUint", offset: -4, length: math.MaxUint, expected: []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Subset(in, tt.offset, tt.length)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Subset(allStrings, 0, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 2, 3, 4}

	tests := []struct {
		name     string
		start    int
		end      int
		expected []int
	}{
		{name: "0,0", start: 0, end: 0, expected: nil},
		{name: "0,1", start: 0, end: 1, expected: []int{0}},
		{name: "0,5", start: 0, end: 5, expected: []int{0, 1, 2, 3, 4}},
		{name: "0,6", start: 0, end: 6, expected: []int{0, 1, 2, 3, 4}},
		{name: "1,1", start: 1, end: 1, expected: nil},
		{name: "1,5", start: 1, end: 5, expected: []int{1, 2, 3, 4}},
		{name: "1,6", start: 1, end: 6, expected: []int{1, 2, 3, 4}},
		{name: "4,5", start: 4, end: 5, expected: []int{4}},
		{name: "5,5", start: 5, end: 5, expected: nil},
		{name: "6,5", start: 6, end: 5, expected: nil},
		{name: "6,6", start: 6, end: 6, expected: nil},
		{name: "1,0", start: 1, end: 0, expected: nil},
		{name: "5,0", start: 5, end: 0, expected: nil},
		{name: "6,4", start: 6, end: 4, expected: nil},
		{name: "6,7", start: 6, end: 7, expected: nil},
		{name: "-10,1", start: -10, end: 1, expected: []int{0}},
		{name: "-1,3", start: -1, end: 3, expected: []int{0, 1, 2}},
		{name: "-10,7", start: -10, end: 7, expected: []int{0, 1, 2, 3, 4}},
		{name: "-10,-1", start: -10, end: -1, expected: nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Slice(in, tt.start, tt.end)

			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Slice(allStrings, 0, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplace(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	tests := []struct {
		name     string
		old      int
		new      int
		n        int
		expected []int
	}{
		{name: "replace 0 with 42, n=2", old: 0, new: 42, n: 2, expected: []int{42, 1, 42, 1, 2, 3, 0}},
		{name: "replace 0 with 42, n=1", old: 0, new: 42, n: 1, expected: []int{42, 1, 0, 1, 2, 3, 0}},
		{name: "replace 0 with 42, n=0", old: 0, new: 42, n: 0, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "replace 0 with 42, n=-1 (first)", old: 0, new: 42, n: -1, expected: []int{42, 1, 42, 1, 2, 3, 42}},
		{name: "replace 0 with 42, n=-1 (second)", old: 0, new: 42, n: -1, expected: []int{42, 1, 42, 1, 2, 3, 42}},
		{name: "replace -1 with 42, n=2", old: -1, new: 42, n: 2, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "replace -1 with 42, n=1", old: -1, new: 42, n: 1, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "replace -1 with 42, n=0", old: -1, new: 42, n: 0, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "replace -1 with 42, n=-1 (first)", old: -1, new: 42, n: -1, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "replace -1 with 42, n=-1 (second)", old: -1, new: 42, n: -1, expected: []int{0, 1, 0, 1, 2, 3, 0}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Replace(in, tt.old, tt.new, tt.n)

			is.Equal(tt.expected, result)
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Replace(allStrings, "0", "2", 1)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	tests := []struct {
		name     string
		old      int
		new      int
		expected []int
	}{
		{name: "replace present value", old: 0, new: 42, expected: []int{42, 1, 42, 1, 2, 3, 42}},
		{name: "replace absent value", old: -1, new: 42, expected: []int{0, 1, 0, 1, 2, 3, 0}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ReplaceAll(in, tt.old, tt.new)

			is.Equal(tt.expected, result)
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := ReplaceAll(allStrings, "0", "2")
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestClone(t *testing.T) {
	t.Parallel()

	t.Run("int slice - mutating original does not affect clone", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		original1 := []int{1, 2, 3, 4, 5}
		result1 := Clone(original1)
		is.Equal([]int{1, 2, 3, 4, 5}, result1)

		// Verify it's a different slice by checking that modifying one doesn't affect the other
		original1[0] = 99
		is.Equal([]int{99, 2, 3, 4, 5}, original1)
		is.Equal([]int{1, 2, 3, 4, 5}, result1)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		original2 := []string{"a", "b", "c"}
		result2 := Clone(original2)
		assert.Equal(t, []string{"a", "b", "c"}, result2)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		original3 := []int{}
		result3 := Clone(original3)
		is.Equal([]int{}, result3)
		is.Empty(result3)
	})

	t.Run("nil slice", func(t *testing.T) {
		t.Parallel()
		var original4 []int
		result4 := Clone(original4)
		assert.Nil(t, result4)
	})

	t.Run("int slice - mutating clone does not affect original", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		original5 := []int{1, 2, 3}
		result5 := Clone(original5)
		result5[0] = 99
		is.Equal([]int{1, 2, 3}, original5) // Original unchanged
		is.Equal([]int{99, 2, 3}, result5)  // Clone changed
	})

	t.Run("named type - mutating clone does not affect original", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		original6 := myStrings{"", "foo", "bar"}
		result6 := Clone(original6)
		result6[0] = "baz"
		is.Equal(myStrings{"", "foo", "bar"}, original6)  // Original unchanged
		is.Equal(myStrings{"baz", "foo", "bar"}, result6) // Clone changed
	})
}

func TestCompact(t *testing.T) {
	t.Parallel()

	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		r1 := Compact([]int{2, 0, 4, 0})
		assert.Equal(t, []int{2, 4}, r1)
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		r2 := Compact([]string{"", "foo", "", "bar", ""})
		assert.Equal(t, []string{"foo", "bar"}, r2)
	})

	t.Run("bool slice", func(t *testing.T) {
		t.Parallel()
		r3 := Compact([]bool{true, false, true, false})
		assert.Equal(t, []bool{true, true}, r3)
	})

	t.Run("slice of structs", func(t *testing.T) {
		t.Parallel()
		type foo struct {
			bar int
			baz string
		}

		// slice of structs
		// If all fields of an element are zero values, Compact removes it.

		r4 := Compact([]foo{
			{bar: 1, baz: "a"}, // all fields are non-zero values
			{bar: 0, baz: ""},  // all fields are zero values
			{bar: 2, baz: ""},  // bar is non-zero
		})

		assert.Equal(t, []foo{{bar: 1, baz: "a"}, {bar: 2, baz: ""}}, r4)
	})

	t.Run("slice of pointers to structs", func(t *testing.T) {
		t.Parallel()
		type foo struct {
			bar int
			baz string
		}

		// slice of pointers to structs
		// If an element is nil, Compact removes it.

		e1, e2, e3 := foo{bar: 1, baz: "a"}, foo{bar: 0, baz: ""}, foo{bar: 2, baz: ""}
		// NOTE: e2 is a zero value of foo, but its pointer &e2 is not a zero value of *foo.
		r5 := Compact([]*foo{&e1, &e2, nil, &e3})

		assert.Equal(t, []*foo{&e1, &e2, &e3}, r5)
	})

	is := assert.New(t)
	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Compact(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestIsSorted(t *testing.T) {
	t.Parallel()

	t.Run("sorted int slice", func(t *testing.T) {
		t.Parallel()
		assert.True(t, IsSorted([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	})

	t.Run("sorted string slice", func(t *testing.T) {
		t.Parallel()
		assert.True(t, IsSorted([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}))
	})

	t.Run("unsorted int slice", func(t *testing.T) {
		t.Parallel()
		assert.False(t, IsSorted([]int{0, 1, 4, 3, 2, 5, 6, 7, 8, 9, 10}))
	})

	t.Run("unsorted string slice", func(t *testing.T) {
		t.Parallel()
		assert.False(t, IsSorted([]string{"a", "b", "d", "c", "e", "f", "g", "h", "i", "j"}))
	})
}

func TestIsSortedBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		iteratee func(s string) int
		expected bool
	}{
		{name: "sorted by length", input: []string{"a", "bb", "ccc"}, iteratee: func(s string) int { return len(s) }, expected: true},
		{name: "unsorted by length", input: []string{"aa", "b", "ccc"}, iteratee: func(s string) int { return len(s) }, expected: false},
		{name: "sorted by numeric value", input: []string{"1", "2", "3", "11"}, iteratee: func(s string) int {
			ret, _ := strconv.Atoi(s)
			return ret
		}, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, IsSortedBy(tt.input, tt.iteratee))
		})
	}
}

func TestSplice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sample := []string{"a", "b", "c", "d", "e", "f", "g"}

	tests := []struct {
		name          string
		input         []string
		pos           int
		values        []string
		checkOriginal bool
		expected      []string
	}{
		{name: "normal case", input: sample, pos: 1, values: []string{"1", "2"}, checkOriginal: true, expected: []string{"a", "1", "2", "b", "c", "d", "e", "f", "g"}},
		{name: "positive overflow", input: sample, pos: 42, values: []string{"1", "2"}, checkOriginal: true, expected: []string{"a", "b", "c", "d", "e", "f", "g", "1", "2"}},
		{name: "negative overflow", input: sample, pos: -42, values: []string{"1", "2"}, checkOriginal: true, expected: []string{"1", "2", "a", "b", "c", "d", "e", "f", "g"}},
		{name: "backward -2", input: sample, pos: -2, values: []string{"1", "2"}, checkOriginal: true, expected: []string{"a", "b", "c", "d", "e", "1", "2", "f", "g"}},
		{name: "backward -7", input: sample, pos: -7, values: []string{"1", "2"}, checkOriginal: true, expected: []string{"1", "2", "a", "b", "c", "d", "e", "f", "g"}},
		{name: "empty input pos 0", input: []string{}, pos: 0, values: []string{"1", "2"}, expected: []string{"1", "2"}},
		{name: "empty input pos 1", input: []string{}, pos: 1, values: []string{"1", "2"}, expected: []string{"1", "2"}},
		{name: "empty input pos -1", input: []string{}, pos: -1, values: []string{"1", "2"}, expected: []string{"1", "2"}},
		{name: "single item pos 0", input: []string{"0"}, pos: 0, values: []string{"1", "2"}, expected: []string{"1", "2", "0"}},
		{name: "single item pos 1", input: []string{"0"}, pos: 1, values: []string{"1", "2"}, expected: []string{"0", "1", "2"}},
		{name: "single item pos -1", input: []string{"0"}, pos: -1, values: []string{"1", "2"}, expected: []string{"1", "2", "0"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Splice(tt.input, tt.pos, tt.values...)

			if tt.checkOriginal {
				is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, sample)
			}
			is.Equal(tt.expected, result)
		})
	}

	t.Run("no side effect on returned slice mutation", func(t *testing.T) {
		t.Parallel()

		// check there is no side effect
		results := Splice(sample, 1)
		results[0] = "b"
		is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, sample)
	})

	// type preserved
	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Splice(allStrings, 1, "1", "2")
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestCut_success(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     []string
		values    []string
		wantLeft  []string
		wantRight []string
	}{
		{name: "case 1", input: []string{"a", "b", "c", "d", "e", "f", "g"}, values: []string{"a", "b"}, wantLeft: []string{}, wantRight: []string{"c", "d", "e", "f", "g"}},
		{name: "case 2", input: []string{"a", "b", "c", "d", "e", "f", "g"}, values: []string{"f", "g"}, wantLeft: []string{"a", "b", "c", "d", "e"}, wantRight: []string{}},
		{name: "case 3", input: []string{"g"}, values: []string{"g"}, wantLeft: []string{}, wantRight: []string{}},
		{name: "case 4", input: []string{"a", "b", "c", "d", "e", "f", "g"}, values: []string{"b", "c"}, wantLeft: []string{"a"}, wantRight: []string{"d", "e", "f", "g"}},
		{name: "case 5", input: []string{"a", "b", "c", "d", "e", "f", "g"}, values: []string{"e", "f"}, wantLeft: []string{"a", "b", "c", "d"}, wantRight: []string{"g"}},
		{name: "case 6", input: []string{"a", "b"}, values: []string{"b"}, wantLeft: []string{"a"}, wantRight: []string{}},
		{name: "case 7", input: []string{"a", "b"}, values: []string{"a"}, wantLeft: []string{}, wantRight: []string{"b"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualLeft, actualRight, result := Cut(tt.input, tt.values)

			is.True(result)
			is.Equal(tt.wantLeft, actualLeft)
			is.Equal(tt.wantRight, actualRight)
		})
	}
}

func TestCut_fail(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     []string
		values    []string
		wantLeft  []string
		wantRight []string
	}{
		{name: "case 1", input: []string{"a", "b", "c", "d", "e", "f", "g"}, values: []string{"z"}, wantLeft: []string{"a", "b", "c", "d", "e", "f", "g"}, wantRight: []string{}},
		{name: "case 2", input: []string{}, values: []string{"z"}, wantLeft: []string{}, wantRight: []string{}},
		{name: "case 3", input: []string{"a"}, values: []string{"z"}, wantLeft: []string{"a"}, wantRight: []string{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualLeft, actualRight, result := Cut(tt.input, tt.values)

			is.False(result)
			is.Equal(tt.wantLeft, actualLeft)
			is.Equal(tt.wantRight, actualRight)
		})
	}
}

type TestCutStruct struct {
	id   int
	data string
}

func TestCutPrefix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// case 1
	tests := []struct {
		name       string
		input      []TestCutStruct
		values     []TestCutStruct
		wantResult bool
		wantAfter  []TestCutStruct
	}{
		{
			name:       "case 1",
			input:      []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
			values:     []TestCutStruct{{id: 1, data: "a"}},
			wantResult: true,
			wantAfter:  []TestCutStruct{{id: 2, data: "a"}, {id: 2, data: "b"}},
		},
		{
			name:       "case 2",
			input:      []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
			values:     []TestCutStruct{},
			wantResult: true,
			wantAfter:  []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
		},
		{
			name:       "case 3",
			input:      []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
			values:     []TestCutStruct{{id: 2, data: "b"}},
			wantResult: false,
			wantAfter:  []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
		},
		{
			name:       "case 4",
			input:      []TestCutStruct{},
			values:     []TestCutStruct{{id: 2, data: "b"}},
			wantResult: false,
			wantAfter:  []TestCutStruct{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualAfter, result := CutPrefix(tt.input, tt.values)

			is.Equal(tt.wantResult, result)
			is.Equal(tt.wantAfter, actualAfter)
		})
	}

	t.Run("case 5 - string slice", func(t *testing.T) {
		t.Parallel()

		actualAfterS, result := CutPrefix([]string{"a", "a", "b"}, []string{})
		is.True(result)
		is.Equal([]string{"a", "a", "b"}, actualAfterS)
	})
}

func TestCutSuffix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// case 1
	tests := []struct {
		name       string
		input      []TestCutStruct
		values     []TestCutStruct
		wantResult bool
		wantBefore []TestCutStruct
	}{
		{
			name:       "case 1",
			input:      []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
			values:     []TestCutStruct{{id: 3, data: "b"}},
			wantResult: false,
			wantBefore: []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
		},
		{
			name:       "case 2",
			input:      []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
			values:     []TestCutStruct{{id: 2, data: "b"}},
			wantResult: true,
			wantBefore: []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}},
		},
		{
			name:       "case 3",
			input:      []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
			values:     []TestCutStruct{},
			wantResult: true,
			wantBefore: []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
		},
		{
			name:       "case 4",
			input:      []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
			values:     []TestCutStruct{{id: 2, data: "a"}},
			wantResult: false,
			wantBefore: []TestCutStruct{{id: 1, data: "a"}, {id: 2, data: "a"}, {id: 2, data: "b"}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualBefore, result := CutSuffix(tt.input, tt.values)

			is.Equal(tt.wantResult, result)
			is.Equal(tt.wantBefore, actualBefore)
		})
	}

	t.Run("case 5 - string slice", func(t *testing.T) {
		t.Parallel()

		actualAfterS, result := CutSuffix([]string{"a", "a", "b"}, []string{})
		is.True(result)
		is.Equal([]string{"a", "a", "b"}, actualAfterS)
	})
}

// TestTrim_smallScan exercises the small-scan path (all cutsets here are
// <= trimSmallCutset). See TestTrim_large for the map-based path.
func TestTrim_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		cutset   []string
		expected []string
	}{
		{name: "trim prefix and suffix", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b"}, expected: []string{"c", "d", "e", "f", "g"}},
		{name: "trim only suffix present in cutset", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e"}},
		{name: "trim everything", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: []string{}},
		{name: "cutset larger than input", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: []string{}},
		{name: "empty cutset", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, Trim(tt.input, tt.cutset))
		})
	}
}

// Trim dispatches on len(cutset) <= trimSmallCutset (8): a cutset of 9 unique
// elements forces the trimLarge path, which the table above never exercises
// (its cutsets are all <= 8).
func TestTrim_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	cutset := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	is.Greater(len(cutset), trimSmallCutset, "sanity check: cutset must exceed trimSmallCutset")

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{name: "cutset repeated around distinct middle", input: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "X", "Y", "a", "b", "c", "d", "e", "f", "g", "h", "i"}, expected: []string{"X", "Y"}},
		{name: "input equals cutset", input: cutset, expected: []string{}},
		{name: "input disjoint from cutset", input: []string{"X", "Y"}, expected: []string{"X", "Y"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, Trim(tt.input, cutset))
		})
	}
}

// TestTrimLeft_smallScan exercises the small-scan path (all cutsets here are
// <= trimSmallCutset). See TestTrimLeft_large for the map-based path.
func TestTrimLeft_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		cutset   []string
		expected []string
	}{
		{name: "trim repeated prefix", input: []string{"a", "a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b"}, expected: []string{"c", "d", "e", "f", "g"}},
		{name: "trim prefix with cutset order reversed", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"b", "a"}, expected: []string{"c", "d", "e", "f", "g"}},
		{name: "cutset not at prefix", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "trim everything", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: []string{}},
		{name: "cutset larger than input", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: []string{}},
		{name: "empty cutset", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, TrimLeft(tt.input, tt.cutset))
		})
	}
}

// TrimLeft dispatches on len(cutset) <= trimSmallCutset (8): a cutset of 9
// unique elements forces the trimLeftLarge path, which the table above never
// exercises (its cutsets are all <= 8).
func TestTrimLeft_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	cutset := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	is.Greater(len(cutset), trimSmallCutset, "sanity check: cutset must exceed trimSmallCutset")

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{name: "prefix matches cutset", input: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "X", "Y"}, expected: []string{"X", "Y"}},
		{name: "input equals cutset", input: cutset, expected: []string{}},
		{name: "input disjoint from cutset", input: []string{"X", "Y"}, expected: []string{"X", "Y"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, TrimLeft(tt.input, cutset))
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		prefix   []string
		expected []string
	}{
		{name: "trim matching prefix", input: []string{"a", "b", "a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"a", "b"}, expected: []string{"c", "d", "e", "f", "g"}},
		{name: "prefix order mismatch", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"b", "a"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "prefix not at start", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "prefix equals whole input", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: []string{}},
		{name: "prefix longer than input", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "empty prefix", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, TrimPrefix(tt.input, tt.prefix))
		})
	}
}

// TestTrimRight_smallScan exercises the small-scan path (all cutsets here are
// <= trimSmallCutset). See TestTrimRight_large for the map-based path.
func TestTrimRight_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		cutset   []string
		expected []string
	}{
		{name: "cutset not at suffix", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "trim repeated suffix", input: []string{"a", "b", "c", "d", "e", "f", "g", "g"}, cutset: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e"}},
		{name: "trim everything", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: []string{}},
		{name: "cutset larger than input", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: []string{}},
		{name: "empty cutset", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, TrimRight(tt.input, tt.cutset))
		})
	}
}

// TrimRight dispatches on len(cutset) <= trimSmallCutset (8): a cutset of 9
// unique elements forces the trimRightLarge path, which the table above
// never exercises (its cutsets are all <= 8).
func TestTrimRight_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	cutset := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	is.Greater(len(cutset), trimSmallCutset, "sanity check: cutset must exceed trimSmallCutset")

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{name: "suffix matches cutset", input: []string{"X", "Y", "a", "b", "c", "d", "e", "f", "g", "h", "i"}, expected: []string{"X", "Y"}},
		{name: "input equals cutset", input: cutset, expected: []string{}},
		{name: "input disjoint from cutset", input: []string{"X", "Y"}, expected: []string{"X", "Y"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, TrimRight(tt.input, cutset))
		})
	}
}

func TestTrimSuffix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		suffix   []string
		expected []string
	}{
		{name: "suffix not at end", input: []string{"a", "b", "c", "d", "e", "f", "g"}, suffix: []string{"a", "b"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "trim matching suffix", input: []string{"a", "b", "c", "d", "e", "f", "g", "f", "g"}, suffix: []string{"f", "g"}, expected: []string{"a", "b", "c", "d", "e"}},
		{name: "suffix order mismatch", input: []string{"a", "b", "c", "d", "e", "f", "g", "f", "g"}, suffix: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e", "f", "g", "f", "g"}},
		{name: "suffix equals whole input", input: []string{"a", "b", "c", "d", "e", "f", "g"}, suffix: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: []string{}},
		{name: "suffix longer than input", input: []string{"a", "b", "c", "d", "e", "f", "g"}, suffix: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "empty suffix", input: []string{"a", "b", "c", "d", "e", "f", "g"}, suffix: []string{}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			is.Equal(tt.expected, TrimSuffix(tt.input, tt.suffix))
		})
	}
}
