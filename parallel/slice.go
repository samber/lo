package parallel

import "sync"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func (_item T, _i int) {
			res := iteratee(_item, _i)

			mu.Lock()
			result[_i] = res
			mu.Unlock()

			wg.Done()
		}(item, i)
	}

	wg.Wait()

	return result
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is call in parallel.
func ForEach[T any](collection []T, iteratee func(T, int)) {
	var wg sync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func (_item T, _i int) {
			iteratee(_item, _i)
			wg.Done()
		}(item, i)
	}

	wg.Wait()
}
