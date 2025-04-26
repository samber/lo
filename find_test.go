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
	t.Parallel()
	is := assert.New(t)

	result1 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 4)
	is.Equal(result2, -1)
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1, ok1 := Find([]string{"a", "b", "c", "d"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})

	result2, ok2 := Find([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(ok1, true)
	is.Equal(result1, "b")
	is.Equal(ok2, false)
	is.Equal(result2, "")
}

func TestFindIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d", "b"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 1)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindLastIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"b", "d", "c", "b", "a"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindLastIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 4)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindOrElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1 := FindOrElse([]string{"a", "b", "c", "d"}, "x", func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})
	result2 := FindOrElse([]string{"foobar"}, "x", func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(result1, "b")
	is.Equal(result2, "x")
}

func TestFindKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 2)
	is.Equal("bar", result1)
	is.True(ok1)

	result2, ok2 := FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 42)
	is.Equal("", result2)
	is.False(ok2)

	type test struct {
		foobar string
	}

	result3, ok3 := FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"foo"})
	is.Equal("foo", result3)
	is.True(ok3)

	result4, ok4 := FindKey(map[string]test{"foo": {"foo"}, "bar": {"bar"}, "baz": {"baz"}}, test{"hello world"})
	is.Equal("", result4)
	is.False(ok4)
}

func TestFindKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return k == "foo"
	})
	is.Equal("foo", result1)
	is.True(ok1)

	result2, ok2 := FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Equal("", result2)
	is.False(ok2)
}

func TestFindUniques(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniques([]int{1, 2, 3})

	is.Equal(3, len(result1))
	is.Equal([]int{1, 2, 3}, result1)

	result2 := FindUniques([]int{1, 2, 2, 3, 1, 2})

	is.Equal(1, len(result2))
	is.Equal([]int{3}, result2)

	result3 := FindUniques([]int{1, 2, 2, 1})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	result4 := FindUniques([]int{})

	is.Equal(0, len(result4))
	is.Equal([]int{}, result4)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := FindUniques(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFindUniquesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindUniquesBy([]int{0, 1, 2}, func(i int) int {
		return i % 3
	})

	is.Equal(3, len(result1))
	is.Equal([]int{0, 1, 2}, result1)

	result2 := FindUniquesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 3
	})

	is.Equal(1, len(result2))
	is.Equal([]int{2}, result2)

	result3 := FindUniquesBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	result4 := FindUniquesBy([]int{}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result4))
	is.Equal([]int{}, result4)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := FindUniquesBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFindDuplicates(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicates([]int{1, 2, 2, 1, 2, 3})

	is.Equal(2, len(result1))
	is.Equal([]int{1, 2}, result1)

	result2 := FindDuplicates([]int{1, 2, 3})

	is.Equal(0, len(result2))
	is.Equal([]int{}, result2)

	result3 := FindDuplicates([]int{})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := FindDuplicates(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestFindDuplicatesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FindDuplicatesBy([]int{3, 4, 5, 6, 7}, func(i int) int {
		return i % 3
	})

	is.Equal(2, len(result1))
	is.Equal([]int{3, 4}, result1)

	result2 := FindDuplicatesBy([]int{0, 1, 2, 3, 4}, func(i int) int {
		return i % 5
	})

	is.Equal(0, len(result2))
	is.Equal([]int{}, result2)

	result3 := FindDuplicatesBy([]int{}, func(i int) int {
		return i % 3
	})

	is.Equal(0, len(result3))
	is.Equal([]int{}, result3)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := FindDuplicatesBy(allStrings, func(i string) string {
		return i
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestMin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Min([]int{1, 2, 3})
	result2 := Min([]int{3, 2, 1})
	result3 := Min([]time.Duration{time.Second, time.Minute, time.Hour})
	result4 := Min([]int{})

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, time.Second)
	is.Equal(result4, 0)
}

func TestMinIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MinIndex([]int{1, 2, 3})
	result2, index2 := MinIndex([]int{3, 2, 1})
	result3, index3 := MinIndex([]time.Duration{time.Second, time.Minute, time.Hour})
	result4, index4 := MinIndex([]int{})

	is.Equal(result1, 1)
	is.Equal(index1, 0)

	is.Equal(result2, 1)
	is.Equal(index2, 2)

	is.Equal(result3, time.Second)
	is.Equal(index3, 0)

	is.Equal(result4, 0)
	is.Equal(index4, -1)
}

func TestMinBy(t *testing.T) {
	t.Parallel()
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

func TestMinIndexBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MinIndexBy([]string{"s1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result2, index2 := MinIndexBy([]string{"string1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	result3, index3 := MinIndexBy([]string{}, func(item string, min string) bool {
		return len(item) < len(min)
	})

	is.Equal(result1, "s1")
	is.Equal(index1, 0)

	is.Equal(result2, "s3")
	is.Equal(index2, 2)

	is.Equal(result3, "")
	is.Equal(index3, -1)
}

func TestEarliest(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	a := time.Now()
	b := a.Add(time.Hour)
	result1 := Earliest(a, b)
	result2 := Earliest()

	is.Equal(result1, a)
	is.Equal(result2, time.Time{})
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
	result1 := EarliestBy([]foo{{t1}, {t2}, {t3}}, func(i foo) time.Time {
		return i.bar
	})
	result2 := EarliestBy([]foo{{t1}}, func(i foo) time.Time {
		return i.bar
	})
	result3 := EarliestBy([]foo{}, func(i foo) time.Time {
		return i.bar
	})

	is.Equal(result1, foo{t3})
	is.Equal(result2, foo{t1})
	is.Equal(result3, foo{})
}

func TestMax(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Max([]int{1, 2, 3})
	result2 := Max([]int{3, 2, 1})
	result3 := Max([]time.Duration{time.Second, time.Minute, time.Hour})
	result4 := Max([]int{})

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, time.Hour)
	is.Equal(result4, 0)
}

func TestMaxIndex(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MaxIndex([]int{1, 2, 3})
	result2, index2 := MaxIndex([]int{3, 2, 1})
	result3, index3 := MaxIndex([]time.Duration{time.Second, time.Minute, time.Hour})
	result4, index4 := MaxIndex([]int{})

	is.Equal(result1, 3)
	is.Equal(index1, 2)

	is.Equal(result2, 3)
	is.Equal(index2, 0)

	is.Equal(result3, time.Hour)
	is.Equal(index3, 2)

	is.Equal(result4, 0)
	is.Equal(index4, -1)
}

func TestMaxBy(t *testing.T) {
	t.Parallel()
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

func TestMaxIndexBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, index1 := MaxIndexBy([]string{"s1", "string2", "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result2, index2 := MaxIndexBy([]string{"string1", "string2", "s3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	result3, index3 := MaxIndexBy([]string{}, func(item string, max string) bool {
		return len(item) > len(max)
	})

	is.Equal(result1, "string2")
	is.Equal(index1, 1)

	is.Equal(result2, "string1")
	is.Equal(index2, 0)

	is.Equal(result3, "")
	is.Equal(index3, -1)
}

func TestLatest(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	a := time.Now()
	b := a.Add(time.Hour)
	result1 := Latest(a, b)
	result2 := Latest()

	is.Equal(result1, b)
	is.Equal(result2, time.Time{})
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
	result1 := LatestBy([]foo{{t1}, {t2}, {t3}}, func(i foo) time.Time {
		return i.bar
	})
	result2 := LatestBy([]foo{{t1}}, func(i foo) time.Time {
		return i.bar
	})
	result3 := LatestBy([]foo{}, func(i foo) time.Time {
		return i.bar
	})

	is.Equal(result1, foo{t2})
	is.Equal(result2, foo{t1})
	is.Equal(result3, foo{})
}

func TestFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := First([]int{1, 2, 3})
	result2, ok2 := First([]int{})

	is.Equal(result1, 1)
	is.Equal(ok1, true)
	is.Equal(result2, 0)
	is.Equal(ok2, false)
}

func TestFirstOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FirstOrEmpty([]int{1, 2, 3})
	result2 := FirstOrEmpty([]int{})
	result3 := FirstOrEmpty([]string{})

	is.Equal(result1, 1)
	is.Equal(result2, 0)
	is.Equal(result3, "")
}

func TestFirstOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FirstOr([]int{1, 2, 3}, 63)
	result2 := FirstOr([]int{}, 23)
	result3 := FirstOr([]string{}, "test")

	is.Equal(result1, 1)
	is.Equal(result2, 23)
	is.Equal(result3, "test")
}

func TestLast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, ok1 := Last([]int{1, 2, 3})
	result2, ok2 := Last([]int{})

	is.Equal(result1, 3)
	is.True(ok1)
	is.Equal(result2, 0)
	is.False(ok2)
}

func TestLastOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastOrEmpty([]int{1, 2, 3})
	result2 := LastOrEmpty([]int{})
	result3 := LastOrEmpty([]string{})

	is.Equal(result1, 3)
	is.Equal(result2, 0)
	is.Equal(result3, "")
}

func TestLastOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastOr([]int{1, 2, 3}, 63)
	result2 := LastOr([]int{}, 23)
	result3 := LastOr([]string{}, "test")

	is.Equal(result1, 3)
	is.Equal(result2, 23)
	is.Equal(result3, "test")
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := Nth([]int{0, 1, 2, 3}, 2)
	result2, err2 := Nth([]int{0, 1, 2, 3}, -2)
	result3, err3 := Nth([]int{0, 1, 2, 3}, 42)
	result4, err4 := Nth([]int{}, 0)
	result5, err5 := Nth([]int{42}, 0)
	result6, err6 := Nth([]int{42}, -1)

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
	is.Equal(result6, 42)
	is.Equal(err6, nil)
}

func TestNthOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	t.Run("Integers", func(t *testing.T) {
		const defaultValue = -1
		intSlice := []int{10, 20, 30, 40, 50}

		is.Equal(30, NthOr(intSlice, 2, defaultValue))
		is.Equal(50, NthOr(intSlice, -1, defaultValue))
		is.Equal(defaultValue, NthOr(intSlice, 5, defaultValue))
	})

	t.Run("Strings", func(t *testing.T) {
		const defaultValue = "none"
		strSlice := []string{"apple", "banana", "cherry", "date"}

		is.Equal("banana", NthOr(strSlice, 1, defaultValue))      // Index 1, expected "banana"
		is.Equal("cherry", NthOr(strSlice, -2, defaultValue))     // Negative index -2, expected "cherry"
		is.Equal(defaultValue, NthOr(strSlice, 10, defaultValue)) // Out of bounds, fallback "none"
	})

	t.Run("Structs", func(t *testing.T) {
		type User struct {
			ID   int
			Name string
		}
		userSlice := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}

		expectedUser := User{ID: 1, Name: "Alice"}
		is.Equal(expectedUser, NthOr(userSlice, 0, User{ID: 0, Name: "Unknown"}))

		expectedUser = User{ID: 3, Name: "Charlie"}
		is.Equal(expectedUser, NthOr(userSlice, -1, User{ID: 0, Name: "Unknown"}))

		expectedUser = User{ID: 0, Name: "Unknown"}
		is.Equal(expectedUser, NthOr(userSlice, 10, User{ID: 0, Name: "Unknown"}))
	})
}

func TestNthOrEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	t.Run("Integers", func(t *testing.T) {
		const defaultValue = 0
		intSlice := []int{10, 20, 30, 40, 50}

		is.Equal(30, NthOrEmpty(intSlice, 2))
		is.Equal(50, NthOrEmpty(intSlice, -1))
		is.Equal(defaultValue, NthOrEmpty(intSlice, 10))
	})

	t.Run("Strings", func(t *testing.T) {
		const defaultValue = ""
		strSlice := []string{"apple", "banana", "cherry", "date"}

		is.Equal("banana", NthOrEmpty(strSlice, 1))
		is.Equal("cherry", NthOrEmpty(strSlice, -2))
		is.Equal(defaultValue, NthOrEmpty(strSlice, 10))
	})

	t.Run("Structs", func(t *testing.T) {
		type User struct {
			ID   int
			Name string
		}
		userSlice := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}

		expectedUser := User{ID: 1, Name: "Alice"}
		is.Equal(expectedUser, NthOrEmpty(userSlice, 0))

		expectedUser = User{ID: 3, Name: "Charlie"}
		is.Equal(expectedUser, NthOrEmpty(userSlice, -1))

		expectedUser = User{ID: 0, Name: ""}
		is.Equal(expectedUser, NthOrEmpty(userSlice, 10))
	})
}

func TestSample(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Sample([]string{"a", "b", "c"})
	result2 := Sample([]string{})

	is.True(Contains([]string{"a", "b", "c"}, result1))
	is.Equal(result2, "")
}

func TestSampleBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := rand.New(rand.NewSource(42))

	result1 := SampleBy([]string{"a", "b", "c"}, r.Intn)
	result2 := SampleBy([]string{}, rand.Intn)

	is.True(Contains([]string{"a", "b", "c"}, result1))
	is.Equal(result2, "")
}

func TestSamples(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Samples([]string{"a", "b", "c"}, 3)
	result2 := Samples([]string{}, 3)

	sort.Strings(result1)

	is.Equal(result1, []string{"a", "b", "c"})
	is.Equal(result2, []string{})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Samples(allStrings, 2)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestSamplesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r := rand.New(rand.NewSource(42))

	result1 := SamplesBy([]string{"a", "b", "c"}, 3, r.Intn)
	result2 := SamplesBy([]string{}, 3, r.Intn)

	sort.Strings(result1)

	is.Equal(result1, []string{"a", "b", "c"})
	is.Equal(result2, []string{})

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := SamplesBy(allStrings, 2, r.Intn)
	is.IsType(nonempty, allStrings, "type preserved")
}
