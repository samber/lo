package noindex

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{2, 4})

	r2 := Filter([]string{"", "foo", "", "bar", ""}, func(x string) bool {
		return len(x) > 0
	})

	is.Equal(r2, []string{"foo", "bar"})
}

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Map([]int{1, 2, 3, 4}, func(x int) string {
		return "Hello"
	})
	result2 := Map([]int64{1, 2, 3, 4}, func(x int64) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "3", "4"})
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterMap([]int64{1, 2, 3, 4}, func(x int64) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), true
		}
		return "", false
	})
	r2 := FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})

	is.Equal(len(r1), 2)
	is.Equal(len(r2), 2)
	is.Equal(r1, []string{"2", "4"})
	is.Equal(r2, []string{"xpu", "xpu"})
}

func TestFlatMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FlatMap([]int{0, 1, 2, 3, 4}, func(x int) []string {
		return []string{"Hello"}
	})
	result2 := FlatMap([]int64{0, 1, 2, 3, 4}, func(x int64) []string {
		result := make([]string, 0, x)
		for i := int64(0); i < x; i++ {
			result = append(result, strconv.FormatInt(x, 10))
		}
		return result
	})

	is.Equal(len(result1), 5)
	is.Equal(len(result2), 10)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"})
}

func TestReduce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reduce([]int{1, 2, 3, 4}, func(agg int, item int) int {
		return agg + item
	}, 0)
	result2 := Reduce([]int{1, 2, 3, 4}, func(agg int, item int) int {
		return agg + item
	}, 10)

	is.Equal(result1, 10)
	is.Equal(result2, 20)
}

func TestReduceRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int) []int {
		return append(agg, item...)
	}, []int{})

	is.Equal(result1, []int{4, 5, 2, 3, 0, 1})
}

func TestForEach(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams := []string{}

	ForEach([]string{"a", "b", "c"}, func(item string) {
		callParams = append(callParams, item)
	})

	is.ElementsMatch([]string{"a", "b", "c"}, callParams)
	is.IsIncreasing(callParams)
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Reject([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{1, 3})

	r2 := Reject([]string{"Smith", "foo", "Domin", "bar", "Olivia"}, func(x string) bool {
		return len(x) > 3
	})

	is.Equal(r2, []string{"foo", "bar"})
}
