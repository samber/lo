//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"math/rand/v2"
	"testing"
)

func TestClampInt8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int8
		min   int8
		max   int8
	}{
		{"empty", []int8{}, -10, 10},
		{"single", []int8{42}, -10, 10},
		{"small", []int8{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 32", make([]int8, 32), 5, 10},
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

			got := ClampInt8x32(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt8x32() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt8x32()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt8x32()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampInt16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int16
		min   int16
		max   int16
	}{
		{"empty", []int16{}, -100, 100},
		{"single", []int16{42}, -10, 10},
		{"small", []int16{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 16", make([]int16, 16), 50, 100},
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

			got := ClampInt16x16(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt16x16() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt16x16()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt16x16()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampInt32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int32
		min   int32
		max   int32
	}{
		{"empty", []int32{}, -100, 100},
		{"single", []int32{42}, -10, 10},
		{"small", []int32{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 8", make([]int32, 8), 50, 100},
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

			got := ClampInt32x8(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt32x8() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt32x8()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt32x8()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampInt64x4(t *testing.T) {
	requireAVX2(t)
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
		{"exactly 4", make([]int64, 4), 50, 100},
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

			got := ClampInt64x4(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt64x4() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt64x4()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt64x4()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint8
		min   uint8
		max   uint8
	}{
		{"empty", []uint8{}, 10, 100},
		{"single", []uint8{42}, 10, 100},
		{"small", []uint8{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 32", make([]uint8, 32), 50, 100},
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

			got := ClampUint8x32(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint8x32() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint8x32()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint8x32()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint16
		min   uint16
		max   uint16
	}{
		{"empty", []uint16{}, 100, 1000},
		{"single", []uint16{42}, 10, 100},
		{"small", []uint16{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 16", make([]uint16, 16), 500, 1000},
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

			got := ClampUint16x16(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint16x16() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint16x16()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint16x16()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint32
		min   uint32
		max   uint32
	}{
		{"empty", []uint32{}, 100, 1000},
		{"single", []uint32{42}, 10, 100},
		{"small", []uint32{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 8", make([]uint32, 8), 500, 1000},
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

			got := ClampUint32x8(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint32x8() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint32x8()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint32x8()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint64x4(t *testing.T) {
	requireAVX2(t)
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
		{"exactly 4", make([]uint64, 4), 500, 1000},
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

			got := ClampUint64x4(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint64x4() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint64x4()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint64x4()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampFloat32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float32
		min   float32
		max   float32
	}{
		{"empty", []float32{}, -10.0, 10.0},
		{"single", []float32{42.5}, -10.0, 10.0},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}, 2.0, 4.0},
		{"exactly 8", make([]float32, 8), -5.0, 10.0},
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

			got := ClampFloat32x8(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampFloat32x8() returned length %d, want %d", len(got), len(tc.input))
			}

			const epsilon = 1e-5
			for i, v := range got {
				if v < tc.min-epsilon || v > tc.max+epsilon {
					t.Errorf("ClampFloat32x8()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if diff := v - expected; diff < -epsilon || diff > epsilon {
					t.Errorf("ClampFloat32x8()[%d] = %v, want %v (original: %v, diff: %v)", i, v, expected, original, diff)
				}
			}
		})
	}
}

func TestClampFloat64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float64
		min   float64
		max   float64
	}{
		{"empty", []float64{}, -10.0, 10.0},
		{"single", []float64{42.5}, -10.0, 10.0},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 2.0, 4.0},
		{"exactly 4", make([]float64, 4), -5.0, 10.0},
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

			got := ClampFloat64x4(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampFloat64x4() returned length %d, want %d", len(got), len(tc.input))
			}

			const epsilon = 1e-10
			for i, v := range got {
				if v < tc.min-epsilon || v > tc.max+epsilon {
					t.Errorf("ClampFloat64x4()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if diff := v - expected; diff < -epsilon || diff > epsilon {
					t.Errorf("ClampFloat64x4()[%d] = %v, want %v (original: %v, diff: %v)", i, v, expected, original, diff)
				}
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX2ClampTypeAlias(t *testing.T) {
	requireAVX2(t)
	input := []myInt32{-5, 0, 10, 15, 20}
	min := myInt32(0)
	max := myInt32(10)
	got := ClampInt32x8(input, min, max)

	for i, v := range got {
		if v < min || v > max {
			t.Errorf("ClampInt32x8()[%d] with type alias = %v, outside range [%v, %v]", i, v, min, max)
		}
		original := input[i]
		expected := original
		if expected < min {
			expected = min
		} else if expected > max {
			expected = max
		}
		if v != expected {
			t.Errorf("ClampInt32x8()[%d] with type alias = %v, want %v (original: %v)", i, v, expected, original)
		}
	}
}
