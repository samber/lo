//go:build go1.23

package it

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		seq      iter.Seq[int]
		target   int
		expected bool
	}{
		{name: "found", seq: values(0, 1, 2, 3, 4, 5), target: 5, expected: true},
		{name: "not found", seq: values(0, 1, 2, 3, 4, 5), target: 6, expected: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Contains(tt.seq, tt.target))
		})
	}
}

func TestContainsBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type a struct {
		A int
		B string
	}

	t.Run("struct type", func(t *testing.T) {
		t.Parallel()

		a1 := values(a{A: 1, B: "1"}, a{A: 2, B: "2"}, a{A: 3, B: "3"})

		tests := []struct {
			name      string
			predicate func(a) bool
			expected  bool
		}{
			{name: "no match", predicate: func(t a) bool { return t.A == 1 && t.B == "2" }, expected: false},
			{name: "match", predicate: func(t a) bool { return t.A == 2 && t.B == "2" }, expected: true},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is.Equal(tt.expected, ContainsBy(a1, tt.predicate))
			})
		}
	})

	t.Run("string type", func(t *testing.T) {
		t.Parallel()

		a2 := values("aaa", "bbb", "ccc")

		tests := []struct {
			name      string
			predicate func(string) bool
			expected  bool
		}{
			{name: "match", predicate: func(t string) bool { return t == "ccc" }, expected: true},
			{name: "no match", predicate: func(t string) bool { return t == "ddd" }, expected: false},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				is.Equal(tt.expected, ContainsBy(a2, tt.predicate))
			})
		}
	})
}

func TestEvery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		args     []int
		expected bool
	}{
		{name: "both present", args: []int{0, 2}, expected: true},
		{name: "one present one missing", args: []int{0, 6}, expected: false},
		{name: "both missing", args: []int{-1, 6}, expected: false},
		{name: "no args", args: nil, expected: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Every(values(0, 1, 2, 3, 4, 5), tt.args...))
		})
	}
}

func TestEveryBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		seq       iter.Seq[int]
		predicate func(int) bool
		expected  bool
	}{
		{name: "all match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 5 }, expected: true},
		{name: "some match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 3 }, expected: false},
		{name: "none match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 0 }, expected: false},
		{name: "empty collection", seq: values[int](), predicate: func(x int) bool { return x < 5 }, expected: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, EveryBy(tt.seq, tt.predicate))
		})
	}
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		args     []int
		expected bool
	}{
		{name: "both present", args: []int{0, 2}, expected: true},
		{name: "one present one missing", args: []int{0, 6}, expected: true},
		{name: "both missing", args: []int{-1, 6}, expected: false},
		{name: "no args", args: nil, expected: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Some(values(0, 1, 2, 3, 4, 5), tt.args...))
		})
	}
}

func TestSomeBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		seq       iter.Seq[int]
		predicate func(int) bool
		expected  bool
	}{
		{name: "some match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 5 }, expected: true},
		{name: "one match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 3 }, expected: true},
		{name: "none match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 0 }, expected: false},
		{name: "empty collection", seq: values[int](), predicate: func(x int) bool { return x < 5 }, expected: false},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, SomeBy(tt.seq, tt.predicate))
		})
	}
}

func TestNone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		args     []int
		expected bool
	}{
		{name: "both present", args: []int{0, 2}, expected: false},
		{name: "one present one missing", args: []int{0, 6}, expected: false},
		{name: "both missing", args: []int{-1, 6}, expected: true},
		{name: "no args", args: nil, expected: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, None(values(0, 1, 2, 3, 4, 5), tt.args...))
		})
	}
}

func TestNoneBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		seq       iter.Seq[int]
		predicate func(int) bool
		expected  bool
	}{
		{name: "some match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 5 }, expected: false},
		{name: "one match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 3 }, expected: false},
		{name: "none match", seq: values(1, 2, 3, 4), predicate: func(x int) bool { return x < 0 }, expected: true},
		{name: "empty collection", seq: values[int](), predicate: func(x int) bool { return x < 5 }, expected: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, NoneBy(tt.seq, tt.predicate))
		})
	}
}

func TestIntersect(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		inputs   []iter.Seq[int]
		expected []int
	}{
		{name: "no sequences", inputs: []iter.Seq[int]{}, expected: nil},
		{name: "single sequence", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5)}, expected: []int{0, 1, 2, 3, 4, 5}},
		{name: "one common element", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(0, 6)}, expected: []int{0}},
		{name: "no common element", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(-1, 6)}, expected: nil},
		{name: "duplicates in first sequence", inputs: []iter.Seq[int]{values(0, 6, 0), values(0, 1, 2, 3, 4, 5)}, expected: []int{0}},
		{name: "duplicates in second sequence", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(0, 6, 0)}, expected: []int{0}},
		{name: "three sequences", inputs: []iter.Seq[int]{values(0, 1, 2), values(1, 2, 3), values(2, 3, 4)}, expected: []int{2}},
		{name: "four sequences no overlap", inputs: []iter.Seq[int]{values(0, 1, 2), values(1, 2, 3), values(2, 3, 4), values(3, 4, 5)}, expected: nil},
		{name: "five sequences with duplicate no overlap", inputs: []iter.Seq[int]{values(0, 1, 2), values(0, 1, 2), values(1, 2, 3), values(2, 3, 4), values(3, 4, 5)}, expected: nil},
		{name: "duplicates within single sequence", inputs: []iter.Seq[int]{values(0, 1, 1)}, expected: []int{0, 1}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := slices.Collect(Intersect(tt.inputs...))
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Intersect(allStrings, allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestIntersectBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	transform := strconv.Itoa

	tests := []struct {
		name     string
		inputs   []iter.Seq[int]
		expected []int
	}{
		{name: "no sequences", inputs: []iter.Seq[int]{}, expected: nil},
		{name: "single sequence", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5)}, expected: []int{0, 1, 2, 3, 4, 5}},
		{name: "one common element", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(0, 6)}, expected: []int{0}},
		{name: "no common element", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(-1, 6)}, expected: nil},
		{name: "duplicates in first sequence", inputs: []iter.Seq[int]{values(0, 6, 0), values(0, 1, 2, 3, 4, 5)}, expected: []int{0}},
		{name: "duplicates in second sequence", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(0, 6, 0)}, expected: []int{0}},
		{name: "three sequences", inputs: []iter.Seq[int]{values(0, 1, 2), values(1, 2, 3), values(2, 3, 4)}, expected: []int{2}},
		{name: "four sequences no overlap", inputs: []iter.Seq[int]{values(0, 1, 2), values(1, 2, 3), values(2, 3, 4), values(3, 4, 5)}, expected: nil},
		{name: "five sequences with duplicate no overlap", inputs: []iter.Seq[int]{values(0, 1, 2), values(0, 1, 2), values(1, 2, 3), values(2, 3, 4), values(3, 4, 5)}, expected: nil},
		{name: "duplicates within single sequence", inputs: []iter.Seq[int]{values(0, 1, 1)}, expected: []int{0, 1}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := slices.Collect(IntersectBy(transform, tt.inputs...))
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := IntersectBy(func(s string) string { return s + s }, allStrings, allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestUnion(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		inputs   []iter.Seq[int]
		expected []int
	}{
		{name: "two sequences with new elements", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(0, 2, 10)}, expected: []int{0, 1, 2, 3, 4, 5, 10}},
		{name: "two disjoint sequences", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(6, 7)}, expected: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{name: "second sequence empty", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values[int]()}, expected: []int{0, 1, 2, 3, 4, 5}},
		{name: "second sequence with duplicates", inputs: []iter.Seq[int]{values(0, 1, 2), values(0, 1, 2, 3, 3)}, expected: []int{0, 1, 2, 3}},
		{name: "two identical sequences", inputs: []iter.Seq[int]{values(0, 1, 2), values(0, 1, 2)}, expected: []int{0, 1, 2}},
		{name: "two empty sequences", inputs: []iter.Seq[int]{values[int](), values[int]()}, expected: nil},
		{name: "three sequences with new elements", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(0, 2, 10), values(0, 1, 11)}, expected: []int{0, 1, 2, 3, 4, 5, 10, 11}},
		{name: "three disjoint sequences", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values(6, 7), values(8, 9)}, expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "three sequences two empty", inputs: []iter.Seq[int]{values(0, 1, 2, 3, 4, 5), values[int](), values[int]()}, expected: []int{0, 1, 2, 3, 4, 5}},
		{name: "three identical sequences", inputs: []iter.Seq[int]{values(0, 1, 2), values(0, 1, 2), values(0, 1, 2)}, expected: []int{0, 1, 2}},
		{name: "three empty sequences", inputs: []iter.Seq[int]{values[int](), values[int](), values[int]()}, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := slices.Collect(Union(tt.inputs...))
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Union(allStrings, allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestWithout(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		seq      iter.Seq[int]
		exclude  []int
		expected []int
	}{
		{name: "removes some elements", seq: values(0, 2, 10), exclude: []int{0, 1, 2, 3, 4, 5}, expected: []int{10}},
		{name: "removes one element", seq: values(0, 7), exclude: []int{0, 1, 2, 3, 4, 5}, expected: []int{7}},
		{name: "empty sequence", seq: values[int](), exclude: []int{0, 1, 2, 3, 4, 5}, expected: nil},
		{name: "removes all elements", seq: values(0, 1, 2), exclude: []int{0, 1, 2}, expected: nil},
		{name: "empty sequence no exclude", seq: values[int](), exclude: nil, expected: nil},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := slices.Collect(Without(tt.seq, tt.exclude...))
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := Without(allStrings, "")
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestWithoutBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type User struct {
		Name string
		Age  int
	}

	// result1 and result3 share the same generic instantiation (K=string), so
	// they're grouped in one table; result2 uses K=int and can't share the
	// row type, so it gets its own subtest per the multi-instantiation rule.
	t.Run("string key", func(t *testing.T) {
		t.Parallel()

		byName := func(item User) string { return item.Name }

		tests := []struct {
			name     string
			seq      iter.Seq[User]
			exclude  []string
			expected []User
		}{
			{name: "excludes by name", seq: values(User{Name: "nick"}, User{Name: "peter"}), exclude: []string{"nick", "lily"}, expected: []User{{Name: "peter"}}},
			{name: "empty sequence no exclude", seq: values[User](), exclude: nil, expected: nil},
		}

		for _, tt := range tests {
			tt := tt //nolint:modernize
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				result := slices.Collect(WithoutBy(tt.seq, byName, tt.exclude...))
				if tt.expected == nil {
					is.Empty(result)
				} else {
					is.Equal(tt.expected, result)
				}
			})
		}
	})

	t.Run("int key", func(t *testing.T) {
		t.Parallel()

		result := WithoutBy(values[User](), func(item User) int { return item.Age }, 1, 2, 3)
		is.Empty(slices.Collect(result))
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := WithoutBy(allStrings, func(string) string { return "" })
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestWithoutNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		seq      iter.Seq[int]
		nths     []int
		expected []int
	}{
		{name: "removes indices", seq: values(5, 6, 7), nths: []int{1, 0}, expected: []int{7}},
		{name: "no indices", seq: values(1, 2), nths: nil, expected: []int{1, 2}},
		{name: "empty sequence", seq: values[int](), nths: nil, expected: nil},
		{name: "out of range indices", seq: values(0, 1, 2, 3), nths: []int{-1, 4}, expected: []int{0, 1, 2, 3}},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := slices.Collect(WithoutNth(tt.seq, tt.nths...))
			if tt.expected == nil {
				is.Empty(result)
			} else {
				is.Equal(tt.expected, result)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings iter.Seq[string]
		allStrings := myStrings(values("", "foo", "bar"))
		nonempty := WithoutNth(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestElementsMatch(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		a        iter.Seq[int]
		b        iter.Seq[int]
		expected bool
	}{
		{name: "empty vs non-empty", a: values[int](), b: values(1), expected: false},
		{name: "different single elements", a: values(1), b: values(2), expected: false},
		{name: "different lengths", a: values(1), b: values(1, 2), expected: false},
		{name: "different element counts", a: values(1, 1, 2), b: values(2, 2, 1), expected: false},
		{name: "same single element", a: values(1), b: values(1), expected: true},
		{name: "same repeated elements", a: values(1, 1), b: values(1, 1), expected: true},
		{name: "same elements different order", a: values(1, 2), b: values(2, 1), expected: true},
		{name: "same multiset different order", a: values(1, 1, 2), b: values(1, 2, 1), expected: true},
	}

	for _, tt := range tests {
		tt := tt //nolint:modernize
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, ElementsMatch(tt.a, tt.b))
		})
	}
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
