package la

import (
	"iter"
)

// Enumerate an iter.Seq to get an iter.Seq2 that yields index and values.
func Enumerate[T any](collection iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		idx := 0

		for v := range collection {
			if !yield(idx, v) {
				return
			}

			idx++
		}
	}
}

// Keys create an iter.Seq of the iter keys.
//
// Notice: you should pay attention to keys semantic which differs from the
// original [lo.Keys] function from the [lo] package. The original one guarantees
// that keys will be unique across at least one map that passed here. An
// iter.Seq2 object doesn't have such a semantic and can contain one key multiple
// times regardless of the number of passed objects.
//
// Might be not the best name here but even in iter.Seq2 type parameters called
// K and V that assume some echo of the key-value semantic.
func Keys[K any, V any](in ...iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for i := range in {
			for k := range in[i] {
				if !yield(k) {
					return
				}
			}
		}
	}
}

// UniqKeys create an iter.Seq of unique keys in the map.
//
// Notice: since iter.Seq2 doesn't guarantee uniqueness of the keys that it
// yields, you might want to use this function instead of Keys in many cases.
// Also, be careful here because if you pass an endless iterator, you can create
// a huge state which tracks keys for uniqueness.
func UniqKeys[K comparable, V any](in ...iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		seen := make(map[K]struct{})

		for i := range in {
			for k := range in[i] {
				if _, exists := seen[k]; exists {
					continue
				}

				seen[k] = struct{}{}

				if !yield(k) {
					return
				}
			}
		}
	}
}

// Values create an iter.Seq of an iter.Seq2 values.
//
// Notice: returned values might have the same keys because iter.Seq2 doesn't
// require keys to be uniq here.
func Values[K any, V any](in ...iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := range in {
			for _, v := range in[i] {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// UniqValues create an iter.Seq of unique values in the iter.Seq2.
//
// Be careful here because if you pass an endless iterator, you can create a huge
// state which tracks values for uniqueness.
func UniqValues[K any, V comparable](in ...iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		seen := make(map[V]struct{})

		for i := range in {
			for _, v := range in[i] {
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
}

// Invert creates an iter.Seq2 composed of the inverted keys and values.
//
// Notice: duplicates handling semantic differs from [lo.Invert] here. An
// iter.Seq2 can contain duplicates, and they will remain after inversion.
func Invert[K any, V any](in iter.Seq2[K, V]) iter.Seq2[V, K] {
	return func(yield func(V, K) bool) {
		for k, v := range in {
			if !yield(v, k) {
				return
			}
		}
	}
}

// Join2 merges multiple iter.Seq2 from left to right.
//
// Notice: resulting iter.Seq2 will contain all key-value pairs regardless of
// their uniqueness because in contrast with the map type, the iter.Seq2 can
// contain many identical keys.
//
// Pay special attention to enumerated iter.Seq2 (this one which you receive from
// the Enumerate function or something similar to them when you have an integer
// value as a key). You may mistakenly expect that their indexes will be added to
// each other, but actually this is not true, and you receive the same index
// multiple times. To get this behavior use the [Follow] function instead.
func Join2[K any, V any, Map ~func(func(K, V) bool)](maps ...Map) Map {
	return func(yield func(K, V) bool) {
		for i := range maps {
			for k, v := range maps[i] {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Follow intended to merge multiple enumerated iter.Seq2 with int keys (like the
// result of the Enumerate function) and re-enumerate them at the same time.
//
// In contrast to Join2 the Follow will ignore the original index value from the
// iter and instead provide an ascending index like if you concatenate the values
// of each iter.Seq2 and then call the Enumerate on the result.
func Follow[V any, Map ~func(func(int, V) bool)](maps ...Map) Map {
	return func(yield func(int, V) bool) {
		idx := 0
		for i := range maps {
			for _, v := range maps[i] {
				if !yield(idx, v) {
					return
				}

				idx++
			}
		}
	}
}
