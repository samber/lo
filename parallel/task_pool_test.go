package parallel

import (
	"runtime"
	"testing"
)

func BenchmarkNewTaskPool(b *testing.B) {
	l := 100000
	c := make([]int, l)
	for i := 0; i < l; i++ {
		c[i] = i
	}

	result := make([]float64, l)

	for n := 0; n < b.N; n++ {
		NewTaskPool[int](c, runtime.NumCPU()/2, func(v, i int) {
			result[i] = float64(v)
		})
	}
}

func BenchmarkMap(b *testing.B) {
	l := 100000
	c := make([]int, l)
	for i := 0; i < l; i++ {
		c[i] = i
	}

	for n := 0; n < b.N; n++ {
		Map[int, float64](c, func(t, _ int) float64 {
			return float64(t)
		})
	}
}
