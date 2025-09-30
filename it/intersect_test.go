//go:build go1.23

package it

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Contains(values(0, 1, 2, 3, 4, 5), 5)
	result2 := Contains(values(0, 1, 2, 3, 4, 5), 6)

	is.True(result1)
	is.False(result2)
}

func TestContainsBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type a struct {
		A int
		B string
	}

	a1 := values(a{A: 1, B: "1"}, a{A: 2, B: "2"}, a{A: 3, B: "3"})
	result1 := ContainsBy(a1, func(t a) bool { return t.A == 1 && t.B == "2" })
	result2 := ContainsBy(a1, func(t a) bool { return t.A == 2 && t.B == "2" })

	a2 := values("aaa", "bbb", "ccc")
	result3 := ContainsBy(a2, func(t string) bool { return t == "ccc" })
	result4 := ContainsBy(a2, func(t string) bool { return t == "ddd" })

	is.False(result1)
	is.True(result2)
	is.True(result3)
	is.False(result4)
}

func TestEvery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Every(values(0, 1, 2, 3, 4, 5), 0, 2)
	result2 := Every(values(0, 1, 2, 3, 4, 5), 0, 6)
	result3 := Every(values(0, 1, 2, 3, 4, 5), -1, 6)
	result4 := Every(values(0, 1, 2, 3, 4, 5))

	is.True(result1)
	is.False(result2)
	is.False(result3)
	is.True(result4)
}

func TestEveryBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := EveryBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 5
	})

	is.True(result1)

	result2 := EveryBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 3
	})

	is.False(result2)

	result3 := EveryBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 0
	})

	is.False(result3)

	result4 := EveryBy(values[int](), func(x int) bool {
		return x < 5
	})

	is.True(result4)
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Some(values(0, 1, 2, 3, 4, 5), 0, 2)
	result2 := Some(values(0, 1, 2, 3, 4, 5), 0, 6)
	result3 := Some(values(0, 1, 2, 3, 4, 5), -1, 6)
	result4 := Some(values(0, 1, 2, 3, 4, 5))

	is.True(result1)
	is.True(result2)
	is.False(result3)
	is.False(result4)
}

func TestSomeBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SomeBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 5
	})

	is.True(result1)

	result2 := SomeBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 3
	})

	is.True(result2)

	result3 := SomeBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 0
	})

	is.False(result3)

	result4 := SomeBy(values[int](), func(x int) bool {
		return x < 5
	})

	is.False(result4)
}

func TestNone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := None(values(0, 1, 2, 3, 4, 5), 0, 2)
	result2 := None(values(0, 1, 2, 3, 4, 5), 0, 6)
	result3 := None(values(0, 1, 2, 3, 4, 5), -1, 6)
	result4 := None(values(0, 1, 2, 3, 4, 5))

	is.False(result1)
	is.False(result2)
	is.True(result3)
	is.True(result4)
}

func TestNoneBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := NoneBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 5
	})

	is.False(result1)

	result2 := NoneBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 3
	})

	is.False(result2)

	result3 := NoneBy(values(1, 2, 3, 4), func(x int) bool {
		return x < 0
	})

	is.True(result3)

	result4 := NoneBy(values[int](), func(x int) bool {
		return x < 5
	})

	is.True(result4)
}

func TestIntersect(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Intersect([]iter.Seq[int]{}...)
	result2 := Intersect(values(0, 1, 2, 3, 4, 5))
	result3 := Intersect(values(0, 1, 2, 3, 4, 5), values(0, 6))
	result4 := Intersect(values(0, 1, 2, 3, 4, 5), values(-1, 6))
	result5 := Intersect(values(0, 6, 0), values(0, 1, 2, 3, 4, 5))
	result6 := Intersect(values(0, 1, 2, 3, 4, 5), values(0, 6, 0))
	result7 := Intersect(values(0, 1, 2), values(1, 2, 3), values(2, 3, 4))
	result8 := Intersect(values(0, 1, 2), values(1, 2, 3), values(2, 3, 4), values(3, 4, 5))
	result9 := Intersect(values(0, 1, 2), values(0, 1, 2), values(1, 2, 3), values(2, 3, 4), values(3, 4, 5))

	is.Empty(slices.Collect(result1))
	is.Equal([]int{0, 1, 2, 3, 4, 5}, slices.Collect(result2))
	is.Equal([]int{0}, slices.Collect(result3))
	is.Empty(slices.Collect(result4))
	is.Equal([]int{0}, slices.Collect(result5))
	is.Equal([]int{0}, slices.Collect(result6))
	is.Equal([]int{2}, slices.Collect(result7))
	is.Empty(slices.Collect(result8))
	is.Empty(slices.Collect(result9))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Intersect(allStrings, allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestUnion(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Union(values(0, 1, 2, 3, 4, 5), values(0, 2, 10))
	result2 := Union(values(0, 1, 2, 3, 4, 5), values(6, 7))
	result3 := Union(values(0, 1, 2, 3, 4, 5), values[int]())
	result4 := Union(values(0, 1, 2), values(0, 1, 2, 3, 3))
	result5 := Union(values(0, 1, 2), values(0, 1, 2))
	result6 := Union(values[int](), values[int]())
	is.Equal([]int{0, 1, 2, 3, 4, 5, 10}, slices.Collect(result1))
	is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7}, slices.Collect(result2))
	is.Equal([]int{0, 1, 2, 3, 4, 5}, slices.Collect(result3))
	is.Equal([]int{0, 1, 2, 3}, slices.Collect(result4))
	is.Equal([]int{0, 1, 2}, slices.Collect(result5))
	is.Empty(slices.Collect(result6))

	result11 := Union(values(0, 1, 2, 3, 4, 5), values(0, 2, 10), values(0, 1, 11))
	result12 := Union(values(0, 1, 2, 3, 4, 5), values(6, 7), values(8, 9))
	result13 := Union(values(0, 1, 2, 3, 4, 5), values[int](), values[int]())
	result14 := Union(values(0, 1, 2), values(0, 1, 2), values(0, 1, 2))
	result15 := Union(values[int](), values[int](), values[int]())
	is.Equal([]int{0, 1, 2, 3, 4, 5, 10, 11}, slices.Collect(result11))
	is.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, slices.Collect(result12))
	is.Equal([]int{0, 1, 2, 3, 4, 5}, slices.Collect(result13))
	is.Equal([]int{0, 1, 2}, slices.Collect(result14))
	is.Empty(slices.Collect(result15))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Union(allStrings, allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestWithout(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Without(values(0, 2, 10), 0, 1, 2, 3, 4, 5)
	result2 := Without(values(0, 7), 0, 1, 2, 3, 4, 5)
	result3 := Without(values[int](), 0, 1, 2, 3, 4, 5)
	result4 := Without(values(0, 1, 2), 0, 1, 2)
	result5 := Without(values[int]())
	is.Equal([]int{10}, slices.Collect(result1))
	is.Equal([]int{7}, slices.Collect(result2))
	is.Empty(slices.Collect(result3))
	is.Empty(slices.Collect(result4))
	is.Empty(slices.Collect(result5))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
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

	result1 := WithoutBy(values(User{Name: "nick"}, User{Name: "peter"}),
		func(item User) string {
			return item.Name
		}, "nick", "lily")
	result2 := WithoutBy(values[User](), func(item User) int { return item.Age }, 1, 2, 3)
	result3 := WithoutBy(values[User](), func(item User) string { return item.Name })
	is.Equal([]User{{Name: "peter"}}, slices.Collect(result1))
	is.Empty(slices.Collect(result2))
	is.Empty(slices.Collect(result3))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := WithoutBy(allStrings, func(string) string { return "" })
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestWithoutNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := WithoutNth(values(5, 6, 7), 1, 0)
	is.Equal([]int{7}, slices.Collect(result1))

	result2 := WithoutNth(values(1, 2))
	is.Equal([]int{1, 2}, slices.Collect(result2))

	result3 := WithoutNth(values[int]())
	is.Empty(slices.Collect(result3))

	result4 := WithoutNth(values(0, 1, 2, 3), -1, 4)
	is.Equal([]int{0, 1, 2, 3}, slices.Collect(result4))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := WithoutNth(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestElementsMatch(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.False(ElementsMatch(values[int](), values(1)))
	is.False(ElementsMatch(values(1), values(2)))
	is.False(ElementsMatch(values(1), values(1, 2)))
	is.False(ElementsMatch(values(1, 1, 2), values(2, 2, 1)))

	is.True(ElementsMatch(values(1), values(1)))
	is.True(ElementsMatch(values(1, 1), values(1, 1)))
	is.True(ElementsMatch(values(1, 2), values(2, 1)))
	is.True(ElementsMatch(values(1, 1, 2), values(1, 2, 1)))
}

func TestElementsMatchBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type someType struct {
		key string
	}

	is.True(ElementsMatchBy(
		values(someType{key: "a"}, someType{key: "b"}),
		values(someType{key: "b"}, someType{key: "a"}),
		func(item someType) string { return item.key },
	))
}
