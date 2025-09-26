//go:build go1.23

package it

import (
	"iter"

	"github.com/samber/lo"
)

// ToSeqPtr returns a sequence of pointers to each value.
func ToSeqPtr[T any](collection iter.Seq[T]) iter.Seq[*T] {
	return Map(collection, lo.ToPtr)
}

// FromSeqPtr returns a sequence with the pointer values.
// Returns a zero value in case of a nil pointer element.
func FromSeqPtr[T any](collection iter.Seq[*T]) iter.Seq[T] {
	return Map(collection, lo.FromPtr)
}

// FromSeqPtrOr returns a sequence with the pointer values or the fallback value.
func FromSeqPtrOr[T any](collection iter.Seq[*T], fallback T) iter.Seq[T] {
	return Map(collection, func(x *T) T { return lo.FromPtrOr(x, fallback) })
}

// ToAnySeq returns a sequence with all elements mapped to `any` type.
func ToAnySeq[T any](collection iter.Seq[T]) iter.Seq[any] {
	return func(yield func(any) bool) {
		for item := range collection {
			if !yield(item) {
				return
			}
		}
	}
}

// FromAnySeq returns a sequence with all elements mapped to a type.
// Panics on type conversion failure.
func FromAnySeq[T any](collection iter.Seq[any]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range collection {
			if t, ok := item.(T); !ok {
				panic("it.FromAnySeq: type conversion failed")
			} else if !yield(t) {
				return
			}
		}
	}
}

// Empty returns an empty sequence.
func Empty[T any]() iter.Seq[T] {
	return func(yield func(T) bool) {}
}

// IsEmpty returns true if the sequence is empty.
func IsEmpty[T any](collection iter.Seq[T]) bool {
	for range collection {
		return false
	}
	return true
}

// IsNotEmpty returns true if the sequence is not empty.
func IsNotEmpty[T any](collection iter.Seq[T]) bool {
	return !IsEmpty(collection)
}

// CoalesceSeq returns the first non-empty sequence.
func CoalesceSeq[T any](v ...iter.Seq[T]) (iter.Seq[T], bool) {
	for i := range v {
		for range v[i] {
			return v[i], true
		}
	}
	return func(yield func(T) bool) {}, false
}

// CoalesceSeqOrEmpty returns the first non-empty sequence.
func CoalesceSeqOrEmpty[T any](v ...iter.Seq[T]) iter.Seq[T] {
	for i := range v {
		for range v[i] {
			return v[i]
		}
	}
	return func(yield func(T) bool) {}
}
