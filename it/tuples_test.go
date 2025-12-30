//go:build go1.23

package it

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/samber/lo"
)

func TestZip(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Zip2(
		values("a", "b"),
		values(1, 2),
	)

	r2 := Zip3(
		values("a", "b", "c"),
		values(1, 2, 3),
		values(4, 5, 6),
	)

	r3 := Zip4(
		values("a", "b", "c", "d"),
		values(1, 2, 3, 4),
		values(5, 6, 7, 8),
		values(true, true, true, true),
	)

	r4 := Zip5(
		values("a", "b", "c", "d", "e"),
		values(1, 2, 3, 4, 5),
		values(6, 7, 8, 9, 10),
		values(true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5),
	)

	r5 := Zip6(
		values("a", "b", "c", "d", "e", "f"),
		values(1, 2, 3, 4, 5, 6),
		values(7, 8, 9, 10, 11, 12),
		values(true, true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5, 0.6),
		values(0.01, 0.02, 0.03, 0.04, 0.05, 0.06),
	)

	r6 := Zip7(
		values("a", "b", "c", "d", "e", "f", "g"),
		values(1, 2, 3, 4, 5, 6, 7),
		values(8, 9, 10, 11, 12, 13, 14),
		values(true, true, true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7),
		values(0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07),
		values[int8](1, 2, 3, 4, 5, 6, 7),
	)

	r7 := Zip8(
		values("a", "b", "c", "d", "e", "f", "g", "h"),
		values(1, 2, 3, 4, 5, 6, 7, 8),
		values(9, 10, 11, 12, 13, 14, 15, 16),
		values(true, true, true, true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8),
		values(0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08),
		values[int8](1, 2, 3, 4, 5, 6, 7, 8),
		values[int16](1, 2, 3, 4, 5, 6, 7, 8),
	)

	r8 := Zip9(
		values("a", "b", "c", "d", "e", "f", "g", "h", "i"),
		values(1, 2, 3, 4, 5, 6, 7, 8, 9),
		values(10, 11, 12, 13, 14, 15, 16, 17, 18),
		values(true, true, true, true, true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9),
		values(0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09),
		values[int8](1, 2, 3, 4, 5, 6, 7, 8, 9),
		values[int16](1, 2, 3, 4, 5, 6, 7, 8, 9),
		values[int32](1, 2, 3, 4, 5, 6, 7, 8, 9),
	)

	assertSeqSupportBreak(t, r1)
	assertSeqSupportBreak(t, r2)
	assertSeqSupportBreak(t, r3)
	assertSeqSupportBreak(t, r4)
	assertSeqSupportBreak(t, r5)
	assertSeqSupportBreak(t, r6)
	assertSeqSupportBreak(t, r7)
	assertSeqSupportBreak(t, r8)

	is.Equal([]lo.Tuple2[string, int]{
		{A: "a", B: 1},
		{A: "b", B: 2},
	}, slices.Collect(r1))

	is.Equal([]lo.Tuple3[string, int, int]{
		{A: "a", B: 1, C: 4},
		{A: "b", B: 2, C: 5},
		{A: "c", B: 3, C: 6},
	}, slices.Collect(r2))

	is.Equal([]lo.Tuple4[string, int, int, bool]{
		{A: "a", B: 1, C: 5, D: true},
		{A: "b", B: 2, C: 6, D: true},
		{A: "c", B: 3, C: 7, D: true},
		{A: "d", B: 4, C: 8, D: true},
	}, slices.Collect(r3))

	is.Equal([]lo.Tuple5[string, int, int, bool, float32]{
		{A: "a", B: 1, C: 6, D: true, E: 0.1},
		{A: "b", B: 2, C: 7, D: true, E: 0.2},
		{A: "c", B: 3, C: 8, D: true, E: 0.3},
		{A: "d", B: 4, C: 9, D: true, E: 0.4},
		{A: "e", B: 5, C: 10, D: true, E: 0.5},
	}, slices.Collect(r4))

	is.Equal([]lo.Tuple6[string, int, int, bool, float32, float64]{
		{A: "a", B: 1, C: 7, D: true, E: 0.1, F: 0.01},
		{A: "b", B: 2, C: 8, D: true, E: 0.2, F: 0.02},
		{A: "c", B: 3, C: 9, D: true, E: 0.3, F: 0.03},
		{A: "d", B: 4, C: 10, D: true, E: 0.4, F: 0.04},
		{A: "e", B: 5, C: 11, D: true, E: 0.5, F: 0.05},
		{A: "f", B: 6, C: 12, D: true, E: 0.6, F: 0.06},
	}, slices.Collect(r5))

	is.Equal([]lo.Tuple7[string, int, int, bool, float32, float64, int8]{
		{A: "a", B: 1, C: 8, D: true, E: 0.1, F: 0.01, G: 1},
		{A: "b", B: 2, C: 9, D: true, E: 0.2, F: 0.02, G: 2},
		{A: "c", B: 3, C: 10, D: true, E: 0.3, F: 0.03, G: 3},
		{A: "d", B: 4, C: 11, D: true, E: 0.4, F: 0.04, G: 4},
		{A: "e", B: 5, C: 12, D: true, E: 0.5, F: 0.05, G: 5},
		{A: "f", B: 6, C: 13, D: true, E: 0.6, F: 0.06, G: 6},
		{A: "g", B: 7, C: 14, D: true, E: 0.7, F: 0.07, G: 7},
	}, slices.Collect(r6))

	is.Equal([]lo.Tuple8[string, int, int, bool, float32, float64, int8, int16]{
		{A: "a", B: 1, C: 9, D: true, E: 0.1, F: 0.01, G: 1, H: 1},
		{A: "b", B: 2, C: 10, D: true, E: 0.2, F: 0.02, G: 2, H: 2},
		{A: "c", B: 3, C: 11, D: true, E: 0.3, F: 0.03, G: 3, H: 3},
		{A: "d", B: 4, C: 12, D: true, E: 0.4, F: 0.04, G: 4, H: 4},
		{A: "e", B: 5, C: 13, D: true, E: 0.5, F: 0.05, G: 5, H: 5},
		{A: "f", B: 6, C: 14, D: true, E: 0.6, F: 0.06, G: 6, H: 6},
		{A: "g", B: 7, C: 15, D: true, E: 0.7, F: 0.07, G: 7, H: 7},
		{A: "h", B: 8, C: 16, D: true, E: 0.8, F: 0.08, G: 8, H: 8},
	}, slices.Collect(r7))

	is.Equal([]lo.Tuple9[string, int, int, bool, float32, float64, int8, int16, int32]{
		{A: "a", B: 1, C: 10, D: true, E: 0.1, F: 0.01, G: 1, H: 1, I: 1},
		{A: "b", B: 2, C: 11, D: true, E: 0.2, F: 0.02, G: 2, H: 2, I: 2},
		{A: "c", B: 3, C: 12, D: true, E: 0.3, F: 0.03, G: 3, H: 3, I: 3},
		{A: "d", B: 4, C: 13, D: true, E: 0.4, F: 0.04, G: 4, H: 4, I: 4},
		{A: "e", B: 5, C: 14, D: true, E: 0.5, F: 0.05, G: 5, H: 5, I: 5},
		{A: "f", B: 6, C: 15, D: true, E: 0.6, F: 0.06, G: 6, H: 6, I: 6},
		{A: "g", B: 7, C: 16, D: true, E: 0.7, F: 0.07, G: 7, H: 7, I: 7},
		{A: "h", B: 8, C: 17, D: true, E: 0.8, F: 0.08, G: 8, H: 8, I: 8},
		{A: "i", B: 9, C: 18, D: true, E: 0.9, F: 0.09, G: 9, H: 9, I: 9},
	}, slices.Collect(r8))
}

func TestZipBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := ZipBy2(
		values("a", "b"),
		values(1, 2),
		lo.T2[string, int],
	)

	r2 := ZipBy3(
		values("a", "b", "c"),
		values(1, 2, 3),
		values(4, 5, 6),
		lo.T3[string, int, int],
	)

	r3 := ZipBy4(
		values("a", "b", "c", "d"),
		values(1, 2, 3, 4),
		values(5, 6, 7, 8),
		values(true, true, true, true),
		lo.T4[string, int, int, bool],
	)

	r4 := ZipBy5(
		values("a", "b", "c", "d", "e"),
		values(1, 2, 3, 4, 5),
		values(6, 7, 8, 9, 10),
		values(true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5),
		lo.T5[string, int, int, bool, float32],
	)

	r5 := ZipBy6(
		values("a", "b", "c", "d", "e", "f"),
		values(1, 2, 3, 4, 5, 6),
		values(7, 8, 9, 10, 11, 12),
		values(true, true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5, 0.6),
		values(0.01, 0.02, 0.03, 0.04, 0.05, 0.06),
		lo.T6[string, int, int, bool, float32, float64],
	)

	r6 := ZipBy7(
		values("a", "b", "c", "d", "e", "f", "g"),
		values(1, 2, 3, 4, 5, 6, 7),
		values(8, 9, 10, 11, 12, 13, 14),
		values(true, true, true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7),
		values(0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07),
		values[int8](1, 2, 3, 4, 5, 6, 7),
		lo.T7[string, int, int, bool, float32, float64, int8],
	)

	r7 := ZipBy8(
		values("a", "b", "c", "d", "e", "f", "g", "h"),
		values(1, 2, 3, 4, 5, 6, 7, 8),
		values(9, 10, 11, 12, 13, 14, 15, 16),
		values(true, true, true, true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8),
		values(0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08),
		values[int8](1, 2, 3, 4, 5, 6, 7, 8),
		values[int16](1, 2, 3, 4, 5, 6, 7, 8),
		lo.T8[string, int, int, bool, float32, float64, int8, int16],
	)

	r8 := ZipBy9(
		values("a", "b", "c", "d", "e", "f", "g", "h", "i"),
		values(1, 2, 3, 4, 5, 6, 7, 8, 9),
		values(10, 11, 12, 13, 14, 15, 16, 17, 18),
		values(true, true, true, true, true, true, true, true, true),
		values[float32](0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9),
		values(0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09),
		values[int8](1, 2, 3, 4, 5, 6, 7, 8, 9),
		values[int16](1, 2, 3, 4, 5, 6, 7, 8, 9),
		values[int32](1, 2, 3, 4, 5, 6, 7, 8, 9),
		lo.T9[string, int, int, bool, float32, float64, int8, int16, int32],
	)

	assertSeqSupportBreak(t, r1)
	assertSeqSupportBreak(t, r2)
	assertSeqSupportBreak(t, r3)
	assertSeqSupportBreak(t, r4)
	assertSeqSupportBreak(t, r5)
	assertSeqSupportBreak(t, r6)
	assertSeqSupportBreak(t, r7)
	assertSeqSupportBreak(t, r8)

	is.Equal([]lo.Tuple2[string, int]{
		{A: "a", B: 1},
		{A: "b", B: 2},
	}, slices.Collect(r1))

	is.Equal([]lo.Tuple3[string, int, int]{
		{A: "a", B: 1, C: 4},
		{A: "b", B: 2, C: 5},
		{A: "c", B: 3, C: 6},
	}, slices.Collect(r2))

	is.Equal([]lo.Tuple4[string, int, int, bool]{
		{A: "a", B: 1, C: 5, D: true},
		{A: "b", B: 2, C: 6, D: true},
		{A: "c", B: 3, C: 7, D: true},
		{A: "d", B: 4, C: 8, D: true},
	}, slices.Collect(r3))

	is.Equal([]lo.Tuple5[string, int, int, bool, float32]{
		{A: "a", B: 1, C: 6, D: true, E: 0.1},
		{A: "b", B: 2, C: 7, D: true, E: 0.2},
		{A: "c", B: 3, C: 8, D: true, E: 0.3},
		{A: "d", B: 4, C: 9, D: true, E: 0.4},
		{A: "e", B: 5, C: 10, D: true, E: 0.5},
	}, slices.Collect(r4))

	is.Equal([]lo.Tuple6[string, int, int, bool, float32, float64]{
		{A: "a", B: 1, C: 7, D: true, E: 0.1, F: 0.01},
		{A: "b", B: 2, C: 8, D: true, E: 0.2, F: 0.02},
		{A: "c", B: 3, C: 9, D: true, E: 0.3, F: 0.03},
		{A: "d", B: 4, C: 10, D: true, E: 0.4, F: 0.04},
		{A: "e", B: 5, C: 11, D: true, E: 0.5, F: 0.05},
		{A: "f", B: 6, C: 12, D: true, E: 0.6, F: 0.06},
	}, slices.Collect(r5))

	is.Equal([]lo.Tuple7[string, int, int, bool, float32, float64, int8]{
		{A: "a", B: 1, C: 8, D: true, E: 0.1, F: 0.01, G: 1},
		{A: "b", B: 2, C: 9, D: true, E: 0.2, F: 0.02, G: 2},
		{A: "c", B: 3, C: 10, D: true, E: 0.3, F: 0.03, G: 3},
		{A: "d", B: 4, C: 11, D: true, E: 0.4, F: 0.04, G: 4},
		{A: "e", B: 5, C: 12, D: true, E: 0.5, F: 0.05, G: 5},
		{A: "f", B: 6, C: 13, D: true, E: 0.6, F: 0.06, G: 6},
		{A: "g", B: 7, C: 14, D: true, E: 0.7, F: 0.07, G: 7},
	}, slices.Collect(r6))

	is.Equal([]lo.Tuple8[string, int, int, bool, float32, float64, int8, int16]{
		{A: "a", B: 1, C: 9, D: true, E: 0.1, F: 0.01, G: 1, H: 1},
		{A: "b", B: 2, C: 10, D: true, E: 0.2, F: 0.02, G: 2, H: 2},
		{A: "c", B: 3, C: 11, D: true, E: 0.3, F: 0.03, G: 3, H: 3},
		{A: "d", B: 4, C: 12, D: true, E: 0.4, F: 0.04, G: 4, H: 4},
		{A: "e", B: 5, C: 13, D: true, E: 0.5, F: 0.05, G: 5, H: 5},
		{A: "f", B: 6, C: 14, D: true, E: 0.6, F: 0.06, G: 6, H: 6},
		{A: "g", B: 7, C: 15, D: true, E: 0.7, F: 0.07, G: 7, H: 7},
		{A: "h", B: 8, C: 16, D: true, E: 0.8, F: 0.08, G: 8, H: 8},
	}, slices.Collect(r7))

	is.Equal([]lo.Tuple9[string, int, int, bool, float32, float64, int8, int16, int32]{
		{A: "a", B: 1, C: 10, D: true, E: 0.1, F: 0.01, G: 1, H: 1, I: 1},
		{A: "b", B: 2, C: 11, D: true, E: 0.2, F: 0.02, G: 2, H: 2, I: 2},
		{A: "c", B: 3, C: 12, D: true, E: 0.3, F: 0.03, G: 3, H: 3, I: 3},
		{A: "d", B: 4, C: 13, D: true, E: 0.4, F: 0.04, G: 4, H: 4, I: 4},
		{A: "e", B: 5, C: 14, D: true, E: 0.5, F: 0.05, G: 5, H: 5, I: 5},
		{A: "f", B: 6, C: 15, D: true, E: 0.6, F: 0.06, G: 6, H: 6, I: 6},
		{A: "g", B: 7, C: 16, D: true, E: 0.7, F: 0.07, G: 7, H: 7, I: 7},
		{A: "h", B: 8, C: 17, D: true, E: 0.8, F: 0.08, G: 8, H: 8, I: 8},
		{A: "i", B: 9, C: 18, D: true, E: 0.9, F: 0.09, G: 9, H: 9, I: 9},
	}, slices.Collect(r8))
}

func TestCrossJoin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	listOne := values("a", "b", "c")
	listTwo := values(1, 2, 3)
	emptyList := values[any]()
	mixedList := values[any](9.6, 4, "foobar")

	results1 := CrossJoin2(emptyList, listTwo)
	is.Empty(slices.Collect(results1))

	results2 := CrossJoin2(listOne, emptyList)
	is.Empty(slices.Collect(results2))

	results3 := CrossJoin2(emptyList, emptyList)
	is.Empty(slices.Collect(results3))

	results4 := CrossJoin2(values("a"), listTwo)
	is.Equal([]lo.Tuple2[string, int]{lo.T2("a", 1), lo.T2("a", 2), lo.T2("a", 3)}, slices.Collect(results4))

	results5 := CrossJoin2(listOne, values(1))
	is.Equal([]lo.Tuple2[string, int]{lo.T2("a", 1), lo.T2("b", 1), lo.T2("c", 1)}, slices.Collect(results5))

	results6 := CrossJoin2(listOne, listTwo)
	is.Equal([]lo.Tuple2[string, int]{lo.T2("a", 1), lo.T2("a", 2), lo.T2("a", 3), lo.T2("b", 1), lo.T2("b", 2), lo.T2("b", 3), lo.T2("c", 1), lo.T2("c", 2), lo.T2("c", 3)}, slices.Collect(results6))

	results7 := CrossJoin2(listOne, mixedList)
	is.Equal([]lo.Tuple2[string, any]{lo.T2[string, any]("a", 9.6), lo.T2[string, any]("a", 4), lo.T2[string, any]("a", "foobar"), lo.T2[string, any]("b", 9.6), lo.T2[string, any]("b", 4), lo.T2[string, any]("b", "foobar"), lo.T2[string, any]("c", 9.6), lo.T2[string, any]("c", 4), lo.T2[string, any]("c", "foobar")}, slices.Collect(results7))
}

func TestCrossJoinBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	listOne := values("a", "b", "c")
	listTwo := values(1, 2, 3)
	emptyList := values[any]()
	mixedList := values[any](9.6, 4, "foobar")

	results1 := CrossJoinBy2(emptyList, listTwo, lo.T2[any, int])
	is.Empty(slices.Collect(results1))

	results2 := CrossJoinBy2(listOne, emptyList, lo.T2[string, any])
	is.Empty(slices.Collect(results2))

	results3 := CrossJoinBy2(emptyList, emptyList, lo.T2[any, any])
	is.Empty(slices.Collect(results3))

	results4 := CrossJoinBy2(values("a"), listTwo, lo.T2[string, int])
	is.Equal([]lo.Tuple2[string, int]{lo.T2("a", 1), lo.T2("a", 2), lo.T2("a", 3)}, slices.Collect(results4))

	results5 := CrossJoinBy2(listOne, values(1), lo.T2[string, int])
	is.Equal([]lo.Tuple2[string, int]{lo.T2("a", 1), lo.T2("b", 1), lo.T2("c", 1)}, slices.Collect(results5))

	results6 := CrossJoinBy2(listOne, listTwo, lo.T2[string, int])
	is.Equal([]lo.Tuple2[string, int]{lo.T2("a", 1), lo.T2("a", 2), lo.T2("a", 3), lo.T2("b", 1), lo.T2("b", 2), lo.T2("b", 3), lo.T2("c", 1), lo.T2("c", 2), lo.T2("c", 3)}, slices.Collect(results6))

	results7 := CrossJoinBy2(listOne, mixedList, lo.T2[string, any])
	is.Equal([]lo.Tuple2[string, any]{lo.T2[string, any]("a", 9.6), lo.T2[string, any]("a", 4), lo.T2[string, any]("a", "foobar"), lo.T2[string, any]("b", 9.6), lo.T2[string, any]("b", 4), lo.T2[string, any]("b", "foobar"), lo.T2[string, any]("c", 9.6), lo.T2[string, any]("c", 4), lo.T2[string, any]("c", "foobar")}, slices.Collect(results7))
}
