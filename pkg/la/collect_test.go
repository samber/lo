package la

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reduce(slices.Values([]int{1, 2, 3, 4}), func(agg int, item int) int {
		return agg + item
	}, 0)
	result2 := Reduce(slices.Values([]int{1, 2, 3, 4}), func(agg int, item int) int {
		return agg + item
	}, 10)

	is.Equal(result1, 10)
	is.Equal(result2, 20)
}

func TestReduce2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reduce2(Enumerate(slices.Values([]int{1, 2, 3, 4})), func(agg int, _ int, item int) int {
		return agg + item
	}, 0)
	result2 := Reduce2(Enumerate(slices.Values([]int{1, 2, 3, 4})), func(agg int, _ int, item int) int {
		return agg + item
	}, 10)

	is.Equal(result1, 10)
	is.Equal(result2, 20)
}

func TestForEach(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}

	ForEach(slices.Values([]string{"a", "b", "c"}), func(item string) {
		callParams1 = append(callParams1, item)
	})

	is.ElementsMatch([]string{"a", "b", "c"}, callParams1)
}

func TestForEach2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	ForEach2(Enumerate(slices.Values([]string{"a", "b", "c"})), func(i int, item string) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.ElementsMatch([]string{"a", "b", "c"}, callParams1)
	is.ElementsMatch([]int{0, 1, 2}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestForEachWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	var callParams1 []string

	ForEachWhile(slices.Values([]string{"a", "b", "c"}), func(item string) bool {
		if item == "c" {
			return false
		}
		callParams1 = append(callParams1, item)

		return true
	})

	is.ElementsMatch([]string{"a", "b"}, callParams1)
}

func TestForEachWhile2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	var callParams1 []string
	var callParams2 []int

	ForEachWhile2(Enumerate(slices.Values([]string{"a", "b", "c"})), func(i int, item string) bool {
		if item == "c" {
			return false
		}
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
		return true
	})

	is.ElementsMatch([]string{"a", "b"}, callParams1)
	is.ElementsMatch([]int{0, 1}, callParams2)
	is.IsIncreasing(callParams2)
}
