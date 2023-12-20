package lo_test

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := lo.Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{2, 4})

	r2 := lo.Filter([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})

	is.Equal(r2, []string{"foo", "bar"})
}

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Map([]int{1, 2, 3, 4}, func(x int, _ int) string {
		return "Hello"
	})
	result2 := lo.Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
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

	r1 := lo.FilterMap([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), true
		}
		return "", false
	})
	r2 := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
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

	result1 := lo.FlatMap([]int{0, 1, 2, 3, 4}, func(x int, _ int) []string {
		return []string{"Hello"}
	})
	result2 := lo.FlatMap([]int64{0, 1, 2, 3, 4}, func(x int64, _ int) []string {
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
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []string{"0", "1", "2"})
}

func TestReduce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	result2 := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 10)

	is.Equal(result1, 10)
	is.Equal(result2, 20)
}

func TestReduceRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, _ int) []int {
		return append(agg, item...)
	}, []int{})

	is.Equal(result1, []int{4, 5, 2, 3, 0, 1})
}

func TestForEach(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	lo.ForEach([]string{"a", "b", "c"}, func(item string, i int) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.ElementsMatch([]string{"a", "b", "c"}, callParams1)
	is.ElementsMatch([]int{0, 1, 2}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Uniq([]int{1, 2, 2, 1})

	is.Equal(len(result1), 2)
	is.Equal(result1, []int{1, 2})
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.UniqBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []int{0, 1, 2})
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
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
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
	result2 := lo.Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	result3 := lo.Chunk([]int{}, 2)
	result4 := lo.Chunk([]int{0}, 2)

	is.Equal(result1, [][]int{{0, 1}, {2, 3}, {4, 5}})
	is.Equal(result2, [][]int{{0, 1}, {2, 3}, {4, 5}, {6}})
	is.Equal(result3, [][]int{})
	is.Equal(result4, [][]int{{0}})
	is.PanicsWithValue("Second parameter must be greater than 0", func() {
		lo.Chunk([]int{0}, 0)
	})
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

	result1 := lo.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)
	result2 := lo.PartitionBy([]int{}, oddEven)

	is.Equal(result1, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Equal(result2, [][]int{})
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Flatten([][]int{{0, 1}, {2, 3, 4, 5}})

	is.Equal(result1, []int{0, 1, 2, 3, 4, 5})
}

func TestInterleave(t *testing.T) {
	tests := []struct {
		name        string
		collections [][]int
		want        []int
	}{
		{
			"nil",
			[][]int{nil},
			[]int{},
		},
		{
			"empty",
			[][]int{},
			[]int{},
		},
		{
			"empties",
			[][]int{{}, {}},
			[]int{},
		},
		{
			"same length",
			[][]int{{1, 3, 5}, {2, 4, 6}},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"different length",
			[][]int{{1, 3, 5, 6}, {2, 4}},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"many slices",
			[][]int{{1}, {2, 5, 8}, {3, 6}, {4, 7, 9, 10}},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lo.Interleave(tt.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interleave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result2 := lo.Shuffle([]int{})

	is.NotEqual(result1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	is.Equal(result2, []int{})
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Reverse([]int{0, 1, 2, 3, 4, 5})
	result2 := lo.Reverse([]int{0, 1, 2, 3, 4, 5, 6})
	result3 := lo.Reverse([]int{})

	is.Equal(result1, []int{5, 4, 3, 2, 1, 0})
	is.Equal(result2, []int{6, 5, 4, 3, 2, 1, 0})
	is.Equal(result3, []int{})
}

func TestFill(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Fill([]foo{{"a"}, {"a"}}, foo{"b"})
	result2 := lo.Fill([]foo{}, foo{"a"})

	is.Equal(result1, []foo{{"b"}, {"b"}})
	is.Equal(result2, []foo{})
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Repeat(2, foo{"a"})
	result2 := lo.Repeat(0, foo{"a"})

	is.Equal(result1, []foo{{"a"}, {"a"}})
	is.Equal(result2, []foo{})
}

func TestRepeatBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	cb := func(i int) int {
		return int(math.Pow(float64(i), 2))
	}

	result1 := lo.RepeatBy(0, cb)
	result2 := lo.RepeatBy(2, cb)
	result3 := lo.RepeatBy(5, cb)

	is.Equal([]int{}, result1)
	is.Equal([]int{0, 1}, result2)
	is.Equal([]int{0, 1, 4, 9, 16}, result3)
}

func TestKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.KeyBy([]string{"a", "aa", "aaa"}, func(str string) int {
		return len(str)
	})

	is.Equal(result1, map[int]string{1: "a", 2: "aa", 3: "aaa"})
}

func TestAssociate(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int) {
		return f.baz, f.bar
	}
	testCases := []struct {
		in     []*foo
		expect map[string]int
	}{
		{
			in:     []*foo{{baz: "apple", bar: 1}},
			expect: map[string]int{"apple": 1},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			expect: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			expect: map[string]int{"apple": 2},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			is := assert.New(t)
			is.Equal(lo.Associate(testCase.in, transform), testCase.expect)
		})
	}
}

func TestSliceToMap(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int) {
		return f.baz, f.bar
	}
	testCases := []struct {
		in     []*foo
		expect map[string]int
	}{
		{
			in:     []*foo{{baz: "apple", bar: 1}},
			expect: map[string]int{"apple": 1},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			expect: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			expect: map[string]int{"apple": 2},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			is := assert.New(t)
			is.Equal(lo.SliceToMap(testCase.in, transform), testCase.expect)
		})
	}
}

func TestDrop(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{1, 2, 3, 4}, lo.Drop([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{2, 3, 4}, lo.Drop([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{3, 4}, lo.Drop([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{4}, lo.Drop([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, lo.Drop([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, lo.Drop([]int{0, 1, 2, 3, 4}, 6))
}

func TestDropRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, lo.DropRight([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{0, 1, 2}, lo.DropRight([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{0, 1}, lo.DropRight([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{0}, lo.DropRight([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, lo.DropRight([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, lo.DropRight([]int{0, 1, 2, 3, 4}, 6))
}

func TestDropWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{4, 5, 6}, lo.DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 4
	}))

	is.Equal([]int{}, lo.DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return true
	}))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, lo.DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t == 10
	}))
}

func TestDropRightWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, lo.DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 3
	}))

	is.Equal([]int{0, 1}, lo.DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 1
	}))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, lo.DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t == 10
	}))

	is.Equal([]int{}, lo.DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 10
	}))
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := lo.Reject([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{1, 3})

	r2 := lo.Reject([]string{"Smith", "foo", "Domin", "bar", "Olivia"}, func(x string, _ int) bool {
		return len(x) > 3
	})

	is.Equal(r2, []string{"foo", "bar"})
}

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	count1 := lo.Count([]int{1, 2, 1}, 1)
	count2 := lo.Count([]int{1, 2, 1}, 3)
	count3 := lo.Count([]int{}, 1)

	is.Equal(count1, 2)
	is.Equal(count2, 0)
	is.Equal(count3, 0)
}

func TestCountBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	count1 := lo.CountBy([]int{1, 2, 1}, func(i int) bool {
		return i < 2
	})

	count2 := lo.CountBy([]int{1, 2, 1}, func(i int) bool {
		return i > 2
	})

	count3 := lo.CountBy([]int{}, func(i int) bool {
		return i <= 2
	})

	is.Equal(count1, 2)
	is.Equal(count2, 0)
	is.Equal(count3, 0)
}

func TestCountValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(map[int]int{}, lo.CountValues([]int{}))
	is.Equal(map[int]int{1: 1, 2: 1}, lo.CountValues([]int{1, 2}))
	is.Equal(map[int]int{1: 1, 2: 2}, lo.CountValues([]int{1, 2, 2}))
	is.Equal(map[string]int{"": 1, "foo": 1, "bar": 1}, lo.CountValues([]string{"foo", "bar", ""}))
	is.Equal(map[string]int{"foo": 1, "bar": 2}, lo.CountValues([]string{"foo", "bar", "bar"}))
}

func TestCountValuesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	oddEven := func(v int) bool {
		return v%2 == 0
	}
	length := func(v string) int {
		return len(v)
	}

	result1 := lo.CountValuesBy([]int{}, oddEven)
	result2 := lo.CountValuesBy([]int{1, 2}, oddEven)
	result3 := lo.CountValuesBy([]int{1, 2, 2}, oddEven)
	result4 := lo.CountValuesBy([]string{"foo", "bar", ""}, length)
	result5 := lo.CountValuesBy([]string{"foo", "bar", "bar"}, length)

	is.Equal(map[bool]int{}, result1)
	is.Equal(map[bool]int{true: 1, false: 1}, result2)
	is.Equal(map[bool]int{true: 2, false: 1}, result3)
	is.Equal(map[int]int{0: 1, 3: 2}, result4)
	is.Equal(map[int]int{3: 3}, result5)
}

func TestSubset(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 2, 3, 4}

	out1 := lo.Subset(in, 0, 0)
	out2 := lo.Subset(in, 10, 2)
	out3 := lo.Subset(in, -10, 2)
	out4 := lo.Subset(in, 0, 10)
	out5 := lo.Subset(in, 0, 2)
	out6 := lo.Subset(in, 2, 2)
	out7 := lo.Subset(in, 2, 5)
	out8 := lo.Subset(in, 2, 3)
	out9 := lo.Subset(in, 2, 4)
	out10 := lo.Subset(in, -2, 4)
	out11 := lo.Subset(in, -4, 1)
	out12 := lo.Subset(in, -4, math.MaxUint)

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

func TestSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 2, 3, 4}

	out1 := lo.Slice(in, 0, 0)
	out2 := lo.Slice(in, 0, 1)
	out3 := lo.Slice(in, 0, 5)
	out4 := lo.Slice(in, 0, 6)
	out5 := lo.Slice(in, 1, 1)
	out6 := lo.Slice(in, 1, 5)
	out7 := lo.Slice(in, 1, 6)
	out8 := lo.Slice(in, 4, 5)
	out9 := lo.Slice(in, 5, 5)
	out10 := lo.Slice(in, 6, 5)
	out11 := lo.Slice(in, 6, 6)
	out12 := lo.Slice(in, 1, 0)
	out13 := lo.Slice(in, 5, 0)
	out14 := lo.Slice(in, 6, 4)
	out15 := lo.Slice(in, 6, 7)
	out16 := lo.Slice(in, -10, 1)
	out17 := lo.Slice(in, -1, 3)
	out18 := lo.Slice(in, -10, 7)

	is.Equal([]int{}, out1)
	is.Equal([]int{0}, out2)
	is.Equal([]int{0, 1, 2, 3, 4}, out3)
	is.Equal([]int{0, 1, 2, 3, 4}, out4)
	is.Equal([]int{}, out5)
	is.Equal([]int{1, 2, 3, 4}, out6)
	is.Equal([]int{1, 2, 3, 4}, out7)
	is.Equal([]int{4}, out8)
	is.Equal([]int{}, out9)
	is.Equal([]int{}, out10)
	is.Equal([]int{}, out11)
	is.Equal([]int{}, out12)
	is.Equal([]int{}, out13)
	is.Equal([]int{}, out14)
	is.Equal([]int{}, out15)
	is.Equal([]int{0}, out16)
	is.Equal([]int{0, 1, 2}, out17)
	is.Equal([]int{0, 1, 2, 3, 4}, out18)
}

func TestReplace(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := lo.Replace(in, 0, 42, 2)
	out2 := lo.Replace(in, 0, 42, 1)
	out3 := lo.Replace(in, 0, 42, 0)
	out4 := lo.Replace(in, 0, 42, -1)
	out5 := lo.Replace(in, 0, 42, -1)
	out6 := lo.Replace(in, -1, 42, 2)
	out7 := lo.Replace(in, -1, 42, 1)
	out8 := lo.Replace(in, -1, 42, 0)
	out9 := lo.Replace(in, -1, 42, -1)
	out10 := lo.Replace(in, -1, 42, -1)

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
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := lo.ReplaceAll(in, 0, 42)
	out2 := lo.ReplaceAll(in, -1, 42)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out1)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out2)
}

func TestCompact(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := lo.Compact([]int{2, 0, 4, 0})

	is.Equal(r1, []int{2, 4})

	r2 := lo.Compact([]string{"", "foo", "", "bar", ""})

	is.Equal(r2, []string{"foo", "bar"})

	r3 := lo.Compact([]bool{true, false, true, false})

	is.Equal(r3, []bool{true, true})

	type foo struct {
		bar int
		baz string
	}

	// slice of structs
	// If all fields of an element are zero values, Compact removes it.

	r4 := lo.Compact([]foo{
		{bar: 1, baz: "a"}, // all fields are non-zero values
		{bar: 0, baz: ""},  // all fields are zero values
		{bar: 2, baz: ""},  // bar is non-zero
	})

	is.Equal(r4, []foo{{bar: 1, baz: "a"}, {bar: 2, baz: ""}})

	// slice of pointers to structs
	// If an element is nil, Compact removes it.

	e1, e2, e3 := foo{bar: 1, baz: "a"}, foo{bar: 0, baz: ""}, foo{bar: 2, baz: ""}
	// NOTE: e2 is a zero value of foo, but its pointer &e2 is not a zero value of *foo.
	r5 := lo.Compact([]*foo{&e1, &e2, nil, &e3})

	is.Equal(r5, []*foo{&e1, &e2, &e3})
}

func TestIsSorted(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(lo.IsSorted([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	is.True(lo.IsSorted([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}))

	is.False(lo.IsSorted([]int{0, 1, 4, 3, 2, 5, 6, 7, 8, 9, 10}))
	is.False(lo.IsSorted([]string{"a", "b", "d", "c", "e", "f", "g", "h", "i", "j"}))
}

func TestIsSortedByKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(lo.IsSortedByKey([]string{"a", "bb", "ccc"}, func(s string) int {
		return len(s)
	}))

	is.False(lo.IsSortedByKey([]string{"aa", "b", "ccc"}, func(s string) int {
		return len(s)
	}))

	is.True(lo.IsSortedByKey([]string{"1", "2", "3", "11"}, func(s string) int {
		ret, _ := strconv.Atoi(s)
		return ret
	}))
}
