//go:build go1.23

package it

import (
	"fmt"
	"iter"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "empty", input: []int{}, expected: 0},
		{name: "four elements", input: []int{1, 2, 3, 4}, expected: 4},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Length(values(tt.input...)))
		})
	}
}

func TestDrain(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var done bool
	list := iter.Seq[int](func(yield func(int) bool) {
		_ = yield(1) && yield(2) && yield(3)

		done = true
	})

	Drain(list)

	is.True(done)
}

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int evens", func(t *testing.T) {
		t.Parallel()
		r1 := Filter(values(1, 2, 3, 4), func(x int) bool {
			return x%2 == 0
		})
		is.Equal([]int{2, 4}, slices.Collect(r1))
	})

	t.Run("non-empty strings", func(t *testing.T) {
		t.Parallel()
		r2 := Filter(values("", "foo", "", "bar", ""), func(x string) bool {
			return len(x) > 0
		})
		is.Equal([]string{"foo", "bar"}, slices.Collect(r2))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Filter(allStrings, func(x string) bool {
			return len(x) > 0
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestFilterI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int evens", func(t *testing.T) {
		t.Parallel()
		r1 := FilterI(values(1, 2, 3, 4), func(x, _ int) bool {
			return x%2 == 0
		})
		is.Equal([]int{2, 4}, slices.Collect(r1))
	})

	t.Run("non-empty strings", func(t *testing.T) {
		t.Parallel()
		r2 := FilterI(values("", "foo", "", "bar", ""), func(x string, _ int) bool {
			return len(x) > 0
		})
		is.Equal([]string{"foo", "bar"}, slices.Collect(r2))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := FilterI(allStrings, func(x string, _ int) bool {
			return len(x) > 0
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int to constant string", func(t *testing.T) {
		t.Parallel()
		result1 := Map(values(1, 2, 3, 4), func(x int) string {
			return "Hello"
		})
		is.Equal([]string{"Hello", "Hello", "Hello", "Hello"}, slices.Collect(result1))
	})

	t.Run("int64 to formatted string", func(t *testing.T) {
		t.Parallel()
		result2 := Map(values[int64](1, 2, 3, 4), func(x int64) string {
			return strconv.FormatInt(x, 10)
		})
		is.Equal([]string{"1", "2", "3", "4"}, slices.Collect(result2))
	})
}

func TestMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int to constant string", func(t *testing.T) {
		t.Parallel()
		result1 := MapI(values(1, 2, 3, 4), func(x, _ int) string {
			return "Hello"
		})
		is.Equal([]string{"Hello", "Hello", "Hello", "Hello"}, slices.Collect(result1))
	})

	t.Run("int64 to formatted string", func(t *testing.T) {
		t.Parallel()
		result2 := MapI(values[int64](1, 2, 3, 4), func(x int64, _ int) string {
			return strconv.FormatInt(x, 10)
		})
		is.Equal([]string{"1", "2", "3", "4"}, slices.Collect(result2))
	})
}

func TestUniqMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type User struct {
		Name string
		age  int
	}

	users := values(User{Name: "Alice", age: 20}, User{Name: "Alex", age: 21}, User{Name: "Alex", age: 22})
	result := UniqMap(users, func(item User) string {
		return item.Name
	})

	is.Equal([]string{"Alice", "Alex"}, slices.Collect(result))
}

func TestUniqMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type User struct {
		Name string
		age  int
	}

	users := values(User{Name: "Alice", age: 20}, User{Name: "Alex", age: 21}, User{Name: "Alex", age: 22})
	result := UniqMapI(users, func(item User, _ int) string {
		return item.Name
	})

	is.Equal([]string{"Alice", "Alex"}, slices.Collect(result))
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int64 evens formatted", func(t *testing.T) {
		t.Parallel()
		r1 := FilterMap(values[int64](1, 2, 3, 4), func(x int64) (string, bool) {
			if x%2 == 0 {
				return strconv.FormatInt(x, 10), true
			}
			return "", false
		})
		is.Equal([]string{"2", "4"}, slices.Collect(r1))
	})

	t.Run("string suffix match", func(t *testing.T) {
		t.Parallel()
		r2 := FilterMap(values("cpu", "gpu", "mouse", "keyboard"), func(x string) (string, bool) {
			if strings.HasSuffix(x, "pu") {
				return "xpu", true
			}
			return "", false
		})
		is.Equal([]string{"xpu", "xpu"}, slices.Collect(r2))
	})
}

func TestFilterMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int64 evens formatted", func(t *testing.T) {
		t.Parallel()
		r1 := FilterMapI(values[int64](1, 2, 3, 4), func(x int64, _ int) (string, bool) {
			if x%2 == 0 {
				return strconv.FormatInt(x, 10), true
			}
			return "", false
		})
		is.Equal([]string{"2", "4"}, slices.Collect(r1))
	})

	t.Run("string suffix match", func(t *testing.T) {
		t.Parallel()
		r2 := FilterMapI(values("cpu", "gpu", "mouse", "keyboard"), func(x string, _ int) (string, bool) {
			if strings.HasSuffix(x, "pu") {
				return "xpu", true
			}
			return "", false
		})
		is.Equal([]string{"xpu", "xpu"}, slices.Collect(r2))
	})
}

func TestFlatMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int to constant string", func(t *testing.T) {
		t.Parallel()
		result1 := FlatMap(values(0, 1, 2, 3, 4), func(x int) iter.Seq[string] {
			return values("Hello")
		})
		is.Equal([]string{"Hello", "Hello", "Hello", "Hello", "Hello"}, slices.Collect(result1))
	})

	t.Run("int64 repeated by value", func(t *testing.T) {
		t.Parallel()
		result2 := FlatMap(values[int64](0, 1, 2, 3, 4), func(x int64) iter.Seq[string] {
			return func(yield func(string) bool) {
				for range x {
					if !yield(strconv.FormatInt(x, 10)) {
						return
					}
				}
			}
		})
		is.Equal([]string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"}, slices.Collect(result2))
	})
}

func TestFlatMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int to constant string", func(t *testing.T) {
		t.Parallel()
		result1 := FlatMapI(values(0, 1, 2, 3, 4), func(x, _ int) iter.Seq[string] {
			return values("Hello")
		})
		is.Equal([]string{"Hello", "Hello", "Hello", "Hello", "Hello"}, slices.Collect(result1))
	})

	t.Run("int64 repeated by value", func(t *testing.T) {
		t.Parallel()
		result2 := FlatMapI(values[int64](0, 1, 2, 3, 4), func(x int64, _ int) iter.Seq[string] {
			return func(yield func(string) bool) {
				for range x {
					if !yield(strconv.FormatInt(x, 10)) {
						return
					}
				}
			}
		})
		is.Equal([]string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"}, slices.Collect(result2))
	})
}

func TestTimes(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	is.Equal([]string{"0", "1", "2"}, slices.Collect(result1))
}

func TestReduce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		seed     int
		expected int
	}{
		{name: "seed zero", seed: 0, expected: 10},
		{name: "seed ten", seed: 10, expected: 20},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Reduce(values(1, 2, 3, 4), func(agg, item int) int {
				return agg + item
			}, tt.seed)
			is.Equal(tt.expected, result)
		})
	}
}

func TestReduceI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		seed     int
		expected int
	}{
		{name: "seed zero", seed: 0, expected: 10},
		{name: "seed ten", seed: 10, expected: 20},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ReduceI(values(1, 2, 3, 4), func(agg, item, _ int) int {
				return agg + item
			}, tt.seed)
			is.Equal(tt.expected, result)
		})
	}
}

func TestReduceLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("slice concatenation", func(t *testing.T) {
		t.Parallel()
		result1 := ReduceLast(values([]int{0, 1}, []int{2, 3}, []int{4, 5}), func(agg, item []int) []int {
			return append(agg, item...)
		}, []int{})
		is.Equal([]int{4, 5, 2, 3, 0, 1}, result1)
	})

	t.Run("int sum", func(t *testing.T) {
		t.Parallel()
		result2 := ReduceLast(values(1, 2, 3, 4), func(agg, item int) int {
			return agg + item
		}, 10)
		is.Equal(20, result2)
	})
}

func TestReduceLastI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("slice concatenation", func(t *testing.T) {
		t.Parallel()
		result1 := ReduceLastI(values([]int{0, 1}, []int{2, 3}, []int{4, 5}), func(agg, item []int, _ int) []int {
			return append(agg, item...)
		}, []int{})
		is.Equal([]int{4, 5, 2, 3, 0, 1}, result1)
	})

	t.Run("int sum", func(t *testing.T) {
		t.Parallel()
		result2 := ReduceLastI(values(1, 2, 3, 4), func(agg, item, _ int) int {
			return agg + item
		}, 10)
		is.Equal(20, result2)
	})
}

func TestForEachI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	ForEachI(values("a", "b", "c"), func(item string, i int) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.Equal([]string{"a", "b", "c"}, callParams1)
	is.Equal([]int{0, 1, 2}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestForEachWhileI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	var callParams1 []string
	var callParams2 []int

	ForEachWhileI(values("a", "b", "c"), func(item string, i int) bool {
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

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int dedup", func(t *testing.T) {
		t.Parallel()
		result1 := Uniq(values(1, 2, 2, 1))
		is.Equal([]int{1, 2}, slices.Collect(result1))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Uniq(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int mod dedup", func(t *testing.T) {
		t.Parallel()
		result1 := UniqBy(values(0, 1, 2, 3, 4, 5), func(i int) int {
			return i % 3
		})
		is.Equal([]int{0, 1, 2}, slices.Collect(result1))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := UniqBy(allStrings, func(i string) string {
			return i
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupBy(values(0, 1, 2, 3, 4, 5), func(i int) int {
		return i % 3
	})
	is.Equal(map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	}, result1)
}

func TestGroupByMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int keys", func(t *testing.T) {
		t.Parallel()
		result1 := GroupByMap(values(0, 1, 2, 3, 4, 5), func(i int) (int, string) {
			return i % 3, strconv.Itoa(i)
		})
		is.Equal(map[int][]string{
			0: {"0", "3"},
			1: {"1", "4"},
			2: {"2", "5"},
		}, result1)
	})

	t.Run("named int type keys", func(t *testing.T) {
		t.Parallel()
		type myInt int
		result2 := GroupByMap(values[myInt](1, 0, 2, 3, 4, 5), func(i myInt) (int, string) {
			return int(i % 3), strconv.Itoa(int(i))
		})
		is.Equal(map[int][]string{
			0: {"0", "3"},
			1: {"1", "4"},
			2: {"2", "5"},
		}, result2)
	})

	t.Run("struct keys", func(t *testing.T) {
		t.Parallel()
		type product struct {
			ID         int64
			CategoryID int64
		}
		products := values(
			product{ID: 1, CategoryID: 1},
			product{ID: 2, CategoryID: 1},
			product{ID: 3, CategoryID: 2},
			product{ID: 4, CategoryID: 3},
			product{ID: 5, CategoryID: 3},
		)
		result3 := GroupByMap(products, func(item product) (int64, string) {
			return item.CategoryID, "Product " + strconv.FormatInt(item.ID, 10)
		})
		is.Equal(map[int64][]string{
			1: {"Product 1", "Product 2"},
			2: {"Product 3"},
			3: {"Product 4", "Product 5"},
		}, result3)
	})
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
		{name: "uneven split", input: []int{0, 1, 2, 3, 4, 5, 6}, size: 2, expected: [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}},
		{name: "empty", input: []int{}, size: 2, expected: nil},
		{name: "single element", input: []int{0}, size: 2, expected: [][]int{{0}}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Chunk(values(tt.input...), tt.size)))
		})
	}

	t.Run("panics on non-positive size", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Chunk: size must be greater than 0", func() {
			Chunk(values(0), 0)
		})
	})
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
		{name: "five elements size three", input: []int{1, 2, 3, 4, 5}, size: 3, expected: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}},
		{name: "six elements size three", input: []int{1, 2, 3, 4, 5, 6}, size: 3, expected: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}},
		{name: "shorter than size", input: []int{1, 2}, size: 3, expected: nil},
		{name: "exactly size", input: []int{1, 2, 3}, size: 3, expected: [][]int{{1, 2, 3}}},
		{name: "size one", input: []int{1, 2, 3, 4}, size: 1, expected: [][]int{{1}, {2}, {3}, {4}}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Window(values(tt.input...), tt.size)))
		})
	}

	t.Run("panics on zero size", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Window: size must be greater than 0", func() {
			Window(values(1, 2, 3), 0)
		})
	})

	t.Run("panics on negative size", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Window: size must be greater than 0", func() {
			Window(values(1, 2, 3), -1)
		})
	})
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
		{name: "step less than size (overlap)", input: []int{1, 2, 3, 4, 5, 6}, size: 3, step: 1, expected: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}},
		{name: "step equals size (no overlap)", input: []int{1, 2, 3, 4, 5, 6}, size: 3, step: 3, expected: [][]int{{1, 2, 3}, {4, 5, 6}}},
		{name: "step greater than size (gaps)", input: []int{1, 2, 3, 4, 5, 6, 7, 8}, size: 2, step: 3, expected: [][]int{{1, 2}, {4, 5}, {7, 8}}},
		{name: "size one step one", input: []int{1, 2, 3, 4}, size: 1, step: 1, expected: [][]int{{1}, {2}, {3}, {4}}},
		{name: "collection shorter than size", input: []int{1, 2}, size: 3, step: 1, expected: nil},
		{name: "size equals step two", input: []int{1, 2, 3, 4, 5, 6}, size: 2, step: 2, expected: [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{name: "overlapping windows step less than size", input: []int{1, 2, 3, 4, 5}, size: 3, step: 2, expected: [][]int{{1, 2, 3}, {3, 4, 5}}},
		{name: "step greater than size with gaps", input: []int{1, 2, 3, 4, 5, 6, 7, 8}, size: 2, step: 4, expected: [][]int{{1, 2}, {5, 6}}},
		{name: "empty collection", input: []int{}, size: 2, step: 1, expected: nil},
		{name: "collection exactly equal to size step 1", input: []int{1, 2, 3}, size: 3, step: 1, expected: [][]int{{1, 2, 3}}},
		{name: "collection exactly equal to size step 3", input: []int{1, 2, 3}, size: 3, step: 3, expected: [][]int{{1, 2, 3}}},
		{name: "collection just larger than size", input: []int{1, 2, 3, 4}, size: 3, step: 1, expected: [][]int{{1, 2, 3}, {2, 3, 4}}},
		{name: "size one step two", input: []int{1, 2, 3, 4, 5}, size: 1, step: 2, expected: [][]int{{1}, {3}, {5}}},
		{name: "size one step three", input: []int{1, 2, 3, 4, 5}, size: 1, step: 3, expected: [][]int{{1}, {4}}},
		{name: "very large step (only first window)", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 2, step: 100, expected: [][]int{{1, 2}}},
		{name: "step one large size (maximum overlap)", input: []int{1, 2, 3, 4, 5}, size: 4, step: 1, expected: [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}}},
		{name: "size equals collection length with step greater than size", input: []int{1, 2, 3}, size: 3, step: 5, expected: [][]int{{1, 2, 3}}},
		{name: "step one size two on single element", input: []int{1}, size: 2, step: 1, expected: nil},
		{name: "large collection with small windows", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 2, step: 2, expected: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}},
		{name: "overlapping step two size four (only full windows)", input: []int{1, 2, 3, 4, 5, 6, 7}, size: 4, step: 2, expected: [][]int{{1, 2, 3, 4}, {3, 4, 5, 6}}},
		{name: "size five step three on seven elements (only full windows)", input: []int{1, 2, 3, 4, 5, 6, 7}, size: 5, step: 3, expected: [][]int{{1, 2, 3, 4, 5}}},
		{name: "collection size exactly size+step-1 (two overlapping windows)", input: []int{1, 2, 3, 4, 5}, size: 3, step: 2, expected: [][]int{{1, 2, 3}, {3, 4, 5}}},
		{name: "large size small step large collection (only full windows)", input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, size: 5, step: 3, expected: [][]int{{1, 2, 3, 4, 5}, {4, 5, 6, 7, 8}, {7, 8, 9, 10, 11}}},
		{name: "size equals step equals one", input: []int{1, 2, 3, 4, 5}, size: 1, step: 1, expected: [][]int{{1}, {2}, {3}, {4}, {5}}},
		{name: "zero elements in collection", input: []int{}, size: 1, step: 1, expected: nil},
		{name: "last window exactly size elements", input: []int{1, 2, 3, 4, 5, 6}, size: 3, step: 3, expected: [][]int{{1, 2, 3}, {4, 5, 6}}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Sliding(values(tt.input...), tt.size, tt.step)))
		})
	}

	t.Run("panics on zero size", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Sliding: size must be greater than 0", func() {
			Sliding(values(1, 2, 3), 0, 1)
		})
	})

	t.Run("panics on zero step", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Sliding: step must be greater than 0", func() {
			Sliding(values(1, 2, 3), 2, 0)
		})
	})

	t.Run("panics on negative step", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Sliding: step must be greater than 0", func() {
			Sliding(values(1, 2, 3), 2, -1)
		})
	})

	t.Run("panics on negative size", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Sliding: size must be greater than 0", func() {
			Sliding(values(1, 2, 3), -1, 1)
		})
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		result17 := Sliding(values("a", "b", "c", "d"), 2, 1)
		is.Equal([][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}}, slices.Collect(result17))
	})

	t.Run("structs", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}
		people := values(
			Person{"Alice", 25},
			Person{"Bob", 30},
			Person{"Charlie", 35},
			Person{"Diana", 40},
		)
		result18 := Sliding(people, 2, 1)
		collected := slices.Collect(result18)
		is.Len(collected, 3)
		is.Equal(Person{"Alice", 25}, collected[0][0])
		is.Equal(Person{"Bob", 30}, collected[0][1])
		is.Equal(Person{"Bob", 30}, collected[1][0])
		is.Equal(Person{"Charlie", 35}, collected[1][1])
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		result23 := Sliding(values(1.1, 2.2, 3.3, 4.4), 2, 1)
		is.Equal([][]float64{{1.1, 2.2}, {2.2, 3.3}, {3.3, 4.4}}, slices.Collect(result23))
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		result24 := Sliding(values(true, false, true, false, true), 3, 2)
		is.Equal([][]bool{{true, false, true}, {true, false, true}}, slices.Collect(result24))
	})

	t.Run("early termination", func(t *testing.T) {
		t.Parallel()
		result27 := Sliding(values(1, 2, 3, 4, 5, 6), 2, 1)
		count := 0
		for window := range result27 {
			count++
			if count == 2 {
				break
			}
			_ = window
		}
		is.Equal(2, count)
	})

	t.Run("pointers", func(t *testing.T) {
		t.Parallel()
		x, y, z := 1, 2, 3
		result32 := Sliding(values(&x, &y, &z), 2, 1)
		collected32 := slices.Collect(result32)
		is.Len(collected32, 2)
		is.Equal(&x, collected32[0][0])
		is.Equal(&y, collected32[0][1])
		is.Equal(&y, collected32[1][0])
		is.Equal(&z, collected32[1][1])
	})
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
		{name: "empty", input: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, PartitionBy(values(tt.input...), oddEven))
		})
	}
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int sequences", func(t *testing.T) {
		t.Parallel()
		result1 := Flatten([]iter.Seq[int]{values(0, 1), values(2, 3, 4, 5)})
		is.Equal([]int{0, 1, 2, 3, 4, 5}, slices.Collect(result1))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Flatten([]myStrings{allStrings})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestConcat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int sequences", func(t *testing.T) {
		t.Parallel()
		result1 := Concat(values(0, 1), values(2, 3, 4, 5))
		is.Equal([]int{0, 1, 2, 3, 4, 5}, slices.Collect(result1))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Concat(allStrings, allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestInterleave(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		in   []iter.Seq[int]
		want []int
	}{
		{
			"empty",
			[]iter.Seq[int]{},
			nil,
		},
		{
			"empties",
			[]iter.Seq[int]{values[int](), values[int]()},
			nil,
		},
		{
			"same length",
			[]iter.Seq[int]{values(1, 3, 5), values(2, 4, 6)},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"different length",
			[]iter.Seq[int]{values(1, 3, 5, 6), values(2, 4)},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"many sequences",
			[]iter.Seq[int]{values(1), values(2, 5, 8), values(3, 6), values(4, 7, 9, 10)},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tc := range testCases {
		tc := tc //nolint:modernize
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, slices.Collect(Interleave(tc.in...)))
		})
	}
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("non-empty shuffle preserves elements", func(t *testing.T) {
		t.Parallel()
		result1 := Shuffle(values(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
		slice1 := slices.Collect(result1)
		is.NotEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, slice1)
		is.ElementsMatch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, slice1)
	})

	t.Run("empty input", func(t *testing.T) {
		t.Parallel()
		result2 := Shuffle(values[int]())
		is.Empty(slices.Collect(result2))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Shuffle(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "six elements", input: []int{0, 1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1, 0}},
		{name: "seven elements", input: []int{0, 1, 2, 3, 4, 5, 6}, expected: []int{6, 5, 4, 3, 2, 1, 0}},
		{name: "empty", input: []int{}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Reverse(values(tt.input...))))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Reverse(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestFill(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []foo
		fillWith foo
		expected []foo
	}{
		{name: "two elements", input: []foo{{"a"}, {"a"}}, fillWith: foo{"b"}, expected: []foo{{"b"}, {"b"}}},
		{name: "empty", input: []foo{}, fillWith: foo{"a"}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Fill(values(tt.input...), tt.fillWith)))
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
		{name: "two repeats", count: 2, value: foo{"a"}, expected: []foo{{"a"}, {"a"}}},
		{name: "zero repeats", count: 0, value: foo{"a"}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Repeat(tt.count, tt.value)))
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
		{name: "zero", count: 0, expected: nil},
		{name: "two", count: 2, expected: []int{0, 1}},
		{name: "five", count: 5, expected: []int{0, 1, 4, 9, 16}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(RepeatBy(tt.count, cb)))
		})
	}
}

func TestKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := KeyBy(values("a", "aa", "aaa"), func(str string) int {
		return len(str)
	})

	is.Equal(map[int]string{1: "a", 2: "aa", 3: "aaa"}, result1)
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
		tc := tc //nolint:modernize
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, Associate(slices.Values(tc.in), transform))
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
		tc := tc //nolint:modernize
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, AssociateI(slices.Values(tc.in), transform))
		})
	}
}

func TestSeqToMap(t *testing.T) {
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
		tc := tc //nolint:modernize
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, SeqToMap(slices.Values(tc.in), transform))
		})
	}
}

func TestSeqToMapI(t *testing.T) {
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
		tc := tc //nolint:modernize
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, SeqToMapI(slices.Values(tc.in), transform))
		})
	}
}

func TestFilterSeqToMap(t *testing.T) {
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
		tc := tc //nolint:modernize
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, FilterSeqToMap(slices.Values(tc.in), transform))
		})
	}
}

func TestFilterSeqToMapI(t *testing.T) {
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
		tc := tc //nolint:modernize
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, FilterSeqToMapI(slices.Values(tc.in), transform))
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
		{name: "empty", input: []int{}, expected: map[int]struct{}{}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Keyify(values(tt.input...)))
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
		{name: "drop none", n: 0, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop one", n: 1, expected: []int{1, 2, 3, 4}},
		{name: "drop two", n: 2, expected: []int{2, 3, 4}},
		{name: "drop three", n: 3, expected: []int{3, 4}},
		{name: "drop four", n: 4, expected: []int{4}},
		{name: "drop all", n: 5, expected: nil},
		{name: "drop more than length", n: 6, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Drop(values(0, 1, 2, 3, 4), tt.n)))
		})
	}

	t.Run("panics on negative n", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Drop: n must not be negative", func() {
			Drop(values(0, 1, 2, 3, 4), -1)
		})
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Drop(allStrings, 2)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestDropLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{name: "drop none", n: 0, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop one", n: 1, expected: []int{0, 1, 2, 3}},
		{name: "drop two", n: 2, expected: []int{0, 1, 2}},
		{name: "drop three", n: 3, expected: []int{0, 1}},
		{name: "drop four", n: 4, expected: []int{0}},
		{name: "drop all", n: 5, expected: nil},
		{name: "drop more than length", n: 6, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(DropLast(values(0, 1, 2, 3, 4), tt.n)))
		})
	}

	t.Run("panics on negative n", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.DropLast: n must not be negative", func() {
			DropLast(values(0, 1, 2, 3, 4), -1)
		})
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := DropLast(allStrings, 2)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestDropWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		predicate func(int) bool
		expected  []int
	}{
		{name: "drop until value found", predicate: func(t int) bool { return t != 4 }, expected: []int{4, 5, 6}},
		{name: "predicate always true", predicate: func(t int) bool { return true }, expected: nil},
		{name: "predicate never true", predicate: func(t int) bool { return t == 10 }, expected: []int{0, 1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(DropWhile(values(0, 1, 2, 3, 4, 5, 6), tt.predicate)))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := DropWhile(allStrings, func(t string) bool {
			return t != "foo"
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestDropLastWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		predicate func(int) bool
		expected  []int
	}{
		{name: "drop trailing until value found", predicate: func(t int) bool { return t != 3 }, expected: []int{0, 1, 2, 3}},
		{name: "drop trailing until earlier value found", predicate: func(t int) bool { return t != 1 }, expected: []int{0, 1}},
		{name: "predicate never true", predicate: func(t int) bool { return t == 10 }, expected: []int{0, 1, 2, 3, 4, 5, 6}},
		{name: "predicate always true", predicate: func(t int) bool { return t != 10 }, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(DropLastWhile(values(0, 1, 2, 3, 4, 5, 6), tt.predicate)))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := DropLastWhile(allStrings, func(t string) bool {
			return t != "foo"
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestTake(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{name: "take three", n: 3, expected: []int{0, 1, 2}},
		{name: "take two", n: 2, expected: []int{0, 1}},
		{name: "take one", n: 1, expected: []int{0}},
		{name: "take none", n: 0, expected: nil},
		{name: "take exactly all", n: 5, expected: []int{0, 1, 2, 3, 4}},
		{name: "take more than length", n: 10, expected: []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Take(values(0, 1, 2, 3, 4), tt.n)))
		})
	}

	t.Run("panics on negative n", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.Take: n must not be negative", func() {
			Take(values(0, 1, 2, 3, 4), -1)
		})
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Take(allStrings, 2)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestTakeWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		predicate func(int) bool
		expected  []int
	}{
		{name: "take until value found", predicate: func(t int) bool { return t != 4 }, expected: []int{0, 1, 2, 3}},
		{name: "predicate always true", predicate: func(t int) bool { return true }, expected: []int{0, 1, 2, 3, 4, 5, 6}},
		{name: "predicate always false", predicate: func(t int) bool { return false }, expected: nil},
		{name: "predicate on threshold", predicate: func(t int) bool { return t < 3 }, expected: []int{0, 1, 2}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(TakeWhile(values(0, 1, 2, 3, 4, 5, 6), tt.predicate)))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := TakeWhile(allStrings, func(t string) bool {
			return t != "bar"
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestTakeFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	isEven := func(item int) bool {
		return item%2 == 0
	}

	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{name: "take two evens", input: []int{1, 2, 3, 4, 5, 6}, n: 2, expected: []int{2, 4}},
		{name: "take zero", input: []int{1, 2, 3, 4, 5, 6}, n: 0, expected: nil},
		{name: "no matches", input: []int{1, 3, 5}, n: 2, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(TakeFilter(values(tt.input...), tt.n, isEven)))
		})
	}

	t.Run("panics on negative n", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.TakeFilterI: n must not be negative", func() {
			TakeFilter(values(1, 2, 3), -1, func(item int) bool { return true })
		})
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar", "baz"))
		nonempty := TakeFilter(allStrings, 2, func(item string) bool {
			return item != ""
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestTakeFilterI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     []int
		n         int
		predicate func(int, int) bool
		expected  []int
	}{
		{
			name:      "take two evens with index bound",
			input:     []int{1, 2, 3, 4, 5, 6},
			n:         2,
			predicate: func(item, index int) bool { return item%2 == 0 && index < 4 },
			expected:  []int{2, 4},
		},
		{
			name:      "take more than matches available",
			input:     []int{1, 2, 3, 4, 5, 6},
			n:         10,
			predicate: func(item, _ int) bool { return item%2 == 0 },
			expected:  []int{2, 4, 6},
		},
		{
			name:      "take zero",
			input:     []int{1, 2, 3, 4, 5, 6},
			n:         0,
			predicate: func(item, index int) bool { return item%2 == 0 && index < 4 },
			expected:  nil,
		},
		{
			name:      "no matches",
			input:     []int{1, 3, 5},
			n:         2,
			predicate: func(item, _ int) bool { return item%2 == 0 },
			expected:  nil,
		},
		{
			name:      "take one odd",
			input:     []int{1, 2, 3, 4, 5},
			n:         1,
			predicate: func(item, _ int) bool { return item%2 != 0 },
			expected:  []int{1},
		},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(TakeFilterI(values(tt.input...), tt.n, tt.predicate)))
		})
	}

	t.Run("panics on negative n", func(t *testing.T) {
		t.Parallel()
		is.PanicsWithValue("it.TakeFilterI: n must not be negative", func() {
			TakeFilterI(values(1, 2, 3), -1, func(item, _ int) bool { return true })
		})
	})
}

func TestDropByIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		indices  []int
		expected []int
	}{
		{name: "drop first", input: []int{0, 1, 2, 3, 4}, indices: []int{0}, expected: []int{1, 2, 3, 4}},
		{name: "drop first three in order", input: []int{0, 1, 2, 3, 4}, indices: []int{0, 1, 2}, expected: []int{3, 4}},
		{name: "drop three indices out of order", input: []int{0, 1, 2, 3, 4}, indices: []int{3, 1, 0}, expected: []int{2, 4}},
		{name: "drop middle", input: []int{0, 1, 2, 3, 4}, indices: []int{2}, expected: []int{0, 1, 3, 4}},
		{name: "drop last", input: []int{0, 1, 2, 3, 4}, indices: []int{4}, expected: []int{0, 1, 2, 3}},
		{name: "drop out of range index", input: []int{0, 1, 2, 3, 4}, indices: []int{5}, expected: []int{0, 1, 2, 3, 4}},
		{name: "drop far out of range index", input: []int{0, 1, 2, 3, 4}, indices: []int{100}, expected: []int{0, 1, 2, 3, 4}},
		{name: "empty collection", input: []int{}, indices: []int{0, 1}, expected: nil},
		{name: "single element drop first and second", input: []int{42}, indices: []int{0, 1}, expected: nil},
		{name: "single element drop second and first", input: []int{42}, indices: []int{1, 0}, expected: nil},
		{name: "empty collection drop one index", input: []int{}, indices: []int{1}, expected: nil},
		{name: "single element drop it", input: []int{1}, indices: []int{0}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(DropByIndex(values(tt.input...), tt.indices...)))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := DropByIndex(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int evens rejected", func(t *testing.T) {
		t.Parallel()
		r1 := Reject(values(1, 2, 3, 4), func(x int) bool {
			return x%2 == 0
		})
		is.Equal([]int{1, 3}, slices.Collect(r1))
	})

	t.Run("long names rejected", func(t *testing.T) {
		t.Parallel()
		r2 := Reject(values("Smith", "foo", "Domin", "bar", "Olivia"), func(x string) bool {
			return len(x) > 3
		})
		is.Equal([]string{"foo", "bar"}, slices.Collect(r2))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Reject(allStrings, func(x string) bool {
			return len(x) > 0
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestRejectI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int evens rejected", func(t *testing.T) {
		t.Parallel()
		r1 := RejectI(values(1, 2, 3, 4), func(x, _ int) bool {
			return x%2 == 0
		})
		is.Equal([]int{1, 3}, slices.Collect(r1))
	})

	t.Run("long names rejected", func(t *testing.T) {
		t.Parallel()
		r2 := RejectI(values("Smith", "foo", "Domin", "bar", "Olivia"), func(x string, _ int) bool {
			return len(x) > 3
		})
		is.Equal([]string{"foo", "bar"}, slices.Collect(r2))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := RejectI(allStrings, func(x string, _ int) bool {
			return len(x) > 0
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestRejectMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int64 evens kept", func(t *testing.T) {
		t.Parallel()
		r1 := RejectMap(values[int64](1, 2, 3, 4), func(x int64) (string, bool) {
			if x%2 == 0 {
				return strconv.FormatInt(x, 10), false
			}
			return "", true
		})
		is.Equal([]string{"2", "4"}, slices.Collect(r1))
	})

	t.Run("string suffix kept", func(t *testing.T) {
		t.Parallel()
		r2 := RejectMap(values("cpu", "gpu", "mouse", "keyboard"), func(x string) (string, bool) {
			if strings.HasSuffix(x, "pu") {
				return "xpu", false
			}
			return "", true
		})
		is.Equal([]string{"xpu", "xpu"}, slices.Collect(r2))
	})
}

func TestRejectMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int64 evens kept", func(t *testing.T) {
		t.Parallel()
		r1 := RejectMapI(values[int64](1, 2, 3, 4), func(x int64, _ int) (string, bool) {
			if x%2 == 0 {
				return strconv.FormatInt(x, 10), false
			}
			return "", true
		})
		is.Equal([]string{"2", "4"}, slices.Collect(r1))
	})

	t.Run("string suffix kept", func(t *testing.T) {
		t.Parallel()
		r2 := RejectMapI(values("cpu", "gpu", "mouse", "keyboard"), func(x string, _ int) (string, bool) {
			if strings.HasSuffix(x, "pu") {
				return "xpu", false
			}
			return "", true
		})
		is.Equal([]string{"xpu", "xpu"}, slices.Collect(r2))
	})
}

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		needle   int
		expected int
	}{
		{name: "value present twice", input: []int{1, 2, 1}, needle: 1, expected: 2},
		{name: "value absent", input: []int{1, 2, 1}, needle: 3, expected: 0},
		{name: "empty collection", input: []int{}, needle: 1, expected: 0},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Count(values(tt.input...), tt.needle))
		})
	}
}

func TestCountBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  int
	}{
		{name: "matches found", input: []int{1, 2, 1}, predicate: func(i int) bool { return i < 2 }, expected: 2},
		{name: "no matches", input: []int{1, 2, 1}, predicate: func(i int) bool { return i > 2 }, expected: 0},
		{name: "empty collection", input: []int{}, predicate: func(i int) bool { return i <= 2 }, expected: 0},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, CountBy(values(tt.input...), tt.predicate))
		})
	}
}

func TestCountValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is.Empty(CountValues(values[int]()))
		is.Equal(map[int]int{1: 1, 2: 1}, CountValues(values(1, 2)))
		is.Equal(map[int]int{1: 1, 2: 2}, CountValues(values(1, 2, 2)))
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		is.Equal(map[string]int{"": 1, "foo": 1, "bar": 1}, CountValues(values("foo", "bar", "")))
		is.Equal(map[string]int{"foo": 1, "bar": 2}, CountValues(values("foo", "bar", "bar")))
	})
}

func TestCountValuesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	oddEven := func(v int) bool {
		return v%2 == 0
	}
	length := func(v string) int {
		return len(v)
	}

	t.Run("ints by oddEven", func(t *testing.T) {
		t.Parallel()
		is.Empty(CountValuesBy(values[int](), oddEven))
		is.Equal(map[bool]int{true: 1, false: 1}, CountValuesBy(values(1, 2), oddEven))
		is.Equal(map[bool]int{true: 2, false: 1}, CountValuesBy(values(1, 2, 2), oddEven))
	})

	t.Run("strings by length", func(t *testing.T) {
		t.Parallel()
		is.Equal(map[int]int{0: 1, 3: 2}, CountValuesBy(values("foo", "bar", ""), length))
		is.Equal(map[int]int{3: 3}, CountValuesBy(values("foo", "bar", "bar"), length))
	})
}

func TestSubset(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		offset   int
		length   int
		expected []int
	}{
		{name: "zero length", offset: 0, length: 0, expected: nil},
		{name: "offset beyond collection", offset: 10, length: 2, expected: nil},
		{name: "length beyond collection", offset: 0, length: 10, expected: []int{0, 1, 2, 3, 4}},
		{name: "beginning slice", offset: 0, length: 2, expected: []int{0, 1}},
		{name: "middle slice exact length", offset: 2, length: 2, expected: []int{2, 3}},
		{name: "middle to end length exceeds", offset: 2, length: 5, expected: []int{2, 3, 4}},
		{name: "middle to end exact remaining", offset: 2, length: 3, expected: []int{2, 3, 4}},
		{name: "middle to end length larger than remaining", offset: 2, length: 4, expected: []int{2, 3, 4}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Subset(values(0, 1, 2, 3, 4), tt.offset, tt.length)))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Subset(allStrings, 0, 2)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		start    int
		end      int
		expected []int
	}{
		{name: "empty range at start", start: 0, end: 0, expected: nil},
		{name: "single element at start", start: 0, end: 1, expected: []int{0}},
		{name: "full range", start: 0, end: 5, expected: []int{0, 1, 2, 3, 4}},
		{name: "end beyond collection", start: 0, end: 6, expected: []int{0, 1, 2, 3, 4}},
		{name: "empty range in middle", start: 1, end: 1, expected: nil},
		{name: "from middle to end", start: 1, end: 5, expected: []int{1, 2, 3, 4}},
		{name: "from middle beyond end", start: 1, end: 6, expected: []int{1, 2, 3, 4}},
		{name: "last element", start: 4, end: 5, expected: []int{4}},
		{name: "empty range at end", start: 5, end: 5, expected: nil},
		{name: "start beyond collection end before", start: 6, end: 5, expected: nil},
		{name: "start and end beyond collection", start: 6, end: 6, expected: nil},
		{name: "start after end", start: 1, end: 0, expected: nil},
		{name: "start after end both large", start: 5, end: 0, expected: nil},
		{name: "start beyond collection end within", start: 6, end: 4, expected: nil},
		{name: "start and end both beyond collection", start: 6, end: 7, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Slice(values(0, 1, 2, 3, 4), tt.start, tt.end)))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Slice(allStrings, 0, 2)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestReplace(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		old      int
		new      int
		n        int
		expected []int
	}{
		{name: "replace first two occurrences", old: 0, new: 42, n: 2, expected: []int{42, 1, 42, 1, 2, 3, 0}},
		{name: "replace first occurrence", old: 0, new: 42, n: 1, expected: []int{42, 1, 0, 1, 2, 3, 0}},
		{name: "replace zero occurrences (n=0)", old: 0, new: 42, n: 0, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "replace all occurrences (n=-1) call one", old: 0, new: 42, n: -1, expected: []int{42, 1, 42, 1, 2, 3, 42}},
		{name: "replace all occurrences (n=-1) call two (stateless)", old: 0, new: 42, n: -1, expected: []int{42, 1, 42, 1, 2, 3, 42}},
		{name: "old value not present n=2", old: -1, new: 42, n: 2, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "old value not present n=1", old: -1, new: 42, n: 1, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "old value not present n=0", old: -1, new: 42, n: 0, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "old value not present n=-1 call one", old: -1, new: 42, n: -1, expected: []int{0, 1, 0, 1, 2, 3, 0}},
		{name: "old value not present n=-1 call two (stateless)", old: -1, new: 42, n: -1, expected: []int{0, 1, 0, 1, 2, 3, 0}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			in := values(0, 1, 0, 1, 2, 3, 0)
			is.Equal(tt.expected, slices.Collect(Replace(in, tt.old, tt.new, tt.n)))
		})
	}

	t.Run("repeated iteration yields consistent results", func(t *testing.T) {
		t.Parallel()
		in := values(0, 1, 0, 1, 2, 3, 0)
		out1 := Replace(in, 0, 42, 2)
		is.Equal([]int{42, 1, 42, 1, 2, 3, 0}, slices.Collect(out1))
		is.Equal([]int{42, 1, 42, 1, 2, 3, 0}, slices.Collect(out1)) // check no counter mutation
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Replace(allStrings, "0", "2", 1)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		old      int
		new      int
		expected []int
	}{
		{name: "old value present", old: 0, new: 42, expected: []int{42, 1, 42, 1, 2, 3, 42}},
		{name: "old value absent", old: -1, new: 42, expected: []int{0, 1, 0, 1, 2, 3, 0}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			in := values(0, 1, 0, 1, 2, 3, 0)
			is.Equal(tt.expected, slices.Collect(ReplaceAll(in, tt.old, tt.new)))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := ReplaceAll(allStrings, "0", "2")
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestCompact(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		r1 := Compact(values(2, 0, 4, 0))
		is.Equal([]int{2, 4}, slices.Collect(r1))
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		r2 := Compact(values("", "foo", "", "bar", ""))
		is.Equal([]string{"foo", "bar"}, slices.Collect(r2))
	})

	t.Run("bools", func(t *testing.T) {
		t.Parallel()
		r3 := Compact(values(true, false, true, false))
		is.Equal([]bool{true, true}, slices.Collect(r3))
	})

	t.Run("structs", func(t *testing.T) {
		t.Parallel()
		type foo struct {
			bar int
			baz string
		}

		// sequence of structs
		// If all fields of an element are zero values, Compact removes it.

		r4 := Compact(values(
			foo{bar: 1, baz: "a"}, // all fields are non-zero values
			foo{bar: 0, baz: ""},  // all fields are zero values
			foo{bar: 2, baz: ""},  // bar is non-zero
		))

		is.Equal([]foo{{bar: 1, baz: "a"}, {bar: 2, baz: ""}}, slices.Collect(r4))
	})

	t.Run("pointers to structs", func(t *testing.T) {
		t.Parallel()
		type foo struct {
			bar int
			baz string
		}

		// sequence of pointers to structs
		// If an element is nil, Compact removes it.

		e1, e2, e3 := foo{bar: 1, baz: "a"}, foo{bar: 0, baz: ""}, foo{bar: 2, baz: ""}
		// NOTE: e2 is a zero value of foo, but its pointer &e2 is not a zero value of *foo.
		r5 := Compact(values(&e1, &e2, nil, &e3))

		is.Equal([]*foo{&e1, &e2, &e3}, slices.Collect(r5))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Compact(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestIsSorted(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("ints", func(t *testing.T) {
		t.Parallel()
		is.True(IsSorted(values(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)))
		is.False(IsSorted(values(0, 1, 4, 3, 2, 5, 6, 7, 8, 9, 10)))
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		is.True(IsSorted(values("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")))
		is.False(IsSorted(values("a", "b", "d", "c", "e", "f", "g", "h", "i", "j")))
	})
}

func TestIsSortedBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     []string
		transform func(string) int
		expected  bool
	}{
		{name: "sorted by length", input: []string{"a", "bb", "ccc"}, transform: func(s string) int { return len(s) }, expected: true},
		{name: "not sorted by length", input: []string{"aa", "b", "ccc"}, transform: func(s string) int { return len(s) }, expected: false},
		{name: "sorted by numeric value", input: []string{"1", "2", "3", "11"}, transform: func(s string) int {
			ret, _ := strconv.Atoi(s)
			return ret
		}, expected: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, IsSortedBy(values(tt.input...), tt.transform))
		})
	}
}

func TestSplice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("normal case", func(t *testing.T) {
		t.Parallel()
		sample := values("a", "b", "c", "d", "e", "f", "g")
		results := slices.Collect(Splice(sample, 1, "1", "2"))
		is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(sample))
		is.Equal([]string{"a", "1", "2", "b", "c", "d", "e", "f", "g"}, results)
	})

	t.Run("positive overflow", func(t *testing.T) {
		t.Parallel()
		sample := values("a", "b", "c", "d", "e", "f", "g")
		results := slices.Collect(Splice(sample, 42, "1", "2"))
		is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(sample))
		is.Equal([]string{"a", "b", "c", "d", "e", "f", "g", "1", "2"}, results)
	})

	t.Run("other cases", func(t *testing.T) {
		t.Parallel()
		is.Equal([]string{"1", "2"}, slices.Collect(Splice(values[string](), 0, "1", "2")))
		is.Equal([]string{"1", "2"}, slices.Collect(Splice(values[string](), 1, "1", "2")))
		is.Equal([]string{"1", "2", "0"}, slices.Collect(Splice(values("0"), 0, "1", "2")))
		is.Equal([]string{"0", "1", "2"}, slices.Collect(Splice(values("0"), 1, "1", "2")))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Splice(allStrings, 1, "1", "2")
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestCutPrefix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name          string
		input         []string
		prefix        []string
		expectedFound bool
		expectedAfter []string
	}{
		{name: "prefix matches", input: []string{"a", "a", "b"}, prefix: []string{"a"}, expectedFound: true, expectedAfter: []string{"a", "b"}},
		{name: "prefix matches (stateless repeat)", input: []string{"a", "a", "b"}, prefix: []string{"a"}, expectedFound: true, expectedAfter: []string{"a", "b"}},
		{name: "prefix does not match", input: []string{"a", "a", "b"}, prefix: []string{"b"}, expectedFound: false, expectedAfter: []string{"a", "a", "b"}},
		{name: "empty collection", input: []string{}, prefix: []string{"b"}, expectedFound: false, expectedAfter: nil},
		{name: "empty prefix", input: []string{"a", "a", "b"}, prefix: []string{}, expectedFound: true, expectedAfter: []string{"a", "a", "b"}},
		{name: "prefix longer than collection", input: []string{"a", "a", "b"}, prefix: []string{"a", "a", "b", "b"}, expectedFound: false, expectedAfter: []string{"a", "a", "b"}},
		{name: "prefix mismatched in the middle", input: []string{"a", "a", "b"}, prefix: []string{"a", "b"}, expectedFound: false, expectedAfter: []string{"a", "a", "b"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, result := CutPrefix(values(tt.input...), tt.prefix)
			is.Equal(tt.expectedFound, result)
			is.Equal(tt.expectedAfter, slices.Collect(actual))
		})
	}
}

func TestCutSuffix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name          string
		suffix        []string
		expectedFound bool
		expectedAfter []string
	}{
		{name: "suffix does not match", suffix: []string{"c"}, expectedFound: false, expectedAfter: []string{"a", "a", "b"}},
		{name: "suffix matches", suffix: []string{"b"}, expectedFound: true, expectedAfter: []string{"a", "a"}},
		{name: "empty suffix", suffix: []string{}, expectedFound: true, expectedAfter: []string{"a", "a", "b"}},
		{name: "suffix mismatched", suffix: []string{"a"}, expectedFound: false, expectedAfter: []string{"a", "a", "b"}},
		{name: "empty suffix (stateless repeat)", suffix: []string{}, expectedFound: true, expectedAfter: []string{"a", "a", "b"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, result := CutSuffix(values("a", "a", "b"), tt.suffix)
			is.Equal(tt.expectedFound, result)
			is.Equal(tt.expectedAfter, slices.Collect(actual))
		})
	}
}

func TestTrim(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		cutset   []string
		expected []string
	}{
		{name: "trim leading a b", cutset: []string{"a", "b"}, expected: []string{"c", "d", "e", "f", "g"}},
		{name: "trim trailing g f", cutset: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e"}},
		{name: "trim all elements", cutset: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: nil},
		{name: "cutset superset of elements", cutset: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: nil},
		{name: "empty cutset", cutset: nil, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Trim(values("a", "b", "c", "d", "e", "f", "g"), tt.cutset...)))
		})
	}
}

func TestTrimFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		cutset   []string
		expected []string
	}{
		{name: "trim first two matching leading elements", input: []string{"a", "a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b"}, expected: []string{"c", "d", "e", "f", "g"}},
		{name: "trim leading b a", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"b", "a"}, expected: []string{"c", "d", "e", "f", "g"}},
		{name: "cutset does not match leading elements", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "trim all elements", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: nil},
		{name: "cutset superset of elements", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: nil},
		{name: "empty cutset", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: nil, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(TrimFirst(values(tt.input...), tt.cutset...)))
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
		{name: "prefix matches leading elements", input: []string{"a", "b", "a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"a", "b"}, expected: []string{"c", "d", "e", "f", "g"}},
		{name: "prefix order mismatch", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"b", "a"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "prefix does not match leading elements", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "prefix equals entire collection", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: nil},
		{name: "prefix longer than collection", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "empty prefix", input: []string{"a", "b", "c", "d", "e", "f", "g"}, prefix: []string{}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(TrimPrefix(values(tt.input...), tt.prefix)))
		})
	}
}

func TestTrimLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		cutset   []string
		expected []string
	}{
		{name: "cutset does not match trailing elements", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "trim trailing g f", input: []string{"a", "b", "c", "d", "e", "f", "g", "g"}, cutset: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e"}},
		{name: "trim all elements", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: nil},
		{name: "cutset superset of elements", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: nil},
		{name: "empty cutset", input: []string{"a", "b", "c", "d", "e", "f", "g"}, cutset: nil, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(TrimLast(values(tt.input...), tt.cutset...)))
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
		{name: "suffix does not match trailing elements", input: []string{"a", "b", "c", "d", "e", "f", "g"}, suffix: []string{"a", "b"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "suffix matches trailing elements", input: []string{"a", "b", "c", "d", "e", "f", "g", "f", "g"}, suffix: []string{"f", "g"}, expected: []string{"a", "b", "c", "d", "e"}},
		{name: "suffix order mismatch", input: []string{"a", "b", "c", "d", "e", "f", "g", "f", "g"}, suffix: []string{"g", "f"}, expected: []string{"a", "b", "c", "d", "e", "f", "g", "f", "g"}},
		{name: "suffix matches with repeated element earlier", input: []string{"a", "b", "c", "d", "e", "f", "f", "g"}, suffix: []string{"f", "g"}, expected: []string{"a", "b", "c", "d", "e", "f"}},
		{name: "suffix equals entire collection", input: []string{"a", "b", "c", "d", "e", "f", "g"}, suffix: []string{"a", "b", "c", "d", "e", "f", "g"}, expected: nil},
		{name: "suffix longer than collection", input: []string{"a", "b", "c", "d", "e", "f", "g"}, suffix: []string{"a", "b", "c", "d", "e", "f", "g", "h"}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
		{name: "empty suffix", input: []string{"a", "b", "c", "d", "e", "f", "g"}, suffix: []string{}, expected: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(TrimSuffix(values(tt.input...), tt.suffix)))
		})
	}
}

func TestBuffer(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		from     int
		to       int
		size     int
		expected [][]int
	}{
		{name: "full batches", from: 1, to: 6, size: 2, expected: [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{name: "partial last batch", from: 1, to: 5, size: 2, expected: [][]int{{1, 2}, {3, 4}, {5}}},
		{name: "empty channel", from: 1, to: 0, size: 2, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, slices.Collect(Buffer(RangeFrom(tt.from, tt.to), tt.size)))
		})
	}

	t.Run("early termination", func(t *testing.T) {
		t.Parallel()
		batches4 := slices.Collect(Take(Buffer(RangeFrom(1, 6), 2), 1))
		is.Equal([][]int{{1, 2}}, batches4)
	})
}

func TestSeqToSeq2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []string
		expected map[int]string
	}{
		{name: "two elements", input: []string{"foo", "bar"}, expected: map[int]string{0: "foo", 1: "bar"}},
		{name: "empty", input: []string{}, expected: map[int]string{}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, maps.Collect(SeqToSeq2(values(tt.input...))))
		})
	}
}
