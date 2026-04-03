package parallel

import "sync"

// Map manipulates a slice and transforms it to a slice of another type.
// `transform` is called in parallel. Result keep the same order.
// An optional WithConcurrency option can be provided to limit the number
// of goroutines running at the same time. When set, only that many worker
// goroutines are started (a bounded worker pool), rather than one per item.
// Play: https://go.dev/play/p/sCJaB3quRMC
func Map[T, R any](collection []T, transform func(item T, index int) R, opts ...Option) []R {
	result := make([]R, len(collection))
	forEach(collection, func(item T, i int) {
		result[i] = transform(item, i)
	}, buildOptions(opts))
	return result
}

// MapErr manipulates a slice and transforms it to a slice of another type.
// `transform` is called in parallel. Result keep the same order.
// Returns the first error encountered and stops processing further items.
// When WithConcurrency is set, a bounded worker pool is used instead of
// one goroutine per item. Supports WithContext for cancellation.
func MapErr[T, R any](collection []T, transform func(item T, index int) (R, error), opts ...Option) ([]R, error) {
	result := make([]R, len(collection))
	err := forEachErr(collection, func(item T, i int) error {
		r, err := transform(item, i)
		if err != nil {
			return err
		}
		result[i] = r
		return nil
	}, buildOptions(opts))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ForEach iterates over elements of collection and invokes callback for each element.
// `iteratee` is called in parallel.
// An optional WithConcurrency option can be provided to limit the number
// of goroutines running at the same time. When set, only that many worker
// goroutines are started (a bounded worker pool), rather than one per item.
// Play: https://go.dev/play/p/sCJaB3quRMC
func ForEach[T any](collection []T, callback func(item T, index int), opts ...Option) {
	forEach(collection, callback, buildOptions(opts))
}

// ForEachErr iterates over elements of collection and invokes callback for each element.
// `callback` is called in parallel.
// Returns the first error encountered and stops processing further items.
// When WithConcurrency is set, a bounded worker pool is used instead of
// one goroutine per item. Supports WithContext for cancellation.
func ForEachErr[T any](collection []T, callback func(item T, index int) error, opts ...Option) error {
	return forEachErr(collection, callback, buildOptions(opts))
}

// Times invokes the iteratee n times, returning a slice of the results of each invocation.
// The iteratee is invoked with index as argument.
// `iteratee` is called in parallel.
// An optional WithConcurrency option can be provided to limit the number
// of goroutines running at the same time. When set, only that many worker
// goroutines are started (a bounded worker pool), rather than one per item.
// Play: https://go.dev/play/p/ZNnWNcJ4Au-
func Times[T any](count int, iteratee func(index int) T, opts ...Option) []T {
	result := make([]T, count)
	_ = runErr(count, func(i int) error {
		result[i] = iteratee(i)
		return nil
	}, buildOptions(opts))
	return result
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
// The order of grouped values is determined by the order they occur in the collection.
// `iteratee` is called in parallel.
// An optional WithConcurrency option can be provided to limit the number
// of goroutines running at the same time.
// Play: https://go.dev/play/p/EkyvA0gw4dj
func GroupBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U, opts ...Option) map[U]Slice {
	result := map[U]Slice{}

	keys := Map(collection, func(item T, _ int) U {
		return iteratee(item)
	}, opts...)

	for i, item := range collection {
		result[keys[i]] = append(result[keys[i]], item)
	}

	return result
}

// PartitionBy returns a slice of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
// The order of groups is determined by their first appearance in the collection.
// `iteratee` is called in parallel.
// An optional WithConcurrency option can be provided to limit the number
// of goroutines running at the same time.
// Play: https://go.dev/play/p/GwBQdMgx2nC
func PartitionBy[T any, K comparable, Slice ~[]T](collection Slice, iteratee func(item T) K, opts ...Option) []Slice {
	result := []Slice{}
	seen := map[K]int{}

	keys := Map(collection, func(item T, _ int) K {
		return iteratee(item)
	}, opts...)

	for i := range collection {
		resultIndex, ok := seen[keys[i]]
		if ok {
			result[resultIndex] = append(result[resultIndex], collection[i])
		} else {
			seen[keys[i]] = len(result)
			result = append(result, Slice{collection[i]})
		}
	}

	return result
}

// forEach executes fn for each element in collection, in parallel.
func forEach[T any](collection []T, fn func(T, int), o options) {
	_ = forEachErr(collection, func(item T, i int) error {
		fn(item, i)
		return nil
	}, o)
}

// forEachErr executes fn for each element in collection, in parallel, with error handling.
func forEachErr[T any](collection []T, fn func(T, int) error, o options) error {
	return runErr(len(collection), func(i int) error {
		return fn(collection[i], i)
	}, o)
}

// runErr is the core parallel executor. It runs fn for each index 0..n-1 using a
// bounded worker pool. Returns the first error; stops scheduling on error or context cancellation.
func runErr(n int, fn func(int) error, o options) error {
	if n == 0 {
		return nil
	}
	concurrency := o.concurrency
	if concurrency <= 0 {
		concurrency = n
	}

	workers := minInt(concurrency, n)
	work := make(chan int)

	// First error wins: buffer of 1 means only the first push succeeds;
	// subsequent errors hit the default branch and are discarded.
	errCh := make(chan error, 1)

	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := range work {
				if err := fn(i); err != nil {
					select {
					case errCh <- err:
					default:
					}
				}
			}
		}()
	}

	// A nil channel blocks forever in select, so if no context was provided
	// the ctxDone case is effectively disabled.
	var ctxDone <-chan struct{}
	if o.ctx != nil {
		ctxDone = o.ctx.Done()
	}

	// Dispatch work. The select races work sends against errors and context
	// cancellation, so we stop dispatching as soon as either fires.
	for i := 0; i < n; i++ {
		select {
		case work <- i:
		case err := <-errCh:
			close(work)
			wg.Wait()
			return err
		case <-ctxDone:
			close(work)
			wg.Wait()
			return o.ctx.Err()
		}
	}

	// All items dispatched. Wait for in-flight workers, then check if any
	// of them produced an error we haven't seen yet.
	close(work)
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
