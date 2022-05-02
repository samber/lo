package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	is := assert.New(t)

	result1 := ToPtr[[]int]([]int{1, 2})

	is.Equal(*result1, []int{1, 2})
}

func TestToSlicePtr(t *testing.T) {
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := ToSlicePtr[string]([]string{str1, str2})

	is.Equal(result1, []*string{&str1, &str2})
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

func TestSafe(t *testing.T) {
	is := assert.New(t)

	type a struct {
		foo *string
	}
	type b struct {
		a *a
	}
	type c struct {
		b *b
	}

	v1 := &c{b: &b{a: &a{foo: ToPtr("foobar")}}}
	v2 := &c{b: &b{a: nil}}
	v3 := map[string]map[string]map[string]int{
		"foo": map[string]map[string]int{
			"bar": map[string]int{
				"baz": 42,
			},
		},
	}

	r1, ok1 := Safe(func() string { return *v1.b.a.foo })
	r2, ok2 := Safe(func() string { return *v2.b.a.foo })
	r3, ok3 := Safe(func() *string { return v2.b.a.foo })
	r4, ok4 := Safe(func() int { return v3["foo"]["bar"]["baz"] })
	r5, ok5 := Safe(func() int { return v3["hello"]["world"]["foobar"] })

	is.Equal("foobar", r1)
	is.True(ok1)

	is.Equal("", r2)
	is.False(ok2)

	is.Nil(r3)
	is.False(ok3)

	is.Equal(42, r4)
	is.True(ok4)

	is.Equal(0, r5)
	is.True(ok5)

	is.Panics(func() {
		_, _ = Safe(func() string { panic("hello world") })
	})
}
