//go:build go1.23

package it_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/samber/lo/it"
)

func assertRangeBounds[T comparable](t *testing.T, seq iter.Seq[T], want []T) {
	t.Helper()
	var got []T
	for item := range seq {
		got = append(got, item)
		if len(got) > len(want) {
			break
		}
	}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRangeWithSteps_bounds(t *testing.T) {
	t.Run("uint8_wrap", func(t *testing.T) {
		assertRangeBounds(t, it.RangeWithSteps[uint8](254, 255, 2), []uint8{254})
	})
	t.Run("int8_wrap", func(t *testing.T) {
		assertRangeBounds(t, it.RangeWithSteps[int8](126, 127, 2), []int8{126})
	})
	t.Run("int8_underflow", func(t *testing.T) {
		assertRangeBounds(t, it.RangeWithSteps[int8](-127, -128, -2), []int8{-127})
	})
	t.Run("int64_wrap", func(t *testing.T) {
		const max = int64(1<<63 - 1)
		assertRangeBounds(t, it.RangeWithSteps(max-1, max, 2), []int64{max - 1})
	})
	t.Run("int64_underflow", func(t *testing.T) {
		const min = int64(-1 << 63)
		assertRangeBounds(t, it.RangeWithSteps(min+1, min, -2), []int64{min + 1})
	})
	t.Run("uint64_wrap", func(t *testing.T) {
		const max = uint64(1<<64 - 1)
		assertRangeBounds(t, it.RangeWithSteps(max-1, max, 2), []uint64{max - 1})
	})
	t.Run("named_uint8", func(t *testing.T) {
		type level uint8
		assertRangeBounds(t, it.RangeWithSteps[level](254, 255, 2), []level{254})
	})
	t.Run("float32_stagnation", func(t *testing.T) {
		assertRangeBounds(t, it.RangeWithSteps[float32](1<<24, 1<<24+4, 1), []float32{1 << 24})
	})
	t.Run("float64_stagnation", func(t *testing.T) {
		assertRangeBounds(t, it.RangeWithSteps[float64](1<<53, 1<<53+4, 1), []float64{1 << 53})
	})
	t.Run("descending_float64_stagnation", func(t *testing.T) {
		assertRangeBounds(t, it.RangeWithSteps[float64](-1<<53, -1<<53-4, -1), []float64{-1 << 53})
	})
	t.Run("ordinary_ascending", func(t *testing.T) {
		assertRangeBounds(t, it.RangeWithSteps(1, 6, 2), []int{1, 3, 5})
	})
	t.Run("ordinary_descending", func(t *testing.T) {
		assertRangeBounds(t, it.RangeWithSteps(5, 0, -2), []int{5, 3, 1})
	})
}
