//go:build go1.23

package it

import (
	"iter"
	"slices"
	"testing"
)

func intersectionOneShot(items ...int) iter.Seq[int] {
	ch := make(chan int, len(items))
	for _, item := range items {
		ch <- item
	}
	close(ch)
	return ChannelToSeq(ch)
}

func TestIntersect_singleUse(t *testing.T) {
	for _, test := range []struct {
		name            string
		input, expected []int
	}{
		{name: "empty"},
		{name: "single", input: []int{4}, expected: []int{4}},
		{name: "duplicates", input: []int{3, 1, 3, 2, 1}, expected: []int{3, 1, 2}},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := slices.Collect(Intersect(intersectionOneShot(test.input...)))
			if !slices.Equal(got, test.expected) {
				t.Fatalf("got %v, want %v", got, test.expected)
			}
		})
	}
}

func TestIntersectBy_singleUse(t *testing.T) {
	for _, test := range []struct {
		name            string
		input, expected []int
	}{
		{name: "empty"},
		{name: "single", input: []int{4}, expected: []int{4}},
		{name: "duplicate keys", input: []int{13, 11, 23, 12, 21}, expected: []int{13, 11, 12}},
	} {
		t.Run(test.name, func(t *testing.T) {
			calls := 0
			got := slices.Collect(IntersectBy(func(item int) int {
				calls++
				return item % 10
			}, intersectionOneShot(test.input...)))
			if !slices.Equal(got, test.expected) || calls != len(test.input) {
				t.Fatalf("got %v with %d transforms, want %v with %d", got, calls, test.expected, len(test.input))
			}
		})
	}
}

func TestIntersect_singleUseLazyBreak(t *testing.T) {
	for _, keyed := range []bool{false, true} {
		name := "plain"
		if keyed {
			name = "keyed"
		}
		t.Run(name, func(t *testing.T) {
			reads := 0
			input := func(yield func(int) bool) {
				for _, item := range []int{7, 4, 7} {
					reads++
					if !yield(item) {
						return
					}
				}
			}
			result := Intersect(input)
			if keyed {
				result = IntersectBy(func(item int) int { return item % 3 }, input)
			}
			if reads != 0 {
				t.Fatal("input consumed before iteration")
			}
			for item := range result {
				if item != 7 {
					t.Fatalf("first item = %d, want 7", item)
				}
				break
			}
			if reads != 1 {
				t.Fatalf("early break consumed %d items, want 1", reads)
			}
		})
	}
}

func TestIntersect_singleUseControls(t *testing.T) {
	t.Run("named and restartable", func(t *testing.T) {
		type namedSequence iter.Seq[int]
		input := namedSequence(slices.Values([]int{3, 1, 3}))
		var plain namedSequence = Intersect(input)
		var keyed namedSequence = IntersectBy(func(item int) int { return item }, input)
		for _, seq := range []namedSequence{plain, keyed} {
			for repeat := 0; repeat < 2; repeat++ {
				if got := slices.Collect(iter.Seq[int](seq)); !slices.Equal(got, []int{3, 1}) {
					t.Fatalf("repeat %d: got %v", repeat, got)
				}
			}
		}
	})
	t.Run("two independent one-shot inputs", func(t *testing.T) {
		got := slices.Collect(Intersect(intersectionOneShot(3, 1, 3, 2), intersectionOneShot(2, 3, 2)))
		if !slices.Equal(got, []int{3, 2}) {
			t.Fatalf("got %v", got)
		}
		got = slices.Collect(IntersectBy(func(item int) int { return item % 10 },
			intersectionOneShot(13, 11, 23, 12), intersectionOneShot(22, 3, 32)))
		if !slices.Equal(got, []int{13, 12}) {
			t.Fatalf("got %v", got)
		}
	})
	t.Run("zero inputs", func(t *testing.T) {
		if got := slices.Collect(Intersect[int, iter.Seq[int]]()); len(got) != 0 {
			t.Fatalf("got %v", got)
		}
		if got := slices.Collect(IntersectBy[int, int, iter.Seq[int]](func(item int) int { return item })); len(got) != 0 {
			t.Fatalf("got %v", got)
		}
	})
}
