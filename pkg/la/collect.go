package la

import "iter"

// Reduce reduces an iter.Seq to a value which is the accumulated result of
// running each element in the iter.Seq through accumulator, where each
// successive invocation is supplied with the return value of the previous.
//
// Since this function requires collecting all values from the iterator – in many
// cases, it will be better to use [lo.Reduce] after [slices.Collect] on the
// iterator instead of this function.
func Reduce[T any, R any](collection iter.Seq[T], accumulator func(agg R, item T) R, initial R) R {
	for v := range collection {
		initial = accumulator(initial, v)
	}

	return initial
}

// Reduce2 reduces an iter.Seq2 to a value which is the accumulated result of
// running each element in the iter.Seq2 through accumulator, where each
// successive invocation is supplied with the return value of the previous.
//
// Since this function requires collecting all values from the iterator – in many
// cases, it will be better to use [lo.Reduce] after [Tuples] and
// [slices.Collect] on the iterator instead of this function.
func Reduce2[K, V any, R any](collection iter.Seq2[K, V], accumulator func(agg R, key K, val V) R, initial R) R {
	for k, v := range collection {
		initial = accumulator(initial, k, v)
	}

	return initial
}

// ForEach iterates over elements of an iter.Seq and invokes iteratee for each element.
func ForEach[T any](collection iter.Seq[T], iteratee func(item T)) {
	for v := range collection {
		iteratee(v)
	}
}

// ForEach2 iterates over elements of an iter.Seq2 and invokes iteratee for each element.
func ForEach2[K, V any](collection iter.Seq2[K, V], iteratee func(key K, val V)) {
	for k, v := range collection {
		iteratee(k, v)
	}
}

// ForEachWhile iterates over elements of an iter.Seq and invokes iteratee for
// each element, the returned value decides to continue or break, like do
// while().
func ForEachWhile[T any](collection iter.Seq[T], iteratee func(item T) (goon bool)) {
	for v := range collection {
		if !iteratee(v) {
			break
		}
	}
}

// ForEachWhile2 iterates over elements of an iter.Seq and invokes iteratee for
// each element, the returned value decides to continue or break, like do
// while().
func ForEachWhile2[K, V any](collection iter.Seq2[K, V], iteratee func(key K, val V) (goon bool)) {
	for k, v := range collection {
		if !iteratee(k, v) {
			break
		}
	}
}

// KeyValues create two parallel iterators where the first yields keys and the
// second – values of the original iterator.
//
// To achieve that, it is necessary to walk through an original iterator twice,
// so if your iterator is not support that – avoid this function.
func KeyValues[K any, V any](in iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V]) {
	return func(yield func(K) bool) {
			for k := range in {
				if !yield(k) {
					return
				}
			}
		},
		func(yield func(V) bool) {
			for _, v := range in {
				if !yield(v) {
					return
				}
			}
		}
}
