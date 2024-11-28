package la

import (
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"iter"
	"maps"
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(Filter(slices.Values([]int{1, 2, 3, 4}), func(x int) bool {
		return x%2 == 0
	}))
	is.Equal(r1, []int{2, 4})

	r2 := slices.Collect(Filter(slices.Values([]string{"", "foo", "", "bar", ""}), func(x string) bool {
		return len(x) > 0
	}))
	is.Equal(r2, []string{"foo", "bar"})
}

func TestFilter2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(Tuples(Filter2(Enumerate(slices.Values([]int{1, 2, 3, 4})), func(_, x int) bool {
		return x%2 == 0
	})))
	is.Equal(r1, []lo.Tuple2[int, int]{{1, 2}, {3, 4}})

	r2 := slices.Collect(Tuples(Filter2(Enumerate(slices.Values([]string{"", "foo", "", "bar", ""})), func(_ int, x string) bool {
		return len(x) > 0
	})))
	is.Equal(r2, []lo.Tuple2[int, string]{{1, "foo"}, {3, "bar"}})
}

func TestFilterByKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := maps.Collect(FilterByKeys(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}), []string{"foo", "baz", "qux"}))

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})

	type myIter iter.Seq2[string, int]
	var before myIter = func(yield func(string, int) bool) {
		for k, v := range map[string]int{"": 0, "foobar": 6, "baz": 3} {
			if !yield(k, v) {
				return
			}
		}
	}

	after := FilterByKeys(before, []string{"foobar", "baz"})
	is.IsType(after, before, "type preserved")
}

func TestFilterByValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := maps.Collect(FilterByValues(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}), []int{1, 3}))

	is.Equal(r1, map[string]int{"foo": 1, "baz": 3})

	type myIter iter.Seq2[string, int]
	var before myIter = func(yield func(string, int) bool) {
		for k, v := range map[string]int{"": 0, "foobar": 6, "baz": 3} {
			if !yield(k, v) {
				return
			}
		}
	}

	after := FilterByValues(before, []int{0, 3})
	is.IsType(after, before, "type preserved")
}

func TestReject(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(Reject(slices.Values([]int{1, 2, 3, 4}), func(x int) bool {
		return x%2 == 0
	}))

	is.Equal(r1, []int{1, 3})

	r2 := slices.Collect(Reject(slices.Values([]string{"Smith", "foo", "Domin", "bar", "Olivia"}), func(x string) bool {
		return len(x) > 3
	}))

	is.Equal(r2, []string{"foo", "bar"})

	type myStrings iter.Seq[string]
	var allStrings myStrings = func(yield func(string) bool) {
		for _, str := range []string{"", "foo", "bar"} {
			if !yield(str) {
				return
			}
		}
	}
	nonempty := Reject(allStrings, func(x string) bool {
		return len(x) > 0
	})
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReject2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := maps.Collect(Reject2(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}), func(key string, value int) bool {
		return value%2 == 1
	}))

	is.Equal(r1, map[string]int{"bar": 2})

	type myIter iter.Seq2[string, int]
	var before myIter = func(yield func(string, int) bool) {
		for k, v := range map[string]int{"": 0, "foobar": 6, "baz": 3} {
			if !yield(k, v) {
				return
			}
		}
	}

	after := Reject2(before, func(key string, value int) bool { return true })
	is.IsType(after, before, "type preserved")
}

func TestRejectByKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := maps.Collect(RejectByKeys(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}), []string{"foo", "baz", "qux"}))

	is.Equal(r1, map[string]int{"bar": 2})

	type myIter iter.Seq2[string, int]
	var before myIter = func(yield func(string, int) bool) {
		for k, v := range map[string]int{"": 0, "foobar": 6, "baz": 3} {
			if !yield(k, v) {
				return
			}
		}
	}

	after := RejectByKeys(before, []string{"foobar", "baz"})
	is.IsType(after, before, "type preserved")
}

func TestRejectByValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := maps.Collect(RejectByValues(maps.All(map[string]int{"foo": 1, "bar": 2, "baz": 3}), []int{1, 3}))

	is.Equal(r1, map[string]int{"bar": 2})

	type myIter iter.Seq2[string, int]
	var before myIter = func(yield func(string, int) bool) {
		for k, v := range map[string]int{"": 0, "foobar": 6, "baz": 3} {
			if !yield(k, v) {
				return
			}
		}
	}

	after := RejectByValues(before, []int{0, 3})
	is.IsType(after, before, "type preserved")
}

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(Uniq(slices.Values([]int{1, 2, 2, 1})))

	is.Equal(len(result1), 2)
	is.Equal(result1, []int{1, 2})
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := slices.Collect(UniqBy(slices.Values([]int{0, 1, 2, 3, 4, 5}), func(i int) int {
		return i % 3
	}))

	is.Equal(len(result1), 3)
	is.Equal(result1, []int{0, 1, 2})
}

func TestReplace(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := slices.Values([]int{0, 1, 0, 1, 2, 3, 0})

	out1 := slices.Collect(Replace(in, 0, 42, 2))
	out2 := slices.Collect(Replace(in, 0, 42, 1))
	out3 := slices.Collect(Replace(in, 0, 42, 0))
	out4 := slices.Collect(Replace(in, 0, 42, -1))
	out5 := slices.Collect(Replace(in, 0, 42, -1))
	out6 := slices.Collect(Replace(in, -1, 42, 2))
	out7 := slices.Collect(Replace(in, -1, 42, 1))
	out8 := slices.Collect(Replace(in, -1, 42, 0))
	out9 := slices.Collect(Replace(in, -1, 42, -1))
	out10 := slices.Collect(Replace(in, -1, 42, -1))

	is.Equal([]int{42, 1, 42, 1, 2, 3, 0}, out1)
	is.Equal([]int{42, 1, 0, 1, 2, 3, 0}, out2)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out3)
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out4)
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out5)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out6)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out7)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out8)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out9)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out10)

	type myIter iter.Seq[string]
	var allStrings myIter = func(yield func(string) bool) {
		for _, str := range []string{"", "foo", "bar"} {
			if !yield(str) {
				return
			}
		}
	}

	nonempty := Replace(allStrings, "0", "2", 1)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplace2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := Enumerate(slices.Values([]int{0, 1, 0, 1, 2, 3, 0}))

	out1 := slices.Collect(Tuples(Replace2(in, 0, 42, 2)))
	out2 := slices.Collect(Tuples(Replace2(in, 0, 42, 1)))
	out3 := slices.Collect(Tuples(Replace2(in, 0, 42, 0)))
	out4 := slices.Collect(Tuples(Replace2(in, 0, 42, -1)))
	out5 := slices.Collect(Tuples(Replace2(in, 0, 42, -1)))
	out6 := slices.Collect(Tuples(Replace2(in, -1, 42, 2)))
	out7 := slices.Collect(Tuples(Replace2(in, -1, 42, 1)))
	out8 := slices.Collect(Tuples(Replace2(in, -1, 42, 0)))
	out9 := slices.Collect(Tuples(Replace2(in, -1, 42, -1)))
	out10 := slices.Collect(Tuples(Replace2(in, -1, 42, -1)))

	is.Equal([]lo.Tuple2[int, int]{{0, 42}, {1, 1}, {2, 42}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out1)
	is.Equal([]lo.Tuple2[int, int]{{0, 42}, {1, 1}, {2, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out2)
	is.Equal([]lo.Tuple2[int, int]{{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out3)
	is.Equal([]lo.Tuple2[int, int]{{0, 42}, {1, 1}, {2, 42}, {3, 1}, {4, 2}, {5, 3}, {6, 42}}, out4)
	is.Equal([]lo.Tuple2[int, int]{{0, 42}, {1, 1}, {2, 42}, {3, 1}, {4, 2}, {5, 3}, {6, 42}}, out5)
	is.Equal([]lo.Tuple2[int, int]{{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out6)
	is.Equal([]lo.Tuple2[int, int]{{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out7)
	is.Equal([]lo.Tuple2[int, int]{{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out8)
	is.Equal([]lo.Tuple2[int, int]{{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out9)
	is.Equal([]lo.Tuple2[int, int]{{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out10)

	type myIter iter.Seq2[int, string]
	var allStrings myIter = func(yield func(int, string) bool) {
		for i, str := range []string{"", "foo", "bar"} {
			if !yield(i, str) {
				return
			}
		}
	}

	nonempty := Replace2(allStrings, "0", "2", 1)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := slices.Values([]int{0, 1, 0, 1, 2, 3, 0})

	out1 := slices.Collect(ReplaceAll(in, 0, 42))
	out2 := slices.Collect(ReplaceAll(in, -1, 42))

	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out1)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out2)

	type myIter iter.Seq[string]
	var allStrings myIter = func(yield func(string) bool) {
		for _, str := range []string{"", "foo", "bar"} {
			if !yield(str) {
				return
			}
		}
	}
	nonempty := ReplaceAll(allStrings, "0", "2")
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestReplaceAll2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := Enumerate(slices.Values([]int{0, 1, 0, 1, 2, 3, 0}))

	out1 := slices.Collect(Tuples(ReplaceAll2(in, 0, 42)))
	out2 := slices.Collect(Tuples(ReplaceAll2(in, -1, 42)))

	is.Equal([]lo.Tuple2[int, int]{{0, 42}, {1, 1}, {2, 42}, {3, 1}, {4, 2}, {5, 3}, {6, 42}}, out1)
	is.Equal([]lo.Tuple2[int, int]{{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 2}, {5, 3}, {6, 0}}, out2)

	type myIter iter.Seq2[int, string]
	var allStrings myIter = func(yield func(int, string) bool) {
		for i, str := range []string{"", "foo", "bar"} {
			if !yield(i, str) {
				return
			}
		}
	}

	nonempty := ReplaceAll2(allStrings, "0", "2")
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestCompact(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(Compact(slices.Values([]int{2, 0, 4, 0})))

	is.Equal(r1, []int{2, 4})

	r2 := slices.Collect(Compact(slices.Values([]string{"", "foo", "", "bar", ""})))

	is.Equal(r2, []string{"foo", "bar"})

	r3 := slices.Collect(Compact(slices.Values([]bool{true, false, true, false})))

	is.Equal(r3, []bool{true, true})

	type foo struct {
		bar int
		baz string
	}

	// slice of structs
	// If all fields of an element are zero values, Compact removes it.

	r4 := slices.Collect(Compact(slices.Values([]foo{
		{bar: 1, baz: "a"}, // all fields are non-zero values
		{bar: 0, baz: ""},  // all fields are zero values
		{bar: 2, baz: ""},  // bar is non-zero
	})))

	is.Equal(r4, []foo{{bar: 1, baz: "a"}, {bar: 2, baz: ""}})

	// slice of pointers to structs
	// If an element is nil, Compact removes it.

	e1, e2, e3 := foo{bar: 1, baz: "a"}, foo{bar: 0, baz: ""}, foo{bar: 2, baz: ""}
	// NOTE: e2 is a zero value of foo, but its pointer &e2 is not a zero value of *foo.
	r5 := slices.Collect(Compact(slices.Values([]*foo{&e1, &e2, nil, &e3})))

	is.Equal(r5, []*foo{&e1, &e2, &e3})

	type myIter iter.Seq[string]
	var allStrings myIter = func(yield func(string) bool) {
		for _, str := range []string{"", "foo", "bar"} {
			if !yield(str) {
				return
			}
		}
	}

	nonempty := Compact(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}

func TestCompact2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := slices.Collect(Tuples(Compact2(Enumerate(slices.Values([]int{2, 0, 4, 0})))))

	is.Equal(r1, []lo.Tuple2[int, int]{{0, 2}, {2, 4}})

	r2 := slices.Collect(Tuples(Compact2(Enumerate(slices.Values([]string{"", "foo", "", "bar", ""})))))

	is.Equal(r2, []lo.Tuple2[int, string]{{1, "foo"}, {3, "bar"}})

	r3 := slices.Collect(Tuples(Compact2(Enumerate(slices.Values([]bool{true, false, true, false})))))

	is.Equal(r3, []lo.Tuple2[int, bool]{{0, true}, {2, true}})

	type foo struct {
		bar int
		baz string
	}

	// slice of structs
	// If all fields of an element are zero values, Compact removes it.

	r4 := slices.Collect(Tuples(Compact2(Enumerate(slices.Values([]foo{
		{bar: 1, baz: "a"}, // all fields are non-zero values
		{bar: 0, baz: ""},  // all fields are zero values
		{bar: 2, baz: ""},  // bar is non-zero
	})))))

	is.Equal(r4, []lo.Tuple2[int, foo]{{0, foo{bar: 1, baz: "a"}}, {2, foo{bar: 2, baz: ""}}})

	// slice of pointers to structs
	// If an element is nil, Compact removes it.

	e1, e2, e3 := foo{bar: 1, baz: "a"}, foo{bar: 0, baz: ""}, foo{bar: 2, baz: ""}
	// NOTE: e2 is a zero value of foo, but its pointer &e2 is not a zero value of *foo.
	r5 := slices.Collect(Tuples(Compact2(Enumerate(slices.Values([]*foo{&e1, &e2, nil, &e3})))))

	is.Equal(r5, []lo.Tuple2[int, *foo]{{0, &e1}, {1, &e2}, {3, &e3}})

	type myIter iter.Seq2[int, string]
	var allStrings myIter = func(yield func(int, string) bool) {
		for i, str := range []string{"", "foo", "bar"} {
			if !yield(i, str) {
				return
			}
		}
	}

	nonempty := Compact2(allStrings)
	is.IsType(nonempty, allStrings, "type preserved")
}
