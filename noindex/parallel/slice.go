package parallel

import lop "github.com/samber/lo/parallel"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(item T) R) []R {
	return lop.Map(collection, func(item T, _ int) R {
		return iteratee(item)
	})
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
// `iteratee` is call in parallel.
func ForEach[T any](collection []T, iteratee func(item T)) {
	lop.ForEach(collection, func(item T, _ int) {
		iteratee(item)
	})
}
