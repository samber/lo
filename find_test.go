package lo

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 4)
	is.Equal(result2, -1)
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1, ok1 := Find([]string{"a", "b", "c", "d"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})

	result2, ok2 := Find([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(ok1, true)
	is.Equal(result1, "b")
	is.Equal(ok2, false)
	is.Equal(result2, "")
}

func TestFindIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d", "b"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 1)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindLastIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"b", "d", "c", "b", "a"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindLastIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 4)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindOrElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1 := FindOrElse([]string{"a", "b", "c", "d"}, "x", func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})
	result2 := FindOrElse([]string{"foobar"}, "x", func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(result1, "b")
	is.Equal(result2, "x")
}

func TestFindKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 2)
	is.Equal("bar", result1)
	is.True(ok1)

	result2, ok2 := FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 42)
	is.Equal("", result2)
	is.False(ok2)

	type test struct {
		foobar string
	}

	result3, ok3 := FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"foo"})
	is.Equal("foo", result3)
	is.True(ok3)

	result4, ok4 := FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"hello world"})
	is.Equal("", result4)
	is.False(ok4)
}

func TestFindKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return k == "foo"
	})
	is.Equal("foo", result1)
	is.True(ok1)

	result2, ok2 := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Equal("", result2)
	is.False(ok2)
}

func TestFindUniques(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniques([]int{1, 2, 3})

	is.Equal(3, len(result1))
	is.Equal([]int{1, 2, 3}, result1)

	result2 := FindUniques([]int{1, 2, 2, 3, 1, 2})

	is.Equal(1, len(result2))
	is.Equal([]int{3}, result2)

	result3 := FindUniques([]int{1, 2, 2, 1})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	result4 := FindUniques([]int{})

	is.Equal(0, len(result4))
	is.Equal([]int{}, result4)
}

func TestFindUniquesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniquesBy([]int{0, 1, 2}, func(i int) int {
		return i % 3
	})

	is.Equal(3, len(result1))
	is.Equal([]int{0, 1, 2}, result1)

	result2 := FindUniquesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 3
	})

	is.Equal(1, len(result2))
	is.Equal([]int{2}, result2)

	result3 := FindUniquesBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	result4 := FindUniquesBy([]int{}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result4))
	is.Equal([]int{}, result4)
}

func TestFindDuplicates(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicates([]int{1, 2, 2, 1, 2, 3})

	is.Equal(2, len(result1))
	is.Equal([]int{1, 2}, result1)

	result2 := FindDuplicates([]int{1, 2, 3})

	is.Equal(0, len(result2))
	is.Equal([]int{}, result2)

	result3 := FindDuplicates([]int{})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)
}

func TestFindDuplicatesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicatesBy([]int{3, 4, 5, 6, 7}, func(i int) int {
		return i % 3
	})

	is.Equal(2, len(result1))
	is.Equal([]int{3, 4}, result1)

	result2 := FindDuplicatesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 5
	})

	is.Equal(0, len(result2))
	is.Equal([]int{}, result2)

	result3 := FindDuplicatesBy([]int{}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)
}

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Min([]int{1, 2, 3})
	result2 := Min([]int{3, 2, 1})
	result3 := Min([]int{})

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 0)
}

func TestMinValue(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MinValue(map[string]int{"1": 1, "2": 2, "3": 3})
	result2 := MinValue(map[string]int{"3": 3, "2": 2, "1": 1})
	result3 := MinValue(map[string]int{})

	is.Equal(result1.Value, 1)
	is.Equal(result2.Value, 1)
	is.Equal(result3.Value, 0)
}

func TestMinBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MinBy([]string{"s1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result2 := MinBy([]string{"string1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result3 := MinBy([]string{}, func(item string, min string) bool {
		return len(item) < len(min)
	})

	is.Equal(result1, "s1")
	is.Equal(result2, "s3")
	is.Equal(result3, "")
}

func TestMinValueBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MinValueBy(map[int]string{1: "s1", 2: "string2", 3: "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result2 := MinValueBy(map[int]string{1: "string1", 2: "string2", 3: "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result3 := MinValueBy(map[int]string{}, func(item string, min string) bool {
		return len(item) < len(min)
	})

	// map can't guarantee order
	is.Contains([]string{"s1", "s3"}, result1.Value)
	is.Equal(result2.Value, "s3")
	is.Equal(result3.Value, "")
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Max([]int{1, 2, 3})
	result2 := Max([]int{3, 2, 1})
	result3 := Max([]int{})

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, 0)
}

func TestMaxValue(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MaxValue(map[string]int{"1": 1, "2": 2, "3": 3})
	result2 := MaxValue(map[string]int{"3": 3, "2": 2, "1": 1})
	result3 := MaxValue(map[string]int{})

	is.Equal(result1.Value, 3)
	is.Equal(result2.Value, 3)
	is.Equal(result3.Value, 0)
}

func TestMaxBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MaxBy([]string{"s1", "string2", "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result2 := MaxBy([]string{"string1", "string2", "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result3 := MaxBy([]string{}, func(item string, max string) bool {
		return len(item) > len(max)
	})

	is.Equal(result1, "string2")
	is.Equal(result2, "string1")
	is.Equal(result3, "")
}

func TestMaxValueBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MaxValueBy(map[int]string{1: "s1", 2: "string2", 3: "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result2 := MaxValueBy(map[int]string{1: "string1", 2: "string2", 3: "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result3 := MaxValueBy(map[int]string{}, func(item string, max string) bool {
		return len(item) > len(max)
	})

	is.Equal(result1.Value, "string2")
	// map can't guarantee order
	is.Contains([]string{"string2", "string1"}, result2.Value)
	is.Equal(result3.Value, "")
}

func TestAnyKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := AnyKey(map[int]int{1: 1, 2: 2, 3: 3})
	result2, err2 := AnyKey(map[int]int{})

	is.Contains([]int{1, 2, 3}, result1)
	is.Equal(err1, nil)
	is.Equal(result2, 0)
	is.Equal(err2, fmt.Errorf("AnyKey: cannot extract the first key of an empty map"))
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := Last([]int{1, 2, 3})
	result2, err2 := Last([]int{})

	is.Equal(result1, 3)
	is.Equal(err1, nil)
	is.Equal(result2, 0)
	is.Equal(err2, fmt.Errorf("last: cannot extract the last element of an empty slice"))
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := Nth([]int{0, 1, 2, 3}, 2)
	result2, err2 := Nth([]int{0, 1, 2, 3}, -2)
	result3, err3 := Nth([]int{0, 1, 2, 3}, 42)
	result4, err4 := Nth([]int{}, 0)
	result5, err5 := Nth([]int{42}, 0)
	result6, err6 := Nth([]int{42}, -1)

	is.Equal(result1, 2)
	is.Equal(err1, nil)
	is.Equal(result2, 2)
	is.Equal(err2, nil)
	is.Equal(result3, 0)
	is.Equal(err3, fmt.Errorf("nth: 42 out of slice bounds"))
	is.Equal(result4, 0)
	is.Equal(err4, fmt.Errorf("nth: 0 out of slice bounds"))
	is.Equal(result5, 42)
	is.Equal(err5, nil)
	is.Equal(result6, 42)
	is.Equal(err6, nil)
}

func TestSample(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Sample([]string{"a", "b", "c"})
	result2 := Sample([]string{})

	is.True(Contains([]string{"a", "b", "c"}, result1))
	is.Equal(result2, "")
}

func TestSamples(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Samples([]string{"a", "b", "c"}, 3)
	result2 := Samples([]string{}, 3)

	sort.Strings(result1)

	is.Equal(result1, []string{"a", "b", "c"})
	is.Equal(result2, []string{})
}
