package lo

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	is.Equal(r1, []int{2, 4})

	r2 := Filter([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.Equal(r2, []string{"foo", "bar"})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Filter(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestMap(t *testing.T) {
	t.Parallel()
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

func TestUniqMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type User struct {
		Name string
		age  int
	}

	users := []User{{Name: "Alice", age: 20}, {Name: "Alex", age: 21}, {Name: "Alex", age: 22}}
	result := UniqMap(users, func(item User, index int) string {
		return item.Name
	})

	is.Equal(result, []string{"Alice", "Alex"})
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []string{"0", "1", "2"})
}

func TestReduce(t *testing.T) {
	t.Parallel()
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

func TestReduceRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, _ int) []int {
		return append(agg, item...)
	}, []int{})

	is.Equal(result1, []int{4, 5, 2, 3, 0, 1})

	type collection []int
	result3 := ReduceRight(collection{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 10)
	is.Equal(result3, 20)
}

func TestForEach(t *testing.T) {
	t.Parallel()
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

func TestForEachWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	var callParams1 []string
	var callParams2 []int

	ForEachWhile([]string{"a", "b", "c"}, func(item string, i int) bool {
		if item == "c" {
			return false
		}
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
		return true
	})

	is.ElementsMatch([]string{"a", "b"}, callParams1)
	is.ElementsMatch([]int{0, 1}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Uniq([]int{1, 2, 2, 1})

	is.Equal(len(result1), 2)
	is.Equal(result1, []int{1, 2})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Uniq(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := UniqBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []int{0, 1, 2})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := UniqBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
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

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := GroupBy(allStrings, func(i string) int {
		return 42
	})
	is.IsType(nonempty[42], allStrings, "type preserved")
}

func TestGroupByMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupByMap([]int{0, 1, 2, 3, 4, 5}, func(i int) (int, string) {
		return i % 3, strconv.Itoa(i)
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, map[int][]string{
		0: {"0", "3"},
		1: {"1", "4"},
		2: {"2", "5"},
	})

	type myInt int
	type myInts []myInt
	result2 := GroupByMap(myInts{1, 0, 2, 3, 4, 5}, func(i myInt) (int, string) {
		return int(i % 3), strconv.Itoa(int(i))
	})

	is.Equal(len(result2), 3)
	is.Equal(result2, map[int][]string{
		0: {"0", "3"},
		1: {"1", "4"},
		2: {"2", "5"},
	})

	type product struct {
		ID         int64
		CategoryID int64
	}
	products := []product{
		{ID: 1, CategoryID: 1},
		{ID: 2, CategoryID: 1},
		{ID: 3, CategoryID: 2},
		{ID: 4, CategoryID: 3},
		{ID: 5, CategoryID: 3},
	}
	result3 := GroupByMap(products, func(item product) (int64, string) {
		return item.CategoryID, "Product " + strconv.FormatInt(item.ID, 10)
	})

	is.Equal(len(result3), 3)
	is.Equal(result3, map[int64][]string{
		1: {"Product 1", "Product 2"},
		2: {"Product 3"},
		3: {"Product 4", "Product 5"},
	})
}

func TestChunk(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
	result2 := Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	result3 := Chunk([]int{}, 2)
	result4 := Chunk([]int{0}, 2)

	is.Equal(result1, [][]int{{0, 1}, {2, 3}, {4, 5}})
	is.Equal(result2, [][]int{{0, 1}, {2, 3}, {4, 5}, {6}})
	is.Equal(result3, [][]int{})
	is.Equal(result4, [][]int{{0}})
	is.PanicsWithValue("Second parameter must be greater than 0", func() {
		Chunk([]int{0}, 0)
	})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Chunk(allStrings, 2)
	is.IsType(nonempty[0], allStrings, "type preserved")

	// appending to a chunk should not affect original array
	originalArray := []int{0, 1, 2, 3, 4, 5}
	result5 := Chunk(originalArray, 2)
	result5[0] = append(result5[0], 6)
	is.Equal(originalArray, []int{0, 1, 2, 3, 4, 5})
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

	is.Equal(result1, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Equal(result2, [][]int{})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := PartitionBy(allStrings, func(item string) int {
		return len(item)
	})
	is.IsType(nonempty[0], allStrings, "type preserved")
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Flatten([][]int{{0, 1}, {2, 3, 4, 5}})

	is.Equal(result1, []int{0, 1, 2, 3, 4, 5})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Flatten([]myStrings{allStrings})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestInterleave(t *testing.T) {
	is := assert.New(t)

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
			if got := Interleave(tt.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interleave() = %v, want %v", got, tt.want)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Interleave(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result2 := Shuffle([]int{})

	is.NotEqual(result1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	is.Equal(result2, []int{})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Shuffle(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reverse([]int{0, 1, 2, 3, 4, 5})
	result2 := Reverse([]int{0, 1, 2, 3, 4, 5, 6})
	result3 := Reverse([]int{})

	is.Equal(result1, []int{5, 4, 3, 2, 1, 0})
	is.Equal(result2, []int{6, 5, 4, 3, 2, 1, 0})
	is.Equal(result3, []int{})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Reverse(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFill(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Fill([]foo{{"a"}, {"a"}}, foo{"b"})
	result2 := Fill([]foo{}, foo{"a"})

	is.Equal(result1, []foo{{"b"}, {"b"}})
	is.Equal(result2, []foo{})
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Repeat(2, foo{"a"})
	result2 := Repeat(0, foo{"a"})

	is.Equal(result1, []foo{{"a"}, {"a"}})
	is.Equal(result2, []foo{})
}

func TestRepeatBy(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	is := assert.New(t)

	result1 := KeyBy([]string{"a", "aa", "aaa"}, func(str string) int {
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
			is.Equal(Associate(testCase.in, transform), testCase.expect)
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
			is.Equal(SliceToMap(testCase.in, transform), testCase.expect)
		})
	}
}

func TestFilterSliceToMap(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int, bool) {
		return f.baz, f.bar, f.bar > 1
	}
	testCases := []struct {
		in     []*foo
		expect map[string]int
	}{
		{
			in:     []*foo{{baz: "apple", bar: 1}},
			expect: map[string]int{},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			expect: map[string]int{"banana": 2},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			expect: map[string]int{"apple": 2},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			is := assert.New(t)
			is.Equal(FilterSliceToMap(testCase.in, transform), testCase.expect)
		})
	}
}

func TestKeyify(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Keyify([]int{1, 2, 3, 4})
	result2 := Keyify([]int{1, 1, 1, 2})
	result3 := Keyify([]int{})
	is.Equal(result1, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}})
	is.Equal(result2, map[int]struct{}{1: {}, 2: {}})
	is.Equal(result3, map[int]struct{}{})
}

func TestDrop(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{1, 2, 3, 4}, Drop([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{2, 3, 4}, Drop([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{3, 4}, Drop([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{4}, Drop([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, Drop([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, Drop([]int{0, 1, 2, 3, 4}, 6))

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Drop(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, DropRight([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{0, 1, 2}, DropRight([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{0, 1}, DropRight([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{0}, DropRight([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, DropRight([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, DropRight([]int{0, 1, 2, 3, 4}, 6))

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := DropRight(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropWhile(t *testing.T) {
	t.Parallel()
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

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := DropWhile(allStrings, func(t string) bool {
		return t != "foo"
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropRightWhile(t *testing.T) {
	t.Parallel()
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

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := DropRightWhile(allStrings, func(t string) bool {
		return t != "foo"
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropByIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{1, 2, 3, 4}, DropByIndex([]int{0, 1, 2, 3, 4}, 0))
	is.Equal([]int{3, 4}, DropByIndex([]int{0, 1, 2, 3, 4}, 0, 1, 2))
	is.Equal([]int{0, 4}, DropByIndex([]int{0, 1, 2, 3, 4}, -4, -2, -3))
	is.Equal([]int{0, 2, 3, 4}, DropByIndex([]int{0, 1, 2, 3, 4}, -4, -4))
	is.Equal([]int{2, 4}, DropByIndex([]int{0, 1, 2, 3, 4}, 3, 1, 0))
	is.Equal([]int{0, 1, 3, 4}, DropByIndex([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{0, 1, 2, 3}, DropByIndex([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{0, 1, 2, 3, 4}, DropByIndex([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{0, 1, 2, 3, 4}, DropByIndex([]int{0, 1, 2, 3, 4}, 100))
	is.Equal([]int{0, 1, 2, 3}, DropByIndex([]int{0, 1, 2, 3, 4}, -1))
	is.Equal([]int{}, DropByIndex([]int{}, 0, 1))
	is.Equal([]int{}, DropByIndex([]int{42}, 0, 1))
	is.Equal([]int{}, DropByIndex([]int{42}, 1, 0))
	is.Equal([]int{}, DropByIndex([]int{}, 1))
	is.Equal([]int{}, DropByIndex([]int{1}, 0))
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Reject([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{1, 3})

	r2 := Reject([]string{"Smith", "foo", "Domin", "bar", "Olivia"}, func(x string, _ int) bool {
		return len(x) > 3
	})

	is.Equal(r2, []string{"foo", "bar"})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Reject(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestRejectMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := RejectMap([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), false
		}
		return "", true
	})
	r2 := RejectMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", false
		}
		return "", true
	})

	is.Equal(len(r1), 2)
	is.Equal(len(r2), 2)
	is.Equal(r1, []string{"2", "4"})
	is.Equal(r2, []string{"xpu", "xpu"})
}

func TestFilterReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	left1, right1 := FilterReject([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(left1, []int{2, 4})
	is.Equal(right1, []int{1, 3})

	left2, right2 := FilterReject([]string{"Smith", "foo", "Domin", "bar", "Olivia"}, func(x string, _ int) bool {
		return len(x) > 3
	})

	is.Equal(left2, []string{"Smith", "Domin", "Olivia"})
	is.Equal(right2, []string{"foo", "bar"})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	a, b := FilterReject(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.IsType(a, allStrings, "type preserved")
	is.IsType(b, allStrings, "type preserved")
}

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	count1 := Count([]int{1, 2, 1}, 1)
	count2 := Count([]int{1, 2, 1}, 3)
	count3 := Count([]int{}, 1)

	is.Equal(count1, 2)
	is.Equal(count2, 0)
	is.Equal(count3, 0)
}

func TestCountBy(t *testing.T) {
	t.Parallel()
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

func TestCountValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(map[int]int{}, CountValues([]int{}))
	is.Equal(map[int]int{1: 1, 2: 1}, CountValues([]int{1, 2}))
	is.Equal(map[int]int{1: 1, 2: 2}, CountValues([]int{1, 2, 2}))
	is.Equal(map[string]int{"": 1, "foo": 1, "bar": 1}, CountValues([]string{"foo", "bar", ""}))
	is.Equal(map[string]int{"foo": 1, "bar": 2}, CountValues([]string{"foo", "bar", "bar"}))
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

	result1 := CountValuesBy([]int{}, oddEven)
	result2 := CountValuesBy([]int{1, 2}, oddEven)
	result3 := CountValuesBy([]int{1, 2, 2}, oddEven)
	result4 := CountValuesBy([]string{"foo", "bar", ""}, length)
	result5 := CountValuesBy([]string{"foo", "bar", "bar"}, length)

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

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Subset(allStrings, 0, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 2, 3, 4}

	out1 := Slice(in, 0, 0)
	out2 := Slice(in, 0, 1)
	out3 := Slice(in, 0, 5)
	out4 := Slice(in, 0, 6)
	out5 := Slice(in, 1, 1)
	out6 := Slice(in, 1, 5)
	out7 := Slice(in, 1, 6)
	out8 := Slice(in, 4, 5)
	out9 := Slice(in, 5, 5)
	out10 := Slice(in, 6, 5)
	out11 := Slice(in, 6, 6)
	out12 := Slice(in, 1, 0)
	out13 := Slice(in, 5, 0)
	out14 := Slice(in, 6, 4)
	out15 := Slice(in, 6, 7)
	out16 := Slice(in, -10, 1)
	out17 := Slice(in, -1, 3)
	out18 := Slice(in, -10, 7)

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

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Slice(allStrings, 0, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplace(t *testing.T) {
	t.Parallel()
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

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Replace(allStrings, "0", "2", 1)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := ReplaceAll(in, 0, 42)
	out2 := ReplaceAll(in, -1, 42)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out1)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out2)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := ReplaceAll(allStrings, "0", "2")
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestClone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Test with int slice
	original1 := []int{1, 2, 3, 4, 5}
	result1 := Clone(original1)
	is.Equal(result1, []int{1, 2, 3, 4, 5})

	// Verify it's a different slice by checking that modifying one doesn't affect the other
	original1[0] = 99
	is.Equal(original1, []int{99, 2, 3, 4, 5})
	is.Equal(result1, []int{1, 2, 3, 4, 5})

	// Test with string slice
	original2 := []string{"a", "b", "c"}
	result2 := Clone(original2)
	is.Equal(result2, []string{"a", "b", "c"})

	// Test with empty slice
	original3 := []int{}
	result3 := Clone(original3)
	is.Equal(result3, []int{})
	is.Equal(len(result3), 0)

	// Test with nil slice
	var original4 []int
	result4 := Clone(original4)
	is.Nil(result4)

	// Verify shallow copy behavior - modifying clone doesn't affect original
	original5 := []int{1, 2, 3}
	result5 := Clone(original5)
	result5[0] = 99
	is.Equal(original5, []int{1, 2, 3}) // Original unchanged
	is.Equal(result5, []int{99, 2, 3})  // Clone changed
}

func TestCompact(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Compact([]int{2, 0, 4, 0})

	is.Equal(r1, []int{2, 4})

	r2 := Compact([]string{"", "foo", "", "bar", ""})

	is.Equal(r2, []string{"foo", "bar"})

	r3 := Compact([]bool{true, false, true, false})

	is.Equal(r3, []bool{true, true})

	type foo struct {
		bar int
		baz string
	}

	// slice of structs
	// If all fields of an element are zero values, Compact removes it.

	r4 := Compact([]foo{
		{bar: 1, baz: "a"}, // all fields are non-zero values
		{bar: 0, baz: ""},  // all fields are zero values
		{bar: 2, baz: ""},  // bar is non-zero
	})

	is.Equal(r4, []foo{{bar: 1, baz: "a"}, {bar: 2, baz: ""}})

	// slice of pointers to structs
	// If an element is nil, Compact removes it.

	e1, e2, e3 := foo{bar: 1, baz: "a"}, foo{bar: 0, baz: ""}, foo{bar: 2, baz: ""}
	// NOTE: e2 is a zero value of foo, but its pointer &e2 is not a zero value of *foo.
	r5 := Compact([]*foo{&e1, &e2, nil, &e3})

	is.Equal(r5, []*foo{&e1, &e2, &e3})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Compact(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestIsSorted(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(IsSorted([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	is.True(IsSorted([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}))

	is.False(IsSorted([]int{0, 1, 4, 3, 2, 5, 6, 7, 8, 9, 10}))
	is.False(IsSorted([]string{"a", "b", "d", "c", "e", "f", "g", "h", "i", "j"}))
}

func TestIsSortedByKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(IsSortedByKey([]string{"a", "bb", "ccc"}, func(s string) int {
		return len(s)
	}))

	is.False(IsSortedByKey([]string{"aa", "b", "ccc"}, func(s string) int {
		return len(s)
	}))

	is.True(IsSortedByKey([]string{"1", "2", "3", "11"}, func(s string) int {
		ret, _ := strconv.Atoi(s)
		return ret
	}))
}

func TestSplice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sample := []string{"a", "b", "c", "d", "e", "f", "g"}

	// normal case
	results := Splice(sample, 1, "1", "2")
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, sample)
	is.Equal([]string{"a", "1", "2", "b", "c", "d", "e", "f", "g"}, results)

	// check there is no side effect
	results = Splice(sample, 1)
	results[0] = "b"
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, sample)

	// positive overflow
	results = Splice(sample, 42, "1", "2")
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, sample)
	is.Equal(results, []string{"a", "b", "c", "d", "e", "f", "g", "1", "2"})

	// negative overflow
	results = Splice(sample, -42, "1", "2")
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, sample)
	is.Equal(results, []string{"1", "2", "a", "b", "c", "d", "e", "f", "g"})

	// backward
	results = Splice(sample, -2, "1", "2")
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, sample)
	is.Equal(results, []string{"a", "b", "c", "d", "e", "1", "2", "f", "g"})

	results = Splice(sample, -7, "1", "2")
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, sample)
	is.Equal(results, []string{"1", "2", "a", "b", "c", "d", "e", "f", "g"})

	// other
	is.Equal([]string{"1", "2"}, Splice([]string{}, 0, "1", "2"))
	is.Equal([]string{"1", "2"}, Splice([]string{}, 1, "1", "2"))
	is.Equal([]string{"1", "2"}, Splice([]string{}, -1, "1", "2"))
	is.Equal([]string{"1", "2", "0"}, Splice([]string{"0"}, 0, "1", "2"))
	is.Equal([]string{"0", "1", "2"}, Splice([]string{"0"}, 1, "1", "2"))
	is.Equal([]string{"1", "2", "0"}, Splice([]string{"0"}, -1, "1", "2"))

	// type preserved
	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Splice(allStrings, 1, "1", "2")
	is.IsType(nonempty, allStrings, "type preserved")
}
