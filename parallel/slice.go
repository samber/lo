package parallel

import "sync"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	if len(collection) == 0 {
		return []R{}
	}

	if len(collection) == 1 {
		return []R{iteratee(collection[0], 0)}
	}

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
// `iteratee` is call in parallel.
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	if len(collection) == 0 {
		return
	}

	if len(collection) == 1 {
		iteratee(collection[0], 0)
		return
	}

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
func Times[T any](count int, iteratee func(index int) T) []T {
	if count == 0 {
		return []T{}
	}

	if count == 1 {
		return []T{iteratee(0)}
	}

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
func GroupBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) map[U]Slice {
	result := map[U]Slice{}

	if len(collection) == 0 {
		return result
	}

	if len(collection) == 1 {
		key := iteratee(collection[0])
		result[key] = collection
		return result
	}

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
func PartitionBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K) []Slice {
	if len(collection) == 0 {
		return []Slice{}
	}

	if len(collection) == 1 {
		return []Slice{collection}
	}

	result := make([]Slice, 0, 1)
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
