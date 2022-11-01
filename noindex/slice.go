package noindex

import (
	"github.com/samber/lo"
)

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[V any](collection []V, predicate func(item V) bool) []V {
	return lo.Filter(collection, func(item V, _ int) bool {
		return predicate(item)
	})
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T) R) []R {
	return lo.Map(collection, func(item T, _ int) R {
		return iteratee(item)
	})
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T any, R any](collection []T, callback func(item T) (R, bool)) []R {
	return lo.FilterMap(collection, func(item T, _ int) (R, bool) {
		return callback(item)
	})
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
func FlatMap[T any, R any](collection []T, iteratee func(item T) []R) []R {
	return lo.FlatMap(collection, func(item T, _ int) []R {
		return iteratee(item)
	})
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T) R, initial R) R {
	return lo.Reduce(collection, func(agg R, item T, _ int) R {
		return accumulator(agg, item)
	}, initial)
}

// ReduceRight helper is like Reduce except that it iterates over elements of collection from right to left.
func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T) R, initial R) R {
	return lo.ReduceRight(collection, func(agg R, item T, _ int) R {
		return accumulator(agg, item)
	}, initial)
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach[T any](collection []T, iteratee func(item T)) {
	lo.ForEach(collection, func(item T, _ int) {
		iteratee(item)
	})
}

// Reject is the opposite of Filter, this method returns the elements of collection that predicate does not return truthy for.
func Reject[V any](collection []V, predicate func(item V) bool) []V {
	return lo.Reject(collection, func(item V, _ int) bool {
		return predicate(item)
	})
}
