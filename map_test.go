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

	result1 := MapToSlice(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x int, v int) string {
		return fmt.Sprintf("%d_%d", x, v)
	})
	result2 := MapToSlice(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x int, _ int) string {
		return strconv.FormatInt(int64(x), 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.ElementsMatch(result1, []string{"1_1", "2_2", "3_3", "4_4"})
	is.ElementsMatch(result2, []string{"1", "2", "3", "4"})
}
