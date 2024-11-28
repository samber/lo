package la

import (
	"fmt"
	"github.com/samber/lo"
	"iter"
	"maps"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fromSlices[T any](res ...[]T) []iter.Seq[T] {
	out := make([]iter.Seq[T], len(res))

	for idx, v := range res {
		out[idx] = slices.Values(v)
	}

	return out
}

func fromSlicesEnumerated[T any](res ...[]T) []iter.Seq2[int, T] {
	out := make([]iter.Seq2[int, T], len(res))

	for idx, v := range res {
		out[idx] = slices.All(v)
	}

	return out
}

func fromMaps[K comparable, V any](res ...map[K]V) []iter.Seq2[K, V] {
	out := make([]iter.Seq2[K, V], len(res))

	for idx, v := range res {
		out[idx] = maps.All(v)
	}

	return out
}

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Map(slices.Values([]int{1, 2, 3, 4}), func(x int) string {
		return "Hello"
	}))
	result2 := slices.Collect(Map(slices.Values([]int64{1, 2, 3, 4}), func(x int64) string {
		return strconv.FormatInt(x, 10)
	}))

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "3", "4"})
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(FilterMap(slices.Values([]int64{1, 2, 3, 4}), func(x int64) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), true
		}
		return "", false
	}))
	r2 := slices.Collect(FilterMap(slices.Values([]string{"cpu", "gpu", "mouse", "keyboard"}), func(x string) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	}))

	is.Equal(len(r1), 2)
	is.Equal(len(r2), 2)
	is.Equal(r1, []string{"2", "4"})
	is.Equal(r2, []string{"xpu", "xpu"})
}

func TestFilterMap2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(Tuples(FilterMap2(
		Enumerate(slices.Values([]int64{1, 2, 3, 4})),
		func(_ int, x int64) (string, bool) {
			if x%2 == 0 {
				return strconv.FormatInt(x, 10), true
			}

			return "", false
		},
	)))
	r2 := slices.Collect(Tuples(FilterMap2(
		Enumerate(slices.Values([]string{"cpu", "gpu", "mouse", "keyboard"})),
		func(_ int, x string) (string, bool) {
			if strings.HasSuffix(x, "pu") {
				return "xpu", true
			}
			return "", false
		})))

	is.Equal(len(r1), 2)
	is.Equal(len(r2), 2)
	is.Equal(r1, []lo.Tuple2[int, string]{{1, "2"}, {3, "4"}})
	is.Equal(r2, []lo.Tuple2[int, string]{{0, "xpu"}, {1, "xpu"}})
}

func TestFlatMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(FlatMap(slices.Values([]int{0, 1, 2, 3, 4}), func(x int) iter.Seq[string] {
		return slices.Values([]string{"Hello"})
	}))
	result2 := slices.Collect(FlatMap(slices.Values([]int64{0, 1, 2, 3, 4}), func(x int64) iter.Seq[string] {
		return func(yield func(string) bool) {
			for i := int64(0); i < x; i++ {
				if !yield(strconv.FormatInt(x, 10)) {
					return
				}
			}
		}
	}))

	is.Equal(len(result1), 5)
	is.Equal(len(result2), 10)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"})
}

func TestFlatMapEnumerated(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Tuples(FlatMapEnumerated(
		Enumerate(slices.Values([]int{0, 1, 2, 3, 4})),
		func(_ int, x int) iter.Seq[string] {
			return slices.Values([]string{"Hello"})
		},
	)))
	result2 := slices.Collect(Tuples(FlatMapEnumerated(
		Enumerate(slices.Values([]int64{0, 1, 2, 3, 4})),
		func(_ int, x int64) iter.Seq[string] {
			return func(yield func(string) bool) {
				for i := int64(0); i < x; i++ {
					if !yield(strconv.FormatInt(x, 10)) {
						return
					}
				}
			}
		},
	)))

	is.Equal(len(result1), 5)
	is.Equal(len(result2), 10)
	is.Equal(result1, []lo.Tuple2[int, string]{{0, "Hello"}, {1, "Hello"}, {2, "Hello"}, {3, "Hello"}, {4, "Hello"}})
	is.Equal(result2, []lo.Tuple2[int, string]{{0, "1"}, {1, "2"}, {2, "2"}, {3, "3"}, {4, "3"}, {5, "3"}, {6, "4"}, {7, "4"}, {8, "4"}, {9, "4"}})
}

func TestChunk(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Chunk(slices.Values([]int{0, 1, 2, 3, 4, 5}), 2))
	result2 := slices.Collect(Chunk(slices.Values([]int{0, 1, 2, 3, 4, 5, 6}), 2))
	result3 := slices.Collect(Chunk(slices.Values([]int{}), 2))
	result4 := slices.Collect(Chunk(slices.Values([]int{0}), 2))

	is.Equal(result1, [][]int{{0, 1}, {2, 3}, {4, 5}})
	is.Equal(result2, [][]int{{0, 1}, {2, 3}, {4, 5}, {6}})
	is.Equal(result3, ([][]int)(nil))
	is.Equal(result4, [][]int{{0}})
	is.PanicsWithValue("Second parameter must be greater than 0", func() {
		Chunk(slices.Values([]int{0}), 0)
	})

	// appending to a chunk should not affect original array
	originalArray := []int{0, 1, 2, 3, 4, 5}
	result5 := slices.Collect(Chunk(slices.Values(originalArray), 2))
	result5[0] = append(result5[0], 6)
	is.Equal(originalArray, []int{0, 1, 2, 3, 4, 5})
}

func TestChunk2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Map(Chunk2(Enumerate(slices.Values([]int{0, 1, 2, 3, 4, 5})), 2), maps.Collect))
	result2 := slices.Collect(Map(Chunk2(Enumerate(slices.Values([]int{0, 1, 2, 3, 4, 5, 6})), 2), maps.Collect))
	result3 := slices.Collect(Map(Chunk2(Enumerate(slices.Values([]int{})), 2), maps.Collect))
	result4 := slices.Collect(Map(Chunk2(Enumerate(slices.Values([]int{0})), 2), maps.Collect))

	is.Equal(result1, []map[int]int{{0: 0, 1: 1}, {2: 2, 3: 3}, {4: 4, 5: 5}})
	is.Equal(result2, []map[int]int{{0: 0, 1: 1}, {2: 2, 3: 3}, {4: 4, 5: 5}, {6: 6}})
	is.Equal(result3, ([]map[int]int)(nil))
	is.Equal(result4, []map[int]int{{0: 0}})
	is.PanicsWithValue("Second parameter must be greater than 0", func() {
		Chunk(slices.Values([]int{0}), 0)
	})
}

func TestFlattenSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(FlattenSlice(slices.Values([][]int{{0, 1}, {2, 3, 4, 5}})))

	is.Equal(result1, []int{0, 1, 2, 3, 4, 5})
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Flatten(slices.Values([]iter.Seq[int]{
		slices.Values([]int{0, 1}),
		slices.Values([]int{2, 3, 4, 5}),
	})))

	is.Equal(result1, []int{0, 1, 2, 3, 4, 5})
}

func TestInterleave(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name        string
		collections []iter.Seq[int]
		want        []int
	}{
		{
			"nil",
			fromSlices(([]int)(nil)),
			([]int)(nil),
		},
		{
			"empties",
			fromSlices([]int{}, []int{}),
			([]int)(nil),
		},
		{
			"same length",
			fromSlices([][]int{{1, 3, 5}, {2, 4, 6}}...),
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"different length",
			fromSlices([][]int{{1, 3, 5, 6}, {2, 4}}...),
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"many slices",
			fromSlices([][]int{{1}, {2, 5, 8}, {3, 6}, {4, 7, 9, 10}}...),
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Collect(Interleave(tt.collections...)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interleave() = %v, want %v", got, tt.want)
			}
		})
	}

	type myIter iter.Seq[string]
	var allStrings myIter = func(yield func(string) bool) {
		for _, str := range []string{"", "foo", "bar"} {
			if !yield(str) {
				return
			}
		}
	}

	nonempty := Interleave(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestInterleave2(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name        string
		collections []iter.Seq2[int, int]
		want        []lo.Tuple2[int, int]
	}{
		{
			"nil",
			fromSlicesEnumerated(([]int)(nil)),
			([]lo.Tuple2[int, int])(nil),
		},
		{
			"empties",
			fromSlicesEnumerated([]int{}, []int{}),
			([]lo.Tuple2[int, int])(nil),
		},
		{
			"same length",
			fromSlicesEnumerated([][]int{{1, 3, 5}, {2, 4, 6}}...),
			[]lo.Tuple2[int, int]{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
		},
		{
			"different length",
			fromSlicesEnumerated([][]int{{1, 3, 5, 6}, {2, 4}}...),
			[]lo.Tuple2[int, int]{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {3, 6}},
		},
		{
			"many slices",
			fromSlicesEnumerated([][]int{{1}, {2, 5, 8}, {3, 6}, {4, 7, 9, 10}}...),
			[]lo.Tuple2[int, int]{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 5}, {1, 6}, {1, 7}, {2, 8}, {2, 9}, {3, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Collect(Tuples(Interleave2(tt.collections...))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interleave() = %v, want %v", got, tt.want)
			}
		})
	}

	type myIter iter.Seq2[int, string]
	var allStrings myIter = func(yield func(int, string) bool) {
		for i, str := range []string{"", "foo", "bar"} {
			if !yield(i, str) {
				return
			}
		}
	}

	nonempty := Interleave2(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := maps.Collect(KeyBy(slices.Values([]string{"a", "aa", "aaa"}), func(str string) int {
		return len(str)
	}))

	is.Equal(result1, map[int]string{1: "a", 2: "aa", 3: "aaa"})
}

func TestSeqToSeq2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := maps.Collect(SeqToSeq2(slices.Values([]string{"a", "aa", "aaa"}), func(str string) int {
		return len(str)
	}))

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
		in     iter.Seq[*foo]
		expect map[string]int
	}{
		{
			in:     slices.Values([]*foo{{baz: "apple", bar: 1}}),
			expect: map[string]int{"apple": 1},
		},
		{
			in:     slices.Values([]*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}}),
			expect: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:     slices.Values([]*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}}),
			expect: map[string]int{"apple": 2},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			is := assert.New(t)
			is.Equal(maps.Collect(Associate(testCase.in, transform)), testCase.expect)
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
		in     iter.Seq[*foo]
		expect map[string]int
	}{
		{
			in:     slices.Values([]*foo{{baz: "apple", bar: 1}}),
			expect: map[string]int{"apple": 1},
		},
		{
			in:     slices.Values([]*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}}),
			expect: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:     slices.Values([]*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}}),
			expect: map[string]int{"apple": 2},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			is := assert.New(t)
			is.Equal(maps.Collect(SliceToMap(testCase.in, transform)), testCase.expect)
		})
	}
}

func TestRejectMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(RejectMap(slices.Values([]int64{1, 2, 3, 4}), func(x int64) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), false
		}
		return "", true
	}))

	r2 := slices.Collect(RejectMap(slices.Values([]string{"cpu", "gpu", "mouse", "keyboard"}), func(x string) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", false
		}
		return "", true
	}))

	is.Equal(len(r1), 2)
	is.Equal(len(r2), 2)
	is.Equal(r1, []string{"2", "4"})
	is.Equal(r2, []string{"xpu", "xpu"})
}

func TestMapKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := maps.Collect(MapKeys(maps.All(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}), func(x int, _ int) string {
		return "Hello"
	}))
	result2 := maps.Collect(MapKeys(maps.All(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}), func(_ int, v int) string {
		return strconv.FormatInt(int64(v), 10)
	}))

	is.Equal(len(result1), 1)
	is.Equal(len(result2), 4)
	is.Equal(result2, map[string]int{"1": 1, "2": 2, "3": 3, "4": 4})
}

func TestMapValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := maps.Collect(MapValues(maps.All(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}), func(x int, _ int) string {
		return "Hello"
	}))
	result2 := maps.Collect(MapValues(maps.All(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}), func(x int, _ int) string {
		return strconv.FormatInt(int64(x), 10)
	}))

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, map[int]string{1: "Hello", 2: "Hello", 3: "Hello", 4: "Hello"})
	is.Equal(result2, map[int]string{1: "1", 2: "2", 3: "3", 4: "4"})
}

func map2Test[I any, O any](t *testing.T, in map[string]I, iteratee func(string, I) (string, O), expected map[string]O) {
	is := assert.New(t)
	result := maps.Collect(Map2(maps.All(in), iteratee))
	is.Equal(result, expected)
}

func map2TestIter[I any, O any](t *testing.T, in iter.Seq2[string, I], iteratee func(string, I) (string, O), expected map[string]O) {
	is := assert.New(t)
	result := maps.Collect(Map2(in, iteratee))
	is.Equal(result, expected)
}

func TestMap2(t *testing.T) {
	map2Test(t, map[string]int{"foo": 1, "bar": 2}, func(k string, v int) (string, int) {
		return k, v + 1
	}, map[string]int{"foo": 2, "bar": 3})
	map2Test(t, map[string]int{"foo": 1, "bar": 2}, func(k string, v int) (string, string) {
		return k, k + strconv.Itoa(v)
	}, map[string]string{"foo": "foo1", "bar": "bar2"})
	map2Test(t, map[string]int{"foo": 1, "bar": 2}, func(k string, v int) (string, string) {
		return k, strconv.Itoa(v) + k
	}, map[string]string{"foo": "1foo", "bar": "2bar"})

	// NoMutation
	{
		is := assert.New(t)
		r1 := map[string]int{"foo": 1, "bar": 2}
		maps.Collect(Map2(maps.All(r1), func(k string, v int) (string, string) {
			return k, strconv.Itoa(v) + "!!"
		}))
		is.Equal(r1, map[string]int{"foo": 1, "bar": 2})
	}
	// EmptyInput
	{
		map2Test(t, map[string]int{}, func(k string, v int) (string, string) {
			return k, strconv.Itoa(v) + "!!"
		}, map[string]string{})

		map2Test(t, map[string]any{}, func(k string, v any) (string, any) {
			return k, v
		}, map[string]any{})
	}
	// Identity
	{
		map2Test(t, map[string]int{"foo": 1, "bar": 2}, func(k string, v int) (string, int) {
			return k, v
		}, map[string]int{"foo": 1, "bar": 2})
		map2Test(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(k string, v any) (string, any) {
			return k, v
		}, map[string]any{"foo": 1, "bar": "2", "ccc": true})
	}
	// ToConstantEntry
	{
		map2Test(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(k string, v any) (string, any) {
			return "key", "value"
		}, map[string]any{"key": "value"})
		map2Test(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(k string, v any) (string, any) {
			return "b", 5
		}, map[string]any{"b": 5})
	}

	// OverlappingKeys
	{
		map2TestIter(t,
			FromTuples([]lo.Tuple2[string, any]{{"foo", 1}, {"foo2", 2}, {"Foo", 2}, {"Foo2", "2"}, {"bar", "2"}, {"ccc", true}}),
			func(k string, v any) (string, any) {
				return string(k[0]), v
			},
			map[string]any{"F": "2", "b": "2", "c": true, "f": 2},
		)
		map2TestIter(t,
			FromTuples([]lo.Tuple2[string, string]{{"foo", "1"}, {"foo2", "2"}, {"Foo", "2"}, {"Foo2", "2"}, {"bar", "2"}, {"ccc", "true"}}),
			func(k string, v string) (string, string) {
				return v, k
			},
			map[string]string{"1": "foo", "2": "bar", "true": "ccc"},
		)
	}
	//NormalMappers
	{
		map2Test(t, map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"}, func(k string, v string) (string, string) {
			return k, k + v
		}, map[string]string{"Foo": "Foo2", "Foo2": "Foo22", "bar": "bar2", "ccc": "ccctrue", "foo": "foo1", "foo2": "foo22"})

		map2Test(t, map[string]struct {
			name string
			age  int
		}{"1-11-1": {name: "foo", age: 1}, "2-22-2": {name: "bar", age: 2}}, func(k string, v struct {
			name string
			age  int
		},
		) (string, string) {
			return v.name, k
		}, map[string]string{"bar": "2-22-2", "foo": "1-11-1"})
	}
}

func BenchmarkJoin2(b *testing.B) {
	counts := []int{32768, 1024, 128, 32, 2}

	allDifferentMap := func(b *testing.B, n int) []iter.Seq2[string, int] {
		defer b.ResetTimer()
		m := make([]iter.Seq2[string, int], 0)
		for i := 0; i < n; i++ {
			m = append(m, maps.All(map[string]int{
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
			}))
		}
		return m
	}

	allTheSameMap := func(b *testing.B, n int) []iter.Seq2[string, int] {
		defer b.ResetTimer()
		m := make([]iter.Seq2[string, int], 0)
		for i := 0; i < n; i++ {
			m = append(m, maps.All(map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
				"e": 5,
				"f": 6,
			}))
		}
		return m
	}

	for _, count := range counts {
		differentMap := allDifferentMap(b, count)
		sameMap := allTheSameMap(b, count)

		b.Run(fmt.Sprintf("%d", count), func(b *testing.B) {
			testcase := []struct {
				name  string
				maps  []iter.Seq2[string, int]
				count int
			}{
				{"different", differentMap, count},
				{"same", sameMap, 6},
			}

			for _, tc := range testcase {
				b.Run(tc.name, func(b *testing.B) {
					b.ResetTimer()
					for n := 0; n < b.N; n++ {
						result := Join2(tc.maps...)
						_ = CollectMap(result, WithMapCapacity(tc.count))
					}
				})
			}
		})
	}
}
