package parallel

import (
	"github.com/samber/lo/optional"
	"runtime"
	"sync"
)

var DefaultPoolSize = runtime.NumCPU() / 2

type Options struct {
	PoolSize optional.Value[int]
}

func getPoolSize(options []*Options) int {
	if options != nil && len(options) >= 0 {
		return options[0].PoolSize.
			IfPresent().
			OrElse(DefaultPoolSize)
	}

	return DefaultPoolSize
}

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is called in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(T, int) R, options ...*Options) []R {
	result := make([]R, len(collection))
	TaskPool[T](len(collection), getPoolSize(options), func(i int) {
		result[i] = iteratee(collection[i], i)
	})
	return result
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is called in parallel.
func ForEach[T any](collection []T, iteratee func(T, int), options ...*Options) {
	TaskPool[T](len(collection), getPoolSize(options), func(i int) {
		iteratee(collection[i], i)
	})
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
// `iteratee` is called in parallel.
func Times[T any](count int, iteratee func(int) T, options ...*Options) []T {
	result := make([]T, count)
	TaskPool[int](count, getPoolSize(options), func(i int) {
		result[i] = iteratee(i)
	})
	return result
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
// `iteratee` is call in parallel.
func GroupBy[T any, U comparable](collection []T, iteratee func(T) U, options ...*Options) map[U][]T {
	result := map[U][]T{}
	resultMu := &sync.Mutex{}

	TaskPool[T](len(collection), getPoolSize(options), func(i int) {
		key := iteratee(collection[i])

		resultMu.Lock()
		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}
		resultMu.Unlock()

		resultMu.Lock()
		result[key] = append(result[key], collection[i])
		resultMu.Unlock()
	})
	return result
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// `iteratee` is call in parallel.
func PartitionBy[T any, K comparable](collection []T, iteratee func(x T) K, options ...*Options) [][]T {
	result := [][]T{}
	groups := GroupBy(collection, iteratee, options...)

	for _, v := range groups {
		result = append(result, v)
	}

	return result
}
