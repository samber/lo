package lo

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		item     int
		expected bool
	}{
		{name: "item present", item: 5, expected: true},
		{name: "item absent", item: 6, expected: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Contains([]int{0, 1, 2, 3, 4, 5}, tt.item))
		})
	}
}

func TestContainsBy(t *testing.T) {
	t.Parallel()

	type a struct {
		A int
		B string
	}

	t.Run("struct predicate: partial match on A only", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a1 := []a{{A: 1, B: "1"}, {A: 2, B: "2"}, {A: 3, B: "3"}}
		is.False(ContainsBy(a1, func(t a) bool { return t.A == 1 && t.B == "2" }))
	})

	t.Run("struct predicate: exact match", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a1 := []a{{A: 1, B: "1"}, {A: 2, B: "2"}, {A: 3, B: "3"}}
		is.True(ContainsBy(a1, func(t a) bool { return t.A == 2 && t.B == "2" }))
	})

	t.Run("string predicate: found", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a2 := []string{"aaa", "bbb", "ccc"}
		is.True(ContainsBy(a2, func(t string) bool { return t == "ccc" }))
	})

	t.Run("string predicate: not found", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		a2 := []string{"aaa", "bbb", "ccc"}
		is.False(ContainsBy(a2, func(t string) bool { return t == "ddd" }))
	})
}

// TestEvery_smallScan exercises the small-scan path (all subsets here are
// <= everySmallSubset). See TestEvery_large for the map-based path.
func TestEvery_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		subset   []int
		expected bool
	}{
		{name: "subset fully present", subset: []int{0, 2}, expected: true},
		{name: "one element missing", subset: []int{0, 6}, expected: false},
		{name: "all elements missing", subset: []int{-1, 6}, expected: false},
		{name: "empty subset", subset: []int{}, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Every([]int{0, 1, 2, 3, 4, 5}, tt.subset))
		})
	}
}

// Every dispatches on len(subset) <= everySmallSubset (8): a subset of 9
// elements forces the everyLarge path, which the table above never
// exercises (its subsets are all <= 2 elements).
func TestEvery_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	collection := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	subsetPresent := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	is.Greater(len(subsetPresent), everySmallSubset, "sanity check: subset must exceed everySmallSubset")

	tests := []struct {
		name     string
		subset   []int
		expected bool
	}{
		{name: "large subset present", subset: subsetPresent, expected: true},
		{name: "large subset missing element", subset: []int{0, 1, 2, 3, 4, 5, 6, 7, 10}, expected: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Every(collection, tt.subset))
		})
	}
}

func TestEveryBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		predicate  func(int) bool
		expected   bool
	}{
		{name: "all match", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 5 }, expected: true},
		{name: "some match", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 3 }, expected: false},
		{name: "none match", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 0 }, expected: false},
		{name: "empty collection", collection: []int{}, predicate: func(x int) bool { return x < 5 }, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, EveryBy(tt.collection, tt.predicate))
		})
	}
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		subset   []int
		expected bool
	}{
		{name: "some elements present", subset: []int{0, 2}, expected: true},
		{name: "one element present", subset: []int{0, 6}, expected: true},
		{name: "no elements present", subset: []int{-1, 6}, expected: false},
		{name: "empty subset", subset: []int{}, expected: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Some([]int{0, 1, 2, 3, 4, 5}, tt.subset))
		})
	}
}

func TestSomeBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		predicate  func(int) bool
		expected   bool
	}{
		{name: "some match", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 5 }, expected: true},
		{name: "one matches", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 3 }, expected: true},
		{name: "none match", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 0 }, expected: false},
		{name: "empty collection", collection: []int{}, predicate: func(x int) bool { return x < 5 }, expected: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, SomeBy(tt.collection, tt.predicate))
		})
	}
}

func TestNone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		subset   []int
		expected bool
	}{
		{name: "some elements present", subset: []int{0, 2}, expected: false},
		{name: "one element present", subset: []int{0, 6}, expected: false},
		{name: "no elements present", subset: []int{-1, 6}, expected: true},
		{name: "empty subset", subset: []int{}, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, None([]int{0, 1, 2, 3, 4, 5}, tt.subset))
		})
	}
}

func TestNoneBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name       string
		collection []int
		predicate  func(int) bool
		expected   bool
	}{
		{name: "all match (none is false)", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 5 }, expected: false},
		{name: "some match", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 3 }, expected: false},
		{name: "none match", collection: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x < 0 }, expected: true},
		{name: "empty collection", collection: []int{}, predicate: func(x int) bool { return x < 5 }, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, NoneBy(tt.collection, tt.predicate))
		})
	}
}

func TestIntersect(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{name: "no lists", lists: [][]int{}, expected: []int{}},
		{name: "single list", lists: [][]int{{1}}, expected: []int{1}},
		{name: "two lists overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 2}}, expected: []int{0, 2}},
		{name: "two lists partial overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 6}}, expected: []int{0}},
		{name: "two lists no overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {-1, 6}}, expected: []int{}},
		{name: "reversed order still overlaps", lists: [][]int{{0, 6}, {0, 1, 2, 3, 4, 5}}, expected: []int{0}},
		{name: "duplicate elements in first list", lists: [][]int{{0, 6, 0}, {0, 1, 2, 3, 4, 5}}, expected: []int{0}},
		{name: "three lists overlap", lists: [][]int{{0, 6, 0, 3}, {0, 1, 2, 3, 4, 5}, {0, 6}}, expected: []int{0}},
		{name: "three lists no common element", lists: [][]int{{0, 6, 0, 3}, {0, 1, 2, 3, 4, 5}, {1, 6}}, expected: []int{}},
		{name: "three lists disjoint", lists: [][]int{{0, 1, 1}, {2}, {3}}, expected: []int{}},
		{name: "single list with duplicates", lists: [][]int{{0, 1, 1}}, expected: []int{0, 1}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.ElementsMatch(tt.expected, Intersect(tt.lists...))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := Intersect(allStrings, allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})

	t.Run("type preserved: map path (single list)", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		// single-list call: len(lists) != 2, so it takes the map path too.
		nonemptyLarge := Intersect(allStrings)
		is.IsType(nonemptyLarge, allStrings, "type preserved (map path)")
	})
}

func TestIntersectBy(t *testing.T) {
	t.Parallel()

	type User struct {
		ID   int
		Name string
	}

	list1 := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	list2 := []User{
		{ID: 2, Name: "Robert"},
		{ID: 3, Name: "Charlie"},
		{ID: 4, Name: "Alice"},
	}

	t.Run("by ID", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		intersectByID := IntersectBy(func(u User) int {
			return u.ID
		}, list1, list2)
		is.ElementsMatch(intersectByID, []User{{ID: 2, Name: "Bob"}, {ID: 3, Name: "Charlie"}})
	})

	t.Run("by Name", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		intersectByName := IntersectBy(func(u User) string {
			return u.Name
		}, list1, list2)
		is.ElementsMatch(intersectByName, []User{{ID: 3, Name: "Charlie"}, {ID: 1, Name: "Alice"}})
	})

	t.Run("by ID and Name combined", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		intersectByIDAndName := IntersectBy(func(u User) string {
			return strconv.Itoa(u.ID) + u.Name
		}, list1, list2)
		is.ElementsMatch(intersectByIDAndName, []User{{ID: 3, Name: "Charlie"}})
	})

	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{name: "three lists via string conversion", lists: [][]int{{0, 6, 0, 3}, {0, 1, 2, 3, 4, 5}, {0, 6}}, expected: []int{0}},
		{name: "single list with duplicates via string conversion", lists: [][]int{{0, 1, 1}}, expected: []int{0, 1}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ElementsMatch(t, tt.expected, IntersectBy(strconv.Itoa, tt.lists...))
		})
	}
}

// TestDifference_smallScan exercises the small-scan path (all lists here are
// <= differenceSmallThreshold). See TestDifference_large for the map-based
// path.
func TestDifference_smallScan(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name      string
		list1     []int
		list2     []int
		wantLeft  []int
		wantRight []int
	}{
		{name: "some elements differ", list1: []int{0, 1, 2, 3, 4, 5}, list2: []int{0, 2, 6}, wantLeft: []int{1, 3, 4, 5}, wantRight: []int{6}},
		{name: "no overlap", list1: []int{1, 2, 3, 4, 5}, list2: []int{0, 6}, wantLeft: []int{1, 2, 3, 4, 5}, wantRight: []int{0, 6}},
		{name: "identical lists", list1: []int{0, 1, 2, 3, 4, 5}, list2: []int{0, 1, 2, 3, 4, 5}, wantLeft: []int{}, wantRight: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			left, right := Difference(tt.list1, tt.list2)
			is.Equal(tt.wantLeft, left)
			is.Equal(tt.wantRight, right)
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		a, b := Difference(allStrings, allStrings)
		is.IsType(a, allStrings, "type preserved")
		is.IsType(b, allStrings, "type preserved")
	})
}

// Difference dispatches on len(list1) <= differenceSmallThreshold &&
// len(list2) <= differenceSmallThreshold (8): a pair of 9-element lists
// forces the differenceLarge path, which the table above never
// exercises (its lists are all <= 6 elements).
func TestDifference_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	list1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	list2 := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
	is.Greater(len(list1), differenceSmallThreshold, "sanity check: list1 must exceed differenceSmallThreshold")
	is.Greater(len(list2), differenceSmallThreshold, "sanity check: list2 must exceed differenceSmallThreshold")

	same := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	tests := []struct {
		name      string
		list1     []int
		list2     []int
		wantLeft  []int
		wantRight []int
	}{
		{name: "partial overlap", list1: list1, list2: list2, wantLeft: []int{0, 1, 2, 3}, wantRight: []int{9, 10, 11, 12}},
		{name: "identical lists", list1: same, list2: same, wantLeft: []int{}, wantRight: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			left, right := Difference(tt.list1, tt.list2)
			is.Equal(tt.wantLeft, left)
			is.Equal(tt.wantRight, right)
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
		is.Greater(len(allStrings), differenceSmallThreshold, "sanity check: allStrings must exceed differenceSmallThreshold")
		a, b := Difference(allStrings, allStrings)
		is.Empty(a)
		is.Empty(b)
		is.IsType(a, allStrings, "type preserved")
		is.IsType(b, allStrings, "type preserved")
	})
}

func TestUnion(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{name: "two lists with overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 2, 10}}, expected: []int{0, 1, 2, 3, 4, 5, 10}},
		{name: "two disjoint lists", lists: [][]int{{0, 1, 2, 3, 4, 5}, {6, 7}}, expected: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{name: "second list empty", lists: [][]int{{0, 1, 2, 3, 4, 5}, {}}, expected: []int{0, 1, 2, 3, 4, 5}},
		{name: "duplicate elements deduped", lists: [][]int{{0, 1, 2}, {0, 1, 2, 3, 3}}, expected: []int{0, 1, 2, 3}},
		{name: "identical lists", lists: [][]int{{0, 1, 2}, {0, 1, 2}}, expected: []int{0, 1, 2}},
		{name: "both lists empty", lists: [][]int{{}, {}}, expected: []int{}},
		{name: "three lists with overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 2, 10}, {0, 1, 11}}, expected: []int{0, 1, 2, 3, 4, 5, 10, 11}},
		{name: "three disjoint lists", lists: [][]int{{0, 1, 2, 3, 4, 5}, {6, 7}, {8, 9}}, expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "three lists two empty", lists: [][]int{{0, 1, 2, 3, 4, 5}, {}, {}}, expected: []int{0, 1, 2, 3, 4, 5}},
		{name: "three identical lists", lists: [][]int{{0, 1, 2}, {0, 1, 2}, {0, 1, 2}}, expected: []int{0, 1, 2}},
		{name: "three empty lists", lists: [][]int{{}, {}, {}}, expected: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Union(tt.lists...))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := Union(allStrings, allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})

	t.Run("type preserved: map path (large input)", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		// total element count > unionSmallThreshold, so this takes the map path.
		largeStrings := myStrings{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
		nonemptyLarge := Union(largeStrings, largeStrings)
		is.Equal(largeStrings, nonemptyLarge)
		is.IsType(nonemptyLarge, largeStrings, "type preserved (map path)")
	})
}

func TestUnionBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testFunc := func(i int) int {
		return i / 2
	}

	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{name: "two lists with overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 2, 10}}, expected: []int{0, 2, 4, 10}},
		{name: "two disjoint lists", lists: [][]int{{0, 1, 2, 3, 4, 5}, {6, 7}}, expected: []int{0, 2, 4, 6}},
		{name: "second list empty", lists: [][]int{{0, 1, 2, 3, 4, 5}, {}}, expected: []int{0, 2, 4}},
		{name: "identical lists", lists: [][]int{{0, 1, 2}, {0, 1, 2}}, expected: []int{0, 2}},
		{name: "both lists empty", lists: [][]int{{}, {}}, expected: []int{}},
		{name: "three lists with overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 2, 10}, {0, 1, 11}}, expected: []int{0, 2, 4, 10}},
		{name: "three disjoint lists", lists: [][]int{{0, 1, 2, 3, 4, 5}, {6, 7}, {8, 9}}, expected: []int{0, 2, 4, 6, 8}},
		{name: "three lists two empty", lists: [][]int{{0, 1, 2, 3, 4, 5}, {}, {}}, expected: []int{0, 2, 4}},
		{name: "three identical lists", lists: [][]int{{0, 1, 2}, {0, 1, 2}, {0, 1, 2}}, expected: []int{0, 2}},
		{name: "three empty lists", lists: [][]int{{}, {}, {}}, expected: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, UnionBy(testFunc, tt.lists...))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"foo", "bar", "baz"}
		nonempty := UnionBy(func(s string) string { return s }, allStrings, allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestUnionByErr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	testFunc := func(i int) (int, error) {
		return i / 2, nil
	}

	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{name: "two lists with overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 2, 10}}, expected: []int{0, 2, 4, 10}},
		{name: "two disjoint lists", lists: [][]int{{0, 1, 2, 3, 4, 5}, {6, 7}}, expected: []int{0, 2, 4, 6}},
		{name: "second list empty", lists: [][]int{{0, 1, 2, 3, 4, 5}, {}}, expected: []int{0, 2, 4}},
		{name: "identical lists", lists: [][]int{{0, 1, 2}, {0, 1, 2}}, expected: []int{0, 2}},
		{name: "both lists empty", lists: [][]int{{}, {}}, expected: []int{}},
		{name: "three lists with overlap", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 2, 10}, {0, 1, 11}}, expected: []int{0, 2, 4, 10}},
		{name: "three disjoint lists", lists: [][]int{{0, 1, 2, 3, 4, 5}, {6, 7}, {8, 9}}, expected: []int{0, 2, 4, 6, 8}},
		{name: "three lists two empty", lists: [][]int{{0, 1, 2, 3, 4, 5}, {}, {}}, expected: []int{0, 2, 4}},
		{name: "three identical lists", lists: [][]int{{0, 1, 2}, {0, 1, 2}, {0, 1, 2}}, expected: []int{0, 2}},
		{name: "three empty lists", lists: [][]int{{}, {}, {}}, expected: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := UnionByErr(testFunc, tt.lists...)
			is.NoError(err)
			is.Equal(tt.expected, result)
		})
	}

	// Test error case
	errorTests := []struct {
		name          string
		lists         [][]int
		wantCallCount int
	}{
		{name: "error in first list", lists: [][]int{{0, 1, 2, 3, 4, 5}, {0, 2, 10}}, wantCallCount: 3},
		{name: "error in second list", lists: [][]int{{0, 1, 3, 4, 5}, {2, 10}}, wantCallCount: 6},
	}

	for _, tt := range errorTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callCount := 0
			errFunc := func(i int) (int, error) {
				callCount++
				if i == 2 {
					return 0, assert.AnError
				}
				return i / 2, nil
			}

			result, err := UnionByErr(errFunc, tt.lists...)
			is.ErrorIs(err, assert.AnError)
			is.Nil(result)
			is.Equal(tt.wantCallCount, callCount, "should stop at first error")
		})
	}
}

// TestWithout_small exercises the small-scan path (all exclude lists here are
// <= withoutSmallExcludeThreshold). See TestWithout_large for the map-based path.
func TestWithout_small(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.LessOrEqual(3, withoutSmallExcludeThreshold, "sanity check: exclude must not exceed withoutSmallExcludeThreshold")

	tests := []struct {
		name     string
		input    []int
		exclude  []int
		expected []int
	}{
		{name: "exclude everything", input: []int{0, 1, 2}, exclude: []int{0, 1, 2}, expected: []int{}},
		{name: "empty input", input: []int{}, exclude: []int{}, expected: []int{}},
		{name: "exclude one element", input: []int{0, 1, 2, 3}, exclude: []int{1}, expected: []int{0, 2, 3}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Without(tt.input, tt.exclude...))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := Without(allStrings, "")
		is.Equal(myStrings{"foo", "bar"}, nonempty)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// TestWithout_large exercises the map-based path (all exclude lists here
// exceed withoutSmallExcludeThreshold). See TestWithout_small for the linear-scan path.
func TestWithout_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	exclude := []int{0, 1, 2, 3, 4, 5}
	is.Greater(len(exclude), withoutSmallExcludeThreshold, "sanity check: exclude must exceed withoutSmallExcludeThreshold")

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "some elements remain", input: []int{0, 2, 10}, expected: []int{10}},
		{name: "one element remains", input: []int{0, 7}, expected: []int{7}},
		{name: "empty input", input: []int{}, expected: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, Without(tt.input, exclude...))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := Without(allStrings, "a", "b", "c", "d", "e")
		is.Equal(allStrings, nonempty)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// TestWithoutBy_small exercises the small-scan path (all exclude lists
// here are <= withoutSmallExcludeThreshold). See TestWithoutBy_large for
// the map-based path.
func TestWithoutBy_small(t *testing.T) {
	t.Parallel()

	type User struct {
		Name string
		Age  int
	}

	t.Run("exclude by name", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := WithoutBy([]User{{Name: "nick"}, {Name: "peter"}},
			func(item User) string {
				return item.Name
			}, "nick", "lily")
		is.Equal([]User{{Name: "peter"}}, result)
	})

	t.Run("empty collection with int key", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := WithoutBy([]User{}, func(item User) int { return item.Age }, 1, 2, 3)
		is.Empty(result)
	})

	t.Run("empty collection, no exclude keys", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := WithoutBy([]User{}, func(item User) string { return item.Name })
		is.Empty(result)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := WithoutBy(allStrings, func(s string) string {
			return s
		})
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

// WithoutBy dispatches on len(exclude) <= withoutSmallExcludeThreshold (4): an
// exclude list of 5 keys forces the withoutByLarge path, which the table
// above never exercises (its exclude lists are all <= 3 elements).
func TestWithoutBy_large(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	byKey := func(v int) int { return v }
	collection := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	tests := []struct {
		name     string
		exclude  []int
		expected []int
	}{
		{name: "excludes some", exclude: []int{0, 1, 2, 3, 4}, expected: []int{5, 6, 7, 8, 9}},
		{name: "excludes all", exclude: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		is.Greater(len(tt.exclude), withoutSmallExcludeThreshold, "sanity check: exclude must exceed withoutSmallExcludeThreshold")
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, WithoutBy(collection, byKey, tt.exclude...))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
		nonempty := WithoutBy(allStrings, func(s string) string { return s }, "z", "y", "x", "w", "v")
		is.Equal(allStrings, nonempty)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestWithoutByErr(t *testing.T) {
	t.Parallel()

	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name          string
		input         []User
		iteratee      func(User) (string, error)
		exclude       []string
		want          []User
		wantErr       string
		wantCallCount int
	}{
		{
			name:  "exclude by name",
			input: []User{{Name: "nick"}, {Name: "peter"}},
			iteratee: func(item User) (string, error) {
				return item.Name, nil
			},
			exclude:       []string{"nick", "lily"},
			want:          []User{{Name: "peter"}},
			wantErr:       "",
			wantCallCount: 2,
		},
		{
			name:  "empty exclude list",
			input: []User{{Name: "nick"}, {Name: "peter"}},
			iteratee: func(item User) (string, error) {
				return item.Name, nil
			},
			exclude:       []string{},
			want:          []User{{Name: "nick"}, {Name: "peter"}},
			wantErr:       "",
			wantCallCount: 2,
		},
		{
			name:  "error on second element",
			input: []User{{Name: "nick"}, {Name: "peter"}, {Name: "lily"}},
			iteratee: func(item User) (string, error) {
				if item.Name == "peter" {
					return "", errors.New("peter not allowed")
				}
				return item.Name, nil
			},
			exclude:       []string{"nick"},
			want:          nil,
			wantErr:       "peter not allowed",
			wantCallCount: 2, // stops early at error
		},
		{
			name:  "error on first element",
			input: []User{{Name: "nick"}, {Name: "peter"}},
			iteratee: func(item User) (string, error) {
				return "", errors.New("first element error")
			},
			exclude:       []string{"nick"},
			want:          nil,
			wantErr:       "first element error",
			wantCallCount: 1,
		},
		{
			name:  "all excluded",
			input: []User{{Name: "nick"}, {Name: "peter"}},
			iteratee: func(item User) (string, error) {
				return item.Name, nil
			},
			exclude:       []string{"nick", "peter", "lily"},
			want:          []User{},
			wantErr:       "",
			wantCallCount: 2,
		},
		{
			name:  "none excluded",
			input: []User{{Name: "nick"}, {Name: "peter"}},
			iteratee: func(item User) (string, error) {
				return item.Name, nil
			},
			exclude:       []string{"alice"},
			want:          []User{{Name: "nick"}, {Name: "peter"}},
			wantErr:       "",
			wantCallCount: 2,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			callCount := 0
			wrappedIteratee := func(item User) (string, error) {
				callCount++
				return tt.iteratee(item)
			}

			got, err := WithoutByErr(tt.input, wrappedIteratee, tt.exclude...)

			if tt.wantErr != "" {
				is.Error(err)
				is.Equal(tt.wantErr, err.Error())
				is.Nil(got)
				if tt.wantCallCount > 0 {
					is.Equal(tt.wantCallCount, callCount, "should stop early on error")
				}
			} else {
				is.NoError(err)
				is.Equal(tt.want, got)
				is.Equal(tt.wantCallCount, callCount)
			}
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty, err := WithoutByErr(allStrings, func(s string) (string, error) {
			return s, nil
		})

		is.NoError(err)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestWithoutEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "removes leading zero", input: []int{0, 1, 2}, expected: []int{1, 2}},
		{name: "no zero to remove", input: []int{1, 2}, expected: []int{1, 2}},
		{name: "empty input", input: []int{}, expected: []int{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, WithoutEmpty(tt.input))
		})
	}

	t.Run("removes nil pointers", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		result := WithoutEmpty([]*int{ToPtr(0), ToPtr(1), nil, ToPtr(2)})
		is.Equal([]*int{ToPtr(0), ToPtr(1), ToPtr(2)}, result)
	})

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := WithoutEmpty(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})
}

func TestWithoutNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    []int
		nths     []int
		expected []int
	}{
		{name: "removes indices in reverse order", input: []int{5, 6, 7}, nths: []int{1, 0}, expected: []int{7}},
		{name: "no indices to remove", input: []int{1, 2}, nths: nil, expected: []int{1, 2}},
		{name: "empty input", input: []int{}, nths: nil, expected: []int{}},
		{name: "negative and out-of-range indices are ignored", input: []int{0, 1, 2, 3}, nths: []int{-1, 4}, expected: []int{0, 1, 2, 3}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, WithoutNth(tt.input, tt.nths...))
		})
	}

	t.Run("type preserved", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		type myStrings []string
		allStrings := myStrings{"", "foo", "bar"}
		nonempty := WithoutNth(allStrings)
		is.IsType(nonempty, allStrings, "type preserved")
	})

	t.Run("supports non-comparable element types", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		// This works for non-comparable as well
		result := WithoutNth([]func() int{func() int { return 1 }, func() int { return 2 }, func() int { return 3 }}, 1)
		is.Equal([]int{1, 3}, Map(result, func(f func() int, _ int) int { return f() }))
	})
}

func TestElementsMatch(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected bool
	}{
		{name: "empty vs non-empty", list1: []int{}, list2: []int{1}, expected: false},
		{name: "different single elements", list1: []int{1}, list2: []int{2}, expected: false},
		{name: "different lengths", list1: []int{1}, list2: []int{1, 2}, expected: false},
		{name: "different element counts", list1: []int{1, 1, 2}, list2: []int{2, 2, 1}, expected: false},
		{name: "empty vs nil", list1: []int{}, list2: nil, expected: true},
		{name: "single matching element", list1: []int{1}, list2: []int{1}, expected: true},
		{name: "duplicate matching elements", list1: []int{1, 1}, list2: []int{1, 1}, expected: true},
		{name: "same elements different order", list1: []int{1, 2}, list2: []int{2, 1}, expected: true},
		{name: "same multiset different order", list1: []int{1, 1, 2}, list2: []int{1, 2, 1}, expected: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			is.Equal(tt.expected, ElementsMatch(tt.list1, tt.list2))
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
		[]someType{{key: "a"}, {key: "b"}},
		[]someType{{key: "b"}, {key: "a"}},
		func(item someType) string { return item.key },
	))
}
