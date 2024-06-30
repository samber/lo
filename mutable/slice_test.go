package mutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := []int{1, 2, 3, 4}
	Filter(&r1, func(x int) bool {
		return x%2 == 0
	})
	is.Equal([]int{2, 4}, r1)

	r2 := []string{"", "foo", "", "bar", ""}
	Filter(&r2, func(x string) bool {
		return len(x) > 0
	})
	is.Equal([]string{"foo", "bar"}, r2)
}

func TestFilterI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := []int{1, 2, 3, 4}
	FilterI(&r1, func(x int, _ int) bool {
		return x%2 == 0
	})
	is.Equal([]int{2, 4}, r1)

	r2 := []string{"", "foo", "", "bar", ""}
	FilterI(&r2, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.Equal([]string{"foo", "bar"}, r2)
}

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := []int{1, 2, 2, 1}
	Uniq(&result1)
	is.Equal(len(result1), 2)
	is.Equal(result1, []int{1, 2})
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := []int{0, 1, 2, 3, 4, 5}
	UniqBy(&result1, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []int{0, 1, 2})
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Shuffle(result1)
	is.NotEqual(result1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result2 := []int{}
	Shuffle(result2)
	is.Equal(result2, []int{})
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := []int{0, 1, 2, 3, 4, 5}
	Reverse(result1)
	is.Equal(result1, []int{5, 4, 3, 2, 1, 0})

	result2 := []int{0, 1, 2, 3, 4, 5, 6}
	Reverse(result2)
	is.Equal(result2, []int{6, 5, 4, 3, 2, 1, 0})

	result3 := []int{}
	Reverse(result3)
	is.Equal(result3, []int{})
}
