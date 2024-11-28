package la

import (
	"maps"
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Collect(Keys(maps.All(map[string]int{"foo": 1, "bar": 2})))
	sort.Strings(r1)
	is.Equal(r1, []string{"bar", "foo"})

	r2 := Collect(Keys(maps.All(map[string]int{})))
	is.Empty(r2)

	r3 := Collect(Keys(maps.All(map[string]int{"foo": 1, "bar": 2}), maps.All(map[string]int{"baz": 3})))
	sort.Strings(r3)
	is.Equal(r3, []string{"bar", "baz", "foo"})

	r4 := Collect(Keys[string, int]())
	is.Equal(r4, []string{})

	r5 := Collect(Keys(maps.All(map[string]int{"foo": 1, "bar": 2}), maps.All(map[string]int{"bar": 3})))
	sort.Strings(r5)
	is.Equal(r5, []string{"bar", "bar", "foo"})
}

func TestUniqKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Collect(UniqKeys(maps.All(map[string]int{"foo": 1, "bar": 2})))
	sort.Strings(r1)
	is.Equal(r1, []string{"bar", "foo"})

	r2 := Collect(UniqKeys(maps.All(map[string]int{})))
	is.Empty(r2)

	r3 := Collect(UniqKeys(maps.All(map[string]int{"foo": 1, "bar": 2}), maps.All(map[string]int{"baz": 3})))
	sort.Strings(r3)
	is.Equal(r3, []string{"bar", "baz", "foo"})

	r4 := Collect(UniqKeys[string, int]())
	is.Equal(r4, []string{})

	r5 := Collect(UniqKeys(maps.All(map[string]int{"foo": 1, "bar": 2}), maps.All(map[string]int{"foo": 1, "bar": 3})))
	sort.Strings(r5)
	is.Equal(r5, []string{"bar", "foo"})

	// check order
	r6 := Collect(UniqKeys(maps.All(map[string]int{"foo": 1}), maps.All(map[string]int{"bar": 3})))
	is.Equal(r6, []string{"foo", "bar"})
}

func TestValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Collect(Values(maps.All(map[string]int{"foo": 1, "bar": 2})))
	sort.Ints(r1)
	is.Equal(r1, []int{1, 2})

	r2 := Collect(Values(maps.All(map[string]int{})))
	is.Empty(r2)

	r3 := Collect(Values(maps.All(map[string]int{"foo": 1, "bar": 2}), maps.All(map[string]int{"baz": 3})))
	sort.Ints(r3)
	is.Equal(r3, []int{1, 2, 3})

	r4 := Collect(Values[string, int]())
	is.Equal(r4, []int{})

	r5 := Collect(Values(maps.All(map[string]int{"foo": 1, "bar": 2}), maps.All(map[string]int{"foo": 1, "bar": 3})))
	sort.Ints(r5)
	is.Equal(r5, []int{1, 1, 2, 3})
}

func TestUniqValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Collect(UniqValues(maps.All(map[string]int{"foo": 1, "bar": 2})))
	sort.Ints(r1)
	is.Equal(r1, []int{1, 2})

	r2 := Collect(UniqValues(maps.All(map[string]int{})))
	is.Empty(r2)

	r3 := Collect(UniqValues(maps.All(map[string]int{"foo": 1, "bar": 2}), maps.All(map[string]int{"baz": 3})))
	sort.Ints(r3)
	is.Equal(r3, []int{1, 2, 3})

	r4 := Collect(UniqValues[string, int]())
	is.Equal(r4, []int{})

	r5 := Collect(UniqValues(maps.All(map[string]int{"foo": 1, "bar": 2}), maps.All(map[string]int{"foo": 1, "bar": 3})))
	sort.Ints(r5)
	is.Equal(r5, []int{1, 2, 3})

	r6 := Collect(UniqValues(maps.All(map[string]int{"foo": 1, "bar": 1}), maps.All(map[string]int{"foo": 1, "bar": 3})))
	sort.Ints(r6)
	is.Equal(r6, []int{1, 3})

	// check order
	r7 := Collect(UniqValues(maps.All(map[string]int{"foo": 1}), maps.All(map[string]int{"bar": 3})))
	is.Equal(r7, []int{1, 3})
}

func TestInvert(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := CollectMap(Invert(maps.All(map[string]int{"a": 1, "b": 2})))
	r2 := CollectMap(Invert(maps.All(map[string]int{"a": 1, "b": 2, "c": 1})))

	is.Len(r1, 2)
	is.EqualValues(map[int]string{1: "a", 2: "b"}, r1)
	is.Len(r2, 2)
}

func TestJoin2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := CollectMap(Join2(maps.All(map[string]int{"a": 1, "b": 2}), maps.All(map[string]int{"b": 3, "c": 4})))

	is.Len(result1, 3)
	is.Equal(result1, map[string]int{"a": 1, "b": 3, "c": 4})

	type myIter func(yield func(string, int) bool)
	before := myIter(maps.All(map[string]int{"": 0, "foobar": 6, "baz": 3}))
	after := Join2(before, before)
	is.IsType(after, before, "type preserved")
}

func TestFollow(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := CollectMap(Follow(
		Enumerate(slices.Values([]string{"zero", "one", "two"})),
		Enumerate(slices.Values([]string{"three", "four", "five"})),
	))

	is.Len(result1, 6)
	is.Equal(result1, map[int]string{0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five"})

	type myIter func(yield func(int, string) bool)
	before := myIter(maps.All(map[int]string{0: "", 6: "foobar", 3: "baz"}))
	after := Follow(before, before)
	is.IsType(after, before, "type preserved")
}
