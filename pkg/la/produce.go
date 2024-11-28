package la

import (
	"github.com/samber/lo"
	"iter"
)

// Times invoke the iteratee n-times lazily, returning an iter.Seq yields each
// invocation by request. Keep in mind that if you re-start iteration of the
// iterator, you will initiate a new sequence and indexes will be re-started from
// zero.
func Times[T any](count int, iteratee func(index int) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range count {
			if !yield(iteratee(i)) {
				return
			}
		}
	}
}

// Times2 invoke the iteratee func n-times lazily, returning an iter.Seq2 where
// the key part will be set to index of invocation and the value is a result of
// iteratee invocation.
//
// Keep in mind that if you re-start iteration of the iterator, you will initiate
// a new sequence and indexes will be re-started from zero.
func Times2[T any](count int, iteratee func(index int) T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := range count {
			if !yield(i, iteratee(i)) {
				return
			}
		}
	}
}

// Repeat builds an iter.Seq with N copies of initial value.
func Repeat[T lo.Clonable[T]](count int, initial T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for range count {
			if !yield(initial.Clone()) {
				return
			}
		}
	}
}

// RepeatBy builds an iter.Seq with values returned by N calls of callback.
func RepeatBy[T any](count int, predicate func(index int) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range count {
			if !yield(predicate(i)) {
				return
			}
		}
	}
}
