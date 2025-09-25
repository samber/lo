//go:build go1.23

package it

import (
	"iter"

	"github.com/samber/lo"
)

// ToPtr returns a sequence of pointers to each value.
func ToPtr[T any](collection iter.Seq[T]) iter.Seq[*T] {
	return Map(collection, lo.ToPtr)
}

// FromPtr returns a sequence with the pointer values.
// Returns a zero value in case of a nil pointer element.
func FromPtr[T any](collection iter.Seq[*T]) iter.Seq[T] {
	return Map(collection, lo.FromPtr)
}

// FromPtrOr returns a sequence with the pointer values or the fallback value.
func FromPtrOr[T any](collection iter.Seq[*T], fallback T) iter.Seq[T] {
	return Map(collection, func(x *T) T { return lo.FromPtrOr(x, fallback) })
}

// ToAny returns a sequence with all elements mapped to `any` type.
func ToAny[T any](collection iter.Seq[T]) iter.Seq[any] {
	return func(yield func(any) bool) {
		for item := range collection {
			if !yield(item) {
				return
			}
		}
	}
}

// FromAny returns a sequence with all elements mapped to a type.
// Panics on type conversion failure.
func FromAny[T any](collection iter.Seq[any]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range collection {
			if t, ok := item.(T); !ok {
				panic("it.FromAny: type conversion failed")
			} else if !yield(t) {
				return
			}
		}
	}
}

// Coalesce returns the first non-empty sequence.
func Coalesce[T any](v ...iter.Seq[T]) (iter.Seq[T], bool) {
	for i := range v {
		for range v[i] {
			return v[i], true
		}
	}
	return func(yield func(T) bool) {}, false
}

// CoalesceOrEmpty returns the first non-empty sequence.
func CoalesceOrEmpty[T any](v ...iter.Seq[T]) iter.Seq[T] {
	for i := range v {
		for range v[i] {
			return v[i]
		}
	}
	return func(yield func(T) bool) {}
}
