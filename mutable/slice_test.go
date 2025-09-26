package mutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	input1 := []int{1, 2, 3, 4}
	r1 := Filter(input1, func(x int) bool {
		return x%2 == 0
	})

	is.Equal([]int{2, 4, 3, 4}, input1)
	is.Equal([]int{2, 4}, r1)

	input2 := []string{"", "foo", "", "bar", ""}
	r2 := Filter(input2, func(x string) bool {
		return len(x) > 0
	})

	is.Equal([]string{"foo", "bar", "", "bar", ""}, input2)
	is.Equal([]string{"foo", "bar"}, r2)
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

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list := []int{1, 2, 3, 4}
	Map(list, func(x int) int {
		return x * 2
	})
	is.Equal([]int{2, 4, 6, 8}, list)

	list = []int{1, 2, 3, 4}
	Map(list, func(x int) int {
		return x * 4
	})
	is.Equal([]int{4, 8, 12, 16}, list)
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
	is := assert.New(t)

	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Shuffle(list)
	is.NotEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, list)

	list = []int{}
	Shuffle(list)
	is.Empty(list)
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list := []int{0, 1, 2, 3, 4, 5}
	Reverse(list)
	is.Equal([]int{5, 4, 3, 2, 1, 0}, list)

	list = []int{0, 1, 2, 3, 4, 5, 6}
	Reverse(list)
	is.Equal([]int{6, 5, 4, 3, 2, 1, 0}, list)

	list = []int{}
	Reverse(list)
	is.Empty(list)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	Reverse(allStrings)
	is.IsType(myStrings{"", "foo", "bar"}, allStrings, "type preserved")
}

func TestFill(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list1 := []string{"a", "0"}
	Fill(list1, "b")
	is.Equal([]string{"b", "b"}, list1)

	list2 := []string{}
	Fill(list2, "b")
	is.Empty(list2)
}
