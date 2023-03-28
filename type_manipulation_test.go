package lo_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.ToPtr([]int{1, 2})

	is.Equal(*result1, []int{1, 2})
}

func TestEmptyableToPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Nil(lo.EmptyableToPtr(0))
	is.Nil(lo.EmptyableToPtr(""))
	is.Nil(lo.EmptyableToPtr[[]int](nil))
	is.Nil(lo.EmptyableToPtr[map[int]int](nil))
	is.Nil(lo.EmptyableToPtr[error](nil))

	is.Equal(*lo.EmptyableToPtr(42), 42)
	is.Equal(*lo.EmptyableToPtr("nonempty"), "nonempty")
	is.Equal(*lo.EmptyableToPtr([]int{}), []int{})
	is.Equal(*lo.EmptyableToPtr([]int{1, 2}), []int{1, 2})
	is.Equal(*lo.EmptyableToPtr(map[int]int{}), map[int]int{})
	is.Equal(*lo.EmptyableToPtr(assert.AnError), assert.AnError)
}

func TestFromPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	ptr := &str1

	is.Equal("foo", lo.FromPtr(ptr))
	is.Equal("", lo.FromPtr[string](nil))
	is.Equal(0, lo.FromPtr[int](nil))
	is.Nil(lo.FromPtr[*string](nil))
	is.EqualValues(ptr, lo.FromPtr(&ptr))
}

func TestFromPtrOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	const fallbackStr = "fallback"
	str := "foo"
	ptrStr := &str

	const fallbackInt = -1
	i := 9
	ptrInt := &i

	is.Equal(str, lo.FromPtrOr(ptrStr, fallbackStr))
	is.Equal(fallbackStr, lo.FromPtrOr(nil, fallbackStr))
	is.Equal(i, lo.FromPtrOr(ptrInt, fallbackInt))
	is.Equal(fallbackInt, lo.FromPtrOr(nil, fallbackInt))
}

func TestToSlicePtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := lo.ToSlicePtr([]string{str1, str2})

	is.Equal(result1, []*string{&str1, &str2})
}

func TestToAnySlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in1 := []int{0, 1, 2, 3}
	in2 := []int{}
	out1 := lo.ToAnySlice(in1)
	out2 := lo.ToAnySlice(in2)

	is.Equal([]any{0, 1, 2, 3}, out1)
	is.Equal([]any{}, out2)
}

func TestFromAnySlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.NotPanics(func() {
		out1, ok1 := lo.FromAnySlice[string]([]any{"foobar", 42})
		out2, ok2 := lo.FromAnySlice[string]([]any{"foobar", "42"})

		is.Equal([]string{}, out1)
		is.False(ok1)
		is.Equal([]string{"foobar", "42"}, out2)
		is.True(ok2)
	})
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	//nolint:unused
	type test struct{}

	is.Empty(lo.Empty[string]())
	is.Empty(lo.Empty[int64]())
	is.Empty(lo.Empty[test]())
	is.Empty(lo.Empty[chan string]())
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	//nolint:unused
	type test struct {
		foobar string
	}

	is.True(lo.IsEmpty(""))
	is.False(lo.IsEmpty("foo"))
	is.True(lo.IsEmpty[int64](0))
	is.False(lo.IsEmpty[int64](42))
	is.True(lo.IsEmpty(test{foobar: ""}))
	is.False(lo.IsEmpty(test{foobar: "foo"}))
}

func TestIsNotEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	//nolint:unused
	type test struct {
		foobar string
	}

	is.False(lo.IsNotEmpty(""))
	is.True(lo.IsNotEmpty("foo"))
	is.False(lo.IsNotEmpty[int64](0))
	is.True(lo.IsNotEmpty[int64](42))
	is.False(lo.IsNotEmpty(test{foobar: ""}))
	is.True(lo.IsNotEmpty(test{foobar: "foo"}))
}

func TestCoalesce(t *testing.T) {
	t.Parallel()
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

	result1, ok1 := lo.Coalesce[int]()
	result2, ok2 := lo.Coalesce(3)
	result3, ok3 := lo.Coalesce(nil, nilStr)
	result4, ok4 := lo.Coalesce(nilStr, str1)
	result5, ok5 := lo.Coalesce(nilStr, str1, str2)
	result6, ok6 := lo.Coalesce(str1, str2, nilStr)
	result7, ok7 := lo.Coalesce(0, 1, 2, 3)
	result8, ok8 := lo.Coalesce(zeroStruct)
	result9, ok9 := lo.Coalesce(zeroStruct, struct1)
	result10, ok10 := lo.Coalesce(zeroStruct, struct1, struct2)

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
