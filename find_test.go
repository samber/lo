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
	is := assert.New(t)

	result1 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}

func TestLastIndexOf(t *testing.T) {
	is := assert.New(t)

	result1 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 4)
	is.Equal(result2, -1)
}

func TestFind(t *testing.T) {
	is := assert.New(t)

	result1, ok1 := Find([]string{"a", "b", "c", "d"}, func(i string) bool {
		return i == "b"
	})
	result2, ok2 := Find([]string{"foobar"}, func(i string) bool {
		return i == "b"
	})

	is.Equal(ok1, true)
	is.Equal(result1, "b")
	is.Equal(ok2, false)
	is.Equal(result2, "")
}

func TestFindIndexOf(t *testing.T) {
	is := assert.New(t)

	item1, index1, ok1 := FindIndexOf([]string{"a", "b", "c", "d", "b"}, func(i string) bool {
		return i == "b"
	})
	item2, index2, ok2 := FindIndexOf([]string{"foobar"}, func(i string) bool {
		return i == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 1)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindLastIndexOf(t *testing.T) {
	is := assert.New(t)

	item1, index1, ok1 := FindLastIndexOf([]string{"a", "b", "c", "d", "b"}, func(i string) bool {
		return i == "b"
	})
	item2, index2, ok2 := FindLastIndexOf([]string{"foobar"}, func(i string) bool {
		return i == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 4)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindOrElse(t *testing.T) {
	is := assert.New(t)

	result1 := FindOrElse([]string{"a", "b", "c", "d"}, "x", func(i string) bool {
		return i == "b"
	})
	result2 := FindOrElse([]string{"foobar"}, "x", func(i string) bool {
		return i == "b"
	})

	is.Equal(result1, "b")
	is.Equal(result2, "x")
}

func TestMin(t *testing.T) {
	is := assert.New(t)

	result1 := Min([]int{1, 2, 3})
	result2 := Min([]int{3, 2, 1})
	result3 := Min([]int{})

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 0)
}

func TestMinBy(t *testing.T) {
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

func TestMax(t *testing.T) {
	is := assert.New(t)

	result1 := Max([]int{1, 2, 3})
	result2 := Max([]int{3, 2, 1})
	result3 := Max([]int{})

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, 0)
}

func TestMaxBy(t *testing.T) {
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

func TestLast(t *testing.T) {
	is := assert.New(t)

	result1, err1 := Last([]int{1, 2, 3})
	result2, err2 := Last([]int{})

	is.Equal(result1, 3)
	is.Equal(err1, nil)
	is.Equal(result2, 0)
	is.Equal(err2, fmt.Errorf("last: cannot extract the last element of an empty slice"))
}

func TestNth(t *testing.T) {
	is := assert.New(t)

	result1, err1 := Nth([]int{0, 1, 2, 3}, 2)
	result2, err2 := Nth([]int{0, 1, 2, 3}, -2)
	result3, err3 := Nth([]int{0, 1, 2, 3}, 42)
	result4, err4 := Nth([]int{}, 0)
	result5, err5 := Nth([]int{42}, 0)

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
}

func TestSample(t *testing.T) {
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Sample([]string{"a", "b", "c"})
	result2 := Sample([]string{})

	is.True(Contains([]string{"a", "b", "c"}, result1))
	is.Equal(result2, "")
}

func TestSamples(t *testing.T) {
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Samples([]string{"a", "b", "c"}, 3)
	result2 := Samples([]string{}, 3)

	sort.Strings(result1)

	is.Equal(result1, []string{"a", "b", "c"})
	is.Equal(result2, []string{})
}
