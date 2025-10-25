//go:build go1.23

package it

import (
	"fmt"
	"iter"
	"math"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Length(values[int]())
	r2 := Length(values(1, 2, 3, 4))

	is.Zero(r1)
	is.Equal(4, r2)
}

func TestDrain(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var done bool
	list := iter.Seq[int](func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
		done = true
	})

	Drain(list)

	is.True(done)
}

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Filter(values(1, 2, 3, 4), func(x int) bool {
		return x%2 == 0
	})
	is.Equal([]int{2, 4}, slices.Collect(r1))

	r2 := Filter(values("", "foo", "", "bar", ""), func(x string) bool {
		return len(x) > 0
	})
	is.Equal([]string{"foo", "bar"}, slices.Collect(r2))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Filter(allStrings, func(x string) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFilterI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterI(values(1, 2, 3, 4), func(x, _ int) bool {
		return x%2 == 0
	})
	is.Equal([]int{2, 4}, slices.Collect(r1))

	r2 := FilterI(values("", "foo", "", "bar", ""), func(x string, _ int) bool {
		return len(x) > 0
	})
	is.Equal([]string{"foo", "bar"}, slices.Collect(r2))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := FilterI(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Map(values(1, 2, 3, 4), func(x int) string {
		return "Hello"
	})
	result2 := Map(values[int64](1, 2, 3, 4), func(x int64) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal([]string{"Hello", "Hello", "Hello", "Hello"}, slices.Collect(result1))
	is.Equal([]string{"1", "2", "3", "4"}, slices.Collect(result2))
}

func TestMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapI(values(1, 2, 3, 4), func(x, _ int) string {
		return "Hello"
	})
	result2 := MapI(values[int64](1, 2, 3, 4), func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal([]string{"Hello", "Hello", "Hello", "Hello"}, slices.Collect(result1))
	is.Equal([]string{"1", "2", "3", "4"}, slices.Collect(result2))
}

func TestUniqMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type User struct {
		Name string
		age  int
	}

	users := values(User{Name: "Alice", age: 20}, User{Name: "Alex", age: 21}, User{Name: "Alex", age: 22})
	result := UniqMap(users, func(item User) string {
		return item.Name
	})

	is.Equal([]string{"Alice", "Alex"}, slices.Collect(result))
}

func TestUniqMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type User struct {
		Name string
		age  int
	}

	users := values(User{Name: "Alice", age: 20}, User{Name: "Alex", age: 21}, User{Name: "Alex", age: 22})
	result := UniqMapI(users, func(item User, _ int) string {
		return item.Name
	})

	is.Equal([]string{"Alice", "Alex"}, slices.Collect(result))
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterMap(values[int64](1, 2, 3, 4), func(x int64) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), true
		}
		return "", false
	})
	r2 := FilterMap(values("cpu", "gpu", "mouse", "keyboard"), func(x string) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})

	is.Equal([]string{"2", "4"}, slices.Collect(r1))
	is.Equal([]string{"xpu", "xpu"}, slices.Collect(r2))
}

func TestFilterMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterMapI(values[int64](1, 2, 3, 4), func(x int64, _ int) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), true
		}
		return "", false
	})
	r2 := FilterMapI(values("cpu", "gpu", "mouse", "keyboard"), func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})

	is.Equal([]string{"2", "4"}, slices.Collect(r1))
	is.Equal([]string{"xpu", "xpu"}, slices.Collect(r2))
}

func TestFlatMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FlatMap(values(0, 1, 2, 3, 4), func(x int) iter.Seq[string] {
		return values("Hello")
	})
	result2 := FlatMap(values[int64](0, 1, 2, 3, 4), func(x int64) iter.Seq[string] {
		return func(yield func(string) bool) {
			for range x {
				if !yield(strconv.FormatInt(x, 10)) {
					return
				}
			}
		}
	})

	is.Equal([]string{"Hello", "Hello", "Hello", "Hello", "Hello"}, slices.Collect(result1))
	is.Equal([]string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"}, slices.Collect(result2))
}

func TestFlatMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FlatMapI(values(0, 1, 2, 3, 4), func(x, _ int) iter.Seq[string] {
		return values("Hello")
	})
	result2 := FlatMapI(values[int64](0, 1, 2, 3, 4), func(x int64, _ int) iter.Seq[string] {
		return func(yield func(string) bool) {
			for range x {
				if !yield(strconv.FormatInt(x, 10)) {
					return
				}
			}
		}
	})

	is.Equal([]string{"Hello", "Hello", "Hello", "Hello", "Hello"}, slices.Collect(result1))
	is.Equal([]string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"}, slices.Collect(result2))
}

func TestTimes(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	is.Equal([]string{"0", "1", "2"}, slices.Collect(result1))
}

func TestReduce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reduce(values(1, 2, 3, 4), func(agg, item int) int {
		return agg + item
	}, 0)
	result2 := Reduce(values(1, 2, 3, 4), func(agg, item int) int {
		return agg + item
	}, 10)

	is.Equal(10, result1)
	is.Equal(20, result2)
}

func TestReduceI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ReduceI(values(1, 2, 3, 4), func(agg, item, _ int) int {
		return agg + item
	}, 0)
	result2 := ReduceI(values(1, 2, 3, 4), func(agg, item, _ int) int {
		return agg + item
	}, 10)

	is.Equal(10, result1)
	is.Equal(20, result2)
}

func TestReduceLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ReduceLast(values([]int{0, 1}, []int{2, 3}, []int{4, 5}), func(agg, item []int) []int {
		return append(agg, item...)
	}, []int{})
	is.Equal([]int{4, 5, 2, 3, 0, 1}, result1)

	result2 := ReduceLast(values(1, 2, 3, 4), func(agg, item int) int {
		return agg + item
	}, 10)
	is.Equal(20, result2)
}

func TestReduceLastI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ReduceLastI(values([]int{0, 1}, []int{2, 3}, []int{4, 5}), func(agg, item []int, _ int) []int {
		return append(agg, item...)
	}, []int{})
	is.Equal([]int{4, 5, 2, 3, 0, 1}, result1)

	result2 := ReduceLastI(values(1, 2, 3, 4), func(agg, item, _ int) int {
		return agg + item
	}, 10)
	is.Equal(20, result2)
}

func TestForEachI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	ForEachI(values("a", "b", "c"), func(item string, i int) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.Equal([]string{"a", "b", "c"}, callParams1)
	is.Equal([]int{0, 1, 2}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestForEachWhileI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	var callParams1 []string
	var callParams2 []int

	ForEachWhileI(values("a", "b", "c"), func(item string, i int) bool {
		if item == "c" {
			return false
		}
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
		return true
	})

	is.Equal([]string{"a", "b"}, callParams1)
	is.Equal([]int{0, 1}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Uniq(values(1, 2, 2, 1))
	is.Equal([]int{1, 2}, slices.Collect(result1))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Uniq(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := UniqBy(values(0, 1, 2, 3, 4, 5), func(i int) int {
		return i % 3
	})
	is.Equal([]int{0, 1, 2}, slices.Collect(result1))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := UniqBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupBy(values(0, 1, 2, 3, 4, 5), func(i int) int {
		return i % 3
	})
	is.Equal(map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	}, result1)
}

func TestGroupByMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupByMap(values(0, 1, 2, 3, 4, 5), func(i int) (int, string) {
		return i % 3, strconv.Itoa(i)
	})
	is.Equal(map[int][]string{
		0: {"0", "3"},
		1: {"1", "4"},
		2: {"2", "5"},
	}, result1)

	type myInt int
	result2 := GroupByMap(values[myInt](1, 0, 2, 3, 4, 5), func(i myInt) (int, string) {
		return int(i % 3), strconv.Itoa(int(i))
	})
	is.Equal(map[int][]string{
		0: {"0", "3"},
		1: {"1", "4"},
		2: {"2", "5"},
	}, result2)

	type product struct {
		ID         int64
		CategoryID int64
	}
	products := values(
		product{ID: 1, CategoryID: 1},
		product{ID: 2, CategoryID: 1},
		product{ID: 3, CategoryID: 2},
		product{ID: 4, CategoryID: 3},
		product{ID: 5, CategoryID: 3},
	)
	result3 := GroupByMap(products, func(item product) (int64, string) {
		return item.CategoryID, "Product " + strconv.FormatInt(item.ID, 10)
	})
	is.Equal(map[int64][]string{
		1: {"Product 1", "Product 2"},
		2: {"Product 3"},
		3: {"Product 4", "Product 5"},
	}, result3)
}

func TestChunk(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Chunk(values(0, 1, 2, 3, 4, 5), 2)
	result2 := Chunk(values(0, 1, 2, 3, 4, 5, 6), 2)
	result3 := Chunk(values[int](), 2)
	result4 := Chunk(values(0), 2)

	is.Equal([][]int{{0, 1}, {2, 3}, {4, 5}}, slices.Collect(result1))
	is.Equal([][]int{{0, 1}, {2, 3}, {4, 5}, {6}}, slices.Collect(result2))
	is.Empty(slices.Collect(result3))
	is.Equal([][]int{{0}}, slices.Collect(result4))
	is.PanicsWithValue("it.Chunk: size must be greater than 0", func() {
		Chunk(values(0), 0)
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

	result1 := PartitionBy(values(-2, -1, 0, 1, 2, 3, 4, 5), oddEven)
	result2 := PartitionBy(values[int](), oddEven)

	is.Equal([][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}}, result1)
	is.Empty(result2)
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Flatten([]iter.Seq[int]{values(0, 1), values(2, 3, 4, 5)})

	is.Equal([]int{0, 1, 2, 3, 4, 5}, slices.Collect(result1))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Flatten([]myStrings{allStrings})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestConcat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Concat(values(0, 1), values(2, 3, 4, 5))
	is.Equal([]int{0, 1, 2, 3, 4, 5}, slices.Collect(result1))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Concat(allStrings, allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestInterleave(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		in   []iter.Seq[int]
		want []int
	}{
		{
			"empty",
			[]iter.Seq[int]{},
			nil,
		},
		{
			"empties",
			[]iter.Seq[int]{values[int](), values[int]()},
			nil,
		},
		{
			"same length",
			[]iter.Seq[int]{values(1, 3, 5), values(2, 4, 6)},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"different length",
			[]iter.Seq[int]{values(1, 3, 5, 6), values(2, 4)},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"many sequences",
			[]iter.Seq[int]{values(1), values(2, 5, 8), values(3, 6), values(4, 7, 9, 10)},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, slices.Collect(Interleave(tc.in...)))
		})
	}
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Shuffle(values(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	result2 := Shuffle(values[int]())

	slice1 := slices.Collect(result1)
	is.NotEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, slice1)
	is.ElementsMatch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, slice1)
	is.Empty(slices.Collect(result2))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Shuffle(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reverse(values(0, 1, 2, 3, 4, 5))
	result2 := Reverse(values(0, 1, 2, 3, 4, 5, 6))
	result3 := Reverse(values[int]())

	is.Equal([]int{5, 4, 3, 2, 1, 0}, slices.Collect(result1))
	is.Equal([]int{6, 5, 4, 3, 2, 1, 0}, slices.Collect(result2))
	is.Empty(slices.Collect(result3))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Reverse(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFill(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Fill(values(foo{"a"}, foo{"a"}), foo{"b"})
	result2 := Fill(values[foo](), foo{"a"})

	is.Equal([]foo{{"b"}, {"b"}}, slices.Collect(result1))
	is.Empty(slices.Collect(result2))
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Repeat(2, foo{"a"})
	result2 := Repeat(0, foo{"a"})

	is.Equal([]foo{{"a"}, {"a"}}, slices.Collect(result1))
	is.Empty(slices.Collect(result2))
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

	is.Empty(slices.Collect(result1))
	is.Equal([]int{0, 1}, slices.Collect(result2))
	is.Equal([]int{0, 1, 4, 9, 16}, slices.Collect(result3))
}

func TestKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := KeyBy(values("a", "aa", "aaa"), func(str string) int {
		return len(str)
	})

	is.Equal(map[int]string{1: "a", 2: "aa", 3: "aaa"}, result1)
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
		in   []*foo
		want map[string]int
	}{
		{
			in:   []*foo{{baz: "apple", bar: 1}},
			want: map[string]int{"apple": 1},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			want: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			want: map[string]int{"apple": 2},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, Associate(slices.Values(tc.in), transform))
		})
	}
}

func TestAssociateI(t *testing.T) {
	t.Parallel()

	transform := func(s string, i int) (int, string) {
		return i % 2, s
	}
	testCases := []struct {
		in   []string
		want map[int]string
	}{
		{
			in:   []string{"zero"},
			want: map[int]string{0: "zero"},
		},
		{
			in:   []string{"zero", "one"},
			want: map[int]string{0: "zero", 1: "one"},
		},
		{
			in:   []string{"two", "one", "zero"},
			want: map[int]string{0: "zero", 1: "one"},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, AssociateI(slices.Values(tc.in), transform))
		})
	}
}

func TestSeqToMap(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int) {
		return f.baz, f.bar
	}
	testCases := []struct {
		in   []*foo
		want map[string]int
	}{
		{
			in:   []*foo{{baz: "apple", bar: 1}},
			want: map[string]int{"apple": 1},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			want: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			want: map[string]int{"apple": 2},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, SeqToMap(slices.Values(tc.in), transform))
		})
	}
}

func TestSeqToMapI(t *testing.T) {
	t.Parallel()

	transform := func(s string, i int) (int, string) {
		return i % 2, s
	}
	testCases := []struct {
		in   []string
		want map[int]string
	}{
		{
			in:   []string{"zero"},
			want: map[int]string{0: "zero"},
		},
		{
			in:   []string{"zero", "one"},
			want: map[int]string{0: "zero", 1: "one"},
		},
		{
			in:   []string{"two", "one", "zero"},
			want: map[int]string{0: "zero", 1: "one"},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, SeqToMapI(slices.Values(tc.in), transform))
		})
	}
}

func TestFilterSeqToMap(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int, bool) {
		return f.baz, f.bar, f.bar > 1
	}
	testCases := []struct {
		in   []*foo
		want map[string]int
	}{
		{
			in:   []*foo{{baz: "apple", bar: 1}},
			want: map[string]int{},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			want: map[string]int{"banana": 2},
		},
		{
			in:   []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			want: map[string]int{"apple": 2},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, FilterSeqToMap(slices.Values(tc.in), transform))
		})
	}
}

func TestFilterSeqToMapI(t *testing.T) {
	t.Parallel()

	transform := func(s string, i int) (int, string, bool) {
		return i % 5, s, i%2 == 0
	}
	testCases := []struct {
		in   []string
		want map[int]string
	}{
		{
			in:   []string{"zero"},
			want: map[int]string{0: "zero"},
		},
		{
			in:   []string{"zero", "one", "two", "three", "four"},
			want: map[int]string{0: "zero", 2: "two", 4: "four"},
		},
		{
			in:   []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"},
			want: map[int]string{0: "ten", 1: "six", 2: "two", 3: "eight", 4: "four"},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, FilterSeqToMapI(slices.Values(tc.in), transform))
		})
	}
}

func TestKeyify(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Keyify(values(1, 2, 3, 4))
	result2 := Keyify(values(1, 1, 1, 2))
	result3 := Keyify(values[int]())
	is.Equal(map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}, result1)
	is.Equal(map[int]struct{}{1: {}, 2: {}}, result2)
	is.Empty(result3)
}

func TestDrop(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3, 4}, slices.Collect(Drop(values(0, 1, 2, 3, 4), 0)))
	is.Equal([]int{1, 2, 3, 4}, slices.Collect(Drop(values(0, 1, 2, 3, 4), 1)))
	is.Equal([]int{2, 3, 4}, slices.Collect(Drop(values(0, 1, 2, 3, 4), 2)))
	is.Equal([]int{3, 4}, slices.Collect(Drop(values(0, 1, 2, 3, 4), 3)))
	is.Equal([]int{4}, slices.Collect(Drop(values(0, 1, 2, 3, 4), 4)))
	is.Empty(slices.Collect(Drop(values(0, 1, 2, 3, 4), 5)))
	is.Empty(slices.Collect(Drop(values(0, 1, 2, 3, 4), 6)))

	is.PanicsWithValue("it.Drop: n must not be negative", func() {
		Drop(values(0, 1, 2, 3, 4), -1)
	})

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Drop(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3, 4}, slices.Collect(DropLast(values(0, 1, 2, 3, 4), 0)))
	is.Equal([]int{0, 1, 2, 3}, slices.Collect(DropLast(values(0, 1, 2, 3, 4), 1)))
	is.Equal([]int{0, 1, 2}, slices.Collect(DropLast(values(0, 1, 2, 3, 4), 2)))
	is.Equal([]int{0, 1}, slices.Collect(DropLast(values(0, 1, 2, 3, 4), 3)))
	is.Equal([]int{0}, slices.Collect(DropLast(values(0, 1, 2, 3, 4), 4)))
	is.Empty(slices.Collect(DropLast(values(0, 1, 2, 3, 4), 5)))
	is.Empty(slices.Collect(DropLast(values(0, 1, 2, 3, 4), 6)))

	is.PanicsWithValue("it.DropLast: n must not be negative", func() {
		DropLast(values(0, 1, 2, 3, 4), -1)
	})

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := DropLast(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{4, 5, 6}, slices.Collect(DropWhile(values(0, 1, 2, 3, 4, 5, 6), func(t int) bool {
		return t != 4
	})))

	is.Empty(slices.Collect(DropWhile(values(0, 1, 2, 3, 4, 5, 6), func(t int) bool {
		return true
	})))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, slices.Collect(DropWhile(values(0, 1, 2, 3, 4, 5, 6), func(t int) bool {
		return t == 10
	})))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := DropWhile(allStrings, func(t string) bool {
		return t != "foo"
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropLastWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, slices.Collect(DropLastWhile(values(0, 1, 2, 3, 4, 5, 6), func(t int) bool {
		return t != 3
	})))

	is.Equal([]int{0, 1}, slices.Collect(DropLastWhile(values(0, 1, 2, 3, 4, 5, 6), func(t int) bool {
		return t != 1
	})))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, slices.Collect(DropLastWhile(values(0, 1, 2, 3, 4, 5, 6), func(t int) bool {
		return t == 10
	})))

	is.Empty(slices.Collect(DropLastWhile(values(0, 1, 2, 3, 4, 5, 6), func(t int) bool {
		return t != 10
	})))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := DropLastWhile(allStrings, func(t string) bool {
		return t != "foo"
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDropByIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{1, 2, 3, 4}, slices.Collect(DropByIndex(values(0, 1, 2, 3, 4), 0)))
	is.Equal([]int{3, 4}, slices.Collect(DropByIndex(values(0, 1, 2, 3, 4), 0, 1, 2)))
	is.Equal([]int{2, 4}, slices.Collect(DropByIndex(values(0, 1, 2, 3, 4), 3, 1, 0)))
	is.Equal([]int{0, 1, 3, 4}, slices.Collect(DropByIndex(values(0, 1, 2, 3, 4), 2)))
	is.Equal([]int{0, 1, 2, 3}, slices.Collect(DropByIndex(values(0, 1, 2, 3, 4), 4)))
	is.Equal([]int{0, 1, 2, 3, 4}, slices.Collect(DropByIndex(values(0, 1, 2, 3, 4), 5)))
	is.Equal([]int{0, 1, 2, 3, 4}, slices.Collect(DropByIndex(values(0, 1, 2, 3, 4), 100)))
	is.Empty(slices.Collect(DropByIndex(values[int](), 0, 1)))
	is.Empty(slices.Collect(DropByIndex(values(42), 0, 1)))
	is.Empty(slices.Collect(DropByIndex(values(42), 1, 0)))
	is.Empty(slices.Collect(DropByIndex(values[int](), 1)))
	is.Empty(slices.Collect(DropByIndex(values(1), 0)))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := DropByIndex(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Reject(values(1, 2, 3, 4), func(x int) bool {
		return x%2 == 0
	})

	is.Equal([]int{1, 3}, slices.Collect(r1))

	r2 := Reject(values("Smith", "foo", "Domin", "bar", "Olivia"), func(x string) bool {
		return len(x) > 3
	})

	is.Equal([]string{"foo", "bar"}, slices.Collect(r2))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Reject(allStrings, func(x string) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestRejectI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := RejectI(values(1, 2, 3, 4), func(x, _ int) bool {
		return x%2 == 0
	})

	is.Equal([]int{1, 3}, slices.Collect(r1))

	r2 := RejectI(values("Smith", "foo", "Domin", "bar", "Olivia"), func(x string, _ int) bool {
		return len(x) > 3
	})

	is.Equal([]string{"foo", "bar"}, slices.Collect(r2))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := RejectI(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestRejectMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := RejectMap(values[int64](1, 2, 3, 4), func(x int64) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), false
		}
		return "", true
	})
	r2 := RejectMap(values("cpu", "gpu", "mouse", "keyboard"), func(x string) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", false
		}
		return "", true
	})

	is.Equal([]string{"2", "4"}, slices.Collect(r1))
	is.Equal([]string{"xpu", "xpu"}, slices.Collect(r2))
}

func TestRejectMapI(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := RejectMapI(values[int64](1, 2, 3, 4), func(x int64, _ int) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), false
		}
		return "", true
	})
	r2 := RejectMapI(values("cpu", "gpu", "mouse", "keyboard"), func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", false
		}
		return "", true
	})

	is.Equal([]string{"2", "4"}, slices.Collect(r1))
	is.Equal([]string{"xpu", "xpu"}, slices.Collect(r2))
}

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	count1 := Count(values(1, 2, 1), 1)
	count2 := Count(values(1, 2, 1), 3)
	count3 := Count(values[int](), 1)

	is.Equal(2, count1)
	is.Zero(count2)
	is.Zero(count3)
}

func TestCountBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	count1 := CountBy(values(1, 2, 1), func(i int) bool {
		return i < 2
	})
	count2 := CountBy(values(1, 2, 1), func(i int) bool {
		return i > 2
	})
	count3 := CountBy(values[int](), func(i int) bool {
		return i <= 2
	})

	is.Equal(2, count1)
	is.Zero(count2)
	is.Zero(count3)
}

func TestCountValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Empty(CountValues(values[int]()))
	is.Equal(map[int]int{1: 1, 2: 1}, CountValues(values(1, 2)))
	is.Equal(map[int]int{1: 1, 2: 2}, CountValues(values(1, 2, 2)))
	is.Equal(map[string]int{"": 1, "foo": 1, "bar": 1}, CountValues(values("foo", "bar", "")))
	is.Equal(map[string]int{"foo": 1, "bar": 2}, CountValues(values("foo", "bar", "bar")))
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

	result1 := CountValuesBy(values[int](), oddEven)
	result2 := CountValuesBy(values(1, 2), oddEven)
	result3 := CountValuesBy(values(1, 2, 2), oddEven)
	result4 := CountValuesBy(values("foo", "bar", ""), length)
	result5 := CountValuesBy(values("foo", "bar", "bar"), length)

	is.Empty(result1)
	is.Equal(map[bool]int{true: 1, false: 1}, result2)
	is.Equal(map[bool]int{true: 2, false: 1}, result3)
	is.Equal(map[int]int{0: 1, 3: 2}, result4)
	is.Equal(map[int]int{3: 3}, result5)
}

func TestSubset(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := values(0, 1, 2, 3, 4)

	out1 := Subset(in, 0, 0)
	out2 := Subset(in, 10, 2)
	out4 := Subset(in, 0, 10)
	out5 := Subset(in, 0, 2)
	out6 := Subset(in, 2, 2)
	out7 := Subset(in, 2, 5)
	out8 := Subset(in, 2, 3)
	out9 := Subset(in, 2, 4)

	is.Empty(slices.Collect(out1))
	is.Empty(slices.Collect(out2))
	is.Equal([]int{0, 1, 2, 3, 4}, slices.Collect(out4))
	is.Equal([]int{0, 1}, slices.Collect(out5))
	is.Equal([]int{2, 3}, slices.Collect(out6))
	is.Equal([]int{2, 3, 4}, slices.Collect(out7))
	is.Equal([]int{2, 3, 4}, slices.Collect(out8))
	is.Equal([]int{2, 3, 4}, slices.Collect(out9))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Subset(allStrings, 0, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := values(0, 1, 2, 3, 4)

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

	is.Empty(slices.Collect(out1))
	is.Equal([]int{0}, slices.Collect(out2))
	is.Equal([]int{0, 1, 2, 3, 4}, slices.Collect(out3))
	is.Equal([]int{0, 1, 2, 3, 4}, slices.Collect(out4))
	is.Empty(slices.Collect(out5))
	is.Equal([]int{1, 2, 3, 4}, slices.Collect(out6))
	is.Equal([]int{1, 2, 3, 4}, slices.Collect(out7))
	is.Equal([]int{4}, slices.Collect(out8))
	is.Empty(slices.Collect(out9))
	is.Empty(slices.Collect(out10))
	is.Empty(slices.Collect(out11))
	is.Empty(slices.Collect(out12))
	is.Empty(slices.Collect(out13))
	is.Empty(slices.Collect(out14))
	is.Empty(slices.Collect(out15))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Slice(allStrings, 0, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplace(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := values(0, 1, 0, 1, 2, 3, 0)

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

	is.Equal([]int{42, 1, 42, 1, 2, 3, 0}, slices.Collect(out1))
	is.Equal([]int{42, 1, 0, 1, 2, 3, 0}, slices.Collect(out2))
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, slices.Collect(out3))
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, slices.Collect(out4))
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, slices.Collect(out5))
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, slices.Collect(out6))
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, slices.Collect(out7))
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, slices.Collect(out8))
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, slices.Collect(out9))
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, slices.Collect(out10))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Replace(allStrings, "0", "2", 1)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := values(0, 1, 0, 1, 2, 3, 0)

	out1 := ReplaceAll(in, 0, 42)
	out2 := ReplaceAll(in, -1, 42)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, slices.Collect(out1))
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, slices.Collect(out2))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := ReplaceAll(allStrings, "0", "2")
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestCompact(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Compact(values(2, 0, 4, 0))

	is.Equal([]int{2, 4}, slices.Collect(r1))

	r2 := Compact(values("", "foo", "", "bar", ""))

	is.Equal([]string{"foo", "bar"}, slices.Collect(r2))

	r3 := Compact(values(true, false, true, false))

	is.Equal([]bool{true, true}, slices.Collect(r3))

	type foo struct {
		bar int
		baz string
	}

	// sequence of structs
	// If all fields of an element are zero values, Compact removes it.

	r4 := Compact(values(
		foo{bar: 1, baz: "a"}, // all fields are non-zero values
		foo{bar: 0, baz: ""},  // all fields are zero values
		foo{bar: 2, baz: ""},  // bar is non-zero
	))

	is.Equal([]foo{{bar: 1, baz: "a"}, {bar: 2, baz: ""}}, slices.Collect(r4))

	// sequence of pointers to structs
	// If an element is nil, Compact removes it.

	e1, e2, e3 := foo{bar: 1, baz: "a"}, foo{bar: 0, baz: ""}, foo{bar: 2, baz: ""}
	// NOTE: e2 is a zero value of foo, but its pointer &e2 is not a zero value of *foo.
	r5 := Compact(values(&e1, &e2, nil, &e3))

	is.Equal([]*foo{&e1, &e2, &e3}, slices.Collect(r5))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Compact(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestIsSorted(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(IsSorted(values(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)))
	is.True(IsSorted(values("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")))

	is.False(IsSorted(values(0, 1, 4, 3, 2, 5, 6, 7, 8, 9, 10)))
	is.False(IsSorted(values("a", "b", "d", "c", "e", "f", "g", "h", "i", "j")))
}

func TestIsSortedBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(IsSortedBy(values("a", "bb", "ccc"), func(s string) int {
		return len(s)
	}))

	is.False(IsSortedBy(values("aa", "b", "ccc"), func(s string) int {
		return len(s)
	}))

	is.True(IsSortedBy(values("1", "2", "3", "11"), func(s string) int {
		ret, _ := strconv.Atoi(s)
		return ret
	}))
}

func TestSplice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sample := values("a", "b", "c", "d", "e", "f", "g")

	// normal case
	results := slices.Collect(Splice(sample, 1, "1", "2"))
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(sample))
	is.Equal([]string{"a", "1", "2", "b", "c", "d", "e", "f", "g"}, results)

	// positive overflow
	results = slices.Collect(Splice(sample, 42, "1", "2"))
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(sample))
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g", "1", "2"}, results)

	// other
	is.Equal([]string{"1", "2"}, slices.Collect(Splice(values[string](), 0, "1", "2")))
	is.Equal([]string{"1", "2"}, slices.Collect(Splice(values[string](), 1, "1", "2")))
	is.Equal([]string{"1", "2", "0"}, slices.Collect(Splice(values("0"), 0, "1", "2")))
	is.Equal([]string{"0", "1", "2"}, slices.Collect(Splice(values("0"), 1, "1", "2")))

	// type preserved
	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Splice(allStrings, 1, "1", "2")
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestCutPrefix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	actual, result := CutPrefix(values("a", "a", "b"), []string{"a"})
	is.True(result)
	is.Equal([]string{"a", "b"}, slices.Collect(actual))

	actual, result = CutPrefix(values("a", "a", "b"), []string{"a"})
	is.True(result)
	is.Equal([]string{"a", "b"}, slices.Collect(actual))

	actual, result = CutPrefix(values("a", "a", "b"), []string{"b"})
	is.False(result)
	is.Equal([]string{"a", "a", "b"}, slices.Collect(actual))

	actual, result = CutPrefix(values[string](), []string{"b"})
	is.False(result)
	is.Empty(slices.Collect(actual))

	actual, result = CutPrefix(values("a", "a", "b"), []string{})
	is.True(result)
	is.Equal([]string{"a", "a", "b"}, slices.Collect(actual))

	actual, result = CutPrefix(values("a", "a", "b"), []string{"a", "a", "b", "b"})
	is.False(result)
	is.Equal([]string{"a", "a", "b"}, slices.Collect(actual))

	actual, result = CutPrefix(values("a", "a", "b"), []string{"a", "b"})
	is.False(result)
	is.Equal([]string{"a", "a", "b"}, slices.Collect(actual))
}

func TestCutSuffix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	actual, result := CutSuffix(values("a", "a", "b"), []string{"c"})
	is.False(result)
	is.Equal([]string{"a", "a", "b"}, slices.Collect(actual))

	actual, result = CutSuffix(values("a", "a", "b"), []string{"b"})
	is.True(result)
	is.Equal([]string{"a", "a"}, slices.Collect(actual))

	actual, result = CutSuffix(values("a", "a", "b"), []string{})
	is.True(result)
	is.Equal([]string{"a", "a", "b"}, slices.Collect(actual))

	actual, result = CutSuffix(values("a", "a", "b"), []string{"a"})
	is.False(result)
	is.Equal([]string{"a", "a", "b"}, slices.Collect(actual))

	actual, result = CutSuffix(values("a", "a", "b"), []string{})
	is.True(result)
	is.Equal([]string{"a", "a", "b"}, slices.Collect(actual))
}

func TestTrim(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	actual := Trim(values("a", "b", "c", "d", "e", "f", "g"), "a", "b")
	is.Equal([]string{"c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = Trim(values("a", "b", "c", "d", "e", "f", "g"), "g", "f")
	is.Equal([]string{"a", "b", "c", "d", "e"}, slices.Collect(actual))
	actual = Trim(values("a", "b", "c", "d", "e", "f", "g"), "a", "b", "c", "d", "e", "f", "g")
	is.Empty(slices.Collect(actual))
	actual = Trim(values("a", "b", "c", "d", "e", "f", "g"), "a", "b", "c", "d", "e", "f", "g", "h")
	is.Empty(slices.Collect(actual))
	actual = Trim(values("a", "b", "c", "d", "e", "f", "g"))
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
}

func TestTrimFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	actual := TrimFirst(values("a", "a", "b", "c", "d", "e", "f", "g"), "a", "b")
	is.Equal([]string{"c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimFirst(values("a", "b", "c", "d", "e", "f", "g"), "b", "a")
	is.Equal([]string{"c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimFirst(values("a", "b", "c", "d", "e", "f", "g"), "g", "f")
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimFirst(values("a", "b", "c", "d", "e", "f", "g"), "a", "b", "c", "d", "e", "f", "g")
	is.Empty(slices.Collect(actual))
	actual = TrimFirst(values("a", "b", "c", "d", "e", "f", "g"), "a", "b", "c", "d", "e", "f", "g", "h")
	is.Empty(slices.Collect(actual))
	actual = TrimFirst(values("a", "b", "c", "d", "e", "f", "g"))
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
}

func TestTrimPrefix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	actual := TrimPrefix(values("a", "b", "a", "b", "c", "d", "e", "f", "g"), []string{"a", "b"})
	is.Equal([]string{"c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimPrefix(values("a", "b", "c", "d", "e", "f", "g"), []string{"b", "a"})
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimPrefix(values("a", "b", "c", "d", "e", "f", "g"), []string{"g", "f"})
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimPrefix(values("a", "b", "c", "d", "e", "f", "g"), []string{"a", "b", "c", "d", "e", "f", "g"})
	is.Empty(slices.Collect(actual))
	actual = TrimPrefix(values("a", "b", "c", "d", "e", "f", "g"), []string{"a", "b", "c", "d", "e", "f", "g", "h"})
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimPrefix(values("a", "b", "c", "d", "e", "f", "g"), []string{})
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
}

func TestTrimLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	actual := TrimLast(values("a", "b", "c", "d", "e", "f", "g"), "a", "b")
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimLast(values("a", "b", "c", "d", "e", "f", "g", "g"), "g", "f")
	is.Equal([]string{"a", "b", "c", "d", "e"}, slices.Collect(actual))
	actual = TrimLast(values("a", "b", "c", "d", "e", "f", "g"), "a", "b", "c", "d", "e", "f", "g")
	is.Empty(slices.Collect(actual))
	actual = TrimLast(values("a", "b", "c", "d", "e", "f", "g"), "a", "b", "c", "d", "e", "f", "g", "h")
	is.Empty(slices.Collect(actual))
	actual = TrimLast(values("a", "b", "c", "d", "e", "f", "g"))
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
}

func TestTrimSuffix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	actual := TrimSuffix(values("a", "b", "c", "d", "e", "f", "g"), []string{"a", "b"})
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimSuffix(values("a", "b", "c", "d", "e", "f", "g", "f", "g"), []string{"f", "g"})
	is.Equal([]string{"a", "b", "c", "d", "e"}, slices.Collect(actual))
	actual = TrimSuffix(values("a", "b", "c", "d", "e", "f", "g", "f", "g"), []string{"g", "f"})
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g", "f", "g"}, slices.Collect(actual))
	actual = TrimSuffix(values("a", "b", "c", "d", "e", "f", "f", "g"), []string{"f", "g"})
	is.Equal([]string{"a", "b", "c", "d", "e", "f"}, slices.Collect(actual))
	actual = TrimSuffix(values("a", "b", "c", "d", "e", "f", "g"), []string{"a", "b", "c", "d", "e", "f", "g"})
	is.Empty(slices.Collect(actual))
	actual = TrimSuffix(values("a", "b", "c", "d", "e", "f", "g"), []string{"a", "b", "c", "d", "e", "f", "g", "h"})
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
	actual = TrimSuffix(values("a", "b", "c", "d", "e", "f", "g"), []string{})
	is.Equal([]string{"a", "b", "c", "d", "e", "f", "g"}, slices.Collect(actual))
}
