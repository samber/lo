//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"math/rand/v2"
	"testing"

	"github.com/samber/lo"
)

func TestSumInt8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int8, 16)},
		{"large", make([]int8, 1000)},
		{"negative", []int8{-1, -2, -3, 4, 5}},
		{"mixed", []int8{-128, 0, 127, 1, -1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int8(rand.IntN(256) - 128)
				}
			}

			got := SumInt8x16(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt8x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int16, 8)},
		{"large", make([]int16, 1000)},
		{"negative", []int16{-1, -2, -3, 4, 5}},
		{"mixed", []int16{-32768, 0, 32767, 1, -1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int16(rand.IntN(65536) - 32768)
				}
			}

			got := SumInt16x8(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt16x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 4", make([]int32, 4)},
		{"large", make([]int32, 1000)},
		{"negative", []int32{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int32(rand.Int32())
				}
			}

			got := SumInt32x4(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt32x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 2", []int64{1, 2}},
		{"large", make([]int64, 1000)},
		{"negative", []int64{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Int64()
				}
			}

			got := SumInt64x2(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt64x2() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint8, 16)},
		{"large", make([]uint8, 1000)},
		{"max values", []uint8{255, 255, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := SumUint8x16(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint8x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint16, 8)},
		{"large", make([]uint16, 1000)},
		{"max values", []uint16{65535, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := SumUint16x8(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint16x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 4", make([]uint32, 4)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := SumUint32x4(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint32x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 2", []uint64{1, 2}},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := SumUint64x2(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint64x2() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumFloat32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 4", []float32{1.0, 2.0, 3.0, 4.0}},
		{"large", make([]float32, 1000)},
		{"negative", []float32{-1.1, -2.2, 3.3, 4.4}},
		{"zeros", []float32{0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float32()
				}
			}

			got := SumFloat32x4(tc.input)
			want := lo.Sum(tc.input)

			const epsilon = 1e-2
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumFloat32x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestSumFloat64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 2", []float64{1.0, 2.0}},
		{"large", make([]float64, 1000)},
		{"negative", []float64{-1.1, -2.2, 3.3, 4.4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float64()
				}
			}

			got := SumFloat64x2(tc.input)
			want := lo.Sum(tc.input)

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumFloat64x2() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestSSETypeAlias(t *testing.T) {
	input := []myInt8{1, 2, 3, 4, 5}
	got := SumInt8x16(input)
	want := lo.Sum(input)

	if got != want {
		t.Errorf("SumInt8x16() with type alias = %v, want %v", got, want)
	}
}

func TestClampInt8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []int8
		min   int8
		max   int8
	}{
		{"empty", []int8{}, -10, 10},
		{"single", []int8{42}, -10, 10},
		{"small", []int8{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 16", make([]int8, 16), 5, 10},
		{"large", make([]int8, 1000), -5, 5},
		{"all below min", []int8{-10, -20, -30}, -5, 10},
		{"all above max", []int8{20, 30, 40}, -10, 10},
		{"already clamped", []int8{5, 6, 7}, 5, 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int8(rand.IntN(256) - 128)
				}
			}

			got := ClampInt8x16(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt8x16() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt8x16()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				// Check that the value was properly clamped from the original
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt8x16()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampInt16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []int16
		min   int16
		max   int16
	}{
		{"empty", []int16{}, -100, 100},
		{"single", []int16{42}, -10, 10},
		{"small", []int16{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 8", make([]int16, 8), 50, 100},
		{"large", make([]int16, 1000), -50, 50},
		{"all below min", []int16{-100, -200, -300}, -50, 100},
		{"all above max", []int16{200, 300, 400}, -100, 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int16(rand.IntN(65536) - 32768)
				}
			}

			got := ClampInt16x8(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt16x8() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt16x8()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt16x8()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampInt32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []int32
		min   int32
		max   int32
	}{
		{"empty", []int32{}, -100, 100},
		{"single", []int32{42}, -10, 10},
		{"small", []int32{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 4", []int32{1, 100, -50, 200}, 0, 100},
		{"large", make([]int32, 1000), -50, 50},
		{"negative range", []int32{-100, -50, 0, 50, 100}, -30, -10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Int32()
				}
			}

			got := ClampInt32x4(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt32x4() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt32x4()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt32x4()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampInt64x2(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
		min   int64
		max   int64
	}{
		{"empty", []int64{}, -100, 100},
		{"single", []int64{42}, -10, 10},
		{"small", []int64{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 2", []int64{-100, 200}, -50, 50},
		{"large", make([]int64, 1000), -50, 50},
		{"all below min", []int64{-1000, -2000, -3000}, -500, 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Int64()
				}
			}

			got := ClampInt64x2(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt64x2() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt64x2()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt64x2()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint8
		min   uint8
		max   uint8
	}{
		{"empty", []uint8{}, 10, 100},
		{"single", []uint8{42}, 10, 100},
		{"small", []uint8{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 16", make([]uint8, 16), 50, 100},
		{"large", make([]uint8, 1000), 50, 200},
		{"all below min", []uint8{1, 2, 3}, 10, 100},
		{"all above max", []uint8{200, 225, 250}, 50, 150},
		{"max values", []uint8{255, 255, 255}, 100, 200},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := ClampUint8x16(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint8x16() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint8x16()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint8x16()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint16
		min   uint16
		max   uint16
	}{
		{"empty", []uint16{}, 100, 1000},
		{"single", []uint16{42}, 10, 100},
		{"small", []uint16{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 8", make([]uint16, 8), 500, 1000},
		{"large", make([]uint16, 1000), 500, 5000},
		{"all below min", []uint16{1, 2, 3}, 10, 100},
		{"all above max", []uint16{2000, 3000, 4000}, 100, 1000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := ClampUint16x8(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint16x8() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint16x8()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint16x8()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint32
		min   uint32
		max   uint32
	}{
		{"empty", []uint32{}, 100, 1000},
		{"single", []uint32{42}, 10, 100},
		{"small", []uint32{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 4", []uint32{1, 1000, 50, 2000}, 100, 1000},
		{"large", make([]uint32, 1000), 500, 5000},
		{"all below min", []uint32{1, 2, 3}, 10, 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := ClampUint32x4(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint32x4() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint32x4()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint32x4()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint64x2(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
		min   uint64
		max   uint64
	}{
		{"empty", []uint64{}, 100, 1000},
		{"single", []uint64{42}, 10, 100},
		{"small", []uint64{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 2", []uint64{50, 2000}, 100, 1000},
		{"large", make([]uint64, 1000), 500, 5000},
		{"all below min", []uint64{1, 2, 3}, 10, 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := ClampUint64x2(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint64x2() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint64x2()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint64x2()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampFloat32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []float32
		min   float32
		max   float32
	}{
		{"empty", []float32{}, -10.0, 10.0},
		{"single", []float32{42.5}, -10.0, 10.0},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}, 2.0, 4.0},
		{"exactly 4", []float32{1.0, 10.0, -5.0, 20.0}, -5.0, 10.0},
		{"large", make([]float32, 1000), -5.0, 5.0},
		{"negative range", []float32{-10.0, -5.0, 0.0, 5.0, 10.0}, -3.0, -1.0},
		{"all below min", []float32{-20.0, -30.0, -40.0}, -10.0, 10.0},
		{"all above max", []float32{20.0, 30.0, 40.0}, -10.0, 10.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float32()*200 - 100
				}
			}

			got := ClampFloat32x4(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampFloat32x4() returned length %d, want %d", len(got), len(tc.input))
			}

			const epsilon = 1e-5
			for i, v := range got {
				if v < tc.min-epsilon || v > tc.max+epsilon {
					t.Errorf("ClampFloat32x4()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if diff := v - expected; diff < -epsilon || diff > epsilon {
					t.Errorf("ClampFloat32x4()[%d] = %v, want %v (original: %v, diff: %v)", i, v, expected, original, diff)
				}
			}
		})
	}
}

func TestClampFloat64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []float64
		min   float64
		max   float64
	}{
		{"empty", []float64{}, -10.0, 10.0},
		{"single", []float64{42.5}, -10.0, 10.0},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 2.0, 4.0},
		{"exactly 2", []float64{-100.0, 200.0}, -50.0, 50.0},
		{"large", make([]float64, 1000), -5.0, 5.0},
		{"negative range", []float64{-10.0, -5.0, 0.0, 5.0, 10.0}, -3.0, -1.0},
		{"all below min", []float64{-20.0, -30.0, -40.0}, -10.0, 10.0},
		{"all above max", []float64{20.0, 30.0, 40.0}, -10.0, 10.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float64()*200 - 100
				}
			}

			got := ClampFloat64x2(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampFloat64x2() returned length %d, want %d", len(got), len(tc.input))
			}

			const epsilon = 1e-10
			for i, v := range got {
				if v < tc.min-epsilon || v > tc.max+epsilon {
					t.Errorf("ClampFloat64x2()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if diff := v - expected; diff < -epsilon || diff > epsilon {
					t.Errorf("ClampFloat64x2()[%d] = %v, want %v (original: %v, diff: %v)", i, v, expected, original, diff)
				}
			}
		})
	}
}

// Test type aliases work correctly
func TestSSEClampTypeAlias(t *testing.T) {
	input := []myInt8{-5, 0, 10, 15, 20}
	min := myInt8(0)
	max := myInt8(10)
	got := ClampInt8x16(input, min, max)

	for i, v := range got {
		if v < min || v > max {
			t.Errorf("ClampInt8x16()[%d] with type alias = %v, outside range [%v, %v]", i, v, min, max)
		}
		original := input[i]
		expected := original
		if expected < min {
			expected = min
		} else if expected > max {
			expected = max
		}
		if v != expected {
			t.Errorf("ClampInt8x16()[%d] with type alias = %v, want %v (original: %v)", i, v, expected, original)
		}
	}
}

func TestMeanInt8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int8, 16)},
		{"large", make([]int8, 1000)},
		{"negative", []int8{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int8(rand.IntN(256) - 128)
				}
			}

			got := MeanInt8x16(tc.input)
			want := int8(0)
			if len(tc.input) > 0 {
				want = int8(lo.Sum(tc.input) / int8(len(tc.input)))
			}

			if got != want {
				t.Errorf("MeanInt8x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int16, 8)},
		{"large", make([]int16, 1000)},
		{"negative", []int16{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int16(rand.IntN(65536) - 32768)
				}
			}

			got := MeanInt16x8(tc.input)
			want := int16(0)
			if len(tc.input) > 0 {
				want = int16(lo.Sum(tc.input) / int16(len(tc.input)))
			}

			if got != want {
				t.Errorf("MeanInt16x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 4", make([]int32, 4)},
		{"large", make([]int32, 1000)},
		{"negative", []int32{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int32(rand.Int32())
				}
			}

			got := MeanInt32x4(tc.input)
			want := int32(0)
			if len(tc.input) > 0 {
				want = int32(lo.Sum(tc.input) / int32(len(tc.input)))
			}

			if got != want {
				t.Errorf("MeanInt32x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 2", []int64{1, 2}},
		{"large", make([]int64, 1000)},
		{"negative", []int64{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Int64()
				}
			}

			got := MeanInt64x2(tc.input)
			want := int64(0)
			if len(tc.input) > 0 {
				want = int64(lo.Sum(tc.input) / int64(len(tc.input)))
			}

			if got != want {
				t.Errorf("MeanInt64x2() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint8, 16)},
		{"large", make([]uint8, 1000)},
		{"max values", []uint8{255, 255, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := MeanUint8x16(tc.input)
			sum := lo.Sum(tc.input)
			want := uint8(0)
			if len(tc.input) > 0 {
				want = uint8(uint64(sum) / uint64(len(tc.input)))
			}

			if got != want {
				t.Errorf("MeanUint8x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint16, 8)},
		{"large", make([]uint16, 1000)},
		{"max values", []uint16{65535, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := MeanUint16x8(tc.input)
			sum := lo.Sum(tc.input)
			want := uint16(0)
			if len(tc.input) > 0 {
				want = uint16(uint64(sum) / uint64(len(tc.input)))
			}

			if got != want {
				t.Errorf("MeanUint16x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 4", make([]uint32, 4)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MeanUint32x4(tc.input)
			want := uint32(0)
			if len(tc.input) > 0 {
				want = uint32(uint32(lo.Sum(tc.input)) / uint32(len(tc.input)))
			}

			if got != want {
				t.Errorf("MeanUint32x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 2", []uint64{1, 2}},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MeanUint64x2(tc.input)
			want := uint64(0)
			if len(tc.input) > 0 {
				want = uint64(uint64(lo.Sum(tc.input)) / uint64(len(tc.input)))
			}

			if got != want {
				t.Errorf("MeanUint64x2() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanFloat32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 4", []float32{1.0, 2.0, 3.0, 4.0}},
		{"large", make([]float32, 1000)},
		{"negative", []float32{-1.1, -2.2, 3.3, 4.4}},
		{"zeros", []float32{0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float32()
				}
			}

			got := MeanFloat32x4(tc.input)
			want := float32(0)
			if len(tc.input) > 0 {
				want = float32(lo.Sum(tc.input) / float32(len(tc.input)))
			}

			const epsilon = 1e-5
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MeanFloat32x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMeanFloat64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 2", []float64{1.0, 2.0}},
		{"large", make([]float64, 1000)},
		{"negative", []float64{-1.1, -2.2, 3.3, 4.4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float64()
				}
			}

			got := MeanFloat64x2(tc.input)
			want := float64(0)
			if len(tc.input) > 0 {
				want = float64(lo.Sum(tc.input) / float64(len(tc.input)))
			}

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MeanFloat64x2() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestSSEMeanTypeAlias(t *testing.T) {
	input := []myInt8{1, 2, 3, 4, 5}
	got := MeanInt8x16(input)
	sum := int64(0)
	for _, v := range input {
		sum += int64(v)
	}
	want := myInt8(sum / int64(len(input)))

	if got != want {
		t.Errorf("MeanInt8x16() with type alias = %v, want %v", got, want)
	}
}

func TestMinInt8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int8, 16)},
		{"large", make([]int8, 1000)},
		{"negative", []int8{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int8(rand.IntN(256) - 128)
				}
			}

			got := MinInt8x16(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt8x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinInt16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int16, 8)},
		{"large", make([]int16, 1000)},
		{"negative", []int16{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int16(rand.IntN(65536) - 32768)
				}
			}

			got := MinInt16x8(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt16x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinInt32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 4", []int32{1, 2, 3, 4}},
		{"large", make([]int32, 1000)},
		{"negative", []int32{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Int32()
				}
			}

			got := MinInt32x4(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt32x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinInt64x2(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 2", []int64{1, 2}},
		{"large", make([]int64, 1000)},
		{"negative", []int64{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Int64()
				}
			}

			got := MinInt64x2(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt64x2() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint8, 16)},
		{"large", make([]uint8, 1000)},
		{"max values", []uint8{255, 100, 50}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := MinUint8x16(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint8x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint16, 8)},
		{"large", make([]uint16, 1000)},
		{"max values", []uint16{65535, 1000, 500}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := MinUint16x8(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint16x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 4", []uint32{1, 2, 3, 4}},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MinUint32x4(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint32x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint64x2(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 2", []uint64{1, 2}},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MinUint64x2(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint64x2() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinFloat32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 4", []float32{1.0, 2.0, 3.0, 4.0}},
		{"large", make([]float32, 1000)},
		{"negative", []float32{-1.1, -2.2, 3.3, 4.4}},
		{"zeros", []float32{0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float32()
				}
			}

			got := MinFloat32x4(tc.input)
			want := lo.Min(tc.input)

			const epsilon = 1e-5
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MinFloat32x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMinFloat64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 2", []float64{1.0, 2.0}},
		{"large", make([]float64, 1000)},
		{"negative", []float64{-1.1, -2.2, 3.3, 4.4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float64()
				}
			}

			got := MinFloat64x2(tc.input)
			want := lo.Min(tc.input)

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MinFloat64x2() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestSSEMinTypeAlias(t *testing.T) {
	input := []myInt8{5, 2, 8, 1, 9}
	got := MinInt8x16(input)
	want := myInt8(1)

	if got != want {
		t.Errorf("MinInt8x16() with type alias = %v, want %v", got, want)
	}
}

func TestMaxInt8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int8, 16)},
		{"large", make([]int8, 1000)},
		{"negative", []int8{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int8(rand.IntN(256) - 128)
				}
			}

			got := MaxInt8x16(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt8x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxInt16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int16, 8)},
		{"large", make([]int16, 1000)},
		{"negative", []int16{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = int16(rand.IntN(65536) - 32768)
				}
			}

			got := MaxInt16x8(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt16x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxInt32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 4", []int32{1, 2, 3, 4}},
		{"large", make([]int32, 1000)},
		{"negative", []int32{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Int32()
				}
			}

			got := MaxInt32x4(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt32x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxInt64x2(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 2", []int64{1, 2}},
		{"large", make([]int64, 1000)},
		{"negative", []int64{-1, -2, -3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Int64()
				}
			}

			got := MaxInt64x2(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt64x2() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint8x16(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint8, 16)},
		{"large", make([]uint8, 1000)},
		{"max values", []uint8{255, 100, 50}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := MaxUint8x16(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint8x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint16x8(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint16, 8)},
		{"large", make([]uint16, 1000)},
		{"max values", []uint16{65535, 1000, 500}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := MaxUint16x8(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint16x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 4", []uint32{1, 2, 3, 4}},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MaxUint32x4(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint32x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint64x2(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 2", []uint64{1, 2}},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MaxUint64x2(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint64x2() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxFloat32x4(t *testing.T) {
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 4", []float32{1.0, 2.0, 3.0, 4.0}},
		{"large", make([]float32, 1000)},
		{"negative", []float32{-1.1, -2.2, 3.3, 4.4}},
		{"zeros", []float32{0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float32()
				}
			}

			got := MaxFloat32x4(tc.input)
			want := lo.Max(tc.input)

			const epsilon = 1e-5
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MaxFloat32x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMaxFloat64x2(t *testing.T) {
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 2", []float64{1.0, 2.0}},
		{"large", make([]float64, 1000)},
		{"negative", []float64{-1.1, -2.2, 3.3, 4.4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float64()
				}
			}

			got := MaxFloat64x2(tc.input)
			want := lo.Max(tc.input)

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MaxFloat64x2() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestSSEMaxTypeAlias(t *testing.T) {
	input := []myInt8{5, 2, 8, 1, 9}
	got := MaxInt8x16(input)
	want := myInt8(9)

	if got != want {
		t.Errorf("MaxInt8x16() with type alias = %v, want %v", got, want)
	}
}
