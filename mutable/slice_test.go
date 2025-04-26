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

	is.Equal(input1, []int{2, 4, 3, 4})
	is.Equal(r1, []int{2, 4})

	input2 := []string{"", "foo", "", "bar", ""}
	r2 := Filter(input2, func(x string) bool {
		return len(x) > 0
	})

	is.Equal(input2, []string{"foo", "bar", "", "bar", ""})
	is.Equal(r2, []string{"foo", "bar"})
}

func TestFilterI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterI([]int{1, 2, 3, 4}, func(x int, i int) bool {
		is.Equal(i, x-1)
		return x%2 == 0
	})

	is.Equal(r1, []int{2, 4})
}

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list := []int{1, 2, 3, 4}
	Map(list, func(x int) int {
		return x * 2
	})
	is.Equal(len(list), 4)
	is.Equal(list, []int{2, 4, 6, 8})

	list = []int{1, 2, 3, 4}
	Map(list, func(x int) int {
		return x * 4
	})
	is.Equal(len(list), 4)
	is.Equal(list, []int{4, 8, 12, 16})
}

func TestMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list := []int{1, 2, 3, 4}
	MapI(list, func(x int, index int) int {
		is.Equal(index, x-1)
		return x * 2
	})
	is.Equal(len(list), 4)
	is.Equal(list, []int{2, 4, 6, 8})

	list = []int{1, 2, 3, 4}
	MapI(list, func(x int, index int) int {
		is.Equal(index, x-1)
		return x * 4
	})
	is.Equal(len(list), 4)
	is.Equal(list, []int{4, 8, 12, 16})
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Shuffle(list)
	is.NotEqual(list, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	list = []int{}
	Shuffle(list)
	is.Equal(list, []int{})
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list := []int{0, 1, 2, 3, 4, 5}
	Reverse(list)
	is.Equal(list, []int{5, 4, 3, 2, 1, 0})

	list = []int{0, 1, 2, 3, 4, 5, 6}
	Reverse(list)
	is.Equal(list, []int{6, 5, 4, 3, 2, 1, 0})

	list = []int{}
	Reverse(list)
	is.Equal(list, []int{})

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
	is.Equal([]string{}, list2)
}
