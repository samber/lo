package la

import (
	"github.com/samber/lo"
	"iter"
	"slices"
)

// Map manipulates an iter.Seq and transforms it to an iter.Seq of another type.
func Map[T any, R any](collection iter.Seq[T], iteratee func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range collection {
			if yield(iteratee(v)) {
				continue
			}

			break
		}
	}
}

// Map2 manipulates an iter.Seq2 entries and transforms it to an iter.Seq2 of
// another type.
func Map2[K1 any, V1 any, K2 any, V2 any](in iter.Seq2[K1, V1], iteratee func(key K1, value V1) (K2, V2)) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		for k, v := range in {
			if k2, v2 := iteratee(k, v); !yield(k2, v2) {
				return
			}
		}
	}
}

// MapKeys manipulates an iter.Seq2 keys and transforms it to an iter.Seq2 of
// another type.
func MapKeys[K any, V any, R any](in iter.Seq2[K, V], iteratee func(value V, key K) R) iter.Seq2[R, V] {
	return func(yield func(R, V) bool) {
		for k, v := range in {
			if !yield(iteratee(v, k), v) {
				return
			}
		}
	}
}

// MapValues manipulates an iter.Seq2 values and transforms it to an iter.Seq2 of
// another type.
func MapValues[K any, V any, R any](in iter.Seq2[K, V], iteratee func(value V, key K) R) iter.Seq2[K, R] {
	return func(yield func(K, R) bool) {
		for k, v := range in {
			if !yield(k, iteratee(v, k)) {
				return
			}
		}
	}
}

// FilterMap returns an iter.Seq, which obtained after both filtering and mapping
// using the given callback function.
//
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap[T any, R any](collection iter.Seq[T], callback func(item T) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range collection {
			r, ok := callback(v)
			if !ok || yield(r) {
				continue
			}

			return
		}
	}
}

// FilterMap2 returns an iter.Seq2, which obtained after both filtering and mapping
// using the given callback function.
//
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
func FilterMap2[K any, V any, R any](collection iter.Seq2[K, V], callback func(K, V) (R, bool)) iter.Seq2[K, R] {
	return func(yield func(K, R) bool) {
		for k, v := range collection {
			r, ok := callback(k, v)
			if !ok || yield(k, r) {
				continue
			}

			return
		}
	}
}

// FlatMap manipulates an iter.Seq and transforms and flattens it to an iter.Seq of another type.
//
// The transform function can either return an iter.Seq or `nil`, and in the `nil` case
// no value is yielded to the final iter.Seq.
func FlatMap[T any, R any](collection iter.Seq[T], iteratee func(item T) iter.Seq[R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range collection {
			resIter := iteratee(v)
			if resIter == nil {
				continue
			}

			for vv := range resIter {
				if !yield(vv) {
					return
				}
			}
		}
	}
}

// FlatMapEnumerated manipulates an iter.Seq2 in the special case – when the key
// represents an index (like when your iter.Seq2 was produced by the Enumerate
// function). This function transforms and flattens it to an iter.Seq2 of another
// type and re-enumerates with proper indexes. All the original indexes will be
// ignored.
//
// The transform function can either return an iter.Seq2 or `nil`, and in the `nil` case
// no value is yielded to the final iter.Seq2.
func FlatMapEnumerated[T any, R any](collection iter.Seq2[int, T], iteratee func(int, T) iter.Seq[R]) iter.Seq2[int, R] {
	return func(yield func(int, R) bool) {
		idx := 0

		for _, v := range collection {
			resIter := iteratee(idx, v)
			if resIter == nil {
				continue
			}

			for vv := range resIter {
				if !yield(idx, vv) {
					return
				}

				idx++
			}
		}
	}
}

// Chunk returns an iter.Seq of elements split into slices the length of size. If
// iter.Seq can't be split evenly, the final chunk will be the remaining
// elements.
//
// Iterator will be collected lazily, not more than one chunk at once. Also,
// internally will be allocated chunk-sized slice to store temporary items
// collected from the chunk.
func Chunk[T any](collection iter.Seq[T], size int) iter.Seq[[]T] {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}

	return func(yield func([]T) bool) {
		chunk := make([]T, 0, size)

		for v := range collection {
			chunk = append(chunk, v)

			if len(chunk) == cap(chunk) {
				// the chunk here will be reused so it can't be yielded, so it should be copied
				res := make([]T, size)

				copy(res, chunk)
				chunk = chunk[0:0]

				if !yield(res) {
					return
				}
			}
		}

		if len(chunk) != 0 {
			yield(chunk)
		}
	}
}

// Chunk2 returns an iter.Seq of elements split into iter.Seq2 the length of the
// `size`. If iter.Seq2 can't be split evenly, the final chunk will be the
// remaining elements.
//
// Iterator will be collected lazily, not more than one chunk at once. Also,
// internally will be allocated chunk-sized slice to store temporary items
// collected from the chunk.
func Chunk2[K, V any, Map ~func(func(K, V) bool)](collection Map, size int) iter.Seq[Map] {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}

	return func(yield func(Map) bool) {
		chunk := make([]lo.Tuple2[K, V], 0, size)

		for k, v := range collection {
			chunk = append(chunk, lo.T2(k, v))

			if len(chunk) == cap(chunk) {
				// the chunk here will be reused so it can't be yielded, so it should be copied
				res := make([]lo.Tuple2[K, V], size)

				copy(res, chunk)
				chunk = chunk[0:0]

				if !yield((Map)(FromTuples(res))) {
					return
				}
			}
		}

		if len(chunk) != 0 {
			yield((Map)(FromTuples(chunk)))
		}
	}
}

// FlattenSlice returns an iter.Seq a single level deep.
func FlattenSlice[T any](collection iter.Seq[[]T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for sl := range collection {
			for _, v := range sl {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Flatten returns an iter.Seq a single level deep.
func Flatten[T any](collection iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for sl := range collection {
			for v := range sl {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Interleave round-robin alternating input slice of an iter.Seq and sequentially
// yield them in a result iterator.
func Interleave[T any, Iter ~func(func(T) bool)](collections ...Iter) Iter {
	return func(yield func(T) bool) {
		remains := len(collections)
		if remains == 0 {
			return
		}

		iters := make([]func() (T, bool), remains)
		stops := make([]func(), remains)

		defer func() {
			for _, stop := range stops {
				if stop != nil {
					stop()
				}
			}
		}()

		// convert to pull style iterators and sets stop functions to end their
		// iterations if the returned iterator stops early.
		for idx, it := range collections {
			iters[idx], stops[idx] = iter.Pull(iter.Seq[T](it))
		}

		for i := 0; ; i = (i + 1) % remains {
			v, valid := iters[i]()
			if !valid {
				// if the iterator exhausted – remove it from the iterators list
				// and try next to get value
				iters = slices.Delete(iters, i, i+1)
				remains--
				// maintain the current position because we removed one of iters from
				// it but yield nothing.
				i--

				if remains > 0 {
					continue
				}

				break
			}

			// if value can't be yielded – there is no need to continue iterations, and it is
			// necessary to stop all iterations via deferred function and return.
			if !yield(v) {
				return
			}
		}
	}
}

// Interleave2 round-robin alternating input slice of an iter.Seq2 and sequentially
// yield them in a result iterator.
func Interleave2[K, V any, Map ~func(func(K, V) bool)](collections ...Map) Map {
	return func(yield func(K, V) bool) {
		remains := len(collections)
		if remains == 0 {
			return
		}

		iters := make([]func() (K, V, bool), remains)
		stops := make([]func(), remains)

		defer func() {
			for _, stop := range stops {
				if stop != nil {
					stop()
				}
			}
		}()

		// convert to pull style iterators and sets stop functions to end their
		// iterations if the returned iterator stops early.
		for idx, it := range collections {
			iters[idx], stops[idx] = iter.Pull2(iter.Seq2[K, V](it))
		}

		for i := 0; ; i = (i + 1) % remains {
			k, v, valid := iters[i]()
			if !valid {
				// if the iterator exhausted – remove it from the iterators list
				// and try next to get value
				iters = slices.Delete(iters, i, i+1)
				remains--
				// maintain the current position because we removed one of iters from
				// it but yield nothing.
				i--

				if remains > 0 {
					continue
				}

				break
			}

			// if value can't be yielded – there is no need to continue iterations, and it is
			// necessary to stop all iterations via deferred function and return.
			if !yield(k, v) {
				return
			}
		}
	}
}
