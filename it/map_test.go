//go:build go1.23

package it

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

	tests := []struct {
		name     string
		maps     []map[string]int
		expected []string
	}{
		{name: "single map", maps: []map[string]int{{"foo": 1, "bar": 2}}, expected: []string{"bar", "foo"}},
		{name: "empty map", maps: []map[string]int{{}}, expected: []string{}},
		{name: "multiple maps", maps: []map[string]int{{"foo": 1, "bar": 2}, {"baz": 3}}, expected: []string{"bar", "baz", "foo"}},
		{name: "no maps", maps: nil, expected: []string{}},
		{name: "duplicate keys across maps", maps: []map[string]int{{"foo": 1, "bar": 2}, {"bar": 3}}, expected: []string{"bar", "bar", "foo"}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.ElementsMatch(tt.expected, slices.Collect(Keys(tt.maps...)))
		})
	}
}

func TestUniqKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		maps     []map[string]int
		expected []string
		exact    bool
	}{
		{name: "single map", maps: []map[string]int{{"foo": 1, "bar": 2}}, expected: []string{"bar", "foo"}},
		{name: "empty map", maps: []map[string]int{{}}, expected: []string{}},
		{name: "multiple maps", maps: []map[string]int{{"foo": 1, "bar": 2}, {"baz": 3}}, expected: []string{"bar", "baz", "foo"}},
		{name: "no maps", maps: nil, expected: []string{}},
		{name: "duplicate keys across maps are deduplicated", maps: []map[string]int{{"foo": 1, "bar": 2}, {"foo": 1, "bar": 3}}, expected: []string{"bar", "foo"}},
		{name: "preserves first-seen order", maps: []map[string]int{{"foo": 1}, {"bar": 3}}, expected: []string{"foo", "bar"}, exact: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := slices.Collect(UniqKeys(tt.maps...))
			if tt.exact {
				is.Equal(tt.expected, result)
			} else {
				is.ElementsMatch(tt.expected, result)
			}
		})
	}
}

func TestValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		maps     []map[string]int
		expected []int
	}{
		{name: "single map", maps: []map[string]int{{"foo": 1, "bar": 2}}, expected: []int{1, 2}},
		{name: "empty map", maps: []map[string]int{{}}, expected: []int{}},
		{name: "multiple maps", maps: []map[string]int{{"foo": 1, "bar": 2}, {"baz": 3}}, expected: []int{1, 2, 3}},
		{name: "no maps", maps: nil, expected: []int{}},
		{name: "duplicate keys across maps", maps: []map[string]int{{"foo": 1, "bar": 2}, {"foo": 1, "bar": 3}}, expected: []int{1, 1, 2, 3}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.ElementsMatch(tt.expected, slices.Collect(Values(tt.maps...)))
		})
	}
}

func TestUniqValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		maps     []map[string]int
		expected []int
		exact    bool
	}{
		{name: "single map", maps: []map[string]int{{"foo": 1, "bar": 2}}, expected: []int{1, 2}},
		{name: "empty map", maps: []map[string]int{{}}, expected: []int{}},
		{name: "multiple maps", maps: []map[string]int{{"foo": 1, "bar": 2}, {"baz": 3}}, expected: []int{1, 2, 3}},
		{name: "no maps", maps: nil, expected: []int{}},
		{name: "duplicate values across maps", maps: []map[string]int{{"foo": 1, "bar": 2}, {"foo": 1, "bar": 3}}, expected: []int{1, 2, 3}},
		{name: "duplicate values within and across maps", maps: []map[string]int{{"foo": 1, "bar": 1}, {"foo": 1, "bar": 3}}, expected: []int{1, 3}},
		{name: "preserves first-seen order", maps: []map[string]int{{"foo": 1}, {"bar": 3}}, expected: []int{1, 3}, exact: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := slices.Collect(UniqValues(tt.maps...))
			if tt.exact {
				is.Equal(tt.expected, result)
			} else {
				is.ElementsMatch(tt.expected, result)
			}
		})
	}
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

func TestInvert(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("no collisions", func(t *testing.T) {
		t.Parallel()
		r1 := Invert(maps.All(map[string]int{"a": 1, "b": 2}))
		is.Equal(map[int]string{1: "a", 2: "b"}, maps.Collect(r1))
	})

	t.Run("colliding values keep one key", func(t *testing.T) {
		t.Parallel()
		r2 := Invert(maps.All(map[string]int{"a": 1, "b": 2, "c": 1}))
		is.Len(maps.Collect(r2), 2)
	})
}

func TestAssign(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("merges maps with later keys overriding earlier ones", func(t *testing.T) {
		t.Parallel()
		result1 := Assign(values(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4}))
		is.Equal(map[string]int{"a": 1, "b": 3, "c": 4}, result1)
	})

	t.Run("preserves custom map type", func(t *testing.T) {
		t.Parallel()
		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after := Assign(values(before, before))
		is.IsType(myMap{}, after, "type preserved")
	})
}

func TestChunkEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("chunk counts", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name        string
			input       map[string]int
			size        int
			expectedLen int
		}{
			{name: "5 entries chunked by 2", input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, size: 2, expectedLen: 3},
			{name: "5 entries chunked by 3", input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, size: 3, expectedLen: 2},
			{name: "empty map", input: map[string]int{}, size: 2, expectedLen: 0},
			{name: "single entry chunked by 2", input: map[string]int{"a": 1}, size: 2, expectedLen: 1},
			{name: "2 entries chunked by 1", input: map[string]int{"a": 1, "b": 2}, size: 1, expectedLen: 2},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is.Len(slices.Collect(ChunkEntries(tt.input, tt.size)), tt.expectedLen)
			})
		}
	})

	t.Run("panics on non-positive size", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name string
			size int
		}{
			{name: "zero size", size: 0},
			{name: "negative size", size: -1},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is.PanicsWithValue("it.ChunkEntries: size must be greater than 0", func() {
					ChunkEntries(map[string]int{"a": 1}, tt.size)
				})
			})
		}
	})

	t.Run("chunks a map of struct values", func(t *testing.T) {
		t.Parallel()

		type myStruct struct {
			Name  string
			Value int
		}

		allStructs := []myStruct{{"one", 1}, {"two", 2}, {"three", 3}}
		nonempty := ChunkEntries(map[string]myStruct{"a": allStructs[0], "b": allStructs[1], "c": allStructs[2]}, 2)
		is.Len(slices.Collect(nonempty), 2)
	})

	t.Run("mutating a returned chunk does not affect the original map", func(t *testing.T) {
		t.Parallel()

		originalMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
		result := slices.Collect(ChunkEntries(originalMap, 2))
		for k := range result[0] {
			result[0][k] = 10
		}
		is.Equal(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, originalMap)
	})
}

func TestMapToSeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		input     map[int]int
		transform func(k, v int) string
		expected  []string
	}{
		{
			name:      "formats key and value",
			input:     map[int]int{1: 5, 2: 6, 3: 7, 4: 8},
			transform: func(k, v int) string { return fmt.Sprintf("%d_%d", k, v) },
			expected:  []string{"1_5", "2_6", "3_7", "4_8"},
		},
		{
			name:      "formats key only",
			input:     map[int]int{1: 5, 2: 6, 3: 7, 4: 8},
			transform: func(k, _ int) string { return strconv.FormatInt(int64(k), 10) },
			expected:  []string{"1", "2", "3", "4"},
		},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.ElementsMatch(tt.expected, slices.Collect(MapToSeq(tt.input, tt.transform)))
		})
	}
}

func TestFilterMapToSeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    map[int]int
		filter   func(k, v int) (string, bool)
		expected []string
	}{
		{
			name:     "formats key and value for even keys",
			input:    map[int]int{1: 5, 2: 6, 3: 7, 4: 8},
			filter:   func(k, v int) (string, bool) { return fmt.Sprintf("%d_%d", k, v), k%2 == 0 },
			expected: []string{"2_6", "4_8"},
		},
		{
			name:     "formats key only for even keys",
			input:    map[int]int{1: 5, 2: 6, 3: 7, 4: 8},
			filter:   func(k, _ int) (string, bool) { return strconv.FormatInt(int64(k), 10), k%2 == 0 },
			expected: []string{"2", "4"},
		},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.ElementsMatch(tt.expected, slices.Collect(FilterMapToSeq(tt.input, tt.filter)))
		})
	}
}

func TestFilterKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("int keys filtered by matching string value", func(t *testing.T) {
		t.Parallel()
		result1 := FilterKeys(map[int]string{1: "foo", 2: "bar", 3: "baz"}, func(k int, v string) bool {
			return v == "foo"
		})
		is.Equal([]int{1}, slices.Collect(result1))
	})

	t.Run("string keys with predicate always false", func(t *testing.T) {
		t.Parallel()
		result2 := FilterKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
			return false
		})
		is.Empty(slices.Collect(result2))
	})
}

func TestFilterValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	t.Run("string values filtered by matching predicate", func(t *testing.T) {
		t.Parallel()
		result1 := FilterValues(map[int]string{1: "foo", 2: "bar", 3: "baz"}, func(k int, v string) bool {
			return v == "foo"
		})
		is.Equal([]string{"foo"}, slices.Collect(result1))
	})

	t.Run("int values with predicate always false", func(t *testing.T) {
		t.Parallel()
		result2 := FilterValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
			return false
		})
		is.Empty(slices.Collect(result2))
	})
}

func TestSeq2KeyToSeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{name: "non-empty map", input: map[string]int{"foo": 4, "bar": 5}, expected: []string{"foo", "bar"}},
		{name: "empty map", input: map[string]int{}, expected: []string{}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.ElementsMatch(tt.expected, slices.Collect(Seq2KeyToSeq(maps.All(tt.input))))
		})
	}
}

func TestSeq2ValueToSeq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    map[string]int
		expected []int
	}{
		{name: "non-empty map", input: map[string]int{"foo": 4, "bar": 5}, expected: []int{4, 5}},
		{name: "empty map", input: map[string]int{}, expected: []int{}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.ElementsMatch(tt.expected, slices.Collect(Seq2ValueToSeq(maps.All(tt.input))))
		})
	}
}

func BenchmarkAssign(b *testing.B) {
	counts := []int{32768, 1024, 128, 32, 2}

	allDifferentMap := func(b *testing.B, n int) []map[string]int {
		b.Helper()
		defer b.ResetTimer()
		m := make([]map[string]int, 0, n)
		for i := range n {
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
		m := make([]map[string]int, 0, n)
		for range n {
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
					for range b.N {
						result := Assign(values(tc.in...))
						_ = result
					}
				})
			}
		})
	}
}
