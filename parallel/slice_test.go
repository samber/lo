package parallel

import (
	"sort"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Map([]int{1, 2, 3, 4}, func(x, _ int) string {
		return "Hello"
	})
	result2 := Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal([]string{"Hello", "Hello", "Hello", "Hello"}, result1)
	is.Equal([]string{"1", "2", "3", "4"}, result2)
}

func TestForEach(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var counter uint64
	collection := []int{1, 2, 3, 4}
	ForEach(collection, func(x, i int) {
		atomic.AddUint64(&counter, 1)
	})

	is.Equal(uint64(4), atomic.LoadUint64(&counter))
}

func TestTimes(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal([]string{"0", "1", "2"}, result1)
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	// order
	for x := range result1 {
		sort.Ints(result1[x])
	}

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

	result1 := PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)
	result2 := PartitionBy([]int{}, oddEven)

	// order
	sort.Slice(result1, func(i, j int) bool {
		return result1[i][0] < result1[j][0]
	})
	for x := range result1 {
		sort.Ints(result1[x])
	}

	is.ElementsMatch(result1, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Empty(result2)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := PartitionBy(allStrings, func(item string) int {
		return len(item)
	})
	is.IsType(nonempty[0], allStrings, "type preserved")
}

func TestFilter(t *testing.T) {
	is := assert.New(t)

	r1 := Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})

	is.Equal(r1, []int{2, 4})

	r2 := Filter([]string{"", "bar", "", "foo", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})

	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	is.Equal(r2, []string{"bar", "foo"})
}
