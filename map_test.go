package lo

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []map[string]int
		want  []string
	}{
		{name: "single map", input: []map[string]int{{"foo": 1, "bar": 2}}, want: []string{"bar", "foo"}},
		{name: "empty map", input: []map[string]int{{}}, want: []string{}},
		{name: "multiple maps", input: []map[string]int{{"foo": 1, "bar": 2}, {"baz": 3}}, want: []string{"bar", "baz", "foo"}},
		{name: "no maps passed", input: []map[string]int{}, want: []string{}},
		{name: "duplicate keys across maps", input: []map[string]int{{"foo": 1, "bar": 2}, {"bar": 3}}, want: []string{"bar", "bar", "foo"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := Keys(tt.input...)
			is.ElementsMatch(tt.want, got)
		})
	}
}

// TestUniqKeys_zero exercises the len(in) == 0 branch: no maps at all.
func TestUniqKeys_zero(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := UniqKeys[string, int]()
	is.Empty(r1)
}

// TestUniqKeys_single exercises the len(in) == 1 branch, which delegates to Keys().
func TestUniqKeys_single(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input map[string]int
		want  []string
	}{
		{name: "basic map", input: map[string]int{"foo": 1, "bar": 2}, want: []string{"bar", "foo"}},
		{name: "empty map", input: map[string]int{}, want: []string{}},
		{name: "nil map is a valid, empty map", input: nil, want: []string{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := UniqKeys(tt.input)
			is.ElementsMatch(tt.want, got)
		})
	}
}

// TestUniqKeys_multiple exercises the len(in) > 1 branch, which dedups via a seen-set.
func TestUniqKeys_multiple(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        []map[string]int
		want         []string
		orderMatters bool // true when the expected order is deterministic and must match exactly
	}{
		{name: "dedup across maps", input: []map[string]int{{"foo": 1, "bar": 2}, {"baz": 3}}, want: []string{"bar", "baz", "foo"}},
		{name: "dedup same key across maps", input: []map[string]int{{"foo": 1, "bar": 2}, {"foo": 1, "bar": 3}}, want: []string{"bar", "foo"}},
		{name: "check order", input: []map[string]int{{"foo": 1}, {"bar": 3}}, want: []string{"foo", "bar"}, orderMatters: true},
		{name: "multiple maps, all empty", input: []map[string]int{{}, {}}, want: []string{}},
		{name: "dedup across more than two maps, keeping first-occurrence order", input: []map[string]int{{"foo": 1}, {"foo": 2, "bar": 3}, {"bar": 4, "baz": 5}}, want: []string{"foo", "bar", "baz"}, orderMatters: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := UniqKeys(tt.input...)

			if tt.orderMatters {
				is.Equal(tt.want, got)
			} else {
				is.ElementsMatch(tt.want, got)
			}
		})
	}
}

func TestHasKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		m    map[string]int
		key  string
		want bool
	}{
		{name: "missing key", m: map[string]int{"foo": 1}, key: "bar", want: false},
		{name: "existing key", m: map[string]int{"foo": 1}, key: "foo", want: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := HasKey(tt.m, tt.key)
			is.Equal(tt.want, got)
		})
	}
}

func TestValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []map[string]int
		want  []int
	}{
		{name: "single map", input: []map[string]int{{"foo": 1, "bar": 2}}, want: []int{1, 2}},
		{name: "empty map", input: []map[string]int{{}}, want: []int{}},
		{name: "multiple maps", input: []map[string]int{{"foo": 1, "bar": 2}, {"baz": 3}}, want: []int{1, 2, 3}},
		{name: "no maps passed", input: []map[string]int{}, want: []int{}},
		{name: "duplicate keys across maps", input: []map[string]int{{"foo": 1, "bar": 2}, {"foo": 1, "bar": 3}}, want: []int{1, 1, 2, 3}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := Values(tt.input...)
			is.ElementsMatch(tt.want, got)
		})
	}
}

func TestUniqValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        []map[string]int
		want         []int
		orderMatters bool // true when the expected order is deterministic and must match exactly
	}{
		{name: "single map", input: []map[string]int{{"foo": 1, "bar": 2}}, want: []int{1, 2}},
		{name: "empty map", input: []map[string]int{{}}, want: []int{}},
		{name: "multiple maps", input: []map[string]int{{"foo": 1, "bar": 2}, {"baz": 3}}, want: []int{1, 2, 3}},
		{name: "no maps passed", input: []map[string]int{}, want: []int{}},
		{name: "dedup same key across maps, keep both distinct values", input: []map[string]int{{"foo": 1, "bar": 2}, {"foo": 1, "bar": 3}}, want: []int{1, 2, 3}},
		{name: "dedup identical values across maps", input: []map[string]int{{"foo": 1, "bar": 1}, {"foo": 1, "bar": 3}}, want: []int{1, 3}},
		{name: "check order", input: []map[string]int{{"foo": 1}, {"bar": 3}}, want: []int{1, 3}, orderMatters: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := UniqValues(tt.input...)

			if tt.orderMatters {
				is.Equal(tt.want, got)
			} else {
				is.ElementsMatch(tt.want, got)
			}
		})
	}
}

func TestValueOr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		m        map[string]int
		key      string
		fallback int
		want     int
	}{
		{name: "missing key returns fallback", m: map[string]int{"foo": 1}, key: "bar", fallback: 2, want: 2},
		{name: "existing key returns value", m: map[string]int{"foo": 1}, key: "foo", fallback: 2, want: 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := ValueOr(tt.m, tt.key, tt.fallback)
			is.Equal(tt.want, got)
		})
	}
}

func TestPickBy(t *testing.T) {
	t.Parallel()

	t.Run("filters entries by predicate", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := PickBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
			return value%2 == 1
		})

		is.Equal(map[string]int{"foo": 1, "baz": 3}, r1)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after := PickBy(before, func(key string, value int) bool { return true })
		is.IsType(after, before, "type preserved")
	})
}

func TestPickByErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         map[string]int
		predicate     func(string, int) (bool, error)
		want          map[string]int
		wantErr       string
		wantCallCount int // 0 means don't check (maps have no order)
	}{
		{
			name:  "filter odd values",
			input: map[string]int{"foo": 1, "bar": 2, "baz": 3},
			predicate: func(key string, value int) (bool, error) {
				return value%2 == 1, nil
			},
			want:          map[string]int{"foo": 1, "baz": 3},
			wantErr:       "",
			wantCallCount: 0, // map iteration order is not deterministic
		},
		{
			name:  "empty map",
			input: map[string]int{},
			predicate: func(key string, value int) (bool, error) {
				return true, nil
			},
			want:          map[string]int{},
			wantErr:       "",
			wantCallCount: 0,
		},
		{
			name:  "error on specific key",
			input: map[string]int{"foo": 1, "bar": 2, "baz": 3},
			predicate: func(key string, value int) (bool, error) {
				if key == "bar" {
					return false, errors.New("bar not allowed")
				}
				return true, nil
			},
			want:          nil,
			wantErr:       "bar not allowed",
			wantCallCount: 0, // map iteration order is not deterministic
		},
		{
			name:  "filter all out",
			input: map[string]int{"foo": 2, "bar": 4, "baz": 6},
			predicate: func(key string, value int) (bool, error) {
				return value%2 == 1, nil
			},
			want:          map[string]int{},
			wantErr:       "",
			wantCallCount: 0,
		},
		{
			name:  "filter all in",
			input: map[string]int{"foo": 1, "bar": 3, "baz": 5},
			predicate: func(key string, value int) (bool, error) {
				return value%2 == 1, nil
			},
			want:          map[string]int{"foo": 1, "bar": 3, "baz": 5},
			wantErr:       "",
			wantCallCount: 0,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got, err := PickByErr(tt.input, tt.predicate)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
			} else {
				is.NoError(err)
				is.Equal(tt.want, got)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after, err := PickByErr(before, func(key string, value int) (bool, error) {
			return true, nil
		})

		is.NoError(err)
		is.IsType(after, before, "type preserved")
	})
}

func TestPickByKeys(t *testing.T) {
	t.Parallel()

	t.Run("picks entries by keys", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := PickByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz", "qux"})

		is.Equal(map[string]int{"foo": 1, "baz": 3}, r1)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after := PickByKeys(before, []string{"foobar", "baz"})
		is.IsType(after, before, "type preserved")
	})
}

func TestPickByValues(t *testing.T) {
	t.Parallel()

	t.Run("picks entries by values", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := PickByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

		is.Equal(map[string]int{"foo": 1, "baz": 3}, r1)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after := PickByValues(before, []int{0, 3})
		is.IsType(after, before, "type preserved")
	})
}

func TestOmitBy(t *testing.T) {
	t.Parallel()

	t.Run("omits entries by predicate", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := OmitBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
			return value%2 == 1
		})

		is.Equal(map[string]int{"bar": 2}, r1)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after := PickBy(before, func(key string, value int) bool { return true })
		is.IsType(after, before, "type preserved")
	})
}

func TestOmitByErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         map[string]int
		predicate     func(string, int) (bool, error)
		want          map[string]int
		wantErr       string
		wantCallCount int // 0 means don't check (maps have no order)
	}{
		{
			name:  "omit odd values",
			input: map[string]int{"foo": 1, "bar": 2, "baz": 3},
			predicate: func(key string, value int) (bool, error) {
				return value%2 == 1, nil
			},
			want:          map[string]int{"bar": 2},
			wantErr:       "",
			wantCallCount: 0, // map iteration order is not deterministic
		},
		{
			name:  "empty map",
			input: map[string]int{},
			predicate: func(key string, value int) (bool, error) {
				return true, nil
			},
			want:          map[string]int{},
			wantErr:       "",
			wantCallCount: 0,
		},
		{
			name:  "error on specific key",
			input: map[string]int{"foo": 1, "bar": 2, "baz": 3},
			predicate: func(key string, value int) (bool, error) {
				if key == "bar" {
					return false, errors.New("bar not allowed")
				}
				return true, nil
			},
			want:          nil,
			wantErr:       "bar not allowed",
			wantCallCount: 0, // map iteration order is not deterministic
		},
		{
			name:  "omit all",
			input: map[string]int{"foo": 1, "bar": 3, "baz": 5},
			predicate: func(key string, value int) (bool, error) {
				return value%2 == 1, nil
			},
			want:          map[string]int{},
			wantErr:       "",
			wantCallCount: 0,
		},
		{
			name:  "omit none",
			input: map[string]int{"foo": 2, "bar": 4, "baz": 6},
			predicate: func(key string, value int) (bool, error) {
				return value%2 == 1, nil
			},
			want:          map[string]int{"foo": 2, "bar": 4, "baz": 6},
			wantErr:       "",
			wantCallCount: 0,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got, err := OmitByErr(tt.input, tt.predicate)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
			} else {
				is.NoError(err)
				is.Equal(tt.want, got)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after, err := OmitByErr(before, func(key string, value int) (bool, error) {
			return false, nil
		})

		is.NoError(err)
		is.IsType(after, before, "type preserved")
	})
}

func TestOmitByKeys(t *testing.T) {
	t.Parallel()

	t.Run("omits entries by keys", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := OmitByKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz", "qux"})

		is.Equal(map[string]int{"bar": 2}, r1)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after := OmitByKeys(before, []string{"foobar", "baz"})
		is.IsType(after, before, "type preserved")
	})
}

func TestOmitByValues(t *testing.T) {
	t.Parallel()

	t.Run("omits entries by values", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := OmitByValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})

		is.Equal(map[string]int{"bar": 2}, r1)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after := OmitByValues(before, []int{0, 3})
		is.IsType(after, before, "type preserved")
	})
}

func TestEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Entries(map[string]int{"foo": 1, "bar": 2})
	is.ElementsMatch(r1, []Entry[string, int]{
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

func TestToPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := ToPairs(map[string]int{"baz": 3, "qux": 4})
	is.ElementsMatch(r1, []Entry[string, int]{
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

func TestFromEntries(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FromEntries([]Entry[string, int]{
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
	is.Equal(1, r1["foo"])
	is.Equal(2, r1["bar"])
}

func TestFromPairs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FromPairs([]Entry[string, int]{
		{
			Key:   "baz",
			Value: 3,
		},
		{
			Key:   "qux",
			Value: 4,
		},
	})

	is.Len(r1, 2)
	is.Equal(3, r1["baz"])
	is.Equal(4, r1["qux"])
}

func TestInvert(t *testing.T) {
	t.Parallel()

	t.Run("unique values invert cleanly", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := Invert(map[string]int{"a": 1, "b": 2})
		is.Equal(map[int]string{1: "a", 2: "b"}, r1)
	})

	t.Run("duplicate values collapse - only length is deterministic", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r2 := Invert(map[string]int{"a": 1, "b": 2, "c": 1})
		is.Len(r2, 2)
	})
}

func TestAssign(t *testing.T) {
	t.Parallel()

	t.Run("merges maps, later maps win", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result1 := Assign(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})
		is.Equal(map[string]int{"a": 1, "b": 3, "c": 4}, result1)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myMap map[string]int
		before := myMap{"": 0, "foobar": 6, "baz": 3}
		after := Assign(before, before)
		is.IsType(after, before, "type preserved")
	})
}

func TestChunkEntries(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   map[string]int
		size    int
		wantLen int
	}{
		{name: "5 entries, size 2", input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, size: 2, wantLen: 3},
		{name: "5 entries, size 3", input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, size: 3, wantLen: 2},
		{name: "empty map", input: map[string]int{}, size: 2, wantLen: 0},
		{name: "1 entry, size 2", input: map[string]int{"a": 1}, size: 2, wantLen: 1},
		{name: "2 entries, size 1", input: map[string]int{"a": 1, "b": 2}, size: 1, wantLen: 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := ChunkEntries(tt.input, tt.size)
			is.Len(got, tt.wantLen)
		})
	}

	t.Run("panics when size is not greater than 0", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		is.PanicsWithValue("lo.ChunkEntries: size must be greater than 0", func() {
			ChunkEntries(map[string]int{"a": 1}, 0)
		})
		is.PanicsWithValue("lo.ChunkEntries: size must be greater than 0", func() {
			ChunkEntries(map[string]int{"a": 1}, -1)
		})
	})

	t.Run("supports non-primitive value types", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStruct struct {
			Name  string
			Value int
		}

		allStructs := []myStruct{{"one", 1}, {"two", 2}, {"three", 3}}
		nonempty := ChunkEntries(map[string]myStruct{"a": allStructs[0], "b": allStructs[1], "c": allStructs[2]}, 2)
		is.Len(nonempty, 2)
	})

	t.Run("does not mutate the original map", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		originalMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
		result := ChunkEntries(originalMap, 2)
		for k := range result[0] {
			result[0][k] = 10
		}
		is.Equal(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, originalMap)
	})
}

func TestMapKeys(t *testing.T) {
	t.Parallel()

	t.Run("constant key collapses map", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := MapKeys(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x, _ int) string {
			return "Hello"
		})
		is.Len(result, 1)
	})

	t.Run("format value as key", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := MapKeys(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(_, v int) string {
			return strconv.FormatInt(int64(v), 10)
		})
		is.Equal(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}, result)
	})
}

func TestMapKeysErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[int]int
		iteratee func(int, int) (string, error)
		want     map[string]int
		wantErr  string
	}{
		{
			name:  "format value as key",
			input: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			iteratee: func(_, v int) (string, error) {
				return strconv.FormatInt(int64(v), 10), nil
			},
			want: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4},
		},
		{
			name:  "error on specific key - returns first error",
			input: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			iteratee: func(v, _ int) (string, error) {
				if v == 3 {
					return "", fmt.Errorf("error at %d", v)
				}
				return strconv.FormatInt(int64(v), 10), nil
			},
			wantErr: "error at 3",
		},
		{
			name:     "empty map",
			input:    map[int]int{},
			iteratee: func(v, _ int) (string, error) { return strconv.FormatInt(int64(v), 10), nil },
			want:     map[string]int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got, err := MapKeysErr(tt.input, tt.iteratee)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
			} else {
				is.NoError(err)
				is.Equal(tt.want, got)
			}
		})
	}

	t.Run("constant key collapses map", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result, err := MapKeysErr(map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(x, _ int) (string, error) {
			return "Hello", nil
		})
		is.NoError(err)
		is.Len(result, 1)
	})

	t.Run("all keys collide - iteration order is non-deterministic", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		// check that value is one of expected values, since map iteration order is not deterministic
		result, err := MapKeysErr(map[int]int{1: 1, 2: 2, 3: 3}, func(_, _ int) (string, error) {
			return "same", nil
		})
		is.NoError(err)
		is.Len(result, 1)
		is.Contains(result, "same")
		is.Contains([]int{1, 2, 3}, result["same"])
	})
}

func TestMapValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[int]int
		iteratee func(int, int) string
		want     map[int]string
	}{
		{
			name:  "constant value transformation",
			input: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			iteratee: func(x, _ int) string {
				return "Hello"
			},
			want: map[int]string{1: "Hello", 2: "Hello", 3: "Hello", 4: "Hello"},
		},
		{
			name:  "format key as value",
			input: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			iteratee: func(x, _ int) string {
				return strconv.FormatInt(int64(x), 10)
			},
			want: map[int]string{1: "1", 2: "2", 3: "3", 4: "4"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := MapValues(tt.input, tt.iteratee)
			is.Equal(tt.want, got)
		})
	}
}

func TestMapValuesErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[int]int
		iteratee func(int, int) (string, error)
		want     map[int]string
		wantErr  string
	}{
		{
			name:  "successful transformation",
			input: map[int]int{1: 1, 2: 2, 3: 3},
			iteratee: func(x, _ int) (string, error) {
				return strconv.FormatInt(int64(x), 10), nil
			},
			want:    map[int]string{1: "1", 2: "2", 3: "3"},
			wantErr: "",
		},
		{
			name:  "constant value transformation",
			input: map[int]int{1: 1, 2: 2, 3: 3},
			iteratee: func(_, _ int) (string, error) {
				return "Hello", nil
			},
			want:    map[int]string{1: "Hello", 2: "Hello", 3: "Hello"},
			wantErr: "",
		},
		{
			name:  "error on specific value",
			input: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
			iteratee: func(x, _ int) (string, error) {
				if x == 3 {
					return "", fmt.Errorf("error at %d", x)
				}
				return strconv.FormatInt(int64(x), 10), nil
			},
			want:    nil,
			wantErr: "error at 3",
		},
		{
			name:  "error on first value",
			input: map[int]int{1: 1, 2: 2},
			iteratee: func(x, _ int) (string, error) {
				if x == 1 {
					return "", errors.New("cannot process 1")
				}
				return strconv.FormatInt(int64(x), 10), nil
			},
			want:    nil,
			wantErr: "cannot process 1",
		},
		{
			name:     "empty map",
			input:    map[int]int{},
			iteratee: func(x, _ int) (string, error) { return strconv.FormatInt(int64(x), 10), nil },
			want:     map[int]string{},
			wantErr:  "",
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got, err := MapValuesErr(tt.input, tt.iteratee)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
			} else {
				is.NoError(err)
				is.Equal(tt.want, got)
			}
		})
	}
}

func TestMapEntries(t *testing.T) {
	t.Parallel()

	t.Run("Normal", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r1 := MapEntries(map[string]int{"foo": 1, "bar": 2},
			func(k string, v int) (string, int) {
				return k, v + 1
			})
		is.Equal(map[string]int{"foo": 2, "bar": 3}, r1)

		r2 := MapEntries(map[string]int{"foo": 1, "bar": 2},
			func(k string, v int) (string, string) {
				return k, k + strconv.Itoa(v)
			})
		is.Equal(map[string]string{"foo": "foo1", "bar": "bar2"}, r2)

		r3 := MapEntries(map[string]int{"foo": 1, "bar": 2},
			func(k string, v int) (string, string) {
				return k, strconv.Itoa(v) + k
			})
		is.Equal(map[string]string{"foo": "1foo", "bar": "2bar"}, r3)
	})

	t.Run("NoMutation", func(t *testing.T) {
		t.Parallel()

		r1 := map[string]int{"foo": 1, "bar": 2}
		MapEntries(r1, func(k string, v int) (string, string) {
			return k, strconv.Itoa(v) + "!!"
		})
		assert.Equal(t, map[string]int{"foo": 1, "bar": 2}, r1)
	})

	t.Run("EmptyInput", func(t *testing.T) {
		t.Parallel()

		r1 := MapEntries(map[string]int{},
			func(k string, v int) (string, string) {
				return k, strconv.Itoa(v) + "!!"
			})
		assert.Empty(t, r1)

		r2 := MapEntries(map[string]any{},
			func(k string, v any) (string, any) {
				return k, v
			})
		assert.Empty(t, r2)
	})

	t.Run("Identity", func(t *testing.T) {
		t.Parallel()

		r1 := MapEntries(map[string]int{"foo": 1, "bar": 2},
			func(k string, v int) (string, int) {
				return k, v
			})
		assert.Equal(t, map[string]int{"foo": 1, "bar": 2}, r1)

		r2 := MapEntries(map[string]any{"foo": 1, "bar": "2", "ccc": true},
			func(k string, v any) (string, any) {
				return k, v
			})
		assert.Equal(t, map[string]any{"foo": 1, "bar": "2", "ccc": true}, r2)
	})

	t.Run("ToConstantEntry", func(t *testing.T) {
		t.Parallel()

		r1 := MapEntries(map[string]any{"foo": 1, "bar": "2", "ccc": true},
			func(k string, v any) (string, any) {
				return "key", "value"
			})
		assert.Equal(t, map[string]any{"key": "value"}, r1)

		r2 := MapEntries(map[string]any{"foo": 1, "bar": "2", "ccc": true},
			func(k string, v any) (string, any) {
				return "b", 5
			})
		assert.Equal(t, map[string]any{"b": 5}, r2)
	})

	// // because using range over map, the order is not guaranteed
	// // this test is not deterministic
	// t.Run("OverlappingKeys", func(t *testing.T) {
	// 		t.Parallel()
	//
	// 	r1 := MapEntries(map[string]any{"foo": 1, "foo2": 2, "Foo": 2, "Foo2": "2", "bar": "2", "ccc": true},
	// 		func(k string, v any) (string, any) {
	// 			return string(k[0]), v
	// 		})
	// 	assert.Equal(t, map[string]any{"F": "2", "b": "2", "c": true, "f": 1}, r1)
	//
	// 	r2 := MapEntries(map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"},
	// 		func(k, v string) (string, string) {
	// 			return v, k
	// 		})
	// 	assert.Equal(t, map[string]string{"1": "foo", "2": "bar", "true": "ccc"}, r2)
	// })

	t.Run("NormalMappers", func(t *testing.T) {
		t.Parallel()

		r1 := MapEntries(map[string]string{"foo": "1", "foo2": "2", "Foo": "2", "Foo2": "2", "bar": "2", "ccc": "true"},
			func(k, v string) (string, string) {
				return k, k + v
			})
		assert.Equal(t, map[string]string{"Foo": "Foo2", "Foo2": "Foo22", "bar": "bar2", "ccc": "ccctrue", "foo": "foo1", "foo2": "foo22"}, r1)

		type myStruct struct {
			name string
			age  int
		}
		r2 := MapEntries(map[string]myStruct{"1-11-1": {name: "foo", age: 1}, "2-22-2": {name: "bar", age: 2}},
			func(k string, v myStruct) (string, string) {
				return v.name, k
			})
		assert.Equal(t, map[string]string{"bar": "2-22-2", "foo": "1-11-1"}, r2)
	})
}

func TestMapEntriesErr(t *testing.T) {
	t.Parallel()

	t.Run("same types", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			input    map[string]int
			iteratee func(string, int) (string, int, error)
			want     map[string]int
			wantErr  string
		}{
			{
				name:  "increment values",
				input: map[string]int{"foo": 1, "bar": 2},
				iteratee: func(k string, v int) (string, int, error) {
					return k, v + 1, nil
				},
				want: map[string]int{"foo": 2, "bar": 3},
			},
			{
				name:  "error on specific key",
				input: map[string]int{"foo": 1, "bar": 2, "baz": 3},
				iteratee: func(k string, v int) (string, int, error) {
					if k == "bar" {
						return "", 0, errors.New("bar not allowed")
					}
					return k, v, nil
				},
				wantErr: "bar not allowed",
			},
			{
				name:  "empty map",
				input: map[string]int{},
				iteratee: func(k string, v int) (string, int, error) {
					return k, v, nil
				},
				want: map[string]int{},
			},
			{
				name:  "to constant entry",
				input: map[string]int{"foo": 1, "bar": 2, "baz": 3},
				iteratee: func(k string, v int) (string, int, error) {
					return "key", 0, nil
				},
				want: map[string]int{"key": 0},
			},
			{
				name:  "identity",
				input: map[string]int{"foo": 1, "bar": 2},
				iteratee: func(k string, v int) (string, int, error) {
					return k, v, nil
				},
				want: map[string]int{"foo": 1, "bar": 2},
			},
		}

		for _, tt := range tests {
			tt := tt // capture range variable
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				got, err := MapEntriesErr(tt.input, tt.iteratee)

				if tt.wantErr != "" {
					is.Error(err)
					is.Equal(tt.wantErr, err.Error())
					is.Nil(got)
				} else {
					is.NoError(err)
					is.Equal(tt.want, got)
				}
			})
		}
	})

	t.Run("different value type", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			input    map[string]int
			iteratee func(string, int) (string, string, error)
			want     map[string]string
			wantErr  string
		}{
			{
				name:  "transform both key and value",
				input: map[string]int{"foo": 1, "bar": 2},
				iteratee: func(k string, v int) (string, string, error) {
					return k, k + strconv.Itoa(v), nil
				},
				want: map[string]string{"foo": "foo1", "bar": "bar2"},
			},
			{
				name:  "error on specific value",
				input: map[string]int{"foo": 1, "bar": 2, "baz": 3},
				iteratee: func(k string, v int) (string, string, error) {
					if v == 2 {
						return "", "", fmt.Errorf("even value not allowed: %d", v)
					}
					return k, strconv.Itoa(v), nil
				},
				wantErr: "even value not allowed: 2",
			},
			{
				name:  "empty map",
				input: map[string]int{},
				iteratee: func(k string, v int) (string, string, error) {
					return k, strconv.Itoa(v), nil
				},
				want: map[string]string{},
			},
		}

		for _, tt := range tests {
			tt := tt // capture range variable
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				got, err := MapEntriesErr(tt.input, tt.iteratee)

				if tt.wantErr != "" {
					is.Error(err)
					is.Equal(tt.wantErr, err.Error())
					is.Nil(got)
				} else {
					is.NoError(err)
					is.Equal(tt.want, got)
				}
			})
		}
	})

	t.Run("invert map", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name     string
			input    map[string]int
			iteratee func(string, int) (int, string, error)
			want     map[int]string
			wantErr  string
		}{
			{
				name:  "successful invert",
				input: map[string]int{"a": 1, "b": 2},
				iteratee: func(k string, v int) (int, string, error) {
					return v, k, nil
				},
				want: map[int]string{1: "a", 2: "b"},
			},
			{
				name:  "error on specific key",
				input: map[string]int{"a": 1, "b": 2},
				iteratee: func(k string, v int) (int, string, error) {
					if k == "b" {
						return 0, "", errors.New("cannot invert b")
					}
					return v, k, nil
				},
				wantErr: "cannot invert b",
			},
		}

		for _, tt := range tests {
			tt := tt // capture range variable
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is := assert.New(t)

				got, err := MapEntriesErr(tt.input, tt.iteratee)

				if tt.wantErr != "" {
					is.Error(err)
					is.Equal(tt.wantErr, err.Error())
					is.Nil(got)
				} else {
					is.NoError(err)
					is.Equal(tt.want, got)
				}
			})
		}
	})
}

func TestMapToSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[int]int
		iteratee func(int, int) string
		want     []string
	}{
		{
			name:  "combine key and value",
			input: map[int]int{1: 5, 2: 6, 3: 7, 4: 8},
			iteratee: func(k, v int) string {
				return fmt.Sprintf("%d_%d", k, v)
			},
			want: []string{"1_5", "2_6", "3_7", "4_8"},
		},
		{
			name:  "key only",
			input: map[int]int{1: 5, 2: 6, 3: 7, 4: 8},
			iteratee: func(k, _ int) string {
				return strconv.FormatInt(int64(k), 10)
			},
			want: []string{"1", "2", "3", "4"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := MapToSlice(tt.input, tt.iteratee)
			is.ElementsMatch(tt.want, got)
		})
	}
}

func TestMapToSliceErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         map[int]int64
		iteratee      func(int, int64) (string, error)
		want          []string
		wantErr       string
		wantCallCount int // 0 means don't check (maps have no order)
	}{
		{
			name:  "successful transformation",
			input: map[int]int64{1: 5, 2: 6, 3: 7},
			iteratee: func(k int, v int64) (string, error) {
				return fmt.Sprintf("%d_%d", k, v), nil
			},
			want:          []string{"1_5", "2_6", "3_7"},
			wantErr:       "",
			wantCallCount: 0, // map iteration order is not deterministic
		},
		{
			name:  "empty map",
			input: map[int]int64{},
			iteratee: func(k int, v int64) (string, error) {
				return fmt.Sprintf("%d_%d", k, v), nil
			},
			want:          []string{},
			wantErr:       "",
			wantCallCount: 0,
		},
		{
			name:  "error on specific key",
			input: map[int]int64{1: 5, 2: 6, 3: 7},
			iteratee: func(k int, v int64) (string, error) {
				if k == 2 {
					return "", errors.New("key 2 not allowed")
				}
				return fmt.Sprintf("%d_%d", k, v), nil
			},
			want:          nil,
			wantErr:       "key 2 not allowed",
			wantCallCount: 0, // map iteration order is not deterministic
		},
		{
			name:  "constant value",
			input: map[int]int64{1: 5, 2: 6, 3: 7},
			iteratee: func(k int, v int64) (string, error) {
				return "constant", nil
			},
			want:          []string{"constant", "constant", "constant"},
			wantErr:       "",
			wantCallCount: 0,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got, err := MapToSliceErr(tt.input, tt.iteratee)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
			} else {
				is.NoError(err)
				is.ElementsMatch(tt.want, got)
			}
		})
	}
}

func TestFilterMapToSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[int]int
		iteratee func(int, int) (string, bool)
		want     []string
	}{
		{
			name:  "filter even keys, combine key and value",
			input: map[int]int{1: 5, 2: 6, 3: 7, 4: 8},
			iteratee: func(k, v int) (string, bool) {
				return fmt.Sprintf("%d_%d", k, v), k%2 == 0
			},
			want: []string{"2_6", "4_8"},
		},
		{
			name:  "filter even keys, key only",
			input: map[int]int{1: 5, 2: 6, 3: 7, 4: 8},
			iteratee: func(k, _ int) (string, bool) {
				return strconv.FormatInt(int64(k), 10), k%2 == 0
			},
			want: []string{"2", "4"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got := FilterMapToSlice(tt.input, tt.iteratee)
			is.ElementsMatch(tt.want, got)
		})
	}
}

func TestFilterMapToSliceErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    map[int]int64
		iteratee func(int, int64) (string, bool, error)
		want     []string
		wantErr  string
	}{
		{
			name:  "filter even keys and transform",
			input: map[int]int64{1: 5, 2: 6, 3: 7, 4: 8},
			iteratee: func(k int, v int64) (string, bool, error) {
				return fmt.Sprintf("%d_%d", k, v), k%2 == 0, nil
			},
			want: []string{"2_6", "4_8"},
		},
		{
			name:  "empty map",
			input: map[int]int64{},
			iteratee: func(k int, v int64) (string, bool, error) {
				return fmt.Sprintf("%d_%d", k, v), true, nil
			},
			want: []string{},
		},
		{
			name:  "filter all out",
			input: map[int]int64{1: 2, 2: 4, 3: 6},
			iteratee: func(k int, v int64) (string, bool, error) {
				return fmt.Sprintf("%d_%d", k, v), false, nil
			},
			want: []string{},
		},
		{
			name:  "filter all in",
			input: map[int]int64{1: 5, 2: 6, 3: 7},
			iteratee: func(k int, v int64) (string, bool, error) {
				return fmt.Sprintf("%d_%d", k, v), true, nil
			},
			want: []string{"1_5", "2_6", "3_7"},
		},
		{
			name:  "constant value",
			input: map[int]int64{1: 5, 2: 6, 3: 7},
			iteratee: func(k int, v int64) (string, bool, error) {
				return "constant", true, nil
			},
			want: []string{"constant", "constant", "constant"},
		},
		{
			name:  "error on specific key",
			input: map[int]int64{1: 5, 2: 6, 3: 7},
			iteratee: func(k int, v int64) (string, bool, error) {
				if k == 2 {
					return "", false, errors.New("key 2 not allowed")
				}
				return fmt.Sprintf("%d_%d", k, v), true, nil
			},
			wantErr: "key 2 not allowed",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got, err := FilterMapToSliceErr(tt.input, tt.iteratee)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
			} else {
				is.NoError(err)
				is.ElementsMatch(tt.want, got)
			}
		})
	}
}

func TestFilterKeys(t *testing.T) {
	t.Parallel()

	t.Run("int keys, string values", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := FilterKeys(map[int]string{1: "foo", 2: "bar", 3: "baz"}, func(k int, v string) bool {
			return v == "foo"
		})
		is.Equal([]int{1}, result)
	})

	t.Run("string keys, int values, filter all out", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := FilterKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
			return false
		})
		is.Empty(result)
	})
}

func TestFilterValues(t *testing.T) {
	t.Parallel()

	t.Run("int keys, string values", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := FilterValues(map[int]string{1: "foo", 2: "bar", 3: "baz"}, func(k int, v string) bool {
			return v == "foo"
		})
		is.Equal([]string{"foo"}, result)
	})

	t.Run("string keys, int values, filter all out", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := FilterValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
			return false
		})
		is.Empty(result)
	})
}

func TestFilterKeysErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     map[int]string
		predicate func(int, string) (bool, error)
		want      []int
		wantErr   string
	}{
		{
			name:  "filter by value",
			input: map[int]string{1: "foo", 2: "bar", 3: "baz"},
			predicate: func(k int, v string) (bool, error) {
				return v == "foo", nil
			},
			want: []int{1},
		},
		{
			name:  "empty map",
			input: map[int]string{},
			predicate: func(k int, v string) (bool, error) {
				return true, nil
			},
			want: []int{},
		},
		{
			name:  "filter all out",
			input: map[int]string{1: "foo", 2: "bar", 3: "baz"},
			predicate: func(k int, v string) (bool, error) {
				return false, nil
			},
			want: []int{},
		},
		{
			name:  "filter all in",
			input: map[int]string{1: "foo", 2: "bar", 3: "baz"},
			predicate: func(k int, v string) (bool, error) {
				return true, nil
			},
			want: []int{1, 2, 3},
		},
		{
			name:  "error on specific key",
			input: map[int]string{1: "foo", 2: "bar", 3: "baz"},
			predicate: func(k int, v string) (bool, error) {
				if k == 2 {
					return false, errors.New("key 2 not allowed")
				}
				return true, nil
			},
			wantErr: "key 2 not allowed",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got, err := FilterKeysErr(tt.input, tt.predicate)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
			} else {
				is.NoError(err)
				is.ElementsMatch(tt.want, got)
			}
		})
	}
}

func TestFilterValuesErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     map[int]string
		predicate func(int, string) (bool, error)
		want      []string
		wantErr   string
	}{
		{
			name:  "filter by value",
			input: map[int]string{1: "foo", 2: "bar", 3: "baz"},
			predicate: func(k int, v string) (bool, error) {
				return v == "foo", nil
			},
			want: []string{"foo"},
		},
		{
			name:  "empty map",
			input: map[int]string{},
			predicate: func(k int, v string) (bool, error) {
				return true, nil
			},
			want: []string{},
		},
		{
			name:  "filter all out",
			input: map[int]string{1: "foo", 2: "bar", 3: "baz"},
			predicate: func(k int, v string) (bool, error) {
				return false, nil
			},
			want: []string{},
		},
		{
			name:  "filter all in",
			input: map[int]string{1: "foo", 2: "bar", 3: "baz"},
			predicate: func(k int, v string) (bool, error) {
				return true, nil
			},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name:  "error on specific key",
			input: map[int]string{1: "foo", 2: "bar", 3: "baz"},
			predicate: func(k int, v string) (bool, error) {
				if k == 2 {
					return false, errors.New("key 2 not allowed")
				}
				return true, nil
			},
			wantErr: "key 2 not allowed",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			got, err := FilterValuesErr(tt.input, tt.predicate)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
			} else {
				is.NoError(err)
				is.ElementsMatch(tt.want, got)
			}
		})
	}
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
						result := Assign(tc.in...)
						_ = result
					}
				})
			}
		})
	}
}
