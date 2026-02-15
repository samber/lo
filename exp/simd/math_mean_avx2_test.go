//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"math/rand/v2"
	"testing"

	"github.com/samber/lo"
)

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
			want := int8(0)
			if len(tc.input) > 0 {
				want = int8(lo.Sum(tc.input) / int8(len(tc.input)))
			}

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
			want := int16(0)
			if len(tc.input) > 0 {
				want = int16(lo.Sum(tc.input) / int16(len(tc.input)))
			}

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
			want := int32(0)
			if len(tc.input) > 0 {
				want = int32(lo.Sum(tc.input) / int32(len(tc.input)))
			}

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
			want := int64(0)
			if len(tc.input) > 0 {
				want = int64(lo.Sum(tc.input) / int64(len(tc.input)))
			}

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
			sum := lo.Sum(tc.input)
			want := uint8(0)
			if len(tc.input) > 0 {
				want = uint8(uint8(sum) / uint8(len(tc.input)))
			}

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
			sum := lo.Sum(tc.input)
			want := uint16(0)
			if len(tc.input) > 0 {
				want = uint16(uint16(sum) / uint16(len(tc.input)))
			}

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
			sum := lo.Sum(tc.input)
			want := uint32(0)
			if len(tc.input) > 0 {
				want = uint32(uint32(sum) / uint32(len(tc.input)))
			}

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
			want := uint64(0)
			if len(tc.input) > 0 {
				want = uint64(lo.Sum(tc.input) / uint64(len(tc.input)))
			}

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
			want := float32(0)
			if len(tc.input) > 0 {
				want = float32(lo.Sum(tc.input) / float32(len(tc.input)))
			}

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
			want := float64(0)
			if len(tc.input) > 0 {
				want = float64(lo.Sum(tc.input) / float64(len(tc.input)))
			}

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
	sum := int64(0)
	for _, v := range input {
		sum += int64(v)
	}
	want := myInt16(sum / int64(len(input)))

	if got != want {
		t.Errorf("MeanInt16x16() with type alias = %v, want %v", got, want)
	}
}
