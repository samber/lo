package lo

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	is := assert.New(t)

	r1 := Keys(map[string]int{"foo": 1, "bar": 2})
	sort.Strings(r1)

	is.Equal(r1, []string{"bar", "foo"})
}

func TestValues(t *testing.T) {
	is := assert.New(t)

	r1 := Values(map[string]int{"foo": 1, "bar": 2})
	sort.Ints(r1)

	is.Equal(r1, []int{1, 2})
}

func TestPickBy(t *testing.T) {
	is := assert.New(t)

	r1 := PickBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})
}

func TestPickByKeys(t *testing.T) {
	is := assert.New(t)

	r1 := PickByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})
}

func TestPickByValues(t *testing.T) {
	is := assert.New(t)

	r1 := PickByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})
}

func TestOmitBy(t *testing.T) {
	is := assert.New(t)

	r1 := OmitBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})

	is.Equal(r1, map[string]int{"bar": 2})
}

func TestOmitByKeys(t *testing.T) {
	is := assert.New(t)

	r1 := OmitByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})

	is.Equal(r1, map[string]int{"bar": 2})
}

func TestOmitByValues(t *testing.T) {
	is := assert.New(t)

	r1 := OmitByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

	is.Equal(r1, map[string]int{"bar": 2})
}

func TestEntries(t *testing.T) {
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

func TestFromEntries(t *testing.T) {
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

func TestInvert(t *testing.T) {
	is := assert.New(t)

	r1 := Invert(map[string]int{"a": 1, "b": 2})
	r2 := Invert(map[string]int{"a": 1, "b": 2, "c": 1})

	is.Len(r1, 2)
	is.EqualValues(map[int]string{1: "a", 2: "b"}, r1)
	is.Len(r2, 2)
}

func TestAssign(t *testing.T) {
	is := assert.New(t)

	result1 := Assign(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})

	is.Len(result1, 3)
	is.Equal(result1, map[string]int{"a": 1, "b": 3, "c": 4})
}

func TestMapKeys(t *testing.T) {
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

func TestMapToSlice(t *testing.T) {
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

func mapEntriesTest[I any, E any](t *testing.T, in map[string]I, iteratee func(entry Entry[string, I]) Entry[string, E], expected map[string]E) {
	is := assert.New(t)
	result := MapEntries(in, iteratee)
	is.Equal(result, expected)
}

func TestMapEntries(t *testing.T) {

	mapEntriesTest(t, map[string]int{"foo": 1, "bar": 2}, func(entry Entry[string, int]) Entry[string, int] {
		return Entry[string, int]{entry.Key, entry.Value + 1}
	}, map[string]int{"foo": 2, "bar": 3})
	mapEntriesTest(t, map[string]int{"foo": 1, "bar": 2}, func(entry Entry[string, int]) Entry[string, string] {
		return Entry[string, string]{entry.Key, entry.Key + strconv.Itoa(entry.Value)}
	}, map[string]string{"foo": "foo1", "bar": "bar2"})
	mapEntriesTest(t, map[string]int{"foo": 1, "bar": 2}, func(entry Entry[string, int]) Entry[string, string] {
		return Entry[string, string]{entry.Key, strconv.Itoa(entry.Value) + entry.Key}
	}, map[string]string{"foo": "1foo", "bar": "2bar"})

	// NoMutation
	{
		is := assert.New(t)
		r1 := map[string]int{"foo": 1, "bar": 2}
		MapEntries(r1, func(entry Entry[string, int]) Entry[string, string] {
			return Entry[string, string]{
				Key:   entry.Key,
				Value: strconv.Itoa(entry.Value) + "!!",
			}
		})
		is.Equal(r1, map[string]int{"foo": 1, "bar": 2})
	}
	// EmptyInput
	{
		mapEntriesTest(t, map[string]int{}, func(entry Entry[string, int]) Entry[string, string] {
			return Entry[string, string]{
				Key:   entry.Key,
				Value: strconv.Itoa(entry.Value) + "!!",
			}
		}, map[string]string{})

		mapEntriesTest(t, map[string]any{}, func(entry Entry[string, any]) Entry[string, any] {
			return entry
		}, map[string]any{})
	}
	// Identity
	{
		mapEntriesTest(t, map[string]int{"foo": 1, "bar": 2}, func(entry Entry[string, int]) Entry[string, int] {
			return entry
		}, map[string]int{"foo": 1, "bar": 2})
		mapEntriesTest(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(entry Entry[string, any]) Entry[string, any] {
			return entry
		}, map[string]any{"foo": 1, "bar": "2", "ccc": true})
	}
	// ToConstantEntry
	{
		mapEntriesTest(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(entry Entry[string, any]) Entry[string, any] {
			return Entry[string, any]{Key: "key", Value: "value"}
		}, map[string]any{"key": "value"})
		mapEntriesTest(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, func(entry Entry[string, any]) Entry[string, any] {
			return Entry[string, any]{Key: "b", Value: 5}
		}, map[string]any{"b": 5})
	}

	//// OverlappingKeys
	//// because using range over map, the order is not guaranteed
	//// this test is not deterministic
	//{
	//	mapEntriesTest(t, map[string]any{"foo": 1, "foo2": 2, "Foo": 2, "Foo2": "2", "bar": "2", "ccc": true}, func(entry Entry[string, any]) Entry[string, any] {
	//		return Entry[string, any]{Key: string(entry.Key[0]), Value: entry.Value}
	//	}, map[string]any{"F": "2", "b": "2", "c": true, "f": 2})
	//	mapEntriesTest(t, map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"}, func(entry Entry[string, string]) Entry[string, string] {
	//		return Entry[string, string]{Key: entry.Value, Value: entry.Key}
	//	}, map[string]string{"1": "foo", "2": "bar", "true": "ccc"})
	//}
	// NormalMappers
	{
		mapEntriesTest(t, map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"}, func(entry Entry[string, string]) Entry[string, string] {
			return Entry[string, string]{Key: entry.Key, Value: entry.Key + entry.Value}
		}, map[string]string{"Foo": "Foo2", "Foo2": "Foo22", "bar": "bar2", "ccc": "ccctrue", "foo": "foo1", "foo2": "foo22"})

		mapEntriesTest(t, map[string]struct {
			name string
			age  int
		}{"1-11-1": {name: "foo", age: 1}, "2-22-2": {name: "bar", age: 2}}, func(entry Entry[string, struct {
			name string
			age  int
		}]) Entry[string, string] {
			return Entry[string, string]{Key: entry.Value.name, Value: entry.Key}
		}, map[string]string{"bar": "2-22-2", "foo": "1-11-1"})
	}
}
