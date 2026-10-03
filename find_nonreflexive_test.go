package lo_test

import (
	"math"
	"testing"

	"github.com/samber/lo"
)

func TestFindUniques_nonReflexiveValues(t *testing.T) {
	t.Parallel()
	type floats []float64
	for _, size := range []int{1, 2, 8, 9} {
		input := make(floats, size)
		for i := range input {
			input[i] = math.NaN()
		}
		got := lo.FindUniques(input)
		if len(got) != size {
			t.Errorf("size %d: got %d values, want %d", size, len(got), size)
		}
		for i, value := range got {
			if !math.IsNaN(value) {
				t.Errorf("size %d index %d: got %v, want NaN", size, i, value)
			}
		}
	}
	input := []float64{2, math.NaN(), 2, 1, math.NaN(), 3}
	got := lo.FindUniques(input)
	if len(got) != 4 {
		t.Fatalf("mixed values: got %v, want [NaN 1 NaN 3]", got)
	}
	if !math.IsNaN(got[0]) || got[1] != 1 || !math.IsNaN(got[2]) || got[3] != 3 {
		t.Errorf("unexpected order: %v", got)
	}
}

func TestFindUniques_nonReflexiveCompositeValues(t *testing.T) {
	t.Parallel()
	type key struct{ value complex128 }
	structured := []key{{complex(math.NaN(), 1)}, {complex(math.NaN(), 1)}}
	if got := lo.FindUniques(structured); len(got) != len(structured) {
		t.Errorf("non-reflexive struct keys: got %v", got)
	}
	if got := lo.FindUniques([]float64{0, math.Copysign(0, -1), 1}); len(got) != 1 || got[0] != 1 {
		t.Errorf("signed zeros must still compare equal: %v", got)
	}
	if got := lo.FindUniques([]float64(nil)); got == nil || len(got) != 0 {
		t.Errorf("empty output: %v", got)
	}
}

func TestFindUniquesBy_nonReflexiveKeys(t *testing.T) {
	t.Parallel()
	type entries []int
	for _, size := range []int{1, 2, 8, 9} {
		input := make(entries, size)
		for i := range input {
			input[i] = i
		}
		got := lo.FindUniquesBy(input, func(int) float64 { return math.NaN() })
		if len(got) != size {
			t.Errorf("size %d: got %v, want %v", size, got, input)
		}
		for i, value := range got {
			if value != input[i] {
				t.Errorf("size %d index %d: got %d, want %d", size, i, value, input[i])
			}
		}
	}
	input := entries{0, 1, 2, 3, 4, 5}
	calls := 0
	got := lo.FindUniquesBy(input, func(value int) float64 {
		calls++
		if value == 1 || value == 4 {
			return math.NaN()
		}
		if value == 0 || value == 2 {
			return 2
		}
		return float64(value)
	})
	want := entries{1, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("mixed keys: got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("mixed keys: got %v, want %v", got, want)
		}
	}
	if calls != len(input) {
		t.Errorf("small path called iteratee %d times, want %d", calls, len(input))
	}
}
