package parallel

import "sync"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
// You also can control the conurrency limit by optional ParallelOption to limit the maximum number of
// concurrent `iteratee` goroutines running at the same time, just like
// `parallel.Map(list, iteratee, parallel.Option().Concurrency(10))`.
func Map[T any, R any](collection []T, iteratee func(T, int) R, options ...*ParallelOption) []R {
	result := make([]R, len(collection))

	handler := func(item T, ix int) {
		result[ix] = iteratee(item, ix)
	}

	ForEach(collection, handler, options...)

	return result
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is call in parallel.
// You also can control the conurrency limit by optional ParallelOption to limit the maximum number of
// concurrent `iteratee` goroutines running at the same time, just like
// `parallel.ForEach(list, iteratee, parallel.Option().Concurrency(10))`.
func ForEach[T any](collection []T, iteratee func(T, int), options ...*ParallelOption) {
	var wg sync.WaitGroup
	var concurrencyLimiter chan bool

	option := mergeOptions(options)
	if option.concurrencySetted {
		concurrencyLimiter = make(chan bool, option.concurrency)
	}

	wg.Add(len(collection))

	for i, item := range collection {
		if concurrencyLimiter != nil {
			concurrencyLimiter <- true
		}
		go func(_item T, _i int) {
			iteratee(_item, _i)
			wg.Done()
			if concurrencyLimiter != nil {
				<-concurrencyLimiter
			}
		}(item, i)
	}

	wg.Wait()
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
// `iteratee` is call in parallel.
// You also can control the conurrency limit by optional ParallelOption to limit the maximum number of
// concurrent `iteratee` goroutines running at the same time, just like
// `parallel.Times(count, iteratee, parallel.Option().Concurrency(10))`.
func Times[T any](count int, iteratee func(int) T, options ...*ParallelOption) []T {
	var concurrencyLimiter chan bool

	option := mergeOptions(options)
	if option.concurrencySetted {
		concurrencyLimiter = make(chan bool, option.concurrency)
	}

	result := make([]T, count)

	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		if concurrencyLimiter != nil {
			concurrencyLimiter <- true
		}
		go func(_i int) {
			defer func() {
				wg.Done()
				if concurrencyLimiter != nil {
					<-concurrencyLimiter
				}
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
// You also can control the conurrency limit by optional ParallelOption to limit the maximum number of
// concurrent `iteratee` goroutines running at the same time, just like
// `parallel.GroupBy(list, iteratee, parallel.Option().Concurrency(10))`.
func GroupBy[T any, U comparable](collection []T, iteratee func(T) U, options ...*ParallelOption) map[U][]T {
	result := map[U][]T{}
	var mu sync.Mutex

	handler := func(item T, ix int) {
		key := iteratee(item)

		mu.Lock()

		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}
		result[key] = append(result[key], item)

		mu.Unlock()
	}

	ForEach(collection, handler, options...)

	return result
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// `iteratee` is call in parallel.
// You also can control the conurrency limit by optional ParallelOption to limit the maximum number of
// concurrent `iteratee` goroutines running at the same time, just like
// `parallel.PartitionBy(list, iteratee, parallel.Option().Concurrency(10))`.
func PartitionBy[T any, K comparable](collection []T, iteratee func(x T) K, options ...*ParallelOption) [][]T {
	result := [][]T{}

	seen := map[K]int{}
	var mu sync.Mutex

	handler := func(item T, ix int) {
		key := iteratee(item)

		mu.Lock()
		defer mu.Unlock()

		resultIndex, ok := seen[key]
		if !ok {
			resultIndex = len(result)
			seen[key] = resultIndex
			result = append(result, []T{})
		}

		result[resultIndex] = append(result[resultIndex], item)
	}

	ForEach(collection, handler, options...)

	return result
}
