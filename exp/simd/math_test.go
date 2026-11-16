//go:build goexperiment.simd

package simd

import (
	"fmt"
	"math"
	"testing"
)

// makeInts returns n deterministic pseudo-random ints in [-1001, 1001]. The magnitude is
// kept small so that sums up to a few hundred elements stay well under float32's 24-bit
// exact-integer range: lane-partitioned float sums then remain bit-identical to a strictly
// sequential scalar sum, and equality assertions don't need a tolerance.
func makeInts(n, seed int) []int {
	out := make([]int, n)
	for i := range out {
		h := uint32(seed)*2654435761 + uint32(i)*40503 //nolint:gosec // deterministic hash, not crypto
		out[i] = int(h%2003) - 1001
	}
	return out
}

// makeData is makeInts converted to T.
func makeData[T numeric](n, seed int) []T {
	ints := makeInts(n, seed)
	out := make([]T, n)
	for i, v := range ints {
		out[i] = T(v)
	}
	return out
}

// ---- scalar reference implementations, mirroring the lo.* algorithms exactly ----

func scalarSum[T numeric](s []T) T {
	var total T
	for i := range s {
		total += s[i]
	}
	return total
}

func scalarMean[T numeric](s []T) T {
	length := T(len(s))
	if length == 0 {
		return 0
	}
	return scalarSum(s) / length
}

func scalarSumBy[TItem any, R numeric](collection []TItem, iteratee func(TItem) R) R {
	var sum R
	for i := range collection {
		sum += iteratee(collection[i])
	}
	return sum
}

func scalarMeanBy[TItem any, R numeric](collection []TItem, iteratee func(TItem) R) R {
	length := R(len(collection))
	if length == 0 {
		return 0
	}
	return scalarSumBy(collection, iteratee) / length
}

// ---- generic table-driven checks ----

// checkSum exercises wrapper (and, when non-nil, the unexported kernel) across sizesAround
// lanes, comparing against a strictly sequential scalar sum. Passing a nil kernel (e.g. for
// named-type instantiations, whose kernel has a different concrete signature) skips the
// kernel-level assertion but still exercises the wrapper's unsafe reinterpret path.
func checkSum[T numeric](t *testing.T, lanes int, wrapper, kernel func([]T) T) {
	t.Helper()
	for _, n := range sizesAround(lanes) {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			t.Parallel()
			input := makeData[T](n, 1)
			want := scalarSum(input)
			if got := wrapper(input); got != want {
				t.Errorf("wrapper = %v, want %v", got, want)
			}
			if n > 0 && kernel != nil {
				if got := kernel(input); got != want {
					t.Errorf("kernel = %v, want %v", got, want)
				}
			}
		})
	}
}

func checkMean[T numeric](t *testing.T, lanes int, wrapper func([]T) T) {
	t.Helper()
	for _, n := range sizesAround(lanes) {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			t.Parallel()
			input := makeData[T](n, 2)
			want := scalarMean(input)
			if got := wrapper(input); got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func checkSumBy[T numeric](t *testing.T, lanes int, wrapper func([]int, func(int) T) T) {
	t.Helper()
	iteratee := func(v int) T { return T(v) }
	for _, n := range sizesAround(lanes) {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			t.Parallel()
			collection := makeInts(n, 3)
			want := scalarSumBy(collection, iteratee)
			if got := wrapper(collection, iteratee); got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func checkMeanBy[T numeric](t *testing.T, lanes int, wrapper func([]int, func(int) T) T) {
	t.Helper()
	iteratee := func(v int) T { return T(v) }
	for _, n := range sizesAround(lanes) {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			t.Parallel()
			collection := makeInts(n, 4)
			want := scalarMeanBy(collection, iteratee)
			if got := wrapper(collection, iteratee); got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// checkClamp exercises the Clamp wrapper (and, when non-nil, the unexported kernel) with
// mn <= mx. min > mx and NaN are covered by dedicated tests below.
func checkClamp[T numeric, S ~[]T](t *testing.T, lanes int, mn, mx T, wrapper func(S, T, T) S, kernel func(dst, src []T, mn, mx T)) {
	t.Helper()
	for _, n := range sizesAround(lanes) {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			t.Parallel()
			input := S(makeData[T](n, 5))
			got := wrapper(input, mn, mx)

			if n == 0 {
				if len(got) != 0 {
					t.Fatalf("expected empty result, got %v", got)
				}
				return
			}
			if len(got) != len(input) {
				t.Fatalf("len(got) = %d, want %d", len(got), len(input))
			}
			for i := range input {
				want := clampScalar(input[i], mn, mx)
				if got[i] != want {
					t.Errorf("index %d: got %v, want %v", i, got[i], want)
				}
			}

			if kernel != nil {
				dst := make([]T, n)
				kernel(dst, []T(input), mn, mx)
				for i := range input {
					want := clampScalar(input[i], mn, mx)
					if dst[i] != want {
						t.Errorf("kernel index %d: got %v, want %v", i, dst[i], want)
					}
				}
			}
		})
	}
}

// ---- Sum ----

func TestSum(t *testing.T) {
	t.Parallel()
	t.Run("int8", func(t *testing.T) { t.Parallel(); checkSum(t, lanes8(), SumInt8[int8], sumInt8) })
	t.Run("int16", func(t *testing.T) { t.Parallel(); checkSum(t, lanes16(), SumInt16[int16], sumInt16) })
	t.Run("int32", func(t *testing.T) { t.Parallel(); checkSum(t, lanes32(), SumInt32[int32], sumInt32) })
	t.Run("int64", func(t *testing.T) { t.Parallel(); checkSum(t, lanes64(), SumInt64[int64], sumInt64) })
	t.Run("uint8", func(t *testing.T) { t.Parallel(); checkSum(t, lanes8(), SumUint8[uint8], sumUint8) })
	t.Run("uint16", func(t *testing.T) { t.Parallel(); checkSum(t, lanes16(), SumUint16[uint16], sumUint16) })
	t.Run("uint32", func(t *testing.T) { t.Parallel(); checkSum(t, lanes32(), SumUint32[uint32], sumUint32) })
	t.Run("uint64", func(t *testing.T) { t.Parallel(); checkSum(t, lanes64(), SumUint64[uint64], sumUint64) })
	t.Run("float32", func(t *testing.T) { t.Parallel(); checkSum(t, lanes32(), SumFloat32[float32], sumFloat32) })
	t.Run("float64", func(t *testing.T) { t.Parallel(); checkSum(t, lanes64(), SumFloat64[float64], sumFloat64) })

	// Named types exercise the unsafe reinterpret path in each wrapper.
	t.Run("myInt8", func(t *testing.T) { t.Parallel(); checkSum[myInt8](t, lanes8(), SumInt8[myInt8], nil) })
	t.Run("myInt16", func(t *testing.T) { t.Parallel(); checkSum[myInt16](t, lanes16(), SumInt16[myInt16], nil) })
	t.Run("myInt32", func(t *testing.T) { t.Parallel(); checkSum[myInt32](t, lanes32(), SumInt32[myInt32], nil) })
	t.Run("myInt64", func(t *testing.T) { t.Parallel(); checkSum[myInt64](t, lanes64(), SumInt64[myInt64], nil) })
	t.Run("myUint8", func(t *testing.T) { t.Parallel(); checkSum[myUint8](t, lanes8(), SumUint8[myUint8], nil) })
	t.Run("myUint16", func(t *testing.T) { t.Parallel(); checkSum[myUint16](t, lanes16(), SumUint16[myUint16], nil) })
	t.Run("myUint32", func(t *testing.T) { t.Parallel(); checkSum[myUint32](t, lanes32(), SumUint32[myUint32], nil) })
	t.Run("myUint64", func(t *testing.T) { t.Parallel(); checkSum[myUint64](t, lanes64(), SumUint64[myUint64], nil) })
	t.Run("myFloat32", func(t *testing.T) { t.Parallel(); checkSum[myFloat32](t, lanes32(), SumFloat32[myFloat32], nil) })
	t.Run("myFloat64", func(t *testing.T) { t.Parallel(); checkSum[myFloat64](t, lanes64(), SumFloat64[myFloat64], nil) })
}

// TestSumOverflow pins the documented wraparound behaviour across the vector/tail boundary:
// it must match lo.Sum's wraparound (both wrap identically) regardless of vector width.
func TestSumOverflow(t *testing.T) {
	t.Parallel()
	if got := SumInt8([]int8{127, 1}); got != -128 {
		t.Errorf("SumInt8 overflow = %v, want -128", got)
	}
	if got := SumUint8([]uint8{255, 2}); got != 1 {
		t.Errorf("SumUint8 overflow = %v, want 1", got)
	}
	if got := SumInt16([]int16{32767, 1}); got != -32768 {
		t.Errorf("SumInt16 overflow = %v, want -32768", got)
	}
	// Overflow that spans a full vector plus the scalar/LoadPart tail.
	big := make([]int8, lanes8()+3)
	for i := range big {
		big[i] = 100
	}
	want := scalarSum(big)
	if got := SumInt8(big); got != want {
		t.Errorf("SumInt8 vector+tail overflow = %v, want %v", got, want)
	}
}

// ---- Mean ----

func TestMean(t *testing.T) {
	t.Parallel()
	t.Run("int8", func(t *testing.T) { t.Parallel(); checkMean(t, lanes8(), MeanInt8[int8]) })
	t.Run("int16", func(t *testing.T) { t.Parallel(); checkMean(t, lanes16(), MeanInt16[int16]) })
	t.Run("int32", func(t *testing.T) { t.Parallel(); checkMean(t, lanes32(), MeanInt32[int32]) })
	t.Run("int64", func(t *testing.T) { t.Parallel(); checkMean(t, lanes64(), MeanInt64[int64]) })
	t.Run("uint8", func(t *testing.T) { t.Parallel(); checkMean(t, lanes8(), MeanUint8[uint8]) })
	t.Run("uint16", func(t *testing.T) { t.Parallel(); checkMean(t, lanes16(), MeanUint16[uint16]) })
	t.Run("uint32", func(t *testing.T) { t.Parallel(); checkMean(t, lanes32(), MeanUint32[uint32]) })
	t.Run("uint64", func(t *testing.T) { t.Parallel(); checkMean(t, lanes64(), MeanUint64[uint64]) })
	t.Run("float32", func(t *testing.T) { t.Parallel(); checkMean(t, lanes32(), MeanFloat32[float32]) })
	t.Run("float64", func(t *testing.T) { t.Parallel(); checkMean(t, lanes64(), MeanFloat64[float64]) })
}

// TestMeanInt8Divisor pins the pre-existing lo.Mean quirk: the divisor is T(len(collection)),
// so on int8 a length above 127 wraps (300 truncates to 44). This is not "fixed" here.
func TestMeanInt8Divisor(t *testing.T) {
	t.Parallel()
	input := make([]int8, 300)
	for i := range input {
		input[i] = 1
	}
	want := scalarMean(input) // divides by int8(300) == 44, same truncation as lo.Mean
	if got := MeanInt8(input); got != want {
		t.Errorf("MeanInt8 = %v, want %v (int8(300) divisor quirk)", got, want)
	}
}

// ---- SumBy / MeanBy ----

func TestSumBy(t *testing.T) {
	t.Parallel()
	t.Run("int8", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes8(), SumByInt8[int, int8]) })
	t.Run("int16", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes16(), SumByInt16[int, int16]) })
	t.Run("int32", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes32(), SumByInt32[int, int32]) })
	t.Run("int64", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes64(), SumByInt64[int, int64]) })
	t.Run("uint8", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes8(), SumByUint8[int, uint8]) })
	t.Run("uint16", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes16(), SumByUint16[int, uint16]) })
	t.Run("uint32", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes32(), SumByUint32[int, uint32]) })
	t.Run("uint64", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes64(), SumByUint64[int, uint64]) })
	t.Run("float32", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes32(), SumByFloat32[int, float32]) })
	t.Run("float64", func(t *testing.T) { t.Parallel(); checkSumBy(t, lanes64(), SumByFloat64[int, float64]) })
}

func TestMeanBy(t *testing.T) {
	t.Parallel()
	t.Run("int8", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes8(), MeanByInt8[int, int8]) })
	t.Run("int16", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes16(), MeanByInt16[int, int16]) })
	t.Run("int32", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes32(), MeanByInt32[int, int32]) })
	t.Run("int64", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes64(), MeanByInt64[int, int64]) })
	t.Run("uint8", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes8(), MeanByUint8[int, uint8]) })
	t.Run("uint16", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes16(), MeanByUint16[int, uint16]) })
	t.Run("uint32", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes32(), MeanByUint32[int, uint32]) })
	t.Run("uint64", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes64(), MeanByUint64[int, uint64]) })
	t.Run("float32", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes32(), MeanByFloat32[int, float32]) })
	t.Run("float64", func(t *testing.T) { t.Parallel(); checkMeanBy(t, lanes64(), MeanByFloat64[int, float64]) })
}

// TestSumByAllocs pins the documented allocation: one (the iteratee output buffer) on the
// SIMD path, zero on the scalar fallback, which calls lo.SumBy directly without building
// that buffer.
//
//nolint:paralleltest // AllocsPerRun reads process-wide allocation counters; parallel siblings would pollute the count.
func TestSumByAllocs(t *testing.T) {
	collection := makeInts(64, 7)
	iteratee := func(v int) int32 { return int32(v) }
	allocs := testing.AllocsPerRun(100, func() {
		_ = SumByInt32(collection, iteratee)
	})
	want := 1.0
	if !useSIMD {
		want = 0
	}
	if allocs != want {
		t.Errorf("SumByInt32 allocs = %v, want %v", allocs, want)
	}
}

// ---- Clamp ----

func TestClamp(t *testing.T) {
	t.Parallel()
	t.Run("int8", func(t *testing.T) {
		t.Parallel()
		checkClamp[int8](t, lanes8(), -50, 50, ClampInt8[int8, []int8], clampInt8)
	})
	t.Run("int16", func(t *testing.T) {
		t.Parallel()
		checkClamp[int16](t, lanes16(), -500, 500, ClampInt16[int16, []int16], clampInt16)
	})
	t.Run("int32", func(t *testing.T) {
		t.Parallel()
		checkClamp[int32](t, lanes32(), -500, 500, ClampInt32[int32, []int32], clampInt32)
	})
	t.Run("int64", func(t *testing.T) {
		t.Parallel()
		checkClamp[int64](t, lanes64(), -500, 500, ClampInt64[int64, []int64], clampInt64)
	})
	t.Run("uint8", func(t *testing.T) {
		t.Parallel()
		checkClamp[uint8](t, lanes8(), 10, 200, ClampUint8[uint8, []uint8], clampUint8)
	})
	t.Run("uint16", func(t *testing.T) {
		t.Parallel()
		checkClamp[uint16](t, lanes16(), 10, 1500, ClampUint16[uint16, []uint16], clampUint16)
	})
	t.Run("uint32", func(t *testing.T) {
		t.Parallel()
		checkClamp[uint32](t, lanes32(), 10, 1500, ClampUint32[uint32, []uint32], clampUint32)
	})
	t.Run("uint64", func(t *testing.T) {
		t.Parallel()
		checkClamp[uint64](t, lanes64(), 10, 1500, ClampUint64[uint64, []uint64], clampUint64)
	})
	t.Run("float32", func(t *testing.T) {
		t.Parallel()
		checkClamp[float32](t, lanes32(), -500, 500, ClampFloat32[float32, []float32], clampFloat32)
	})
	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		checkClamp[float64](t, lanes64(), -500, 500, ClampFloat64[float64, []float64], clampFloat64)
	})

	// Named slice + element types exercise the unsafe reinterpret path.
	t.Run("myInt32", func(t *testing.T) {
		t.Parallel()
		checkClamp[myInt32](t, lanes32(), -500, 500, ClampInt32[myInt32, []myInt32], nil)
	})
}

// TestClampMinGreaterThanMax pins the head/tail consistency fix: when mn > mx, every
// element clamps to mn, both inside and outside the vector loop.
func TestClampMinGreaterThanMax(t *testing.T) {
	t.Parallel()
	input := makeData[int32](2*lanes32()+3, 8)
	got := ClampInt32(input, int32(50), int32(-50))
	for i, v := range got {
		if v != 50 {
			t.Errorf("index %d: got %v, want 50 (mn)", i, v)
		}
	}
}

// TestClampMinEqualsMax pins the constant-output case.
func TestClampMinEqualsMax(t *testing.T) {
	t.Parallel()
	input := makeData[int32](2*lanes32()+1, 9)
	got := ClampInt32(input, int32(7), int32(7))
	for i, v := range got {
		if v != 7 {
			t.Errorf("index %d: got %v, want 7", i, v)
		}
	}
}

// TestClampEmpty pins the "returns the input slice unchanged" aliasing behaviour.
func TestClampEmpty(t *testing.T) {
	t.Parallel()
	var input []int32
	got := ClampInt32(input, int32(0), int32(10))
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}

	nonNilEmpty := []int32{}
	got2 := ClampInt32(nonNilEmpty, int32(0), int32(10))
	if len(got2) != 0 {
		t.Errorf("got %v, want empty", got2)
	}
}

// TestClampAliasing pins that non-empty Clamp allocates a fresh result and never mutates
// the input.
func TestClampAliasing(t *testing.T) {
	t.Parallel()
	input := makeData[int32](2*lanes32()+1, 10)
	inputCopy := append([]int32(nil), input...)

	got := ClampInt32(input, int32(-100), int32(100))

	for i := range input {
		if input[i] != inputCopy[i] {
			t.Fatalf("Clamp mutated input at index %d: %v != %v", i, input[i], inputCopy[i])
		}
	}
	if len(got) > 0 && len(input) > 0 && &got[0] == &input[0] {
		t.Error("Clamp result aliases the input slice")
	}
}

// TestClampNaN pins that Clamp returns a NaN element unchanged, matching lo.Clamp exactly
// (its `<`/`>` comparisons are always false against NaN) on every architecture, both through
// the wrapper and directly through the kernel.
func TestClampNaN(t *testing.T) {
	t.Parallel()
	nan32 := float32(math.NaN())
	if got := ClampFloat32([]float32{1, nan32, -5}, -1, 1)[1]; !math.IsNaN(float64(got)) {
		t.Errorf("ClampFloat32 NaN element = %v, want NaN", got)
	}
	dst32 := make([]float32, 3)
	clampFloat32(dst32, []float32{1, nan32, -5}, -1, 1)
	if !math.IsNaN(float64(dst32[1])) {
		t.Errorf("clampFloat32 NaN element = %v, want NaN", dst32[1])
	}

	nan64 := math.NaN()
	if got := ClampFloat64([]float64{1, nan64, -5}, -1, 1)[1]; !math.IsNaN(got) {
		t.Errorf("ClampFloat64 NaN element = %v, want NaN", got)
	}
	dst64 := make([]float64, 3)
	clampFloat64(dst64, []float64{1, nan64, -5}, -1, 1)
	if !math.IsNaN(dst64[1]) {
		t.Errorf("clampFloat64 NaN element = %v, want NaN", dst64[1])
	}
}
