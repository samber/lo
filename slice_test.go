package lo

import (
	"math"
	"strconv"
	"strings"
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

	r1 := Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{2, 4})

	r2 := Filter([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})

	is.Equal(r2, []string{"foo", "bar"})
}

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

	r1 := FilterMap([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, bool) {
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
	is.Equal(r1, []string{"2", "4"})
	is.Equal(r2, []string{"xpu", "xpu"})
}

func TestFlatMap(t *testing.T) {
	is := assert.New(t)

	result1 := FlatMap([]int{0, 1, 2, 3, 4}, func(x int, _ int) []string {
		return []string{"Hello"}
	})
	result2 := FlatMap([]int64{0, 1, 2, 3, 4}, func(x int64, _ int) []string {
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

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []string{"0", "1", "2"})
}

func TestReduce(t *testing.T) {
	is := assert.New(t)

	result1 := Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	result2 := Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 10)

	is.Equal(result1, 10)
	is.Equal(result2, 20)
}

func TestForEach(t *testing.T) {
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	ForEach([]string{"a", "b", "c"}, func(item string, i int) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.ElementsMatch([]string{"a", "b", "c"}, callParams1)
	is.ElementsMatch([]int{0, 1, 2}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestUniq(t *testing.T) {
	is := assert.New(t)

	result1 := Uniq([]int{1, 2, 2, 1})

	is.Equal(len(result1), 2)
	is.Equal(result1, []int{1, 2})
}

func TestUniqBy(t *testing.T) {
	is := assert.New(t)

	result1 := UniqBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []int{0, 1, 2})
}

func TestGroupBy(t *testing.T) {
	is := assert.New(t)

	result1 := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	})
}

func TestChunk(t *testing.T) {
	is := assert.New(t)

	result1 := Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
	result2 := Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	result3 := Chunk([]int{}, 2)
	result4 := Chunk([]int{0}, 2)

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

	result1 := PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)
	result2 := PartitionBy([]int{}, oddEven)

	is.Equal(result1, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Equal(result2, [][]int{})
}

func TestFlatten(t *testing.T) {
	is := assert.New(t)

	result1 := Flatten([][]int{{0, 1}, {2, 3, 4, 5}})

	is.Equal(result1, []int{0, 1, 2, 3, 4, 5})
}

func TestShuffle(t *testing.T) {
	is := assert.New(t)

	result1 := Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result2 := Shuffle([]int{})

	is.NotEqual(result1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	is.Equal(result2, []int{})
}

func TestReverse(t *testing.T) {
	is := assert.New(t)

	result1 := Reverse([]int{0, 1, 2, 3, 4, 5})
	result2 := Reverse([]int{0, 1, 2, 3, 4, 5, 6})
	result3 := Reverse([]int{})

	is.Equal(result1, []int{5, 4, 3, 2, 1, 0})
	is.Equal(result2, []int{6, 5, 4, 3, 2, 1, 0})
	is.Equal(result3, []int{})
}

func TestFill(t *testing.T) {
	is := assert.New(t)

	result1 := Fill([]foo{{"a"}, {"a"}}, foo{"b"})
	result2 := Fill([]foo{}, foo{"a"})

	is.Equal(result1, []foo{{"b"}, {"b"}})
	is.Equal(result2, []foo{})
}

func TestRepeat(t *testing.T) {
	is := assert.New(t)

	result1 := Repeat(2, foo{"a"})
	result2 := Repeat(0, foo{"a"})

	is.Equal(result1, []foo{{"a"}, {"a"}})
	is.Equal(result2, []foo{})
}

func TestRepeatBy(t *testing.T) {
	is := assert.New(t)

	cb := func(i int) int {
		return int(math.Pow(float64(i), 2))
	}

	result1 := RepeatBy(0, cb)
	result2 := RepeatBy(2, cb)
	result3 := RepeatBy(5, cb)

	is.Equal([]int{}, result1)
	is.Equal([]int{0, 1}, result2)
	is.Equal([]int{0, 1, 4, 9, 16}, result3)
}

func TestKeyBy(t *testing.T) {
	is := assert.New(t)

	result1 := KeyBy([]string{"a", "aa", "aaa"}, func(str string) int {
		return len(str)
	})

	is.Equal(result1, map[int]string{1: "a", 2: "aa", 3: "aaa"})
}

func TestDrop(t *testing.T) {
	is := assert.New(t)

	is.Equal([]int{1, 2, 3, 4}, Drop([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{2, 3, 4}, Drop([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{3, 4}, Drop([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{4}, Drop([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, Drop([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, Drop([]int{0, 1, 2, 3, 4}, 6))
}

func TestDropRight(t *testing.T) {
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, DropRight([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{0, 1, 2}, DropRight([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{0, 1}, DropRight([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{0}, DropRight([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, DropRight([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, DropRight([]int{0, 1, 2, 3, 4}, 6))
}

func TestDropWhile(t *testing.T) {
	is := assert.New(t)

	is.Equal([]int{4, 5, 6}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 4
	}))

	is.Equal([]int{}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return true
	}))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t == 10
	}))
}

func TestDropRightWhile(t *testing.T) {
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 3
	}))

	is.Equal([]int{0, 1}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 1
	}))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t == 10
	}))

	is.Equal([]int{}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 10
	}))
}

func TestReject(t *testing.T) {
	is := assert.New(t)

	r1 := Reject([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{1, 3})

	r2 := Reject([]string{"Smith", "foo", "Domin", "bar", "Olivia"}, func(x string, _ int) bool {
		return len(x) > 3
	})

	is.Equal(r2, []string{"foo", "bar"})
}

func TestCount(t *testing.T) {
	is := assert.New(t)

	count1 := Count([]int{1, 2, 1}, 1)
	count2 := Count([]int{1, 2, 1}, 3)
	count3 := Count([]int{}, 1)

	is.Equal(count1, 2)
	is.Equal(count2, 0)
	is.Equal(count3, 0)
}

func TestCountBy(t *testing.T) {
	is := assert.New(t)

	count1 := CountBy([]int{1, 2, 1}, func(i int) bool {
		return i < 2
	})

	count2 := CountBy([]int{1, 2, 1}, func(i int) bool {
		return i > 2
	})

	count3 := CountBy([]int{}, func(i int) bool {
		return i <= 2
	})

	is.Equal(count1, 2)
	is.Equal(count2, 0)
	is.Equal(count3, 0)
}

func TestSubset(t *testing.T) {
	is := assert.New(t)

	in := []int{0, 1, 2, 3, 4}

	out1 := Subset(in, 0, 0)
	out2 := Subset(in, 10, 2)
	out3 := Subset(in, -10, 2)
	out4 := Subset(in, 0, 10)
	out5 := Subset(in, 0, 2)
	out6 := Subset(in, 2, 2)
	out7 := Subset(in, 2, 5)
	out8 := Subset(in, 2, 3)
	out9 := Subset(in, 2, 4)
	out10 := Subset(in, -2, 4)
	out11 := Subset(in, -4, 1)
	out12 := Subset(in, -4, math.MaxUint)

	is.Equal([]int{}, out1)
	is.Equal([]int{}, out2)
	is.Equal([]int{0, 1}, out3)
	is.Equal([]int{0, 1, 2, 3, 4}, out4)
	is.Equal([]int{0, 1}, out5)
	is.Equal([]int{2, 3}, out6)
	is.Equal([]int{2, 3, 4}, out7)
	is.Equal([]int{2, 3, 4}, out8)
	is.Equal([]int{2, 3, 4}, out9)
	is.Equal([]int{3, 4}, out10)
	is.Equal([]int{1}, out11)
	is.Equal([]int{1, 2, 3, 4}, out12)
}

func TestReplace(t *testing.T) {
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := Replace(in, 0, 42, 2)
	out2 := Replace(in, 0, 42, 1)
	out3 := Replace(in, 0, 42, 0)
	out4 := Replace(in, 0, 42, -1)
	out5 := Replace(in, 0, 42, -1)
	out6 := Replace(in, -1, 42, 2)
	out7 := Replace(in, -1, 42, 1)
	out8 := Replace(in, -1, 42, 0)
	out9 := Replace(in, -1, 42, -1)
	out10 := Replace(in, -1, 42, -1)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 0}, out1)
	is.Equal([]int{42, 1, 0, 1, 2, 3, 0}, out2)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out3)
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out4)
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out5)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out6)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out7)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out8)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out9)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out10)
}

func TestReplaceAll(t *testing.T) {
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := ReplaceAll(in, 0, 42)
	out2 := ReplaceAll(in, -1, 42)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out1)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out2)
}
