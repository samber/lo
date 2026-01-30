//go:build go1.23

package it

import (
	"iter"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/samber/lo/internal/xrand"
)

func TestIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := IndexOf(values(0, 1, 2, 1, 2, 3), 2)
	result2 := IndexOf(values(0, 1, 2, 1, 2, 3), 6)

	is.Equal(2, result1)
	is.Equal(-1, result2)
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastIndexOf(values(0, 1, 2, 1, 2, 3), 2)
	result2 := LastIndexOf(values(0, 1, 2, 1, 2, 3), 6)

	is.Equal(4, result1)
	is.Equal(-1, result2)
}

func TestHasPrefix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(HasPrefix(values(1, 2, 3, 4), 1, 2, 3, 4))
	is.True(HasPrefix(values(1, 2, 3, 4), 1, 2))
	is.False(HasPrefix(values(1, 2, 3, 4), 42))
	is.False(HasPrefix(values(1, 2), 1, 2, 3, 4))
	is.True(HasPrefix(values(1, 2, 3, 4)))
}

func TestHasSuffix(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(HasSuffix(values(1, 2, 3, 4), 1, 2, 3, 4))
	is.True(HasSuffix(values(1, 2, 3, 4), 3, 4))
	is.True(HasSuffix(values(1, 2, 3, 4, 5), 3, 4, 5))
	is.False(HasSuffix(values(1, 2, 3, 4), 42))
	is.False(HasSuffix(values(1, 2), 1, 2, 3, 4))
	is.True(HasSuffix(values(1, 2, 3, 4)))
	is.False(HasSuffix(values(0), 0, 0))
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1, ok1 := Find(values("a", "b", "c", "d"), func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})

	result2, ok2 := Find(values("foobar"), func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.True(ok1)
	is.Equal("b", result1)
	is.False(ok2)
	is.Empty(result2)
}

func TestFindIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindIndexOf(values("a", "b", "c", "d", "b"), func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d", "b"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindIndexOf(values("foobar"), func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal("b", item1)
	is.True(ok1)
	is.Equal(1, index1)
	is.Empty(item2)
	is.False(ok2)
	is.Equal(-1, index2)
}

func TestFindLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	item1, index1, ok1 := FindLastIndexOf(values("a", "b", "c", "d", "b"), func(item string) bool {
		return item == "b"
	})
	item2, index2, ok2 := FindLastIndexOf(values("foobar"), func(item string) bool {
		return item == "b"
	})

	is.Equal("b", item1)
	is.True(ok1)
	is.Equal(4, index1)
	is.Empty(item2)
	is.False(ok2)
	is.Equal(-1, index2)
}

func TestFindOrElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1 := FindOrElse(values("a", "b", "c", "d"), "x", func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})
	result2 := FindOrElse(values("foobar"), "x", func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal("b", result1)
	is.Equal("x", result2)
}

func TestFindUniques(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniques(values(1, 2, 3))
	is.Equal([]int{1, 2, 3}, slices.Collect(result1))

	result2 := FindUniques(values(1, 2, 2, 3, 1, 2))
	is.Equal([]int{3}, slices.Collect(result2))

	result3 := FindUniques(values(1, 2, 2, 1))
	is.Empty(slices.Collect(result3))

	result4 := FindUniques(values[int]())
	is.Empty(slices.Collect(result4))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := FindUniques(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFindUniquesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniquesBy(values(0, 1, 2), func(i int) int {
		return i % 3
	})
	is.Equal([]int{0, 1, 2}, slices.Collect(result1))

	result2 := FindUniquesBy(values(0, 1, 2, 3, 4), func(i int) int {
		return i % 3
	})
	is.Equal([]int{2}, slices.Collect(result2))

	result3 := FindUniquesBy(values(0, 1, 2, 3, 4, 5), func(i int) int {
		return i % 3
	})
	is.Empty(slices.Collect(result3))

	result4 := FindUniquesBy(values[int](), func(i int) int {
		return i % 3
	})
	is.Empty(slices.Collect(result4))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := FindUniquesBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFindDuplicates(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicates(values(1, 2, 2, 1, 2, 3))
	is.Equal([]int{2, 1}, slices.Collect(result1))

	result2 := FindDuplicates(values(1, 2, 3))
	is.Empty(slices.Collect(result2))

	result3 := FindDuplicates(values[int]())
	is.Empty(slices.Collect(result3))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := FindDuplicates(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFindDuplicatesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicatesBy(values(3, 4, 5, 6, 7), func(i int) int {
		return i % 3
	})
	is.Equal([]int{3, 4}, slices.Collect(result1))

	result2 := FindDuplicatesBy(values(0, 1, 2, 3, 4), func(i int) int {
		return i % 5
	})
	is.Empty(slices.Collect(result2))

	result3 := FindDuplicatesBy(values[int](), func(i int) int {
		return i % 3
	})
	is.Empty(slices.Collect(result3))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := FindDuplicatesBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Min(values(1, 2, 3))
	result2 := Min(values(3, 2, 1))
	result3 := Min(values(time.Second, time.Minute, time.Hour))
	result4 := Min(values[int]())

	is.Equal(1, result1)
	is.Equal(1, result2)
	is.Equal(time.Second, result3)
	is.Zero(result4)
}

func TestMinIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MinIndex(values(1, 2, 3))
	result2, index2 := MinIndex(values(3, 2, 1))
	result3, index3 := MinIndex(values(time.Second, time.Minute, time.Hour))
	result4, index4 := MinIndex(values[int]())

	is.Equal(1, result1)
	is.Zero(index1)

	is.Equal(1, result2)
	is.Equal(2, index2)

	is.Equal(time.Second, result3)
	is.Zero(index3)

	is.Zero(result4)
	is.Equal(-1, index4)
}

func TestMinBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MinBy(values("s1", "string2", "s3"), func(item, mIn string) bool {
		return len(item) < len(mIn)
	})
	result2 := MinBy(values("string1", "string2", "s3"), func(item, mIn string) bool {
		return len(item) < len(mIn)
	})
	result3 := MinBy(values[string](), func(item, mIn string) bool {
		return len(item) < len(mIn)
	})

	is.Equal("s1", result1)
	is.Equal("s3", result2)
	is.Empty(result3)
}

func TestMinIndexBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MinIndexBy(values("s1", "string2", "s3"), func(item, mIn string) bool {
		return len(item) < len(mIn)
	})
	result2, index2 := MinIndexBy(values("string1", "string2", "s3"), func(item, mIn string) bool {
		return len(item) < len(mIn)
	})
	result3, index3 := MinIndexBy(values[string](), func(item, mIn string) bool {
		return len(item) < len(mIn)
	})

	is.Equal("s1", result1)
	is.Zero(index1)

	is.Equal("s3", result2)
	is.Equal(2, index2)

	is.Empty(result3)
	is.Equal(-1, index3)
}

func TestEarliest(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	a := time.Now()
	b := a.Add(time.Hour)
	result1 := Earliest(values(a, b))
	result2 := Earliest(values[time.Time]())

	is.Equal(a, result1)
	is.Zero(result2)
}

func TestEarliestBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type foo struct {
		bar time.Time
	}

	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	t3 := t1.Add(-time.Hour)
	result1 := EarliestBy(values(foo{t1}, foo{t2}, foo{t3}), func(i foo) time.Time {
		return i.bar
	})
	result2 := EarliestBy(values(foo{t1}), func(i foo) time.Time {
		return i.bar
	})
	result3 := EarliestBy(values[foo](), func(i foo) time.Time {
		return i.bar
	})

	is.Equal(foo{t3}, result1)
	is.Equal(foo{t1}, result2)
	is.Zero(result3)
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Max(values(1, 2, 3))
	result2 := Max(values(3, 2, 1))
	result3 := Max(values(time.Second, time.Minute, time.Hour))
	result4 := Max(values[int]())

	is.Equal(3, result1)
	is.Equal(3, result2)
	is.Equal(time.Hour, result3)
	is.Zero(result4)
}

func TestMaxIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MaxIndex(values(1, 2, 3))
	result2, index2 := MaxIndex(values(3, 2, 1))
	result3, index3 := MaxIndex(values(time.Second, time.Minute, time.Hour))
	result4, index4 := MaxIndex(values[int]())

	is.Equal(3, result1)
	is.Equal(2, index1)

	is.Equal(3, result2)
	is.Zero(index2)

	is.Equal(time.Hour, result3)
	is.Equal(2, index3)

	is.Zero(result4)
	is.Equal(-1, index4)
}

func TestMaxBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MaxBy(values("s1", "string2", "s3"), func(item, mAx string) bool {
		return len(item) > len(mAx)
	})
	result2 := MaxBy(values("string1", "string2", "s3"), func(item, mAx string) bool {
		return len(item) > len(mAx)
	})
	result3 := MaxBy(values[string](), func(item, mAx string) bool {
		return len(item) > len(mAx)
	})

	is.Equal("string2", result1)
	is.Equal("string1", result2)
	is.Empty(result3)
}

func TestMaxIndexBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MaxIndexBy(values("s1", "string2", "s3"), func(item, mAx string) bool {
		return len(item) > len(mAx)
	})
	result2, index2 := MaxIndexBy(values("string1", "string2", "s3"), func(item, mAx string) bool {
		return len(item) > len(mAx)
	})
	result3, index3 := MaxIndexBy(values[string](), func(item, mAx string) bool {
		return len(item) > len(mAx)
	})

	is.Equal("string2", result1)
	is.Equal(1, index1)

	is.Equal("string1", result2)
	is.Zero(index2)

	is.Empty(result3)
	is.Equal(-1, index3)
}

func TestLatest(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	a := time.Now()
	b := a.Add(time.Hour)
	result1 := Latest(values(a, b))
	result2 := Latest(values[time.Time]())

	is.Equal(b, result1)
	is.Zero(result2)
}

func TestLatestBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type foo struct {
		bar time.Time
	}

	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	t3 := t1.Add(-time.Hour)
	result1 := LatestBy(values(foo{t1}, foo{t2}, foo{t3}), func(i foo) time.Time {
		return i.bar
	})
	result2 := LatestBy(values(foo{t1}), func(i foo) time.Time {
		return i.bar
	})
	result3 := LatestBy(values[foo](), func(i foo) time.Time {
		return i.bar
	})

	is.Equal(foo{t2}, result1)
	is.Equal(foo{t1}, result2)
	is.Zero(result3)
}

func TestFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := First(values(1, 2, 3))
	result2, ok2 := First(values[int]())

	is.Equal(1, result1)
	is.True(ok1)
	is.Zero(result2)
	is.False(ok2)
}

func TestFirstOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FirstOrEmpty(values(1, 2, 3))
	result2 := FirstOrEmpty(values[int]())
	result3 := FirstOrEmpty(values[string]())

	is.Equal(1, result1)
	is.Zero(result2)
	is.Empty(result3)
}

func TestFirstOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FirstOr(values(1, 2, 3), 63)
	result2 := FirstOr(values[int](), 23)
	result3 := FirstOr(values[string](), "test")

	is.Equal(1, result1)
	is.Equal(23, result2)
	is.Equal("test", result3)
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := Last(values(1, 2, 3))
	result2, ok2 := Last(values[int]())

	is.Equal(3, result1)
	is.True(ok1)
	is.Zero(result2)
	is.False(ok2)
}

func TestLastOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastOrEmpty(values(1, 2, 3))
	result2 := LastOrEmpty(values[int]())
	result3 := LastOrEmpty(values[string]())

	is.Equal(3, result1)
	is.Zero(result2)
	is.Empty(result3)
}

func TestLastOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastOr(values(1, 2, 3), 63)
	result2 := LastOr(values[int](), 23)
	result3 := LastOr(values[string](), "test")

	is.Equal(3, result1)
	is.Equal(23, result2)
	is.Equal("test", result3)
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := Nth(values(0, 1, 2, 3), 2)
	result2, err2 := Nth(values(0, 1, 2, 3), -2)
	result3, err3 := Nth(values(0, 1, 2, 3), 42)
	result4, err4 := Nth(values[int](), 0)
	result5, err5 := Nth(values(42), 0)
	result6, err6 := Nth(values(42), -1)

	is.Equal(2, result1)
	is.NoError(err1)
	is.Zero(result2)
	is.EqualError(err2, "nth: -2 out of bounds")
	is.Zero(result3)
	is.EqualError(err3, "nth: 42 out of bounds")
	is.Zero(result4)
	is.EqualError(err4, "nth: 0 out of bounds")
	is.Equal(42, result5)
	is.NoError(err5)
	is.Zero(result6)
	is.EqualError(err6, "nth: -1 out of bounds")
}

func TestNthOr(t *testing.T) {
	t.Parallel()

	t.Run("Integers", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		const defaultValue = -1
		ints := values(10, 20, 30, 40, 50)

		is.Equal(30, NthOr(ints, 2, defaultValue))
		is.Equal(defaultValue, NthOr(ints, -1, defaultValue))
		is.Equal(defaultValue, NthOr(ints, 5, defaultValue))
	})

	t.Run("Strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		const defaultValue = "none"
		strs := values("apple", "banana", "cherry", "date")

		is.Equal("banana", NthOr(strs, 1, defaultValue))      // Index 1, expected "banana"
		is.Equal(defaultValue, NthOr(strs, -2, defaultValue)) // Negative index -2, expected "cherry"
		is.Equal(defaultValue, NthOr(strs, 10, defaultValue)) // Out of bounds, fallback "none"
	})

	t.Run("Structs", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type User struct {
			ID   int
			Name string
		}
		users := values(
			User{ID: 1, Name: "Alice"},
			User{ID: 2, Name: "Bob"},
			User{ID: 3, Name: "Charlie"},
		)
		defaultValue := User{ID: 0, Name: "Unknown"}

		is.Equal(User{ID: 1, Name: "Alice"}, NthOr(users, 0, defaultValue))
		is.Equal(defaultValue, NthOr(users, -1, defaultValue))
		is.Equal(defaultValue, NthOr(users, 10, defaultValue))
	})
}

func TestNthOrEmpty(t *testing.T) {
	t.Parallel()

	t.Run("Integers", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		ints := values(10, 20, 30, 40, 50)

		is.Equal(30, NthOrEmpty(ints, 2))
		is.Zero(NthOrEmpty(ints, -1))
		is.Zero(NthOrEmpty(ints, 10))
	})

	t.Run("Strings", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		strs := values("apple", "banana", "cherry", "date")

		is.Equal("banana", NthOrEmpty(strs, 1))
		is.Empty(NthOrEmpty(strs, -2))
		is.Empty(NthOrEmpty(strs, 10))
	})

	t.Run("Structs", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type User struct {
			ID   int
			Name string
		}
		users := values(
			User{ID: 1, Name: "Alice"},
			User{ID: 2, Name: "Bob"},
			User{ID: 3, Name: "Charlie"},
		)

		is.Equal(User{ID: 1, Name: "Alice"}, NthOrEmpty(users, 0))
		is.Zero(NthOrEmpty(users, -1))
		is.Zero(NthOrEmpty(users, 10))
	})
}

func TestSample(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Sample(values("a", "b", "c"))
	result2 := Sample(values[string]())

	is.True(Contains(values("a", "b", "c"), result1))
	is.Empty(result2)
}

func TestSampleBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SampleBy(values("a", "b", "c"), xrand.IntN)
	result2 := SampleBy(values[string](), xrand.IntN)

	is.True(Contains(values("a", "b", "c"), result1))
	is.Empty(result2)
}

func TestSamples(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Samples(values("a", "b", "c"), 3)
	result2 := Samples(values[string](), 3)

	is.ElementsMatch(slices.Collect(result1), []string{"a", "b", "c"})
	is.Empty(slices.Collect(result2))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := Samples(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestSamplesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SamplesBy(values("a", "b", "c"), 3, xrand.IntN)
	result2 := SamplesBy(values[string](), 3, xrand.IntN)
	result3 := SamplesBy(values("a", "b", "c"), 3, func(n int) int { return n - 1 })
	result4 := SamplesBy(values("a", "b", "c"), 3, func(int) int { return 0 })
	result5 := SamplesBy(values("a", "b", "c"), 0, func(int) int { return 1 })
	result6 := SamplesBy(values("a", "b", "c"), -1, nil)

	// index out of range [1] with length 1
	is.Panics(func() {
		SamplesBy(values("a", "b", "c"), 3, func(int) int { return 1 })
	})

	is.ElementsMatch(slices.Collect(result1), []string{"a", "b", "c"})
	is.Empty(slices.Collect(result2))
	is.Equal([]string{"c", "b", "a"}, slices.Collect(result3))
	is.Equal([]string{"a", "c", "b"}, slices.Collect(result4))
	is.Empty(slices.Collect(result5))
	is.Empty(slices.Collect(result6))

	type myStrings iter.Seq[string]
	allStrings := myStrings(values("", "foo", "bar"))
	nonempty := SamplesBy(allStrings, 2, xrand.IntN)
	is.IsType(nonempty, allStrings, "type preserved")
}
