//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"math/rand/v2"
	"testing"

	"github.com/samber/lo"
)

func TestSumInt8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 32", make([]int8, 32)},
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

			got := SumInt8x32(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int16, 16)},
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

			got := SumInt16x16(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int32, 8)},
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

			got := SumInt32x8(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumInt64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 4", make([]int64, 4)},
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

			got := SumInt64x4(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumInt64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 32", make([]uint8, 32)},
		{"large", make([]uint8, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := SumUint8x32(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint16, 16)},
		{"large", make([]uint16, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := SumUint16x16(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint32, 8)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := SumUint32x8(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumUint64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 4", make([]uint64, 4)},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := SumUint64x4(tc.input)
			want := lo.Sum(tc.input)

			if got != want {
				t.Errorf("SumUint64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumFloat32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 8", make([]float32, 8)},
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

			got := SumFloat32x8(tc.input)
			want := lo.Sum(tc.input)

			const epsilon = 1e-2
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumFloat32x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestSumFloat64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 4", make([]float64, 4)},
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

			got := SumFloat64x4(tc.input)
			want := lo.Sum(tc.input)

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumFloat64x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX2TypeAlias(t *testing.T) {
	requireAVX2(t)
	input := []myInt16{1, 2, 3, 4, 5}
	got := SumInt16x16(input)
	want := lo.Sum(input)

	if got != want {
		t.Errorf("SumInt16x16() with type alias = %v, want %v", got, want)
	}
}

func TestMeanInt8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 32", make([]int8, 32)},
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

			got := MeanInt8x32(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanInt8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int16, 16)},
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

			got := MeanInt16x16(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanInt16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int32, 8)},
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

			got := MeanInt32x8(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanInt32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanInt64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 4", make([]int64, 4)},
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

			got := MeanInt64x4(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanInt64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 32", make([]uint8, 32)},
		{"large", make([]uint8, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint8(rand.IntN(256))
				}
			}

			got := MeanUint8x32(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanUint8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint16, 16)},
		{"large", make([]uint16, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = uint16(rand.IntN(65536))
				}
			}

			got := MeanUint16x16(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanUint16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint32, 8)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MeanUint32x8(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanUint32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanUint64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 4", make([]uint64, 4)},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MeanUint64x4(tc.input)
			want := lo.Mean(tc.input)

			if got != want {
				t.Errorf("MeanUint64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMeanFloat32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 8", make([]float32, 8)},
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

			got := MeanFloat32x8(tc.input)
			want := lo.Mean(tc.input)

			const epsilon = 1e-5
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MeanFloat32x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMeanFloat64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 4", make([]float64, 4)},
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

			got := MeanFloat64x4(tc.input)
			want := lo.Mean(tc.input)

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MeanFloat64x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX2MeanTypeAlias(t *testing.T) {
	requireAVX2(t)
	input := []myInt16{1, 2, 3, 4, 5}
	got := MeanInt16x16(input)
	want := lo.Mean(input)

	if got != want {
		t.Errorf("MeanInt16x16() with type alias = %v, want %v", got, want)
	}
}

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

func TestMinInt8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 32", make([]int8, 32)},
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

			got := MinInt8x32(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinInt16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int16, 16)},
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

			got := MinInt16x16(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinInt32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int32, 8)},
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

			got := MinInt32x8(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinInt64x4(t *testing.T) {
	requireAVX2(t)
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 4", make([]int64, 4)},
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

			got := MinInt64x4(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinInt64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 32", make([]uint8, 32)},
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

			got := MinUint8x32(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint16, 16)},
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

			got := MinUint16x16(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint32, 8)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MinUint32x8(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinUint64x4(t *testing.T) {
	requireAVX2(t)
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 4", make([]uint64, 4)},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MinUint64x4(tc.input)
			want := lo.Min(tc.input)

			if got != want {
				t.Errorf("MinUint64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMinFloat32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 8", make([]float32, 8)},
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

			got := MinFloat32x8(tc.input)
			want := lo.Min(tc.input)

			const epsilon = 1e-5
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MinFloat32x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMinFloat64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 4", []float64{1.0, 2.0, 3.0, 4.0}},
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

			got := MinFloat64x4(tc.input)
			want := lo.Min(tc.input)

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MinFloat64x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX2MinTypeAlias(t *testing.T) {
	requireAVX2(t)
	input := []myInt32{5, 2, 8, 1, 9}
	got := MinInt32x8(input)
	want := myInt32(1)

	if got != want {
		t.Errorf("MinInt32x8() with type alias = %v, want %v", got, want)
	}
}

func TestMaxInt8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int8
	}{
		{"empty", []int8{}},
		{"single", []int8{42}},
		{"small", []int8{1, 2, 3, 4, 5}},
		{"exactly 32", make([]int8, 32)},
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

			got := MaxInt8x32(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxInt16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int16
	}{
		{"empty", []int16{}},
		{"single", []int16{42}},
		{"small", []int16{1, 2, 3, 4, 5}},
		{"exactly 16", make([]int16, 16)},
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

			got := MaxInt16x16(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxInt32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []int32
	}{
		{"empty", []int32{}},
		{"single", []int32{42}},
		{"small", []int32{1, 2, 3, 4, 5}},
		{"exactly 8", make([]int32, 8)},
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

			got := MaxInt32x8(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxInt64x4(t *testing.T) {
	requireAVX2(t)
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []int64
	}{
		{"empty", []int64{}},
		{"single", []int64{42}},
		{"small", []int64{1, 2, 3, 4, 5}},
		{"exactly 4", make([]int64, 4)},
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

			got := MaxInt64x4(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxInt64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint8x32(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint8
	}{
		{"empty", []uint8{}},
		{"single", []uint8{42}},
		{"small", []uint8{1, 2, 3, 4, 5}},
		{"exactly 32", make([]uint8, 32)},
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

			got := MaxUint8x32(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint16x16(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint16
	}{
		{"empty", []uint16{}},
		{"single", []uint16{42}},
		{"small", []uint16{1, 2, 3, 4, 5}},
		{"exactly 16", make([]uint16, 16)},
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

			got := MaxUint16x16(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []uint32
	}{
		{"empty", []uint32{}},
		{"single", []uint32{42}},
		{"small", []uint32{1, 2, 3, 4, 5}},
		{"exactly 8", make([]uint32, 8)},
		{"large", make([]uint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint32()
				}
			}

			got := MaxUint32x8(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxUint64x4(t *testing.T) {
	requireAVX2(t)
	requireAVX512(t)
	testCases := []struct {
		name  string
		input []uint64
	}{
		{"empty", []uint64{}},
		{"single", []uint64{42}},
		{"small", []uint64{1, 2, 3, 4, 5}},
		{"exactly 4", make([]uint64, 4)},
		{"large", make([]uint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0] == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i] = rand.Uint64()
				}
			}

			got := MaxUint64x4(tc.input)
			want := lo.Max(tc.input)

			if got != want {
				t.Errorf("MaxUint64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestMaxFloat32x8(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float32
	}{
		{"empty", []float32{}},
		{"single", []float32{42.5}},
		{"small", []float32{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 8", make([]float32, 8)},
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

			got := MaxFloat32x8(tc.input)
			want := lo.Max(tc.input)

			const epsilon = 1e-5
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MaxFloat32x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestMaxFloat64x4(t *testing.T) {
	requireAVX2(t)
	testCases := []struct {
		name  string
		input []float64
	}{
		{"empty", []float64{}},
		{"single", []float64{42.5}},
		{"small", []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{"exactly 4", []float64{1.0, 2.0, 3.0, 4.0}},
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

			got := MaxFloat64x4(tc.input)
			want := lo.Max(tc.input)

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("MaxFloat64x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX2MaxTypeAlias(t *testing.T) {
	requireAVX2(t)
	input := []myInt32{5, 2, 8, 1, 9}
	got := MaxInt32x8(input)
	want := myInt32(9)

	if got != want {
		t.Errorf("MaxInt32x8() with type alias = %v, want %v", got, want)
	}
}

// SumBy tests

type avx2Item struct {
	Value      int8
	Weight     int8
	Multiplier int8
}

func TestSumByInt8x32(t *testing.T) {
	requireAVX2(t)

	testCases := []struct {
		name  string
		input []avx2Item
	}{
		{"empty", []avx2Item{}},
		{"single", []avx2Item{{Value: 42}}},
		{"small", []avx2Item{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}},
		{"exactly 32", make([]avx2Item, 32)},
		{"large", make([]avx2Item, 1000)},
		{"negative", []avx2Item{{Value: -1}, {Value: -2}, {Value: -3}, {Value: 4}, {Value: 5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = int8(rand.IntN(256) - 128)
				}
			}

			// Using Value field as the iteratee
			got := SumByInt8x32(tc.input, func(i avx2Item) int8 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2Item, _ int) int8 { return i.Value }))

			if got != want {
				t.Errorf("SumByInt8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByInt16x16(t *testing.T) {
	requireAVX2(t)

	type avx2ItemInt16 struct {
		Value int16
	}

	testCases := []struct {
		name  string
		input []avx2ItemInt16
	}{
		{"empty", []avx2ItemInt16{}},
		{"single", []avx2ItemInt16{{Value: 42}}},
		{"small", []avx2ItemInt16{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 16", make([]avx2ItemInt16, 16)},
		{"large", make([]avx2ItemInt16, 1000)},
		{"negative", []avx2ItemInt16{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = int16(rand.IntN(65536) - 32768)
				}
			}

			got := SumByInt16x16(tc.input, func(i avx2ItemInt16) int16 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2ItemInt16, _ int) int16 { return i.Value }))

			if got != want {
				t.Errorf("SumByInt16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByInt32x8(t *testing.T) {
	requireAVX2(t)

	type avx2ItemInt32 struct {
		Value int32
	}

	testCases := []struct {
		name  string
		input []avx2ItemInt32
	}{
		{"empty", []avx2ItemInt32{}},
		{"single", []avx2ItemInt32{{Value: 42}}},
		{"small", []avx2ItemInt32{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 8", make([]avx2ItemInt32, 8)},
		{"large", make([]avx2ItemInt32, 1000)},
		{"negative", []avx2ItemInt32{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Int32()
				}
			}

			got := SumByInt32x8(tc.input, func(i avx2ItemInt32) int32 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2ItemInt32, _ int) int32 { return i.Value }))

			if got != want {
				t.Errorf("SumByInt32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByInt64x4(t *testing.T) {
	requireAVX2(t)

	type avx2ItemInt64 struct {
		Value int64
	}

	testCases := []struct {
		name  string
		input []avx2ItemInt64
	}{
		{"empty", []avx2ItemInt64{}},
		{"single", []avx2ItemInt64{{Value: 42}}},
		{"small", []avx2ItemInt64{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 4", []avx2ItemInt64{{1}, {2}, {3}, {4}}},
		{"large", make([]avx2ItemInt64, 1000)},
		{"negative", []avx2ItemInt64{{-1}, {-2}, {-3}, {4}, {5}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Int64()
				}
			}

			got := SumByInt64x4(tc.input, func(i avx2ItemInt64) int64 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2ItemInt64, _ int) int64 { return i.Value }))

			if got != want {
				t.Errorf("SumByInt64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByUint8x32(t *testing.T) {
	requireAVX2(t)

	type avx2ItemUint8 struct {
		Value uint8
	}

	testCases := []struct {
		name  string
		input []avx2ItemUint8
	}{
		{"empty", []avx2ItemUint8{}},
		{"single", []avx2ItemUint8{{Value: 42}}},
		{"small", []avx2ItemUint8{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 32", make([]avx2ItemUint8, 32)},
		{"large", make([]avx2ItemUint8, 1000)},
		{"max values", []avx2ItemUint8{{Value: 255}, {Value: 255}, {Value: 1}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = uint8(rand.IntN(256))
				}
			}

			got := SumByUint8x32(tc.input, func(i avx2ItemUint8) uint8 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2ItemUint8, _ int) uint8 { return i.Value }))

			if got != want {
				t.Errorf("SumByUint8x32() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByUint16x16(t *testing.T) {
	requireAVX2(t)

	type avx2ItemUint16 struct {
		Value uint16
	}

	testCases := []struct {
		name  string
		input []avx2ItemUint16
	}{
		{"empty", []avx2ItemUint16{}},
		{"single", []avx2ItemUint16{{Value: 42}}},
		{"small", []avx2ItemUint16{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 16", make([]avx2ItemUint16, 16)},
		{"large", make([]avx2ItemUint16, 1000)},
		{"max values", []avx2ItemUint16{{Value: 65535}, {Value: 1}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = uint16(rand.IntN(65536))
				}
			}

			got := SumByUint16x16(tc.input, func(i avx2ItemUint16) uint16 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2ItemUint16, _ int) uint16 { return i.Value }))

			if got != want {
				t.Errorf("SumByUint16x16() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByUint32x8(t *testing.T) {
	requireAVX2(t)

	type avx2ItemUint32 struct {
		Value uint32
	}

	testCases := []struct {
		name  string
		input []avx2ItemUint32
	}{
		{"empty", []avx2ItemUint32{}},
		{"single", []avx2ItemUint32{{Value: 42}}},
		{"small", []avx2ItemUint32{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 8", make([]avx2ItemUint32, 8)},
		{"large", make([]avx2ItemUint32, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Uint32()
				}
			}

			got := SumByUint32x8(tc.input, func(i avx2ItemUint32) uint32 { return i.Value })
					want := lo.Sum(lo.Map(tc.input, func(i avx2ItemUint32, _ int) uint32 { return i.Value }))

			if got != want {
				t.Errorf("SumByUint32x8() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByUint64x4(t *testing.T) {
	requireAVX2(t)

	type avx2ItemUint64 struct {
		Value uint64
	}

	testCases := []struct {
		name  string
		input []avx2ItemUint64
	}{
		{"empty", []avx2ItemUint64{}},
		{"single", []avx2ItemUint64{{Value: 42}}},
		{"small", []avx2ItemUint64{{1}, {2}, {3}, {4}, {5}}},
		{"exactly 4", []avx2ItemUint64{{1}, {2}, {3}, {4}}},
		{"large", make([]avx2ItemUint64, 1000)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Uint64()
				}
			}

			got := SumByUint64x4(tc.input, func(i avx2ItemUint64) uint64 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2ItemUint64, _ int) uint64 { return i.Value }))

			if got != want {
				t.Errorf("SumByUint64x4() = %v, want %v", got, want)
			}
		})
	}
}

func TestSumByFloat32x8(t *testing.T) {
	requireAVX2(t)

	type avx2ItemFloat32 struct {
		Value float32
	}

	testCases := []struct {
		name  string
		input []avx2ItemFloat32
	}{
		{"empty", []avx2ItemFloat32{}},
		{"single", []avx2ItemFloat32{{Value: 42.5}}},
		{"small", []avx2ItemFloat32{{1.1}, {2.2}, {3.3}, {4.4}, {5.5}}},
		{"exactly 8", make([]avx2ItemFloat32, 8)},
		{"large", make([]avx2ItemFloat32, 1000)},
		{"negative", []avx2ItemFloat32{{-1.1}, {-2.2}, {3.3}, {4.4}}},
		{"zeros", []avx2ItemFloat32{{0}, {0}, {0}, {0}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Float32()
				}
			}

			got := SumByFloat32x8(tc.input, func(i avx2ItemFloat32) float32 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2ItemFloat32, _ int) float32 { return i.Value }))

			const epsilon = 1e-3
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumByFloat32x8() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

func TestSumByFloat64x4(t *testing.T) {
	requireAVX2(t)

	type avx2ItemFloat64 struct {
		Value float64
	}

	testCases := []struct {
		name  string
		input []avx2ItemFloat64
	}{
		{"empty", []avx2ItemFloat64{}},
		{"single", []avx2ItemFloat64{{Value: 42.5}}},
		{"small", []avx2ItemFloat64{{1.1}, {2.2}, {3.3}, {4.4}, {5.5}}},
		{"exactly 4", []avx2ItemFloat64{{1.0}, {2.0}, {3.0}, {4.0}}},
		{"large", make([]avx2ItemFloat64, 1000)},
		{"negative", []avx2ItemFloat64{{-1.1}, {-2.2}, {3.3}, {4.4}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.input) > 0 && tc.input[0].Value == 0 && len(tc.input) > 6 {
				for i := range tc.input {
					tc.input[i].Value = rand.Float64()
				}
			}

			got := SumByFloat64x4(tc.input, func(i avx2ItemFloat64) float64 { return i.Value })
			want := lo.Sum(lo.Map(tc.input, func(i avx2ItemFloat64, _ int) float64 { return i.Value }))

			const epsilon = 1e-10
			if diff := got - want; diff < -epsilon || diff > epsilon {
				t.Errorf("SumByFloat64x4() = %v, want %v (diff: %v)", got, want, diff)
			}
		})
	}
}

// Test type alias works correctly for SumBy
func TestAVX2SumByTypeAlias(t *testing.T) {
	requireAVX2(t)

	type myAVX2Item struct {
		Value myInt8
	}

	input := []myAVX2Item{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}
	got := SumByInt8x32(input, func(i myAVX2Item) myInt8 { return i.Value })
	want := myInt8(15)

	if got != want {
		t.Errorf("SumByInt8x32() with type alias = %v, want %v", got, want)
	}
}
