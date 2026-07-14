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

	t.Run("constant transform", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Map([]int{1, 2, 3, 4}, func(x, _ int) string {
			return "Hello"
		})
		is.Equal([]string{"Hello", "Hello", "Hello", "Hello"}, result)
	})

	t.Run("formats int64 as string", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
			return strconv.FormatInt(x, 10)
		})
		is.Equal([]string{"1", "2", "3", "4"}, result)
	})
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

	t.Run("groups by modulo", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
			return i % 3
		})

		// order
		for x := range result {
			sort.Ints(result[x])
		}

		is.Equal(map[int][]int{
			0: {0, 3},
			1: {1, 4},
			2: {2, 5},
		}, result)
	})

	t.Run("preserves named slice type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := GroupBy(allStrings, func(i string) int {
			return 42
		})
		is.IsType(nonempty[42], allStrings, "type preserved")
	})
}

func TestPartitionBy(t *testing.T) {
	t.Parallel()

	oddEven := func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	}

	t.Run("partitions a non-empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)

		// order
		sort.Slice(result, func(i, j int) bool {
			return result[i][0] < result[j][0]
		})
		for x := range result {
			sort.Ints(result[x])
		}

		is.ElementsMatch(result, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	})

	t.Run("empty collection", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := PartitionBy([]int{}, oddEven)
		is.Empty(result)
	})

	t.Run("preserves named slice type", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := PartitionBy(allStrings, func(item string) int {
			return len(item)
		})
		is.IsType(nonempty[0], allStrings, "type preserved")
	})
}
