package mutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
