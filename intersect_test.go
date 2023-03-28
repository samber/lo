package lo_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Contains([]int{0, 1, 2, 3, 4, 5}, 5)
	result2 := lo.Contains([]int{0, 1, 2, 3, 4, 5}, 6)

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
	result1 := lo.ContainsBy(a1, func(t a) bool { return t.A == 1 && t.B == "2" })
	result2 := lo.ContainsBy(a1, func(t a) bool { return t.A == 2 && t.B == "2" })

	a2 := []string{"aaa", "bbb", "ccc"}
	result3 := lo.ContainsBy(a2, func(t string) bool { return t == "ccc" })
	result4 := lo.ContainsBy(a2, func(t string) bool { return t == "ddd" })

	is.Equal(result1, false)
	is.Equal(result2, true)
	is.Equal(result3, true)
	is.Equal(result4, false)
}

func TestEvery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := lo.Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := lo.Every([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := lo.Every([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.True(result1)
	is.False(result2)
	is.False(result3)
	is.True(result4)
}

func TestEveryBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.True(result1)

	result2 := lo.EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.False(result2)

	result3 := lo.EveryBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.False(result3)

	result4 := lo.EveryBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.True(result4)
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := lo.Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := lo.Some([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := lo.Some([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.True(result1)
	is.True(result2)
	is.False(result3)
	is.False(result4)
}

func TestSomeBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.True(result1)

	result2 := lo.SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.True(result2)

	result3 := lo.SomeBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.False(result3)

	result4 := lo.SomeBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.False(result4)
}

func TestNone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.None([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := lo.None([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := lo.None([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := lo.None([]int{0, 1, 2, 3, 4, 5}, []int{})

	is.False(result1)
	is.False(result2)
	is.True(result3)
	is.True(result4)
}

func TestNoneBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})

	is.False(result1)

	result2 := lo.NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})

	is.False(result2)

	result3 := lo.NoneBy([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})

	is.True(result3)

	result4 := lo.NoneBy([]int{}, func(x int) bool {
		return x < 5
	})

	is.True(result4)
}

func TestIntersect(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Intersect([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := lo.Intersect([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := lo.Intersect([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := lo.Intersect([]int{0, 6}, []int{0, 1, 2, 3, 4, 5})
	result5 := lo.Intersect([]int{0, 6, 0}, []int{0, 1, 2, 3, 4, 5})

	is.Equal(result1, []int{0, 2})
	is.Equal(result2, []int{0})
	is.Equal(result3, []int{})
	is.Equal(result4, []int{0})
	is.Equal(result5, []int{0})
}

func TestDifference(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	left1, right1 := lo.Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	is.Equal(left1, []int{1, 3, 4, 5})
	is.Equal(right1, []int{6})

	left2, right2 := lo.Difference([]int{1, 2, 3, 4, 5}, []int{0, 6})
	is.Equal(left2, []int{1, 2, 3, 4, 5})
	is.Equal(right2, []int{0, 6})

	left3, right3 := lo.Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5})
	is.Equal(left3, []int{})
	is.Equal(right3, []int{})
}

func TestUnion(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	result2 := lo.Union([]int{0, 1, 2, 3, 4, 5}, []int{6, 7})
	result3 := lo.Union([]int{0, 1, 2, 3, 4, 5}, []int{})
	result4 := lo.Union([]int{0, 1, 2}, []int{0, 1, 2})
	result5 := lo.Union([]int{}, []int{})
	is.Equal(result1, []int{0, 1, 2, 3, 4, 5, 10})
	is.Equal(result2, []int{0, 1, 2, 3, 4, 5, 6, 7})
	is.Equal(result3, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result4, []int{0, 1, 2})
	is.Equal(result5, []int{})

	result11 := lo.Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10}, []int{0, 1, 11})
	result12 := lo.Union([]int{0, 1, 2, 3, 4, 5}, []int{6, 7}, []int{8, 9})
	result13 := lo.Union([]int{0, 1, 2, 3, 4, 5}, []int{}, []int{})
	result14 := lo.Union([]int{0, 1, 2}, []int{0, 1, 2}, []int{0, 1, 2})
	result15 := lo.Union([]int{}, []int{}, []int{})
	is.Equal(result11, []int{0, 1, 2, 3, 4, 5, 10, 11})
	is.Equal(result12, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	is.Equal(result13, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result14, []int{0, 1, 2})
	is.Equal(result15, []int{})
}

func TestWithout(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Without([]int{0, 2, 10}, 0, 1, 2, 3, 4, 5)
	result2 := lo.Without([]int{0, 7}, 0, 1, 2, 3, 4, 5)
	result3 := lo.Without([]int{}, 0, 1, 2, 3, 4, 5)
	result4 := lo.Without([]int{0, 1, 2}, 0, 1, 2)
	result5 := lo.Without([]int{})
	is.Equal(result1, []int{10})
	is.Equal(result2, []int{7})
	is.Equal(result3, []int{})
	is.Equal(result4, []int{})
	is.Equal(result5, []int{})
}

func TestWithoutEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.WithoutEmpty([]int{0, 1, 2})
	result2 := lo.WithoutEmpty([]int{1, 2})
	result3 := lo.WithoutEmpty([]int{})
	is.Equal(result1, []int{1, 2})
	is.Equal(result2, []int{1, 2})
	is.Equal(result3, []int{})
}
