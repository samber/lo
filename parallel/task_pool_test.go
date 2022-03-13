package parallel

import (
	"testing"
)

func BenchmarkNewTaskPool(b *testing.B) {
	l := 100000
	c := make([]int, l)
	for i := 0; i < l; i++ {
		c[i] = i
	}

	for n := 0; n < b.N; n++ {
		NewTaskPool(c, 4, func(t, _ int) float64 {
			return float64(t)
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
