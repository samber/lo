package iter

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(Keys(map[string]int{"foo": 1, "bar": 2}))
	is.ElementsMatch(r1, []string{"bar", "foo"})

	r2 := slices.Collect(Keys(map[string]int{}))
	is.Empty(r2)

	r3 := slices.Collect(Keys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3}))
	is.ElementsMatch(r3, []string{"bar", "baz", "foo"})

	r4 := slices.Collect(Keys[string, int]())
	is.Empty(r4)

	r5 := slices.Collect(Keys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"bar": 3}))
	is.ElementsMatch(r5, []string{"bar", "bar", "foo"})
}

func TestUniqKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(UniqKeys(map[string]int{"foo": 1, "bar": 2}))
	is.ElementsMatch(r1, []string{"bar", "foo"})

	r2 := slices.Collect(UniqKeys(map[string]int{}))
	is.Empty(r2)

	r3 := slices.Collect(UniqKeys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3}))
	is.ElementsMatch(r3, []string{"bar", "baz", "foo"})

	r4 := slices.Collect(UniqKeys[string, int]())
	is.Empty(r4)

	r5 := slices.Collect(UniqKeys(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3}))
	is.ElementsMatch(r5, []string{"bar", "foo"})

	// check order
	r6 := slices.Collect(UniqKeys(map[string]int{"foo": 1}, map[string]int{"bar": 3}))
	is.Equal([]string{"foo", "bar"}, r6)
}

func TestValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(Values(map[string]int{"foo": 1, "bar": 2}))
	is.ElementsMatch(r1, []int{1, 2})

	r2 := slices.Collect(Values(map[string]int{}))
	is.Empty(r2)

	r3 := slices.Collect(Values(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3}))
	is.ElementsMatch(r3, []int{1, 2, 3})

	r4 := slices.Collect(Values[string, int]())
	is.Empty(r4)

	r5 := slices.Collect(Values(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3}))
	is.ElementsMatch(r5, []int{1, 1, 2, 3})
}

func TestUniqValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(UniqValues(map[string]int{"foo": 1, "bar": 2}))
	is.ElementsMatch(r1, []int{1, 2})

	r2 := slices.Collect(UniqValues(map[string]int{}))
	is.Empty(r2)

	r3 := slices.Collect(UniqValues(map[string]int{"foo": 1, "bar": 2}, map[string]int{"baz": 3}))
	is.ElementsMatch(r3, []int{1, 2, 3})

	r4 := slices.Collect(UniqValues[string, int]())
	is.Empty(r4)

	r5 := slices.Collect(UniqValues(map[string]int{"foo": 1, "bar": 2}, map[string]int{"foo": 1, "bar": 3}))
	is.ElementsMatch(r5, []int{1, 2, 3})

	r6 := slices.Collect(UniqValues(map[string]int{"foo": 1, "bar": 1}, map[string]int{"foo": 1, "bar": 3}))
	is.ElementsMatch(r6, []int{1, 3})

	// check order
	r7 := slices.Collect(UniqValues(map[string]int{"foo": 1}, map[string]int{"bar": 3}))
	is.Equal([]int{1, 3}, r7)
}

func TestEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := maps.Collect(Entries(map[string]int{"foo": 1, "bar": 2}))
	is.Equal(map[string]int{"foo": 1, "bar": 2}, r1)
}

func TestToPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := maps.Collect(ToPairs(map[string]int{"foo": 1, "bar": 2}))
	is.Equal(map[string]int{"foo": 1, "bar": 2}, r1)
}

func TestFromEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FromEntries(maps.All(map[string]int{"foo": 1, "bar": 2}))

	is.Len(r1, 2)
	is.Equal(1, r1["foo"])
	is.Equal(2, r1["bar"])
}

func TestFromPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FromPairs(maps.All(map[string]int{"baz": 3, "qux": 4}))

	is.Len(r1, 2)
	is.Equal(3, r1["baz"])
	is.Equal(4, r1["qux"])
}

func TestAssign(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Assign(values(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4}))
	is.Equal(map[string]int{"a": 1, "b": 3, "c": 4}, result1)

	type myMap map[string]int
	before := myMap{"": 0, "foobar": 6, "baz": 3}
	after := Assign(values(before, before))
	is.IsType(myMap{}, after, "type preserved")
}

func TestChunkEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ChunkEntries(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, 2)
	result2 := ChunkEntries(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, 3)
	result3 := ChunkEntries(map[string]int{}, 2)
	result4 := ChunkEntries(map[string]int{"a": 1}, 2)
	result5 := ChunkEntries(map[string]int{"a": 1, "b": 2}, 1)

	is.Len(slices.Collect(result1), 3)
	is.Len(slices.Collect(result2), 2)
	is.Empty(slices.Collect(result3))
	is.Len(slices.Collect(result4), 1)
	is.Len(slices.Collect(result5), 2)

	is.PanicsWithValue("iter.ChunkEntries: size must be greater than 0", func() {
		ChunkEntries(map[string]int{"a": 1}, 0)
	})
	is.PanicsWithValue("iter.ChunkEntries: size must be greater than 0", func() {
		ChunkEntries(map[string]int{"a": 1}, -1)
	})

	type myStruct struct {
		Name  string
		Value int
	}

	allStructs := []myStruct{{"one", 1}, {"two", 2}, {"three", 3}}
	nonempty := ChunkEntries(map[string]myStruct{"a": allStructs[0], "b": allStructs[1], "c": allStructs[2]}, 2)
	is.Len(slices.Collect(nonempty), 2)

	originalMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result6 := slices.Collect(ChunkEntries(originalMap, 2))
	for k := range result6[0] {
		result6[0][k] = 10
	}
	is.Equal(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, originalMap)
}

func TestFromMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FromMap(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k, v int) string {
		return fmt.Sprintf("%d_%d", k, v)
	})
	result2 := FromMap(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k, _ int) string {
		return strconv.FormatInt(int64(k), 10)
	})

	is.ElementsMatch(slices.Collect(result1), []string{"1_5", "2_6", "3_7", "4_8"})
	is.ElementsMatch(slices.Collect(result2), []string{"1", "2", "3", "4"})
}

func TestFilterFromMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FilterFromMap(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k, v int) (string, bool) {
		return fmt.Sprintf("%d_%d", k, v), k%2 == 0
	})
	result2 := FilterFromMap(map[int]int{1: 5, 2: 6, 3: 7, 4: 8}, func(k, _ int) (string, bool) {
		return strconv.FormatInt(int64(k), 10), k%2 == 0
	})

	is.ElementsMatch(slices.Collect(result1), []string{"2_6", "4_8"})
	is.ElementsMatch(slices.Collect(result2), []string{"2", "4"})
}

func TestFilterKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FilterKeys(map[int]string{1: "foo", 2: "bar", 3: "baz"}, func(k int, v string) bool {
		return v == "foo"
	})
	is.Equal([]int{1}, slices.Collect(result1))

	result2 := FilterKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Empty(slices.Collect(result2))
}

func TestFilterValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FilterValues(map[int]string{1: "foo", 2: "bar", 3: "baz"}, func(k int, v string) bool {
		return v == "foo"
	})
	is.Equal([]string{"foo"}, slices.Collect(result1))

	result2 := FilterValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	is.Empty(slices.Collect(result2))
}

func BenchmarkAssign(b *testing.B) {
	counts := []int{32768, 1024, 128, 32, 2}

	allDifferentMap := func(b *testing.B, n int) []map[string]int {
		b.Helper()
		defer b.ResetTimer()
		m := make([]map[string]int, 0)
		for i := 0; i < n; i++ {
			m = append(m, map[string]int{
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
				strconv.Itoa(i): i,
			},
			)
		}
		return m
	}

	allTheSameMap := func(b *testing.B, n int) []map[string]int {
		b.Helper()
		defer b.ResetTimer()
		m := make([]map[string]int, 0)
		for i := 0; i < n; i++ {
			m = append(m, map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
				"e": 5,
				"f": 6,
			},
			)
		}
		return m
	}

	for _, count := range counts {
		differentMap := allDifferentMap(b, count)
		sameMap := allTheSameMap(b, count)

		b.Run(strconv.Itoa(count), func(b *testing.B) {
			testCases := []struct {
				name string
				in   []map[string]int
			}{
				{"different", differentMap},
				{"same", sameMap},
			}

			for _, tc := range testCases {
				b.Run(tc.name, func(b *testing.B) {
					b.ResetTimer()
					for n := 0; n < b.N; n++ {
						result := Assign(values(tc.in...))
						_ = result
					}
				})
			}
		})
	}
}
