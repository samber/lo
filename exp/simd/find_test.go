//go:build goexperiment.simd

package simd

import (
	"fmt"
	"math"
	"testing"
)

// checkMinMax exercises the Min/Max wrappers and their unexported kernels across
// sizesAround lanes.
func checkMinMax[T numeric](t *testing.T, lanes int, minWrapper, maxWrapper, minKernel, maxKernel func([]T) T) {
	t.Helper()
	for _, n := range sizesAround(lanes) {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			t.Parallel()
			input := makeData[T](n, 11)

			if n == 0 {
				if got := minWrapper(input); got != 0 {
					t.Errorf("min empty = %v, want 0", got)
				}
				if got := maxWrapper(input); got != 0 {
					t.Errorf("max empty = %v, want 0", got)
				}
				return
			}

			wantMin, wantMax := input[0], input[0]
			for _, v := range input[1:] {
				if v < wantMin {
					wantMin = v
				}
				if v > wantMax {
					wantMax = v
				}
			}

			if got := minWrapper(input); got != wantMin {
				t.Errorf("min wrapper = %v, want %v", got, wantMin)
			}
			if minKernel != nil {
				if got := minKernel(input); got != wantMin {
					t.Errorf("min kernel = %v, want %v", got, wantMin)
				}
			}
			if got := maxWrapper(input); got != wantMax {
				t.Errorf("max wrapper = %v, want %v", got, wantMax)
			}
			if maxKernel == nil {
				return
			}
			if got := maxKernel(input); got != wantMax {
				t.Errorf("max kernel = %v, want %v", got, wantMax)
			}
		})
	}
}

func TestMinMax(t *testing.T) {
	t.Parallel()
	t.Run("int8", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes8(), MinInt8[int8], MaxInt8[int8], minInt8, maxInt8)
	})
	t.Run("int16", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes16(), MinInt16[int16], MaxInt16[int16], minInt16, maxInt16)
	})
	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes32(), MinInt32[int32], MaxInt32[int32], minInt32, maxInt32)
	})
	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes64(), MinInt64[int64], MaxInt64[int64], minInt64, maxInt64)
	})
	t.Run("uint8", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes8(), MinUint8[uint8], MaxUint8[uint8], minUint8, maxUint8)
	})
	t.Run("uint16", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes16(), MinUint16[uint16], MaxUint16[uint16], minUint16, maxUint16)
	})
	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes32(), MinUint32[uint32], MaxUint32[uint32], minUint32, maxUint32)
	})
	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes64(), MinUint64[uint64], MaxUint64[uint64], minUint64, maxUint64)
	})
	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes32(), MinFloat32[float32], MaxFloat32[float32], minFloat32, maxFloat32)
	})
	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		checkMinMax(t, lanes64(), MinFloat64[float64], MaxFloat64[float64], minFloat64, maxFloat64)
	})

	// Named types exercise the unsafe reinterpret path; kernels have a different concrete
	// signature for named types, so they are skipped here (nil).
	t.Run("myInt32", func(t *testing.T) {
		t.Parallel()
		checkMinMax[myInt32](t, lanes32(), MinInt32[myInt32], MaxInt32[myInt32], nil, nil)
	})
}

// TestMinMaxSingleElement pins the len==1 scalar prologue, which every type's kernel takes
// when the slice is shorter than one vector.
func TestMinMaxSingleElement(t *testing.T) {
	t.Parallel()
	if got := MinInt32([]int32{42}); got != 42 {
		t.Errorf("MinInt32([42]) = %v, want 42", got)
	}
	if got := MaxInt32([]int32{42}); got != 42 {
		t.Errorf("MaxInt32([42]) = %v, want 42", got)
	}
}

// TestMinMaxNaN pins that Min/Max's NaN handling matches lo.Min/lo.Max exactly, on every
// architecture: NaN elsewhere in the slice is ignored, but NaN at index 0 poisons every
// comparison (NaN < x and NaN > x are always false) and propagates to the result.
func TestMinMaxNaN(t *testing.T) {
	t.Parallel()
	nan32 := float32(math.NaN())
	nanElsewhere := []float32{1, nan32, -5, 3}
	nanFirst := []float32{nan32, 1, -5, 3}

	if got := MinFloat32(nanElsewhere); got != -5 {
		t.Errorf("MinFloat32(NaN elsewhere) = %v, want -5", got)
	}
	if got := minFloat32(nanElsewhere); got != -5 {
		t.Errorf("minFloat32(NaN elsewhere) = %v, want -5", got)
	}
	if got := MaxFloat32(nanElsewhere); got != 3 {
		t.Errorf("MaxFloat32(NaN elsewhere) = %v, want 3", got)
	}
	if got := maxFloat32(nanElsewhere); got != 3 {
		t.Errorf("maxFloat32(NaN elsewhere) = %v, want 3", got)
	}

	if got := MinFloat32(nanFirst); !math.IsNaN(float64(got)) {
		t.Errorf("MinFloat32(NaN first) = %v, want NaN", got)
	}
	if got := minFloat32(nanFirst); !math.IsNaN(float64(got)) {
		t.Errorf("minFloat32(NaN first) = %v, want NaN", got)
	}
	if got := MaxFloat32(nanFirst); !math.IsNaN(float64(got)) {
		t.Errorf("MaxFloat32(NaN first) = %v, want NaN", got)
	}
	if got := maxFloat32(nanFirst); !math.IsNaN(float64(got)) {
		t.Errorf("maxFloat32(NaN first) = %v, want NaN", got)
	}
}

// TestMinMaxAllocs pins that Min/Max never allocate.
//
//nolint:paralleltest // AllocsPerRun reads process-wide allocation counters; parallel siblings would pollute the count.
func TestMinMaxAllocs(t *testing.T) {
	input := makeData[int32](2*lanes32()+3, 12)
	allocs := testing.AllocsPerRun(100, func() {
		_ = MinInt32(input)
		_ = MaxInt32(input)
	})
	if allocs != 0 {
		t.Errorf("Min/MaxInt32 allocs = %v, want 0", allocs)
	}
}
