package lo

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Keys(map[string]int{"foo": 1, "bar": 2})
	sort.Strings(r1)
	is.Equal(r1, []string{"bar", "foo"})

	r2 := Keys(map[string]int{})
	is.Empty(r2)

	r3 := Keys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
	sort.Strings(r3)
	is.Equal(r3, []string{"bar", "baz", "foo"})

	r4 := Keys[string, int]()
	is.Equal(r4, []string{})

	r5 := Keys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"bar": 3})
	sort.Strings(r5)
	is.Equal(r5, []string{"bar", "bar", "foo"})
}

func TestUniqKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := UniqKeys(map[string]int{"foo": 1, "bar": 2})
	sort.Strings(r1)
	is.Equal(r1, []string{"bar", "foo"})

	r2 := UniqKeys(map[string]int{})
	is.Empty(r2)

	r3 := UniqKeys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
	sort.Strings(r3)
	is.Equal(r3, []string{"bar", "baz", "foo"})

	r4 := UniqKeys[string, int]()
	is.Equal(r4, []string{})

	r5 := UniqKeys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3})
	sort.Strings(r5)
	is.Equal(r5, []string{"bar", "foo"})

	// check order
	r6 := UniqKeys(map[string]int{"foo": 1}, map[string]int{"bar": 3})
	is.Equal(r6, []string{"foo", "bar"})
}

func TestHasKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := HasKey(map[string]int{"foo": 1}, "bar")
	is.False(r1)

	r2 := HasKey(map[string]int{"foo": 1}, "foo")
	is.True(r2)
}

func TestValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Values(map[string]int{"foo": 1, "bar": 2})
	sort.Ints(r1)
	is.Equal(r1, []int{1, 2})

	r2 := Values(map[string]int{})
	is.Empty(r2)

	r3 := Values(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
	sort.Ints(r3)
	is.Equal(r3, []int{1, 2, 3})

	r4 := Values[string, int]()
	is.Equal(r4, []int{})

	r5 := Values(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3})
	sort.Ints(r5)
	is.Equal(r5, []int{1, 1, 2, 3})
}

func TestUniqValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := UniqValues(map[string]int{"foo": 1, "bar": 2})
	sort.Ints(r1)
	is.Equal(r1, []int{1, 2})

	r2 := UniqValues(map[string]int{})
	is.Empty(r2)

	r3 := UniqValues(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
	sort.Ints(r3)
	is.Equal(r3, []int{1, 2, 3})

	r4 := UniqValues[string, int]()
	is.Equal(r4, []int{})

	r5 := UniqValues(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3})
	sort.Ints(r5)
	is.Equal(r5, []int{1, 2, 3})

	r6 := UniqValues(map[string]int{"foo": 1, "bar": 1}, map[string]int{"foo": 1, "bar": 3})
	sort.Ints(r6)
	is.Equal(r6, []int{1, 3})

	// check order
	r7 := UniqValues(map[string]int{"foo": 1}, map[string]int{"bar": 3})
	is.Equal(r7, []int{1, 3})
}

func TestValueOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := ValueOr(map[string]int{"foo": 1}, "bar", 2)
	is.Equal(r1, 2)

	r2 := ValueOr(map[string]int{"foo": 1}, "foo", 2)
	is.Equal(r2, 1)
}

func TestPickBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := PickBy(before, func(key string, value int) bool { return true })
	is.IsType(after, before, "type preserved")
}

func TestPickByKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz", "qux"})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := PickByKeys(before, []string{"foobar", "baz"})
	is.IsType(after, before, "type preserved")
}

func TestPickByValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := PickByValues(before, []int{0, 3})
	is.IsType(after, before, "type preserved")
}

func TestOmitBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := OmitBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})

	is.Equal(r1, map[string]int{"bar": 2})

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := PickBy(before, func(key string, value int) bool { return true })
	is.IsType(after, before, "type preserved")
}

func TestOmitByKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := OmitByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz", "qux"})

	is.Equal(r1, map[string]int{"bar": 2})

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := OmitByKeys(before, []string{"foobar", "baz"})
	is.IsType(after, before, "type preserved")
}

func TestOmitByValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := OmitByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

	is.Equal(r1, map[string]int{"bar": 2})

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := OmitByValues(before, []int{0, 3})
	is.IsType(after, before, "type preserved")
}

func TestEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Entries(map[string]int{"foo": 1, "bar": 2})

	sort.Slice(r1, func(i, j int) bool {
		return r1[i].Value < r1[j].Value
	})
	is.EqualValues(r1, []Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
	})
}

func TestToPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := ToPairs(map[string]int{"baz": 3, "qux": 4})

	sort.Slice(r1, func(i, j int) bool {
		return r1[i].Value < r1[j].Value
	})
	is.EqualValues(r1, []Entry[string, int]{
		{
			Key:   "baz",
			Value: 3,
		},
		{
			Key:   "qux",
			Value: 4,
		},
	})
}

func TestFromEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FromEntries([]Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
	})

	is.Len(r1, 2)
	is.Equal(r1["foo"], 1)
	is.Equal(r1["bar"], 2)
}

func TestFromPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FromPairs([]Entry[string, int]{
		{
			Key:   "baz",
			Value: 3,
		},
		{
			Key:   "qux",
			Value: 4,
		},
	})

	is.Len(r1, 2)
	is.Equal(r1["baz"], 3)
	is.Equal(r1["qux"], 4)
}

func TestInvert(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Invert(map[string]int{"a": 1, "b": 2})
	r2 := Invert(map[string]int{"a": 1, "b": 2, "c": 1})

	is.Len(r1, 2)
	is.EqualValues(map[int]string{1: "a", 2: "b"}, r1)
	is.Len(r2, 2)
}

func TestAssign(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Assign(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})

	is.Len(result1, 3)
	is.Equal(result1, map[string]int{"a": 1, "b": 3, "c": 4})

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := Assign(before, before)
	is.IsType(after, before, "type preserved")
}

func TestChunkEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ChunkEntries(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, 2)
	result2 := ChunkEntries(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, 3)
	result3 := ChunkEntries(map[string]int{}, 2)
	result4 := ChunkEntries(map[string]int{"a": 1}, 2)
	result5 := ChunkEntries(map[string]int{"a": 1, "b": 2}, 1)

	expectedCount1 := 3
	expectedCount2 := 2
	expectedCount3 := 0
	expectedCount4 := 1
	expectedCount5 := 2

	is.Len(result1, expectedCount1)
	is.Len(result2, expectedCount2)
	is.Len(result3, expectedCount3)
	is.Len(result4, expectedCount4)
	is.Len(result5, expectedCount5)

	is.PanicsWithValue("The chunk size must be greater than 0", func() {
		ChunkEntries(map[string]int{"a": 1}, 0)
	})
	is.PanicsWithValue("The chunk size must be greater than 0", func() {
		ChunkEntries(map[string]int{"a": 1}, -1)
	})

	type myStruct struct {
		Name  string
		Value int
	}

	allStructs := []myStruct{{"one", 1}, {"two", 2}, {"three", 3}}
	nonempty := ChunkEntries(map[string]myStruct{"a": allStructs[0], "b": allStructs[1], "c": allStructs[2]}, 2)
	is.Len(nonempty, 2)

	originalMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result6 := ChunkEntries(originalMap, 2)
	for k := range result6[0] {
		result6[0][k] = 10
	}
	is.Equal(originalMap, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5})
}

func TestMapKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapKeys(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x int, _ int) string {
		return "Hello"
	})
	result2 := MapKeys(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(_ int, v int) string {
		return strconv.FormatInt(int64(v), 10)
	})

	is.Equal(len(result1), 1)
	is.Equal(len(result2), 4)
	is.Equal(result2, map[string]int{"1": 1, "2": 2, "3": 3, "4": 4})
}

func TestMapValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapValues(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x int, _ int) string {
		return "Hello"
	})
	result2 := MapValues(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x int, _ int) string {
		return strconv.FormatInt(int64(x), 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, map[int]string{1: "Hello", 2: "Hello", 3: "Hello", 4: "Hello"})
	is.Equal(result2, map[int]string{1: "1", 2: "2", 3: "3", 4: "4"})
}

func mapEntriesTest[I any, O any](t *testing.T, in map[string]I, iteratee func(string, I) (string, O), expected map[string]O) {
	is := assert.New(t)
	result := MapEntries(in, iteratee)
	is.Equal(result, expected)
}

func TestMapEntries(t *testing.T) {
	mapEntriesTest(t, map[string]int{"foo": 1, "bar": 2}, func(k string, v int) (string, int) {
		return k, v + 1
	}, map[string]int{"foo": 2, "bar": 3})
	mapEntriesTest(t, map[string]int{"foo": 1, "bar": 2}, func(k string, v int) (string, string) {
		return k, k + strconv.Itoa(v)
	}, map[string]string{"foo": "foo1", "bar": "bar2"})
	mapEntriesTest(t, map[string]int{"foo": 1, "bar": 2}, func(k string, v int) (string, string) {
		return k, strconv.Itoa(v) + k
	}, map[string]string{"foo": "1foo", "bar": "2bar"})

	// NoMutation
	{
		is := assert.New(t)
		r1 := map[string]int{"foo": 1, "bar": 2}
		MapEntries(r1, func(k string, v int) (string, string) {
			return k, strconv.Itoa(v) + "!!"
		})
		is.Equal(r1, map[string]int{"foo": 1, "bar": 2})
	}
	// EmptyInput
	{
		mapEntriesTest(t, map[string]int{}, func(k string, v int) (string, string) {
			return k, strconv.Itoa(v) + "!!"
		}, map[string]string{})

		mapEntriesTest(t, map[string]any{}, func(k string, v any) (string, any) {
			return k, v
		}, map[string]any{})
	}
	// Identity
	{
		mapEntriesTest(t, map[string]int{"foo": 1, "bar": 2}, func(k string, v int) (string, int) {
			return k, v
		}, map[string]int{"foo": 1, "bar": 2})
		mapEntriesTest(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(k string, v any) (string, any) {
			return k, v
		}, map[string]any{"foo": 1, "bar": "2", "ccc": true})
	}
	// ToConstantEntry
	{
		mapEntriesTest(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(k string, v any) (string, any) {
			return "key", "value"
		}, map[string]any{"key": "value"})
		mapEntriesTest(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(k string, v any) (string, any) {
			return "b", 5
		}, map[string]any{"b": 5})
	}

	//// OverlappingKeys
	//// because using range over map, the order is not guaranteed
	//// this test is not deterministic
	//{
	//	mapEntriesTest(t, map[string]any{"foo": 1, "foo2": 2, "Foo": 2, "Foo2": "2", "bar": "2", "ccc": true}, func(k string, v any) (string, any) {
	//		return string(k[0]), v
	//	}, map[string]any{"F": "2", "b": "2", "c": true, "f": 2})
	//	mapEntriesTest(t, map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"}, func(k string, v string) (string, string) {
	//		return v, k
	//	}, map[string]string{"1": "foo", "2": "bar", "true": "ccc"})
	//}
	//NormalMappers
	{
		mapEntriesTest(t, map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"}, func(k string, v string) (string, string) {
			return k, k + v
		}, map[string]string{"Foo": "Foo2", "Foo2": "Foo22", "bar": "bar2", "ccc": "ccctrue", "foo": "foo1", "foo2": "foo22"})

		mapEntriesTest(t, map[string]struct {
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

func TestMapToSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapToSlice(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k int, v int) string {
		return fmt.Sprintf("%d_%d", k, v)
	})
	result2 := MapToSlice(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k int, _ int) string {
		return strconv.FormatInt(int64(k), 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.ElementsMatch(result1, []string{"1_5", "2_6", "3_7", "4_8"})
	is.ElementsMatch(result2, []string{"1", "2", "3", "4"})
}

func BenchmarkAssign(b *testing.B) {
	counts := []int{32768, 1024, 128, 32, 2}

	allDifferentMap := func(b *testing.B, n int) []map[string]int {
		defer b.ResetTimer()
		m := make([]map[string]int, 0)
		for i := 0; i < n; i++ {
			m = append(m, map[string]int{
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
			},
			)
		}
		return m
	}

	allTheSameMap := func(b *testing.B, n int) []map[string]int {
		defer b.ResetTimer()
		m := make([]map[string]int, 0)
		for i := 0; i < n; i++ {
			m = append(m, map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
				"e": 5,
				"f": 6,
			},
			)
		}
		return m
	}

	for _, count := range counts {
		differentMap := allDifferentMap(b, count)
		sameMap := allTheSameMap(b, count)

		b.Run(fmt.Sprintf("%d", count), func(b *testing.B) {
			testcase := []struct {
				name string
				maps []map[string]int
			}{
				{"different", differentMap},
				{"same", sameMap},
			}

			for _, tc := range testcase {
				b.Run(tc.name, func(b *testing.B) {
					b.ResetTimer()
					for n := 0; n < b.N; n++ {
						result := Assign(tc.maps...)
						_ = result
					}
				})
			}
		})

	}
}
