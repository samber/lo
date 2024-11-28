package la

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"maps"
	"slices"
	"sort"
	"strconv"
	"testing"
)

func TestEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Collect(Entries(maps.All(map[string]int{"foo": 1, "bar": 2})))

	sort.Slice(r1, func(i, j int) bool {
		return r1[i].Value < r1[j].Value
	})
	is.EqualValues(r1, []lo.Entry[string, int]{
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
	t.Parallel()
	is := assert.New(t)

	r1 := CollectMap(FromEntries(slices.Values([]lo.Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
	})))

	is.Len(r1, 2)
	is.Equal(r1["foo"], 1)
	is.Equal(r1["bar"], 2)
}

func TestToPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Collect(Pairs(maps.All(map[string]int{"baz": 3, "qux": 4})))

	sort.Slice(r1, func(i, j int) bool {
		return r1[i].Value < r1[j].Value
	})
	is.EqualValues(r1, []lo.Entry[string, int]{
		{
			Key:   "baz",
			Value: 3,
		},
		{
			Key:   "qux",
			Value: 4,
		},
	})
}

func TestFromPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := CollectMap(FromPairs(slices.Values([]lo.Entry[string, int]{
		{
			Key:   "baz",
			Value: 3,
		},
		{
			Key:   "qux",
			Value: 4,
		},
	})))

	is.Len(r1, 2)
	is.Equal(r1["baz"], 3)
	is.Equal(r1["qux"], 4)
}

func TestSeq2ToSeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Seq2ToSeq(maps.All(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}), func(k int, v int) string {
		return fmt.Sprintf("%d_%d", k, v)
	}))
	result2 := slices.Collect(Seq2ToSeq(maps.All(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}), func(k int, _ int) string {
		return strconv.FormatInt(int64(k), 10)
	}))

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.ElementsMatch(result1, []string{"1_5", "2_6", "3_7", "4_8"})
	is.ElementsMatch(result2, []string{"1", "2", "3", "4"})
}
