package parallel

import (
	"sort"
	"strconv"
	"testing"
	"time"

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

func TestMapConcurrency(t *testing.T) {
	is := assert.New(t)

	num := 100
	concurrency := 5
	duration := 100 * time.Millisecond

	list := []int{}
	expected := []string{}
	for i := 0; i < num; i++ {
		list = append(list, i)
		expected = append(expected, strconv.Itoa(i))
	}

	handler := func(item int, ix int) string {
		time.Sleep(duration)
		return strconv.Itoa(item)
	}

	startAt := time.Now()
	result := Map(list, handler, WithConcurrency(concurrency))
	endAt := time.Now()

	is.Equal(result, expected)
	is.Equal(endAt.Sub(startAt).Round(duration), time.Duration(num/concurrency)*duration)
}

func TestTimes(t *testing.T) {
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []string{"0", "1", "2"})
}

func TestTimesConcurrency(t *testing.T) {
	is := assert.New(t)

	num := 100
	concurrency := 5
	duration := 100 * time.Millisecond

	expected := []string{}
	for i := 0; i < num; i++ {
		expected = append(expected, strconv.Itoa(i))
	}

	startAt := time.Now()
	result := Times(num, func(i int) string {
		time.Sleep(duration)
		return strconv.FormatInt(int64(i), 10)
	}, WithConcurrency(concurrency))
	endAt := time.Now()

	is.Equal(len(result), num)
	is.Equal(result, expected)
	is.Equal(endAt.Sub(startAt).Round(duration), time.Duration(num/concurrency)*duration)
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

func TestGroupByConcurrency(t *testing.T) {
	is := assert.New(t)

	concurrency := 2
	duration := 100 * time.Millisecond

	startAt := time.Now()
	result := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		time.Sleep(duration)
		return i % 3
	}, WithConcurrency(concurrency))
	endAt := time.Now()

	// order
	for x := range result {
		sort.Slice(result[x], func(i, j int) bool {
			return result[x][i] < result[x][j]
		})
	}

	is.EqualValues(len(result), 3)
	is.EqualValues(result, map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	})
	is.Equal(endAt.Sub(startAt).Round(duration), time.Duration(6/concurrency)*duration)
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

func TestPartitionByConcurrency(t *testing.T) {
	is := assert.New(t)

	concurrency := 2
	duration := 100 * time.Millisecond

	oddEven := func(x int) string {
		time.Sleep(duration)
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	}

	startAt := time.Now()
	result := PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven, WithConcurrency(concurrency))
	endAt := time.Now()

	// order
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	for x := range result {
		sort.Slice(result[x], func(i, j int) bool {
			return result[x][i] < result[x][j]
		})
	}

	is.ElementsMatch(result, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Equal(endAt.Sub(startAt).Round(duration), time.Duration(8/concurrency)*duration)
}
