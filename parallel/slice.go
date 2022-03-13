package parallel

import (
	"runtime"
	"sync"
)

var DefaultPoolSize = runtime.NumCPU() / 2

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is called in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))
	NewTaskPool(collection, DefaultPoolSize, func(v T, i int) {
		result[i] = iteratee(v, i)
	})
	return result
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is called in parallel.
func ForEach[T any](collection []T, iteratee func(T, int)) {
	NewTaskPool(collection, DefaultPoolSize, iteratee)
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
// `iteratee` is called in parallel.
func Times[T any](count int, iteratee func(int) T) []T {
	result := make([]T, count)
	NewTaskPool[int](make([]int, count), DefaultPoolSize, func(_ int, i int) {
		result[i] = iteratee(i)
	})
	return result
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
// `iteratee` is call in parallel.
func GroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := map[U][]T{}
	resultMu := &sync.Mutex{}

	NewTaskPool[T](collection, DefaultPoolSize, func(v T, _ int) {
		key := iteratee(v)

		resultMu.Lock()
		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}
		resultMu.Unlock()

		resultMu.Lock()
		result[key] = append(result[key], v)
		resultMu.Unlock()
	})
	return result
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// `iteratee` is call in parallel.
func PartitionBy[T any, K comparable](collection []T, iteratee func(x T) K) [][]T {
	result := [][]T{}
	seen := map[K]int{}

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(collection))

	for _, item := range collection {
		go func(_item T) {
			key := iteratee(_item)

			mu.Lock()

			resultIndex, ok := seen[key]
			if !ok {
				resultIndex = len(result)
				seen[key] = resultIndex
				result = append(result, []T{})
			}

			result[resultIndex] = append(result[resultIndex], _item)

			mu.Unlock()
			wg.Done()
		}(item)
	}

	wg.Wait()

	return result
}
