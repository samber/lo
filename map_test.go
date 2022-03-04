package lo

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	is := assert.New(t)

	r1 := Keys[string, int](map[string]int{"foo": 1, "bar": 2})

	sort.Sort(sort.StringSlice(r1))
	is.Equal(r1, []string{"bar", "foo"})
}

func TestValues(t *testing.T) {
	is := assert.New(t)

	r1 := Values[string, int](map[string]int{"foo": 1, "bar": 2})

	sort.Sort(sort.IntSlice(r1))
	is.Equal(r1, []int{1, 2})
}

func TestEntries(t *testing.T) {
	is := assert.New(t)

	r1 := Entries[string, int](map[string]int{"foo": 1, "bar": 2})

	is.EqualValues(r1, []Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
	})
}

func TestFromEntries(t *testing.T) {
	is := assert.New(t)

	r1 := FromEntries[string, int]([]Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
	})

	is.Len(r1, 2)
	is.Equal(r1["foo"], 1)
	is.Equal(r1["bar"], 2)
}

func TestAssign(t *testing.T) {
	is := assert.New(t)

	result1 := Assign[string, int](map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})

	is.Len(result1, 3)
	is.Equal(result1, map[string]int{"a": 1, "b": 3, "c": 4})
}
