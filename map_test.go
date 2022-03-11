package lo

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/expect"
)

func TestKeys(t *testing.T) {
	is := assert.New(t)

	r1 := Keys[string, int](map[string]int{"foo": 1, "bar": 2})
	sort.Strings(r1)

	is.Equal(r1, []string{"bar", "foo"})
}

func TestValues(t *testing.T) {
	is := assert.New(t)

	r1 := Values[string, int](map[string]int{"foo": 1, "bar": 2})
	sort.Ints(r1)

	is.Equal(r1, []int{1, 2})
}

func TestEntries(t *testing.T) {
	is := assert.New(t)

	r1 := Entries[string, int](map[string]int{"foo": 1, "bar": 2})

	sort.Slice(r1, func(i, j int) bool {
		return r1[i].Value < r1[j].Value
	})
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

func TestMapValues(t *testing.T) {
	is := assert.New(t)

	result1 := MapValues[int, int, string](map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x int, _ int) string {
		return "Hello"
	})
	result2 := MapValues[int, int64, string](map[int]int64{1: 1, 2: 2, 3: 3, 4: 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, map[int]string{1: "Hello", 2: "Hello", 3: "Hello", 4: "Hello"})
	is.Equal(result2, map[int]string{1: "1", 2: "2", 3: "3", 4: "4"})
}


type Character struct {
	dir  string
	code int
}
func TestKeyBy(t *testing.T){
	r1 := KeyBy[Character, string]([]Character{
		{dir:"left", code:97},
		{dir:"right", code:100},
	},func(char Character)string {
		return string(rune(char.code))
	})

	expect := Character{dir:"left",code:97}

	if r1["a"] != expect {
		t.Fatalf("Expect %+v, got %+v", expect, r1["a"] )
	}
}