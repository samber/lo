//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"math/rand/v2"
	"testing"

	"github.com/samber/lo"
)

func TestSumInt8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 64", make([]int8, 64)},
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

			got := SumInt8x64(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 32", make([]int16, 32)},
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

			got := SumInt16x32(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int32, 16)},
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

			got := SumInt32x16(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt32x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int64, 8)},
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

			got := SumInt64x8(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 64", make([]uint8, 64)},
		{"large", make([]uint8, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := SumUint8x64(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 32", make([]uint16, 32)},
		{"large", make([]uint16, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := SumUint16x32(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint32, 16)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := SumUint32x16(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint32x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint64, 8)},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := SumUint64x8(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumFloat32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 16", make([]float32, 16)},
		{"large", make([]float32, 1000)},
		{"negative", []float32{-1.1, -2.2, 3.3, 4.4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float32()
				}
			}

			got := SumFloat32x16(tc.input)
			want := lo.Sum(tc.input)

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumFloat32x16() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestSumFloat64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 8", make([]float64, 8)},
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

			got := SumFloat64x8(tc.input)
			want := lo.Sum(tc.input)

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumFloat64x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX512TypeAlias(t *testing.T) {
	requireAVX512(t)
	input := []myInt32{1, 2, 3, 4, 5}
	got := SumInt32x16(input)
	want := lo.Sum(input)

	if got != want {
		t.Errorf("SumInt32x16() with type alias = %v, want %v", got, want)
	}
}

func TestMeanInt8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 64", make([]int8, 64)},
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

			got := MeanInt8x64(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanInt8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 32", make([]int16, 32)},
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

			got := MeanInt16x32(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanInt16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int32, 16)},
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

			got := MeanInt32x16(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanInt32x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int64, 8)},
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

			got := MeanInt64x8(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanInt64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 64", make([]uint8, 64)},
		{"large", make([]uint8, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := MeanUint8x64(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanUint8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 32", make([]uint16, 32)},
		{"large", make([]uint16, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := MeanUint16x32(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanUint16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint32, 16)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MeanUint32x16(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanUint32x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint64, 8)},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MeanUint64x8(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanUint64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanFloat32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 16", make([]float32, 16)},
		{"large", make([]float32, 1000)},
		{"negative", []float32{-1.1, -2.2, 3.3, 4.4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Float32()
				}
			}

			got := MeanFloat32x16(tc.input)
			want := lo.Mean(tc.input)

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MeanFloat32x16() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMeanFloat64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 8", make([]float64, 8)},
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

			got := MeanFloat64x8(tc.input)
			want := lo.Mean(tc.input)

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MeanFloat64x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX512MeanTypeAlias(t *testing.T) {
	requireAVX512(t)
	input := []myInt32{1, 2, 3, 4, 5}
	got := MeanInt32x16(input)
	want := lo.Mean(input)

	if got != want {
		t.Errorf("MeanInt32x16() with type alias = %v, want %v", got, want)
	}
}

func TestClampInt8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int8
		min   int8
		max   int8
	}{
		{"empty", []int8{}, -10, 10},
		{"single", []int8{42}, -10, 10},
		{"small", []int8{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 64", make([]int8, 64), 5, 10},
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

			got := ClampInt8x64(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt8x64() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt8x64()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt8x64()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampInt16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int16
		min   int16
		max   int16
	}{
		{"empty", []int16{}, -100, 100},
		{"single", []int16{42}, -10, 10},
		{"small", []int16{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 32", make([]int16, 32), 50, 100},
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

			got := ClampInt16x32(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt16x32() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt16x32()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt16x32()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampInt32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int32
		min   int32
		max   int32
	}{
		{"empty", []int32{}, -100, 100},
		{"single", []int32{42}, -10, 10},
		{"small", []int32{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 16", make([]int32, 16), 50, 100},
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

			got := ClampInt32x16(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt32x16() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt32x16()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt32x16()[%d] = %v, want %v (original: %v)", i, v, expected, original)
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

func TestClampInt64x8(t *testing.T) {
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
		{"exactly 8", make([]int64, 8), 50, 100},
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

			got := ClampInt64x8(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampInt64x8() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampInt64x8()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampInt64x8()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint8
		min   uint8
		max   uint8
	}{
		{"empty", []uint8{}, 10, 100},
		{"single", []uint8{42}, 10, 100},
		{"small", []uint8{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 64", make([]uint8, 64), 50, 100},
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

			got := ClampUint8x64(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint8x64() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint8x64()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint8x64()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint16
		min   uint16
		max   uint16
	}{
		{"empty", []uint16{}, 100, 1000},
		{"single", []uint16{42}, 10, 100},
		{"small", []uint16{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 32", make([]uint16, 32), 500, 1000},
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

			got := ClampUint16x32(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint16x32() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint16x32()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint16x32()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampUint32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint32
		min   uint32
		max   uint32
	}{
		{"empty", []uint32{}, 100, 1000},
		{"single", []uint32{42}, 10, 100},
		{"small", []uint32{1, 2, 3, 4, 5}, 2, 4},
		{"exactly 16", make([]uint32, 16), 500, 1000},
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

			got := ClampUint32x16(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint32x16() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint32x16()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint32x16()[%d] = %v, want %v (original: %v)", i, v, expected, original)
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

func TestClampUint64x8(t *testing.T) {
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
		{"exactly 8", make([]uint64, 8), 500, 1000},
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

			got := ClampUint64x8(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampUint64x8() returned length %d, want %d", len(got), len(tc.input))
			}

			for i, v := range got {
				if v < tc.min || v > tc.max {
					t.Errorf("ClampUint64x8()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if v != expected {
					t.Errorf("ClampUint64x8()[%d] = %v, want %v (original: %v)", i, v, expected, original)
				}
			}
		})
	}
}

func TestClampFloat32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float32
		min   float32
		max   float32
	}{
		{"empty", []float32{}, -10.0, 10.0},
		{"single", []float32{42.5}, -10.0, 10.0},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}, 2.0, 4.0},
		{"exactly 16", make([]float32, 16), -5.0, 10.0},
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

			got := ClampFloat32x16(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampFloat32x16() returned length %d, want %d", len(got), len(tc.input))
			}

			const epsilon = 1e-3
			for i, v := range got {
				if v < tc.min-epsilon || v > tc.max+epsilon {
					t.Errorf("ClampFloat32x16()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if diff := v - expected; diff < -epsilon || diff > epsilon {
					t.Errorf("ClampFloat32x16()[%d] = %v, want %v (original: %v, diff: %v)", i, v, expected, original, diff)
				}
			}
		})
	}
}

func TestClampFloat64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float64
		min   float64
		max   float64
	}{
		{"empty", []float64{}, -10.0, 10.0},
		{"single", []float64{42.5}, -10.0, 10.0},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 2.0, 4.0},
		{"exactly 8", make([]float64, 8), -5.0, 10.0},
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

			got := ClampFloat64x8(tc.input, tc.min, tc.max)

			if len(got) != len(tc.input) {
				t.Errorf("ClampFloat64x8() returned length %d, want %d", len(got), len(tc.input))
			}

			const epsilon = 1e-3
			for i, v := range got {
				if v < tc.min-epsilon || v > tc.max+epsilon {
					t.Errorf("ClampFloat64x8()[%d] = %v, outside range [%v, %v]", i, v, tc.min, tc.max)
				}
				original := tc.input[i]
				expected := original
				if expected < tc.min {
					expected = tc.min
				} else if expected > tc.max {
					expected = tc.max
				}
				if diff := v - expected; diff < -epsilon || diff > epsilon {
					t.Errorf("ClampFloat64x8()[%d] = %v, want %v (original: %v, diff: %v)", i, v, expected, original, diff)
				}
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX512ClampTypeAlias(t *testing.T) {
	requireAVX512(t)
	input := []myInt32{-5, 0, 10, 15, 20}
	min := myInt32(0)
	max := myInt32(10)
	got := ClampInt32x16(input, min, max)

	for i, v := range got {
		if v < min || v > max {
			t.Errorf("ClampInt32x16()[%d] with type alias = %v, outside range [%v, %v]", i, v, min, max)
		}
		original := input[i]
		expected := original
		if expected < min {
			expected = min
		} else if expected > max {
			expected = max
		}
		if v != expected {
			t.Errorf("ClampInt32x16()[%d] with type alias = %v, want %v (original: %v)", i, v, expected, original)
		}
	}
}

func TestMinInt8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 64", make([]int8, 64)},
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

			got := MinInt8x64(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinInt16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 32", make([]int16, 32)},
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

			got := MinInt16x32(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinInt32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int32, 16)},
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

			got := MinInt32x16(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt32x16() = %v, want %v", got, want)
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

func TestMinInt64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int64, 8)},
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

			got := MinInt64x8(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 64", make([]uint8, 64)},
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

			got := MinUint8x64(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 32", make([]uint16, 32)},
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

			got := MinUint16x32(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint32, 16)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MinUint32x16(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint32x16() = %v, want %v", got, want)
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

func TestMinUint64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint64, 8)},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MinUint64x8(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinFloat32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 16", make([]float32, 16)},
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

			got := MinFloat32x16(tc.input)
			want := lo.Min(tc.input)

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MinFloat32x16() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMinFloat64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 8", []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}},
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

			got := MinFloat64x8(tc.input)
			want := lo.Min(tc.input)

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MinFloat64x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX512MinTypeAlias(t *testing.T) {
	requireAVX512(t)
	input := []myInt32{5, 2, 8, 1, 9}
	got := MinInt32x16(input)
	want := myInt32(1)

	if got != want {
		t.Errorf("MinInt32x16() with type alias = %v, want %v", got, want)
	}
}

func TestMaxInt8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 64", make([]int8, 64)},
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

			got := MaxInt8x64(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxInt16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 32", make([]int16, 32)},
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

			got := MaxInt16x32(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxInt32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int32, 16)},
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

			got := MaxInt32x16(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt32x16() = %v, want %v", got, want)
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

func TestMaxInt64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int64, 8)},
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

			got := MaxInt64x8(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 64", make([]uint8, 64)},
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

			got := MaxUint8x64(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 32", make([]uint16, 32)},
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

			got := MaxUint16x32(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint32, 16)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MaxUint32x16(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint32x16() = %v, want %v", got, want)
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

func TestMaxUint64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint64, 8)},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MaxUint64x8(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxFloat32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 16", make([]float32, 16)},
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

			got := MaxFloat32x16(tc.input)
			want := lo.Max(tc.input)

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MaxFloat32x16() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMaxFloat64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 8", []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}},
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

			got := MaxFloat64x8(tc.input)
			want := lo.Max(tc.input)

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MaxFloat64x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX512MaxTypeAlias(t *testing.T) {
	requireAVX512(t)
	input := []myInt32{5, 2, 8, 1, 9}
	got := MaxInt32x16(input)
	want := myInt32(9)

	if got != want {
		t.Errorf("MaxInt32x16() with type alias = %v, want %v", got, want)
	}
}

// SumBy tests

type avx512Item struct {
	Value      int8
	Weight     int8
	Multiplier int8
}

func TestSumByInt8x64(t *testing.T) {
	requireAVX512(t)

	testCases := []struct {
		name  string
		input []avx512Item
	}{
		{"empty", []avx512Item{}},
		{"single", []avx512Item{{Value: 42}}},
		{"small", []avx512Item{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}},
		{"exactly 64", make([]avx512Item, 64)},
		{"large", make([]avx512Item, 1000)},
		{"negative", []avx512Item{{Value: -1}, {Value: -2}, {Value: -3}, {Value: 4}, {Value: 5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = int8(rand.IntN(256) - 128)
				}
			}

			// Using Value field as the iteratee
			got := SumByInt8x64(tc.input, func(i avx512Item) int8 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512Item, _ int) int8 { return i.Value }))

			if got != want {
				t.Errorf("SumByInt8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByInt16x32(t *testing.T) {
	requireAVX512(t)

	type avx512ItemInt16 struct {
		Value int16
	}

	testCases := []struct {
		name  string
		input []avx512ItemInt16
	}{
		{"empty", []avx512ItemInt16{}},
		{"single", []avx512ItemInt16{{Value: 42}}},
		{"small", []avx512ItemInt16{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 32", make([]avx512ItemInt16, 32)},
		{"large", make([]avx512ItemInt16, 1000)},
		{"negative", []avx512ItemInt16{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = int16(rand.IntN(65536) - 32768)
				}
			}

			got := SumByInt16x32(tc.input, func(i avx512ItemInt16) int16 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemInt16, _ int) int16 { return i.Value }))

			if got != want {
				t.Errorf("SumByInt16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByInt32x16(t *testing.T) {
	requireAVX512(t)

	type avx512ItemInt32 struct {
		Value int32
	}

	testCases := []struct {
		name  string
		input []avx512ItemInt32
	}{
		{"empty", []avx512ItemInt32{}},
		{"single", []avx512ItemInt32{{Value: 42}}},
		{"small", []avx512ItemInt32{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 16", make([]avx512ItemInt32, 16)},
		{"large", make([]avx512ItemInt32, 1000)},
		{"negative", []avx512ItemInt32{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Int32()
				}
			}

			got := SumByInt32x16(tc.input, func(i avx512ItemInt32) int32 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemInt32, _ int) int32 { return i.Value }))

			if got != want {
				t.Errorf("SumByInt32x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByInt64x8(t *testing.T) {
	requireAVX512(t)

	type avx512ItemInt64 struct {
		Value int64
	}

	testCases := []struct {
		name  string
		input []avx512ItemInt64
	}{
		{"empty", []avx512ItemInt64{}},
		{"single", []avx512ItemInt64{{Value: 42}}},
		{"small", []avx512ItemInt64{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 8", make([]avx512ItemInt64, 8)},
		{"large", make([]avx512ItemInt64, 1000)},
		{"negative", []avx512ItemInt64{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Int64()
				}
			}

			got := SumByInt64x8(tc.input, func(i avx512ItemInt64) int64 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemInt64, _ int) int64 { return i.Value }))

			if got != want {
				t.Errorf("SumByInt64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByUint8x64(t *testing.T) {
	requireAVX512(t)

	type avx512ItemUint8 struct {
		Value uint8
	}

	testCases := []struct {
		name  string
		input []avx512ItemUint8
	}{
		{"empty", []avx512ItemUint8{}},
		{"single", []avx512ItemUint8{{Value: 42}}},
		{"small", []avx512ItemUint8{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 64", make([]avx512ItemUint8, 64)},
		{"large", make([]avx512ItemUint8, 1000)},
		{"max values", []avx512ItemUint8{{Value: 255}, {Value: 255}, {Value: 1}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = uint8(rand.IntN(256))
				}
			}

			got := SumByUint8x64(tc.input, func(i avx512ItemUint8) uint8 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemUint8, _ int) uint8 { return i.Value }))

			if got != want {
				t.Errorf("SumByUint8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByUint16x32(t *testing.T) {
	requireAVX512(t)

	type avx512ItemUint16 struct {
		Value uint16
	}

	testCases := []struct {
		name  string
		input []avx512ItemUint16
	}{
		{"empty", []avx512ItemUint16{}},
		{"single", []avx512ItemUint16{{Value: 42}}},
		{"small", []avx512ItemUint16{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 32", make([]avx512ItemUint16, 32)},
		{"large", make([]avx512ItemUint16, 1000)},
		{"max values", []avx512ItemUint16{{Value: 65535}, {Value: 1}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = uint16(rand.IntN(65536))
				}
			}

			got := SumByUint16x32(tc.input, func(i avx512ItemUint16) uint16 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemUint16, _ int) uint16 { return i.Value }))

			if got != want {
				t.Errorf("SumByUint16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByUint32x16(t *testing.T) {
	requireAVX512(t)

	type avx512ItemUint32 struct {
		Value uint32
	}

	testCases := []struct {
		name  string
		input []avx512ItemUint32
	}{
		{"empty", []avx512ItemUint32{}},
		{"single", []avx512ItemUint32{{Value: 42}}},
		{"small", []avx512ItemUint32{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 16", make([]avx512ItemUint32, 16)},
		{"large", make([]avx512ItemUint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Uint32()
				}
			}

			got := SumByUint32x16(tc.input, func(i avx512ItemUint32) uint32 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemUint32, _ int) uint32 { return i.Value }))

			if got != want {
				t.Errorf("SumByUint32x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByUint64x8(t *testing.T) {
	requireAVX512(t)

	type avx512ItemUint64 struct {
		Value uint64
	}

	testCases := []struct {
		name  string
		input []avx512ItemUint64
	}{
		{"empty", []avx512ItemUint64{}},
		{"single", []avx512ItemUint64{{Value: 42}}},
		{"small", []avx512ItemUint64{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 8", make([]avx512ItemUint64, 8)},
		{"large", make([]avx512ItemUint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Uint64()
				}
			}

			got := SumByUint64x8(tc.input, func(i avx512ItemUint64) uint64 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemUint64, _ int) uint64 { return i.Value }))

			if got != want {
				t.Errorf("SumByUint64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByFloat32x16(t *testing.T) {
	requireAVX512(t)

	type avx512ItemFloat32 struct {
		Value float32
	}

	testCases := []struct {
		name  string
		input []avx512ItemFloat32
	}{
		{"empty", []avx512ItemFloat32{}},
		{"single", []avx512ItemFloat32{{Value: 42.5}}},
		{"small", []avx512ItemFloat32{{1.1}, {2.2}, {3.3}, {4.4}, {5.5}}},
		{"exactly 16", make([]avx512ItemFloat32, 16)},
		{"large", make([]avx512ItemFloat32, 1000)},
		{"negative", []avx512ItemFloat32{{-1.1}, {-2.2}, {3.3}, {4.4}}},
		{"zeros", []avx512ItemFloat32{{0}, {0}, {0}, {0}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Float32()
				}
			}

			got := SumByFloat32x16(tc.input, func(i avx512ItemFloat32) float32 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemFloat32, _ int) float32 { return i.Value }))

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumByFloat32x16() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestSumByFloat64x8(t *testing.T) {
	requireAVX512(t)

	type avx512ItemFloat64 struct {
		Value float64
	}

	testCases := []struct {
		name  string
		input []avx512ItemFloat64
	}{
		{"empty", []avx512ItemFloat64{}},
		{"single", []avx512ItemFloat64{{Value: 42.5}}},
		{"small", []avx512ItemFloat64{{1.1}, {2.2}, {3.3}, {4.4}, {5.5}}},
		{"exactly 8", make([]avx512ItemFloat64, 8)},
		{"large", make([]avx512ItemFloat64, 1000)},
		{"negative", []avx512ItemFloat64{{-1.1}, {-2.2}, {3.3}, {4.4}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Float64()
				}
			}

			got := SumByFloat64x8(tc.input, func(i avx512ItemFloat64) float64 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx512ItemFloat64, _ int) float64 { return i.Value }))

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumByFloat64x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type alias works correctly for SumBy
func TestAVX512SumByTypeAlias(t *testing.T) {
	requireAVX512(t)

	type myAVX512Item struct {
		Value myInt8
	}

	input := []myAVX512Item{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}
	got := SumByInt8x64(input, func(i myAVX512Item) myInt8 { return i.Value })
	want := myInt8(15)

	if got != want {
		t.Errorf("SumByInt8x64() with type alias = %v, want %v", got, want)
	}
}

// MeanBy tests

func TestMeanByInt8x64(t *testing.T) {
	requireAVX512(t)

	testCases := []struct {
		name  string
		input []avx512Item
	}{
		{"empty", []avx512Item{}},
		{"single", []avx512Item{{Value: 42}}},
		{"small", []avx512Item{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}},
		{"exactly 64", make([]avx512Item, 64)},
		{"large", make([]avx512Item, 1000)},
		{"negative", []avx512Item{{Value: -1}, {Value: -2}, {Value: -3}, {Value: 4}, {Value: 5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = int8(rand.IntN(256) - 128)
				}
			}

			got := MeanByInt8x64(tc.input, func(i avx512Item) int8 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512Item, _ int) int8 { return i.Value }))

			if got != want {
				t.Errorf("MeanByInt8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanByInt16x32(t *testing.T) {
	requireAVX512(t)

	type avx512ItemInt16 struct {
		Value int16
	}

	testCases := []struct {
		name  string
		input []avx512ItemInt16
	}{
		{"empty", []avx512ItemInt16{}},
		{"single", []avx512ItemInt16{{Value: 42}}},
		{"small", []avx512ItemInt16{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 32", make([]avx512ItemInt16, 32)},
		{"large", make([]avx512ItemInt16, 1000)},
		{"negative", []avx512ItemInt16{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = int16(rand.IntN(65536) - 32768)
				}
			}

			got := MeanByInt16x32(tc.input, func(i avx512ItemInt16) int16 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemInt16, _ int) int16 { return i.Value }))

			if got != want {
				t.Errorf("MeanByInt16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanByInt32x16(t *testing.T) {
	requireAVX512(t)

	type avx512ItemInt32 struct {
		Value int32
	}

	testCases := []struct {
		name  string
		input []avx512ItemInt32
	}{
		{"empty", []avx512ItemInt32{}},
		{"single", []avx512ItemInt32{{Value: 42}}},
		{"small", []avx512ItemInt32{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 16", make([]avx512ItemInt32, 16)},
		{"large", make([]avx512ItemInt32, 1000)},
		{"negative", []avx512ItemInt32{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Int32()
				}
			}

			got := MeanByInt32x16(tc.input, func(i avx512ItemInt32) int32 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemInt32, _ int) int32 { return i.Value }))

			if got != want {
				t.Errorf("MeanByInt32x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanByInt64x8(t *testing.T) {
	requireAVX512(t)

	type avx512ItemInt64 struct {
		Value int64
	}

	testCases := []struct {
		name  string
		input []avx512ItemInt64
	}{
		{"empty", []avx512ItemInt64{}},
		{"single", []avx512ItemInt64{{Value: 42}}},
		{"small", []avx512ItemInt64{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 8", make([]avx512ItemInt64, 8)},
		{"large", make([]avx512ItemInt64, 1000)},
		{"negative", []avx512ItemInt64{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Int64()
				}
			}

			got := MeanByInt64x8(tc.input, func(i avx512ItemInt64) int64 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemInt64, _ int) int64 { return i.Value }))

			if got != want {
				t.Errorf("MeanByInt64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanByUint8x64(t *testing.T) {
	requireAVX512(t)

	type avx512ItemUint8 struct {
		Value uint8
	}

	testCases := []struct {
		name  string
		input []avx512ItemUint8
	}{
		{"empty", []avx512ItemUint8{}},
		{"single", []avx512ItemUint8{{Value: 42}}},
		{"small", []avx512ItemUint8{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 64", make([]avx512ItemUint8, 64)},
		{"large", make([]avx512ItemUint8, 1000)},
		{"max values", []avx512ItemUint8{{Value: 255}, {Value: 255}, {Value: 1}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = uint8(rand.IntN(256))
				}
			}

			got := MeanByUint8x64(tc.input, func(i avx512ItemUint8) uint8 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemUint8, _ int) uint8 { return i.Value }))

			if got != want {
				t.Errorf("MeanByUint8x64() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanByUint16x32(t *testing.T) {
	requireAVX512(t)

	type avx512ItemUint16 struct {
		Value uint16
	}

	testCases := []struct {
		name  string
		input []avx512ItemUint16
	}{
		{"empty", []avx512ItemUint16{}},
		{"single", []avx512ItemUint16{{Value: 42}}},
		{"small", []avx512ItemUint16{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 32", make([]avx512ItemUint16, 32)},
		{"large", make([]avx512ItemUint16, 1000)},
		{"max values", []avx512ItemUint16{{Value: 65535}, {Value: 1}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = uint16(rand.IntN(65536))
				}
			}

			got := MeanByUint16x32(tc.input, func(i avx512ItemUint16) uint16 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemUint16, _ int) uint16 { return i.Value }))

			if got != want {
				t.Errorf("MeanByUint16x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanByUint32x16(t *testing.T) {
	requireAVX512(t)

	type avx512ItemUint32 struct {
		Value uint32
	}

	testCases := []struct {
		name  string
		input []avx512ItemUint32
	}{
		{"empty", []avx512ItemUint32{}},
		{"single", []avx512ItemUint32{{Value: 42}}},
		{"small", []avx512ItemUint32{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 16", make([]avx512ItemUint32, 16)},
		{"large", make([]avx512ItemUint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Uint32()
				}
			}

			got := MeanByUint32x16(tc.input, func(i avx512ItemUint32) uint32 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemUint32, _ int) uint32 { return i.Value }))

			if got != want {
				t.Errorf("MeanByUint32x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanByUint64x8(t *testing.T) {
	requireAVX512(t)

	type avx512ItemUint64 struct {
		Value uint64
	}

	testCases := []struct {
		name  string
		input []avx512ItemUint64
	}{
		{"empty", []avx512ItemUint64{}},
		{"single", []avx512ItemUint64{{Value: 42}}},
		{"small", []avx512ItemUint64{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 8", make([]avx512ItemUint64, 8)},
		{"large", make([]avx512ItemUint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Uint64()
				}
			}

			got := MeanByUint64x8(tc.input, func(i avx512ItemUint64) uint64 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemUint64, _ int) uint64 { return i.Value }))

			if got != want {
				t.Errorf("MeanByUint64x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanByFloat32x16(t *testing.T) {
	requireAVX512(t)

	type avx512ItemFloat32 struct {
		Value float32
	}

	testCases := []struct {
		name  string
		input []avx512ItemFloat32
	}{
		{"empty", []avx512ItemFloat32{}},
		{"single", []avx512ItemFloat32{{Value: 42.5}}},
		{"small", []avx512ItemFloat32{{1.1}, {2.2}, {3.3}, {4.4}, {5.5}}},
		{"exactly 16", make([]avx512ItemFloat32, 16)},
		{"large", make([]avx512ItemFloat32, 1000)},
		{"negative", []avx512ItemFloat32{{-1.1}, {-2.2}, {3.3}, {4.4}}},
		{"zeros", []avx512ItemFloat32{{0}, {0}, {0}, {0}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Float32()
				}
			}

			got := MeanByFloat32x16(tc.input, func(i avx512ItemFloat32) float32 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemFloat32, _ int) float32 { return i.Value }))

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MeanByFloat32x16() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMeanByFloat64x8(t *testing.T) {
	requireAVX512(t)

	type avx512ItemFloat64 struct {
		Value float64
	}

	testCases := []struct {
		name  string
		input []avx512ItemFloat64
	}{
		{"empty", []avx512ItemFloat64{}},
		{"single", []avx512ItemFloat64{{Value: 42.5}}},
		{"small", []avx512ItemFloat64{{1.1}, {2.2}, {3.3}, {4.4}, {5.5}}},
		{"exactly 8", make([]avx512ItemFloat64, 8)},
		{"large", make([]avx512ItemFloat64, 1000)},
		{"negative", []avx512ItemFloat64{{-1.1}, {-2.2}, {3.3}, {4.4}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Float64()
				}
			}

			got := MeanByFloat64x8(tc.input, func(i avx512ItemFloat64) float64 { return i.Value })
			want := lo.Mean(lo.Map(tc.input, func(i avx512ItemFloat64, _ int) float64 { return i.Value }))

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MeanByFloat64x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type alias works correctly for MeanBy
func TestAVX512MeanByTypeAlias(t *testing.T) {
	requireAVX512(t)

	type myAVX512Item struct {
		Value myInt8
	}

	input := []myAVX512Item{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}
	got := MeanByInt8x64(input, func(i myAVX512Item) myInt8 { return i.Value })
	want := myInt8(3)

	if got != want {
		t.Errorf("MeanByInt8x64() with type alias = %v, want %v", got, want)
	}
}
