package la

import (
	"github.com/samber/lo"
	"iter"
)

type capacityOpts struct {
	capacity int
}

type capacityOpt func(opts *capacityOpts)

// FromTuples creates a new iter.Seq2 from the [lo.Tuple2] where the first value
// will represent key and second – value.
func FromTuples[K, V any](res []lo.Tuple2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, vv := range res {
			if !yield(vv.A, vv.B) {
				return
			}
		}
	}
}

type ToMapOpt capacityOpt

// WithMapCapacity sets the capacity of an internal map which will be filled
// with the content of the iterator.
func WithMapCapacity(capacity int) ToMapOpt {
	return func(opts *capacityOpts) {
		opts.capacity = capacity
	}
}

// CollectMap creates a new map from the compatible iterator with keys and values.
// Since it will return a map, it might be costly to call that function because
// it will walk through the iterator.
//
// If the iterator has the same key more than one time – only the last value will
// be set into the returned map.
func CollectMap[K comparable, V any](in iter.Seq2[K, V], opts ...ToMapOpt) map[K]V {
	options := lo.Reduce(opts, func(agg *capacityOpts, item ToMapOpt, _ int) *capacityOpts {
		item(agg)

		return agg
	}, &capacityOpts{})

	out := make(map[K]V, options.capacity)
	for k, v := range in {
		out[k] = v
	}

	return out
}

// CollectToMap yield all values from the iterator to a pre-allocated slice.
//
// Might be useful if you know better how many items are in an iterator or want
// to reuse an already allocated map.
//
// Passed map will not be cleared before filling.
func CollectToMap[K comparable, V any, Map ~map[K]V](in iter.Seq2[K, V], out Map) Map {
	for k, v := range in {
		out[k] = v
	}

	return out
}

type ToSliceOpt capacityOpt

// WithSliceCapacity sets the capacity of an internal slice which will be filled
// with the content of the iterator.
func WithSliceCapacity(capacity int) ToSliceOpt {
	return func(opts *capacityOpts) {
		opts.capacity = capacity
	}
}

// Collect creates a slice that contains all values from the iterator.
//
// This func reflects the stdlib slices.Collect function but allows you to create
// a pre-allocated target slice to avoid many allocations.
func Collect[T any](in iter.Seq[T], opts ...ToSliceOpt) []T {
	options := lo.Reduce(opts, func(agg *capacityOpts, item ToSliceOpt, _ int) *capacityOpts {
		item(agg)

		return agg
	}, &capacityOpts{})

	out := make([]T, 0, options.capacity)
	for v := range in {
		out = append(out, v)
	}

	return out
}

// CollectTo yield all values from the iterator to a pre-allocated slice.
//
// Might be useful if you know better how many items are in an iterator or want
// to reuse an already allocated slice.
//
// Iterator will be appended to the passed slice.
func CollectTo[T any, Slice ~[]T](in iter.Seq[T], out Slice) Slice {
	for v := range in {
		out = append(out, v)
	}

	return out
}

// Seq2ToSeq transforms an iter.Seq2 into an iter.Seq based on a specific iteratee.
func Seq2ToSeq[K any, V any, R any](in iter.Seq2[K, V], iteratee func(key K, value V) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for k, v := range in {
			if !yield(iteratee(k, v)) {
				return
			}
		}
	}
}

// Entries transform an iter.Seq2 into iter.Seq of key/value pairs.
func Entries[K comparable, V any](in iter.Seq2[K, V]) iter.Seq[lo.Entry[K, V]] {
	return func(yield func(lo.Entry[K, V]) bool) {
		for k, v := range in {
			if !yield(lo.Entry[K, V]{Key: k, Value: v}) {
				return
			}
		}
	}
}

// FromEntries transforms an iter.Seq of key/value pairs into an iter.Seq2.
func FromEntries[K comparable, V any](entries iter.Seq[lo.Entry[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k := range entries {
			if !yield(k.Key, k.Value) {
				return
			}
		}
	}
}

// Pairs transform an iter.Seq2 into an iter.Seq of key/value pairs.
// Alias of Entries().
func Pairs[K comparable, V any](in iter.Seq2[K, V]) iter.Seq[lo.Entry[K, V]] {
	return Entries(in)
}

// Tuples transform an iter.Seq2 into an iter.Seq of [lo.Tuple2].
func Tuples[K comparable, V any](in iter.Seq2[K, V]) iter.Seq[lo.Tuple2[K, V]] {
	return func(yield func(tuple2 lo.Tuple2[K, V]) bool) {
		for k, v := range in {
			if !yield(lo.T2(k, v)) {
				return
			}
		}
	}
}

// FromPairs transform an iter.Seq2 into an iter.Seq of key/value pairs.
// Alias of FromEntries().
func FromPairs[K comparable, V any](entries iter.Seq[lo.Entry[K, V]]) iter.Seq2[K, V] {
	return FromEntries(entries)
}

// KeyBy transforms an iter.Seq to an iter.Seq2 based on a pivot callback.
func KeyBy[K any, V any](collection iter.Seq[V], iteratee func(item V) K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v := range collection {
			if !yield(iteratee(v), v) {
				return
			}
		}
	}
}

// SeqToSeq2 transforms an iter.Seq to an iter.Seq2 based on a pivot callback.
//
// Alias to: KeyBy
func SeqToSeq2[K any, V any](collection iter.Seq[V], iteratee func(item V) K) iter.Seq2[K, V] {
	return KeyBy(collection, iteratee)
}

// Associate returns an iter.Seq2 containing key-value pairs provided by the
// transform function applied to elements of the given iter.Seq.
//
// If any of two pairs have the same key, all of them will be returned in
// iter.Seq2 because iter.Seq2 can contain the same key more than once.
//
// The order of iteration will be the same as in the original iter.Seq.
func Associate[T any, K any, V any](collection iter.Seq[T], transform func(item T) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for v := range collection {
			if !yield(transform(v)) {
				return
			}
		}
	}
}

// SliceToMap returns an iter.Seq2 containing key-value pairs provided by the
// transform function applied to elements of the given iter.Seq.
//
// If any of two pairs have the same key, all of them will be returned in
// iter.Seq2 because iter.Seq2 can contain the same key more than once.
//
// The order of iteration will be the same as in the original iter.Seq.
// Alias of Associate().
func SliceToMap[T any, K any, V any](collection iter.Seq[T], transform func(item T) (K, V)) iter.Seq2[K, V] {
	return Associate(collection, transform)
}
