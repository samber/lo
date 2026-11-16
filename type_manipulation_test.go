package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNil(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var s *string
	var i *int
	var b *bool
	var ifaceWithNilValue any = (*string)(nil) //nolint:staticcheck

	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{name: "zero int", input: 0, expected: false},
		{name: "zero struct", input: struct{}{}, expected: false},
		{name: "nil *string", input: s, expected: true},
		{name: "nil *int", input: i, expected: true},
		{name: "nil *bool", input: b, expected: true},
		{name: "interface wrapping nil pointer", input: ifaceWithNilValue, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			is.Equal(tt.expected, IsNil(tt.input))
		})
	}

	is.False(ifaceWithNilValue == nil) //nolint:staticcheck,testifylint
}

func TestIsNotNil(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	var s *string
	var i *int
	var b *bool
	var ifaceWithNilValue any = (*string)(nil) //nolint:staticcheck

	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{name: "zero int", input: 0, expected: true},
		{name: "zero struct", input: struct{}{}, expected: true},
		{name: "nil *string", input: s, expected: false},
		{name: "nil *int", input: i, expected: false},
		{name: "nil *bool", input: b, expected: false},
		{name: "interface wrapping nil pointer", input: ifaceWithNilValue, expected: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)
			is.Equal(tt.expected, IsNotNil(tt.input))
		})
	}

	is.True(ifaceWithNilValue != nil) //nolint:staticcheck,testifylint
}

func TestToPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ToPtr([]int{1, 2})

	is.Equal([]int{1, 2}, *result1)
}

func TestNil(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	nilFloat64 := Nil[float64]()
	var expNilFloat64 *float64

	nilString := Nil[string]()
	var expNilString *string

	is.Equal(expNilFloat64, nilFloat64)
	is.Nil(nilFloat64)
	is.NotEqual(nil, nilFloat64) //nolint:testifylint

	is.Equal(expNilString, nilString)
	is.Nil(nilString)
	is.NotEqual(nil, nilString) //nolint:testifylint

	is.NotEqual(nilString, nilFloat64)
}

func TestEmptyableToPtr(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Nil(EmptyableToPtr(0))
		is.Equal(42, *EmptyableToPtr(42))
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Nil(EmptyableToPtr(""))
		is.Equal("nonempty", *EmptyableToPtr("nonempty"))
	})

	t.Run("slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Nil(EmptyableToPtr[[]int](nil))
		is.Empty(*EmptyableToPtr([]int{}))
		is.Equal([]int{1, 2}, *EmptyableToPtr([]int{1, 2}))
	})

	t.Run("map", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Nil(EmptyableToPtr[map[int]int](nil))
		is.Empty(*EmptyableToPtr(map[int]int{}))
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Nil(EmptyableToPtr[error](nil))
		is.Equal(assert.AnError, *EmptyableToPtr(assert.AnError))
	})
}

func TestFromPtr(t *testing.T) {
	t.Parallel()

	str1 := "foo"
	ptr := &str1

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Equal("foo", FromPtr(ptr))
		is.Empty(FromPtr[string](nil))
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Zero(FromPtr[int](nil))
	})

	t.Run("*string", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Nil(FromPtr[*string](nil))
		is.Equal(ptr, FromPtr(&ptr))
	})
}

func TestFromPtrOr(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		const fallbackStr = "fallback"
		str := "foo"
		ptrStr := &str

		is.Equal(str, FromPtrOr(ptrStr, fallbackStr))
		is.Equal(fallbackStr, FromPtrOr(nil, fallbackStr))
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		const fallbackInt = -1
		i := 9
		ptrInt := &i

		is.Equal(i, FromPtrOr(ptrInt, fallbackInt))
		is.Equal(fallbackInt, FromPtrOr(nil, fallbackInt))
	})
}

func TestToSlicePtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := ToSlicePtr([]string{str1, str2})

	is.Equal([]*string{&str1, &str2}, result1)
}

func TestFromSlicePtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := FromSlicePtr([]*string{&str1, &str2, nil})

	is.Equal([]string{str1, str2, ""}, result1)
}

func TestFromSlicePtrOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := FromSlicePtrOr([]*string{&str1, &str2, nil}, "fallback")

	is.Equal([]string{str1, str2, "fallback"}, result1)
}

func TestToAnySlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		expected []any
	}{
		{name: "non-empty slice", input: []int{0, 1, 2, 3}, expected: []any{0, 1, 2, 3}},
		{name: "empty slice", input: []int{}, expected: []any{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			is.Equal(tt.expected, ToAnySlice(tt.input))
		})
	}
}

func TestFromAnySlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       []any
		expectedOut []string
		expectedOk  bool
	}{
		{name: "mismatched types", input: []any{"foobar", 42}, expectedOut: []string{}, expectedOk: false},
		{name: "matching types", input: []any{"foobar", "42"}, expectedOut: []string{"foobar", "42"}, expectedOk: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			is.NotPanics(func() {
				out, ok := FromAnySlice[string](tt.input)
				is.Equal(tt.expectedOut, out)
				is.Equal(tt.expectedOk, ok)
			})
		})
	}
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Empty(Empty[string]())
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Empty(Empty[int64]())
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type test struct{}

		is.Empty(Empty[test]())
	})

	t.Run("chan", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Empty(Empty[chan string]())
	})

	t.Run("slice", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Nil(Empty[[]int]())
	})

	t.Run("map", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.Nil(Empty[map[string]int]())
	})
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.True(IsEmpty(""))
		is.False(IsEmpty("foo"))
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.True(IsEmpty[int64](0))
		is.False(IsEmpty[int64](42))
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type test struct {
			foobar string
		}

		is.True(IsEmpty(test{foobar: ""}))
		is.False(IsEmpty(test{foobar: "foo"}))
	})
}

func TestIsNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.False(IsNotEmpty(""))
		is.True(IsNotEmpty("foo"))
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.False(IsNotEmpty[int64](0))
		is.True(IsNotEmpty[int64](42))
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type test struct {
			foobar string
		}

		is.False(IsNotEmpty(test{foobar: ""}))
		is.True(IsNotEmpty(test{foobar: "foo"}))
	})
}

func TestCoalesce(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name       string
			values     []int
			expected   int
			expectedOk bool
		}{
			{name: "no values", values: nil, expected: 0, expectedOk: false},
			{name: "single non-zero value", values: []int{3}, expected: 3, expectedOk: true},
			{name: "first non-zero of many", values: []int{0, 1, 2, 3}, expected: 1, expectedOk: true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				result, ok := Coalesce(tt.values...)
				is.Equal(tt.expected, result)
				is.Equal(tt.expectedOk, ok)
			})
		}
	})

	t.Run("*string", func(t *testing.T) {
		t.Parallel()

		newStr := func(v string) *string { return &v }
		var nilStr *string
		str1 := newStr("str1")
		str2 := newStr("str2")

		tests := []struct {
			name       string
			values     []*string
			expected   *string
			expectedOk bool
		}{
			{name: "all nil", values: []*string{nil, nilStr}, expected: nil, expectedOk: false},
			{name: "nilStr then str1", values: []*string{nilStr, str1}, expected: str1, expectedOk: true},
			{name: "nilStr then str1 str2", values: []*string{nilStr, str1, str2}, expected: str1, expectedOk: true},
			{name: "str1 str2 nilStr", values: []*string{str1, str2, nilStr}, expected: str1, expectedOk: true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				result, ok := Coalesce(tt.values...)
				is.Equal(tt.expected, result)
				is.Equal(tt.expectedOk, ok)
			})
		}
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()

		type structType struct {
			field1 int
			field2 float64
		}
		var zeroStruct structType
		struct1 := structType{1, 1.0}
		struct2 := structType{2, 2.0}

		tests := []struct {
			name       string
			values     []structType
			expected   structType
			expectedOk bool
		}{
			{name: "zero struct only", values: []structType{zeroStruct}, expected: zeroStruct, expectedOk: false},
			{name: "zero struct then struct1", values: []structType{zeroStruct, struct1}, expected: struct1, expectedOk: true},
			{name: "zero struct then struct1 struct2", values: []structType{zeroStruct, struct1, struct2}, expected: struct1, expectedOk: true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				result, ok := Coalesce(tt.values...)
				is.Equal(tt.expected, result)
				is.Equal(tt.expectedOk, ok)
			})
		}
	})
}

func TestCoalesceOrEmpty(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			values   []int
			expected int
		}{
			{name: "no values", values: nil, expected: 0},
			{name: "single non-zero value", values: []int{3}, expected: 3},
			{name: "first non-zero of many", values: []int{0, 1, 2, 3}, expected: 1},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				is.Equal(tt.expected, CoalesceOrEmpty(tt.values...))
			})
		}
	})

	t.Run("*string", func(t *testing.T) {
		t.Parallel()

		newStr := func(v string) *string { return &v }
		var nilStr *string
		str1 := newStr("str1")
		str2 := newStr("str2")

		tests := []struct {
			name     string
			values   []*string
			expected *string
		}{
			{name: "all nil", values: []*string{nil, nilStr}, expected: nil},
			{name: "nilStr then str1", values: []*string{nilStr, str1}, expected: str1},
			{name: "nilStr then str1 str2", values: []*string{nilStr, str1, str2}, expected: str1},
			{name: "str1 str2 nilStr", values: []*string{str1, str2, nilStr}, expected: str1},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				is.Equal(tt.expected, CoalesceOrEmpty(tt.values...))
			})
		}
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()

		type structType struct {
			field1 int
			field2 float64
		}
		var zeroStruct structType
		struct1 := structType{1, 1.0}
		struct2 := structType{2, 2.0}

		tests := []struct {
			name     string
			values   []structType
			expected structType
		}{
			{name: "zero struct only", values: []structType{zeroStruct}, expected: zeroStruct},
			{name: "zero struct then struct1", values: []structType{zeroStruct, struct1}, expected: struct1},
			{name: "zero struct then struct1 struct2", values: []structType{zeroStruct, struct1, struct2}, expected: struct1},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				is.Equal(tt.expected, CoalesceOrEmpty(tt.values...))
			})
		}
	})
}

func TestCoalesceSlice(t *testing.T) {
	t.Parallel()

	var sliceNil []int
	slice0 := []int{}
	slice1 := []int{1}
	slice2 := []int{1, 2}

	tests := []struct {
		name       string
		values     [][]int
		expected   []int
		expectedOk bool
	}{
		{name: "no values", values: nil, expected: []int{}, expectedOk: false},
		{name: "single nil slice", values: [][]int{nil}, expected: []int{}, expectedOk: false},
		{name: "single nil-valued slice var", values: [][]int{sliceNil}, expected: []int{}, expectedOk: false},
		{name: "single empty slice", values: [][]int{slice0}, expected: []int{}, expectedOk: false},
		{name: "nil, nil-valued var, empty slice", values: [][]int{nil, sliceNil, slice0}, expected: []int{}, expectedOk: false},
		{name: "single two-element slice", values: [][]int{slice2}, expected: slice2, expectedOk: true},
		{name: "single one-element slice", values: [][]int{slice1}, expected: slice1, expectedOk: true},
		{name: "one-element then two-element", values: [][]int{slice1, slice2}, expected: slice1, expectedOk: true},
		{name: "two-element then one-element", values: [][]int{slice2, slice1}, expected: slice2, expectedOk: true},
		{name: "nil, empty, one, two", values: [][]int{sliceNil, slice0, slice1, slice2}, expected: slice1, expectedOk: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			result, ok := CoalesceSlice(tt.values...)
			is.NotNil(result)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedOk, ok)
		})
	}
}

func TestCoalesceSliceOrEmpty(t *testing.T) {
	t.Parallel()

	var sliceNil []int
	slice0 := []int{}
	slice1 := []int{1}
	slice2 := []int{1, 2}

	tests := []struct {
		name     string
		values   [][]int
		expected []int
	}{
		{name: "no values", values: nil, expected: []int{}},
		{name: "single nil slice", values: [][]int{nil}, expected: []int{}},
		{name: "single nil-valued slice var", values: [][]int{sliceNil}, expected: []int{}},
		{name: "single empty slice", values: [][]int{slice0}, expected: []int{}},
		{name: "nil, nil-valued var, empty slice", values: [][]int{nil, sliceNil, slice0}, expected: []int{}},
		{name: "single two-element slice", values: [][]int{slice2}, expected: slice2},
		{name: "single one-element slice", values: [][]int{slice1}, expected: slice1},
		{name: "one-element then two-element", values: [][]int{slice1, slice2}, expected: slice1},
		{name: "two-element then one-element", values: [][]int{slice2, slice1}, expected: slice2},
		{name: "nil, empty, one, two", values: [][]int{sliceNil, slice0, slice1, slice2}, expected: slice1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			result := CoalesceSliceOrEmpty(tt.values...)
			is.NotNil(result)
			is.Equal(tt.expected, result)
		})
	}
}

func TestCoalesceMap(t *testing.T) {
	t.Parallel()

	var mapNil map[int]int
	map0 := map[int]int{}
	map1 := map[int]int{1: 1}
	map2 := map[int]int{1: 1, 2: 2}

	tests := []struct {
		name       string
		values     []map[int]int
		expected   map[int]int
		expectedOk bool
	}{
		{name: "no values", values: nil, expected: map[int]int{}, expectedOk: false},
		{name: "single nil map", values: []map[int]int{nil}, expected: map[int]int{}, expectedOk: false},
		{name: "single nil-valued map var", values: []map[int]int{mapNil}, expected: map[int]int{}, expectedOk: false},
		{name: "single empty map", values: []map[int]int{map0}, expected: map[int]int{}, expectedOk: false},
		{name: "nil, nil-valued var, empty map", values: []map[int]int{nil, mapNil, map0}, expected: map[int]int{}, expectedOk: false},
		{name: "single two-entry map", values: []map[int]int{map2}, expected: map2, expectedOk: true},
		{name: "single one-entry map", values: []map[int]int{map1}, expected: map1, expectedOk: true},
		{name: "one-entry then two-entry", values: []map[int]int{map1, map2}, expected: map1, expectedOk: true},
		{name: "two-entry then one-entry", values: []map[int]int{map2, map1}, expected: map2, expectedOk: true},
		{name: "nil, empty, one, two", values: []map[int]int{mapNil, map0, map1, map2}, expected: map1, expectedOk: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			result, ok := CoalesceMap(tt.values...)
			is.NotNil(result)
			is.Equal(tt.expected, result)
			is.Equal(tt.expectedOk, ok)
		})
	}
}

func TestCoalesceMapOrEmpty(t *testing.T) {
	t.Parallel()

	var mapNil map[int]int
	map0 := map[int]int{}
	map1 := map[int]int{1: 1}
	map2 := map[int]int{1: 1, 2: 2}

	tests := []struct {
		name     string
		values   []map[int]int
		expected map[int]int
	}{
		{name: "no values", values: nil, expected: map[int]int{}},
		{name: "single nil map", values: []map[int]int{nil}, expected: map[int]int{}},
		{name: "single nil-valued map var", values: []map[int]int{mapNil}, expected: map[int]int{}},
		{name: "single empty map", values: []map[int]int{map0}, expected: map[int]int{}},
		{name: "nil, nil-valued var, empty map", values: []map[int]int{nil, mapNil, map0}, expected: map[int]int{}},
		{name: "single two-entry map", values: []map[int]int{map2}, expected: map2},
		{name: "single one-entry map", values: []map[int]int{map1}, expected: map1},
		{name: "one-entry then two-entry", values: []map[int]int{map1, map2}, expected: map1},
		{name: "two-entry then one-entry", values: []map[int]int{map2, map1}, expected: map2},
		{name: "nil, empty, one, two", values: []map[int]int{mapNil, map0, map1, map2}, expected: map1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			result := CoalesceMapOrEmpty(tt.values...)
			is.NotNil(result)
			is.Equal(tt.expected, result)
		})
	}
}
