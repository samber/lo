package lo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Contains([]int{0, 1, 2, 3, 4, 5}, 5)
	result2 := Contains([]int{0, 1, 2, 3, 4, 5}, 6)

	is.Equal(result1, true)
	is.Equal(result2, false)
}

func TestContainsBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type a struct {
		A int
		B string
	}

	a1 := []a{{A: 1, B: "1"}, {A: 2, B: "2"}, {A: 3, B: "3"}}
	result1 := ContainsBy(a1, func(t a) bool { return t.A == 1 && t.B == "2" })
	result2 := ContainsBy(a1, func(t a) bool { return t.A == 2 && t.B == "2" })

	a2 := []string{"aaa", "bbb", "ccc"}
	result3 := ContainsBy(a2, func(t string) bool { return t == "ccc" })
	result4 := ContainsBy(a2, func(t string) bool { return t == "ddd" })

	is.Equal(result1, false)
	is.Equal(result2, true)
	is.Equal(result3, true)
	is.Equal(result4, false)
}

func TestEvery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Every([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := Every([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.True(result1)
	is.False(result2)
	is.False(result3)
	is.True(result4)
}

func TestEveryBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.True(result1)

	result2 := EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.False(result2)

	result3 := EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.False(result3)

	result4 := EveryBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.True(result4)
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Some([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := Some([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.True(result1)
	is.True(result2)
	is.False(result3)
	is.False(result4)
}

func TestSomeBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.True(result1)

	result2 := SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.True(result2)

	result3 := SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.False(result3)

	result4 := SomeBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.False(result4)
}

func TestNone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := None([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := None([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := None([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := None([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.False(result1)
	is.False(result2)
	is.True(result3)
	is.True(result4)
}

func TestNoneBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.False(result1)

	result2 := NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.False(result2)

	result3 := NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.True(result3)

	result4 := NoneBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.True(result4)
}

func TestIntersect(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Intersect([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Intersect([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Intersect([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := Intersect([]int{0, 6}, []int{0, 1, 2, 3, 4, 5})
	result5 := Intersect([]int{0, 6, 0}, []int{0, 1, 2, 3, 4, 5})

	is.Equal(result1, []int{0, 2})
	is.Equal(result2, []int{0})
	is.Equal(result3, []int{})
	is.Equal(result4, []int{0})
	is.Equal(result5, []int{0})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Intersect(allStrings, allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestDifference(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	left1, right1 := Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	is.Equal(left1, []int{1, 3, 4, 5})
	is.Equal(right1, []int{6})

	left2, right2 := Difference([]int{1, 2, 3, 4, 5}, []int{0, 6})
	is.Equal(left2, []int{1, 2, 3, 4, 5})
	is.Equal(right2, []int{0, 6})

	left3, right3 := Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5})
	is.Equal(left3, []int{})
	is.Equal(right3, []int{})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	a, b := Difference(allStrings, allStrings)
	is.IsType(a, allStrings, "type preserved")
	is.IsType(b, allStrings, "type preserved")
}

func TestUnion(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	result2 := Union([]int{0, 1, 2, 3, 4, 5}, []int{6, 7})
	result3 := Union([]int{0, 1, 2, 3, 4, 5}, []int{})
	result4 := Union([]int{0, 1, 2}, []int{0, 1, 2})
	result5 := Union([]int{}, []int{})
	is.Equal(result1, []int{0, 1, 2, 3, 4, 5, 10})
	is.Equal(result2, []int{0, 1, 2, 3, 4, 5, 6, 7})
	is.Equal(result3, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result4, []int{0, 1, 2})
	is.Equal(result5, []int{})

	result11 := Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10}, []int{0, 1, 11})
	result12 := Union([]int{0, 1, 2, 3, 4, 5}, []int{6, 7}, []int{8, 9})
	result13 := Union([]int{0, 1, 2, 3, 4, 5}, []int{}, []int{})
	result14 := Union([]int{0, 1, 2}, []int{0, 1, 2}, []int{0, 1, 2})
	result15 := Union([]int{}, []int{}, []int{})
	is.Equal(result11, []int{0, 1, 2, 3, 4, 5, 10, 11})
	is.Equal(result12, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	is.Equal(result13, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result14, []int{0, 1, 2})
	is.Equal(result15, []int{})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Union(allStrings, allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestWithout(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Without([]int{0, 2, 10}, 0, 1, 2, 3, 4, 5)
	result2 := Without([]int{0, 7}, 0, 1, 2, 3, 4, 5)
	result3 := Without([]int{}, 0, 1, 2, 3, 4, 5)
	result4 := Without([]int{0, 1, 2}, 0, 1, 2)
	result5 := Without([]int{})
	is.Equal(result1, []int{10})
	is.Equal(result2, []int{7})
	is.Equal(result3, []int{})
	is.Equal(result4, []int{})
	is.Equal(result5, []int{})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Without(allStrings, "")
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestWithoutBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type User struct {
		Name string
		Age  int
	}

	result1 := WithoutBy([]User{{Name: "nick"}, {Name: "peter"}},
		func(item User) string {
			return item.Name
		}, "nick", "lily")
	result2 := WithoutBy([]User{}, func(item User) int { return item.Age }, 1, 2, 3)
	result3 := WithoutBy([]User{}, func(item User) string { return item.Name })
	is.Equal(result1, []User{{Name: "peter"}})
	is.Equal(result2, []User{})
	is.Equal(result3, []User{})
}

func TestWithoutEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := WithoutEmpty([]int{0, 1, 2})
	result2 := WithoutEmpty([]int{1, 2})
	result3 := WithoutEmpty([]int{})
	result4 := WithoutEmpty([]*int{ToPtr(0), ToPtr(1), nil, ToPtr(2)})
	is.Equal(result1, []int{1, 2})
	is.Equal(result2, []int{1, 2})
	is.Equal(result3, []int{})
	is.Equal(result4, []*int{ToPtr(0), ToPtr(1), ToPtr(2)})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := WithoutEmpty(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestWithoutNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := WithoutNth([]int{5, 6, 7}, 1, 0)
	is.Equal([]int{7}, result1)

	result2 := WithoutNth([]int{1, 2})
	is.Equal([]int{1, 2}, result2)

	result3 := WithoutNth([]int{})
	is.Equal([]int{}, result3)

	result4 := WithoutNth([]int{0, 1, 2, 3}, -1, 4)
	is.Equal([]int{0, 1, 2, 3}, result4)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := WithoutNth(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestXor(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}

	tests := []struct {
		expl string
		args args
		want []int
	}{
		// although nil and empty slices are functionally equivalent
		// one is indeed nil, and the other is not
		// https://go.dev/play/p/eV55dA7Y4NT
		{
			expl: "nil, nil",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: []int{},
		},
		{
			expl: "empty, empty",
			args: args{
				list1: []int{},
				list2: []int{},
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
			expl: "empty, nil",
			args: args{
				list1: []int{},
				list2: nil,
			},
			want: []int{},
		},
		{
			expl: "zero",
			args: args{
				list1: []int{0, 1, 2},
				list2: []int{1, 2, 3},
			},
			want: []int{0, 3},
		},
		{
			// ya, it's a pun
			expl: "one with nothing",
			args: args{
				list1: []int{1},
				list2: []int{},
			},
			want: []int{1},
		},
		{
			// ooo, I'm on a roll
			expl: "all for nothing",
			args: args{
				list1: []int{1, 2, 3},
				list2: []int{},
			},
			want: []int{1, 2, 3},
		},
		{
			// realistically, I can't stop at this point
			// brace yourself
			expl: "all for one",
			args: args{
				list1: []int{1, 2, 3},
				list2: []int{1},
			},
			want: []int{2, 3},
		},
		{
			// you knew it was coming...
			expl: "one for all",
			args: args{
				list1: []int{1},
				list2: []int{1, 2, 3},
			},
			want: []int{2, 3},
		},
		{
			// putting my chips down
			expl: "all in",
			args: args{
				list1: []int{1, 2, 3},
				list2: []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			expl: "creamy filling",
			args: args{
				list1: []int{1, 2, 3},
				list2: []int{2},
			},
			want: []int{1, 3},
		},
		{
			expl: "end of the line",
			args: args{
				list1: []int{1, 2, 3},
				list2: []int{3},
			},
			want: []int{1, 2},
		},
		{
			expl: "going negative",
			args: args{
				list1: []int{-1, 2, 3},
				list2: []int{3},
			},
			want: []int{-1, 2},
		},
		{
			expl: "xor negative",
			args: args{
				list1: []int{-1, 2, 3},
				list2: []int{-1},
			},
			want: []int{2, 3},
		},
		{
			expl: "nothing alike",
			args: args{
				list1: []int{1, 2, 3},
				list2: []int{-1, 4},
			},
			want: []int{1, 2, 3, -1, 4},
		},
		{
			expl: "distinctly similar",
			args: args{
				list1: []int{1, 2, 2, 3, 3, 5},
				list2: []int{1, 3},
			},
			want: []int{2, 5},
		},
		{
			// buzz lightyear
			expl: "duplicates... everywhere...",
			args: args{
				list1: []int{1, 2, 2, 3, 3, 5},
				list2: []int{1, 2, 2, 3, 3, 5},
			},
			want: []int{},
		},
		{
			expl: "order matters, first come first served",
			args: args{
				list1: []int{3, 2, 1},
				list2: []int{1, 2, 2, 3, 3, 5},
			},
			want: []int{5},
		},
		{
			expl: "flip it and reverse it",
			args: args{
				list1: []int{1, 2, 3, 4, 5},
				list2: []int{5, 4, 3, 2, 1},
			},
			want: []int{},
		},
		{
			expl: "donut",
			args: args{
				list1: []int{1, 2, 3, 4, 5},
				list2: []int{3, 4, 5, 6, 7},
			},
			want: []int{1, 2, 6, 7},
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf(
			"when {%s} xor {%s} expect {%s} (%s)",
			prettyFormatSlice(tt.args.list1),
			prettyFormatSlice(tt.args.list2),
			prettyFormatSlice(tt.want),
			tt.expl,
		)
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Xor(tt.args.list1, tt.args.list2), "Xor(%v, %v)", tt.args.list1, tt.args.list2)
		})
	}
}
