package la

import (
	"github.com/samber/lo"
	"iter"
)

// Filter iterates over elements of an iter.Seq, returning an iter.Seq, which
// yields only values where the predicate returns true.
func Filter[T any, Iter ~func(func(T) bool)](collection Iter, predicate func(item T) bool) Iter {
	return func(yield func(v T) bool) {
		for v := range collection {
			if !predicate(v) || yield(v) {
				continue
			}

			return
		}
	}
}

// Filter2 returns the same iter.Seq2 type filtered by the given predicate.
func Filter2[K any, V any, Map ~func(func(K, V) bool)](in Map, predicate func(key K, value V) bool) Map {
	return func(yield func(K, V) bool) {
		for k, v := range in {
			if !predicate(k, v) || yield(k, v) {
				continue
			}

			return
		}
	}
}

// FilterByKeys returns the same iter.Seq2 type filtered by given keys.
//
// Notice: semantic here differs from [lo.PickByKeys] because in the map you can
// effectively iterate over keys and pick only those that are requested, but with
// the iter.Seq2 this is not possible, and we should iterate over elements and
// throw away those who don't match.
func FilterByKeys[K comparable, V any, Map ~func(func(K, V) bool)](in Map, keys []K) Map {
	return func(yield func(K, V) bool) {
		for k, v := range in {
			if !lo.Contains(keys, k) || yield(k, v) {
				continue
			}

			return
		}
	}
}

// FilterByValues returns the same iter.Seq2 type filtered by given values.
func FilterByValues[K any, V comparable, Map ~func(func(K, V) bool)](in Map, values []V) Map {
	return func(yield func(K, V) bool) {
		for k, v := range in {
			if !lo.Contains(values, v) || yield(k, v) {
				continue
			}

			return
		}
	}
}

// Reject is the opposite of Filter, this method returns the elements of the
// iter.Seq that predicate doesn't return truth for.
func Reject[T any, Slice ~func(func(T) bool)](collection Slice, predicate func(item T) bool) Slice {
	return func(yield func(v T) bool) {
		for v := range collection {
			if predicate(v) || yield(v) {
				continue
			}

			return
		}
	}
}

// Reject2 is the opposite of Filter2, this method returns the elements of the
// iter.Seq that predicate doesn't return truth for.
func Reject2[K any, V any, Map ~func(func(K, V) bool)](in Map, predicate func(key K, value V) bool) Map {
	return func(yield func(K, V) bool) {
		for k, v := range in {
			if predicate(k, v) || yield(k, v) {
				continue
			}

			return
		}
	}
}

// RejectByKeys returns the same iter.Seq2 type filtered by given keys.
//
// Notice: semantic here differs from the [lo.OmitByKeys] because in the map you
// can effectively iterate over keys and drop those which should be omitted, but
// with the iter.Seq2 this is not possible, and we should iterate over elements
// and throw away those who don't match.
func RejectByKeys[K comparable, V any, Map ~func(func(K, V) bool)](in Map, keys []K) Map {
	return func(yield func(K, V) bool) {
		for k, v := range in {
			if lo.Contains(keys, k) || yield(k, v) {
				continue
			}

			return
		}
	}
}

// RejectByValues returns the same iter.Seq2 type filtered by given values.
//
// Notice: semantic here differs from the [lo.OmitByValues] because in the map you
// can effectively iterate over keys and drop those which should be omitted, but
// with the iter.Seq2 this is not possible, and we should iterate over elements
// and throw away those who don't match.
func RejectByValues[K any, V comparable, Map ~func(func(K, V) bool)](in Map, values []V) Map {
	return func(yield func(K, V) bool) {
		for k, v := range in {
			if lo.Contains(values, v) || yield(k, v) {
				continue
			}

			return
		}
	}
}

// RejectMap is the opposite of FilterMap, this method returns an iter.Seq which
// obtained after both filtering and mapping using the given callback function.
//
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func RejectMap[T any, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range collection {
			r, ok := callback(v)
			if ok || yield(r) {
				continue
			}

			return
		}
	}
}

// Uniq returns a duplicate-free version of an iter.Seq, in which only the first
// occurrence of each element is kept. The order of result values is determined
// by the order they occur in the iter.Seq.
//
// Be careful here because if you pass an endless iterator, you can create a huge
// state which tracks values for uniqueness.
func Uniq[T comparable, Iter ~func(func(T) bool)](collection Iter) Iter {
	return func(yield func(T) bool) {
		seen := map[T]struct{}{}

		for v := range collection {
			if _, exists := seen[v]; exists {
				continue
			}

			seen[v] = struct{}{}

			if !yield(v) {
				return
			}
		}
	}
}

// UniqBy returns a duplicate-free version of an iter.Seq, in which only the
// first occurrence of each element is kept. The order of result values is
// determined by the order they occur in the iter.Seq. It accepts `iteratee`
// which is invoked for each element in iter.Seq to generate the criterion by
// which uniqueness is computed.
//
// Be careful here because if you pass an endless iterator, you can create a huge
// state which tracks values for uniqueness.
func UniqBy[T any, U comparable, Iter ~func(func(T) bool)](collection Iter, iteratee func(item T) U) Iter {
	return func(yield func(T) bool) {
		seen := map[U]struct{}{}

		for v := range collection {
			c := iteratee(v)

			if _, exists := seen[c]; exists {
				continue
			}

			seen[c] = struct{}{}

			if !yield(v) {
				return
			}
		}
	}
}

// Replace returns an iter.Seq with the first n non-overlapping instances of the
// `old` replaced by the `new`.
//
// If n is less than 0, then all elements will be replaced.
func Replace[T comparable, Slice ~func(func(T) bool)](collection Slice, old T, new T, n int) Slice {
	return func(yield func(T) bool) {
		// When iterator iterated more than once and passed parameter captured by the
		// function, it might be unexpected that iterator created once but the iterator
		// function will be called each time when you iterate it. So if you capture the
		// parameter, you should avoid mutating it OR take an effort to restore the
		// original value.
		replaces := 0

		for v := range collection {
			if (n < 0 || n > replaces) && v == old {
				v = new
				n--
			}

			if !yield(v) {
				return
			}
		}
	}
}

// Replace2 return an iter.Seq2 with the first n non-overlapping instances of the
// `old` values replaced by the `new`. The keys remain the same.
//
// If n is less than 0, then all elements will be replaced.
func Replace2[K any, V comparable, Map ~func(func(K, V) bool)](collection Map, old V, new V, n int) Map {
	return func(yield func(K, V) bool) {
		// When iterator iterated more than once and passed parameter captured by the
		// function, it might be unexpected that iterator created once but the iterator
		// function will be called each time when you iterate it. So if you capture the
		// parameter, you should avoid mutating it OR take an effort to restore the
		// original value.
		replaces := 0

		for k, v := range collection {
			if (n < 0 || n > replaces) && v == old {
				v = new
				replaces++
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

// ReplaceAll return iter.Seq with all non-overlapping instances of the `old`
// replaced by the `new`
func ReplaceAll[T comparable, Slice ~func(func(T) bool)](collection Slice, old T, new T) Slice {
	return Replace(collection, old, new, -1)
}

// ReplaceAll2 return iter.Seq with all non-overlapping instances of the `old`
// values replaced by the `new`. The keys remain the same.
func ReplaceAll2[K any, V comparable, Map ~func(func(K, V) bool)](collection Map, old V, new V) Map {
	return Replace2(collection, old, new, -1)
}

// Compact returns an iter.Seq of all non-zero elements.
func Compact[T comparable, Slice ~func(func(T) bool)](collection Slice) Slice {
	var zero T

	return func(yield func(T) bool) {
		for v := range collection {
			if v == zero || yield(v) {
				continue
			}

			return
		}
	}
}

// Compact2 returns an iter.Seq2 of all non-zero elements.
func Compact2[K any, V comparable, Map ~func(func(K, V) bool)](collection Map) Map {
	var zero V

	return func(yield func(K, V) bool) {
		for k, v := range collection {
			if v == zero || yield(k, v) {
				continue
			}

			return
		}
	}
}
