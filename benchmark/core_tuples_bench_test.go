package benchmark

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
)

func BenchmarkZip2_Equal(b *testing.B) {
	for _, n := range lengths {
		a := genSliceInt(n)
		s := genSliceString(n)
		b.Run(fmt.Sprintf("n_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lo.Zip2(a, s)
			}
		})
	}
}

func BenchmarkZip2_Unequal(b *testing.B) {
	for _, n := range lengths {
		a := genSliceInt(n)
		s := genSliceString(n / 2)
		b.Run(fmt.Sprintf("n_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lo.Zip2(a, s)
			}
		})
	}
}

func BenchmarkUnzip2(b *testing.B) {
	for _, n := range lengths {
		tuples := make([]lo.Tuple2[int, string], n)
		for i := range tuples {
			tuples[i] = lo.Tuple2[int, string]{A: i, B: "x"}
		}
		b.Run(fmt.Sprintf("n_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lo.Unzip2(tuples)
			}
		})
	}
}
