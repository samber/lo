package parallel

import (
	"runtime"
	"sync"
	"testing"
)

const l = 100000

var c []int

func init() {
	c = make([]int, l)
	for i := 0; i < l; i++ {
		c[i] = i
	}
}

func BenchmarkNewTaskPool(b *testing.B) {
	result := make([]float64, l)

	for n := 0; n < b.N; n++ {
		NewTaskPool[int](c, runtime.NumCPU()/2, func(v, i int) {
			result[i] = float64(v)
		})
	}
}

func oldMap[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	var wg sync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func(_item T, _i int) {
			res := iteratee(_item, _i)

			result[_i] = res

			wg.Done()
		}(item, i)
	}

	wg.Wait()

	return result
}

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		oldMap[int, float64](c, func(t, _ int) float64 {
			return float64(t)
		})
	}
}
