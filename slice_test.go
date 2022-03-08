package lo

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}

func TestFilter(t *testing.T) {
	is := assert.New(t)

	r1 := Filter[int]([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{2, 4})

	r2 := Filter[string]([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})

	is.Equal(r2, []string{"foo", "bar"})
}

func TestMap(t *testing.T) {
	is := assert.New(t)

	result1 := Map[int, string]([]int{1, 2, 3, 4}, func(x int, _ int) string {
		return "Hello"
	})
	result2 := Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "3", "4"})
}

func TestFlatMap(t *testing.T) {
	is := assert.New(t)

	result1 := FlatMap[int, string]([]int{0, 1, 2, 3, 4}, func(x int, _ int) []string {
		return []string{"Hello"}
	})
	result2 := FlatMap[int64, string]([]int64{0, 1, 2, 3, 4}, func(x int64, _ int) []string {
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

func TestTimes(t *testing.T) {
	is := assert.New(t)

	result1 := Times[string](3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []string{"0", "1", "2"})
}

func TestReduce(t *testing.T) {
	is := assert.New(t)

	result1 := Reduce[int, int]([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	result2 := Reduce[int, int]([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 10)

	is.Equal(result1, 10)
	is.Equal(result2, 20)
}

func TestUniq(t *testing.T) {
	is := assert.New(t)

	result1 := Uniq[int]([]int{1, 2, 2, 1})

	is.Equal(len(result1), 2)
	is.Equal(result1, []int{1, 2})
}

func TestUniqBy(t *testing.T) {
	is := assert.New(t)

	result1 := UniqBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []int{0, 1, 2})
}

func TestGroupBy(t *testing.T) {
	is := assert.New(t)

	result1 := GroupBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, map[int][]int{
		0: []int{0, 3},
		1: []int{1, 4},
		2: []int{2, 5},
	})
}

func TestChunk(t *testing.T) {
	is := assert.New(t)

	result1 := Chunk[int]([]int{0, 1, 2, 3, 4, 5}, 2)
	result2 := Chunk[int]([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	result3 := Chunk[int]([]int{}, 2)
	result4 := Chunk[int]([]int{0}, 2)

	is.Equal(result1, [][]int{{0, 1}, {2, 3}, {4, 5}})
	is.Equal(result2, [][]int{{0, 1}, {2, 3}, {4, 5}, {6}})
	is.Equal(result3, [][]int{})
	is.Equal(result4, [][]int{{0}})
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

	result1 := PartitionBy[int, string]([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)
	result2 := PartitionBy[int, string]([]int{}, oddEven)

	is.Equal(result1, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Equal(result2, [][]int{})
}

func TestFlatten(t *testing.T) {
	is := assert.New(t)

	result1 := Flatten[int]([][]int{{0, 1}, {2, 3, 4, 5}})

	is.Equal(result1, []int{0, 1, 2, 3, 4, 5})
}

func TestShuffle(t *testing.T) {
	is := assert.New(t)

	result1 := Shuffle[int]([]int{0, 1, 2, 3, 4, 5})
	result2 := Shuffle[int]([]int{})

	is.NotEqual(result1, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result2, []int{})
}

func TestReverse(t *testing.T) {
	is := assert.New(t)

	result1 := Reverse[int]([]int{0, 1, 2, 3, 4, 5})
	result2 := Reverse[int]([]int{0, 1, 2, 3, 4, 5, 6})
	result3 := Reverse[int]([]int{})

	is.Equal(result1, []int{5, 4, 3, 2, 1, 0})
	is.Equal(result2, []int{6, 5, 4, 3, 2, 1, 0})
	is.Equal(result3, []int{})
}

func TestFill(t *testing.T) {
	is := assert.New(t)

	result1 := Fill[foo]([]foo{foo{"a"}, foo{"a"}}, foo{"b"})
	result2 := Fill[foo]([]foo{}, foo{"a"})

	is.Equal(result1, []foo{foo{"b"}, foo{"b"}})
	is.Equal(result2, []foo{})
}

func TestRepeat(t *testing.T) {
	is := assert.New(t)

	result1 := Repeat[foo](2, foo{"a"})
	result2 := Repeat[foo](0, foo{"a"})

	is.Equal(result1, []foo{foo{"a"}, foo{"a"}})
	is.Equal(result2, []foo{})
}

func TestToMap(t *testing.T) {
	is := assert.New(t)

	result1 := ToMap[int, string]([]string{"a", "aa", "aaa"}, func(str string) int {
		return len(str)
	})

	is.Equal(result1, map[int]string{1: "a", 2: "aa", 3: "aaa"})
}
