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
