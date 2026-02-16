//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"math/rand/v2"
	"testing"

	"github.com/samber/lo"
)

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
