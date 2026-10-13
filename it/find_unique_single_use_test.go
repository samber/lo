//go:build go1.23

package it_test

import (
	"iter"
	"reflect"
	"slices"
	"testing"

	"github.com/samber/lo/it"
)

func TestFindUniques_singleUse(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []int
		want  []int
	}{
		{"mixed", []int{4, 2, 4, 1, 3, 2, 5}, []int{1, 3, 5}},
		{"all_unique", []int{3, 1, 2}, []int{3, 1, 2}},
		{"all_duplicates", []int{2, 1, 2, 1}, nil},
		{"empty", nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ch := make(chan int, len(tc.input))
			for _, value := range tc.input {
				ch <- value
			}
			close(ch)
			if got := slices.Collect(it.FindUniques(it.ChannelToSeq(ch))); !slices.Equal(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindUniquesBy_singleUse(t *testing.T) {
	type item struct {
		key   int
		value []string
	}
	input := []item{{1, []string{"duplicate"}}, {2, []string{"first"}}, {1, []string{"duplicate again"}}, {3, []string{"second"}}}
	ch := make(chan item, len(input))
	for _, value := range input {
		ch <- value
	}
	close(ch)
	calls := 0
	got := slices.Collect(it.FindUniquesBy(it.ChannelToSeq(ch), func(value item) int {
		calls++
		return value.key
	}))
	want := []item{input[1], input[3]}
	if !reflect.DeepEqual(got, want) || calls != len(input) {
		t.Fatalf("got %v and %d transform calls, want %v and %d", got, calls, want, len(input))
	}
}

func TestFindUniquesBy_singleTraversal(t *testing.T) {
	type namedSeq iter.Seq[int]
	traversals, calls := 0, 0
	input := namedSeq(func(yield func(int) bool) {
		traversals++
		for _, value := range []int{4, 1, 4, 2, 3} {
			if !yield(value) {
				return
			}
		}
	})
	result := it.FindUniquesBy(input, func(value int) int { calls++; return value })
	if traversals != 0 || calls != 0 {
		t.Fatal("input consumed before iteration")
	}
	var sameType namedSeq = result
	for value := range sameType {
		if value != 1 {
			t.Fatalf("first unique = %d, want 1", value)
		}
		break
	}
	if traversals != 1 || calls != 5 {
		t.Fatalf("got %d traversals and %d transforms, want 1 and 5", traversals, calls)
	}
	if got := slices.Collect(iter.Seq[int](result)); !slices.Equal(got, []int{1, 2, 3}) || traversals != 2 || calls != 10 {
		t.Fatalf("restartable input: got %v, %d traversals and %d transforms", got, traversals, calls)
	}
}
