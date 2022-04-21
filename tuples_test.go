package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestT(t *testing.T) {
	is := assert.New(t)

	r1 := T2[string, int]("a", 1)
	r2 := T3[string, int, float32]("b", 2, 3.0)
	r3 := T4[string, int, float32, bool]("c", 3, 4.0, true)
	r4 := T5[string, int, float32, bool, string]("d", 4, 5.0, false, "e")
	r5 := T6[string, int, float32, bool, string, int]("f", 5, 6.0, true, "g", 7)
	r6 := T7[string, int, float32, bool, string, int, float64]("h", 6, 7.0, false, "i", 8, 9.0)
	r7 := T8[string, int, float32, bool, string, int, float64, bool]("j", 7, 8.0, true, "k", 9, 10.0, false)
	r8 := T9[string, int, float32, bool, string, int, float64, bool, string]("l", 8, 9.0, false, "m", 10, 11.0, true, "n")

	is.Equal(r1, Tuple2[string, int]{A: "a", B: 1})
	is.Equal(r2, Tuple3[string, int, float32]{A: "b", B: 2, C: 3.0})
	is.Equal(r3, Tuple4[string, int, float32, bool]{A: "c", B: 3, C: 4.0, D: true})
	is.Equal(r4, Tuple5[string, int, float32, bool, string]{A: "d", B: 4, C: 5.0, D: false, E: "e"})
	is.Equal(r5, Tuple6[string, int, float32, bool, string, int]{A: "f", B: 5, C: 6.0, D: true, E: "g", F: 7})
	is.Equal(r6, Tuple7[string, int, float32, bool, string, int, float64]{A: "h", B: 6, C: 7.0, D: false, E: "i", F: 8, G: 9.0})
	is.Equal(r7, Tuple8[string, int, float32, bool, string, int, float64, bool]{A: "j", B: 7, C: 8.0, D: true, E: "k", F: 9, G: 10.0, H: false})
	is.Equal(r8, Tuple9[string, int, float32, bool, string, int, float64, bool, string]{A: "l", B: 8, C: 9.0, D: false, E: "m", F: 10, G: 11.0, H: true, I: "n"})
}

func TestZip(t *testing.T) {
	is := assert.New(t)

	r1 := Zip2[string, int](
		[]string{"a", "b"},
		[]int{1, 2},
	)

	r2 := Zip3[string, int, int](
		[]string{"a", "b", "c"},
		[]int{1, 2, 3}, []int{4, 5, 6},
	)

	r3 := Zip4[string, int, int, bool](
		[]string{"a", "b", "c", "d"},
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]bool{true, true, true, true},
	)

	r4 := Zip5[string, int, int, bool, float32](
		[]string{"a", "b", "c", "d", "e"},
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]bool{true, true, true, true, true},
		[]float32{0.1, 0.2, 0.3, 0.4, 0.5},
	)

	r5 := Zip6[string, int, int, bool, float32, float64](
		[]string{"a", "b", "c", "d", "e", "f"},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{7, 8, 9, 10, 11, 12},
		[]bool{true, true, true, true, true, true},
		[]float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6},
		[]float64{0.01, 0.02, 0.03, 0.04, 0.05, 0.06},
	)

	r6 := Zip7[string, int, int, bool, float32, float64, int8](
		[]string{"a", "b", "c", "d", "e", "f", "g"},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{8, 9, 10, 11, 12, 13, 14},
		[]bool{true, true, true, true, true, true, true},
		[]float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7},
		[]float64{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07},
		[]int8{1, 2, 3, 4, 5, 6, 7},
	)

	r7 := Zip8[string, int, int, bool, float32, float64, int8, int16](
		[]string{"a", "b", "c", "d", "e", "f", "g", "h"},
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]int{9, 10, 11, 12, 13, 14, 15, 16},
		[]bool{true, true, true, true, true, true, true, true},
		[]float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8},
		[]float64{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08},
		[]int8{1, 2, 3, 4, 5, 6, 7, 8},
		[]int16{1, 2, 3, 4, 5, 6, 7, 8},
	)

	r8 := Zip9[string, int, int, bool, float32, float64, int8, int16, int32](
		[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{10, 11, 12, 13, 14, 15, 16, 17, 18},
		[]bool{true, true, true, true, true, true, true, true, true},
		[]float32{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9},
		[]float64{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09},
		[]int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int16{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
	)

	is.Equal(r1, []Tuple2[string, int]{
		{A: "a", B: 1},
		{A: "b", B: 2},
	})

	is.Equal(r2, []Tuple3[string, int, int]{
		{A: "a", B: 1, C: 4},
		{A: "b", B: 2, C: 5},
		{A: "c", B: 3, C: 6},
	})

	is.Equal(r3, []Tuple4[string, int, int, bool]{
		{A: "a", B: 1, C: 5, D: true},
		{A: "b", B: 2, C: 6, D: true},
		{A: "c", B: 3, C: 7, D: true},
		{A: "d", B: 4, C: 8, D: true},
	})

	is.Equal(r4, []Tuple5[string, int, int, bool, float32]{
		{A: "a", B: 1, C: 6, D: true, E: 0.1},
		{A: "b", B: 2, C: 7, D: true, E: 0.2},
		{A: "c", B: 3, C: 8, D: true, E: 0.3},
		{A: "d", B: 4, C: 9, D: true, E: 0.4},
		{A: "e", B: 5, C: 10, D: true, E: 0.5},
	})

	is.Equal(r5, []Tuple6[string, int, int, bool, float32, float64]{
		{A: "a", B: 1, C: 7, D: true, E: 0.1, F: 0.01},
		{A: "b", B: 2, C: 8, D: true, E: 0.2, F: 0.02},
		{A: "c", B: 3, C: 9, D: true, E: 0.3, F: 0.03},
		{A: "d", B: 4, C: 10, D: true, E: 0.4, F: 0.04},
		{A: "e", B: 5, C: 11, D: true, E: 0.5, F: 0.05},
		{A: "f", B: 6, C: 12, D: true, E: 0.6, F: 0.06},
	})

	is.Equal(r6, []Tuple7[string, int, int, bool, float32, float64, int8]{
		{A: "a", B: 1, C: 8, D: true, E: 0.1, F: 0.01, G: 1},
		{A: "b", B: 2, C: 9, D: true, E: 0.2, F: 0.02, G: 2},
		{A: "c", B: 3, C: 10, D: true, E: 0.3, F: 0.03, G: 3},
		{A: "d", B: 4, C: 11, D: true, E: 0.4, F: 0.04, G: 4},
		{A: "e", B: 5, C: 12, D: true, E: 0.5, F: 0.05, G: 5},
		{A: "f", B: 6, C: 13, D: true, E: 0.6, F: 0.06, G: 6},
		{A: "g", B: 7, C: 14, D: true, E: 0.7, F: 0.07, G: 7},
	})

	is.Equal(r7, []Tuple8[string, int, int, bool, float32, float64, int8, int16]{
		{A: "a", B: 1, C: 9, D: true, E: 0.1, F: 0.01, G: 1, H: 1},
		{A: "b", B: 2, C: 10, D: true, E: 0.2, F: 0.02, G: 2, H: 2},
		{A: "c", B: 3, C: 11, D: true, E: 0.3, F: 0.03, G: 3, H: 3},
		{A: "d", B: 4, C: 12, D: true, E: 0.4, F: 0.04, G: 4, H: 4},
		{A: "e", B: 5, C: 13, D: true, E: 0.5, F: 0.05, G: 5, H: 5},
		{A: "f", B: 6, C: 14, D: true, E: 0.6, F: 0.06, G: 6, H: 6},
		{A: "g", B: 7, C: 15, D: true, E: 0.7, F: 0.07, G: 7, H: 7},
		{A: "h", B: 8, C: 16, D: true, E: 0.8, F: 0.08, G: 8, H: 8},
	})

	is.Equal(r8, []Tuple9[string, int, int, bool, float32, float64, int8, int16, int32]{
		{A: "a", B: 1, C: 10, D: true, E: 0.1, F: 0.01, G: 1, H: 1, I: 1},
		{A: "b", B: 2, C: 11, D: true, E: 0.2, F: 0.02, G: 2, H: 2, I: 2},
		{A: "c", B: 3, C: 12, D: true, E: 0.3, F: 0.03, G: 3, H: 3, I: 3},
		{A: "d", B: 4, C: 13, D: true, E: 0.4, F: 0.04, G: 4, H: 4, I: 4},
		{A: "e", B: 5, C: 14, D: true, E: 0.5, F: 0.05, G: 5, H: 5, I: 5},
		{A: "f", B: 6, C: 15, D: true, E: 0.6, F: 0.06, G: 6, H: 6, I: 6},
		{A: "g", B: 7, C: 16, D: true, E: 0.7, F: 0.07, G: 7, H: 7, I: 7},
		{A: "h", B: 8, C: 17, D: true, E: 0.8, F: 0.08, G: 8, H: 8, I: 8},
		{A: "i", B: 9, C: 18, D: true, E: 0.9, F: 0.09, G: 9, H: 9, I: 9},
	})
}

func TestUnzip(t *testing.T) {
	is := assert.New(t)

	r1, r2 := Unzip2[string, int]([]Tuple2[string, int]{{A: "a", B: 1}, {A: "b", B: 2}})

	is.Equal(r1, []string{"a", "b"})
	is.Equal(r2, []int{1, 2})
}
