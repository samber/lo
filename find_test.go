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

	result1 := IndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := IndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}

func TestLastIndexOf(t *testing.T) {
	is := assert.New(t)

	result1 := LastIndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := LastIndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 4)
	is.Equal(result2, -1)
}

func TestFind(t *testing.T) {
	is := assert.New(t)

	result1, ok1 := Find[string]([]string{"a", "b", "c", "d"}, func(i string) bool {
		return i == "b"
	})
	result2, ok2 := Find[string]([]string{"foobar"}, func(i string) bool {
		return i == "b"
	})

	is.Equal(ok1, true)
	is.Equal(result1, "b")
	is.Equal(ok2, false)
	is.Equal(result2, "")
}

func TestFindIndexOf(t *testing.T) {
	is := assert.New(t)

	item1, index1, ok1 := FindIndexOf[string]([]string{"a", "b", "c", "d", "b"}, func(i string) bool {
		return i == "b"
	})
	item2, index2, ok2 := FindIndexOf[string]([]string{"foobar"}, func(i string) bool {
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

	item1, index1, ok1 := FindLastIndexOf[string]([]string{"a", "b", "c", "d", "b"}, func(i string) bool {
		return i == "b"
	})
	item2, index2, ok2 := FindLastIndexOf[string]([]string{"foobar"}, func(i string) bool {
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

	result1 := FindOrElse[string]([]string{"a", "b", "c", "d"}, "x", func(i string) bool {
		return i == "b"
	})
	result2 := FindOrElse[string]([]string{"foobar"}, "x", func(i string) bool {
		return i == "b"
	})

	is.Equal(result1, "b")
	is.Equal(result2, "x")
}

func TestMin(t *testing.T) {
	is := assert.New(t)

	result1 := Min[int]([]int{1, 2, 3})
	result2 := Min[int]([]int{3, 2, 1})
	result3 := Min[int]([]int{})

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 0)
}

func TestMinBy(t *testing.T) {
	is := assert.New(t)

	result1 := MinBy[string]([]string{"s1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result2 := MinBy[string]([]string{"string1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result3 := MinBy[string]([]string{}, func(item string, min string) bool {
		return len(item) < len(min)
	})

	is.Equal(result1, "s1")
	is.Equal(result2, "s3")
	is.Equal(result3, "")
}

func TestMax(t *testing.T) {
	is := assert.New(t)

	result1 := Max[int]([]int{1, 2, 3})
	result2 := Max[int]([]int{3, 2, 1})
	result3 := Max[int]([]int{})

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, 0)
}

func TestMaxBy(t *testing.T) {
	is := assert.New(t)

	result1 := MaxBy[string]([]string{"s1", "string2", "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result2 := MaxBy[string]([]string{"string1", "string2", "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result3 := MaxBy[string]([]string{}, func(item string, max string) bool {
		return len(item) > len(max)
	})

	is.Equal(result1, "string2")
	is.Equal(result2, "string1")
	is.Equal(result3, "")
}

func TestLast(t *testing.T) {
	is := assert.New(t)

	result1, err1 := Last[int]([]int{1, 2, 3})
	result2, err2 := Last[int]([]int{})

	is.Equal(result1, 3)
	is.Equal(err1, nil)
	is.Equal(result2, 0)
	is.Equal(err2, fmt.Errorf("last: cannot extract the last element of an empty slice"))
}

func TestNth(t *testing.T) {
	is := assert.New(t)

	result1, err1 := Nth[int]([]int{0, 1, 2, 3}, 2)
	result2, err2 := Nth[int]([]int{0, 1, 2, 3}, -2)
	result3, err3 := Nth[int]([]int{0, 1, 2, 3}, 42)
	result4, err4 := Nth[int]([]int{}, 0)
	result5, err5 := Nth[int]([]int{42}, 0)

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

	result1 := Sample[string]([]string{"a", "b", "c"})
	result2 := Sample[string]([]string{})

	is.True(Contains[string]([]string{"a", "b", "c"}, result1))
	is.Equal(result2, "")
}

func TestSamples(t *testing.T) {
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Samples[string]([]string{"a", "b", "c"}, 3)
	result2 := Samples[string]([]string{}, 3)

	sort.Strings(result1)

	is.Equal(result1, []string{"a", "b", "c"})
	is.Equal(result2, []string{})
}

func TestCoalesce(t *testing.T) {
	is := assert.New(t)

	newStr := func(v string) *string { return &v }
	var nilStr *string
	str1 := newStr("str1")
	str2 := newStr("str2")

	type structType struct {
		field1 int
		field2 float64
	}
	var zeroStruct structType
	struct1 := structType{1, 1.0}
	struct2 := structType{2, 2.0}

	result1, ok1 := Coalesce[int]()
	result2, ok2 := Coalesce(3)
	result3, ok3 := Coalesce[*string](nil, nilStr)
	result4, ok4 := Coalesce(nilStr, str1)
	result5, ok5 := Coalesce(nilStr, str1, str2)
	result6, ok6 := Coalesce(str1, str2, nilStr)
	result7, ok7 := Coalesce(0, 1, 2, 3)
	result8, ok8 := Coalesce(zeroStruct)
	result9, ok9 := Coalesce(zeroStruct, struct1)
	result10, ok10 := Coalesce(zeroStruct, struct1, struct2)

	is.Equal(0, result1)
	is.False(ok1)

	is.Equal(3, result2)
	is.True(ok2)

	is.Nil(result3)
	is.False(ok3)

	is.Equal(str1, result4)
	is.True(ok4)

	is.Equal(str1, result5)
	is.True(ok5)

	is.Equal(str1, result6)
	is.True(ok6)

	is.Equal(result7, 1)
	is.True(ok7)

	is.Equal(result8, zeroStruct)
	is.False(ok8)

	is.Equal(result9, struct1)
	is.True(ok9)

	is.Equal(result10, struct1)
	is.True(ok10)
}
