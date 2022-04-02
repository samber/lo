package parallel

import "sync"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(T, int) R, options ...*ParallelOption) []R {
	result := make([]R, len(collection))

	handler := func (item T, ix int) {
		result[ix] = iteratee(item, ix)
	}

	ForEach(collection, handler, options...)

	return result
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is call in parallel.
func ForEach[T any](collection []T, iteratee func(T, int), options ...*ParallelOption) {
	var wg sync.WaitGroup
	var concurrencyChn chan bool

	option := mergeOptions(options)
	if option.concurrencySetted {
		concurrencyChn = make(chan bool, option.concurrency)
	} else {
		concurrencyChn = make(chan bool, len(collection))
	}

	for i, item := range collection {
		wg.Add(1)
		concurrencyChn <- true
		go func(_item T, _i int) {
			iteratee(_item, _i)
			wg.Done()
			<- concurrencyChn
		}(item, i)
	}

	wg.Wait()
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
// `iteratee` is call in parallel.
func Times[T any](count int, iteratee func(int) T, options ...*ParallelOption) []T {
	var concurrencyChn chan bool

	option := mergeOptions(options)
	if option.concurrencySetted {
		concurrencyChn = make(chan bool, option.concurrency)
	} else {
		concurrencyChn = make(chan bool, count)
	}

	var wg sync.WaitGroup
	result := make([]T, count)
	for i := 0; i < count; i++ {
		wg.Add(1)
		concurrencyChn <- true
		go func(_i int) {
			defer func() {
				wg.Done()
				<- concurrencyChn
			}()
			item := iteratee(_i)
			result[_i] = item
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

			if _, ok := result[key]; !ok {
				result[key] = []T{}
			}

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
