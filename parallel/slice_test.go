package parallel

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	is := assert.New(t)

	result1 := Map([]int{1, 2, 3, 4}, func(x int, _ int) string {
		return "Hello"
	})
	result2 := Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "3", "4"})
}

func TestFilterMap(t *testing.T) {
	is := assert.New(t)

	r1 := FilterMap([]int64{1, 2, 3, 2, 1}, func(x int64, _ int) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), true
		}
		return "", false
	})
	r2 := FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})

	is.Equal(len(r1), 2)
	is.Equal(len(r2), 2)
	is.Equal(r1, []string{"2", "2"})
	is.Equal(r2, []string{"xpu", "xpu"})
}

func TestTimes(t *testing.T) {
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []string{"0", "1", "2"})
}

func TestGroupBy(t *testing.T) {
	is := assert.New(t)

	result1 := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	// order
	for x := range result1 {
		sort.Slice(result1[x], func(i, j int) bool {
			return result1[x][i] < result1[x][j]
		})
	}

	is.EqualValues(len(result1), 3)
	is.EqualValues(result1, map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	})
}

func TestPartitionBy(t *testing.T) {
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
		sort.Slice(result1[x], func(i, j int) bool {
			return result1[x][i] < result1[x][j]
		})
	}

	is.ElementsMatch(result1, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Equal(result2, [][]int{})
}
