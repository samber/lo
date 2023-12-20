package lo_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := lo.IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := lo.LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 4)
	is.Equal(result2, -1)
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1, ok1 := lo.Find([]string{"a", "b", "c", "d"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})

	result2, ok2 := lo.Find([]string{"foobar"}, func(item string) bool {
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
	item1, index1, ok1 := lo.FindIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d", "b"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := lo.FindIndexOf([]string{"foobar"}, func(item string) bool {
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
	item1, index1, ok1 := lo.FindLastIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"b", "d", "c", "b", "a"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := lo.FindLastIndexOf([]string{"foobar"}, func(item string) bool {
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
	result1 := lo.FindOrElse([]string{"a", "b", "c", "d"}, "x", func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})
	result2 := lo.FindOrElse([]string{"foobar"}, "x", func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(result1, "b")
	is.Equal(result2, "x")
}

func TestFindKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := lo.FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 2)
	is.Equal("bar", result1)
	is.True(ok1)

	result2, ok2 := lo.FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 42)
	is.Equal("", result2)
	is.False(ok2)

	type test struct {
		foobar string
	}

	result3, ok3 := lo.FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"foo"})
	is.Equal("foo", result3)
	is.True(ok3)

	result4, ok4 := lo.FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"hello world"})
	is.Equal("", result4)
	is.False(ok4)
}

func TestFindKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := lo.FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return k == "foo"
	})
	is.Equal("foo", result1)
	is.True(ok1)

	result2, ok2 := lo.FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Equal("", result2)
	is.False(ok2)
}

func TestFindUniques(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.FindUniques([]int{1, 2, 3})

	is.Equal(3, len(result1))
	is.Equal([]int{1, 2, 3}, result1)

	result2 := lo.FindUniques([]int{1, 2, 2, 3, 1, 2})

	is.Equal(1, len(result2))
	is.Equal([]int{3}, result2)

	result3 := lo.FindUniques([]int{1, 2, 2, 1})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	result4 := lo.FindUniques([]int{})

	is.Equal(0, len(result4))
	is.Equal([]int{}, result4)
}

func TestFindUniquesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.FindUniquesBy([]int{0, 1, 2}, func(i int) int {
		return i % 3
	})

	is.Equal(3, len(result1))
	is.Equal([]int{0, 1, 2}, result1)

	result2 := lo.FindUniquesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 3
	})

	is.Equal(1, len(result2))
	is.Equal([]int{2}, result2)

	result3 := lo.FindUniquesBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	result4 := lo.FindUniquesBy([]int{}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result4))
	is.Equal([]int{}, result4)
}

func TestFindDuplicates(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.FindDuplicates([]int{1, 2, 2, 1, 2, 3})

	is.Equal(2, len(result1))
	is.Equal([]int{1, 2}, result1)

	result2 := lo.FindDuplicates([]int{1, 2, 3})

	is.Equal(0, len(result2))
	is.Equal([]int{}, result2)

	result3 := lo.FindDuplicates([]int{})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)
}

func TestFindDuplicatesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.FindDuplicatesBy([]int{3, 4, 5, 6, 7}, func(i int) int {
		return i % 3
	})

	is.Equal(2, len(result1))
	is.Equal([]int{3, 4}, result1)

	result2 := lo.FindDuplicatesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 5
	})

	is.Equal(0, len(result2))
	is.Equal([]int{}, result2)

	result3 := lo.FindDuplicatesBy([]int{}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)
}

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Min([]int{1, 2, 3})
	result2 := lo.Min([]int{3, 2, 1})
	result3 := lo.Min([]int{})

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 0)
}

func TestMinBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.MinBy([]string{"s1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result2 := lo.MinBy([]string{"string1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result3 := lo.MinBy([]string{}, func(item string, min string) bool {
		return len(item) < len(min)
	})

	is.Equal(result1, "s1")
	is.Equal(result2, "s3")
	is.Equal(result3, "")
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Max([]int{1, 2, 3})
	result2 := lo.Max([]int{3, 2, 1})
	result3 := lo.Max([]int{})

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, 0)
}

func TestMaxBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.MaxBy([]string{"s1", "string2", "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result2 := lo.MaxBy([]string{"string1", "string2", "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result3 := lo.MaxBy([]string{}, func(item string, max string) bool {
		return len(item) > len(max)
	})

	is.Equal(result1, "string2")
	is.Equal(result2, "string1")
	is.Equal(result3, "")
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := lo.Last([]int{1, 2, 3})
	result2, err2 := lo.Last([]int{})

	is.Equal(result1, 3)
	is.Equal(err1, nil)
	is.Equal(result2, 0)
	is.Equal(err2, fmt.Errorf("last: cannot extract the last element of an empty slice"))
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := lo.Nth([]int{0, 1, 2, 3}, 2)
	result2, err2 := lo.Nth([]int{0, 1, 2, 3}, -2)
	result3, err3 := lo.Nth([]int{0, 1, 2, 3}, 42)
	result4, err4 := lo.Nth([]int{}, 0)
	result5, err5 := lo.Nth([]int{42}, 0)
	result6, err6 := lo.Nth([]int{42}, -1)

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

	result1 := lo.Sample([]string{"a", "b", "c"})
	result2 := lo.Sample([]string{})

	is.True(lo.Contains([]string{"a", "b", "c"}, result1))
	is.Equal(result2, "")
}

func TestSamples(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := lo.Samples([]string{"a", "b", "c"}, 3)
	result2 := lo.Samples([]string{}, 3)

	sort.Strings(result1)

	is.Equal(result1, []string{"a", "b", "c"})
	is.Equal(result2, []string{})
}
