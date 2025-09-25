//go:build go1.23

package it

import (
	"iter"
	"maps"
)

// Keys creates a sequence of the map keys.
func Keys[K comparable, V any](in ...map[K]V) iter.Seq[K] {
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

// UniqKeys creates a sequence of unique keys in the map.
func UniqKeys[K comparable, V any](in ...map[K]V) iter.Seq[K] {
	return func(yield func(K) bool) {
		size := 0
		for i := range in {
			size += len(in[i])
		}

		seen := make(map[K]struct{}, size)

		for i := range in {
			for k := range in[i] {
				if _, exists := seen[k]; exists {
					continue
				}
				if !yield(k) {
					return
				}
				seen[k] = struct{}{}
			}
		}
	}
}

// Values creates a sequence of the map values.
func Values[K comparable, V any](in ...map[K]V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := range in {
			for k := range in[i] {
				if !yield(in[i][k]) {
					return
				}
			}
		}
	}
}

// UniqValues creates a sequence of unique values in the map.
func UniqValues[K, V comparable](in ...map[K]V) iter.Seq[V] {
	return func(yield func(V) bool) {
		size := 0
		for i := range in {
			size += len(in[i])
		}

		seen := make(map[V]struct{}, size)

		for i := range in {
			for k := range in[i] {
				val := in[i][k]
				if _, exists := seen[val]; exists {
					continue
				}
				if !yield(val) {
					return
				}
				seen[val] = struct{}{}
			}
		}
	}
}

// Entries transforms a map into a sequence of key/value pairs.
func Entries[K comparable, V any](in ...map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, m := range in {
			for k, v := range m {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// ToPairs transforms a map into a sequence of key/value pairs.
// Alias of Entries().
func ToPairs[K comparable, V any](in ...map[K]V) iter.Seq2[K, V] {
	return Entries(in...)
}

// FromEntries transforms a sequence of key/value pairs into a map.
func FromEntries[K comparable, V any](entries ...iter.Seq2[K, V]) map[K]V {
	m := make(map[K]V)
	for _, e := range entries {
		maps.Insert(m, e)
	}
	return m
}

// FromPairs transforms a sequence of key/value pairs into a map.
// Alias of FromEntries().
func FromPairs[K comparable, V any](entries ...iter.Seq2[K, V]) map[K]V {
	return FromEntries(entries...)
}

// Invert creates a sequence composed of inverted keys and values.
func Invert[K, V comparable](in iter.Seq2[K, V]) iter.Seq2[V, K] {
	return func(yield func(V, K) bool) {
		for k, v := range in {
			if !yield(v, k) {
				return
			}
		}
	}
}

// Assign merges multiple sequences of maps from left to right.
func Assign[K comparable, V any, Map ~map[K]V](maps ...iter.Seq[Map]) Map {
	out := make(Map)

	for i := range maps {
		for item := range maps[i] {
			for k := range item {
				out[k] = item[k]
			}
		}
	}

	return out
}

// ChunkEntries splits a map into a sequence of elements in groups of length equal to its size. If the map cannot be split evenly,
// the final chunk will contain the remaining elements.
func ChunkEntries[K comparable, V any](m map[K]V, size int) iter.Seq[map[K]V] {
	if size <= 0 {
		panic("it.ChunkEntries: size must be greater than 0")
	}

	return func(yield func(map[K]V) bool) {
		result := make(map[K]V, size)
		for k, v := range m {
			result[k] = v
			if len(result) == size {
				if !yield(result) {
					return
				}
				result = make(map[K]V, size)
			}
		}
		if len(result) > 0 {
			yield(result)
		}
	}
}

// FromMap transforms a map into a sequence based on specified iteratee.
func FromMap[K comparable, V, R any](in map[K]V, iteratee func(key K, value V) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for k := range in {
			if !yield(iteratee(k, in[k])) {
				return
			}
		}
	}
}

// FilterFromMap transforms a map into a sequence based on specified iteratee.
// The iteratee returns a value and a boolean. If the boolean is true, the value is added to the result sequence.
// If the boolean is false, the value is not added to the result sequence.
// The order of the keys in the input map is not specified and the order of the keys in the output sequence is not guaranteed.
func FilterFromMap[K comparable, V, R any](in map[K]V, iteratee func(key K, value V) (R, bool)) iter.Seq[R] {
	return func(yield func(R) bool) {
		for k := range in {
			if v, ok := iteratee(k, in[k]); ok && !yield(v) {
				return
			}
		}
	}
}

// FilterKeys transforms a map into a sequence based on predicate returns true for specific elements.
// It is a mix of Filter and Keys.
func FilterKeys[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range in {
			if predicate(k, in[k]) && !yield(k) {
				return
			}
		}
	}
}

// FilterValues transforms a map into a sequence based on predicate returns true for specific elements.
// It is a mix of Filter and Values.
func FilterValues[K comparable, V any](in map[K]V, predicate func(key K, value V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for k := range in {
			if predicate(k, in[k]) && !yield(in[k]) {
				return
			}
		}
	}
}
