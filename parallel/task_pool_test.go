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

func normalMap[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

var callback = func(t, _ int) float64 {
	return float64(t)
}

func BenchmarkNormalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		normalMap[int, float64](c, callback)
	}
}

func BenchmarkNewTaskPool(b *testing.B) {
	result := make([]float64, l)

	for n := 0; n < b.N; n++ {
		TaskPool[int](len(c), runtime.NumCPU(), func(i int) {
			result[i] = float64(c[i])
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
		oldMap[int, float64](c, callback)
	}
}
