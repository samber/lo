package lo

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Keys(map[string]int{"foo": 1, "bar": 2})
	is.ElementsMatch(r1, []string{"bar", "foo"})

	r2 := Keys(map[string]int{})
	is.Empty(r2)

	r3 := Keys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
	is.ElementsMatch(r3, []string{"bar", "baz", "foo"})

	r4 := Keys[string, int]()
	is.Empty(r4)

	r5 := Keys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"bar": 3})
	is.ElementsMatch(r5, []string{"bar", "bar", "foo"})
}

func TestUniqKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := UniqKeys(map[string]int{"foo": 1, "bar": 2})
	is.ElementsMatch(r1, []string{"bar", "foo"})

	r2 := UniqKeys(map[string]int{})
	is.Empty(r2)

	r3 := UniqKeys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
	is.ElementsMatch(r3, []string{"bar", "baz", "foo"})

	r4 := UniqKeys[string, int]()
	is.Empty(r4)

	r5 := UniqKeys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3})
	is.ElementsMatch(r5, []string{"bar", "foo"})

	// check order
	r6 := UniqKeys(map[string]int{"foo": 1}, map[string]int{"bar": 3})
	is.Equal([]string{"foo", "bar"}, r6)
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
	is.ElementsMatch(r1, []int{1, 2})

	r2 := Values(map[string]int{})
	is.Empty(r2)

	r3 := Values(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
	is.ElementsMatch(r3, []int{1, 2, 3})

	r4 := Values[string, int]()
	is.Empty(r4)

	r5 := Values(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3})
	is.ElementsMatch(r5, []int{1, 1, 2, 3})
}

func TestUniqValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := UniqValues(map[string]int{"foo": 1, "bar": 2})
	is.ElementsMatch(r1, []int{1, 2})

	r2 := UniqValues(map[string]int{})
	is.Empty(r2)

	r3 := UniqValues(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3})
	is.ElementsMatch(r3, []int{1, 2, 3})

	r4 := UniqValues[string, int]()
	is.Empty(r4)

	r5 := UniqValues(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3})
	is.ElementsMatch(r5, []int{1, 2, 3})

	r6 := UniqValues(map[string]int{"foo": 1, "bar": 1}, map[string]int{"foo": 1, "bar": 3})
	is.ElementsMatch(r6, []int{1, 3})

	// check order
	r7 := UniqValues(map[string]int{"foo": 1}, map[string]int{"bar": 3})
	is.Equal([]int{1, 3}, r7)
}

func TestValueOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := ValueOr(map[string]int{"foo": 1}, "bar", 2)
	is.Equal(2, r1)

	r2 := ValueOr(map[string]int{"foo": 1}, "foo", 2)
	is.Equal(1, r2)
}

func TestPickBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})

	is.Equal(map[string]int{"foo": 1, "baz": 3}, r1)

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := PickBy(before, func(key string, value int) bool { return true })
	is.IsType(after, before, "type preserved")
}

func TestPickByKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz", "qux"})

	is.Equal(map[string]int{"foo": 1, "baz": 3}, r1)

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := PickByKeys(before, []string{"foobar", "baz"})
	is.IsType(after, before, "type preserved")
}

func TestPickByValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := PickByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

	is.Equal(map[string]int{"foo": 1, "baz": 3}, r1)

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

	is.Equal(map[string]int{"bar": 2}, r1)

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := PickBy(before, func(key string, value int) bool { return true })
	is.IsType(after, before, "type preserved")
}

func TestOmitByKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := OmitByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz", "qux"})

	is.Equal(map[string]int{"bar": 2}, r1)

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := OmitByKeys(before, []string{"foobar", "baz"})
	is.IsType(after, before, "type preserved")
}

func TestOmitByValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := OmitByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

	is.Equal(map[string]int{"bar": 2}, r1)

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := OmitByValues(before, []int{0, 3})
	is.IsType(after, before, "type preserved")
}

func TestEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Entries(map[string]int{"foo": 1, "bar": 2})
	is.ElementsMatch(r1, []Entry[string, int]{
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
	is.ElementsMatch(r1, []Entry[string, int]{
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
	is.Equal(1, r1["foo"])
	is.Equal(2, r1["bar"])
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
	is.Equal(3, r1["baz"])
	is.Equal(4, r1["qux"])
}

func TestInvert(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Invert(map[string]int{"a": 1, "b": 2})
	r2 := Invert(map[string]int{"a": 1, "b": 2, "c": 1})

	is.Equal(map[int]string{1: "a", 2: "b"}, r1)
	is.Len(r2, 2)
}

func TestAssign(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Assign(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})
	is.Equal(map[string]int{"a": 1, "b": 3, "c": 4}, result1)

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

	is.Len(result1, 3)
	is.Len(result2, 2)
	is.Empty(result3)
	is.Len(result4, 1)
	is.Len(result5, 2)

	is.PanicsWithValue("lo.ChunkEntries: size must be greater than 0", func() {
		ChunkEntries(map[string]int{"a": 1}, 0)
	})
	is.PanicsWithValue("lo.ChunkEntries: size must be greater than 0", func() {
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
	is.Equal(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, originalMap)
}

func TestMapKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapKeys(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x, _ int) string {
		return "Hello"
	})
	result2 := MapKeys(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(_, v int) string {
		return strconv.FormatInt(int64(v), 10)
	})

	is.Len(result1, 1)
	is.Equal(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}, result2)
}

func TestMapValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapValues(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x, _ int) string {
		return "Hello"
	})
	result2 := MapValues(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x, _ int) string {
		return strconv.FormatInt(int64(x), 10)
	})

	is.Equal(map[int]string{1: "Hello", 2: "Hello", 3: "Hello", 4: "Hello"}, result1)
	is.Equal(map[int]string{1: "1", 2: "2", 3: "3", 4: "4"}, result2)
}

func TestMapEntries(t *testing.T) {
	t.Parallel()

	t.Run("Normal", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := MapEntries(map[string]int{"foo": 1, "bar": 2},
			func(k string, v int) (string, int) {
				return k, v + 1
			})
		is.Equal(map[string]int{"foo": 2, "bar": 3}, r1)

		r2 := MapEntries(map[string]int{"foo": 1, "bar": 2},
			func(k string, v int) (string, string) {
				return k, k + strconv.Itoa(v)
			})
		is.Equal(map[string]string{"foo": "foo1", "bar": "bar2"}, r2)

		r3 := MapEntries(map[string]int{"foo": 1, "bar": 2},
			func(k string, v int) (string, string) {
				return k, strconv.Itoa(v) + k
			})
		is.Equal(map[string]string{"foo": "1foo", "bar": "2bar"}, r3)
	})

	t.Run("NoMutation", func(t *testing.T) {
		t.Parallel()

		r1 := map[string]int{"foo": 1, "bar": 2}
		MapEntries(r1, func(k string, v int) (string, string) {
			return k, strconv.Itoa(v) + "!!"
		})
		assert.Equal(t, map[string]int{"foo": 1, "bar": 2}, r1)
	})

	t.Run("EmptyInput", func(t *testing.T) {
		t.Parallel()

		r1 := MapEntries(map[string]int{},
			func(k string, v int) (string, string) {
				return k, strconv.Itoa(v) + "!!"
			})
		assert.Empty(t, r1)

		r2 := MapEntries(map[string]any{},
			func(k string, v any) (string, any) {
				return k, v
			})
		assert.Empty(t, r2)
	})

	t.Run("Identity", func(t *testing.T) {
		t.Parallel()

		r1 := MapEntries(map[string]int{"foo": 1, "bar": 2},
			func(k string, v int) (string, int) {
				return k, v
			})
		assert.Equal(t, map[string]int{"foo": 1, "bar": 2}, r1)

		r2 := MapEntries(map[string]any{"foo": 1, "bar": "2", "ccc": true},
			func(k string, v any) (string, any) {
				return k, v
			})
		assert.Equal(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, r2)
	})

	t.Run("ToConstantEntry", func(t *testing.T) {
		t.Parallel()

		r1 := MapEntries(map[string]any{"foo": 1, "bar": "2", "ccc": true},
			func(k string, v any) (string, any) {
				return "key", "value"
			})
		assert.Equal(t, map[string]any{"key": "value"}, r1)

		r2 := MapEntries(map[string]any{"foo": 1, "bar": "2", "ccc": true},
			func(k string, v any) (string, any) {
				return "b", 5
			})
		assert.Equal(t, map[string]any{"b": 5}, r2)
	})

	// // because using range over map, the order is not guaranteed
	// // this test is not deterministic
	// t.Run("OverlappingKeys", func(t *testing.T) {
	// 		t.Parallel()
	//
	// 	r1 := MapEntries(map[string]any{"foo": 1, "foo2": 2, "Foo": 2, "Foo2": "2", "bar": "2", "ccc": true},
	// 		func(k string, v any) (string, any) {
	// 			return string(k[0]), v
	// 		})
	// 	assert.Equal(t, map[string]any{"F": "2", "b": "2", "c": true, "f": 1}, r1)
	//
	// 	r2 := MapEntries(map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"},
	// 		func(k, v string) (string, string) {
	// 			return v, k
	// 		})
	// 	assert.Equal(t, map[string]string{"1": "foo", "2": "bar", "true": "ccc"}, r2)
	// })

	t.Run("NormalMappers", func(t *testing.T) {
		t.Parallel()

		r1 := MapEntries(map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"},
			func(k, v string) (string, string) {
				return k, k + v
			})
		assert.Equal(t, map[string]string{"Foo": "Foo2", "Foo2": "Foo22", "bar": "bar2", "ccc": "ccctrue", "foo": "foo1", "foo2": "foo22"}, r1)

		type myStruct struct {
			name string
			age  int
		}
		r2 := MapEntries(map[string]myStruct{"1-11-1": {name: "foo", age: 1}, "2-22-2": {name: "bar", age: 2}},
			func(k string, v myStruct) (string, string) {
				return v.name, k
			})
		assert.Equal(t, map[string]string{"bar": "2-22-2", "foo": "1-11-1"}, r2)
	})
}

func TestMapToSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MapToSlice(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k, v int) string {
		return fmt.Sprintf("%d_%d", k, v)
	})
	result2 := MapToSlice(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k, _ int) string {
		return strconv.FormatInt(int64(k), 10)
	})

	is.ElementsMatch(result1, []string{"1_5", "2_6", "3_7", "4_8"})
	is.ElementsMatch(result2, []string{"1", "2", "3", "4"})
}

func TestFilterMapToSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FilterMapToSlice(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k, v int) (string, bool) {
		return fmt.Sprintf("%d_%d", k, v), k%2 == 0
	})
	result2 := FilterMapToSlice(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k, _ int) (string, bool) {
		return strconv.FormatInt(int64(k), 10), k%2 == 0
	})

	is.ElementsMatch(result1, []string{"2_6", "4_8"})
	is.ElementsMatch(result2, []string{"2", "4"})
}

func TestFilterKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FilterKeys(map[int]string{1: "foo", 2: "bar", 3: "baz"}, func(k int, v string) bool {
		return v == "foo"
	})
	is.Equal([]int{1}, result1)

	result2 := FilterKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Empty(result2)
}

func TestFilterValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FilterValues(map[int]string{1: "foo", 2: "bar", 3: "baz"}, func(k int, v string) bool {
		return v == "foo"
	})
	is.Equal([]string{"foo"}, result1)

	result2 := FilterValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Empty(result2)
}

func BenchmarkAssign(b *testing.B) {
	counts := []int{32768, 1024, 128, 32, 2}

	allDifferentMap := func(b *testing.B, n int) []map[string]int {
		b.Helper()
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
		b.Helper()
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

		b.Run(strconv.Itoa(count), func(b *testing.B) {
			testCases := []struct {
				name string
				in   []map[string]int
			}{
				{"different", differentMap},
				{"same", sameMap},
			}

			for _, tc := range testCases {
				b.Run(tc.name, func(b *testing.B) {
					b.ResetTimer()
					for n := 0; n < b.N; n++ {
						result := Assign(tc.in...)
						_ = result
					}
				})
			}
		})
	}
}
