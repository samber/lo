package parallel

import "sync"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is called in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(T, int) R) []R {
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

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is called in parallel.
func ForEach[T any](collection []T, iteratee func(T, int)) {
	var wg sync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func(_item T, _i int) {
			iteratee(_item, _i)
			wg.Done()
		}(item, i)
	}

	wg.Wait()
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
// `iteratee` is called in parallel.
func Times[T any](count int, iteratee func(int) T) []T {
	result := make([]T, count)

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(_i int) {
			item := iteratee(_i)

			result[_i] = item

			wg.Done()
		}(i)
	}

	wg.Wait()

	return result
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
// `iteratee` is called in parallel.
func GroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := map[U][]T{}

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(collection))

	for _, item := range collection {
		go func(_item T) {
			key := iteratee(_item)

			mu.Lock()

			result[key] = append(result[key], _item)

			mu.Unlock()
			wg.Done()
		}(item)
	}

	wg.Wait()

	return result
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// `iteratee` is called in parallel.
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

// Find search an element in a slice based on a predicate. It returns element and true if element was found.
// `predicate` is called in parallel. No order guaranteed, i.e. found element may be not first matching the predicate.
func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	found := make(chan T, len(collection))
	finished := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(len(collection))
	go func() {
		wg.Wait()
		finished <- struct{}{}
	}()

	for _, item := range collection {
		go func(item T) {
			if predicate(item) {
				found <- item
			}

			wg.Done()
		}(item)
	}

	var result T
	select {
	case result = <-found:
		return result, true
	case <-finished:
		return result, false
	}
}
