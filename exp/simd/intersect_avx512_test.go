//go:build go1.26 && goexperiment.simd && amd64

package simd

import (
	"math/rand/v2"
	"testing"

	"github.com/samber/lo"
)

func TestAVX512ContainsInt8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []int8
		target     int8
		expected   bool
	}{
		{"empty", []int8{}, 42, false},
		{"single found", []int8{42}, 42, true},
		{"single not found", []int8{42}, 10, false},
		{"small found", []int8{1, 2, 3, 4, 5}, 3, true},
		{"small not found", []int8{1, 2, 3, 4, 5}, 10, false},
		{"exactly 64 found", make([]int8, 64), 0, true},
		{"exactly 64 not found", make([]int8, 64), 127, false},
		{"large found", make([]int8, 1000), 0, true},
		{"large not found", make([]int8, 100), 127, false},
		{"negative found", []int8{-1, -2, -3, 4, 5}, -2, true},
		{"min value", []int8{-128, 0, 127}, -128, true},
		{"max value", []int8{-128, 0, 127}, 127, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 64 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[100] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = int8(rand.IntN(256) - 128)
						if tc.collection[i] == tc.target {
							tc.collection[i]++
						}
					}
				}
			}

			got := ContainsInt8x64(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsInt8x64() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsInt16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []int16
		target     int16
		expected   bool
	}{
		{"empty", []int16{}, 42, false},
		{"single found", []int16{42}, 42, true},
		{"single not found", []int16{42}, 10, false},
		{"small found", []int16{1, 2, 3, 4, 5}, 3, true},
		{"small not found", []int16{1, 2, 3, 4, 5}, 10, false},
		{"exactly 32 found", make([]int16, 32), 0, true},
		{"exactly 32 not found", make([]int16, 32), 1000, false},
		{"large found", make([]int16, 1000), 0, true},
		{"large not found", make([]int16, 100), 1000, false},
		{"negative found", []int16{-1, -2, -3, 4, 5}, -2, true},
		{"min value", []int16{-32768, 0, 32767}, -32768, true},
		{"max value", []int16{-32768, 0, 32767}, 32767, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 32 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[50] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = int16(rand.IntN(65536) - 32768)
						if tc.collection[i] == tc.target {
							tc.collection[i]++
						}
					}
				}
			}

			got := ContainsInt16x32(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsInt16x32() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsInt32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []int32
		target     int32
		expected   bool
	}{
		{"empty", []int32{}, 42, false},
		{"single found", []int32{42}, 42, true},
		{"single not found", []int32{42}, 10, false},
		{"small found", []int32{1, 2, 3, 4, 5}, 3, true},
		{"small not found", []int32{1, 2, 3, 4, 5}, 10, false},
		{"exactly 16 found", make([]int32, 16), 0, true},
		{"exactly 16 not found", make([]int32, 16), 100, false},
		{"large found", make([]int32, 1000), 0, true},
		{"large not found", make([]int32, 100), 100000, false},
		{"negative found", []int32{-1, -2, -3, 4, 5}, -2, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 16 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[50] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = rand.Int32()
						if tc.collection[i] == tc.target {
							tc.collection[i]++
						}
					}
				}
			}

			got := ContainsInt32x16(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsInt32x16() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsInt64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []int64
		target     int64
		expected   bool
	}{
		{"empty", []int64{}, 42, false},
		{"single found", []int64{42}, 42, true},
		{"single not found", []int64{42}, 10, false},
		{"small found", []int64{1, 2, 3, 4, 5}, 3, true},
		{"small not found", []int64{1, 2, 3, 4, 5}, 10, false},
		{"exactly 8 found", make([]int64, 8), 0, true},
		{"exactly 8 not found", make([]int64, 8), 100, false},
		{"large found", make([]int64, 1000), 0, true},
		{"large not found", make([]int64, 100), 1000000, false},
		{"negative found", []int64{-1, -2, -3, 4, 5}, -2, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 8 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[50] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = rand.Int64()
						if tc.collection[i] == tc.target {
							tc.collection[i]++
						}
					}
				}
			}

			got := ContainsInt64x8(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsInt64x8() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsUint8x64(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []uint8
		target     uint8
		expected   bool
	}{
		{"empty", []uint8{}, 42, false},
		{"single found", []uint8{42}, 42, true},
		{"single not found", []uint8{42}, 10, false},
		{"small found", []uint8{1, 2, 3, 4, 5}, 3, true},
		{"small not found", []uint8{1, 2, 3, 4, 5}, 10, false},
		{"exactly 64 found", make([]uint8, 64), 0, true},
		{"exactly 64 not found", make([]uint8, 64), 255, false},
		{"large found", make([]uint8, 1000), 0, true},
		{"large not found", make([]uint8, 100), 255, false},
		{"max value", []uint8{255, 100, 50}, 255, true},
		{"zero found", []uint8{0, 1, 2, 3}, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 64 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[100] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = uint8(rand.IntN(256))
						if tc.collection[i] == tc.target {
							tc.collection[i] = tc.collection[i] + 1
							if tc.collection[i] == 0 { // wrapped around
								tc.collection[i] = 1
							}
						}
					}
				}
			}

			got := ContainsUint8x64(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsUint8x64() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsUint16x32(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []uint16
		target     uint16
		expected   bool
	}{
		{"empty", []uint16{}, 42, false},
		{"single found", []uint16{42}, 42, true},
		{"single not found", []uint16{42}, 10, false},
		{"small found", []uint16{1, 2, 3, 4, 5}, 3, true},
		{"small not found", []uint16{1, 2, 3, 4, 5}, 10, false},
		{"exactly 32 found", make([]uint16, 32), 0, true},
		{"exactly 32 not found", make([]uint16, 32), 1000, false},
		{"large found", make([]uint16, 1000), 0, true},
		{"large not found", make([]uint16, 100), 1000, false},
		{"max value", []uint16{65535, 1000, 500}, 65535, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 32 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[50] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = uint16(rand.IntN(65536))
						if tc.collection[i] == tc.target {
							tc.collection[i] = tc.collection[i] + 1
							if tc.collection[i] == 0 { // wrapped around
								tc.collection[i] = 1
							}
						}
					}
				}
			}

			got := ContainsUint16x32(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsUint16x32() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsUint32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []uint32
		target     uint32
		expected   bool
	}{
		{"empty", []uint32{}, 42, false},
		{"single found", []uint32{42}, 42, true},
		{"single not found", []uint32{42}, 10, false},
		{"small found", []uint32{1, 2, 3, 4, 5}, 3, true},
		{"small not found", []uint32{1, 2, 3, 4, 5}, 10, false},
		{"exactly 16 found", make([]uint32, 16), 0, true},
		{"exactly 16 not found", make([]uint32, 16), 100, false},
		{"large found", make([]uint32, 1000), 0, true},
		{"large not found", make([]uint32, 100), 100000, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 16 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[50] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = rand.Uint32()
						if tc.collection[i] == tc.target {
							tc.collection[i]++
						}
					}
				}
			}

			got := ContainsUint32x16(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsUint32x16() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsUint64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []uint64
		target     uint64
		expected   bool
	}{
		{"empty", []uint64{}, 42, false},
		{"single found", []uint64{42}, 42, true},
		{"single not found", []uint64{42}, 10, false},
		{"small found", []uint64{1, 2, 3, 4, 5}, 3, true},
		{"small not found", []uint64{1, 2, 3, 4, 5}, 10, false},
		{"exactly 8 found", make([]uint64, 8), 0, true},
		{"exactly 8 not found", make([]uint64, 8), 100, false},
		{"large found", make([]uint64, 1000), 0, true},
		{"large not found", make([]uint64, 100), 1000000, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 8 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[50] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = rand.Uint64()
						if tc.collection[i] == tc.target {
							tc.collection[i]++
						}
					}
				}
			}

			got := ContainsUint64x8(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsUint64x8() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsFloat32x16(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []float32
		target     float32
		expected   bool
	}{
		{"empty", []float32{}, 42.5, false},
		{"single found", []float32{42.5}, 42.5, true},
		{"single not found", []float32{42.5}, 10.0, false},
		{"small found", []float32{1.1, 2.2, 3.3, 4.4, 5.5}, 3.3, true},
		{"small not found", []float32{1.1, 2.2, 3.3, 4.4, 5.5}, 10.0, false},
		{"exactly 16 found", []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}, 16.0, true},
		{"exactly 16 not found", []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}, 100.0, false},
		{"large found", make([]float32, 1000), 0, true},
		{"large not found", make([]float32, 100), 100000.0, false},
		{"negative found", []float32{-1.1, -2.2, 3.3, 4.4}, -2.2, true},
		{"zeros", []float32{0, 0, 0, 0}, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 16 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[50] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = rand.Float32() * 10000
					}
				}
			}

			got := ContainsFloat32x16(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsFloat32x16() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestAVX512ContainsFloat64x8(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name      string
		collection []float64
		target     float64
		expected   bool
	}{
		{"empty", []float64{}, 42.5, false},
		{"single found", []float64{42.5}, 42.5, true},
		{"single not found", []float64{42.5}, 10.0, false},
		{"small found", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 3.3, true},
		{"small not found", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 10.0, false},
		{"exactly 8 found", []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}, 8.0, true},
		{"exactly 8 not found", []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}, 100.0, false},
		{"large found", make([]float64, 1000), 0, true},
		{"large not found", make([]float64, 100), 100000.0, false},
		{"negative found", []float64{-1.1, -2.2, 3.3, 4.4}, -2.2, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.collection) > 8 && tc.collection[0] == 0 {
				if tc.expected {
					tc.collection[50] = tc.target
				} else {
					for i := range tc.collection {
						tc.collection[i] = rand.Float64() * 10000
					}
				}
			}

			got := ContainsFloat64x8(tc.collection, tc.target)
			if got != tc.expected {
				t.Errorf("ContainsFloat64x8() = %v, want %v", got, tc.expected)
			}
		})
	}
}

// Test type aliases work correctly
func TestAVX512ContainsTypeAlias(t *testing.T) {
	requireAVX512(t)
	input := []myInt8{1, 2, 3, 4, 5}
	target := myInt8(3)

	got := ContainsInt8x64(input, target)
	if !got {
		t.Errorf("ContainsInt8x64() with type alias = false, want true")
	}

	target = myInt8(10)
	got = ContainsInt8x64(input, target)
	if got {
		t.Errorf("ContainsInt8x64() with type alias = true, want false")
	}
}

// Test consistency with lo.Contains
func TestAVX512ContainsConsistency(t *testing.T) {
	requireAVX512(t)
	testCases := []struct {
		name string
		int8 []int8
		int16 []int16
		int32 []int32
		int64 []int64
		uint8 []uint8
		uint16 []uint16
		uint32 []uint32
		uint64 []uint64
		float32 []float32
		float64 []float64
	}{
		{
			name: "mixed",
			int8: []int8{-1, 0, 1, 2, 3},
			int16: []int16{-1, 0, 1, 2, 3},
			int32: []int32{-1, 0, 1, 2, 3},
			int64: []int64{-1, 0, 1, 2, 3},
			uint8: []uint8{0, 1, 2, 3, 4},
			uint16: []uint16{0, 1, 2, 3, 4},
			uint32: []uint32{0, 1, 2, 3, 4},
			uint64: []uint64{0, 1, 2, 3, 4},
			float32: []float32{0.0, 1.1, 2.2, 3.3},
			float64: []float64{0.0, 1.1, 2.2, 3.3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.int8) > 0 {
				for _, target := range tc.int8 {
					got := ContainsInt8x64(tc.int8, target)
					want := lo.Contains(tc.int8, target)
					if got != want {
						t.Errorf("ContainsInt8x64() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.int16) > 0 {
				for _, target := range tc.int16 {
					got := ContainsInt16x32(tc.int16, target)
					want := lo.Contains(tc.int16, target)
					if got != want {
						t.Errorf("ContainsInt16x32() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.int32) > 0 {
				for _, target := range tc.int32 {
					got := ContainsInt32x16(tc.int32, target)
					want := lo.Contains(tc.int32, target)
					if got != want {
						t.Errorf("ContainsInt32x16() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.int64) > 0 {
				for _, target := range tc.int64 {
					got := ContainsInt64x8(tc.int64, target)
					want := lo.Contains(tc.int64, target)
					if got != want {
						t.Errorf("ContainsInt64x8() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.uint8) > 0 {
				for _, target := range tc.uint8 {
					got := ContainsUint8x64(tc.uint8, target)
					want := lo.Contains(tc.uint8, target)
					if got != want {
						t.Errorf("ContainsUint8x64() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.uint16) > 0 {
				for _, target := range tc.uint16 {
					got := ContainsUint16x32(tc.uint16, target)
					want := lo.Contains(tc.uint16, target)
					if got != want {
						t.Errorf("ContainsUint16x32() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.uint32) > 0 {
				for _, target := range tc.uint32 {
					got := ContainsUint32x16(tc.uint32, target)
					want := lo.Contains(tc.uint32, target)
					if got != want {
						t.Errorf("ContainsUint32x16() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.uint64) > 0 {
				for _, target := range tc.uint64 {
					got := ContainsUint64x8(tc.uint64, target)
					want := lo.Contains(tc.uint64, target)
					if got != want {
						t.Errorf("ContainsUint64x8() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.float32) > 0 {
				for _, target := range tc.float32 {
					got := ContainsFloat32x16(tc.float32, target)
					want := lo.Contains(tc.float32, target)
					if got != want {
						t.Errorf("ContainsFloat32x16() not consistent with lo.Contains for target %v", target)
					}
				}
			}
			if len(tc.float64) > 0 {
				for _, target := range tc.float64 {
					got := ContainsFloat64x8(tc.float64, target)
					want := lo.Contains(tc.float64, target)
					if got != want {
						t.Errorf("ContainsFloat64x8() not consistent with lo.Contains for target %v", target)
					}
				}
			}
		})
	}
}
