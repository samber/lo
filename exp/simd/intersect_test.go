//go:build goexperiment.simd

package simd

import (
	"fmt"
	"testing"
)

// pickAbsent returns a value of T that is not present in data. It tries a handful of
// candidate values first, falling back to an exhaustive scan over a small range — always
// terminating quickly for the narrow integer types this package accelerates.
func pickAbsent[T numeric](data []T) T {
	present := make(map[T]bool, len(data))
	for _, v := range data {
		present[v] = true
	}
	for _, cand := range []int64{0, 1, -1, 12345, -12345, 999999} {
		v := T(cand)
		if !present[v] {
			return v
		}
	}
	var v T
	for i := range 1 << 16 {
		v = T(i)
		if !present[v] {
			return v
		}
	}
	return v
}

// checkContains exercises the Contains wrapper (and, when non-nil, the unexported kernel)
// with both a guaranteed hit and a guaranteed miss, across sizesAround lanes.
func checkContains[T numeric](t *testing.T, lanes int, wrapper, kernel func([]T, T) bool) {
	t.Helper()
	for _, n := range sizesAround(lanes) {
		t.Run(fmt.Sprintf("n=%d/miss", n), func(t *testing.T) {
			t.Parallel()
			input := makeData[T](n, 21)
			target := pickAbsent(input)
			if got := wrapper(input, target); got {
				t.Errorf("wrapper = true, want false")
			}
			if n > 0 && kernel != nil {
				if got := kernel(input, target); got {
					t.Errorf("kernel = true, want false")
				}
			}
		})

		if n == 0 {
			continue
		}
		t.Run(fmt.Sprintf("n=%d/hit-first", n), func(t *testing.T) {
			t.Parallel()
			input := makeData[T](n, 22)
			target := input[0]
			if got := wrapper(input, target); !got {
				t.Errorf("wrapper = false, want true")
			}
			if kernel != nil {
				if got := kernel(input, target); !got {
					t.Errorf("kernel = false, want true")
				}
			}
		})
		t.Run(fmt.Sprintf("n=%d/hit-last", n), func(t *testing.T) {
			t.Parallel()
			input := makeData[T](n, 23)
			target := input[len(input)-1]
			if got := wrapper(input, target); !got {
				t.Errorf("wrapper = false, want true")
			}
			if kernel != nil {
				if got := kernel(input, target); !got {
					t.Errorf("kernel = false, want true")
				}
			}
		})
	}
}

func TestContains(t *testing.T) {
	t.Parallel()
	t.Run("int8", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes8(), ContainsInt8[int8], containsInt8)
	})
	t.Run("int16", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes16(), ContainsInt16[int16], containsInt16)
	})
	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes32(), ContainsInt32[int32], containsInt32)
	})
	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes64(), ContainsInt64[int64], containsInt64)
	})
	t.Run("uint8", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes8(), ContainsUint8[uint8], containsUint8)
	})
	t.Run("uint16", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes16(), ContainsUint16[uint16], containsUint16)
	})
	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes32(), ContainsUint32[uint32], containsUint32)
	})
	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes64(), ContainsUint64[uint64], containsUint64)
	})
	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes32(), ContainsFloat32[float32], containsFloat32)
	})
	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		checkContains(t, lanes64(), ContainsFloat64[float64], containsFloat64)
	})

	t.Run("myInt32", func(t *testing.T) {
		t.Parallel()
		checkContains[myInt32](t, lanes32(), ContainsInt32[myInt32], nil)
	})
}

// TestContainsZeroTargetTail is the regression test for the LoadPart zero-fill trap:
// LoadPart zero-fills unused lanes, which would produce a false positive for target == 0 if
// the tail ever used it. Every kernel instead uses an overlapping full load for the tail, so
// this must be false for every remainder 1..lanes-1, for every type.
func TestContainsZeroTargetTail(t *testing.T) {
	t.Parallel()

	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		lanes := lanes32()
		for rem := 1; rem < lanes; rem++ {
			n := lanes + rem // one full vector plus a non-zero-length tail
			input := make([]int32, n)
			for i := range input {
				input[i] = int32(i%17) + 1 // guaranteed non-zero
			}
			if got := ContainsInt32(input, int32(0)); got {
				t.Errorf("rem=%d: Contains(..., 0) = true, want false", rem)
			}
		}
	})

	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		lanes := lanes32()
		for rem := 1; rem < lanes; rem++ {
			n := lanes + rem
			input := make([]float32, n)
			for i := range input {
				input[i] = float32(i%17) + 1
			}
			if got := ContainsFloat32(input, float32(0)); got {
				t.Errorf("rem=%d: Contains(..., 0) = true, want false", rem)
			}
		}
	})

	t.Run("int8", func(t *testing.T) {
		t.Parallel()
		lanes := lanes8()
		for rem := 1; rem < lanes; rem++ {
			n := lanes + rem
			input := make([]int8, n)
			for i := range input {
				input[i] = int8(i%17) + 1
			}
			if got := ContainsInt8(input, int8(0)); got {
				t.Errorf("rem=%d: Contains(..., 0) = true, want false", rem)
			}
		}
	})

	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		lanes := lanes64()
		for rem := 1; rem < lanes; rem++ {
			n := lanes + rem
			input := make([]uint64, n)
			for i := range input {
				input[i] = uint64(i%17) + 1
			}
			if got := ContainsUint64(input, uint64(0)); got {
				t.Errorf("rem=%d: Contains(..., 0) = true, want false", rem)
			}
		}
	})
}

// TestContainsOverlapRegionHit pins that a hit only reachable through the overlapping tail
// load (the last element, with a length that leaves a one-element remainder) is found.
func TestContainsOverlapRegionHit(t *testing.T) {
	t.Parallel()
	lanes := lanes32()
	input := make([]int32, lanes+1) // remainder of exactly 1: the overlap covers lanes-1..2*lanes-1
	for i := range input {
		input[i] = 0
	}
	input[len(input)-1] = 99
	if !ContainsInt32(input, int32(99)) {
		t.Error("expected hit in the overlapped tail region")
	}
}

// TestContainsAllocs pins that Contains never allocates.
//
//nolint:paralleltest // AllocsPerRun reads process-wide allocation counters; parallel siblings would pollute the count.
func TestContainsAllocs(t *testing.T) {
	input := makeData[int32](2*lanes32()+3, 24)
	allocs := testing.AllocsPerRun(100, func() {
		_ = ContainsInt32(input, int32(-999999))
	})
	if allocs != 0 {
		t.Errorf("ContainsInt32 allocs = %v, want 0", allocs)
	}
}
