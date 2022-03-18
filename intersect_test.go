package lo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	is := assert.New(t)

	result1 := Contains[int]([]int{0, 1, 2, 3, 4, 5}, 5)
	result2 := Contains[int]([]int{0, 1, 2, 3, 4, 5}, 6)

	is.Equal(result1, true)
	is.Equal(result2, false)
}

func TestContainsBy(t *testing.T) {
	is := assert.New(t)

	type a struct {
		A int
		B string
	}

	a1 := []a{a{A: 1, B: "1"}, a{A: 2, B: "2"}, a{A: 3, B: "3"}}
	result1 := ContainsBy[a](a1, func(t a) bool { return t.A == 1 && t.B == "2" })
	result2 := ContainsBy[a](a1, func(t a) bool { return t.A == 2 && t.B == "2" })

	a2 := []string{"aaa", "bbb", "ccc"}
	result3 := ContainsBy[string](a2, func(t string) bool { return t == "ccc" })
	result4 := ContainsBy[string](a2, func(t string) bool { return t == "ddd" })

	is.Equal(result1, false)
	is.Equal(result2, true)
	is.Equal(result3, true)
	is.Equal(result4, false)
}

func TestSome(t *testing.T) {
	is := assert.New(t)

	result1 := Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})

	is.Equal(result1, true)
	is.Equal(result2, true)
	is.Equal(result3, false)
}

func TestEvery(t *testing.T) {
	is := assert.New(t)

	result1 := Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})

	is.Equal(result1, true)
	is.Equal(result2, false)
	is.Equal(result3, false)
}

func TestIntersect(t *testing.T) {
	is := assert.New(t)

	result1 := Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := Intersect[int]([]int{0, 6}, []int{0, 1, 2, 3, 4, 5})
	result5 := Intersect[int]([]int{0, 6, 0}, []int{0, 1, 2, 3, 4, 5})

	is.Equal(result1, []int{0, 2})
	is.Equal(result2, []int{0})
	is.Equal(result3, []int{})
	is.Equal(result4, []int{0})
	is.Equal(result5, []int{0})
}

func TestDifference(t *testing.T) {
	is := assert.New(t)

	left1, right1 := Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	is.Equal(left1, []int{1, 3, 4, 5})
	is.Equal(right1, []int{6})

	left2, right2 := Difference[int]([]int{1, 2, 3, 4, 5}, []int{0, 6})
	is.Equal(left2, []int{1, 2, 3, 4, 5})
	is.Equal(right2, []int{0, 6})

	left3, right3 := Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5})
	is.Equal(left3, []int{})
	is.Equal(right3, []int{})
}

func TestUnion(t *testing.T) {
	is := assert.New(t)
	result1 := Union[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	result2 := Union[int]([]int{0, 1, 2, 3, 4, 5}, []int{6, 7})
	result3 := Union[int]([]int{0, 1, 2, 3, 4, 5}, []int{})
	result4 := Union[int]([]int{0, 1, 2}, []int{0, 1, 2})
	result5 := Union[int]([]int{}, []int{})
	is.Equal(result1, []int{0, 1, 2, 3, 4, 5, 10})
	is.Equal(result2, []int{0, 1, 2, 3, 4, 5, 6, 7})
	is.Equal(result3, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result4, []int{0, 1, 2})
	is.Equal(result5, []int{})
}

func TestExcept(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		expl string
		args args
		want []int
	}{
		{
			expl: "nils",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: []int{},
		},
		{
			expl: "empties",
			args: args{
				list1: []int{},
				list2: []int{},
			},
			want: []int{},
		},
		{
			expl: "empty, nil",
			args: args{
				list1: []int{},
				list2: nil,
			},
			want: []int{},
		},
		{
			expl: "nil, empty",
			args: args{
				list1: nil,
				list2: []int{},
			},
			want: []int{},
		},
		{
			expl: "one and none",
			args: args{
				list1: []int{1},
				list2: []int{},
			},
			want: []int{1},
		},
		{
			expl: "none and one",
			args: args{
				list1: []int{},
				list2: []int{1},
			},
			want: []int{},
		},
		{
			expl: "none and many",
			args: args{
				list1: []int{},
				list2: []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			expl: "many and none",
			args: args{
				list1: []int{1, 2, 3},
				list2: []int{},
			},
			want: []int{1, 2, 3},
		},
		{
			expl: "exclude identical one",
			args: args{
				list1: []int{1},
				list2: []int{1},
			},
			want: []int{},
		},
		{
			expl: "exclude identical many",
			args: args{
				list1: []int{1, 2, 3},
				list2: []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			expl: "removal from middle",
			args: args{
				list1: []int{1, 2, 3, 4, 5, 6},
				list2: []int{2, 3},
			},
			want: []int{1, 4, 5, 6},
		},
		{
			expl: "removal from beginning",
			args: args{
				list1: []int{1, 2, 3, 4, 5, 6},
				list2: []int{1},
			},
			want: []int{2, 3, 4, 5, 6},
		},
		{
			expl: "removal from end",
			args: args{
				list1: []int{1, 2, 3, 4, 5, 6},
				list2: []int{6},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			expl: "nothing in common",
			args: args{
				list1: []int{1, 2, 3, 4, 5, 6},
				list2: []int{7, 8, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			expl: "maintain order",
			args: args{
				list1: []int{7, 1, 3, 8},
				list2: []int{1},
			},
			want: []int{7, 3, 8},
		},
		{
			expl: "maintain order in chaos",
			args: args{
				list1: []int{7, 1, 3, 8, 5, 4},
				list2: []int{1, 9, 7, 8},
			},
			want: []int{3, 5, 4},
		},
		{
			expl: "exclude heavy",
			args: args{
				list1: []int{3, 1},
				list2: []int{7, 1, 3, 8, 5, 4},
			},
			want: []int{},
		},
		{
			expl: "keep distinct exclude nothing",
			args: args{
				list1: []int{3, 1, 5, 1, 3},
				list2: []int{},
			},
			want: []int{3, 1, 5},
		},
		{
			expl: "keep distinct exclude one",
			args: args{
				list1: []int{3, 1, 5, 1, 3},
				list2: []int{1},
			},
			want: []int{3, 5},
		},
		{
			expl: "keep distinct exclude many",
			args: args{
				list1: []int{3, 1, 5, 1, 3},
				list2: []int{3, 5},
			},
			want: []int{1},
		},
	}

	prettyFormatSlice := func(vals []int) string {
		valsStr := make([]string, len(vals))

		i := 0
		for _, v := range vals {
			valsStr[i] = strconv.Itoa(v)
			i++
		}

		return strings.Join(valsStr, ",")
	}

	for _, tt := range tests {
		name := fmt.Sprintf(
			"when {%s} except {%s} expect {%s} (%s)",
			prettyFormatSlice(tt.args.list1),
			prettyFormatSlice(tt.args.list2),
			prettyFormatSlice(tt.want),
			tt.expl,
		)

		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Except(tt.args.list1, tt.args.list2), "Except(%v, %v)", tt.args.list1, tt.args.list2)
		})
	}
}
