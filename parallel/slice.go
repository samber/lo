package parallel

import "sync"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
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

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
// `callback` is call in parallel. Result is not guaranteed to keep the same order.
func FilterMap[T any, R any](collection []T, callback func(T, int) (R, bool)) []R {
	result := []R{}
	ch := make(chan R, len(collection))

	var wg sync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func(_item T, _i int) {
			if r, ok := callback(_item, _i); ok {
				ch <- r
			}

			wg.Done()
		}(item, i)
	}

	go func() {
		wg.Wait()

		close(ch)
	}()

	for r := range ch {
		result = append(result, r)
	}

	return result
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is call in parallel.
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
// `iteratee` is call in parallel.
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
// `iteratee` is call in parallel.
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
